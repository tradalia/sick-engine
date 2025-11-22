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
	"github.com/tradalia/sick-engine/ast/expression"
	"github.com/tradalia/sick-engine/ast/statement"
	"github.com/tradalia/sick-engine/data"
	"github.com/tradalia/sick-engine/datatype"
	"github.com/tradalia/sick-engine/parser"
	"github.com/tradalia/sick-engine/tool"
)

//=============================================================================

const (
	DefaultPackage = "main"
)

//=============================================================================

var ReservedIdentifiers = map[string]bool{
	"bar"   : true,
	"trade" : true,
	"open"  : true,
	"close" : true,
	"high"  : true,
	"low"   : true,
}

//=============================================================================

func Build(tree parser.IScriptContext) *Script {
	script := NewScript(tree.GetParser().GetInputStream().GetSourceName())

	pack := convertPackage(tree.PackageDef())
	script.PackageName = pack

	for _, cds := range tree.AllConstantsDef() {
		for _,cd := range cds.AllConstantDef() {
			c := convertConstantDef(cd)
			if !script.AddConstant(c) {
				raiseError(tree.GetParser(), "constant's name already used as identifier: "+ c.Name)
			}
		}
	}

	for _, vds := range tree.AllVariablesDef() {
		for _,vd := range vds.AllVariableDef() {
			v := convertVariableDef(vd)
			if !script.AddVariable(v) {
				raiseError(tree.GetParser(), "variable's name already used as identifier: "+ v.Name)
			}
		}
	}

	for _, fd := range tree.AllFunctionDef() {
		f := convertFunctionDef(fd)
		if !script.AddFunction(f) {
			raiseError(tree.GetParser(), "function's name already used as identifier: "+ f.name)
		}
	}

	for _, ed := range tree.AllEnumDef() {
		e := convertEnumDef(ed)
		if !script.AddEnum(e) {
			raiseError(tree.GetParser(), "enum's name already used as identifier: "+ e.Name())
		}
	}

	for _, cd := range tree.AllClassDef() {
		c := convertClassDef(cd)
		if !script.AddClass(c) {
			raiseError(tree.GetParser(), "class's name already used as identifier: "+ c.Name())
		}
	}

	return script
}

//=============================================================================
//===
//=== Top entities
//===
//=============================================================================

func convertPackage(tree parser.IPackageDefContext) string {
	if tree != nil {
		id := tree.IDENTIFIER().GetText()

		if _,ok := ReservedIdentifiers[id]; ok {
			raiseError(tree.GetParser(), "package's name is reserved: "+ id)
		}

		if !tool.StartsWithLowerCase(id) {
			raiseError(tree.GetParser(), "package's name must start with a lowercase character: "+ id)
		}

		return id
	}

	return DefaultPackage
}

//=============================================================================

func convertConstantDef(tree parser.IConstantDefContext) *Constant {
	name  := tree.IDENTIFIER().GetText()
	value := tree.Expression()

	if !tool.StartsWithUpperCase(name) {
		raiseError(tree.GetParser(), "constant's name must start with an uppercase character: "+ name)
	}

	return NewConstant(name, convertExpression(value))
}

//=============================================================================

func convertVariableDef(tree parser.IVariableDefContext) *Variable {
	name  := tree.IDENTIFIER().GetText()
	value := tree.Expression()

	if _,ok := ReservedIdentifiers[name]; ok {
		raiseError(tree.GetParser(), "variable's name is reserved: "+ name)
	}

	if !tool.StartsWithLowerCase(name) {
		raiseError(tree.GetParser(), "variable's name must start with a lowercase character: "+ name)
	}

	return NewVariable(name, convertExpression(value))
}

//=============================================================================

func convertFunctionDef(tree parser.IFunctionDefContext) *Function {
	name := tree.IDENTIFIER().GetText()

	if _,ok := ReservedIdentifiers[name]; ok {
		raiseError(tree.GetParser(), "function's name is reserved: "+ name)
	}

	if !tool.StartsWithLowerCase(name) {
		raiseError(tree.GetParser(), "function's name must start with a lowercase character: "+ name)
	}

	f := NewFunction(name)

	if tree.Class() != nil {
		f.class = tree.Class().GetText()
		if !tool.StartsWithUpperCase(f.class) {
			raiseError(tree.GetParser(), "function's class name must start with an uppercase character: "+ f.class)
		}
	}

	convertFunctionParams (f, tree.Parameters())
	convertFunctionResults(f, tree.Results())
	f.block = convertStatementBlock(tree.Block())

	return f
}

