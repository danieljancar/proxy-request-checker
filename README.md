# Proxy Checker in Go

This is a simple proxy checker written in Go. It checks the links in the `links.json` file and returns a report.

## Getting Started

### Prerequisites

* [Go](https://golang.org/dl/)
* [Git](https://git-scm.com/downloads)

### Installing

```bash
git clone https://github.com/danieljancar/go-proxy-request-checker.git
cd go-proxy-checker
```

### Configuration

The `config/links.json` file contains the proxies to be checked. The format is as follows, make sure to create a `links.json` file in the `config` folder if it does not exist.

```json
[
  {
    "url": "https://xxx.xxx.xxx.xxx:xxxx",
    "expectedStatus": 200
  },
  {
    "url": "https://xxx.xxx.xxx.xxx:xxxx",
    "expectedStatus": 200
  },
  {
    "url": "https://xxx.xxx.xxx.xxx:xxxx",
    "expectedStatus": 403
  }
]
```

## Usage

```bash
go run cmd/main.go
```

## Output

```bash
$ go run main.go
Checking 3 proxies...
Proxy 1:
    IP: xxx.xxx.xxx.xxx
    Port: xxxx
    Country: xxx
    Anonymity: xxx
    HTTPS: xxx
    Speed: xxx
    Status: xxx
Proxy 2:
    IP: xxx.xxx.xxx.xxx
    Port: xxxx
    Country: xxx
    Anonymity: xxx
    HTTPS: xxx
    Speed: xxx
    Status: xxx
Proxy 3:
    IP: xxx.xxx.xxx.xxx
    Port: xxxx
    Country: xxx
    Anonymity: xxx
    HTTPS: xxx
    Speed: xxx
    Status: xxx
    
Finished checking 3 proxies.
--------
Results:
    200: 2
    403: 1
    Failed: 0
    Total: 3

Do you want to save the results? (y/n): y

Enter the name of the file: results

Saved results to results.json 
```

# Contributing

I welcome contributions from everyone. Before you get started, please take a moment to review the [guidelines](.github/CONTRIBUTING.md).

## Branches

This project uses the following branches:

- `master` - The main branch. This branch is used for production releases.
- `develop` - The development branch. This branch is used for development and testing.

## Commit Messages

When you commit, make sure to follow the [Contributing Guide](.github/CONTRIBUTING.md) for commit messages. This will help keep the commit messages clean and consistent across the project. If you don't follow the guide, your pull request will be rejected.

# License

This project is licensed under the [MIT](LICENSE) license. By contributing to this project, you agree that your contributions will be released under the same license. Also, you agree to the [Contributor Covenant Code of Conduct](.github/CODE_OF_CONDUCT.md) and [Developer Certificate of Origin](.github/DCO.md).

## Disclaimer

This project is for educational purposes only. I am not responsible for any misuse of this project
