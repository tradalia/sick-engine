// Code generated from TslParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // TslParser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type TslParser struct {
	*antlr.BaseParser
}

var TslParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tslparserParserInit() {
	staticData := &TslParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'package'", "'params'", "'func'", "'if'", "'else'", "'for'", "'const'",
		"'true'", "'false'", "'return'", "'and'", "'or'", "'not'", "'null'",
		"'this'", "'of'", "'int'", "'real'", "'bool'", "'string'", "'time'",
		"'date'", "'timeseries'", "'enum'", "'class'", "'error'", "'list'",
		"'map'", "", "", "", "", "", "'('", "')'", "'['", "']'", "'{'", "'}'",
		"'='", "'*'", "'/'", "'+'", "'-'", "','", "'.'", "':'", "'''", "'<>'",
		"'<'", "'<='", "'>'", "'>='", "'::'",
	}
	staticData.SymbolicNames = []string{
		"", "PACKAGE", "PARAMS", "FUNC", "IF", "ELSE", "FOR", "CONST", "TRUE",
		"FALSE", "RETURN", "AND", "OR", "NOT", "NULL", "THIS", "OF", "INT",
		"REAL", "BOOL", "STRING", "TIME", "DATE", "TIMESERIES", "ENUM", "CLASS",
		"ERROR", "LIST", "MAP", "IDENTIFIER", "INT_VALUE", "REAL_VALUE", "DIGIT",
		"STRING_VALUE", "L_PAREN", "R_PAREN", "L_BRACKET", "R_BRACKET", "L_CURLY",
		"R_CURLY", "EQUAL", "STAR", "SLASH", "PLUS", "MINUS", "COMMA", "DOT",
		"COLON", "APOS", "NOT_EQUAL", "LESS_THAN", "LESS_OR_EQUAL", "GREATER_THAN",
		"GREATER_OR_EQUAL", "IN_BAR", "WS", "NEWLINE", "BLOCK_COMMENT", "LINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"script", "packageDef", "constantsDef", "constantDef", "constantValue",
		"timeValue", "dateValue", "boolValue", "functionDef", "class", "parameters",
		"parameterDecl", "results", "paramType", "keyType", "enumDef", "enumItem",
		"enumValue", "classDef", "property", "block", "statement", "varDeclaration",
		"ifStatement", "elseIfBlock", "elseBlock", "functionCall", "returnStatement",
		"expression", "identifierExpression", "fqIdentifier", "unaryExpression",
		"expressionInParenthesis",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 58, 328, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 1, 0, 3, 0, 68, 8, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 74,
		8, 0, 10, 0, 12, 0, 77, 9, 0, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 5, 2,
		85, 8, 2, 10, 2, 12, 2, 88, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 102, 8, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 8, 1, 8, 3, 8, 122, 8, 8, 1, 8, 1, 8, 1, 8, 3, 8, 127, 8, 8, 1, 8,
		1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 139, 8,
		10, 10, 10, 12, 10, 142, 9, 10, 3, 10, 144, 8, 10, 1, 10, 3, 10, 147, 8,
		10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 5, 12, 157,
		8, 12, 10, 12, 12, 12, 160, 9, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1,
		13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 3, 13, 180, 8, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 15, 1,
		15, 4, 15, 188, 8, 15, 11, 15, 12, 15, 189, 1, 15, 1, 15, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 3, 16, 199, 8, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		18, 1, 18, 4, 18, 207, 8, 18, 11, 18, 12, 18, 208, 1, 18, 1, 18, 1, 19,
		1, 19, 1, 19, 1, 20, 1, 20, 5, 20, 218, 8, 20, 10, 20, 12, 20, 221, 9,
		20, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 21, 3, 21, 229, 8, 21, 1, 22,
		1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 5, 23, 239, 8, 23, 10,
		23, 12, 23, 242, 9, 23, 1, 23, 3, 23, 245, 8, 23, 1, 24, 1, 24, 1, 24,
		1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 26, 5,
		26, 260, 8, 26, 10, 26, 12, 26, 263, 9, 26, 3, 26, 265, 8, 26, 1, 26, 1,
		26, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28,
		278, 8, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 4, 28, 292, 8, 28, 11, 28, 12, 28, 293, 1, 28,
		1, 28, 1, 28, 4, 28, 299, 8, 28, 11, 28, 12, 28, 300, 5, 28, 303, 8, 28,
		10, 28, 12, 28, 306, 9, 28, 1, 29, 1, 29, 1, 29, 3, 29, 311, 8, 29, 1,
		30, 1, 30, 1, 30, 5, 30, 316, 8, 30, 10, 30, 12, 30, 319, 9, 30, 1, 31,
		1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 1, 32, 1, 32, 0, 1, 56, 33, 0, 2, 4,
		6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42,
		44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 0, 7, 1, 0, 8, 9, 2, 0, 17,
		17, 19, 22, 2, 0, 30, 30, 33, 33, 1, 0, 41, 42, 1, 0, 43, 44, 2, 0, 40,
		40, 49, 53, 2, 0, 13, 13, 43, 44, 345, 0, 67, 1, 0, 0, 0, 2, 78, 1, 0,
		0, 0, 4, 81, 1, 0, 0, 0, 6, 91, 1, 0, 0, 0, 8, 101, 1, 0, 0, 0, 10, 103,
		1, 0, 0, 0, 12, 109, 1, 0, 0, 0, 14, 117, 1, 0, 0, 0, 16, 119, 1, 0, 0,
		0, 18, 130, 1, 0, 0, 0, 20, 134, 1, 0, 0, 0, 22, 150, 1, 0, 0, 0, 24, 153,
		1, 0, 0, 0, 26, 179, 1, 0, 0, 0, 28, 181, 1, 0, 0, 0, 30, 183, 1, 0, 0,
		0, 32, 193, 1, 0, 0, 0, 34, 200, 1, 0, 0, 0, 36, 202, 1, 0, 0, 0, 38, 212,
		1, 0, 0, 0, 40, 215, 1, 0, 0, 0, 42, 228, 1, 0, 0, 0, 44, 230, 1, 0, 0,
		0, 46, 234, 1, 0, 0, 0, 48, 246, 1, 0, 0, 0, 50, 251, 1, 0, 0, 0, 52, 254,
		1, 0, 0, 0, 54, 268, 1, 0, 0, 0, 56, 277, 1, 0, 0, 0, 58, 307, 1, 0, 0,
		0, 60, 312, 1, 0, 0, 0, 62, 320, 1, 0, 0, 0, 64, 323, 1, 0, 0, 0, 66, 68,
		3, 2, 1, 0, 67, 66, 1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 75, 1, 0, 0, 0,
		69, 74, 3, 4, 2, 0, 70, 74, 3, 16, 8, 0, 71, 74, 3, 30, 15, 0, 72, 74,
		3, 36, 18, 0, 73, 69, 1, 0, 0, 0, 73, 70, 1, 0, 0, 0, 73, 71, 1, 0, 0,
		0, 73, 72, 1, 0, 0, 0, 74, 77, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 75, 76,
		1, 0, 0, 0, 76, 1, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 78, 79, 5, 1, 0, 0,
		79, 80, 5, 29, 0, 0, 80, 3, 1, 0, 0, 0, 81, 82, 5, 7, 0, 0, 82, 86, 5,
		34, 0, 0, 83, 85, 3, 6, 3, 0, 84, 83, 1, 0, 0, 0, 85, 88, 1, 0, 0, 0, 86,
		84, 1, 0, 0, 0, 86, 87, 1, 0, 0, 0, 87, 89, 1, 0, 0, 0, 88, 86, 1, 0, 0,
		0, 89, 90, 5, 35, 0, 0, 90, 5, 1, 0, 0, 0, 91, 92, 5, 29, 0, 0, 92, 93,
		5, 40, 0, 0, 93, 94, 3, 8, 4, 0, 94, 7, 1, 0, 0, 0, 95, 102, 5, 30, 0,
		0, 96, 102, 5, 31, 0, 0, 97, 102, 3, 10, 5, 0, 98, 102, 3, 12, 6, 0, 99,
		102, 5, 33, 0, 0, 100, 102, 3, 14, 7, 0, 101, 95, 1, 0, 0, 0, 101, 96,
		1, 0, 0, 0, 101, 97, 1, 0, 0, 0, 101, 98, 1, 0, 0, 0, 101, 99, 1, 0, 0,
		0, 101, 100, 1, 0, 0, 0, 102, 9, 1, 0, 0, 0, 103, 104, 5, 48, 0, 0, 104,
		105, 5, 30, 0, 0, 105, 106, 5, 47, 0, 0, 106, 107, 5, 30, 0, 0, 107, 108,
		5, 48, 0, 0, 108, 11, 1, 0, 0, 0, 109, 110, 5, 48, 0, 0, 110, 111, 5, 30,
		0, 0, 111, 112, 5, 44, 0, 0, 112, 113, 5, 30, 0, 0, 113, 114, 5, 44, 0,
		0, 114, 115, 5, 30, 0, 0, 115, 116, 5, 48, 0, 0, 116, 13, 1, 0, 0, 0, 117,
		118, 7, 0, 0, 0, 118, 15, 1, 0, 0, 0, 119, 121, 5, 3, 0, 0, 120, 122, 3,
		18, 9, 0, 121, 120, 1, 0, 0, 0, 121, 122, 1, 0, 0, 0, 122, 123, 1, 0, 0,
		0, 123, 124, 5, 29, 0, 0, 124, 126, 3, 20, 10, 0, 125, 127, 3, 24, 12,
		0, 126, 125, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 128, 1, 0, 0, 0, 128,
		129, 3, 40, 20, 0, 129, 17, 1, 0, 0, 0, 130, 131, 5, 34, 0, 0, 131, 132,
		5, 29, 0, 0, 132, 133, 5, 35, 0, 0, 133, 19, 1, 0, 0, 0, 134, 143, 5, 34,
		0, 0, 135, 140, 3, 22, 11, 0, 136, 137, 5, 45, 0, 0, 137, 139, 3, 22, 11,
		0, 138, 136, 1, 0, 0, 0, 139, 142, 1, 0, 0, 0, 140, 138, 1, 0, 0, 0, 140,
		141, 1, 0, 0, 0, 141, 144, 1, 0, 0, 0, 142, 140, 1, 0, 0, 0, 143, 135,
		1, 0, 0, 0, 143, 144, 1, 0, 0, 0, 144, 146, 1, 0, 0, 0, 145, 147, 5, 45,
		0, 0, 146, 145, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0,
		148, 149, 5, 35, 0, 0, 149, 21, 1, 0, 0, 0, 150, 151, 5, 29, 0, 0, 151,
		152, 3, 26, 13, 0, 152, 23, 1, 0, 0, 0, 153, 158, 3, 26, 13, 0, 154, 155,
		5, 45, 0, 0, 155, 157, 3, 26, 13, 0, 156, 154, 1, 0, 0, 0, 157, 160, 1,
		0, 0, 0, 158, 156, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 25, 1, 0, 0,
		0, 160, 158, 1, 0, 0, 0, 161, 180, 5, 17, 0, 0, 162, 180, 5, 18, 0, 0,
		163, 180, 5, 19, 0, 0, 164, 180, 5, 20, 0, 0, 165, 180, 5, 21, 0, 0, 166,
		180, 5, 22, 0, 0, 167, 180, 5, 23, 0, 0, 168, 180, 3, 60, 30, 0, 169, 180,
		5, 26, 0, 0, 170, 171, 5, 28, 0, 0, 171, 172, 5, 16, 0, 0, 172, 173, 3,
		28, 14, 0, 173, 174, 5, 47, 0, 0, 174, 175, 3, 26, 13, 0, 175, 180, 1,
		0, 0, 0, 176, 177, 5, 27, 0, 0, 177, 178, 5, 16, 0, 0, 178, 180, 3, 26,
		13, 0, 179, 161, 1, 0, 0, 0, 179, 162, 1, 0, 0, 0, 179, 163, 1, 0, 0, 0,
		179, 164, 1, 0, 0, 0, 179, 165, 1, 0, 0, 0, 179, 166, 1, 0, 0, 0, 179,
		167, 1, 0, 0, 0, 179, 168, 1, 0, 0, 0, 179, 169, 1, 0, 0, 0, 179, 170,
		1, 0, 0, 0, 179, 176, 1, 0, 0, 0, 180, 27, 1, 0, 0, 0, 181, 182, 7, 1,
		0, 0, 182, 29, 1, 0, 0, 0, 183, 184, 5, 24, 0, 0, 184, 185, 5, 29, 0, 0,
		185, 187, 5, 38, 0, 0, 186, 188, 3, 32, 16, 0, 187, 186, 1, 0, 0, 0, 188,
		189, 1, 0, 0, 0, 189, 187, 1, 0, 0, 0, 189, 190, 1, 0, 0, 0, 190, 191,
		1, 0, 0, 0, 191, 192, 5, 39, 0, 0, 192, 31, 1, 0, 0, 0, 193, 198, 5, 29,
		0, 0, 194, 195, 5, 34, 0, 0, 195, 196, 3, 34, 17, 0, 196, 197, 5, 35, 0,
		0, 197, 199, 1, 0, 0, 0, 198, 194, 1, 0, 0, 0, 198, 199, 1, 0, 0, 0, 199,
		33, 1, 0, 0, 0, 200, 201, 7, 2, 0, 0, 201, 35, 1, 0, 0, 0, 202, 203, 5,
		25, 0, 0, 203, 204, 5, 29, 0, 0, 204, 206, 5, 38, 0, 0, 205, 207, 3, 38,
		19, 0, 206, 205, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 206, 1, 0, 0, 0,
		208, 209, 1, 0, 0, 0, 209, 210, 1, 0, 0, 0, 210, 211, 5, 39, 0, 0, 211,
		37, 1, 0, 0, 0, 212, 213, 5, 29, 0, 0, 213, 214, 3, 26, 13, 0, 214, 39,
		1, 0, 0, 0, 215, 219, 5, 38, 0, 0, 216, 218, 3, 42, 21, 0, 217, 216, 1,
		0, 0, 0, 218, 221, 1, 0, 0, 0, 219, 217, 1, 0, 0, 0, 219, 220, 1, 0, 0,
		0, 220, 222, 1, 0, 0, 0, 221, 219, 1, 0, 0, 0, 222, 223, 5, 39, 0, 0, 223,
		41, 1, 0, 0, 0, 224, 229, 3, 44, 22, 0, 225, 229, 3, 46, 23, 0, 226, 229,
		3, 52, 26, 0, 227, 229, 3, 54, 27, 0, 228, 224, 1, 0, 0, 0, 228, 225, 1,
		0, 0, 0, 228, 226, 1, 0, 0, 0, 228, 227, 1, 0, 0, 0, 229, 43, 1, 0, 0,
		0, 230, 231, 3, 60, 30, 0, 231, 232, 5, 40, 0, 0, 232, 233, 3, 56, 28,
		0, 233, 45, 1, 0, 0, 0, 234, 235, 5, 4, 0, 0, 235, 236, 3, 56, 28, 0, 236,
		240, 3, 40, 20, 0, 237, 239, 3, 48, 24, 0, 238, 237, 1, 0, 0, 0, 239, 242,
		1, 0, 0, 0, 240, 238, 1, 0, 0, 0, 240, 241, 1, 0, 0, 0, 241, 244, 1, 0,
		0, 0, 242, 240, 1, 0, 0, 0, 243, 245, 3, 50, 25, 0, 244, 243, 1, 0, 0,
		0, 244, 245, 1, 0, 0, 0, 245, 47, 1, 0, 0, 0, 246, 247, 5, 5, 0, 0, 247,
		248, 5, 4, 0, 0, 248, 249, 3, 56, 28, 0, 249, 250, 3, 40, 20, 0, 250, 49,
		1, 0, 0, 0, 251, 252, 5, 5, 0, 0, 252, 253, 3, 40, 20, 0, 253, 51, 1, 0,
		0, 0, 254, 255, 3, 60, 30, 0, 255, 264, 5, 34, 0, 0, 256, 261, 3, 56, 28,
		0, 257, 258, 5, 45, 0, 0, 258, 260, 3, 56, 28, 0, 259, 257, 1, 0, 0, 0,
		260, 263, 1, 0, 0, 0, 261, 259, 1, 0, 0, 0, 261, 262, 1, 0, 0, 0, 262,
		265, 1, 0, 0, 0, 263, 261, 1, 0, 0, 0, 264, 256, 1, 0, 0, 0, 264, 265,
		1, 0, 0, 0, 265, 266, 1, 0, 0, 0, 266, 267, 5, 35, 0, 0, 267, 53, 1, 0,
		0, 0, 268, 269, 5, 10, 0, 0, 269, 270, 3, 56, 28, 0, 270, 55, 1, 0, 0,
		0, 271, 272, 6, 28, -1, 0, 272, 278, 3, 58, 29, 0, 273, 278, 3, 64, 32,
		0, 274, 278, 3, 62, 31, 0, 275, 278, 3, 52, 26, 0, 276, 278, 3, 8, 4, 0,
		277, 271, 1, 0, 0, 0, 277, 273, 1, 0, 0, 0, 277, 274, 1, 0, 0, 0, 277,
		275, 1, 0, 0, 0, 277, 276, 1, 0, 0, 0, 278, 304, 1, 0, 0, 0, 279, 280,
		10, 5, 0, 0, 280, 281, 7, 3, 0, 0, 281, 303, 3, 56, 28, 6, 282, 283, 10,
		4, 0, 0, 283, 284, 7, 4, 0, 0, 284, 303, 3, 56, 28, 5, 285, 286, 10, 3,
		0, 0, 286, 287, 7, 5, 0, 0, 287, 303, 3, 56, 28, 4, 288, 291, 10, 2, 0,
		0, 289, 290, 5, 11, 0, 0, 290, 292, 3, 56, 28, 0, 291, 289, 1, 0, 0, 0,
		292, 293, 1, 0, 0, 0, 293, 291, 1, 0, 0, 0, 293, 294, 1, 0, 0, 0, 294,
		303, 1, 0, 0, 0, 295, 298, 10, 1, 0, 0, 296, 297, 5, 12, 0, 0, 297, 299,
		3, 56, 28, 0, 298, 296, 1, 0, 0, 0, 299, 300, 1, 0, 0, 0, 300, 298, 1,
		0, 0, 0, 300, 301, 1, 0, 0, 0, 301, 303, 1, 0, 0, 0, 302, 279, 1, 0, 0,
		0, 302, 282, 1, 0, 0, 0, 302, 285, 1, 0, 0, 0, 302, 288, 1, 0, 0, 0, 302,
		295, 1, 0, 0, 0, 303, 306, 1, 0, 0, 0, 304, 302, 1, 0, 0, 0, 304, 305,
		1, 0, 0, 0, 305, 57, 1, 0, 0, 0, 306, 304, 1, 0, 0, 0, 307, 310, 3, 60,
		30, 0, 308, 309, 5, 54, 0, 0, 309, 311, 3, 56, 28, 0, 310, 308, 1, 0, 0,
		0, 310, 311, 1, 0, 0, 0, 311, 59, 1, 0, 0, 0, 312, 317, 5, 29, 0, 0, 313,
		314, 5, 46, 0, 0, 314, 316, 5, 29, 0, 0, 315, 313, 1, 0, 0, 0, 316, 319,
		1, 0, 0, 0, 317, 315, 1, 0, 0, 0, 317, 318, 1, 0, 0, 0, 318, 61, 1, 0,
		0, 0, 319, 317, 1, 0, 0, 0, 320, 321, 7, 6, 0, 0, 321, 322, 3, 56, 28,
		0, 322, 63, 1, 0, 0, 0, 323, 324, 5, 34, 0, 0, 324, 325, 3, 56, 28, 0,
		325, 326, 5, 35, 0, 0, 326, 65, 1, 0, 0, 0, 28, 67, 73, 75, 86, 101, 121,
		126, 140, 143, 146, 158, 179, 189, 198, 208, 219, 228, 240, 244, 261, 264,
		277, 293, 300, 302, 304, 310, 317,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// TslParserInit initializes any static state used to implement TslParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTslParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TslParserInit() {
	staticData := &TslParserParserStaticData
	staticData.once.Do(tslparserParserInit)
}

// NewTslParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTslParser(input antlr.TokenStream) *TslParser {
	TslParserInit()
	this := new(TslParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &TslParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "TslParser.g4"

	return this
}

// TslParser tokens.
const (
	TslParserEOF              = antlr.TokenEOF
	TslParserPACKAGE          = 1
	TslParserPARAMS           = 2
	TslParserFUNC             = 3
	TslParserIF               = 4
	TslParserELSE             = 5
	TslParserFOR              = 6
	TslParserCONST            = 7
	TslParserTRUE             = 8
	TslParserFALSE            = 9
	TslParserRETURN           = 10
	TslParserAND              = 11
	TslParserOR               = 12
	TslParserNOT              = 13
	TslParserNULL             = 14
	TslParserTHIS             = 15
	TslParserOF               = 16
	TslParserINT              = 17
	TslParserREAL             = 18
	TslParserBOOL             = 19
	TslParserSTRING           = 20
	TslParserTIME             = 21
	TslParserDATE             = 22
	TslParserTIMESERIES       = 23
	TslParserENUM             = 24
	TslParserCLASS            = 25
	TslParserERROR            = 26
	TslParserLIST             = 27
	TslParserMAP              = 28
	TslParserIDENTIFIER       = 29
	TslParserINT_VALUE        = 30
	TslParserREAL_VALUE       = 31
	TslParserDIGIT            = 32
	TslParserSTRING_VALUE     = 33
	TslParserL_PAREN          = 34
	TslParserR_PAREN          = 35
	TslParserL_BRACKET        = 36
	TslParserR_BRACKET        = 37
	TslParserL_CURLY          = 38
	TslParserR_CURLY          = 39
	TslParserEQUAL            = 40
	TslParserSTAR             = 41
	TslParserSLASH            = 42
	TslParserPLUS             = 43
	TslParserMINUS            = 44
	TslParserCOMMA            = 45
	TslParserDOT              = 46
	TslParserCOLON            = 47
	TslParserAPOS             = 48
	TslParserNOT_EQUAL        = 49
	TslParserLESS_THAN        = 50
	TslParserLESS_OR_EQUAL    = 51
	TslParserGREATER_THAN     = 52
	TslParserGREATER_OR_EQUAL = 53
	TslParserIN_BAR           = 54
	TslParserWS               = 55
	TslParserNEWLINE          = 56
	TslParserBLOCK_COMMENT    = 57
	TslParserLINE_COMMENT     = 58
)

