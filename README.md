# dustomize - Dockerfile templates

![GitHub Actions CI](https://github.com/develeap/dustomize/actions/workflows/ci.yaml/badge.svg?branch=main)
![GitHub repo size](https://img.shields.io/github/repo-size/develeap/dustomize)
![Bitbucket open issues](https://img.shields.io/bitbucket/issues-raw/develeap/dustomize)
![GitHub all releases](https://img.shields.io/github/downloads/develeap/dustomize/total)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)

<img src="https://d1.awsstatic.com/acs/characters/Logos/Docker-Logo_Horizontel_279x131.b8a5c41e56b77706656d61080f6a0217a3ba356d.png" width="300px">

This project is used to create powerful templates within Dockerfiles

## ⏏️ Options

```go
A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  dustomize parse [flags]

Flags:
  -c, --config string   Values file to read from.
  -f, --file strings    Template files to parse.
  -k, --folder string   Your templates directory to parse.
  -h, --help            help for parse
  -o, --output string   Export parsing to target folder.
```

## 📋 Example

```yaml
# values.yml
Base: alpine

command:
  ls: ls -la
# ...
```

```Dockerfile
# Dockerfile
FROM {{ .Base }}
FROM {{ .command.ls }}
RUN the date today is: {{ now | htmlDate }}
```

```go
// Install the CLI
go install .

// Run the CLI
dustomize parse -c example/values.yaml -k example/dockerfiles/templates/
```

## ℹ️️ Requirements

- [Go](https://golang.org/doc/install) 1.20.x (to build the project)

## 💁🏻 Contributing

This is an open source project. Any contribution would be greatly appreciated!

## 🚩 Issues

If you have found an issue, please report it on the [issue tracker](https://github.com/develeap/dustomize/issues)

## Made with ❤️ in develeap

[<img src="https://media.licdn.com/dms/image/C4D0BAQFXwTP7SFX0QQ/company-logo_200_200/0/1583831070407?e=2147483647&v=beta&t=bWP52NuMxHiQyhMIEe9D7xTNcQMuQDbrTy-ZiVVLCv0" width="50px">](https://www.develeap.com/)
[<img src="https://upload.wikimedia.org/wikipedia/commons/8/81/LinkedIn_icon.svg" width="50px">](https://www.linkedin.com/company/develeap/mycompany/)
