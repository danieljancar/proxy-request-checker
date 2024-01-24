# Proxy Checker in Go

Check if your network or proxy are working as expected, by checking the status code of a request to a website.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/)
- [Git](https://git-scm.com/downloads)

## Configuration

The `config/links/links-config.json` file contains the proxies to be checked.

```json
[
  {
    "url": "https://xxx.xxx.xxx.xxx:xxxx",
    "expectedStatus": 200
  }
]
```

> **Note:** Make sure to create the `links-config.json` file, based on the `links-config.example.json` file, if it does
> not exist.

---

The `.env` file contains the configuration for the proxy checker. The format is as follows, make sure to create a `.env`
file, based on the `.env.example` file, if it does not exist.

> **Note:** Find more information about the configuration in the [docs/configuration.md](docs/CONFIGURATION.md) file.

## Usage

### Go Run

To run the proxy checker on the fly, run the following command:

```bash
/projects/go-proxy-request-checker >
go run cmd/main.go
```

### Binary

To run the proxy checker as a binary, run the following commands:

```bash
/projects/go-proxy-request-checker >
go build -o bin/proxy-checker cmd/main.go 
/projects/go-proxy-request-checker >
bin/proxy-checker
```

> **Note:** Run the binary on the project root directory to avoid any issues with the config files, that might not be
> found.

The report is saved to the `reports` folder by default.

# Known Issues

- [ ] If running the binary, the report is not saved to the specified file path, also the links.json file is not found,
  if the binary isn't run from the project root directory. #3

# Contributing

I welcome contributions from everyone. Before you get started, please take a moment to review
the [guidelines](.github/CONTRIBUTING.md).

## Branches

This project uses the following branches:

- `master` - The main branch. This branch is used for production releases.
- `develop` - The development branch. This branch is used for development and testing.

## Commit Messages

When you commit, make sure to follow the [Contributing Guide](.github/CONTRIBUTING.md) for commit messages. This will
help keep the commit messages clean and consistent across the project. If you don't follow the guide, your pull request
will be rejected.

# License

This project is licensed under the [MIT](LICENSE) license. By contributing to this project, you agree that your
contributions will be released under the same license. Also, you agree to
the [Contributor Covenant Code of Conduct](.github/CODE_OF_CONDUCT.md)
and [Developer Certificate of Origin](.github/DCO.md).

## Disclaimer

This project is for educational purposes only. I am not responsible for any misuse of this project
