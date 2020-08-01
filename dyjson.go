package dyjson

import (
	"encoding/json"
	"strconv"
)

type jsonDataType uint8

const (
	errorDataType jsonDataType = iota
	nullDataType
	objectDataType
	arrayDataType
	stringDataType
	numberDataType
	booleanDataType
)

// JSONValue represents a JSON value, independently of its data type.
type JSONValue struct {
	json.RawMessage

	valObject  map[string]*JSONValue
	valArray   []*JSONValue
	valString  string
	valNumber  float64
	valBoolean bool

	dataType jsonDataType
}

// Parse parses the value into a new JSONValue.
func Parse(value []byte) *JSONValue {
	return &JSONValue{
		RawMessage: value,
	}
}

// ParseString parses the value (as a string) into a new JSONValue.
func ParseString(value string) *JSONValue {
	return Parse([]byte(value))
}

// IsNull checks if the value is null.
func (v *JSONValue) IsNull() bool {
	v.parseNull()
	return v.dataType == nullDataType
}

// IsObject checks if the value is a JSON object.
func (v *JSONValue) IsObject() bool {
	v.parseObject()
	return v.dataType == objectDataType
}

// IsArray checks if the value is a JSON array.
func (v *JSONValue) IsArray() bool {
	v.parseArray()
	return v.dataType == arrayDataType
}

// IsString checks if the value is a string.
func (v *JSONValue) IsString() bool {
	v.parseString()
	return v.dataType == stringDataType
}

// IsNumber checks if the value is a number.
func (v *JSONValue) IsNumber() bool {
	v.parseNumber()
	return v.dataType == numberDataType
}

// IsBoolean checks if the value is a boolean.
func (v *JSONValue) IsBoolean() bool {
	v.parseBoolean()
	return v.dataType == booleanDataType
}

// Object returns the value parsed as a JSON object (JSONValue).
func (v *JSONValue) Object() map[string]*JSONValue {
	v.parseObject()
	return v.valObject
}

// Array returns the value parsed as a JSON array (JSONValue array).
func (v *JSONValue) Array() []*JSONValue {
	v.parseArray()
	return v.valArray
}

// String returns the value parsed as a string.
func (v *JSONValue) String() string {
	v.parseString()
	return v.valString
}

// Number returns the value parsed as a number (float64).
func (v *JSONValue) Number() float64 {
	v.parseNumber()
	return v.valNumber
}

// Boolean returns the value parsed as a boolean (bool).
func (v *JSONValue) Boolean() bool {
	v.parseBoolean()
	return v.valBoolean
}

// Set sets the value as the internal JSON value.
// Useful to update itself when any child's value changes.
func (v *JSONValue) Set() {
	switch {
	case v.IsNull():
		v.SetNull()
	case v.IsObject():
		v.SetObject(v.valObject)
	case v.IsArray():
		v.SetArray(v.valArray)
	case v.IsString():
		v.SetString(v.valString)
	case v.IsNumber():
		v.SetNumber(v.valNumber)
	case v.IsBoolean():
		v.SetBoolean(v.valBoolean)
	}
}

// SetNull sets the value as a JSON null.
func (v *JSONValue) SetNull() {
	v.RawMessage = []byte("null")
	v.dataType = nullDataType
}

// SetObject sets the value as a JSON object.
func (v *JSONValue) SetObject(val map[string]*JSONValue) {

	var i int
	v.RawMessage = []byte{'{'}
	for key, value := range val {

		if i != 0 {
			v.RawMessage = append(v.RawMessage, ',')
		}
		i++

		v.RawMessage = append(v.RawMessage, '"')
		v.RawMessage = append(v.RawMessage, key...)
		v.RawMessage = append(v.RawMessage, '"', ':')
		v.RawMessage = append(v.RawMessage, value.RawMessage...)
	}
	v.RawMessage = append(v.RawMessage, '}')

	v.valObject = val
	v.dataType = objectDataType
}

// SetArray sets the value as a JSON array.
func (v *JSONValue) SetArray(val []*JSONValue) {

	v.RawMessage = []byte{'['}
	for i, value := range val {
		if i != 0 {
			v.RawMessage = append(v.RawMessage, ',')
		}
		v.RawMessage = append(v.RawMessage, value.RawMessage...)
	}
	v.RawMessage = append(v.RawMessage, ']')

	v.valArray = val
	v.dataType = arrayDataType
}

// SetString sets the value as a JSON string.
func (v *JSONValue) SetString(val string) {

	v.RawMessage = []byte{'"'}
	v.RawMessage = append(v.RawMessage, val...)
	v.RawMessage = append(v.RawMessage, '"')

	v.valString = val
	v.dataType = stringDataType
}

// SetNumber sets the value as a JSON number.
func (v *JSONValue) SetNumber(val float64) {
	v.RawMessage = []byte(strconv.FormatFloat(val, 'f', -1, 64))
	v.valNumber = val
	v.dataType = numberDataType
}

// SetBoolean sets the value as a JSON boolean.
func (v *JSONValue) SetBoolean(val bool) {
	if val {
		v.RawMessage = []byte("true")
	} else {
		v.RawMessage = []byte("false")
	}
	v.valBoolean = val
	v.dataType = booleanDataType
}

func (v *JSONValue) parseNull() {
	if v.dataType != 0 {
		return
	}

	var val interface{}
	if json.Unmarshal(v.RawMessage, &val) == nil && val == nil {
		v.dataType = nullDataType
	}
}

func (v *JSONValue) parseObject() {
	if v.dataType != 0 {
		return
	}

	if json.Unmarshal(v.RawMessage, &v.valObject) == nil && v.valObject != nil {
		v.dataType = objectDataType
	}
}

func (v *JSONValue) parseArray() {
	if v.dataType != 0 {
		return
	}

	if json.Unmarshal(v.RawMessage, &v.valArray) == nil && v.valArray != nil {
		v.dataType = arrayDataType
	}
}

func (v *JSONValue) parseString() {
	if v.dataType != 0 {
		return
	}

	if json.Unmarshal(v.RawMessage, &v.valString) == nil {
		v.dataType = stringDataType
	}
}

func (v *JSONValue) parseNumber() {
	if v.dataType != 0 {
		return
	}

	if json.Unmarshal(v.RawMessage, &v.valNumber) == nil {
		v.dataType = numberDataType
	}
}

func (v *JSONValue) parseBoolean() {
	if v.dataType != 0 {
		return
	}

	if json.Unmarshal(v.RawMessage, &v.valBoolean) == nil {
		v.dataType = booleanDataType
	}
}
