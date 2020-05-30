package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	gormbulk "github.com/t-tiger/gorm-bulk-insert/v2"

	"github.com/jinzhu/gorm"
)

func init() {
	// The sentense for -h option
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage of %s:
stdout_as_data_json | %s
`, os.Args[0], os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	bytes, _ := ioutil.ReadAll(os.Stdin)
	var records []Record
	err := json.Unmarshal(bytes, &records)
	if err != nil {
		panic(err)
	}

	// For value must be kind of Struct
	var iRecords []interface{}
	for _, record := range records {
		iRecords = append(iRecords, record)
	}

	err = initMigration()
	if err != nil {
		panic(err)
	}

	db, err := connection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for offset, offsetP, arrLen := 0, onceInsertNum, len(iRecords); offset <= arrLen; offset += offsetP {
		var limit int
		if arrLen < offset+offsetP {
			limit = arrLen
		} else {
			limit = offset + offsetP
		}
		fmt.Println(offset, limit)
		err = gormbulk.BulkInsert(db, iRecords[offset:limit], onceInsertNum)
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("Finished insert into database of SQLite. Check %s.\n", dbFileName)
}

func initMigration() error {
	db, err := gorm.Open("sqlite3", dbFileName)
	if err != nil {
		return err
	}
	defer db.Close()
	db.AutoMigrate(&Record{})

	return nil
}

func connection() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", dbFileName)
	if err != nil {
		return nil, err
	}

	return db, nil
}
