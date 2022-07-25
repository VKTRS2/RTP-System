// Package xsdt
// Do not Edit. This stuff it's been automatically generated.
package xsdt

import (
	"fmt"
	"strconv"
	"time"
)

/*
 * Boolean
 */

const (
	BooleanSample = false
	BooleanZero   = false
)

func (me Boolean) IsValid(optional bool) bool {
	return true
}

func (me Boolean) IsZero() bool {
	return true
}

func ToBoolean(intf interface{}) (Boolean, error) {
	var err error
	b := false
	switch v := intf.(type) {
	case bool:
		b = v
	case string:
		b, err = strconv.ParseBool(v)
		if err != nil {
			return false, err
		}
	default:
		return false, fmt.Errorf("not a boolean %v", intf)
	}

	return Boolean(b), nil
}

func MustToBoolean(intf interface{}) Boolean {
	b, err := ToBoolean(intf)
	if err != nil {
		panic(err)
	}

	return b
}

/*
 * AnyType
 */

const (
	AnyTypeSample = "xsdt.AnyType"
	AnyTypeZero   = ""
)

func (me AnyType) IsValid(optional bool) bool {
	return true
}

func (me AnyType) IsZero() bool {
	return true
}

func ToAnyType(intf interface{}) (AnyType, error) {
	s := fmt.Sprintf("%v", intf)
	return AnyType(s), nil
}

func MustToAnyType(intf interface{}) AnyType {
	s, err := ToAnyType(intf)
	if err != nil {
		panic(err)
	}

	return s
}

/*
 * Decimal
 */

const (
	DecimalSample = "100.00"
	DecimalZero   = ""
)

func (me Decimal) IsValid(optional bool) bool {
	return true
}

func (me Decimal) IsZero() bool {
	return true
}

func ToDecimal(intf interface{}) (Decimal, error) {
	d := ""
	switch v := intf.(type) {
	case float64:
		d = fmt.Sprintf("%v", v)
	case string:
		d = v
	case Decimal:
		return v, nil
	default:
		return "", fmt.Errorf("not a boolean %v", intf)
	}

	return Decimal(d), nil
}

func MustToDecimal(intf interface{}) Decimal {
	b, err := ToDecimal(intf)
	if err != nil {
		panic(err)
	}

	return b
}

/*
 * String
 */

const (
	StringSample = "string value"
	StringZero   = ""
)

func (me String) IsValid(optional bool) bool {
	return true
}

func (me String) IsZero() bool {
	return true
}

func ToString(intf interface{}) (String, error) {
	s := fmt.Sprintf("%v", intf)
	return String(s), nil
}

func MustToString(intf interface{}) String {
	s, err := ToString(intf)
	if err != nil {
		panic(err)
	}

	return s
}

/*
 * Base64Binary
 */

func (me Base64Binary) IsValid(optional bool) bool {
	return true
}

func (me Base64Binary) IsZero() bool {
	return true
}

func ToBase64Binary(intf interface{}) (Base64Binary, error) {
	s := fmt.Sprintf("%v", intf)
	return Base64Binary(s), nil
}

func MustToBase64Binary(intf interface{}) Base64Binary {
	s, err := ToBase64Binary(intf)
	if err != nil {
		panic(err)
	}

	return s
}

/*
 * Date
 */

func (me Date) IsValid(optional bool) bool {
	return true
}

func (me Date) IsZero() bool {
	return me == "" || me == "0001-01-01"
}

func ToDateFromTime(tm time.Time) Date {
	return Date(tm.Format("2006-01-02"))
}

/*
 * DateTime
 */

func (me DateTime) IsValid(optional bool) bool {

	if me.IsZero() {
		if optional {
			return true
		}
		return false
	}

	return true
}

func (me DateTime) IsZero() bool {
	return me == "" || me == "0001-01-01T00:00:00Z"
}

func ToDateTimeFromTime(tm time.Time) DateTime {
	return DateTime(tm.Format(time.RFC3339))
}
