# Ani-Go \[BETA\]

Webserver in Golang for a website dedicated to Anime

### Compile:
```sh
go build -c AniGo.exe Ani.go
```

### Run witout compiling:
```sh
go run Ani.go
```
The server will run on **http://127.0.0.1:8770**

### Warning:
Database is required, the database can be built with these commands
```sh
go run Ani.go --init-db
# or
AniGo.exe --init-db
```
**_the data files are not included_** here on Github since the size is a little bit more than 50MB


## TODO
- Website
- Api
- Views on SQLite
