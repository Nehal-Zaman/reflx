# reflx - Discover Reflected Parameters in URLs

![GitHub License](https://img.shields.io/github/license/Nehal-Zaman/reflx)
![Go Version](https://img.shields.io/github/go-mod/go-version/Nehal-Zaman/reflx)

`reflx` is a powerful tool designed to help you identify reflected parameters in URLs. It automates the process of checking whether input parameters are being reflected in the response content of web URLs. By leveraging this tool, you can efficiently identify potential security vulnerabilities related to parameter reflection.

## Features

- Quickly discover reflected parameters in URLs.
- Multithreaded processing for faster analysis.
- Option to suppress error messages for silent operation.
- Easily read URLs from a file or provide them through stdin.

## Installation

To install `reflx`, make sure you have Go installed on your system, and then execute the following command:

```sh
go install github.com/Nehal-Zaman/reflx@latest
```

## Usage

```sh
reflx -h
```

### Flags

- `-list`: Specify a file containing a list of URLs to analyze.
- `-threads`: Specify the number of threads to use for processing (default is 10).
- `-silent`: Suppress error messages for silent operation.

### Examples

1. Analyze URLs from a file:

```sh
reflx -list url_list.txt
```

2. Analyze URLs from stdin:

```sh
cat url_list.txt | reflx
```

## How It Works

reflx operates by analyzing the parameters in the given URLs and checking if they are being reflected in the response content. It performs the following steps:

1. Reads the list of URLs either from a file or from stdin.
2. Parses each URL and extracts its parameters.
3. Constructs modified URLs by replacing parameter values with randomly generated strings.
4. Makes HTTP requests to the modified URLs and checks if the generated strings are reflected in the response content.
5. Displays findings regarding reflected parameters.

## Contributions

Contributions are welcome! If you find a bug or want to suggest an improvement, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

---

**Disclaimer:** This tool is designed for educational and security research purposes. Use it responsibly and only on web applications you have permission to test. The developers are not responsible for any unauthorized or illegal use of this tool.

For more information, please refer to the [GitHub repository](https://github.com/Nehal-Zaman/reflx).

**Author:** Nehal Zaman (n3hal_)
**Contact:** nehalcodes@gmail.com