// TslParser rules.
const (
	TslParserRULE_script                  = 0
	TslParserRULE_packageDef              = 1
	TslParserRULE_constantsDef            = 2
	TslParserRULE_constantDef             = 3
	TslParserRULE_constantValue           = 4
	TslParserRULE_timeValue               = 5
	TslParserRULE_dateValue               = 6
	TslParserRULE_boolValue               = 7
	TslParserRULE_functionDef             = 8
	TslParserRULE_class                   = 9
	TslParserRULE_parameters              = 10
	TslParserRULE_parameterDecl           = 11
	TslParserRULE_results                 = 12
	TslParserRULE_paramType               = 13
	TslParserRULE_keyType                 = 14
	TslParserRULE_enumDef                 = 15
	TslParserRULE_enumItem                = 16
	TslParserRULE_enumValue               = 17
	TslParserRULE_classDef                = 18
	TslParserRULE_property                = 19
	TslParserRULE_block                   = 20
	TslParserRULE_statement               = 21
	TslParserRULE_varDeclaration          = 22
	TslParserRULE_ifStatement             = 23
	TslParserRULE_elseIfBlock             = 24
	TslParserRULE_elseBlock               = 25
	TslParserRULE_functionCall            = 26
	TslParserRULE_returnStatement         = 27
	TslParserRULE_expression              = 28
	TslParserRULE_identifierExpression    = 29
	TslParserRULE_fqIdentifier            = 30
	TslParserRULE_unaryExpression         = 31
	TslParserRULE_expressionInParenthesis = 32
)

// IScriptContext is an interface to support dynamic dispatch.
type IScriptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PackageDef() IPackageDefContext
	AllConstantsDef() []IConstantsDefContext
	ConstantsDef(i int) IConstantsDefContext
	AllFunctionDef() []IFunctionDefContext
	FunctionDef(i int) IFunctionDefContext
	AllEnumDef() []IEnumDefContext
	EnumDef(i int) IEnumDefContext
	AllClassDef() []IClassDefContext
	ClassDef(i int) IClassDefContext

	// IsScriptContext differentiates from other interfaces.
	IsScriptContext()
}

type ScriptContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScriptContext() *ScriptContext {
	var p = new(ScriptContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_script
	return p
}

func InitEmptyScriptContext(p *ScriptContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_script
}

func (*ScriptContext) IsScriptContext() {}

func NewScriptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScriptContext {
	var p = new(ScriptContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_script

	return p
}

func (s *ScriptContext) GetParser() antlr.Parser { return s.parser }

func (s *ScriptContext) PackageDef() IPackageDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPackageDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPackageDefContext)
}

func (s *ScriptContext) AllConstantsDef() []IConstantsDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstantsDefContext); ok {
			len++
		}
	}

	tst := make([]IConstantsDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstantsDefContext); ok {
			tst[i] = t.(IConstantsDefContext)
			i++
		}
	}

	return tst
}

func (s *ScriptContext) ConstantsDef(i int) IConstantsDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantsDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantsDefContext)
}

func (s *ScriptContext) AllFunctionDef() []IFunctionDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFunctionDefContext); ok {
			len++
		}
	}

	tst := make([]IFunctionDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFunctionDefContext); ok {
			tst[i] = t.(IFunctionDefContext)
			i++
		}
	}

	return tst
}

func (s *ScriptContext) FunctionDef(i int) IFunctionDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDefContext)
}

func (s *ScriptContext) AllEnumDef() []IEnumDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnumDefContext); ok {
			len++
		}
	}

	tst := make([]IEnumDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnumDefContext); ok {
			tst[i] = t.(IEnumDefContext)
			i++
		}
	}

	return tst
}

func (s *ScriptContext) EnumDef(i int) IEnumDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumDefContext)
}

func (s *ScriptContext) AllClassDef() []IClassDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IClassDefContext); ok {
			len++
		}
	}

	tst := make([]IClassDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IClassDefContext); ok {
			tst[i] = t.(IClassDefContext)
			i++
		}
	}

	return tst
}

func (s *ScriptContext) ClassDef(i int) IClassDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassDefContext)
}

func (s *ScriptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScriptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScriptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterScript(s)
	}
}

func (s *ScriptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitScript(s)
	}
}

func (p *TslParser) Script() (localctx IScriptContext) {
	localctx = NewScriptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TslParserRULE_script)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(67)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TslParserPACKAGE {
		{
			p.SetState(66)
			p.PackageDef()
		}

	}
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&50331784) != 0 {
		p.SetState(73)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case TslParserCONST:
			{
				p.SetState(69)
				p.ConstantsDef()
			}

		case TslParserFUNC:
			{
				p.SetState(70)
				p.FunctionDef()
			}

		case TslParserENUM:
			{
				p.SetState(71)
				p.EnumDef()
			}

		case TslParserCLASS:
			{
				p.SetState(72)
				p.ClassDef()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(77)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPackageDefContext is an interface to support dynamic dispatch.
type IPackageDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PACKAGE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode

	// IsPackageDefContext differentiates from other interfaces.
	IsPackageDefContext()
}

type PackageDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageDefContext() *PackageDefContext {
	var p = new(PackageDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_packageDef
	return p
}

func InitEmptyPackageDefContext(p *PackageDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_packageDef
}

func (*PackageDefContext) IsPackageDefContext() {}

func NewPackageDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageDefContext {
	var p = new(PackageDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_packageDef

	return p
}

func (s *PackageDefContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageDefContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(TslParserPACKAGE, 0)
}

func (s *PackageDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TslParserIDENTIFIER, 0)
}

func (s *PackageDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterPackageDef(s)
	}
}

func (s *PackageDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitPackageDef(s)
	}
}

