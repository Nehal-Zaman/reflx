package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {

	urlfile, threads := GetUrlListFileName()
	var urlList = make([]string, 0)

	if urlfile == "" {
		urlList = GetUrlsFromStdin()
	} else {
		raw_contents, err := os.ReadFile(urlfile)
		checkErr(err)

		urlList = strings.Split(string(raw_contents), "\n")
	}

	urlChan := make(chan string)

	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for urlInput := range urlChan {
				CheckReflection(urlInput)
			}
		}()
	}

	for _, urlInput := range urlList {
		urlChan <- urlInput
	}

	close(urlChan)
	wg.Wait()
}

func CheckReflection(urlInput string) {
	urlParsed, err := url.Parse(urlInput)
	if err != nil {
		fmt.Println(err)
		return
	}

	new_query_parts := make([]string, 0)
	if urlParsed.RawQuery != "" {
		params := ParseRawQuery(urlParsed.RawQuery)
		var paramMap = make(map[string]string)

		for _, param := range params {
			if param != "" {
				random_string := getRandomParamValue()
				paramMap[param] = random_string
				new_query_parts = append(new_query_parts, param+"="+random_string)
			}
		}

		new_url := urlParsed.Scheme + "://" + urlParsed.Host + urlParsed.Path + "?" + strings.Join(new_query_parts, "&")
		if urlParsed.Fragment != "" {
			new_url += "#" + urlParsed.Fragment
		}

		resp_contents := MakeHTTPRequest(new_url)
		for k := range paramMap {
			if strings.Contains(resp_contents, paramMap[k]) {
				fmt.Printf("Parameter '%v' is reflected in %v\n", k, new_url)
			}
		}
	}
}

func MakeHTTPRequest(urlInput string) string {
	resp, err := http.Get(urlInput)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	checkErr(err)

	return string(body)

}

func getRandomParamValue() string {
	var characters string = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz012345678"
	random_string := ""

	for i := 0; i < 32; i++ {
		random_string += string(characters[rand.Intn(len(characters))])
	}

	return random_string
}

func ParseRawQuery(raw_query string) []string {
	var params = make([]string, 0)
	ampersand_split := strings.Split(raw_query, "&")

	for _, raw_query_part := range ampersand_split {
		params = append(params, strings.TrimSpace(strings.Split(raw_query_part, "=")[0]))
	}

	return params
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GetUrlListFileName() (string, int) {
	urlfilePtr := flag.String("list", "", "specify a file containing URL list")
	threadsPtr := flag.Int("threads", 10, "specify number of threads to use")
	flag.Parse()

	return *urlfilePtr, *threadsPtr
}

func GetUrlsFromStdin() []string {
	var urlList = make([]string, 0)

	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		urlList = append(urlList, scanner.Text())
	}

	return urlList
}
