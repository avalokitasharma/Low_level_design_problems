package sqldb

import (
	"fmt"
	"log"
)

func Init() {
	dbService := NewDatabase()

	err := dbService.CreateTable("users", []*Column{
		NewColumn("id", TypeInt, Required(), MinValue(1024)),
		NewColumn("username", TypeString, Required(), MaxLength(20)),
	})
	if err != nil {
		log.Fatalf("failed to create table %v", err)
	}
	err = dbService.InsertRecord("users", map[string]any{
		"id":       1030,
		"username": "hi.there@gmail.com",
	})
	if err != nil {
		log.Fatalf("failed to insert record %v", err)
	}
	records, err := dbService.GetRecords("users", nil)
	if err != nil {
		log.Fatalf("Failed to get records: %v", err)
	}

	fmt.Println("All Records:")
	for _, record := range records {
		fmt.Printf("%+v\n", record)
	}
}
