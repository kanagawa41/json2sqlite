package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

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
	var records Records
	json.Unmarshal(bytes, &records)

	err := initMigration()
	if err != nil {
		panic(err)
		return
	}

	db, err := connection()
	if err != nil {
		panic(err)
		return
	}
	defer db.Close()

	var errCnt int
	for _, record := range records {
		errors := db.Create(&record).GetErrors()

		for _, err := range errors {
			errCnt++
			fmt.Fprint(os.Stderr, err)
		}
	}

	fmt.Printf("Finished insert into database of SQLite. Check %s.\n", dbFileName)
	if errCnt > 0 {
		fmt.Printf("%d of %d rows were failed insert.\n", errCnt, len(records))
	}
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
