# Learn some design pattern using Golang

## Set up database
- Go in models/bdd and add your config
- You can create default db called wiki or change NAME_DB constant
- That's it because Gorm will create database schema

## Set up linter with vscode
- Open config vscode file 
    ```sh
    $ sudo nano $HOME/.config/Code/User/settings.json
    ```
- Add properties to config vscode file after first ```{```
    ```json
    "go.lintTool":"golangci-lint",
    "go.lintFlags": [
        "--fast"
    ],
    ```
- Thank's to ```.golangci.yml``` file there is all tools you needed

## Launch linter
```sh
$ golangci-lint run
```

## Launch project
```sh
$ go run main.go
```

[LICENSE](https://github.com/wyllisMonteiro/mvc_go/blob/master/LICENSE)