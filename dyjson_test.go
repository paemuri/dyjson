package dyjson

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	v := Parse([]byte("test"))
	assert.Equal(t, []byte("test"), []byte(v.RawMessage))
}

func TestParseString(t *testing.T) {
	v := ParseString("test")
	assert.Equal(t, []byte("test"), []byte(v.RawMessage))
}

func TestIsNull_ValidInput(t *testing.T) {
	v := ParseString(`null`)
	assert.True(t, v.IsNull())
}

func TestIsNull_InvalidInput(t *testing.T) {
	v := ParseString(`test`)
	assert.False(t, v.IsNull())
}

func TestIsObject_ValidInput(t *testing.T) {
	v := ParseString(`{"batata":true}`)
	assert.True(t, v.IsObject())
}

func TestIsObject_InvalidInput(t *testing.T) {
	v := ParseString(`test`)
	assert.False(t, v.IsObject())
}

func TestIsArray_ValidInput(t *testing.T) {
	v := ParseString(`["batata",true]`)
	assert.True(t, v.IsArray())
}

func TestIsArray_InvalidInput(t *testing.T) {
	v := ParseString(`test`)
	assert.False(t, v.IsArray())
}

func TestIsString_ValidInput(t *testing.T) {
	v := ParseString(`"batata"`)
	assert.True(t, v.IsString())
}

func TestIsString_InvalidInput(t *testing.T) {
	v := ParseString(`test`)
	assert.False(t, v.IsString())
}

func TestIsNumber_ValidInput(t *testing.T) {
	v := ParseString(`1`)
	assert.True(t, v.IsNumber())
}

func TestIsNumber_InvalidInput(t *testing.T) {
	v := ParseString(`test`)
	assert.False(t, v.IsNumber())
}

func TestIsBoolean_ValidInput(t *testing.T) {
	v := ParseString(`true`)
	assert.True(t, v.IsBoolean())
}

func TestIsBoolean_InvalidInput(t *testing.T) {
	v := ParseString(`test`)
	assert.False(t, v.IsBoolean())
}

func TestObject(t *testing.T) {
	v := ParseString(`{"batata":true}`)
	assert.Equal(t, ParseString("true"), v.Object()["batata"])
}

func TestArray(t *testing.T) {
	v := ParseString(`["batata",true]`)
	assert.Equal(t, ParseString("true"), v.Array()[1])
}

func TestString(t *testing.T) {
	v := ParseString(`"batata"`)
	assert.Equal(t, "batata", v.String())
}

func TestNumber(t *testing.T) {
	v := ParseString(`1`)
	assert.Equal(t, 1.0, v.Number())
}

func TestBoolean(t *testing.T) {
	v := ParseString(`true`)
	assert.Equal(t, true, v.Boolean())
}

func TestSet_NullValue(t *testing.T) {
	v := ParseString(`null`)
	v.Set()
	assert.Equal(t, nullDataType, v.dataType)
}

func TestSet_ObjectValue(t *testing.T) {
	v := ParseString(`{"batata":true}`)
	v.Set()
	assert.Equal(t, objectDataType, v.dataType)
}

func TestSet_ArrayValue(t *testing.T) {
	v := ParseString(`["batata",true]`)
	v.Set()
	assert.Equal(t, arrayDataType, v.dataType)
}

func TestSet_StringValue(t *testing.T) {
	v := ParseString(`"batata"`)
	v.Set()
	assert.Equal(t, stringDataType, v.dataType)
}

func TestSet_NumberValue(t *testing.T) {
	v := ParseString(`1`)
	v.Set()
	assert.Equal(t, numberDataType, v.dataType)
}

func TestSet_BooleanValue(t *testing.T) {
	v := ParseString(`true`)
	v.Set()
	assert.Equal(t, booleanDataType, v.dataType)
}

func TestSet_NoValue(t *testing.T) {
	v := Parse(nil)
	v.Set()
	assert.Equal(t, errorDataType, v.dataType)
}

func TestSetNull(t *testing.T) {
	v := Parse(nil)
	v.SetNull()
	assert.Equal(t, nullDataType, v.dataType)
	assert.Equal(t, []byte("null"), []byte(v.RawMessage))
}

func TestSetObject(t *testing.T) {
	v1 := ParseString(`{"batata":true}`)
	v2 := Parse(nil)
	v2.SetObject(v1.Object())
	assert.Equal(t, objectDataType, v2.dataType)
	assert.Equal(t, []byte(v2.RawMessage), []byte(v1.RawMessage))
}

func TestSetArray(t *testing.T) {
	v1 := ParseString(`["batata",true]`)
	v2 := Parse(nil)
	v2.SetArray(v1.Array())
	assert.Equal(t, arrayDataType, v2.dataType)
	assert.Equal(t, []byte(v2.RawMessage), []byte(v1.RawMessage))
}

func TestSetString(t *testing.T) {
	v1 := ParseString(`"batata"`)
	v2 := Parse(nil)
	v2.SetString(v1.String())
	assert.Equal(t, stringDataType, v2.dataType)
	assert.Equal(t, []byte(v2.RawMessage), []byte(v1.RawMessage))
}

func TestSetNumber(t *testing.T) {
	v1 := ParseString(`1`)
	v2 := Parse(nil)
	v2.SetNumber(v1.Number())
	assert.Equal(t, numberDataType, v2.dataType)
	assert.Equal(t, []byte(v2.RawMessage), []byte(v1.RawMessage))
}

func TestSetBoolean(t *testing.T) {
	v1 := ParseString(`true`)
	v2 := Parse(nil)
	v2.SetBoolean(v1.Boolean())
	assert.Equal(t, booleanDataType, v2.dataType)
	assert.Equal(t, []byte(v2.RawMessage), []byte(v1.RawMessage))
}