func (p *TslParser) PackageDef() (localctx IPackageDefContext) {
	localctx = NewPackageDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TslParserRULE_packageDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(78)
		p.Match(TslParserPACKAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(79)
		p.Match(TslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstantsDefContext is an interface to support dynamic dispatch.
type IConstantsDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONST() antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	AllConstantDef() []IConstantDefContext
	ConstantDef(i int) IConstantDefContext

	// IsConstantsDefContext differentiates from other interfaces.
	IsConstantsDefContext()
}

type ConstantsDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantsDefContext() *ConstantsDefContext {
	var p = new(ConstantsDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_constantsDef
	return p
}

func InitEmptyConstantsDefContext(p *ConstantsDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_constantsDef
}

func (*ConstantsDefContext) IsConstantsDefContext() {}

func NewConstantsDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantsDefContext {
	var p = new(ConstantsDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_constantsDef

	return p
}

func (s *ConstantsDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantsDefContext) CONST() antlr.TerminalNode {
	return s.GetToken(TslParserCONST, 0)
}

func (s *ConstantsDefContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserL_PAREN, 0)
}

func (s *ConstantsDefContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserR_PAREN, 0)
}

func (s *ConstantsDefContext) AllConstantDef() []IConstantDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstantDefContext); ok {
			len++
		}
	}

	tst := make([]IConstantDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstantDefContext); ok {
			tst[i] = t.(IConstantDefContext)
			i++
		}
	}

	return tst
}

func (s *ConstantsDefContext) ConstantDef(i int) IConstantDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantDefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantDefContext)
}

func (s *ConstantsDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantsDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantsDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterConstantsDef(s)
	}
}

func (s *ConstantsDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitConstantsDef(s)
	}
}

func (p *TslParser) ConstantsDef() (localctx IConstantsDefContext) {
	localctx = NewConstantsDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TslParserRULE_constantsDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(81)
		p.Match(TslParserCONST)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(82)
		p.Match(TslParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TslParserIDENTIFIER {
		{
			p.SetState(83)
			p.ConstantDef()
		}

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(89)
		p.Match(TslParserR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstantDefContext is an interface to support dynamic dispatch.
type IConstantDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	ConstantValue() IConstantValueContext

	// IsConstantDefContext differentiates from other interfaces.
	IsConstantDefContext()
}

type ConstantDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantDefContext() *ConstantDefContext {
	var p = new(ConstantDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_constantDef
	return p
}

func InitEmptyConstantDefContext(p *ConstantDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_constantDef
}

func (*ConstantDefContext) IsConstantDefContext() {}

func NewConstantDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantDefContext {
	var p = new(ConstantDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_constantDef

	return p
}

func (s *ConstantDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TslParserIDENTIFIER, 0)
}

func (s *ConstantDefContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(TslParserEQUAL, 0)
}

func (s *ConstantDefContext) ConstantValue() IConstantValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantValueContext)
}

func (s *ConstantDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterConstantDef(s)
	}
}

func (s *ConstantDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitConstantDef(s)
	}
}

func (p *TslParser) ConstantDef() (localctx IConstantDefContext) {
	localctx = NewConstantDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TslParserRULE_constantDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(91)
		p.Match(TslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(92)
		p.Match(TslParserEQUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.ConstantValue()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstantValueContext is an interface to support dynamic dispatch.
type IConstantValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT_VALUE() antlr.TerminalNode
	REAL_VALUE() antlr.TerminalNode
	TimeValue() ITimeValueContext
	DateValue() IDateValueContext
	STRING_VALUE() antlr.TerminalNode
	BoolValue() IBoolValueContext

	// IsConstantValueContext differentiates from other interfaces.
	IsConstantValueContext()
}

type ConstantValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantValueContext() *ConstantValueContext {
	var p = new(ConstantValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_constantValue
	return p
}

func InitEmptyConstantValueContext(p *ConstantValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_constantValue
}

func (*ConstantValueContext) IsConstantValueContext() {}

func NewConstantValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantValueContext {
	var p = new(ConstantValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_constantValue

	return p
}

func (s *ConstantValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantValueContext) INT_VALUE() antlr.TerminalNode {
	return s.GetToken(TslParserINT_VALUE, 0)
}

func (s *ConstantValueContext) REAL_VALUE() antlr.TerminalNode {
	return s.GetToken(TslParserREAL_VALUE, 0)
}

func (s *ConstantValueContext) TimeValue() ITimeValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITimeValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITimeValueContext)
}

func (s *ConstantValueContext) DateValue() IDateValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDateValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDateValueContext)
}

func (s *ConstantValueContext) STRING_VALUE() antlr.TerminalNode {
	return s.GetToken(TslParserSTRING_VALUE, 0)
}

func (s *ConstantValueContext) BoolValue() IBoolValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolValueContext)
}

func (s *ConstantValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterConstantValue(s)
	}
}

func (s *ConstantValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitConstantValue(s)
	}
}

func (p *TslParser) ConstantValue() (localctx IConstantValueContext) {
	localctx = NewConstantValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TslParserRULE_constantValue)
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(95)
			p.Match(TslParserINT_VALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(96)
			p.Match(TslParserREAL_VALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(97)
			p.TimeValue()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(98)
			p.DateValue()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(99)
			p.Match(TslParserSTRING_VALUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(100)
			p.BoolValue()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITimeValueContext is an interface to support dynamic dispatch.
type ITimeValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAPOS() []antlr.TerminalNode
	APOS(i int) antlr.TerminalNode
	AllINT_VALUE() []antlr.TerminalNode
	INT_VALUE(i int) antlr.TerminalNode
	COLON() antlr.TerminalNode

	// IsTimeValueContext differentiates from other interfaces.
	IsTimeValueContext()
}

type TimeValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTimeValueContext() *TimeValueContext {
	var p = new(TimeValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_timeValue
	return p
}

func InitEmptyTimeValueContext(p *TimeValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_timeValue
}

func (*TimeValueContext) IsTimeValueContext() {}

func NewTimeValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TimeValueContext {
	var p = new(TimeValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_timeValue

	return p
}

func (s *TimeValueContext) GetParser() antlr.Parser { return s.parser }

func (s *TimeValueContext) AllAPOS() []antlr.TerminalNode {
	return s.GetTokens(TslParserAPOS)
}

func (s *TimeValueContext) APOS(i int) antlr.TerminalNode {
	return s.GetToken(TslParserAPOS, i)
}

func (s *TimeValueContext) AllINT_VALUE() []antlr.TerminalNode {
	return s.GetTokens(TslParserINT_VALUE)
}

func (s *TimeValueContext) INT_VALUE(i int) antlr.TerminalNode {
	return s.GetToken(TslParserINT_VALUE, i)
}

func (s *TimeValueContext) COLON() antlr.TerminalNode {
	return s.GetToken(TslParserCOLON, 0)
}

func (s *TimeValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TimeValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TimeValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterTimeValue(s)
	}
}

func (s *TimeValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitTimeValue(s)
	}
}

func (p *TslParser) TimeValue() (localctx ITimeValueContext) {
	localctx = NewTimeValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, TslParserRULE_timeValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Match(TslParserAPOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)
		p.Match(TslParserINT_VALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(105)
		p.Match(TslParserCOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(106)
		p.Match(TslParserINT_VALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(107)
		p.Match(TslParserAPOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDateValueContext is an interface to support dynamic dispatch.
type IDateValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllAPOS() []antlr.TerminalNode
	APOS(i int) antlr.TerminalNode
	AllINT_VALUE() []antlr.TerminalNode
	INT_VALUE(i int) antlr.TerminalNode
	AllMINUS() []antlr.TerminalNode
	MINUS(i int) antlr.TerminalNode

	// IsDateValueContext differentiates from other interfaces.
	IsDateValueContext()
}

type DateValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDateValueContext() *DateValueContext {
	var p = new(DateValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_dateValue
	return p
}

func InitEmptyDateValueContext(p *DateValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_dateValue
}

func (*DateValueContext) IsDateValueContext() {}

func NewDateValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DateValueContext {
	var p = new(DateValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_dateValue

	return p
}

func (s *DateValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DateValueContext) AllAPOS() []antlr.TerminalNode {
	return s.GetTokens(TslParserAPOS)
}

func (s *DateValueContext) APOS(i int) antlr.TerminalNode {
	return s.GetToken(TslParserAPOS, i)
}

func (s *DateValueContext) AllINT_VALUE() []antlr.TerminalNode {
	return s.GetTokens(TslParserINT_VALUE)
}

func (s *DateValueContext) INT_VALUE(i int) antlr.TerminalNode {
	return s.GetToken(TslParserINT_VALUE, i)
}

func (s *DateValueContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(TslParserMINUS)
}

func (s *DateValueContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(TslParserMINUS, i)
}

func (s *DateValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DateValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterDateValue(s)
	}
}

func (s *DateValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitDateValue(s)
	}
}

func (p *TslParser) DateValue() (localctx IDateValueContext) {
	localctx = NewDateValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, TslParserRULE_dateValue)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(TslParserAPOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(110)
		p.Match(TslParserINT_VALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(111)
		p.Match(TslParserMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(112)
		p.Match(TslParserINT_VALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(113)
		p.Match(TslParserMINUS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(114)
		p.Match(TslParserINT_VALUE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(115)
		p.Match(TslParserAPOS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBoolValueContext is an interface to support dynamic dispatch.
type IBoolValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode

	// IsBoolValueContext differentiates from other interfaces.
	IsBoolValueContext()
}

type BoolValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolValueContext() *BoolValueContext {
	var p = new(BoolValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_boolValue
	return p
}

func InitEmptyBoolValueContext(p *BoolValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_boolValue
}

func (*BoolValueContext) IsBoolValueContext() {}

func NewBoolValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolValueContext {
	var p = new(BoolValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_boolValue

	return p
}

func (s *BoolValueContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolValueContext) TRUE() antlr.TerminalNode {
	return s.GetToken(TslParserTRUE, 0)
}

func (s *BoolValueContext) FALSE() antlr.TerminalNode {
	return s.GetToken(TslParserFALSE, 0)
}

func (s *BoolValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterBoolValue(s)
	}
}

func (s *BoolValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitBoolValue(s)
	}
}

func (p *TslParser) BoolValue() (localctx IBoolValueContext) {
	localctx = NewBoolValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, TslParserRULE_boolValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(117)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TslParserTRUE || _la == TslParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionDefContext is an interface to support dynamic dispatch.
type IFunctionDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FUNC() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	Parameters() IParametersContext
	Block() IBlockContext
	Class() IClassContext
	Results() IResultsContext

	// IsFunctionDefContext differentiates from other interfaces.
	IsFunctionDefContext()
}

type FunctionDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDefContext() *FunctionDefContext {
	var p = new(FunctionDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_functionDef
	return p
}

func InitEmptyFunctionDefContext(p *FunctionDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_functionDef
}

func (*FunctionDefContext) IsFunctionDefContext() {}

func NewFunctionDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDefContext {
	var p = new(FunctionDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_functionDef

	return p
}

func (s *FunctionDefContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDefContext) FUNC() antlr.TerminalNode {
	return s.GetToken(TslParserFUNC, 0)
}

func (s *FunctionDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TslParserIDENTIFIER, 0)
}

func (s *FunctionDefContext) Parameters() IParametersContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParametersContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParametersContext)
}

func (s *FunctionDefContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDefContext) Class() IClassContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassContext)
}

func (s *FunctionDefContext) Results() IResultsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResultsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResultsContext)
}

func (s *FunctionDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterFunctionDef(s)
	}
}

func (s *FunctionDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitFunctionDef(s)
	}
}

func (p *TslParser) FunctionDef() (localctx IFunctionDefContext) {
	localctx = NewFunctionDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, TslParserRULE_functionDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(TslParserFUNC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(121)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TslParserL_PAREN {
		{
			p.SetState(120)
			p.Class()
		}

	}
	{
		p.SetState(123)
		p.Match(TslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(124)
		p.Parameters()
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1023279104) != 0 {
		{
			p.SetState(125)
			p.Results()
		}

	}
	{
		p.SetState(128)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IClassContext is an interface to support dynamic dispatch.
type IClassContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_PAREN() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode

	// IsClassContext differentiates from other interfaces.
	IsClassContext()
}

type ClassContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassContext() *ClassContext {
	var p = new(ClassContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_class
	return p
}

func InitEmptyClassContext(p *ClassContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_class
}

func (*ClassContext) IsClassContext() {}

func NewClassContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassContext {
	var p = new(ClassContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_class

	return p
}

func (s *ClassContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserL_PAREN, 0)
}

func (s *ClassContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TslParserIDENTIFIER, 0)
}

func (s *ClassContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserR_PAREN, 0)
}

func (s *ClassContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterClass(s)
	}
}

func (s *ClassContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitClass(s)
	}
}

func (p *TslParser) Class() (localctx IClassContext) {
	localctx = NewClassContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, TslParserRULE_class)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.Match(TslParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(131)
		p.Match(TslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.Match(TslParserR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParametersContext is an interface to support dynamic dispatch.
type IParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	AllParameterDecl() []IParameterDeclContext
	ParameterDecl(i int) IParameterDeclContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsParametersContext differentiates from other interfaces.
	IsParametersContext()
}

type ParametersContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParametersContext() *ParametersContext {
	var p = new(ParametersContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_parameters
	return p
}

func InitEmptyParametersContext(p *ParametersContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_parameters
}

func (*ParametersContext) IsParametersContext() {}

func NewParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParametersContext {
	var p = new(ParametersContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_parameters

	return p
}

func (s *ParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *ParametersContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserL_PAREN, 0)
}

func (s *ParametersContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserR_PAREN, 0)
}

func (s *ParametersContext) AllParameterDecl() []IParameterDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterDeclContext); ok {
			len++
		}
	}

	tst := make([]IParameterDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterDeclContext); ok {
			tst[i] = t.(IParameterDeclContext)
			i++
		}
	}

	return tst
}

func (s *ParametersContext) ParameterDecl(i int) IParameterDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterDeclContext)
}

func (s *ParametersContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TslParserCOMMA)
}

