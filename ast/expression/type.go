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

//=============================================================================

const (
	TypeToFindout  = -1
	TypeInt        =  0
	TypeReal       =  1
	TypeBool       =  2
	TypeString     =  3
	TypeTime       =  4
	TypeDate       =  5
	TypeTimeseries =  6
	TypeEnum       =  7
	TypeClass      =  8
	TypeList       =  9
	TypeMap        = 10
	TypeError      = 11
)

//=============================================================================

type Type struct {
	Code int8
	Name string
}

//=============================================================================

var IntType        = &Type{ Code: TypeInt,        Name: "int",        }
var RealType       = &Type{ Code: TypeReal,       Name: "real",       }
var BoolType       = &Type{ Code: TypeBool,       Name: "bool",       }
var StringType     = &Type{ Code: TypeString,     Name: "string",     }
var TimeType       = &Type{ Code: TypeTime,       Name: "time",       }
var DateType       = &Type{ Code: TypeDate,       Name: "date",       }
var TimeseriesType = &Type{ Code: TypeTimeseries, Name: "timeseries", }
var ErrorType      = &Type{ Code: TypeError,      Name: "error",      }

//=============================================================================

func NewToFindoutType(typeOrFQId string) *Type {
	return &Type{
		Code: TypeToFindout,
		Name: typeOrFQId,
	}
}

//=============================================================================

func NewEnumType(name string) *Type {
	return &Type{
		Code: TypeEnum,
		Name: name,
	}
}

//=============================================================================

func NewClassType(name string) *Type {
	return &Type{
		Code: TypeClass,
		Name: name,
	}
}

//=============================================================================

func NewListType() *Type {
	return &Type{
		Code: TypeList,
		Name: "list",
	}
}

//=============================================================================

func NewMapType() *Type {
	return &Type{
		Code: TypeMap,
		Name: "map",
	}
}

//=============================================================================
