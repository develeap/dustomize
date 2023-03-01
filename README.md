# dustomize - Dockerfile templates

<img src="https://d1.awsstatic.com/acs/characters/Logos/Docker-Logo_Horizontel_279x131.b8a5c41e56b77706656d61080f6a0217a3ba356d.png" width="300px">

This project is used to create powerful templates within Dockerfiles

## ⏏️ Options

```go
Usage of dustomize:
  -config string
        Config definition
  -export bool
        Enable to export the resulted files
  -folder string
        Folder to parse
  -output string
        Output folder
  -verbose bool
        Enable to view logs
```

## 📋 Example

```yaml
# values.yml
Base: alpine

command:
  copy: cp
  delete: rm
  add: touch
```

```Dockerfile
# Dockerfile
FROM {{ .Base }}
RUN ls -la
```

```go
// terminal command (root dir)
go run . -folder example/dockerfiles/templates/ -config example/values.yaml -export -output example/dockerfiles/outputs
```

## ℹ️️ Requirements

- [Go](https://golang.org/doc/install) 1.20.x (to build the project)

## 💁🏻 Contributing

This is an open source project. Any contribution would be greatly appreciated!

## 🚩 Issues

If you have found an issue, please report it on the [issue tracker](https://github.com/develeap/dustomize/issues)

## 📝 LICENSE

License is MIT

## Made with ❤️ in develeap

[<img src="https://media.licdn.com/dms/image/C4D0BAQFXwTP7SFX0QQ/company-logo_200_200/0/1583831070407?e=2147483647&v=beta&t=bWP52NuMxHiQyhMIEe9D7xTNcQMuQDbrTy-ZiVVLCv0" width="50px">](https://www.develeap.com/)
[<img src="https://upload.wikimedia.org/wikipedia/commons/8/81/LinkedIn_icon.svg" width="50px">](https://www.linkedin.com/company/develeap/mycompany/)
