package db

import (
	"database/sql"
	"io/ioutil"
	"tw.ntust.dripmonitor/logger/helpers"
	"fmt"
)

func SeedDB(db *sql.DB) {
	dat, err := ioutil.ReadFile(helpers.ProjectPath + "/setup/db/seed.sql")
	if err != nil {
		panic(err)
	}
	sqlQuery := string(dat)

	_, err = db.Exec(sqlQuery)
	if err != nil {
		panic(err)
	}

	fmt.Println("DB seeding finished.")
}
