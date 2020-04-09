# json2sqlite
====

# Overview
Insert your json data into SQLite.

## Description
You can analysis json file using SQL with a few steps.
This script is made by golang and need acknowledge about one, but it's easy :)

## Demo
```
$ cp struct-sample.go struct.go
$ cat test.json | go run main.go struct.go
Finished insert into database of SQLite. Check analysis.db.
$ sqlite3 analysis.db
```

## Requirement
* golang
* github.com/jinzhu/gorm

## Usage
* Prepare your json data.
* Create `struct.go`, copied from `struct-sample.go`. 
* Modify `struct.go` as json data struct.
* Then you pass that to this script as stdin.

### Tips
* If you don't want to get the cause of error, you specify "2>", like below:
```
$ cat test.json | go run main.go struct.go 2>
```

## Install
```
$ go mod vendor
```

## Build
### First check your machine
```
$ go env
```
### Do it
In my case, using mac
```
$ env GOOS=darwin GOARCH=amd64 go build -a -v main.go struct.go
```

## Contribution
1. Fork it ( https://github.com/kanagawa41/json2sqlite/fork )
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create new Pull Request
