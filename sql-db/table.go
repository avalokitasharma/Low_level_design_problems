package sqldb

import "fmt"

type Table struct {
	Name    string
	Columns []*Column
	Rows    []map[string]any
}

func NewTable(name string, Columns []*Column) *Table {
	return &Table{
		Name:    name,
		Columns: Columns,
		Rows:    []map[string]any{},
	}
}

func (t *Table) AddRow(r map[string]any) error {
	//col is fixed - so check for value of each col
	// range over col name

	for _, col := range t.Columns {
		value, ok := r[col.Name]
		if !ok {
			if col.Constraints.Required {
				return fmt.Errorf("required column %s is missing", col.Name)
			}
			continue
		}
		// validation for that col
		// if not working - return error
		if err := col.Validate(value); err != nil {
			return err
		}
	}

	// check for unknown column
	for colName := range r {
		found := false
		for _, col := range t.Columns {
			if col.Name == colName {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("unkown column %s", colName)
		}
	}
	safeCopy := make(map[string]any)
	for col, value := range r {
		safeCopy[col] = value
	}
	t.Rows = append(t.Rows, safeCopy)
	fmt.Printf("1 Row added successfully.\n")
	return nil
}

func (t *Table) GetRows(filter map[string]any) []map[string]any {
	// select  *
	if len(filter) == 0 {
		return t.Rows
	}

	// loop over table - check each row
	var matchedRows []map[string]any
	for _, row := range t.Rows {
		matches := true
		for col, val := range filter {
			rowVal, ok := row[col]
			if !ok || val != rowVal {
				matches = false
				break
			}
		}
		if matches {
			matchedRows = append(matchedRows, row)
		}
	}
	return matchedRows
}