func (s *ParametersContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TslParserCOMMA, i)
}

func (s *ParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterParameters(s)
	}
}

func (s *ParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitParameters(s)
	}
}

func (p *TslParser) Parameters() (localctx IParametersContext) {
	localctx = NewParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, TslParserRULE_parameters)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Match(TslParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TslParserIDENTIFIER {
		{
			p.SetState(135)
			p.ParameterDecl()
		}
		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(136)
					p.Match(TslParserCOMMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(137)
					p.ParameterDecl()
				}

			}
			p.SetState(142)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TslParserCOMMA {
		{
			p.SetState(145)
			p.Match(TslParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(148)
		p.Match(TslParserR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterDeclContext is an interface to support dynamic dispatch.
type IParameterDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	ParamType() IParamTypeContext

	// IsParameterDeclContext differentiates from other interfaces.
	IsParameterDeclContext()
}

type ParameterDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterDeclContext() *ParameterDeclContext {
	var p = new(ParameterDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_parameterDecl
	return p
}

func InitEmptyParameterDeclContext(p *ParameterDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_parameterDecl
}

func (*ParameterDeclContext) IsParameterDeclContext() {}

func NewParameterDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterDeclContext {
	var p = new(ParameterDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_parameterDecl

	return p
}

func (s *ParameterDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterDeclContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TslParserIDENTIFIER, 0)
}

func (s *ParameterDeclContext) ParamType() IParamTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamTypeContext)
}

func (s *ParameterDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterParameterDecl(s)
	}
}

func (s *ParameterDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitParameterDecl(s)
	}
}

func (p *TslParser) ParameterDecl() (localctx IParameterDeclContext) {
	localctx = NewParameterDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, TslParserRULE_parameterDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Match(TslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
		p.ParamType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IResultsContext is an interface to support dynamic dispatch.
type IResultsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParamType() []IParamTypeContext
	ParamType(i int) IParamTypeContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsResultsContext differentiates from other interfaces.
	IsResultsContext()
}

type ResultsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResultsContext() *ResultsContext {
	var p = new(ResultsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_results
	return p
}

func InitEmptyResultsContext(p *ResultsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_results
}

func (*ResultsContext) IsResultsContext() {}

func NewResultsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResultsContext {
	var p = new(ResultsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_results

	return p
}

func (s *ResultsContext) GetParser() antlr.Parser { return s.parser }

func (s *ResultsContext) AllParamType() []IParamTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParamTypeContext); ok {
			len++
		}
	}

	tst := make([]IParamTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParamTypeContext); ok {
			tst[i] = t.(IParamTypeContext)
			i++
		}
	}

	return tst
}

func (s *ResultsContext) ParamType(i int) IParamTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamTypeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamTypeContext)
}

func (s *ResultsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TslParserCOMMA)
}

func (s *ResultsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TslParserCOMMA, i)
}

func (s *ResultsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResultsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResultsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterResults(s)
	}
}

func (s *ResultsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitResults(s)
	}
}

func (p *TslParser) Results() (localctx IResultsContext) {
	localctx = NewResultsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, TslParserRULE_results)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(153)
		p.ParamType()
	}
	p.SetState(158)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == TslParserCOMMA {
		{
			p.SetState(154)
			p.Match(TslParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(155)
			p.ParamType()
		}

		p.SetState(160)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParamTypeContext is an interface to support dynamic dispatch.
type IParamTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	REAL() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	STRING() antlr.TerminalNode
	TIME() antlr.TerminalNode
	DATE() antlr.TerminalNode
	TIMESERIES() antlr.TerminalNode
	FqIdentifier() IFqIdentifierContext
	ERROR() antlr.TerminalNode
	MAP() antlr.TerminalNode
	OF() antlr.TerminalNode
	KeyType() IKeyTypeContext
	COLON() antlr.TerminalNode
	ParamType() IParamTypeContext
	LIST() antlr.TerminalNode

	// IsParamTypeContext differentiates from other interfaces.
	IsParamTypeContext()
}

type ParamTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParamTypeContext() *ParamTypeContext {
	var p = new(ParamTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_paramType
	return p
}

func InitEmptyParamTypeContext(p *ParamTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_paramType
}

func (*ParamTypeContext) IsParamTypeContext() {}

func NewParamTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParamTypeContext {
	var p = new(ParamTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_paramType

	return p
}

func (s *ParamTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ParamTypeContext) INT() antlr.TerminalNode {
	return s.GetToken(TslParserINT, 0)
}

func (s *ParamTypeContext) REAL() antlr.TerminalNode {
	return s.GetToken(TslParserREAL, 0)
}

func (s *ParamTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(TslParserBOOL, 0)
}

func (s *ParamTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(TslParserSTRING, 0)
}

func (s *ParamTypeContext) TIME() antlr.TerminalNode {
	return s.GetToken(TslParserTIME, 0)
}

func (s *ParamTypeContext) DATE() antlr.TerminalNode {
	return s.GetToken(TslParserDATE, 0)
}

func (s *ParamTypeContext) TIMESERIES() antlr.TerminalNode {
	return s.GetToken(TslParserTIMESERIES, 0)
}

func (s *ParamTypeContext) FqIdentifier() IFqIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFqIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFqIdentifierContext)
}

func (s *ParamTypeContext) ERROR() antlr.TerminalNode {
	return s.GetToken(TslParserERROR, 0)
}

func (s *ParamTypeContext) MAP() antlr.TerminalNode {
	return s.GetToken(TslParserMAP, 0)
}

func (s *ParamTypeContext) OF() antlr.TerminalNode {
	return s.GetToken(TslParserOF, 0)
}

func (s *ParamTypeContext) KeyType() IKeyTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeyTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeyTypeContext)
}

