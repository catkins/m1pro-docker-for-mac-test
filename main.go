package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"

	_ "github.com/go-sql-driver/mysql"
)

const concurrency = 100
const numRows = 1000

func main() {
	fmt.Printf("sleeping 5\n")
	time.Sleep(5 * time.Second)

	bench("test:example@(mysql)/testdb")
	// bench("test:example@(mysql-native)/testdb")
}

func bench(dbhost string) {
	fmt.Printf("starting benchmark for %s\n", dbhost)

	db, err := sql.Open("mysql", dbhost)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("starting setup for %s\n", dbhost)
	err = setup(db)
	if err != nil {
		panic(err)
	}

	defer teardown(db)

	startTime := time.Now()

	group, _ := errgroup.WithContext(context.Background())
	for i := 0; i < concurrency; i++ {
		group.Go(benchThread(db, i))
	}

	err = group.Wait()

	duration := time.Since(startTime)
	if err != nil {
		fmt.Printf("err: %#v\n", err)
	}
	fmt.Printf("finished inserting %d rows from %d goroutines for %s in %s\n", numRows*concurrency, concurrency, dbhost, duration)
}

func setup(db *sql.DB) error {
	createTable := `CREATE TABLE items (threadNum INT, rowNum INT);`
	_, err := db.Exec(createTable)
	return err
}

func teardown(db *sql.DB) error {
	dropTable := `DROP TABLE items;`
	_, err := db.Exec(dropTable)
	return err
}

func benchThread(db *sql.DB, threadNum int) func() error {
	return func() error {
		for i := 0; i < numRows; i++ {
			query := fmt.Sprintf(`INSERT INTO items (threadNum, rowNum) VALUES (%d, %d);`, threadNum, i)
			_, err := db.Exec(query)
			if err != nil {
				return fmt.Errorf("error - thread: %d - %+v", threadNum, err)
			}
		}

		return nil
	}
}
