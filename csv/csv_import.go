package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	_ = godotenv.Load(".env")
	var myEnv map[string]string
	myEnv, _ = godotenv.Read()

	db, err := sql.Open("mysql", fmt.Sprintf("%s/%s", myEnv["DATABASE_DSN"], myEnv["DATABASE_NAME"]))

	if err != nil {
		log.Fatal(fmt.Sprintf("%s: %s", "Database could not connect", err))
	}

	_, err = db.Exec("DROP TABLE IF EXISTS gender")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table Droped succesfully")

	_, err = db.Exec(fmt.Sprintf("%s %s ("+
		"id int(11) NOT NULL auto_increment,"+
		"name varchar(255) NOT NULL,"+
		"country_code varchar(5) not null,"+
		"gender varchar(1) NOT NULL,"+
		"PRIMARY KEY  (id)"+
		")", "CREATE TABLE IF NOT EXISTS", myEnv["TABLE_NAME"]))

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table created successfully.")

	records := readCsvFile("csv/wgnd_ctry.csv")

	var name string
	var gender string
	var country string

	for _, cells := range records {
		name = cells[0]
		country = cells[1]
		gender = cells[2]

		_, err = db.Exec("INSERT INTO gender(name, country_code, gender) VALUES('" + name + "', '" + country + "', '" + gender + "')")

		if err != nil {
			fmt.Println(fmt.Sprintf("%s: %s ", "Could not inserted record", name, gender))
		}
	}

	fmt.Println("Record insert process completed")
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)

	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records[1:]
}
