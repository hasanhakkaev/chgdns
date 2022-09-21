# chgdns
[![CI](https://github.com/hasanhakkaev/chgdns/workflows/CI/badge.svg?event=push)](https://github.com/hasanhakkaev/chgdns/actions?query=workflow%3ACI)
[![Go Report Card](https://goreportcard.com/badge/github.com/hasanhakkaev/chgdns)](https://goreportcard.com/report/github.com/hasanhakkaev/chgdns)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=hasanhakkaev_chgdns&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=hasanhakkaev_chgdns)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=hasanhakkaev_chgdns&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=hasanhakkaev_chgdns)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=hasanhakkaev_chgdns&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=hasanhakkaev_chgdns)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=hasanhakkaev_chgdns&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=hasanhakkaev_chgdns)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=hasanhakkaev_chgdns&metric=coverage)](https://sonarcloud.io/summary/new_code?id=hasanhakkaev_chgdns)
[![Release](https://img.shields.io/github/release/hasanhakkaev/chgdns.svg)](https://github.com/hasanhakkaev/chgdns/releases/latest)
[![Go version](https://img.shields.io/github/go-mod/go-version/hasanhakkaev/chgdns)](https://github.com/hasanhakkaev/chgdns)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This tool lets you easily change the DNS settings of your Windows.

Install[WIP]
-------
### Choco package manager[WIP]
If you use the [Chocolatey package manager](https://chocolatey.org/) you can install `chgdns` with:

```
choco install chgdns
```
### Binary download[WIP]
You can manually download and install the binary from [Releases](https://github.com/hasanhakkaev/chgdns/releases) page.

####

## Usage[WIP]
You can simply run binary by providing required command line arguments:
```
chgdns.exe -h
```

####
## Development[WIP]
This project requires below tools while developing:
- [Golang 1.17](https://golang.org/doc/go1.17)
- [pre-commit](https://pre-commit.com/)
- [golangci-lint](https://golangci-lint.run/usage/install/) - required by [pre-commit](https://pre-commit.com/)
- [gocyclo](https://github.com/fzipp/gocyclo) - required by [pre-commit](https://pre-commit.com/)

After you installed [pre-commit](https://pre-commit.com/), simply run below command to prepare your development environment:
```shell
$ pre-commit install
```
