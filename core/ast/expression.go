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

//=============================================================================
//===
//=== Expression
//===
//=============================================================================

type Expression interface {
	Eval() (any,error)
	Type() *Type
}

//=============================================================================
//===
//=== Identifier
//===
//=============================================================================

type IdentifierExpression struct {
	name *FQIdentifier
	bar   Expression
}

//=============================================================================

func (e *IdentifierExpression) Eval() (any,error) {
	return nil,nil
}

//=============================================================================

func (e *IdentifierExpression) Type() *Type {
	return nil
}

//=============================================================================

func NewIdentifierExpression(name *FQIdentifier, bar Expression) *IdentifierExpression {
	return &IdentifierExpression{
		name: name,
		bar:  bar,
	}
}

//=============================================================================
//===
//=== Mult
//===
//=============================================================================

type MultExpression struct {
	left  Expression
	right Expression
}

//=============================================================================

func (e *MultExpression) Eval() (any,error) {
	return nil,nil
}

//=============================================================================

func (e *MultExpression) Type() *Type {
	return nil
}

//=============================================================================

func NewMultExpression(left, right Expression) *MultExpression {
	return &MultExpression{
		left:  left,
		right: right,
	}
}

//=============================================================================
//===
//=== Div
//===
//=============================================================================

type DivExpression struct {
	left  Expression
	right Expression
}

//=============================================================================

func (e *DivExpression) Eval() (any,error) {
	return nil,nil
}

//=============================================================================

func (e *DivExpression) Type() *Type {
	return nil
}

//=============================================================================

func NewDivExpression(left, right Expression) *DivExpression {
	return &DivExpression{
		left:  left,
		right: right,
	}
}

//=============================================================================
//===
//=== Add
//===
//=============================================================================

type AddExpression struct {
	left  Expression
	right Expression
}

//=============================================================================

func (e *AddExpression) Eval() (any,error) {
	return nil,nil
}

//=============================================================================

func (e *AddExpression) Type() *Type {
	return nil
}

//=============================================================================

func NewAddExpression(left, right Expression) *AddExpression {
	return &AddExpression{
		left:  left,
		right: right,
	}
}

//=============================================================================
//===
//=== Sub
//===
//=============================================================================

type SubExpression struct {
	left  Expression
	right Expression
}

//=============================================================================

func (e *SubExpression) Eval() (any,error) {
	return nil,nil
}

//=============================================================================

func (e *SubExpression) Type() *Type {
	return nil
}

//=============================================================================

func NewSubExpression(left, right Expression) *SubExpression {
	return &SubExpression{
		left:  left,
		right: right,
	}
}

//=============================================================================
//===
//=== Equal
//===
//=============================================================================

type EqualExpression struct {
	left  Expression
	right Expression
}

//=============================================================================

func (e *EqualExpression) Eval() (any,error) {
	return nil,nil
}

//=============================================================================

func (e *EqualExpression) Type() *Type {
	return nil
}

//=============================================================================

func NewEqualExpression(left, right Expression) *EqualExpression {
	return &EqualExpression{
		left:  left,
		right: right,
	}
}

//=============================================================================
//===
//=== Less or equal
//===
//=============================================================================

type LessOrEqualExpression struct {
	left  Expression
	right Expression
}

//=============================================================================

func (e *LessOrEqualExpression) Eval() (any,error) {
	return nil,nil
}

//=============================================================================

func (e *LessOrEqualExpression) Type() *Type {
	return nil
}

//=============================================================================

func NewLessOrEqualExpression(left, right Expression) *LessOrEqualExpression {
	return &LessOrEqualExpression{
		left:  left,
		right: right,
	}
}

//=============================================================================
//===
//=== Greater or equal
//===
//=============================================================================

type GreaterOrEqualExpression struct {
	left  Expression
	right Expression
}

//=============================================================================

func (e *GreaterOrEqualExpression) Eval() (any,error) {
	return nil,nil
}

//=============================================================================

func (e *GreaterOrEqualExpression) Type() *Type {
	return nil
}

//=============================================================================

func NewGreaterOrEqualExpression(left, right Expression) *GreaterOrEqualExpression {
	return &GreaterOrEqualExpression{
		left:  left,
		right: right,
	}
}