//=============================================================================

func convertFunctionParams(f *Function, params parser.IParametersContext) {
	names := map[string]bool{}

	for _, pd := range params.AllParameterDecl() {
		name  := pd.IDENTIFIER().GetText()
		type_ := pd.Type_().GetText()

		if _,ok := ReservedIdentifiers[name]; ok {
			raiseError(params.GetParser(), "function's parameter is reserved: "+ name)
		}

		if !tool.StartsWithLowerCase(name) {
			raiseError(params.GetParser(), "function's parameter must start with a lowercase character: "+ name)
		}

		if _,ok := names[name]; ok {
			raiseError(params.GetParser(), "function's parameter name is duplicated: "+ name)
		}

		names[name] = true
		t := datatype.NewToFindOutType(type_)
		p := NewParam(name,t)
		f.AddParam(p)
	}
}

//=============================================================================

func convertFunctionResults(f *Function, results parser.IResultsContext) {
	if results != nil {
		for _,r := range results.AllType_() {
			t := datatype.NewToFindOutType(r.GetText())
			f.AddReturnType(t)
		}
	}
}

//=============================================================================

func convertEnumDef(tree parser.IEnumDefContext) *datatype.EnumType {
	name := tree.IDENTIFIER().GetText()
	e := datatype.NewEnumType(name)

	if !tool.StartsWithUpperCase(name) {
		raiseError(tree.GetParser(), "enum's name must start with an uppercase character: "+ name)
	}

	countInts := 0
	countStr  := 0

	for _, ei := range tree.AllEnumItem() {
		i := convertEnumItem(ei)
		if !e.AddItem(i) {
			raiseError(tree.GetParser(), "enum's item already used as identifier: "+ i.Name)
		}

		if i.Code != -1 {
			countInts++
		}

		if i.Value != "" {
			countStr++
		}
	}

	badInts := countInts > 0 && countInts != e.Size()
	badStr  := countStr  > 0 && countStr  != e.Size()

	if badInts && badStr {
		raiseError(tree.GetParser(), "mixing integers and strings in enum: "+ name)
	}

	if countInts == 0 && countStr == 0 {
		e.AssignCodes()
	}

	return e
}

//=============================================================================

func convertEnumItem(tree parser.IEnumItemContext) *datatype.EnumItem {
	name  := tree.IDENTIFIER().GetText()
	code  := -1
	value := ""

	if !tool.StartsWithUpperCase(name) {
		raiseError(tree.GetParser(), "enum item's name must start with an uppercase character: "+ name)
	}

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
	return datatype.NewEnumItem(name, code, value)
}

//=============================================================================

func convertClassDef(tree parser.IClassDefContext) *datatype.ClassType {
	name := tree.IDENTIFIER().GetText()
	c := datatype.NewClassType(name)

	if !tool.StartsWithUpperCase(name) {
		raiseError(tree.GetParser(), "class's name must start with an uppercase character: "+ name)
	}

	for _, p := range tree.AllProperty() {
		c.AddProperty(convertProperty(p))
	}

	return c
}

//=============================================================================

func convertProperty(tree parser.IPropertyContext) *datatype.Property {
	name  := tree.IDENTIFIER().GetText()
	type_ := tree.Type_().GetText()

	if _,ok := ReservedIdentifiers[name]; ok {
		raiseError(tree.GetParser(), "class's property is reserved: "+ name)
	}

	if !tool.StartsWithLowerCase(name) {
		raiseError(tree.GetParser(), "class's property must start with a lowercase character: "+ name)
	}

	t := datatype.NewToFindOutType(type_)

	return datatype.NewProperty(name,t)
}

//=============================================================================
//===
//=== Statements
//===
//=============================================================================

