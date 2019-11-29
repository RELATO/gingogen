# GinGoGen

Converting a MySQL database'schema to a RESTful golang APIs app in the fastest way

Based on https://github.com/dejavuzhou/felix/ginbro

## In a hurry ?
./gingogen model -u root -p [password] -a 172.16.232.100:3306 -d [dbname] -c utf8  -o=github.com/relato/dbresults

 
## GinGoGen Installation
you can install it by `go get` command：
```shell
go get github.com/relato/gingogen
```
the GinGoGen executable binary will locate in $GOPATH/bin
[check GOBIN is in your environment PATH](https://stackoverflow.com/questions/25216765/gobin-not-set-cannot-run-go-install)

## Usage

### 1. `gingogen gen` generate a new Gin+Gorm+MySQL RESTful APIs Application with JWT middleware and auth
example 

`gingogen gen -u root -p Password -a "127.0.0.1:3306" -d databasename -o "github.com/user/awesome" -c utf8 --authTable=users --authPassword=password`
```bash

$ gingogen gen -h
generate a RESTful APIs app with gin and gorm for gophers

Usage:
  gingogen gen [flags]

Examples:
gingogen gen -u root -p password -a "127.0.0.1:38306" -d dbname -c utf8 --authTable=users --authPassword=pw_column -o=github.com/relato/gingogen/out"

Flags:
  -l, --appListen string      app listen Address eg:mojotv.cn, using domain will support gin-TLS (default "127.0.0.1:5555")
      --authPassword string   password bycrpt column (default "password")
      --authTable string      the MySQL login table (default "users")
  -h, --help                  help for gen
  -o, --outPackage string     output package relative to $GOPATH/src

Global Flags:
      --config string          config file (default is $HOME/gingogen.yaml)
  -a, --mysqlAddr string       MySQL host:port (default "127.0.0.1:3306")
  -c, --mysqlCharset string    MySQL charset (default "utf8")
  -d, --mysqlDatabase string   MySQL database name
  -p, --mysqlPassword string   MySQL password (default "password")
  -u, --mysqlUser string       MySQL user name (default "root")
```

### 2. `gingogen bare` generate a bare project with one resource which you have to edit the `config.toml` which is easy for you to customize
```bash
$ gingogen bare -h
create a bare project which its mysql flags are not necessary

Usage:
  gingogen bare [flags]

Examples:
gingogen bare -o=github.com/relato/gingogen/out5"

Flags:
  -h, --help                help for bare
  -o, --outPackage string   output package relative to $GOPATH/src
```
### 3. `gingogen model` generate GORM models of tables in a MySQL database
```bash
$ genbro model -h
generate GORM models of MySQL tables.

Usage:
  gingogen model [flags]

Examples:
gingogen model -u root -p password -a 127.0.0.1:3306 -d venom -c utf8  -o=github.com/relato/gingogen/out_model

Flags:
  -h, --help                help for model
  -o, --outPackage string   eg: models,the models will be created at $GOPATH/src/models

Global Flags:
      --config string          config file (default is $HOME/gingogen.yaml)
  -a, --mysqlAddr string       MySQL host:port (default "127.0.0.1:3306")
  -c, --mysqlCharset string    MySQL charset (default "utf8")
  -d, --mysqlDatabase string   MySQL database name
  -p, --mysqlPassword string   MySQL password (default "password")
  -u, --mysqlUser string       MySQL user name (default "root")
```

## go packages
```shell
go get github.com/gin-contrib/cors
go get github.com/gin-contrib/static
go get github.com/gin-gonic/autotls
go get github.com/gin-gonic/gin
go get github.com/sirupsen/logrus
go get github.com/spf13/viper
go get github.com/spf13/cobra
go get github.com/go-redis/redis
go get github.com/go-sql-driver/mysql
go get github.com/jinzhu/gorm
go get github.com/dgrijalva/jwt-go
```
## How to fix `go get golang.org/x/crypto/bcrypt` and `go get golang.org/x/crypto/text` error
```bash
mkdir -p $GOPATH/src/golang.org/x
cd $GOPATH/src/golang.org/x
git clone https://github.com/golang/crypto
git clone https://github.com/golang/text
```
retry the command`go get github.com/relato/gingogen`

## Info
- resource table'schema which has no "ID","id","Id'" or "iD" will not generate model or route.
- the column which type is json value must be a string which is able to decode into a JSON, when resource is called POST or PATCH from the swaggerUI.
## Thanks
- [ginbro](https://github.com/dejavuzhou/felix/ginbro)
- [swagger Specification](https://swagger.io/specification/)
- [gin-gonic/gin](https://github.com/gin-gonic/gin)
- [GORM](http://gorm.io/)
- [viper](https://github.com/spf13/viper)
- [cobra](https://github.com/spf13/cobra#getting-started)
- [dgrijalva/jwt-go](https://github.com/dgrijalva/jwt-go)
- [base64captcha](https://github.com/mojocn/base64Captcha)
