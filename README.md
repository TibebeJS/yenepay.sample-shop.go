# yenepay.sample-shop.go
![tests](https://github.com/TibebeJS/yenepay.sample-shop.go/workflows/tests/badge.svg)
![linter](https://github.com/TibebeJS/yenepay.sample-shop.go/workflows/linter/badge.svg)
![Deploy](https://github.com/TibebeJS/yenepay.sample-shop.go/workflows/Deploy/badge.svg)
![Go Report Card](https://goreportcard.com/badge/github.com/tibebejs/yenepay.sample-shop.go)

![Heroku](https://heroku-badge.herokuapp.com/?app=yenepay-go-bookshop)

An online book shop app to demonstrate [Yenepay-Go](https://github.com/TibebeJS/yenepay.sdk.go):
- Built with Go ([revel framework](https://github.com/revel/revel))
- Using a PostgreSQL database with [GORP](https://github.com/go-gorp/gorp) (ORM-ish library)
- with working authentication (Interceptors, Models)
- [YenePay](https://www.yenepay.com/) integrated for Payment (using [yenepay.sdk.go](https://github.com/TibebeJS/yenepay.sdk.go) library)

## Quick Start #
### 1. Try using a Deployed/Live Server #
The latest version of this app is continuously deployed (CD pipeline) to Heroku as soon as new changes on master branch are found.

You can use that to quickly try out the application without having to manually build/install.

[Try out the Live Demo](https://yenepay-go-bookshop.herokuapp.com/)

### 2. Manual Installation #
### pre-requisites
 - PostgreSQL
 - Go 1.x (tested with v1.15)
 - `GOPATH` should be properly configured. At minimum `revel` should be resolved in `PATH`.

    You can try running the following (on linux):
    ```
    $ export GOPATH=/home/$USER/go
    $ export PATH=$PATH:$GOPATH/bin
    ```

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


> ### Under Construction #
> ---
> ### Coming soon...