func convertStatementBlock(tree parser.IBlockContext) *statement.Block {
	block := &statement.Block{}

	for _,stmt := range tree.AllStatement() {
		varDecl := stmt.VarDeclaration()
		ifStmt  := stmt.IfStatement()
		callStmt:= stmt.FunctionCall()
		retStmt := stmt.ReturnStatement()

		if varDecl != nil {
			block.Add(convertVarDeclaration(varDecl))
		} else if ifStmt != nil {
			block.Add(convertIfStatement(ifStmt))
		} else if callStmt != nil {
			block.Add(convertFunctionCallStatement(callStmt))
		} else if retStmt != nil {
			block.Add(convertReturnStatement(retStmt))
		} else {
			panic("Unknown statement type : " + stmt.GetText())
		}
	}

	return block
}

//=============================================================================

func convertVarDeclaration(tree parser.IVarDeclarationContext) *statement.VarDeclaration {

	vd := statement.NewVarDeclaration()

	for _, accId := range tree.AllAccessedIdentifier() {
		name := convertFQIdentifier(accId.FqIdentifier())

		var acc expression.Expression

		if accId.AccessorExpression() != nil {
			acc = convertExpression(accId.AccessorExpression().Expression())
		}

		vd.AddIdentifier(expression.NewIdentifierExpression(name, acc))
	}

	for _, e := range tree.AllExpression() {
		expr := convertExpression(e)
		vd.AddExpression(expr)
	}

	return vd
}

//=============================================================================

func convertIfStatement(tree parser.IIfStatementContext) *statement.IfStatement {
	is    := statement.NewIfStatement()
	cond  := convertExpression(tree.Expression())
	block := convertStatementBlock(tree.Block())
	is.AddConditionalBlock(statement.NewConditionalBlock(cond, block))

	for _, eib := range tree.AllElseIfBlock() {
		cond  = convertExpression(eib.Expression())
		block = convertStatementBlock(eib.Block())

		is.AddConditionalBlock(statement.NewConditionalBlock(cond, block))
	}

	if tree.ElseBlock() != nil {
		cond  = expression.NewTrueExpression()
		block = convertStatementBlock(tree.ElseBlock().Block())

		is.AddConditionalBlock(statement.NewConditionalBlock(cond, block))
	}

	return is
}

//=============================================================================

func convertReturnStatement(tree parser.IReturnStatementContext) *statement.ReturnStatement {
	rs := statement.NewReturnStatement()

	for _, e := range tree.AllExpression() {
		expr := convertExpression(e)
		rs.AddExpression(expr)
	}

	return rs
}

//=============================================================================

func convertFunctionCallStatement(tree parser.IFunctionCallContext) *statement.FunctionCallStatement {
	fqid := convertFQIdentifier(tree.FqIdentifier())
	s := statement.NewFunctionCallStatement(fqid)

	for _,e := range tree.ParamsExpression().AllExpression() {
		expr := convertExpression(e)
		s.AddExpression(expr)
	}

	return s
}

//=============================================================================
//===
//=== Expressions
//===
//=============================================================================

func convertExpression(tree parser.IExpressionContext) expression.Expression {
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
	constValue:= tree.ConstantValueExpression()

	if identExpr != nil {
		return convertIdentifierExpression(identExpr)
	} else if exprParen != nil {
		return convertExpression(exprParen.Expression())
	} else if unaryExpr != nil {
		return convertUnaryExpression(unaryExpr)
	} else if constValue != nil {
		return convertConstantValueExpression(constValue)
	} else {
		panic("Unknown expression type : " + tree.GetText())
	}
}

//=============================================================================

