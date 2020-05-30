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

## Usage
* Prepare your json data.
* Create `struct.go`, copied from `struct.go.sample`. 
* Modify `struct.go` as json data struct.
* Then you pass json data to this script as stdin.

### Do it
```
$ cp struct.go.sample struct.go && cat test.json | go run main.go struct.go
$ sqlite3 analysis.db
sqlite> select * from records;
1|Ken|34|1985-10-26 12:00:00+09:00|ken@example.com|Father|1|Tokyo
2|Taro|43|1976-07-13 14:00:00+09:00|taro@example.com|Daughter|3|Osaka
```

## Install
```
$ go get
```

## Build
### First check your machine
```
$ go env
```
### Do it
When using mac
```
$ env GOOS=darwin GOARCH=amd64 go build -a -v main.go struct.go
```

## Contribution
1. Fork it ( https://github.com/kanagawa41/json2sqlite/fork )
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create new Pull Request
