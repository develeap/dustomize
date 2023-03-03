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

[link](./dustomize.yaml)

```yaml
# dustomize.yaml
options:
  displayValues: false # true = display appended values when parsing

import:
  fromFile:
    - example/values.yaml
    - example/values2.yaml
    - example/values3.yaml
  fromUrl:
    - https://run.mocky.io/v3/897c32b8-bf7c-40b4-ace6-2b6d5f68f6ac
    - https://run.mocky.io/v3/7456b035-f644-4ed2-b6c4-e777a9871d7d
    - https://run.mocky.io/v3/18543c68-c50b-464e-a130-86c2bf4574c7
  fromText: |
    Base: alpine

    command:
      copy: cp
      delete: rm
      add: touch

    lines:
      one: this is line number one
      two: this is line number two
      three: this is line number three
      four: this is line number four

    Attended: true

export:
  - template: example/templates/dockerfiles/Dockerfile
    target: example/outputs/dockerfiles/Dockerfile
    description: first example..

  - template: example/templates/dockerfiles/app1.Dockerfile
    target: example/outputs/dockerfiles/app1.Dockerfile
    description: second example..

  - template: example/templates/dockerfiles/app2.Dockerfile
    target: example/outputs/dockerfiles/app2.Dockerfile

  - template: example/templates/texts/TEST.MD
    target: example/outputs/texts/TEST.MD
```

```Dockerfile
# Dockerfile - before
FROM {{ .Base }}
RUN apt install {{ .git.packageName }}={{ .git.packageVersion }}
RUN echo {{ randAlphaNum 20 }}
```

```go
// Install the CLI
go install .

// Run the CLI
dustomize parse // reads local config
```

```Dockerfile
# Dockerfile - after
FROM alpine
RUN apt install git=1:2.9.3-1
RUN echo gN5mBamkiCzMTycytuwC
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
