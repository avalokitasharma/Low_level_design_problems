package sqldb

import (
	"fmt"
	"sync"
)

type Database struct {
	mu     sync.RWMutex
	tables map[string]*Table
}

func NewDatabase() *Database {
	return &Database{tables: make(map[string]*Table)}
}

func (db *Database) CreateTable(name string, columns []*Column) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, exists := db.tables[name]; exists {
		return fmt.Errorf("table %s already exists", name)
	}
	table := NewTable(name, columns)
	db.tables[name] = table
	fmt.Printf("table '%s' created successfully.\n", name)
	return nil
}

func (db *Database) GetTable(name string) (*Table, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	table, ok := db.tables[name]
	if !ok {
		return table, fmt.Errorf("table %s is not found", name)
	}
	return table, nil
}
func (db *Database) DeleteTable(name string) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	if _, ok := db.tables[name]; !ok {
		return fmt.Errorf("table %s is not found", name)
	}

	delete(db.tables, name)
	return nil
}
func (db *Database) GetRecords(tableName string, filter map[string]any) ([]map[string]any, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()

	table, exists := db.tables[tableName]
	if !exists {
		return nil, fmt.Errorf("table %s not found", tableName)
	}

	return table.GetRows(filter), nil

}

func (db *Database) InsertRecord(tableName string, record map[string]any) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	table, exists := db.tables[tableName]
	if !exists {
		return fmt.Errorf("table %s not found", tableName)
	}

	return table.AddRow(record)
}
