# go-mysql-rest-api

A guide for creating RESTful API with Golang and MySQL and Gin.

## Build and Usage

```bash
go mod tidy
go mod verify
go build
```

```bash
./mysql-restful-server
```

```bash
http://localhost:8989
```

## Config

conf.yaml

```yaml
mysql:
  host: "127.0.0.1"
  port: "3306"
  user: "root"
  pwd: "root"
  db: "test"
  enable_auth: "true"
  auth_table: "auth"
  auth_name: "username"
  auth_pwd: "passwd"
```
for security all insert / delete / update operation must login  
so if you want to insert / delete / update data please enable auth  

you have no permission to read or modify any data from ```auth_table```  

TODO： Make the fields corresponding to  ```auth_name``` and ```auth_pwd```    configurable
TODO： Need to guard against the risk of SQL injection attacks

## Features

* Generates API for MySql database
* can not use ${auth_table}

## API

|Method         | Path           | Auth  | Operation  |
| :-------------: |:-------------| :-----:| :-----:| 
| GET     | /api/v1/:tableName/:id | NO | serect * from tableName where id = ? |
| GET     | /api/v1/:tableName?order=desc&page=0&size=20 | NO | select * from tableName where id>0 order by id desc limit 0,20 |
| POST     | /login | NO | login |
| GET     | /api/refresh_token | YES | refresh_token |
| DELETE     | /api/v1/:tableName/:id | YES | delete from tableName where id = ? |
| POST     | /api/v1/:tableName| YES | insert into tableName (data.key) values (data.value) |
| PUT     | /api/v1/:tableName| YES | create a table | 


## Security

### Login API:

```bash
POST /login HTTP/1.1
Host: 127.0.0.1:8989
Content-Type: application/json
Cache-Control: no-cache
{
	"username":"admin",
	"password":"admin"
}

{
    "code": 200,
    "expire": "2018-01-05T15:26:18+08:00",
    "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTUxMzcxNzgsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTUxNTEzMzU3OH0.D48Ada0pVR72nAS_gt8TTxzvtdy2s-OAnoizbmRIhtunciw5905G7QCcJZHqJvxcje4GBWA6e9wvOOEz7OVk9lrsTxPfFHwSnqkoj7ZkCGGkOIYkC-SVaVJB47Ez6yxhQljwHf_LiWVnkLpBN4y4eerqEErh-R4VXmZ9ZnJQdO3V78ZiXdaY2oMAmX7-JYHz6LOlTqjxMsZ8KHUrRRt5mDbLZxw4Ni_Ww-vetR3dNwIsCH_-ExsE6Z9UQlOP_yTo7iD09_sbyeSAB-ZE0e7qnOjgWCgujZJxFPsoWCIQV4O4ONWTpVZxds3eLjWIlyBlyV2LHi85b2f-nmOfRQphDw"
}

```


### Refresh token API:

```bash
GET /api/refresh_token HTTP/1.1
Host: 127.0.0.1:8989
Authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTUxMzcxNzgsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTUxNTEzMzU3OH0.D48Ada0pVR72nAS_gt8TTxzvtdy2s-OAnoizbmRIhtunciw5905G7QCcJZHqJvxcje4GBWA6e9wvOOEz7OVk9lrsTxPfFHwSnqkoj7ZkCGGkOIYkC-SVaVJB47Ez6yxhQljwHf_LiWVnkLpBN4y4eerqEErh-R4VXmZ9ZnJQdO3V78ZiXdaY2oMAmX7-JYHz6LOlTqjxMsZ8KHUrRRt5mDbLZxw4Ni_Ww-vetR3dNwIsCH_-ExsE6Z9UQlOP_yTo7iD09_sbyeSAB-ZE0e7qnOjgWCgujZJxFPsoWCIQV4O4ONWTpVZxds3eLjWIlyBlyV2LHi85b2f-nmOfRQphDw
Cache-Control: no-cache

{
    "code": 200,
    "expire": "2018-01-05T15:26:55+08:00",
    "token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTUxMzcyMTUsImlkIjoiYWRtaW4iLCJvcmlnX2lhdCI6MTUxNTEzMzU3OH0.lWJC6BaK5XC9N1Wc66MnxDJD-JXOCdAFwW7qGeIjRvPn6R5vYWgB559zeOC8bLxqhZW7CFZplzZQwuM9H3AjQuU5u7Iiaw4zjS1c2W180U_dPqUz1IeZA3zLpuSLjvNWAxGS-iw7B7aUmpJ7KC9ubBHLItXenKbiZn77SOys3zgNwLm_BfkoOMZj2GXxZPLderxj7GR06oNeARy_hXTUM4wa4-C83R6x5OH22VJXiXmNhIDBv5m0AiK7JYZmpbMr6gSGTNVhUM5971ww7u64Ly2viSO0_vnPWR-L-zOKZVVjwJAkdzScpxXnHyXOQTSKcrJETh7OBL4lU2TaQm941w"
}
```

## TODO

- ~~security API~~
- ~~DELETE API~~
- ~~POST API~~
- PUT API
- create table API
- alert API
