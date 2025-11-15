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

parser grammar TslParser;

options {
    tokenVocab = TslLexer;
}

//=============================================================================

script
    : packageDef? ( constantsDef | functionDef | enumDef | classDef )*
    ;

//=============================================================================
//===
//=== Top entities
//===
//=============================================================================

packageDef
    : PACKAGE IDENTIFIER
    ;

//=============================================================================

constantsDef
    : CONST L_PAREN ( constantDef )* R_PAREN
    ;

constantDef
    : IDENTIFIER EQUAL constantValue
    ;

constantValue
    : INT_VALUE
    | REAL_VALUE
    | timeValue
    | dateValue
    | STRING_VALUE
    | boolValue
    ;

timeValue
    : APOS INT_VALUE COLON INT_VALUE APOS;

dateValue
    : APOS INT_VALUE MINUS INT_VALUE MINUS INT_VALUE APOS;

boolValue
    : TRUE | FALSE ;

//=============================================================================

functionDef
    : FUNC (class)? IDENTIFIER parameters results? block
    ;

class
    : L_PAREN IDENTIFIER R_PAREN
    ;

parameters
    : L_PAREN (parameterDecl (COMMA parameterDecl)* )? COMMA? R_PAREN
    ;

parameterDecl
    : IDENTIFIER paramType
    ;

results
    : paramType (COMMA paramType)*
    ;

paramType
    : INT
    | REAL
    | BOOL
    | STRING
    | TIME
    | DATE
    | TIMESERIES
    | fqIdentifier
    | ERROR
    | MAP OF keyType COLON paramType
    | LIST OF paramType
    ;

keyType
    : INT
    | BOOL
    | STRING
    | TIME
    | DATE
    ;

//=============================================================================

enumDef
    : ENUM IDENTIFIER L_CURLY (enumItem)+ R_CURLY;

enumItem
    : IDENTIFIER ( L_PAREN enumValue R_PAREN )?
    ;

enumValue
    : INT_VALUE | STRING_VALUE;

//=============================================================================

classDef
    : CLASS IDENTIFIER L_CURLY (property)+ R_CURLY;

property
    : IDENTIFIER paramType
    ;

//=============================================================================
//===
//=== Statements
//===
//=============================================================================

block
    : L_CURLY (statement)* R_CURLY;

//=============================================================================

statement
    : varDeclaration
    | ifStatement
    | functionCall
    | returnStatement
    ;

//=============================================================================

varDeclaration
    : fqIdentifier EQUAL expression
    ;

//=============================================================================

ifStatement
    : IF expression block (elseIfBlock)* (elseBlock)?
    ;

elseIfBlock : ELSE IF expression block;

elseBlock : ELSE block;

//=============================================================================

functionCall
    : fqIdentifier L_PAREN (expression (COMMA expression)* )? R_PAREN
    ;

//=============================================================================

returnStatement : RETURN expression;

//=============================================================================
//===
//=== Expressions
//===
//=============================================================================

expression
    : identifierExpression
    | expressionInParenthesis
    | unaryExpression
    | functionCall
    | constantValue
    | expression (STAR | SLASH) expression
    | expression (PLUS | MINUS) expression
    | expression ( EQUAL | NOT_EQUAL | LESS_THAN | LESS_OR_EQUAL | GREATER_THAN | GREATER_OR_EQUAL ) expression
    | expression ( AND expression )+
    | expression ( OR  expression )+
    ;

identifierExpression
    : fqIdentifier ( IN_BAR expression )?
    ;

fqIdentifier
    : IDENTIFIER ( DOT IDENTIFIER )*
    ;

unaryExpression
    : (MINUS|PLUS|NOT) expression
    ;

expressionInParenthesis
    : L_PAREN expression R_PAREN;


//=============================================================================
