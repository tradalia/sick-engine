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

package core

import (
	"github.com/tradalia/sick-engine/ast"
	"github.com/tradalia/sick-engine/ast/expression"
)

//=============================================================================
//===
//=== Environment
//===
//=============================================================================

type Environment struct {
	constants map[string]*ast.Constant
	functions map[string]*ast.Function
	enums     map[string]*ast.Enum
	classes   map[string]*ast.Class
}

//=============================================================================

func NewEnvironment() *Environment {
	return &Environment{
		constants: make(map[string]*ast.Constant),
		functions: make(map[string]*ast.Function),
		enums    : make(map[string]*ast.Enum),
		classes  : make(map[string]*ast.Class),
	}
}

//=============================================================================

func (e *Environment) AddScript(s *ast.Script) *ParseErrors {
	pes := NewParseErrors(s.Filename)

	for _,c := range s.Constants {
		e.addConstant(s.PackageName, c, pes)
	}

	for _,f := range s.Functions {
		e.addFunction(s.PackageName, f, pes)
	}

	for _,en := range s.Enums {
		e.addEnum(s.PackageName, en, pes)
	}

	for _,c := range s.Classes {
		e.addClass(s.PackageName, c, pes)
	}

	return pes
}

//=============================================================================

func (e *Environment) addConstant(pac string, c *ast.Constant, pes *ParseErrors) {

	fqn := pac + "." + c.Name

	if _, ok := e.constants[fqn]; ok {
		pe := NewParseError(-1, -1, "constant name already exists: " + fqn)
		pes.AddError(pe)
		return
	}

	//pe := checkValueValidity(c.Value)
	//if pe != nil {
	//	pes.AddError(pe)
	//	return
	//}

	e.constants[fqn] = c
}

//=============================================================================

func (e *Environment) addFunction(pac string, f *ast.Function, pes *ParseErrors) {

}

//=============================================================================

func (e *Environment) addEnum(pac string, en *ast.Enum, pes *ParseErrors) {

}

//=============================================================================

func (e *Environment) addClass(pac string, c *ast.Class, pes *ParseErrors) {

}

//=============================================================================

func checkValueValidity(v *expression.Value) *ParseError {
	return nil
}

//=============================================================================

