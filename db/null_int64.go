package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
)

// NullInt64 represents an int64 that may be null.
type NullInt64 struct {
	Int64 int64
	Valid bool // Valid is true if Int64 is not NULL
}

// MarshalJSON try to marshalize to json
func (n NullInt64) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return []byte(fmt.Sprintf("%d", n.Int64)), nil
	}

	return []byte("null"), nil
}

// UnmarshalJSON try to unmarshal data from input
func (n *NullInt64) UnmarshalJSON(b []byte) error {
	if strings.ToLower(string(b)) == "null" {
		n.Int64 = 0
		n.Valid = false
		return nil
	}
	val, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return err
	}
	n.Int64 = val
	n.Valid = !(string(b) == "")

	return nil
}

// Scan implements the Scanner interface.
func (n *NullInt64) Scan(value interface{}) error {
	temp := &sql.NullInt64{}
	if err := temp.Scan(value); err != nil {
		return err
	}

	n.Int64 = temp.Int64
	n.Valid = temp.Valid
	return nil
}

// Value implements the driver Valuer interface.
func (n NullInt64) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Int64, nil
}
