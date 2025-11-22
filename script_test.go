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

package main

import (
	"strings"
	"testing"

	"github.com/tradalia/sick-engine/core"
)

//=============================================================================
//=== Package
//=============================================================================

func TestPackageNameNotReserved(t *testing.T) {
	_, pes := ParseString("package bar")
	assertError(t, pes, "reserved")
}

//=============================================================================

func TestPackageNameInLowercase(t *testing.T) {
	_, pes := ParseString("package A")
	assertError(t, pes, "lowercase")
}

//=============================================================================
//=== Constants
//=============================================================================

func TestConstantNameInUppercase(t *testing.T) {
	_, pes := ParseString("const { a = 4 }")
	assertError(t, pes, "uppercase")
}

//=============================================================================
//=== Functions
//=============================================================================

func TestFunctionNameNotReserved(t *testing.T) {
	_, pes := ParseString("func bar() {}")
	assertError(t, pes, "reserved")
}

//=============================================================================

func TestFunctionNameInLowercase(t *testing.T) {
	_, pes := ParseString("func A() {}")
	assertError(t, pes, "lowercase")
}

//=============================================================================

func TestFunctionClassNameInUppercase(t *testing.T) {
	_, pes := ParseString("func (b) a() {}")
	assertError(t, pes, "uppercase")
}

//=============================================================================
//=== Enums
//=============================================================================

func TestEnumNameInUppercase(t *testing.T) {
	_, pes := ParseString("enum a { B }")
	assertError(t, pes, "uppercase")
}

//=============================================================================

func TestEnumItemNameInUppercase(t *testing.T) {
	_, pes := ParseString("enum A { b }")
	assertError(t, pes, "uppercase")
}

//=============================================================================

func TestEnumWithMicexItems(t *testing.T) {
	_, pes := ParseString("enum A { B(1) C(\"x\") }")
	assertError(t, pes, "mixing")
}

//=============================================================================

func TestEnumWithAutoCodeAssignment(t *testing.T) {
	s, pes := ParseString("enum A { B C }")
	if !pes.IsEmpty() {
		t.Errorf("Errors returned")
	}

	e := s.Enums[0]
	if e.Items()[0].Code != 1 && e.Items()[1].Code != 2 {
		t.Errorf("Bad auto assignment")
	}
}

//=============================================================================
//=== Classes
//=============================================================================

func TestClassNameInUppercase(t *testing.T) {
	_, pes := ParseString("class a { }")
	assertError(t, pes, "uppercase")
}

//=============================================================================

func TestPropertyNameInLowercase(t *testing.T) {
	_, pes := ParseString("class A { B int }")
	assertError(t, pes, "lowercase")
}

//=============================================================================
//=== Identifiers
//=============================================================================

func TestIdentifierNameInLowercase(t *testing.T) {
	_, pes := ParseString("func a() { B=5 }")
	assertError(t, pes, "lowercase")

	_, pes = ParseString("func a() { a.B=5 }")
	assertError(t, pes, "lowercase")
}

//=============================================================================
//===
//=== Private functions
//===
//=============================================================================

func assertError(t *testing.T, pes *core.ParseErrors, expected string) {
	if pes.IsEmpty() {
		t.Errorf("No error returned")
	} else if !strings.Contains(pes.Errors[0].Error, expected) {
		t.Errorf("Error should contain: %s", expected)
	}
}

//=============================================================================