//=============================================================================
//===
//=== Less than
//===
//=============================================================================

type LessThanExpression struct {
	left  Expression
	right Expression
}

//=============================================================================

func (e *LessThanExpression) Eval() (any,error) {
	return nil,nil
}

//=============================================================================

func (e *LessThanExpression) Type() *Type {
	return nil
}

//=============================================================================

func NewLessThanExpression(left, right Expression) *LessThanExpression {
	return &LessThanExpression{
		left:  left,
		right: right,
	}
}

//=============================================================================
//===
//=== Greater than
//===
//=============================================================================

type GreaterThanExpression struct {
	left  Expression
	right Expression
}

//=============================================================================

func (e *GreaterThanExpression) Eval() (any,error) {
	return nil,nil
}

//=============================================================================

func (e *GreaterThanExpression) Type() *Type {
	return nil
}

//=============================================================================

func NewGreaterThanExpression(left, right Expression) *GreaterThanExpression {
	return &GreaterThanExpression{
		left:  left,
		right: right,
	}
}

//=============================================================================
//===
//=== Value
//===
//=============================================================================

type ConstantValueExpression struct {
	Value     any
	ValueType *Type
}

//=============================================================================

func (e *ConstantValueExpression) Eval() (any,error) {
	return nil,nil
}

//=============================================================================

func (e *ConstantValueExpression) Type() *Type {
	return nil
}

//=============================================================================
//===
//=== Function call
//===
//=============================================================================

type FunctionCallExpression struct {
}

//=============================================================================

func (e *FunctionCallExpression) Eval() (any,error) {
	return nil,nil
}

//=============================================================================

func (e *FunctionCallExpression) Type() *Type {
	return nil
}

//=============================================================================
//===
//=== And
//===
//=============================================================================

type AndExpression struct {
	expressions []Expression
}

//=============================================================================

func NewAndExpression(expressions []Expression) *AndExpression {
	return &AndExpression{
		expressions: expressions,
	}
}

//=============================================================================

func (e *AndExpression) Eval() (any,error) {
	for _, cond := range e.expressions {
		eval,err := cond.Eval()
		if err != nil {
			return false,err
		}

		if !eval {
			return false,nil
		}
	}

	return true,nil
}

//=============================================================================

func (e *AndExpression) Type() *Type {
	return nil
}

//=============================================================================
//===
//=== Or
//===
//=============================================================================

type OrExpression struct {
	expressions []Expression
}

//=============================================================================

func NewOrExpression(expressions []Expression) *OrExpression {
	return &OrExpression{
		expressions: expressions,
	}
}

//=============================================================================

func (e *OrExpression) Eval() (any,error) {
	for _, cond := range e.expressions {
		eval,err := cond.Eval()
		if err != nil {
			return false,err
		}

		if !eval {
			return false,nil
		}
	}

	return true,nil
}

//=============================================================================

func (e *OrExpression) Type() *Type {
	return nil
}

//=============================================================================
//===
//=== Not
//===
//=============================================================================

type NotExpression struct {
	expression Expression
}

//=============================================================================

func NewNotExpression(e Expression) *NotExpression {
	return &NotExpression{
		expression: e,
	}
}

//=============================================================================

func (e *NotExpression) Eval() (any,error) {
	res,err := e.expression.Eval()
	if err != nil {
		return false,err
	}

	return !res,nil
}

//=============================================================================

func (e *NotExpression) Type() *Type {
	return nil
}

//=============================================================================
//===
//=== True
//===
//=============================================================================

type TrueExpression struct {
	value *Value
}

//=============================================================================

func NewTrueExpression() *TrueExpression {
	return &TrueExpression{
		value : NewValue(BoolType, true),
	}
}

//=============================================================================

func (e *TrueExpression) Eval() (any,error) {
	return true,nil
}

//=============================================================================

func (e *TrueExpression) Type() *Type {
	return nil
}

//=============================================================================
//===
//=== False
//===
//=============================================================================

type FalseExpression struct {
}

//=============================================================================

func NewFalseExpression() *FalseExpression {
	return &FalseExpression{}
}

//=============================================================================

func (e *FalseExpression) Eval() (any,error) {
	return false,nil
}

//=============================================================================

func (e *FalseExpression) Type() *Type {
	return nil
}

//=============================================================================
