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
	"strconv"

	"github.com/antlr4-go/antlr/v4"
	"github.com/tradalia/sick-engine/parser"
)

//=============================================================================

const (
	DefaultPackage = "main"
)

//=============================================================================

func Build(tree parser.IScriptContext) *Script {
	script := NewScript()

	script.PackageName = convertPackage(tree.PackageDef())

	for _, cds := range tree.AllConstantsDef() {
		for _,cd := range cds.AllConstantDef() {
			script.AddConstant(convertConstantDef(cd))
		}
	}

	for _, fd := range tree.AllFunctionDef() {
		script.AddFunction(convertFunctionDef(fd))
	}

	for _, ed := range tree.AllEnumDef() {
		script.AddEnum(convertEnumDef(ed))
	}

	for _, cd := range tree.AllClassDef() {
		script.AddClass(convertClassDef(cd))
	}

	return script
}

//=============================================================================
//===
//=== Top entities
//===
//=============================================================================

func convertPackage(packDef parser.IPackageDefContext) string {
	if packDef != nil {
		return packDef.IDENTIFIER().GetText()
	}

	return DefaultPackage
}

//=============================================================================

func convertConstantDef(tree parser.IConstantDefContext) *Constant {
	name  := tree.IDENTIFIER().GetText()
	value := tree.ConstantValue()

	c := &Constant{
		Name : name,
	}

	intVal  := value.INT_VALUE()
	realVal := value.REAL_VALUE()
	timeVal := value.TimeValue()
	dateVal := value.DateValue()
	strVal  := value.STRING_VALUE()
	boolVal := value.BoolValue()

	var err error

	if intVal != nil {
		c.Value,err = strconv.Atoi(intVal.GetText())
		if err != nil {
			raiseError(tree.GetParser(), "Invalid integer value : " + intVal.GetText())
		}
	} else if realVal != nil {
		c.Value,err = strconv.ParseFloat(realVal.GetText(), 64)
		if err != nil {
			raiseError(tree.GetParser(), "Invalid real value : " + realVal.GetText())
		}
	} else if timeVal != nil {
		hh,_ := strconv.Atoi(timeVal.INT_VALUE(0).GetText())
		mm,_ := strconv.Atoi(timeVal.INT_VALUE(1).GetText())

		//TODO: Convert to Time object
		_ = hh + mm
	} else if dateVal != nil {
		y,_ := strconv.Atoi(timeVal.INT_VALUE(0).GetText())
		m,_ := strconv.Atoi(timeVal.INT_VALUE(1).GetText())
		d,_ := strconv.Atoi(timeVal.INT_VALUE(2).GetText())

		//TODO: Convert to a Date object
		_ = y + m + d
	} else if strVal != nil {
		text := strVal.GetText()
		c.Value = text[1:len(text)-1]
	} else if boolVal != nil {
		c.Value = boolVal.GetText() == "true"
	} else {
		panic("Unknown constant value : " + value.GetText())
	}

	return c
}

//=============================================================================

func convertFunctionDef(tree parser.IFunctionDefContext) *Function {
	name := tree.IDENTIFIER().GetText()
	f := NewFunction(name)

	if tree.Class() != nil {
		f.class = tree.Class().GetText()
	}

	convertFunctionParams (f, tree.Parameters())
	convertFunctionResults(f, tree.Results())
	f.block = convertStatementBlock(tree.Block())

	return f
}

//=============================================================================

func convertFunctionParams(f *Function, params parser.IParametersContext) {
	for _, pd := range params.AllParameterDecl() {
		name := pd.IDENTIFIER().GetText()
		typ  := pd.ParamType().GetText()

		t := &Type {
			Name: typ,
		}
		p := NewParam(name,t)
		f.AddParam(p)
	}
}

//=============================================================================

func convertFunctionResults(f *Function, results parser.IResultsContext) {
	for _,r := range results.AllParamType() {
		t := &Type{
			Name: r.GetText(),
		}

		f.AddReturnType(t)
	}
}

//=============================================================================

func convertEnumDef(tree parser.IEnumDefContext) *Enum {
	name := tree.IDENTIFIER().GetText()
	e := NewEnum(name)

	for _, ei := range tree.AllEnumItem() {
		e.AddItem(convertEnumItem(ei))
	}

	return e
}

//=============================================================================

func convertEnumItem(tree parser.IEnumItemContext) *EnumItem {
	name  := tree.IDENTIFIER().GetText()
	code  := -1
	value := ""

	enumVal := tree.EnumValue()
	if enumVal != nil {
		intValue := enumVal.INT_VALUE()
		strValue := enumVal.STRING_VALUE()

		if intValue != nil {
			iValue,err := strconv.Atoi(intValue.GetText())
			if err != nil {
				raiseError(tree.GetParser(), "Invalid integer value for enum : " + intValue.GetText())
			}

			code = iValue
		}
		if strValue != nil {
			text := strValue.GetText()
			value = text[1:len(text)-1]
		}
	}
	return NewEnumItem(name, code, value)
}

//=============================================================================

func convertClassDef(tree parser.IClassDefContext) *Class {
	name := tree.IDENTIFIER().GetText()
	c := NewClass(name)

	for _, p := range tree.AllProperty() {
		c.AddProperty(convertProperty(p))
	}

	return c
}

//=============================================================================

func convertProperty(tree parser.IPropertyContext) *Property {
	name := tree.IDENTIFIER().GetText()
	typ  := tree.ParamType().GetText()

	t := &Type {
		Name: typ,
	}

	return NewProperty(name,t)
}

//=============================================================================

