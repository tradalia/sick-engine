//=============================================================================
/*
Copyright © 2025 Andrea Carboni andrea.carboni71@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
//=============================================================================

package expression

import (
	"fmt"

	"github.com/tradalia/sick-engine/datatype"
)

//=============================================================================
//===
//=== Constants
//===
//=============================================================================

var TrueValue  = NewValue( BoolType, true )
var FalseValue = NewValue( BoolType, false)

var TrueValueSet  = NewValueSet(TrueValue)
var FalseValueSet = NewValueSet(FalseValue)

//=============================================================================

func NewIntValue   (i int           ) *Value { return NewValue(IntType,    i) }
func NewRealValue  (r float64       ) *Value { return NewValue(RealType,   r) }
func NewBoolValue  (b bool          ) *Value { return NewValue(BoolType,   b) }
func NewStringValue(s string        ) *Value { return NewValue(StringType, s) }
func NewTimeValue  (t *datatype.Time) *Value { return NewValue(TimeType,   t) }
func NewDateValue  (d *datatype.Date) *Value { return NewValue(DateType,   d) }

//=============================================================================
//===
//=== ValueSet
//===
//=============================================================================

type ValueSet struct {
	values []*Value
}

//=============================================================================

func NewValueSet(v *Value) *ValueSet {
	vs := &ValueSet{}
	vs.AddValue(v)
	return vs
}

//=============================================================================

func (vs *ValueSet) AddValue(value *Value) {
	vs.values = append(vs.values, value)
}

//=============================================================================

func (vs *ValueSet) Size() int {
	return len(vs.values)
}

//=============================================================================
//===
//=== Value
//===
//=============================================================================

type Value struct {
	Type  *Type
	Data  any
}

//=============================================================================

func NewValue(typ *Type, value any) *Value {
	return &Value{
		Type: typ,
		Data: value,
	}
}

//=============================================================================

func (v *Value) Equals(vv *Value) (bool,error) {
	if v.Type != vv.Type {
		return false, fmt.Errorf("expected type %v but got %v", v.Type, vv.Type)
	}

	if v.Data == nil && vv.Data == nil {
		return true, nil
	}

	if v.Data == nil || vv.Data == nil {
		return false, nil
	}

	return v.Data == vv.Data, nil
}

//=============================================================================

func (v *Value) LessThan(vv *Value) (bool,error) {
	if v.Type != vv.Type {
		return false, fmt.Errorf("expected type %v but got %v", v.Type, vv.Type)
	}

	if v.Data == nil || vv.Data == nil {
		return false, nil
	}

	//TODO
	return false, nil
}

//=============================================================================
