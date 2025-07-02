package sqldb

import (
	"fmt"
	"reflect"
)

type ColumnType string

const (
	TypeString = "string"
	TypeInt    = "int"
)

type ColumnConstraint struct {
	Required  bool
	MaxLength *int
	MinValue  *int
}

type Column struct {
	Name        string
	Type        ColumnType
	Constraints ColumnConstraint
}

func NewColumn(name string, colType ColumnType, constraints ...func(*ColumnConstraint)) *Column {
	col := Column{
		Name: name,
		Type: colType,
	}
	for _, constraint := range constraints {
		constraint(&col.Constraints)
	}
	return &col
}
func Required() func(*ColumnConstraint) {
	return func(cc *ColumnConstraint) {
		cc.Required = true
	}
}

func (c *Column) Validate(value any) error {
	if value == nil {
		if c.Constraints.Required {
			return fmt.Errorf("Column %s is required", c.Name)
		}
		return nil
	}

	// type check
	switch c.Type {
	case TypeString:
		strVal, ok := value.(string)

		if !ok {
			return fmt.Errorf("expected string for column %s but got %s", c.Name, strVal)
		}
		if c.Constraints.MaxLength != nil && len(strVal) > *c.Constraints.MaxLength {
			return fmt.Errorf("string length for %s exceeds max length %d", c.Name, *c.Constraints.MaxLength)
		}
	case TypeInt:
		intVal, err := convertToInt(value)
		if err != nil {
			return fmt.Errorf("expected int for column %s but got %T", c.Name, value)
		}
		if c.Constraints.MinValue != nil && int(intVal) < *c.Constraints.MinValue {
			return fmt.Errorf("Min value for %s exceeds  %d", c.Name, *c.Constraints.MinValue)
		}

	}
	return nil
}
func convertToInt(val any) (int64, error) {
	v := reflect.ValueOf(val)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int(), nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(v.Uint()), nil
	default:
		return 0, fmt.Errorf("can not convert %T to int64", val)
	}
}

func MinValue(val int) func(*ColumnConstraint) {
	return func(cc *ColumnConstraint) {
		cc.MinValue = &val
	}
}
func MaxLength(length int) func(*ColumnConstraint) {
	return func(cc *ColumnConstraint) {
		cc.MaxLength = &length
	}
}
