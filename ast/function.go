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

package ast

import (
	"strings"

	"github.com/tradalia/sick-engine/ast/statement"
	"github.com/tradalia/sick-engine/datatype"
)

//=============================================================================
//===
//=== Function
//===
//=============================================================================

type Function struct {
	name     string
	class    string
	params   []*Param
	returns  []datatype.Type
	block    *statement.Block
}

//=============================================================================

func NewFunction(name string) *Function {
	return &Function{
		name : name,
	}
}
//=============================================================================

func (f *Function) Id() string {
	sb := strings.Builder{}
	sb.WriteString(f.name)
	sb.WriteString("|")

	for _,p := range f.params {
		sb.WriteString("|")
		sb.WriteString(p.Type.String())
	}

	return sb.String()
}

//=============================================================================

func (f *Function) AddParam(p *Param) {
	f.params = append(f.params, p)
}

//=============================================================================

func (f *Function) AddReturnType(t datatype.Type) {
	f.returns = append(f.returns, t)
}

//=============================================================================
//===
//=== Param
//===
//=============================================================================

type Param struct {
	Name  string
	Type  datatype.Type
}

//=============================================================================

func NewParam(name string, t datatype.Type) *Param {
	return &Param{
		Name: name,
		Type: t,
	}
}

//=============================================================================