func (s *ParamTypeContext) COLON() antlr.TerminalNode {
	return s.GetToken(TslParserCOLON, 0)
}

func (s *ParamTypeContext) ParamType() IParamTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamTypeContext)
}

func (s *ParamTypeContext) LIST() antlr.TerminalNode {
	return s.GetToken(TslParserLIST, 0)
}

func (s *ParamTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParamTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParamTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterParamType(s)
	}
}

func (s *ParamTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitParamType(s)
	}
}

func (p *TslParser) ParamType() (localctx IParamTypeContext) {
	localctx = NewParamTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, TslParserRULE_paramType)
	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case TslParserINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(161)
			p.Match(TslParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TslParserREAL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(162)
			p.Match(TslParserREAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TslParserBOOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(163)
			p.Match(TslParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TslParserSTRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(164)
			p.Match(TslParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TslParserTIME:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(165)
			p.Match(TslParserTIME)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TslParserDATE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(166)
			p.Match(TslParserDATE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TslParserTIMESERIES:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(167)
			p.Match(TslParserTIMESERIES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TslParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(168)
			p.FqIdentifier()
		}

	case TslParserERROR:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(169)
			p.Match(TslParserERROR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case TslParserMAP:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(170)
			p.Match(TslParserMAP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(171)
			p.Match(TslParserOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(172)
			p.KeyType()
		}
		{
			p.SetState(173)
			p.Match(TslParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(174)
			p.ParamType()
		}

	case TslParserLIST:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(176)
			p.Match(TslParserLIST)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(177)
			p.Match(TslParserOF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(178)
			p.ParamType()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IKeyTypeContext is an interface to support dynamic dispatch.
type IKeyTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	STRING() antlr.TerminalNode
	TIME() antlr.TerminalNode
	DATE() antlr.TerminalNode

	// IsKeyTypeContext differentiates from other interfaces.
	IsKeyTypeContext()
}

type KeyTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeyTypeContext() *KeyTypeContext {
	var p = new(KeyTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_keyType
	return p
}

func InitEmptyKeyTypeContext(p *KeyTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_keyType
}

func (*KeyTypeContext) IsKeyTypeContext() {}

func NewKeyTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyTypeContext {
	var p = new(KeyTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_keyType

	return p
}

func (s *KeyTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyTypeContext) INT() antlr.TerminalNode {
	return s.GetToken(TslParserINT, 0)
}

func (s *KeyTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(TslParserBOOL, 0)
}

func (s *KeyTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(TslParserSTRING, 0)
}

func (s *KeyTypeContext) TIME() antlr.TerminalNode {
	return s.GetToken(TslParserTIME, 0)
}

func (s *KeyTypeContext) DATE() antlr.TerminalNode {
	return s.GetToken(TslParserDATE, 0)
}

func (s *KeyTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterKeyType(s)
	}
}

func (s *KeyTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitKeyType(s)
	}
}

func (p *TslParser) KeyType() (localctx IKeyTypeContext) {
	localctx = NewKeyTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, TslParserRULE_keyType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(181)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7995392) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumDefContext is an interface to support dynamic dispatch.
type IEnumDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENUM() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	L_CURLY() antlr.TerminalNode
	R_CURLY() antlr.TerminalNode
	AllEnumItem() []IEnumItemContext
	EnumItem(i int) IEnumItemContext

	// IsEnumDefContext differentiates from other interfaces.
	IsEnumDefContext()
}

type EnumDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumDefContext() *EnumDefContext {
	var p = new(EnumDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_enumDef
	return p
}

func InitEmptyEnumDefContext(p *EnumDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_enumDef
}

func (*EnumDefContext) IsEnumDefContext() {}

func NewEnumDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDefContext {
	var p = new(EnumDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_enumDef

	return p
}

func (s *EnumDefContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDefContext) ENUM() antlr.TerminalNode {
	return s.GetToken(TslParserENUM, 0)
}

func (s *EnumDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TslParserIDENTIFIER, 0)
}

func (s *EnumDefContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(TslParserL_CURLY, 0)
}

func (s *EnumDefContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(TslParserR_CURLY, 0)
}

func (s *EnumDefContext) AllEnumItem() []IEnumItemContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnumItemContext); ok {
			len++
		}
	}

	tst := make([]IEnumItemContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnumItemContext); ok {
			tst[i] = t.(IEnumItemContext)
			i++
		}
	}

	return tst
}

func (s *EnumDefContext) EnumItem(i int) IEnumItemContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumItemContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumItemContext)
}

func (s *EnumDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterEnumDef(s)
	}
}

func (s *EnumDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitEnumDef(s)
	}
}

func (p *TslParser) EnumDef() (localctx IEnumDefContext) {
	localctx = NewEnumDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, TslParserRULE_enumDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(183)
		p.Match(TslParserENUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(184)
		p.Match(TslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(185)
		p.Match(TslParserL_CURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == TslParserIDENTIFIER {
		{
			p.SetState(186)
			p.EnumItem()
		}

		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(191)
		p.Match(TslParserR_CURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumItemContext is an interface to support dynamic dispatch.
type IEnumItemContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	EnumValue() IEnumValueContext
	R_PAREN() antlr.TerminalNode

	// IsEnumItemContext differentiates from other interfaces.
	IsEnumItemContext()
}

type EnumItemContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumItemContext() *EnumItemContext {
	var p = new(EnumItemContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_enumItem
	return p
}

func InitEmptyEnumItemContext(p *EnumItemContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_enumItem
}

func (*EnumItemContext) IsEnumItemContext() {}

func NewEnumItemContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumItemContext {
	var p = new(EnumItemContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_enumItem

	return p
}

func (s *EnumItemContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumItemContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TslParserIDENTIFIER, 0)
}

func (s *EnumItemContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserL_PAREN, 0)
}

func (s *EnumItemContext) EnumValue() IEnumValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumValueContext)
}

func (s *EnumItemContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserR_PAREN, 0)
}

func (s *EnumItemContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumItemContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumItemContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterEnumItem(s)
	}
}

func (s *EnumItemContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitEnumItem(s)
	}
}

func (p *TslParser) EnumItem() (localctx IEnumItemContext) {
	localctx = NewEnumItemContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, TslParserRULE_enumItem)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.Match(TslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TslParserL_PAREN {
		{
			p.SetState(194)
			p.Match(TslParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(195)
			p.EnumValue()
		}
		{
			p.SetState(196)
			p.Match(TslParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEnumValueContext is an interface to support dynamic dispatch.
type IEnumValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT_VALUE() antlr.TerminalNode
	STRING_VALUE() antlr.TerminalNode

	// IsEnumValueContext differentiates from other interfaces.
	IsEnumValueContext()
}

type EnumValueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueContext() *EnumValueContext {
	var p = new(EnumValueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_enumValue
	return p
}

func InitEmptyEnumValueContext(p *EnumValueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_enumValue
}

func (*EnumValueContext) IsEnumValueContext() {}

func NewEnumValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueContext {
	var p = new(EnumValueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_enumValue

	return p
}

func (s *EnumValueContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueContext) INT_VALUE() antlr.TerminalNode {
	return s.GetToken(TslParserINT_VALUE, 0)
}

func (s *EnumValueContext) STRING_VALUE() antlr.TerminalNode {
	return s.GetToken(TslParserSTRING_VALUE, 0)
}

func (s *EnumValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterEnumValue(s)
	}
}

func (s *EnumValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitEnumValue(s)
	}
}

func (p *TslParser) EnumValue() (localctx IEnumValueContext) {
	localctx = NewEnumValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, TslParserRULE_enumValue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(200)
		_la = p.GetTokenStream().LA(1)

		if !(_la == TslParserINT_VALUE || _la == TslParserSTRING_VALUE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IClassDefContext is an interface to support dynamic dispatch.
type IClassDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CLASS() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	L_CURLY() antlr.TerminalNode
	R_CURLY() antlr.TerminalNode
	AllProperty() []IPropertyContext
	Property(i int) IPropertyContext

	// IsClassDefContext differentiates from other interfaces.
	IsClassDefContext()
}

type ClassDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassDefContext() *ClassDefContext {
	var p = new(ClassDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_classDef
	return p
}

func InitEmptyClassDefContext(p *ClassDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_classDef
}

func (*ClassDefContext) IsClassDefContext() {}

func NewClassDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDefContext {
	var p = new(ClassDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_classDef

	return p
}

func (s *ClassDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassDefContext) CLASS() antlr.TerminalNode {
	return s.GetToken(TslParserCLASS, 0)
}

func (s *ClassDefContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TslParserIDENTIFIER, 0)
}

func (s *ClassDefContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(TslParserL_CURLY, 0)
}

func (s *ClassDefContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(TslParserR_CURLY, 0)
}

func (s *ClassDefContext) AllProperty() []IPropertyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPropertyContext); ok {
			len++
		}
	}

	tst := make([]IPropertyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPropertyContext); ok {
			tst[i] = t.(IPropertyContext)
			i++
		}
	}

	return tst
}

func (s *ClassDefContext) Property(i int) IPropertyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *ClassDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassDefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterClassDef(s)
	}
}

