# go-mysql-rest-api

A guide for creating RESTful API with Golang and MySQL and Gin.

## Build and Usage

```bash
go build
```

```bash
./mysql-rest-api
```

```bash
http://localhost:8989
```

## Config

modify conf.yaml

## Features

* Generates API for MySql database
* no security so just GET API

## API

* GET&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;     /api/v1/:tableName/:id
* GET&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;     /api/v1/:tableName?order=desc&page=0&size=20
* DELETE&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;     /api/v1/:tableName/:id
* POST&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;    /login
* GET&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;     /api/refresh_token


### Login API:

```bash
http -v --json POST localhost:8000/login username=admin password=admin
```


### Refresh token API:

```bash
$ http -v -f GET localhost:8000/auth/refresh_token "Authorization:Bearer xxxxxxxxx"  "Content-Type: application/json"
```

## TODO

- ~~security API~~
- POST API
- PUT API
- ~~DELETE API~~
- create table API
- alert API
