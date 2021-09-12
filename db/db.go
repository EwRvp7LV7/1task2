package db1

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "wordstestdb"
)

//могут быть одинаковые int, поэтому работаем с struct
//using structs slice for sorting
type Pair struct {
	Key   int
	Value string
}

func AddWordsToDB(pll *[]Pair) {

	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// CREATE TABLE words (
	//   id serial NOT NULL PRIMARY KEY,
	//   fit int CHECK (fit > 0),
	//   keyword VARCHAR(50)
	// );
	// CREATE INDEX keyword_idx ON words (keyword);

	for _, p1 := range *pll {
		//EXISTS понимает только с WHERE
		// res, err := db.Exec("IF NOT EXISTS (SELECT * FROM words WHERE keyword = $1) THEN INSERT INTO words(fit,keyword) VALUES($1,$2) returning id END IF;", p1.Key, p1.Value)
		res, err := db.Exec("INSERT INTO words(fit,keyword) SELECT $1, $2 WHERE NOT EXISTS (SELECT 1 FROM words WHERE keyword = CAST($2 AS VARCHAR))", p1.Key, p1.Value)

		if err != nil {
			log.Println(err)
			fmt.Println(err)
		}

		affected, err := res.RowsAffected()
		if err != nil {
			log.Println(err)
		}

		if affected > 0 {
			fmt.Println("added ", p1)
		}
		

	}

}