func (s *ClassDefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitClassDef(s)
	}
}

func (p *TslParser) ClassDef() (localctx IClassDefContext) {
	localctx = NewClassDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, TslParserRULE_classDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(202)
		p.Match(TslParserCLASS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(203)
		p.Match(TslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(204)
		p.Match(TslParserL_CURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(206)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == TslParserIDENTIFIER {
		{
			p.SetState(205)
			p.Property()
		}

		p.SetState(208)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(210)
		p.Match(TslParserR_CURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPropertyContext is an interface to support dynamic dispatch.
type IPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	ParamType() IParamTypeContext

	// IsPropertyContext differentiates from other interfaces.
	IsPropertyContext()
}

type PropertyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyContext() *PropertyContext {
	var p = new(PropertyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_property
	return p
}

func InitEmptyPropertyContext(p *PropertyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_property
}

func (*PropertyContext) IsPropertyContext() {}

func NewPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyContext {
	var p = new(PropertyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_property

	return p
}

func (s *PropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(TslParserIDENTIFIER, 0)
}

func (s *PropertyContext) ParamType() IParamTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParamTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParamTypeContext)
}

func (s *PropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterProperty(s)
	}
}

func (s *PropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitProperty(s)
	}
}

func (p *TslParser) Property() (localctx IPropertyContext) {
	localctx = NewPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, TslParserRULE_property)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(TslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.ParamType()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_CURLY() antlr.TerminalNode
	R_CURLY() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(TslParserL_CURLY, 0)
}

func (s *BlockContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(TslParserR_CURLY, 0)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *TslParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, TslParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(215)
		p.Match(TslParserL_CURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(219)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&536871952) != 0 {
		{
			p.SetState(216)
			p.Statement()
		}

		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(222)
		p.Match(TslParserR_CURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarDeclaration() IVarDeclarationContext
	IfStatement() IIfStatementContext
	FunctionCall() IFunctionCallContext
	ReturnStatement() IReturnStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) VarDeclaration() IVarDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarDeclarationContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *TslParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, TslParserRULE_statement)
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 16, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(224)
			p.VarDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(225)
			p.IfStatement()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(226)
			p.FunctionCall()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(227)
			p.ReturnStatement()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarDeclarationContext is an interface to support dynamic dispatch.
type IVarDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FqIdentifier() IFqIdentifierContext
	EQUAL() antlr.TerminalNode
	Expression() IExpressionContext

	// IsVarDeclarationContext differentiates from other interfaces.
	IsVarDeclarationContext()
}

type VarDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclarationContext() *VarDeclarationContext {
	var p = new(VarDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_varDeclaration
	return p
}

func InitEmptyVarDeclarationContext(p *VarDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_varDeclaration
}

func (*VarDeclarationContext) IsVarDeclarationContext() {}

func NewVarDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclarationContext {
	var p = new(VarDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_varDeclaration

	return p
}

func (s *VarDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclarationContext) FqIdentifier() IFqIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFqIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFqIdentifierContext)
}

func (s *VarDeclarationContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(TslParserEQUAL, 0)
}

func (s *VarDeclarationContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VarDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterVarDeclaration(s)
	}
}

func (s *VarDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitVarDeclaration(s)
	}
}

func (p *TslParser) VarDeclaration() (localctx IVarDeclarationContext) {
	localctx = NewVarDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, TslParserRULE_varDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(230)
		p.FqIdentifier()
	}
	{
		p.SetState(231)
		p.Match(TslParserEQUAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(232)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IF() antlr.TerminalNode
	Expression() IExpressionContext
	Block() IBlockContext
	AllElseIfBlock() []IElseIfBlockContext
	ElseIfBlock(i int) IElseIfBlockContext
	ElseBlock() IElseBlockContext

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) IF() antlr.TerminalNode {
	return s.GetToken(TslParserIF, 0)
}

func (s *IfStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStatementContext) AllElseIfBlock() []IElseIfBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IElseIfBlockContext); ok {
			len++
		}
	}

	tst := make([]IElseIfBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IElseIfBlockContext); ok {
			tst[i] = t.(IElseIfBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) ElseIfBlock(i int) IElseIfBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseIfBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseIfBlockContext)
}

func (s *IfStatementContext) ElseBlock() IElseBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IElseBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IElseBlockContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (p *TslParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, TslParserRULE_ifStatement)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(234)
		p.Match(TslParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(235)
		p.expression(0)
	}
	{
		p.SetState(236)
		p.Block()
	}
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(237)
				p.ElseIfBlock()
			}

		}
		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 17, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == TslParserELSE {
		{
			p.SetState(243)
			p.ElseBlock()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseIfBlockContext is an interface to support dynamic dispatch.
type IElseIfBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	IF() antlr.TerminalNode
	Expression() IExpressionContext
	Block() IBlockContext

	// IsElseIfBlockContext differentiates from other interfaces.
	IsElseIfBlockContext()
}

type ElseIfBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseIfBlockContext() *ElseIfBlockContext {
	var p = new(ElseIfBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_elseIfBlock
	return p
}

func InitEmptyElseIfBlockContext(p *ElseIfBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_elseIfBlock
}

func (*ElseIfBlockContext) IsElseIfBlockContext() {}

func NewElseIfBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseIfBlockContext {
	var p = new(ElseIfBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_elseIfBlock

	return p
}

func (s *ElseIfBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseIfBlockContext) ELSE() antlr.TerminalNode {
	return s.GetToken(TslParserELSE, 0)
}

func (s *ElseIfBlockContext) IF() antlr.TerminalNode {
	return s.GetToken(TslParserIF, 0)
}

func (s *ElseIfBlockContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ElseIfBlockContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseIfBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseIfBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseIfBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterElseIfBlock(s)
	}
}

func (s *ElseIfBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitElseIfBlock(s)
	}
}

func (p *TslParser) ElseIfBlock() (localctx IElseIfBlockContext) {
	localctx = NewElseIfBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, TslParserRULE_elseIfBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(246)
		p.Match(TslParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(247)
		p.Match(TslParserIF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(248)
		p.expression(0)
	}
	{
		p.SetState(249)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IElseBlockContext is an interface to support dynamic dispatch.
type IElseBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ELSE() antlr.TerminalNode
	Block() IBlockContext

	// IsElseBlockContext differentiates from other interfaces.
	IsElseBlockContext()
}

type ElseBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseBlockContext() *ElseBlockContext {
	var p = new(ElseBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_elseBlock
	return p
}

func InitEmptyElseBlockContext(p *ElseBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_elseBlock
}

func (*ElseBlockContext) IsElseBlockContext() {}

func NewElseBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseBlockContext {
	var p = new(ElseBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_elseBlock

	return p
}

func (s *ElseBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseBlockContext) ELSE() antlr.TerminalNode {
	return s.GetToken(TslParserELSE, 0)
}

func (s *ElseBlockContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ElseBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElseBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterElseBlock(s)
	}
}

func (s *ElseBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitElseBlock(s)
	}
}

func (p *TslParser) ElseBlock() (localctx IElseBlockContext) {
	localctx = NewElseBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, TslParserRULE_elseBlock)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(251)
		p.Match(TslParserELSE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(252)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FqIdentifier() IFqIdentifierContext
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) FqIdentifier() IFqIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFqIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFqIdentifierContext)
}

func (s *FunctionCallContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserL_PAREN, 0)
}

func (s *FunctionCallContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserR_PAREN, 0)
}

func (s *FunctionCallContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *FunctionCallContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FunctionCallContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(TslParserCOMMA)
}

