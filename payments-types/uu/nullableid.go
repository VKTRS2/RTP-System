package uu

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"io"
	"reflect"
)

var IDNull NullableID

// NullableID is a UUID where the Nil UUID "00000000-0000-0000-0000-000000000000"
// is interpreted as the null values for SQL and JSON.
type NullableID [16]byte

// NullableIDFromString parses a string as NullableID.
// The Nil UUID "00000000-0000-0000-0000-000000000000"
// is interpreted as NULL.
func NullableIDFromString(s string) (NullableID, error) {
	id, err := IDFromString(s)
	if err != nil {
		return IDNull, err
	}
	return NullableID(id), nil
}

// NullableIDFromBytes creates a new valid NullableID
func NullableIDFromBytes(s []byte) (NullableID, error) {
	id, err := IDFromBytes(s)
	if err != nil {
		return IDNull, err
	}
	return NullableID(id), nil
}

// NullableIDFromPtr creates a new NullableID that be null if ptr is nil.
func NullableIDFromPtr(ptr *ID) NullableID {
	if ptr == nil {
		return IDNull
	}
	return NullableID(*ptr)
}

// Version returns algorithm version used to generate UUID.
func (n NullableID) Version() uint {
	return ID(n).Version()
}

// Variant returns an ID layout variant or IDVariantInvalid if unknown.
func (n NullableID) Variant() uint {
	return ID(n).Variant()
}

// Valid returns if Variant and Version of this UUID are supported.
// A Nil UUID is also valid.
func (n NullableID) Valid() bool {
	return n == IDNull || ID(n).Valid()
}

// Validate returns an error if the Variant and Version of this UUID are not supported.
// A Nil UUID is also valid.
func (n NullableID) Validate() error {
	if n == IDNull {
		return nil
	}
	return ID(n).Validate()
}

// Set sets an ID for this NullableID
func (n *NullableID) Set(id ID) {
	*n = NullableID(id)
}

// SetNull sets the NullableID to null
func (n *NullableID) SetNull() {
	*n = IDNull
}

// Get returns the non nullable ID value
// or panics if the NullableID is null.
// Note: check with IsNull before using Get!
func (n NullableID) Get() ID {
	if n.IsNull() {
		panic("NULL uu.ID")
	}
	return ID(n)
}

// GetOrNil returns the non nullable ID value
// or the Nil UUID if n is null.
// Use Get to ensure getting a non Nil UUID or panic.
func (n NullableID) GetOrNil() ID {
	return ID(n)
}

// PrettyPrint the NullableID in the format xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx or as NULL.
// Implements pretty.Printer.
func (n NullableID) PrettyPrint(w io.Writer) {
	w.Write([]byte(n.StringOr("NULL")))
}

// GoString returns a pseudo Go literal for the ID in the format:
//   uu.ID("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")
// It compiles if uu.ID is replace by uu.IDMustFromString.
func (n NullableID) GoString() string {
	return `uu.ID("` + n.String() + `")`
}

// Hex returns the hex representation without dashes of the UUID
// The returned string is always 32 characters long.
func (n NullableID) Hex() string {
	return ID(n).Hex()
}

// Base64 returns the unpadded base64 URL encoding of the UUID.
// The returned string is always 22 characters long.
func (n NullableID) Base64() string {
	return ID(n).Base64()
}

// IsNull returns true if the NullableID is null.
// IsNull implements the nullable.Nullable interface.
func (n NullableID) IsNull() bool {
	return n == IDNull
}

// IsNotNull returns true if the NullableID is not null.
func (n NullableID) IsNotNull() bool {
	return n != IDNull
}

// String returns the ID as string or "NULL"
func (n NullableID) String() string {
	return n.StringOr("NULL")
}

// StringOr returns the ID as string or the passed nullStr
func (n NullableID) StringOr(nullStr string) string {
	if n.IsNull() {
		return nullStr
	}
	return ID(n).String()
}

// StringBytes returns the canonical string representation of the UUID as byte slice:
// xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
func (n NullableID) StringBytes() []byte {
	return ID(n).StringBytes()
}

// Ptr returns a pointer to this NullableID's value, or a nil pointer if this NullableID is null.
func (n NullableID) Ptr() *ID {
	if n == IDNull {
		return nil
	}
	return (*ID)(&n)
}

// Value implements the driver.Valuer interface.
func (n NullableID) Value() (driver.Value, error) {
	if n == IDNull {
		return nil, nil
	}
	// Delegate to ID Value function
	return ID(n).Value()
}

// Scan implements the sql.Scanner interface.
func (n *NullableID) Scan(src interface{}) error {
	if src == nil {
		*n = IDNull
		return nil
	}
	// Delegate to ID.Scan function
	return (*ID)(n).Scan(src)
}

// UnmarshalJSON implements json.Unmarshaler.
// It supports string and null input. Blank string input does not produce a null ID.
// It also supports unmarshalling a sql.NullString.
func (n *NullableID) UnmarshalJSON(data []byte) error {
	// TODO optimize
	var v interface{}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}

	switch x := v.(type) {
	case string:
		id, err := IDFromString(x)
		if err != nil {
			return err
		}
		*n = NullableID(id)
		return err

	case map[string]interface{}:
		var ns sql.NullString
		err = json.Unmarshal(data, &ns)
		if err != nil || !ns.Valid {
			return err
		}
		id, err := IDFromString(ns.String)
		if err != nil {
			return err
		}
		*n = NullableID(id)
		return err

	case nil:
		*n = IDNull
		return nil

	default:
		return fmt.Errorf("cannot UnmarshalJSON(%s) as uu.NullableID", reflect.TypeOf(v))
	}
}

// MarshalJSON implements json.Marshaler.
func (n NullableID) MarshalJSON() ([]byte, error) {
	if n == IDNull {
		return []byte("null"), nil
	}
	b := make([]byte, 1, 38)
	b[0] = '"'
	b = append(b, n.StringBytes()...)
	b = append(b, '"')
	return b, nil
}

// MarshalText implements encoding.TextMarshaler.
// It will encode a blank string when this String is null.
func (n NullableID) MarshalText() ([]byte, error) {
	if n == IDNull {
		return []byte{}, nil
	}
	return ID(n).MarshalText()
}

// UnmarshalText implements encoding.TextUnmarshaler.
// It will unmarshal to a null String if the input is a blank string.
func (n *NullableID) UnmarshalText(text []byte) (err error) {
	if len(text) == 0 {
		*n = IDNull
		return nil
	}
	return (*ID)(n).UnmarshalText(text)
}

// MarshalBinary implements the encoding.BinaryMarshaler interface.
func (n NullableID) MarshalBinary() (data []byte, err error) {
	return ID(n).MarshalBinary()
}

// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface.
// It will return error if the slice isn't 16 bytes long,
// but does not check the validity of the UUID.
func (n *NullableID) UnmarshalBinary(data []byte) (err error) {
	return (*ID)(n).UnmarshalBinary(data)
}

// NullableIDCompare returns bytes.Compare result of a and b.
func NullableIDCompare(a, b NullableID) int {
	return bytes.Compare(a[:], b[:])
}
