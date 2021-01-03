# yenepay.sample-shop.go
![tests](https://github.com/TibebeJS/yenepay.sample-shop.go/workflows/tests/badge.svg)
[![codecov](https://codecov.io/gh/TibebeJS/yenepay.sample-shop.go/branch/main/graph/badge.svg?token=PQTt3MS57J)](https://codecov.io/gh/TibebeJS/yenepay.sample-shop.go)
![linter](https://github.com/TibebeJS/yenepay.sample-shop.go/workflows/linter/badge.svg)
![Deploy](https://github.com/TibebeJS/yenepay.sample-shop.go/workflows/Deploy/badge.svg)
![Go Report Card](https://goreportcard.com/badge/github.com/tibebejs/yenepay.sample-shop.go)

> ### Under Construction #
> ---
> ### Coming soon...
## How to run #
### 1. Use the Deployed/Live Server #
The latest version of this app is automatically deployed to Heroku using Github Action. You can use that to quickly try out the application without having to manually build/install.

[Try out the Live Demo](https://yenepay-go-bookshop.herokuapp.com/)

### 2. Manual Installation #
### pre-requisite
 - PostgreSQL
 - Go 1.x (tested with v1.15)
 - `GOPATH` should be properly configured (at least `revel` command should resolved)
1. Clone this repository: 
```
$ git clone git@github.com:TibebeJS/yenepay.sample-shop.go.git
```
2. `cd` into the newly created folder:
```
$ cd yenepay.sample-shop.go
```
3. set an `DATABASE_URL` environment variable with the postgres connection string:
```
$ export DATABASE_URL="host=<your_db> user=<your_user> password=<your_password> dbname=<your_db> sslmode=disable"
```
Make sure to replace all the values inside "`<`" and "`>`" with your credential values.

**Example:**
```
$ export DATABASE_URL="host=localhost user=postgres password=postgres dbname=my_db sslmode=disable"
```

4. install dependencies:
```
$ go get github.com/revel/cmd/revel
$ go get github.com/go-gorp/gorp
```
5. remove previous builds and run the application:
```
$ rm -rf target
$ revel run -a .
```
6. Navigate to `http://localhost:9000` on your favourite browser.

### 3. Using Docker/Docker-Compose #

1. Clone this repository:
```
$ git clone git@github.com:TibebeJS/yenepay.sample-shop.go.git
```
2. `cd` into the newly created folder:
```
$ cd yenepay.sample-shop.go
```
3. Start App via Docker-compose:
```
$ docker-compose up --build
```
4. Navigate to `http://localhost:9000` on your favourite browser :)