func (s *FunctionCallContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(TslParserCOMMA, i)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (p *TslParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, TslParserRULE_functionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(254)
		p.FqIdentifier()
	}
	{
		p.SetState(255)
		p.Match(TslParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&307892783686400) != 0 {
		{
			p.SetState(256)
			p.expression(0)
		}
		p.SetState(261)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == TslParserCOMMA {
			{
				p.SetState(257)
				p.Match(TslParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(258)
				p.expression(0)
			}

			p.SetState(263)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(266)
		p.Match(TslParserR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RETURN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_returnStatement
	return p
}

func InitEmptyReturnStatementContext(p *ReturnStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_returnStatement
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) RETURN() antlr.TerminalNode {
	return s.GetToken(TslParserRETURN, 0)
}

func (s *ReturnStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (p *TslParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, TslParserRULE_returnStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(268)
		p.Match(TslParserRETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(269)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IdentifierExpression() IIdentifierExpressionContext
	ExpressionInParenthesis() IExpressionInParenthesisContext
	UnaryExpression() IUnaryExpressionContext
	FunctionCall() IFunctionCallContext
	ConstantValue() IConstantValueContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	STAR() antlr.TerminalNode
	SLASH() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	EQUAL() antlr.TerminalNode
	NOT_EQUAL() antlr.TerminalNode
	LESS_THAN() antlr.TerminalNode
	LESS_OR_EQUAL() antlr.TerminalNode
	GREATER_THAN() antlr.TerminalNode
	GREATER_OR_EQUAL() antlr.TerminalNode
	AllAND() []antlr.TerminalNode
	AND(i int) antlr.TerminalNode
	AllOR() []antlr.TerminalNode
	OR(i int) antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) IdentifierExpression() IIdentifierExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentifierExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentifierExpressionContext)
}

func (s *ExpressionContext) ExpressionInParenthesis() IExpressionInParenthesisContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionInParenthesisContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionInParenthesisContext)
}

func (s *ExpressionContext) UnaryExpression() IUnaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *ExpressionContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ExpressionContext) ConstantValue() IConstantValueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantValueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantValueContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) STAR() antlr.TerminalNode {
	return s.GetToken(TslParserSTAR, 0)
}

func (s *ExpressionContext) SLASH() antlr.TerminalNode {
	return s.GetToken(TslParserSLASH, 0)
}

func (s *ExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(TslParserPLUS, 0)
}

func (s *ExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(TslParserMINUS, 0)
}

func (s *ExpressionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(TslParserEQUAL, 0)
}

func (s *ExpressionContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(TslParserNOT_EQUAL, 0)
}

func (s *ExpressionContext) LESS_THAN() antlr.TerminalNode {
	return s.GetToken(TslParserLESS_THAN, 0)
}

func (s *ExpressionContext) LESS_OR_EQUAL() antlr.TerminalNode {
	return s.GetToken(TslParserLESS_OR_EQUAL, 0)
}

func (s *ExpressionContext) GREATER_THAN() antlr.TerminalNode {
	return s.GetToken(TslParserGREATER_THAN, 0)
}

func (s *ExpressionContext) GREATER_OR_EQUAL() antlr.TerminalNode {
	return s.GetToken(TslParserGREATER_OR_EQUAL, 0)
}

func (s *ExpressionContext) AllAND() []antlr.TerminalNode {
	return s.GetTokens(TslParserAND)
}

func (s *ExpressionContext) AND(i int) antlr.TerminalNode {
	return s.GetToken(TslParserAND, i)
}

func (s *ExpressionContext) AllOR() []antlr.TerminalNode {
	return s.GetTokens(TslParserOR)
}

func (s *ExpressionContext) OR(i int) antlr.TerminalNode {
	return s.GetToken(TslParserOR, i)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *TslParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *TslParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 56
	p.EnterRecursionRule(localctx, 56, TslParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(272)
			p.IdentifierExpression()
		}

	case 2:
		{
			p.SetState(273)
			p.ExpressionInParenthesis()
		}

	case 3:
		{
			p.SetState(274)
			p.UnaryExpression()
		}

	case 4:
		{
			p.SetState(275)
			p.FunctionCall()
		}

	case 5:
		{
			p.SetState(276)
			p.ConstantValue()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(302)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TslParserRULE_expression)
				p.SetState(279)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(280)
					_la = p.GetTokenStream().LA(1)

					if !(_la == TslParserSTAR || _la == TslParserSLASH) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(281)
					p.expression(6)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TslParserRULE_expression)
				p.SetState(282)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(283)
					_la = p.GetTokenStream().LA(1)

					if !(_la == TslParserPLUS || _la == TslParserMINUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(284)
					p.expression(5)
				}

			case 3:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TslParserRULE_expression)
				p.SetState(285)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(286)
					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&17452548067688448) != 0) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(287)
					p.expression(4)
				}

			case 4:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TslParserRULE_expression)
				p.SetState(288)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				p.SetState(291)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = 1
				for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					switch _alt {
					case 1:
						{
							p.SetState(289)
							p.Match(TslParserAND)
							if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
							}
						}
						{
							p.SetState(290)
							p.expression(0)
						}

					default:
						p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
						goto errorExit
					}

					p.SetState(293)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 22, p.GetParserRuleContext())
					if p.HasError() {
						goto errorExit
					}
				}

			case 5:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, TslParserRULE_expression)
				p.SetState(295)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				p.SetState(298)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_alt = 1
				for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
					switch _alt {
					case 1:
						{
							p.SetState(296)
							p.Match(TslParserOR)
							if p.HasError() {
								// Recognition error - abort rule
								goto errorExit
							}
						}
						{
							p.SetState(297)
							p.expression(0)
						}

					default:
						p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
						goto errorExit
					}

					p.SetState(300)
					p.GetErrorHandler().Sync(p)
					_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 23, p.GetParserRuleContext())
					if p.HasError() {
						goto errorExit
					}
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(306)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIdentifierExpressionContext is an interface to support dynamic dispatch.
type IIdentifierExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FqIdentifier() IFqIdentifierContext
	IN_BAR() antlr.TerminalNode
	Expression() IExpressionContext

	// IsIdentifierExpressionContext differentiates from other interfaces.
	IsIdentifierExpressionContext()
}

type IdentifierExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentifierExpressionContext() *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_identifierExpression
	return p
}

func InitEmptyIdentifierExpressionContext(p *IdentifierExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_identifierExpression
}

func (*IdentifierExpressionContext) IsIdentifierExpressionContext() {}

func NewIdentifierExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_identifierExpression

	return p
}

func (s *IdentifierExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentifierExpressionContext) FqIdentifier() IFqIdentifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFqIdentifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFqIdentifierContext)
}

func (s *IdentifierExpressionContext) IN_BAR() antlr.TerminalNode {
	return s.GetToken(TslParserIN_BAR, 0)
}

func (s *IdentifierExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentifierExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitIdentifierExpression(s)
	}
}

func (p *TslParser) IdentifierExpression() (localctx IIdentifierExpressionContext) {
	localctx = NewIdentifierExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, TslParserRULE_identifierExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		p.FqIdentifier()
	}
	p.SetState(310)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 26, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(308)
			p.Match(TslParserIN_BAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(309)
			p.expression(0)
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFqIdentifierContext is an interface to support dynamic dispatch.
type IFqIdentifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIDENTIFIER() []antlr.TerminalNode
	IDENTIFIER(i int) antlr.TerminalNode
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsFqIdentifierContext differentiates from other interfaces.
	IsFqIdentifierContext()
}

type FqIdentifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFqIdentifierContext() *FqIdentifierContext {
	var p = new(FqIdentifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_fqIdentifier
	return p
}

func InitEmptyFqIdentifierContext(p *FqIdentifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_fqIdentifier
}

func (*FqIdentifierContext) IsFqIdentifierContext() {}

func NewFqIdentifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FqIdentifierContext {
	var p = new(FqIdentifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_fqIdentifier

	return p
}

func (s *FqIdentifierContext) GetParser() antlr.Parser { return s.parser }

func (s *FqIdentifierContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(TslParserIDENTIFIER)
}

func (s *FqIdentifierContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(TslParserIDENTIFIER, i)
}

func (s *FqIdentifierContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(TslParserDOT)
}

func (s *FqIdentifierContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(TslParserDOT, i)
}

func (s *FqIdentifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FqIdentifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FqIdentifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterFqIdentifier(s)
	}
}

func (s *FqIdentifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitFqIdentifier(s)
	}
}

func (p *TslParser) FqIdentifier() (localctx IFqIdentifierContext) {
	localctx = NewFqIdentifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, TslParserRULE_fqIdentifier)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(312)
		p.Match(TslParserIDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(317)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(313)
				p.Match(TslParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(314)
				p.Match(TslParserIDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(319)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryExpressionContext is an interface to support dynamic dispatch.
type IUnaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	MINUS() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	NOT() antlr.TerminalNode

	// IsUnaryExpressionContext differentiates from other interfaces.
	IsUnaryExpressionContext()
}

type UnaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExpressionContext() *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_unaryExpression
	return p
}

func InitEmptyUnaryExpressionContext(p *UnaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_unaryExpression
}

func (*UnaryExpressionContext) IsUnaryExpressionContext() {}

func NewUnaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_unaryExpression

	return p
}

func (s *UnaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(TslParserMINUS, 0)
}

func (s *UnaryExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(TslParserPLUS, 0)
}

func (s *UnaryExpressionContext) NOT() antlr.TerminalNode {
	return s.GetToken(TslParserNOT, 0)
}

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

func (p *TslParser) UnaryExpression() (localctx IUnaryExpressionContext) {
	localctx = NewUnaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, TslParserRULE_unaryExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&26388279074816) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(321)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionInParenthesisContext is an interface to support dynamic dispatch.
type IExpressionInParenthesisContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_PAREN() antlr.TerminalNode
	Expression() IExpressionContext
	R_PAREN() antlr.TerminalNode

	// IsExpressionInParenthesisContext differentiates from other interfaces.
	IsExpressionInParenthesisContext()
}

type ExpressionInParenthesisContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionInParenthesisContext() *ExpressionInParenthesisContext {
	var p = new(ExpressionInParenthesisContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_expressionInParenthesis
	return p
}

func InitEmptyExpressionInParenthesisContext(p *ExpressionInParenthesisContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = TslParserRULE_expressionInParenthesis
}

func (*ExpressionInParenthesisContext) IsExpressionInParenthesisContext() {}

func NewExpressionInParenthesisContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionInParenthesisContext {
	var p = new(ExpressionInParenthesisContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = TslParserRULE_expressionInParenthesis

	return p
}

func (s *ExpressionInParenthesisContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionInParenthesisContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserL_PAREN, 0)
}

func (s *ExpressionInParenthesisContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionInParenthesisContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(TslParserR_PAREN, 0)
}

func (s *ExpressionInParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionInParenthesisContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionInParenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.EnterExpressionInParenthesis(s)
	}
}

func (s *ExpressionInParenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TslParserListener); ok {
		listenerT.ExitExpressionInParenthesis(s)
	}
}

func (p *TslParser) ExpressionInParenthesis() (localctx IExpressionInParenthesisContext) {
	localctx = NewExpressionInParenthesisContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, TslParserRULE_expressionInParenthesis)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(323)
		p.Match(TslParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(324)
		p.expression(0)
	}
	{
		p.SetState(325)
		p.Match(TslParserR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *TslParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 28:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TslParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