func convertArithmeticExpression(tree parser.IExpressionContext) expression.Expression {
	if tree.STAR() != nil {
		return expression.NewMultExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	if tree.SLASH() != nil {
		return expression.NewDivExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	if tree.PLUS() != nil {
		return expression.NewAddExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}
	if tree.MINUS() != nil {
		return expression.NewSubExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	return nil
}

//=============================================================================

func convertRelationalExpression(tree parser.IExpressionContext) expression.Expression {
	if tree.EQUAL() != nil {
		return expression.NewEqualExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	if tree.LESS_OR_EQUAL() != nil {
		return expression.NewLessOrEqualExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	if tree.GREATER_OR_EQUAL() != nil {
		return expression.NewGreaterOrEqualExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	if tree.LESS_THAN() != nil {
		return expression.NewLessThanExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	if tree.GREATER_THAN() != nil {
		return expression.NewGreaterThanExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	if tree.NOT_EQUAL() != nil {
		return expression.NewNotEqualExpression(convertExpression(tree.Expression(0)), convertExpression(tree.Expression(1)))
	}

	return nil
}

//=============================================================================

func convertBooleanExpression(tree parser.IExpressionContext) expression.Expression {
	if tree.AllAND() != nil {
		expr := buildExpressionList(tree.AllExpression())
		return expression.NewAndExpression(expr)
	}

	if tree.AllOR() != nil {
		expr := buildExpressionList(tree.AllExpression())
		return expression.NewOrExpression(expr)
	}

	return nil
}

//=============================================================================

func buildExpressionList(list []parser.IExpressionContext) []expression.Expression {
	var res []expression.Expression

	for _, expr := range list {
		res = append(res, convertExpression(expr))
	}

	return res
}

//=============================================================================

func convertIdentifierExpression(tree parser.IIdentifierExpressionContext) expression.Expression {
	name     := convertFQIdentifier(tree.FqIdentifier())
	accessor := tree.AccessorExpression()
	params   := tree.ParamsExpression()

	if params != nil {
		list := buildExpressionList(params.AllExpression())
		return expression.NewFunctionCallExpression(name, list)
	}

	var acc expression.Expression

	if accessor != nil {
		acc = convertExpression(accessor.Expression())
	}

	return expression.NewIdentifierExpression(name, acc)
}

//=============================================================================

func convertUnaryExpression(tree parser.IUnaryExpressionContext) expression.Expression {
	e := convertExpression(tree.Expression())

	if tree.NOT() != nil {
		return expression.NewNotExpression(e)
	}

	if tree.PLUS() != nil {
		return e
	}

	if tree.MINUS() != nil {
		//TODO
		panic("unimplemented")
	}

	panic("Unknown unary expression type : " + tree.GetText())
}

//=============================================================================

func convertConstantValueExpression(tree parser.IConstantValueExpressionContext) expression.Expression {
	value := convertConstantValue(tree)
	return expression.NewConstantValueExpression(value)
}

//=============================================================================

func convertConstantValue(tree parser.IConstantValueExpressionContext) expression.Value {
	//--- Integers

	intVal := tree.INT_VALUE()
	if intVal != nil {
		return convertConstantInt(tree.GetParser(), intVal)
	}

	//--- Reals

	realVal := tree.REAL_VALUE()
	if realVal != nil {
		return convertConstantReal(tree.GetParser(), realVal)
	}

	//--- Time objects

	timeVal := tree.TimeValue()
	if timeVal != nil {
		return convertConstantTime(tree.GetParser(), timeVal)
	}

	//--- Date objects

	dateVal := tree.DateValue()
	if dateVal != nil {
		return convertConstantDate(tree.GetParser(), dateVal)
	}

	//--- Strings

	strVal  := tree.STRING_VALUE()
	if strVal != nil {
		return convertConstantString(strVal)
	}

	//--- Booleans

	boolVal := tree.BoolValue()
	if boolVal != nil {
		v := boolVal.GetText() == "true"

		return expression.NewBoolValue(v)
	}

	//--- Errors

	errVal := tree.ErrorValue()
	if errVal != nil {
		text := errVal.STRING_VALUE().GetText()
		text  = text[1:len(text)-1]

		return expression.NewErrorValue(text)
	}

	//--- Lists

	listVal := tree.ListValue()
	if listVal != nil {
		return convertList(listVal)
	}

	//--- Maps

	mapVal := tree.MapValue()
	if mapVal != nil {
		return convertMap(mapVal)
	}

	//--- Unknown

	panic("Unknown constant expression type : " + tree.GetText())
	return nil
}

//=============================================================================

func convertConstantInt(parser antlr.Parser, val antlr.TerminalNode) expression.Value {
	i,err := strconv.Atoi(val.GetText())
	if err != nil {
		raiseError(parser, "invalid integer value : " + val.GetText())
	}

	return expression.NewIntValue(int64(i))
}

//=============================================================================

func convertConstantReal(parser antlr.Parser, val antlr.TerminalNode) expression.Value {
	f,err := strconv.ParseFloat(val.GetText(), 64)
	if err != nil {
		raiseError(parser, "Invalid real value : " + val.GetText())
	}

	return expression.NewRealValue(f)
}

//=============================================================================

func convertConstantTime(parser antlr.Parser, val parser.ITimeValueContext) expression.Value {
	hh,_ := strconv.Atoi(val.INT_VALUE(0).GetText())
	mm,_ := strconv.Atoi(val.INT_VALUE(1).GetText())

	v := data.NewTime(hh, mm)

	if !v.IsValid() {
		raiseError(parser, "Invalid time value : " + val.GetText())
	}

	return expression.NewTimeValue(v)
}

//=============================================================================

func convertConstantDate(parser antlr.Parser, val parser.IDateValueContext) expression.Value {
	y,_ := strconv.Atoi(val.INT_VALUE(0).GetText())
	m,_ := strconv.Atoi(val.INT_VALUE(1).GetText())
	d,_ := strconv.Atoi(val.INT_VALUE(2).GetText())

	v := data.NewDate(y, m, d)

	if !v.IsValid() {
		raiseError(parser, "Invalid date value : " + val.GetText())
	}

	return expression.NewDateValue(v)
}

//=============================================================================

func convertConstantString(val antlr.TerminalNode) expression.Value {
	text := val.GetText()
	text  = text[1:len(text)-1]

	return expression.NewStringValue(text)
}

//=============================================================================

func convertList(listVal parser.IListValueContext) expression.Value {
	t  := convertType(listVal.ListType().Type_())
	lv := expression.NewListValue(t)

	if listVal.InitialListValues() != nil {
		for _,e := range listVal.InitialListValues().AllExpression() {
			lv.AddExpression(convertExpression(e))
		}
	}

	return lv
}

//=============================================================================

func convertMap(mapVal parser.IMapValueContext) expression.Value {
	kt := datatype.NewToFindOutType(mapVal.MapType().KeyType().GetText())
	vt := convertType(mapVal.MapType().Type_())
	mv := expression.NewMapValue(kt, vt)

	if mapVal.InitialMapValues() != nil {
		for _,me := range mapVal.InitialMapValues().AllKeyValueCouple() {
			k := convertKeyValue  (me.KeyValue())
			e := convertExpression(me.Expression())
			mv.Set(k, e)
		}
	}

	return mv
}

//=============================================================================

func convertKeyValue(tree parser.IKeyValueContext) expression.Value{
	//--- Integers

	intVal  := tree.INT_VALUE()
	if intVal != nil {
		return convertConstantInt(tree.GetParser(), intVal)
	}

	//--- Time objects

	timeVal := tree.TimeValue()
	if timeVal != nil {
		return convertConstantTime(tree.GetParser(), timeVal)
	}

	//--- Date objects

	dateVal := tree.DateValue()
	if dateVal != nil {
		return convertConstantDate(tree.GetParser(), dateVal)
	}

	//--- Strings

	strVal  := tree.STRING_VALUE()
	if strVal != nil {
		return convertConstantString(strVal)
	}

	panic("Unknown constant expression type : " + tree.GetText())
	return nil
}

//=============================================================================

func convertType(tree parser.ITypeContext) datatype.Type {
	if tree.ListType() != nil {
		lt := tree.ListType()
		subType := convertType(lt.Type_())
		return datatype.NewListType(subType)
	}

	if tree.MapType() != nil {
		mt := tree.MapType()
		keyType := datatype.NewToFindOutType(mt.KeyType().GetText())
		valType := convertType(mt.Type_())
		return datatype.NewMapType(keyType, valType)
	}

	return datatype.NewToFindOutType(tree.GetText())
}

//=============================================================================

func convertFQIdentifier(tree parser.IFqIdentifierContext) *expression.FQIdentifier {
	fqi := expression.NewFQIdentifier()

	for _, s := range tree.AllIDENTIFIER() {
		name := s.GetText()
		fqi.AddScope(name)
	}

	return fqi
}

//=============================================================================

func raiseError(p antlr.Parser, message string) {
	p.NotifyErrorListeners(message, nil, nil)
}

//=============================================================================