func convertStatementBlock(tree parser.IBlockContext) *StatementBlock {
	block := &StatementBlock{}

	for _,stmt := range tree.AllStatement() {
		varDecl := stmt.VarDeclaration()
		ifStmt  := stmt.IfStatement()
		retStmt := stmt.ReturnStatement()

		if varDecl != nil {
			block.Add(convertVarDeclaration(varDecl))
		} else if ifStmt != nil {
			block.Add(convertIfStatement(ifStmt))
		} else if retStmt != nil {
			block.Add(convertReturnStatement(retStmt))
		} else {
			panic("Unknown statement type : " + stmt.GetText())
		}
	}

	return block
}

//=============================================================================

func convertVarDeclaration(tree parser.IVarDeclarationContext) *VarDeclaration {
	name := convertFQIdentifier(tree.FqIdentifier())
	expr := convertExpression(tree.Expression())

	vd := &VarDeclaration{
		FQName     : name,
		Expression : expr,
		//TODO: infer type
	}

	return vd
}

//=============================================================================

func convertFQIdentifier(tree parser.IFqIdentifierContext) *FQIdentifier {
	fqi := NewFQIdentifier()

	for _, s := range tree.AllIDENTIFIER() {
		fqi.AddScope(s.GetText())
	}

	return fqi
}

//=============================================================================

func convertIfStatement(tree parser.IIfStatementContext) *IfStatement {
	is    := NewIfStatement()
	cond  := convertExpression(tree.Expression())
	block := convertStatementBlock(tree.Block())
	is.AddConditionalBlock(NewConditionalBlock(cond, block))

	for _, eib := range tree.AllElseIfBlock() {
		cond  = convertExpression(eib.Expression())
		block = convertStatementBlock(eib.Block())

		is.AddConditionalBlock(NewConditionalBlock(cond, block))
	}

	if tree.ElseBlock() != nil {
		cond  = NewTrueExpression()
		block = convertStatementBlock(tree.ElseBlock().Block())

		is.AddConditionalBlock(NewConditionalBlock(cond, block))
	}

	return is
}

//=============================================================================

func convertReturnStatement(tree parser.IReturnStatementContext) *ReturnStatement {
	expr := convertExpression(tree.Expression())
	return NewReturnStatement(expr)
}

//=============================================================================
//TODO
func convertFunctionCall(tree parser.IFunctionCallContext) Expression {

}

//=============================================================================

func convertExpression(tree parser.IExpressionContext) Expression {
	if tree == nil {
		return nil
	}

	//--- First, try with an arithmetic expression

	ae := convertArithmeticExpression(tree)
	if ae != nil {
		return ae
	}

	//--- Then try with a relational expression

	re := convertRelationalExpression(tree)
	if re != nil {
		return re
	}

	//--- Now, try with AND/OR

	be := convertBooleanExpression(tree)
	if be != nil {
		return be
	}

	//--- Lastly, other possibilities

	identExpr := tree.IdentifierExpression()
	exprParen := tree.ExpressionInParenthesis()
	unaryExpr := tree.UnaryExpression()
	funCall   := tree.FunctionCall()
	constValue:= tree.ConstantValue()

	if identExpr != nil {
		return convertIdentifierExpression(identExpr)
	} else if exprParen != nil {
		return convertExpression(exprParen.Expression())
	} else if unaryExpr != nil {
		return convertUnaryExpression(unaryExpr)
	} else if funCall != nil {
		return convertFunctionCall(funCall)
	} else if constValue != nil {
		return convertConstantValue(constValue)
	} else {
		panic("Unknown expression type : " + tree.GetText())
	}
}

//=============================================================================

func convertArithmeticExpression(tree parser.IExpressionContext) Expression {
	if tree.STAR() != nil {
		return NewMultExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	if tree.SLASH() != nil {
		return NewDivExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	if tree.PLUS() != nil {
		return NewAddExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}
	if tree.MINUS() != nil {
		return NewSubExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	return nil
}

//=============================================================================

func convertRelationalExpression(tree parser.IExpressionContext) Expression {
	if tree.EQUAL() != nil {
		return NewEqualExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	if tree.LESS_OR_EQUAL() != nil {
		return NewLessOrEqualExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	if tree.GREATER_OR_EQUAL() != nil {
		return NewGreaterOrEqualExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	if tree.LESS_THAN() != nil {
		return NewLessThanExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	if tree.GREATER_THAN() != nil {
		return NewGreaterThanExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	if tree.NOT_EQUAL() != nil {
		//TODO
	}

	return nil
}

//=============================================================================

func convertBooleanExpression(tree parser.IExpressionContext) Expression {
	if tree.AllAND() != nil {
		expr := buildExpressionList(tree.AllExpression())
		return NewAndExpression(expr)
	}

	if tree.AllOR() != nil {
		expr := buildExpressionList(tree.AllExpression())
		return NewOrExpression(expr)
	}

	return nil
}

//=============================================================================

func buildExpressionList(list []parser.IExpressionContext) []Expression {
	var res []Expression

	for _, expr := range list {
		res = append(res, convertExpression(expr))
	}

	return res
}

//=============================================================================

func convertIdentifierExpression(tree parser.IIdentifierExpressionContext) Expression {
	name := convertFQIdentifier(tree.FqIdentifier())
	bar  := convertExpression(tree.Expression())

	return NewIdentifierExpression(name, bar)
}

//=============================================================================

func convertUnaryExpression(tree parser.IUnaryExpressionContext) Expression {

}

//=============================================================================

func convertConstantValue(tree parser.IConstantValueContext) Expression {

}

//=============================================================================

func raiseError(p antlr.Parser, message string) {
	p.NotifyErrorListeners(message, nil, nil)
}

//=============================================================================
