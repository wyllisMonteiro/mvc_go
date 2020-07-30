# Learn some design pattern using Golang

## Set up database
- Go in models/bdd and add your config
- You can create default db called wiki or change NAME_DB constant
- That's it !!! Because Gorm will create database schema

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
- Thank's to ```.golangci.yml``` file, there are all the necessary tools

## Set up unit tests with vscode
- Just create a file with _test.go name

Example :

For example, you want to create unit tests to a file called **auth.go**
Create a new file called **auth_test.go** at the same location than **auth.go**
In that file, add all functions needed to make tests with **Test** in prefix function name
like **TestLogin()**, **TestLogout()**.
For more details about writting tests check this link : https://medium.com/better-programming/easy-guide-to-unit-testing-in-golang-4fc1e9d96679

## Launch unit tests
```sh
$ go test ./...
```

## Launch linter
```sh
$ golangci-lint run
```

## Launch project
```sh
$ go run main.go
```

[LICENSE](https://github.com/wyllisMonteiro/mvc_go/blob/master/LICENSE)