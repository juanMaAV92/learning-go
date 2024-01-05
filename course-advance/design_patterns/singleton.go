package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct{}

func (db *Database) CreateSingleConnection() {
	fmt.Println("Creating Singleton for database")
	time.Sleep(2 * time.Second)
	fmt.Println("Singleton created")
}

var db *Database
var lock sync.Mutex

func GetDatabaseInstance() *Database {
	lock.Lock()
	defer lock.Unlock()
	if db == nil {
		fmt.Println("Creating database instance")
		db = &Database{}
		db.CreateSingleConnection()
	} else {
		fmt.Println("Database instance already created")
	}
	return db
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetDatabaseInstance()
		}()
	}

	wg.Wait()
}
