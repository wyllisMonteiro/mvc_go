# warticle

# What is warticle
It's a student project for building an application that allows user to create articles. I used strategy design pattern with Golang and MVC architecture.

# Features
- Display all articles
- Display, create and edit article
- Download from browser articles in .csv or .xlsx

## Set up database
- Config database for Golang app:
    - Create a file called `/warticle/.env`
    - Set `USERDB` like this `USERDB=root`
    - Set `PASSDB` like this `PASSDB=root`
    - Set `NAMEDB` like this `NAMEDB=wiki`
- Config database in `docker-compose.yml`
    - Set `MYSQL_ROOT_PASSWORD, MYSQL_DATABASE, MYSQL_USER, MYSQL_PASSWORD`

At the begining there is no articles because db is empty

## Start project
Use `docker-compose up --build -d`

If you need to update Golang code you can execute `docker-compose up --build -d app`

Application at http://localhost:9000

Godoc at http://localhost:6060

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

## Set up unit tests
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

[LICENSE](https://github.com/wyllisMonteiro/mvc_go/blob/master/LICENSE)