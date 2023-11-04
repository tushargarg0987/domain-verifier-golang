# Domain Name Verifier (Go)

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Contributing](#contributing)
- [License](#license)

## Introduction

The Domain Name Verifier is a command-line tool written in the Go programming language that helps you verify the validity and ownership of domain names. It can be used for various purposes, including domain management, domain portfolio verification, and security checks.

This tool uses DNS and WHOIS queries to obtain information about a domain name, helping you determine if a domain is active, who it belongs to, and other relevant details.

## Features

- Check the validity of a domain name.
- Retrieve DNS information, including A, AAAA, MX, and TXT records.
- Perform a WHOIS lookup to get domain ownership information.
- Determine if a domain is available for registration.

## Prerequisites

Before using the Domain Name Verifier, make sure you have the following prerequisites installed on your system:

- Go (Golang) 1.14 or higher
- Internet connectivity for DNS and WHOIS queries

## Installation

You can install the Domain Name Verifier by running the following Go command:

```bash
git clone https://github.com/tushargarg0987/domain-verifier-golang.git
```

## Usage

To use the Domain Name Verifier, run the following command in your terminal:

```bash
cd domain-verifier-golang
go run main.go
```

## Contributing

If you'd like to contribute to this project, please fork the repository and create a pull request. We welcome contributions and improvements from the community.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.