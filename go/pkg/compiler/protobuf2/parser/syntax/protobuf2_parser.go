// Code generated from Protobuf2.g4 by ANTLR 4.13.0. DO NOT EDIT.

package syntax // Protobuf2
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

type Protobuf2Parser struct {
	*antlr.BaseParser
}

var Protobuf2ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func protobuf2ParserInit() {
	staticData := &Protobuf2ParserStaticData
	staticData.LiteralNames = []string{
		"", "'syntax'", "'import'", "'weak'", "'public'", "'package'", "'option'",
		"'repeated'", "'optional'", "'required'", "'group'", "'oneof'", "'map'",
		"'int32'", "'int64'", "'uint32'", "'uint64'", "'sint32'", "'sint64'",
		"'fixed32'", "'fixed64'", "'sfixed32'", "'sfixed64'", "'bool'", "'string'",
		"'double'", "'float'", "'bytes'", "'reserved'", "'extensions'", "'to'",
		"'max'", "'enum'", "'extend'", "'message'", "'service'", "'rpc'", "'stream'",
		"'returns'", "''proto2''", "'\"proto2\"'", "';'", "'='", "'('", "')'",
		"'['", "']'", "'{'", "'}'", "'<'", "'>'", "'.'", "','", "':'", "'+'",
		"'-'",
	}
	staticData.SymbolicNames = []string{
		"", "SYNTAX", "IMPORT", "WEAK", "PUBLIC", "PACKAGE", "OPTION", "REPEATED",
		"OPTIONAL", "REQUIRED", "GROUP", "ONEOF", "MAP", "INT32", "INT64", "UINT32",
		"UINT64", "SINT32", "SINT64", "FIXED32", "FIXED64", "SFIXED32", "SFIXED64",
		"BOOL", "STRING", "DOUBLE", "FLOAT", "BYTES", "RESERVED", "EXTENSIONS",
		"TO", "MAX", "ENUM", "EXTEND", "MESSAGE", "SERVICE", "RPC", "STREAM",
		"RETURNS", "PROTO2_LIT_SINGLE", "PROTO2_LIT_DOBULE", "SEMI", "EQ", "LP",
		"RP", "LB", "RB", "LC", "RC", "LT", "GT", "DOT", "COMMA", "COLON", "PLUS",
		"MINUS", "STR_LIT", "BOOL_LIT", "FLOAT_LIT", "INT_LIT", "IDENTIFIER",
		"WS", "LINE_COMMENT", "COMMENT",
	}
	staticData.RuleNames = []string{
		"proto", "syntax", "importStatement", "packageStatement", "optionStatement",
		"optionName", "fieldLabel", "field", "fieldOptions", "fieldOption",
		"fieldNumber", "group", "oneof", "oneofField", "mapField", "keyType",
		"type_", "extensions", "reserved", "ranges", "range_", "reservedFieldNames",
		"topLevelDef", "enumDef", "enumBody", "enumElement", "enumField", "enumValueOptions",
		"enumValueOption", "messageDef", "messageBody", "messageElement", "extendDef",
		"extendElement", "serviceDef", "serviceElement", "rpc", "stream", "constant",
		"blockLit", "emptyStatement_", "ident", "fullIdent", "messageName",
		"enumName", "fieldName", "groupName", "oneofName", "mapName", "serviceName",
		"rpcName", "streamName", "messageType", "enumType", "intLit", "strLit",
		"boolLit", "floatLit", "keywords",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 63, 551, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		2, 58, 7, 58, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 125, 8, 0, 10,
		0, 12, 0, 128, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1,
		2, 3, 2, 139, 8, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 160,
		8, 5, 1, 5, 1, 5, 3, 5, 164, 8, 5, 3, 5, 166, 8, 5, 1, 6, 1, 6, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 179, 8, 7, 1, 7, 1,
		7, 1, 8, 1, 8, 1, 8, 5, 8, 186, 8, 8, 10, 8, 12, 8, 189, 9, 8, 1, 9, 1,
		9, 1, 9, 1, 9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11,
		1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 5, 12, 210, 8, 12, 10,
		12, 12, 12, 213, 9, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 13, 1, 13,
		1, 13, 1, 13, 1, 13, 3, 13, 225, 8, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		3, 14, 242, 8, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1,
		16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16, 1, 16,
		1, 16, 1, 16, 1, 16, 3, 16, 265, 8, 16, 1, 17, 1, 17, 1, 17, 1, 17, 1,
		18, 1, 18, 1, 18, 3, 18, 274, 8, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19,
		5, 19, 281, 8, 19, 10, 19, 12, 19, 284, 9, 19, 1, 20, 1, 20, 1, 20, 1,
		20, 3, 20, 290, 8, 20, 3, 20, 292, 8, 20, 1, 21, 1, 21, 1, 21, 5, 21, 297,
		8, 21, 10, 21, 12, 21, 300, 9, 21, 1, 22, 1, 22, 1, 22, 1, 22, 3, 22, 306,
		8, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 5, 24, 314, 8, 24, 10,
		24, 12, 24, 317, 9, 24, 1, 24, 1, 24, 1, 25, 1, 25, 1, 25, 3, 25, 324,
		8, 25, 1, 26, 1, 26, 1, 26, 3, 26, 329, 8, 26, 1, 26, 1, 26, 3, 26, 333,
		8, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 1, 27, 5, 27, 341, 8, 27, 10,
		27, 12, 27, 344, 9, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 29,
		1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 5, 30, 358, 8, 30, 10, 30, 12, 30, 361,
		9, 30, 1, 30, 1, 30, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1,
		31, 1, 31, 1, 31, 1, 31, 3, 31, 376, 8, 31, 1, 32, 1, 32, 1, 32, 1, 32,
		5, 32, 382, 8, 32, 10, 32, 12, 32, 385, 9, 32, 1, 32, 1, 32, 1, 33, 1,
		33, 1, 33, 3, 33, 392, 8, 33, 1, 34, 1, 34, 1, 34, 1, 34, 5, 34, 398, 8,
		34, 10, 34, 12, 34, 401, 9, 34, 1, 34, 1, 34, 1, 35, 1, 35, 1, 35, 1, 35,
		3, 35, 409, 8, 35, 1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 415, 8, 36, 1, 36,
		1, 36, 1, 36, 1, 36, 1, 36, 3, 36, 422, 8, 36, 1, 36, 1, 36, 1, 36, 1,
		36, 1, 36, 5, 36, 429, 8, 36, 10, 36, 12, 36, 432, 9, 36, 1, 36, 1, 36,
		3, 36, 436, 8, 36, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1, 37, 1,
		37, 1, 37, 1, 37, 5, 37, 448, 8, 37, 10, 37, 12, 37, 451, 9, 37, 1, 37,
		1, 37, 3, 37, 455, 8, 37, 1, 38, 1, 38, 3, 38, 459, 8, 38, 1, 38, 1, 38,
		3, 38, 463, 8, 38, 1, 38, 1, 38, 1, 38, 1, 38, 3, 38, 469, 8, 38, 1, 39,
		1, 39, 1, 39, 1, 39, 1, 39, 5, 39, 476, 8, 39, 10, 39, 12, 39, 479, 9,
		39, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 3, 41, 487, 8, 41, 1, 42,
		1, 42, 1, 42, 5, 42, 492, 8, 42, 10, 42, 12, 42, 495, 9, 42, 1, 43, 1,
		43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 46, 1, 46, 1, 47, 1, 47, 1, 48, 1, 48,
		1, 49, 1, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 52, 3, 52, 516, 8, 52, 1,
		52, 1, 52, 1, 52, 5, 52, 521, 8, 52, 10, 52, 12, 52, 524, 9, 52, 1, 52,
		1, 52, 1, 53, 3, 53, 529, 8, 53, 1, 53, 1, 53, 1, 53, 5, 53, 534, 8, 53,
		10, 53, 12, 53, 537, 9, 53, 1, 53, 1, 53, 1, 54, 1, 54, 1, 55, 1, 55, 1,
		56, 1, 56, 1, 57, 1, 57, 1, 58, 1, 58, 1, 58, 0, 0, 59, 0, 2, 4, 6, 8,
		10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44,
		46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80,
		82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112,
		114, 116, 0, 7, 1, 0, 39, 40, 1, 0, 3, 4, 1, 0, 7, 9, 1, 0, 13, 24, 1,
		0, 54, 55, 2, 0, 39, 40, 56, 56, 2, 0, 1, 38, 57, 57, 577, 0, 118, 1, 0,
		0, 0, 2, 131, 1, 0, 0, 0, 4, 136, 1, 0, 0, 0, 6, 143, 1, 0, 0, 0, 8, 147,
		1, 0, 0, 0, 10, 165, 1, 0, 0, 0, 12, 167, 1, 0, 0, 0, 14, 169, 1, 0, 0,
		0, 16, 182, 1, 0, 0, 0, 18, 190, 1, 0, 0, 0, 20, 194, 1, 0, 0, 0, 22, 196,
		1, 0, 0, 0, 24, 203, 1, 0, 0, 0, 26, 216, 1, 0, 0, 0, 28, 228, 1, 0, 0,
		0, 30, 245, 1, 0, 0, 0, 32, 264, 1, 0, 0, 0, 34, 266, 1, 0, 0, 0, 36, 270,
		1, 0, 0, 0, 38, 277, 1, 0, 0, 0, 40, 285, 1, 0, 0, 0, 42, 293, 1, 0, 0,
		0, 44, 305, 1, 0, 0, 0, 46, 307, 1, 0, 0, 0, 48, 311, 1, 0, 0, 0, 50, 323,
		1, 0, 0, 0, 52, 325, 1, 0, 0, 0, 54, 336, 1, 0, 0, 0, 56, 347, 1, 0, 0,
		0, 58, 351, 1, 0, 0, 0, 60, 355, 1, 0, 0, 0, 62, 375, 1, 0, 0, 0, 64, 377,
		1, 0, 0, 0, 66, 391, 1, 0, 0, 0, 68, 393, 1, 0, 0, 0, 70, 408, 1, 0, 0,
		0, 72, 410, 1, 0, 0, 0, 74, 437, 1, 0, 0, 0, 76, 468, 1, 0, 0, 0, 78, 470,
		1, 0, 0, 0, 80, 482, 1, 0, 0, 0, 82, 486, 1, 0, 0, 0, 84, 488, 1, 0, 0,
		0, 86, 496, 1, 0, 0, 0, 88, 498, 1, 0, 0, 0, 90, 500, 1, 0, 0, 0, 92, 502,
		1, 0, 0, 0, 94, 504, 1, 0, 0, 0, 96, 506, 1, 0, 0, 0, 98, 508, 1, 0, 0,
		0, 100, 510, 1, 0, 0, 0, 102, 512, 1, 0, 0, 0, 104, 515, 1, 0, 0, 0, 106,
		528, 1, 0, 0, 0, 108, 540, 1, 0, 0, 0, 110, 542, 1, 0, 0, 0, 112, 544,
		1, 0, 0, 0, 114, 546, 1, 0, 0, 0, 116, 548, 1, 0, 0, 0, 118, 126, 3, 2,
		1, 0, 119, 125, 3, 4, 2, 0, 120, 125, 3, 6, 3, 0, 121, 125, 3, 8, 4, 0,
		122, 125, 3, 44, 22, 0, 123, 125, 3, 80, 40, 0, 124, 119, 1, 0, 0, 0, 124,
		120, 1, 0, 0, 0, 124, 121, 1, 0, 0, 0, 124, 122, 1, 0, 0, 0, 124, 123,
		1, 0, 0, 0, 125, 128, 1, 0, 0, 0, 126, 124, 1, 0, 0, 0, 126, 127, 1, 0,
		0, 0, 127, 129, 1, 0, 0, 0, 128, 126, 1, 0, 0, 0, 129, 130, 5, 0, 0, 1,
		130, 1, 1, 0, 0, 0, 131, 132, 5, 1, 0, 0, 132, 133, 5, 42, 0, 0, 133, 134,
		7, 0, 0, 0, 134, 135, 5, 41, 0, 0, 135, 3, 1, 0, 0, 0, 136, 138, 5, 2,
		0, 0, 137, 139, 7, 1, 0, 0, 138, 137, 1, 0, 0, 0, 138, 139, 1, 0, 0, 0,
		139, 140, 1, 0, 0, 0, 140, 141, 3, 110, 55, 0, 141, 142, 5, 41, 0, 0, 142,
		5, 1, 0, 0, 0, 143, 144, 5, 5, 0, 0, 144, 145, 3, 84, 42, 0, 145, 146,
		5, 41, 0, 0, 146, 7, 1, 0, 0, 0, 147, 148, 5, 6, 0, 0, 148, 149, 3, 10,
		5, 0, 149, 150, 5, 42, 0, 0, 150, 151, 3, 76, 38, 0, 151, 152, 5, 41, 0,
		0, 152, 9, 1, 0, 0, 0, 153, 166, 3, 84, 42, 0, 154, 160, 3, 82, 41, 0,
		155, 156, 5, 43, 0, 0, 156, 157, 3, 84, 42, 0, 157, 158, 5, 44, 0, 0, 158,
		160, 1, 0, 0, 0, 159, 154, 1, 0, 0, 0, 159, 155, 1, 0, 0, 0, 160, 163,
		1, 0, 0, 0, 161, 162, 5, 51, 0, 0, 162, 164, 3, 84, 42, 0, 163, 161, 1,
		0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 166, 1, 0, 0, 0, 165, 153, 1, 0, 0,
		0, 165, 159, 1, 0, 0, 0, 166, 11, 1, 0, 0, 0, 167, 168, 7, 2, 0, 0, 168,
		13, 1, 0, 0, 0, 169, 170, 3, 12, 6, 0, 170, 171, 3, 32, 16, 0, 171, 172,
		3, 90, 45, 0, 172, 173, 5, 42, 0, 0, 173, 178, 3, 20, 10, 0, 174, 175,
		5, 45, 0, 0, 175, 176, 3, 16, 8, 0, 176, 177, 5, 46, 0, 0, 177, 179, 1,
		0, 0, 0, 178, 174, 1, 0, 0, 0, 178, 179, 1, 0, 0, 0, 179, 180, 1, 0, 0,
		0, 180, 181, 5, 41, 0, 0, 181, 15, 1, 0, 0, 0, 182, 187, 3, 18, 9, 0, 183,
		184, 5, 52, 0, 0, 184, 186, 3, 18, 9, 0, 185, 183, 1, 0, 0, 0, 186, 189,
		1, 0, 0, 0, 187, 185, 1, 0, 0, 0, 187, 188, 1, 0, 0, 0, 188, 17, 1, 0,
		0, 0, 189, 187, 1, 0, 0, 0, 190, 191, 3, 10, 5, 0, 191, 192, 5, 42, 0,
		0, 192, 193, 3, 76, 38, 0, 193, 19, 1, 0, 0, 0, 194, 195, 3, 108, 54, 0,
		195, 21, 1, 0, 0, 0, 196, 197, 3, 12, 6, 0, 197, 198, 5, 10, 0, 0, 198,
		199, 3, 92, 46, 0, 199, 200, 5, 42, 0, 0, 200, 201, 3, 20, 10, 0, 201,
		202, 3, 60, 30, 0, 202, 23, 1, 0, 0, 0, 203, 204, 5, 11, 0, 0, 204, 205,
		3, 94, 47, 0, 205, 211, 5, 47, 0, 0, 206, 210, 3, 8, 4, 0, 207, 210, 3,
		26, 13, 0, 208, 210, 3, 80, 40, 0, 209, 206, 1, 0, 0, 0, 209, 207, 1, 0,
		0, 0, 209, 208, 1, 0, 0, 0, 210, 213, 1, 0, 0, 0, 211, 209, 1, 0, 0, 0,
		211, 212, 1, 0, 0, 0, 212, 214, 1, 0, 0, 0, 213, 211, 1, 0, 0, 0, 214,
		215, 5, 48, 0, 0, 215, 25, 1, 0, 0, 0, 216, 217, 3, 32, 16, 0, 217, 218,
		3, 90, 45, 0, 218, 219, 5, 42, 0, 0, 219, 224, 3, 20, 10, 0, 220, 221,
		5, 45, 0, 0, 221, 222, 3, 16, 8, 0, 222, 223, 5, 46, 0, 0, 223, 225, 1,
		0, 0, 0, 224, 220, 1, 0, 0, 0, 224, 225, 1, 0, 0, 0, 225, 226, 1, 0, 0,
		0, 226, 227, 5, 41, 0, 0, 227, 27, 1, 0, 0, 0, 228, 229, 5, 12, 0, 0, 229,
		230, 5, 49, 0, 0, 230, 231, 3, 30, 15, 0, 231, 232, 5, 52, 0, 0, 232, 233,
		3, 32, 16, 0, 233, 234, 5, 50, 0, 0, 234, 235, 3, 96, 48, 0, 235, 236,
		5, 42, 0, 0, 236, 241, 3, 20, 10, 0, 237, 238, 5, 45, 0, 0, 238, 239, 3,
		16, 8, 0, 239, 240, 5, 46, 0, 0, 240, 242, 1, 0, 0, 0, 241, 237, 1, 0,
		0, 0, 241, 242, 1, 0, 0, 0, 242, 243, 1, 0, 0, 0, 243, 244, 5, 41, 0, 0,
		244, 29, 1, 0, 0, 0, 245, 246, 7, 3, 0, 0, 246, 31, 1, 0, 0, 0, 247, 265,
		5, 25, 0, 0, 248, 265, 5, 26, 0, 0, 249, 265, 5, 13, 0, 0, 250, 265, 5,
		14, 0, 0, 251, 265, 5, 15, 0, 0, 252, 265, 5, 16, 0, 0, 253, 265, 5, 17,
		0, 0, 254, 265, 5, 18, 0, 0, 255, 265, 5, 19, 0, 0, 256, 265, 5, 20, 0,
		0, 257, 265, 5, 21, 0, 0, 258, 265, 5, 22, 0, 0, 259, 265, 5, 23, 0, 0,
		260, 265, 5, 24, 0, 0, 261, 265, 5, 27, 0, 0, 262, 265, 3, 104, 52, 0,
		263, 265, 3, 106, 53, 0, 264, 247, 1, 0, 0, 0, 264, 248, 1, 0, 0, 0, 264,
		249, 1, 0, 0, 0, 264, 250, 1, 0, 0, 0, 264, 251, 1, 0, 0, 0, 264, 252,
		1, 0, 0, 0, 264, 253, 1, 0, 0, 0, 264, 254, 1, 0, 0, 0, 264, 255, 1, 0,
		0, 0, 264, 256, 1, 0, 0, 0, 264, 257, 1, 0, 0, 0, 264, 258, 1, 0, 0, 0,
		264, 259, 1, 0, 0, 0, 264, 260, 1, 0, 0, 0, 264, 261, 1, 0, 0, 0, 264,
		262, 1, 0, 0, 0, 264, 263, 1, 0, 0, 0, 265, 33, 1, 0, 0, 0, 266, 267, 5,
		29, 0, 0, 267, 268, 3, 38, 19, 0, 268, 269, 5, 41, 0, 0, 269, 35, 1, 0,
		0, 0, 270, 273, 5, 28, 0, 0, 271, 274, 3, 38, 19, 0, 272, 274, 3, 42, 21,
		0, 273, 271, 1, 0, 0, 0, 273, 272, 1, 0, 0, 0, 274, 275, 1, 0, 0, 0, 275,
		276, 5, 41, 0, 0, 276, 37, 1, 0, 0, 0, 277, 282, 3, 40, 20, 0, 278, 279,
		5, 52, 0, 0, 279, 281, 3, 40, 20, 0, 280, 278, 1, 0, 0, 0, 281, 284, 1,
		0, 0, 0, 282, 280, 1, 0, 0, 0, 282, 283, 1, 0, 0, 0, 283, 39, 1, 0, 0,
		0, 284, 282, 1, 0, 0, 0, 285, 291, 3, 108, 54, 0, 286, 289, 5, 30, 0, 0,
		287, 290, 3, 108, 54, 0, 288, 290, 5, 31, 0, 0, 289, 287, 1, 0, 0, 0, 289,
		288, 1, 0, 0, 0, 290, 292, 1, 0, 0, 0, 291, 286, 1, 0, 0, 0, 291, 292,
		1, 0, 0, 0, 292, 41, 1, 0, 0, 0, 293, 298, 3, 110, 55, 0, 294, 295, 5,
		52, 0, 0, 295, 297, 3, 110, 55, 0, 296, 294, 1, 0, 0, 0, 297, 300, 1, 0,
		0, 0, 298, 296, 1, 0, 0, 0, 298, 299, 1, 0, 0, 0, 299, 43, 1, 0, 0, 0,
		300, 298, 1, 0, 0, 0, 301, 306, 3, 58, 29, 0, 302, 306, 3, 46, 23, 0, 303,
		306, 3, 68, 34, 0, 304, 306, 3, 64, 32, 0, 305, 301, 1, 0, 0, 0, 305, 302,
		1, 0, 0, 0, 305, 303, 1, 0, 0, 0, 305, 304, 1, 0, 0, 0, 306, 45, 1, 0,
		0, 0, 307, 308, 5, 32, 0, 0, 308, 309, 3, 88, 44, 0, 309, 310, 3, 48, 24,
		0, 310, 47, 1, 0, 0, 0, 311, 315, 5, 47, 0, 0, 312, 314, 3, 50, 25, 0,
		313, 312, 1, 0, 0, 0, 314, 317, 1, 0, 0, 0, 315, 313, 1, 0, 0, 0, 315,
		316, 1, 0, 0, 0, 316, 318, 1, 0, 0, 0, 317, 315, 1, 0, 0, 0, 318, 319,
		5, 48, 0, 0, 319, 49, 1, 0, 0, 0, 320, 324, 3, 8, 4, 0, 321, 324, 3, 52,
		26, 0, 322, 324, 3, 80, 40, 0, 323, 320, 1, 0, 0, 0, 323, 321, 1, 0, 0,
		0, 323, 322, 1, 0, 0, 0, 324, 51, 1, 0, 0, 0, 325, 326, 3, 82, 41, 0, 326,
		328, 5, 42, 0, 0, 327, 329, 5, 55, 0, 0, 328, 327, 1, 0, 0, 0, 328, 329,
		1, 0, 0, 0, 329, 330, 1, 0, 0, 0, 330, 332, 3, 108, 54, 0, 331, 333, 3,
		54, 27, 0, 332, 331, 1, 0, 0, 0, 332, 333, 1, 0, 0, 0, 333, 334, 1, 0,
		0, 0, 334, 335, 5, 41, 0, 0, 335, 53, 1, 0, 0, 0, 336, 337, 5, 45, 0, 0,
		337, 342, 3, 56, 28, 0, 338, 339, 5, 52, 0, 0, 339, 341, 3, 56, 28, 0,
		340, 338, 1, 0, 0, 0, 341, 344, 1, 0, 0, 0, 342, 340, 1, 0, 0, 0, 342,
		343, 1, 0, 0, 0, 343, 345, 1, 0, 0, 0, 344, 342, 1, 0, 0, 0, 345, 346,
		5, 46, 0, 0, 346, 55, 1, 0, 0, 0, 347, 348, 3, 10, 5, 0, 348, 349, 5, 42,
		0, 0, 349, 350, 3, 76, 38, 0, 350, 57, 1, 0, 0, 0, 351, 352, 5, 34, 0,
		0, 352, 353, 3, 86, 43, 0, 353, 354, 3, 60, 30, 0, 354, 59, 1, 0, 0, 0,
		355, 359, 5, 47, 0, 0, 356, 358, 3, 62, 31, 0, 357, 356, 1, 0, 0, 0, 358,
		361, 1, 0, 0, 0, 359, 357, 1, 0, 0, 0, 359, 360, 1, 0, 0, 0, 360, 362,
		1, 0, 0, 0, 361, 359, 1, 0, 0, 0, 362, 363, 5, 48, 0, 0, 363, 61, 1, 0,
		0, 0, 364, 376, 3, 14, 7, 0, 365, 376, 3, 46, 23, 0, 366, 376, 3, 58, 29,
		0, 367, 376, 3, 64, 32, 0, 368, 376, 3, 8, 4, 0, 369, 376, 3, 24, 12, 0,
		370, 376, 3, 28, 14, 0, 371, 376, 3, 34, 17, 0, 372, 376, 3, 22, 11, 0,
		373, 376, 3, 36, 18, 0, 374, 376, 3, 80, 40, 0, 375, 364, 1, 0, 0, 0, 375,
		365, 1, 0, 0, 0, 375, 366, 1, 0, 0, 0, 375, 367, 1, 0, 0, 0, 375, 368,
		1, 0, 0, 0, 375, 369, 1, 0, 0, 0, 375, 370, 1, 0, 0, 0, 375, 371, 1, 0,
		0, 0, 375, 372, 1, 0, 0, 0, 375, 373, 1, 0, 0, 0, 375, 374, 1, 0, 0, 0,
		376, 63, 1, 0, 0, 0, 377, 378, 5, 33, 0, 0, 378, 379, 3, 104, 52, 0, 379,
		383, 5, 47, 0, 0, 380, 382, 3, 66, 33, 0, 381, 380, 1, 0, 0, 0, 382, 385,
		1, 0, 0, 0, 383, 381, 1, 0, 0, 0, 383, 384, 1, 0, 0, 0, 384, 386, 1, 0,
		0, 0, 385, 383, 1, 0, 0, 0, 386, 387, 5, 48, 0, 0, 387, 65, 1, 0, 0, 0,
		388, 392, 3, 14, 7, 0, 389, 392, 3, 22, 11, 0, 390, 392, 3, 80, 40, 0,
		391, 388, 1, 0, 0, 0, 391, 389, 1, 0, 0, 0, 391, 390, 1, 0, 0, 0, 392,
		67, 1, 0, 0, 0, 393, 394, 5, 35, 0, 0, 394, 395, 3, 98, 49, 0, 395, 399,
		5, 47, 0, 0, 396, 398, 3, 70, 35, 0, 397, 396, 1, 0, 0, 0, 398, 401, 1,
		0, 0, 0, 399, 397, 1, 0, 0, 0, 399, 400, 1, 0, 0, 0, 400, 402, 1, 0, 0,
		0, 401, 399, 1, 0, 0, 0, 402, 403, 5, 48, 0, 0, 403, 69, 1, 0, 0, 0, 404,
		409, 3, 8, 4, 0, 405, 409, 3, 72, 36, 0, 406, 409, 3, 74, 37, 0, 407, 409,
		3, 80, 40, 0, 408, 404, 1, 0, 0, 0, 408, 405, 1, 0, 0, 0, 408, 406, 1,
		0, 0, 0, 408, 407, 1, 0, 0, 0, 409, 71, 1, 0, 0, 0, 410, 411, 5, 36, 0,
		0, 411, 412, 3, 100, 50, 0, 412, 414, 5, 43, 0, 0, 413, 415, 5, 37, 0,
		0, 414, 413, 1, 0, 0, 0, 414, 415, 1, 0, 0, 0, 415, 416, 1, 0, 0, 0, 416,
		417, 3, 104, 52, 0, 417, 418, 5, 44, 0, 0, 418, 419, 5, 38, 0, 0, 419,
		421, 5, 43, 0, 0, 420, 422, 5, 37, 0, 0, 421, 420, 1, 0, 0, 0, 421, 422,
		1, 0, 0, 0, 422, 423, 1, 0, 0, 0, 423, 424, 3, 104, 52, 0, 424, 435, 5,
		44, 0, 0, 425, 430, 5, 47, 0, 0, 426, 429, 3, 8, 4, 0, 427, 429, 3, 80,
		40, 0, 428, 426, 1, 0, 0, 0, 428, 427, 1, 0, 0, 0, 429, 432, 1, 0, 0, 0,
		430, 428, 1, 0, 0, 0, 430, 431, 1, 0, 0, 0, 431, 433, 1, 0, 0, 0, 432,
		430, 1, 0, 0, 0, 433, 436, 5, 48, 0, 0, 434, 436, 5, 41, 0, 0, 435, 425,
		1, 0, 0, 0, 435, 434, 1, 0, 0, 0, 436, 73, 1, 0, 0, 0, 437, 438, 5, 37,
		0, 0, 438, 439, 3, 102, 51, 0, 439, 440, 5, 43, 0, 0, 440, 441, 3, 104,
		52, 0, 441, 442, 5, 52, 0, 0, 442, 443, 3, 104, 52, 0, 443, 454, 5, 44,
		0, 0, 444, 449, 5, 47, 0, 0, 445, 448, 3, 8, 4, 0, 446, 448, 3, 80, 40,
		0, 447, 445, 1, 0, 0, 0, 447, 446, 1, 0, 0, 0, 448, 451, 1, 0, 0, 0, 449,
		447, 1, 0, 0, 0, 449, 450, 1, 0, 0, 0, 450, 452, 1, 0, 0, 0, 451, 449,
		1, 0, 0, 0, 452, 455, 5, 48, 0, 0, 453, 455, 5, 41, 0, 0, 454, 444, 1,
		0, 0, 0, 454, 453, 1, 0, 0, 0, 455, 75, 1, 0, 0, 0, 456, 469, 3, 84, 42,
		0, 457, 459, 7, 4, 0, 0, 458, 457, 1, 0, 0, 0, 458, 459, 1, 0, 0, 0, 459,
		460, 1, 0, 0, 0, 460, 469, 3, 108, 54, 0, 461, 463, 7, 4, 0, 0, 462, 461,
		1, 0, 0, 0, 462, 463, 1, 0, 0, 0, 463, 464, 1, 0, 0, 0, 464, 469, 3, 114,
		57, 0, 465, 469, 3, 110, 55, 0, 466, 469, 3, 112, 56, 0, 467, 469, 3, 78,
		39, 0, 468, 456, 1, 0, 0, 0, 468, 458, 1, 0, 0, 0, 468, 462, 1, 0, 0, 0,
		468, 465, 1, 0, 0, 0, 468, 466, 1, 0, 0, 0, 468, 467, 1, 0, 0, 0, 469,
		77, 1, 0, 0, 0, 470, 477, 5, 47, 0, 0, 471, 472, 3, 82, 41, 0, 472, 473,
		5, 53, 0, 0, 473, 474, 3, 76, 38, 0, 474, 476, 1, 0, 0, 0, 475, 471, 1,
		0, 0, 0, 476, 479, 1, 0, 0, 0, 477, 475, 1, 0, 0, 0, 477, 478, 1, 0, 0,
		0, 478, 480, 1, 0, 0, 0, 479, 477, 1, 0, 0, 0, 480, 481, 5, 48, 0, 0, 481,
		79, 1, 0, 0, 0, 482, 483, 5, 41, 0, 0, 483, 81, 1, 0, 0, 0, 484, 487, 5,
		60, 0, 0, 485, 487, 3, 116, 58, 0, 486, 484, 1, 0, 0, 0, 486, 485, 1, 0,
		0, 0, 487, 83, 1, 0, 0, 0, 488, 493, 3, 82, 41, 0, 489, 490, 5, 51, 0,
		0, 490, 492, 3, 82, 41, 0, 491, 489, 1, 0, 0, 0, 492, 495, 1, 0, 0, 0,
		493, 491, 1, 0, 0, 0, 493, 494, 1, 0, 0, 0, 494, 85, 1, 0, 0, 0, 495, 493,
		1, 0, 0, 0, 496, 497, 3, 82, 41, 0, 497, 87, 1, 0, 0, 0, 498, 499, 3, 82,
		41, 0, 499, 89, 1, 0, 0, 0, 500, 501, 3, 82, 41, 0, 501, 91, 1, 0, 0, 0,
		502, 503, 3, 82, 41, 0, 503, 93, 1, 0, 0, 0, 504, 505, 3, 82, 41, 0, 505,
		95, 1, 0, 0, 0, 506, 507, 3, 82, 41, 0, 507, 97, 1, 0, 0, 0, 508, 509,
		3, 82, 41, 0, 509, 99, 1, 0, 0, 0, 510, 511, 3, 82, 41, 0, 511, 101, 1,
		0, 0, 0, 512, 513, 3, 82, 41, 0, 513, 103, 1, 0, 0, 0, 514, 516, 5, 51,
		0, 0, 515, 514, 1, 0, 0, 0, 515, 516, 1, 0, 0, 0, 516, 522, 1, 0, 0, 0,
		517, 518, 3, 82, 41, 0, 518, 519, 5, 51, 0, 0, 519, 521, 1, 0, 0, 0, 520,
		517, 1, 0, 0, 0, 521, 524, 1, 0, 0, 0, 522, 520, 1, 0, 0, 0, 522, 523,
		1, 0, 0, 0, 523, 525, 1, 0, 0, 0, 524, 522, 1, 0, 0, 0, 525, 526, 3, 86,
		43, 0, 526, 105, 1, 0, 0, 0, 527, 529, 5, 51, 0, 0, 528, 527, 1, 0, 0,
		0, 528, 529, 1, 0, 0, 0, 529, 535, 1, 0, 0, 0, 530, 531, 3, 82, 41, 0,
		531, 532, 5, 51, 0, 0, 532, 534, 1, 0, 0, 0, 533, 530, 1, 0, 0, 0, 534,
		537, 1, 0, 0, 0, 535, 533, 1, 0, 0, 0, 535, 536, 1, 0, 0, 0, 536, 538,
		1, 0, 0, 0, 537, 535, 1, 0, 0, 0, 538, 539, 3, 88, 44, 0, 539, 107, 1,
		0, 0, 0, 540, 541, 5, 59, 0, 0, 541, 109, 1, 0, 0, 0, 542, 543, 7, 5, 0,
		0, 543, 111, 1, 0, 0, 0, 544, 545, 5, 57, 0, 0, 545, 113, 1, 0, 0, 0, 546,
		547, 5, 58, 0, 0, 547, 115, 1, 0, 0, 0, 548, 549, 7, 6, 0, 0, 549, 117,
		1, 0, 0, 0, 48, 124, 126, 138, 159, 163, 165, 178, 187, 209, 211, 224,
		241, 264, 273, 282, 289, 291, 298, 305, 315, 323, 328, 332, 342, 359, 375,
		383, 391, 399, 408, 414, 421, 428, 430, 435, 447, 449, 454, 458, 462, 468,
		477, 486, 493, 515, 522, 528, 535,
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

// Protobuf2ParserInit initializes any static state used to implement Protobuf2Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewProtobuf2Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Protobuf2ParserInit() {
	staticData := &Protobuf2ParserStaticData
	staticData.once.Do(protobuf2ParserInit)
}

// NewProtobuf2Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewProtobuf2Parser(input antlr.TokenStream) *Protobuf2Parser {
	Protobuf2ParserInit()
	this := new(Protobuf2Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &Protobuf2ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Protobuf2.g4"

	return this
}

// Protobuf2Parser tokens.
const (
	Protobuf2ParserEOF               = antlr.TokenEOF
	Protobuf2ParserSYNTAX            = 1
	Protobuf2ParserIMPORT            = 2
	Protobuf2ParserWEAK              = 3
	Protobuf2ParserPUBLIC            = 4
	Protobuf2ParserPACKAGE           = 5
	Protobuf2ParserOPTION            = 6
	Protobuf2ParserREPEATED          = 7
	Protobuf2ParserOPTIONAL          = 8
	Protobuf2ParserREQUIRED          = 9
	Protobuf2ParserGROUP             = 10
	Protobuf2ParserONEOF             = 11
	Protobuf2ParserMAP               = 12
	Protobuf2ParserINT32             = 13
	Protobuf2ParserINT64             = 14
	Protobuf2ParserUINT32            = 15
	Protobuf2ParserUINT64            = 16
	Protobuf2ParserSINT32            = 17
	Protobuf2ParserSINT64            = 18
	Protobuf2ParserFIXED32           = 19
	Protobuf2ParserFIXED64           = 20
	Protobuf2ParserSFIXED32          = 21
	Protobuf2ParserSFIXED64          = 22
	Protobuf2ParserBOOL              = 23
	Protobuf2ParserSTRING            = 24
	Protobuf2ParserDOUBLE            = 25
	Protobuf2ParserFLOAT             = 26
	Protobuf2ParserBYTES             = 27
	Protobuf2ParserRESERVED          = 28
	Protobuf2ParserEXTENSIONS        = 29
	Protobuf2ParserTO                = 30
	Protobuf2ParserMAX               = 31
	Protobuf2ParserENUM              = 32
	Protobuf2ParserEXTEND            = 33
	Protobuf2ParserMESSAGE           = 34
	Protobuf2ParserSERVICE           = 35
	Protobuf2ParserRPC               = 36
	Protobuf2ParserSTREAM            = 37
	Protobuf2ParserRETURNS           = 38
	Protobuf2ParserPROTO2_LIT_SINGLE = 39
	Protobuf2ParserPROTO2_LIT_DOBULE = 40
	Protobuf2ParserSEMI              = 41
	Protobuf2ParserEQ                = 42
	Protobuf2ParserLP                = 43
	Protobuf2ParserRP                = 44
	Protobuf2ParserLB                = 45
	Protobuf2ParserRB                = 46
	Protobuf2ParserLC                = 47
	Protobuf2ParserRC                = 48
	Protobuf2ParserLT                = 49
	Protobuf2ParserGT                = 50
	Protobuf2ParserDOT               = 51
	Protobuf2ParserCOMMA             = 52
	Protobuf2ParserCOLON             = 53
	Protobuf2ParserPLUS              = 54
	Protobuf2ParserMINUS             = 55
	Protobuf2ParserSTR_LIT           = 56
	Protobuf2ParserBOOL_LIT          = 57
	Protobuf2ParserFLOAT_LIT         = 58
	Protobuf2ParserINT_LIT           = 59
	Protobuf2ParserIDENTIFIER        = 60
	Protobuf2ParserWS                = 61
	Protobuf2ParserLINE_COMMENT      = 62
	Protobuf2ParserCOMMENT           = 63
)

// Protobuf2Parser rules.
const (
	Protobuf2ParserRULE_proto              = 0
	Protobuf2ParserRULE_syntax             = 1
	Protobuf2ParserRULE_importStatement    = 2
	Protobuf2ParserRULE_packageStatement   = 3
	Protobuf2ParserRULE_optionStatement    = 4
	Protobuf2ParserRULE_optionName         = 5
	Protobuf2ParserRULE_fieldLabel         = 6
	Protobuf2ParserRULE_field              = 7
	Protobuf2ParserRULE_fieldOptions       = 8
	Protobuf2ParserRULE_fieldOption        = 9
	Protobuf2ParserRULE_fieldNumber        = 10
	Protobuf2ParserRULE_group              = 11
	Protobuf2ParserRULE_oneof              = 12
	Protobuf2ParserRULE_oneofField         = 13
	Protobuf2ParserRULE_mapField           = 14
	Protobuf2ParserRULE_keyType            = 15
	Protobuf2ParserRULE_type_              = 16
	Protobuf2ParserRULE_extensions         = 17
	Protobuf2ParserRULE_reserved           = 18
	Protobuf2ParserRULE_ranges             = 19
	Protobuf2ParserRULE_range_             = 20
	Protobuf2ParserRULE_reservedFieldNames = 21
	Protobuf2ParserRULE_topLevelDef        = 22
	Protobuf2ParserRULE_enumDef            = 23
	Protobuf2ParserRULE_enumBody           = 24
	Protobuf2ParserRULE_enumElement        = 25
	Protobuf2ParserRULE_enumField          = 26
	Protobuf2ParserRULE_enumValueOptions   = 27
	Protobuf2ParserRULE_enumValueOption    = 28
	Protobuf2ParserRULE_messageDef         = 29
	Protobuf2ParserRULE_messageBody        = 30
	Protobuf2ParserRULE_messageElement     = 31
	Protobuf2ParserRULE_extendDef          = 32
	Protobuf2ParserRULE_extendElement      = 33
	Protobuf2ParserRULE_serviceDef         = 34
	Protobuf2ParserRULE_serviceElement     = 35
	Protobuf2ParserRULE_rpc                = 36
	Protobuf2ParserRULE_stream             = 37
	Protobuf2ParserRULE_constant           = 38
	Protobuf2ParserRULE_blockLit           = 39
	Protobuf2ParserRULE_emptyStatement_    = 40
	Protobuf2ParserRULE_ident              = 41
	Protobuf2ParserRULE_fullIdent          = 42
	Protobuf2ParserRULE_messageName        = 43
	Protobuf2ParserRULE_enumName           = 44
	Protobuf2ParserRULE_fieldName          = 45
	Protobuf2ParserRULE_groupName          = 46
	Protobuf2ParserRULE_oneofName          = 47
	Protobuf2ParserRULE_mapName            = 48
	Protobuf2ParserRULE_serviceName        = 49
	Protobuf2ParserRULE_rpcName            = 50
	Protobuf2ParserRULE_streamName         = 51
	Protobuf2ParserRULE_messageType        = 52
	Protobuf2ParserRULE_enumType           = 53
	Protobuf2ParserRULE_intLit             = 54
	Protobuf2ParserRULE_strLit             = 55
	Protobuf2ParserRULE_boolLit            = 56
	Protobuf2ParserRULE_floatLit           = 57
	Protobuf2ParserRULE_keywords           = 58
)

// IProtoContext is an interface to support dynamic dispatch.
type IProtoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Syntax() ISyntaxContext
	EOF() antlr.TerminalNode
	AllImportStatement() []IImportStatementContext
	ImportStatement(i int) IImportStatementContext
	AllPackageStatement() []IPackageStatementContext
	PackageStatement(i int) IPackageStatementContext
	AllOptionStatement() []IOptionStatementContext
	OptionStatement(i int) IOptionStatementContext
	AllTopLevelDef() []ITopLevelDefContext
	TopLevelDef(i int) ITopLevelDefContext
	AllEmptyStatement_() []IEmptyStatement_Context
	EmptyStatement_(i int) IEmptyStatement_Context

	// IsProtoContext differentiates from other interfaces.
	IsProtoContext()
}

type ProtoContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProtoContext() *ProtoContext {
	var p = new(ProtoContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_proto
	return p
}

func InitEmptyProtoContext(p *ProtoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_proto
}

func (*ProtoContext) IsProtoContext() {}

func NewProtoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProtoContext {
	var p = new(ProtoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_proto

	return p
}

func (s *ProtoContext) GetParser() antlr.Parser { return s.parser }

func (s *ProtoContext) Syntax() ISyntaxContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISyntaxContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISyntaxContext)
}

func (s *ProtoContext) EOF() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserEOF, 0)
}

func (s *ProtoContext) AllImportStatement() []IImportStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IImportStatementContext); ok {
			len++
		}
	}

	tst := make([]IImportStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IImportStatementContext); ok {
			tst[i] = t.(IImportStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProtoContext) ImportStatement(i int) IImportStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IImportStatementContext); ok {
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

	return t.(IImportStatementContext)
}

func (s *ProtoContext) AllPackageStatement() []IPackageStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPackageStatementContext); ok {
			len++
		}
	}

	tst := make([]IPackageStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPackageStatementContext); ok {
			tst[i] = t.(IPackageStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProtoContext) PackageStatement(i int) IPackageStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPackageStatementContext); ok {
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

	return t.(IPackageStatementContext)
}

func (s *ProtoContext) AllOptionStatement() []IOptionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOptionStatementContext); ok {
			len++
		}
	}

	tst := make([]IOptionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOptionStatementContext); ok {
			tst[i] = t.(IOptionStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProtoContext) OptionStatement(i int) IOptionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionStatementContext); ok {
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

	return t.(IOptionStatementContext)
}

func (s *ProtoContext) AllTopLevelDef() []ITopLevelDefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITopLevelDefContext); ok {
			len++
		}
	}

	tst := make([]ITopLevelDefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITopLevelDefContext); ok {
			tst[i] = t.(ITopLevelDefContext)
			i++
		}
	}

	return tst
}

func (s *ProtoContext) TopLevelDef(i int) ITopLevelDefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITopLevelDefContext); ok {
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

	return t.(ITopLevelDefContext)
}

func (s *ProtoContext) AllEmptyStatement_() []IEmptyStatement_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			len++
		}
	}

	tst := make([]IEmptyStatement_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEmptyStatement_Context); ok {
			tst[i] = t.(IEmptyStatement_Context)
			i++
		}
	}

	return tst
}

func (s *ProtoContext) EmptyStatement_(i int) IEmptyStatement_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
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

	return t.(IEmptyStatement_Context)
}

func (s *ProtoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProtoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProtoContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitProto(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) Proto() (localctx IProtoContext) {
	localctx = NewProtoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Protobuf2ParserRULE_proto)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(118)
		p.Syntax()
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2263447765092) != 0 {
		p.SetState(124)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case Protobuf2ParserIMPORT:
			{
				p.SetState(119)
				p.ImportStatement()
			}

		case Protobuf2ParserPACKAGE:
			{
				p.SetState(120)
				p.PackageStatement()
			}

		case Protobuf2ParserOPTION:
			{
				p.SetState(121)
				p.OptionStatement()
			}

		case Protobuf2ParserENUM, Protobuf2ParserEXTEND, Protobuf2ParserMESSAGE, Protobuf2ParserSERVICE:
			{
				p.SetState(122)
				p.TopLevelDef()
			}

		case Protobuf2ParserSEMI:
			{
				p.SetState(123)
				p.EmptyStatement_()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(129)
		p.Match(Protobuf2ParserEOF)
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

// ISyntaxContext is an interface to support dynamic dispatch.
type ISyntaxContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SYNTAX() antlr.TerminalNode
	EQ() antlr.TerminalNode
	SEMI() antlr.TerminalNode
	PROTO2_LIT_SINGLE() antlr.TerminalNode
	PROTO2_LIT_DOBULE() antlr.TerminalNode

	// IsSyntaxContext differentiates from other interfaces.
	IsSyntaxContext()
}

type SyntaxContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySyntaxContext() *SyntaxContext {
	var p = new(SyntaxContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_syntax
	return p
}

func InitEmptySyntaxContext(p *SyntaxContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_syntax
}

func (*SyntaxContext) IsSyntaxContext() {}

func NewSyntaxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SyntaxContext {
	var p = new(SyntaxContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_syntax

	return p
}

func (s *SyntaxContext) GetParser() antlr.Parser { return s.parser }

func (s *SyntaxContext) SYNTAX() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSYNTAX, 0)
}

func (s *SyntaxContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserEQ, 0)
}

func (s *SyntaxContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSEMI, 0)
}

func (s *SyntaxContext) PROTO2_LIT_SINGLE() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserPROTO2_LIT_SINGLE, 0)
}

func (s *SyntaxContext) PROTO2_LIT_DOBULE() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserPROTO2_LIT_DOBULE, 0)
}

func (s *SyntaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SyntaxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SyntaxContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitSyntax(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) Syntax() (localctx ISyntaxContext) {
	localctx = NewSyntaxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Protobuf2ParserRULE_syntax)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(Protobuf2ParserSYNTAX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.Match(Protobuf2ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(133)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Protobuf2ParserPROTO2_LIT_SINGLE || _la == Protobuf2ParserPROTO2_LIT_DOBULE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(134)
		p.Match(Protobuf2ParserSEMI)
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

// IImportStatementContext is an interface to support dynamic dispatch.
type IImportStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IMPORT() antlr.TerminalNode
	StrLit() IStrLitContext
	SEMI() antlr.TerminalNode
	WEAK() antlr.TerminalNode
	PUBLIC() antlr.TerminalNode

	// IsImportStatementContext differentiates from other interfaces.
	IsImportStatementContext()
}

type ImportStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImportStatementContext() *ImportStatementContext {
	var p = new(ImportStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_importStatement
	return p
}

func InitEmptyImportStatementContext(p *ImportStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_importStatement
}

func (*ImportStatementContext) IsImportStatementContext() {}

func NewImportStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStatementContext {
	var p = new(ImportStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_importStatement

	return p
}

func (s *ImportStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStatementContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserIMPORT, 0)
}

func (s *ImportStatementContext) StrLit() IStrLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrLitContext)
}

func (s *ImportStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSEMI, 0)
}

func (s *ImportStatementContext) WEAK() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserWEAK, 0)
}

func (s *ImportStatementContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserPUBLIC, 0)
}

func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitImportStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) ImportStatement() (localctx IImportStatementContext) {
	localctx = NewImportStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Protobuf2ParserRULE_importStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		p.Match(Protobuf2ParserIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf2ParserWEAK || _la == Protobuf2ParserPUBLIC {
		{
			p.SetState(137)
			_la = p.GetTokenStream().LA(1)

			if !(_la == Protobuf2ParserWEAK || _la == Protobuf2ParserPUBLIC) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(140)
		p.StrLit()
	}
	{
		p.SetState(141)
		p.Match(Protobuf2ParserSEMI)
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

// IPackageStatementContext is an interface to support dynamic dispatch.
type IPackageStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PACKAGE() antlr.TerminalNode
	FullIdent() IFullIdentContext
	SEMI() antlr.TerminalNode

	// IsPackageStatementContext differentiates from other interfaces.
	IsPackageStatementContext()
}

type PackageStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPackageStatementContext() *PackageStatementContext {
	var p = new(PackageStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_packageStatement
	return p
}

func InitEmptyPackageStatementContext(p *PackageStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_packageStatement
}

func (*PackageStatementContext) IsPackageStatementContext() {}

func NewPackageStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageStatementContext {
	var p = new(PackageStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_packageStatement

	return p
}

func (s *PackageStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageStatementContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserPACKAGE, 0)
}

func (s *PackageStatementContext) FullIdent() IFullIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFullIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *PackageStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSEMI, 0)
}

func (s *PackageStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitPackageStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) PackageStatement() (localctx IPackageStatementContext) {
	localctx = NewPackageStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Protobuf2ParserRULE_packageStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(143)
		p.Match(Protobuf2ParserPACKAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(144)
		p.FullIdent()
	}
	{
		p.SetState(145)
		p.Match(Protobuf2ParserSEMI)
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

// IOptionStatementContext is an interface to support dynamic dispatch.
type IOptionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPTION() antlr.TerminalNode
	OptionName() IOptionNameContext
	EQ() antlr.TerminalNode
	Constant() IConstantContext
	SEMI() antlr.TerminalNode

	// IsOptionStatementContext differentiates from other interfaces.
	IsOptionStatementContext()
}

type OptionStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionStatementContext() *OptionStatementContext {
	var p = new(OptionStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_optionStatement
	return p
}

func InitEmptyOptionStatementContext(p *OptionStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_optionStatement
}

func (*OptionStatementContext) IsOptionStatementContext() {}

func NewOptionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionStatementContext {
	var p = new(OptionStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_optionStatement

	return p
}

func (s *OptionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionStatementContext) OPTION() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserOPTION, 0)
}

func (s *OptionStatementContext) OptionName() IOptionNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *OptionStatementContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserEQ, 0)
}

func (s *OptionStatementContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *OptionStatementContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSEMI, 0)
}

func (s *OptionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitOptionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) OptionStatement() (localctx IOptionStatementContext) {
	localctx = NewOptionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Protobuf2ParserRULE_optionStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		p.Match(Protobuf2ParserOPTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.OptionName()
	}
	{
		p.SetState(149)
		p.Match(Protobuf2ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(150)
		p.Constant()
	}
	{
		p.SetState(151)
		p.Match(Protobuf2ParserSEMI)
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

// IOptionNameContext is an interface to support dynamic dispatch.
type IOptionNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFullIdent() []IFullIdentContext
	FullIdent(i int) IFullIdentContext
	Ident() IIdentContext
	LP() antlr.TerminalNode
	RP() antlr.TerminalNode
	DOT() antlr.TerminalNode

	// IsOptionNameContext differentiates from other interfaces.
	IsOptionNameContext()
}

type OptionNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionNameContext() *OptionNameContext {
	var p = new(OptionNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_optionName
	return p
}

func InitEmptyOptionNameContext(p *OptionNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_optionName
}

func (*OptionNameContext) IsOptionNameContext() {}

func NewOptionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionNameContext {
	var p = new(OptionNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_optionName

	return p
}

func (s *OptionNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionNameContext) AllFullIdent() []IFullIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFullIdentContext); ok {
			len++
		}
	}

	tst := make([]IFullIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFullIdentContext); ok {
			tst[i] = t.(IFullIdentContext)
			i++
		}
	}

	return tst
}

func (s *OptionNameContext) FullIdent(i int) IFullIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFullIdentContext); ok {
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

	return t.(IFullIdentContext)
}

func (s *OptionNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *OptionNameContext) LP() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserLP, 0)
}

func (s *OptionNameContext) RP() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserRP, 0)
}

func (s *OptionNameContext) DOT() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserDOT, 0)
}

func (s *OptionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitOptionName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) OptionName() (localctx IOptionNameContext) {
	localctx = NewOptionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, Protobuf2ParserRULE_optionName)
	var _la int

	p.SetState(165)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.FullIdent()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case Protobuf2ParserSYNTAX, Protobuf2ParserIMPORT, Protobuf2ParserWEAK, Protobuf2ParserPUBLIC, Protobuf2ParserPACKAGE, Protobuf2ParserOPTION, Protobuf2ParserREPEATED, Protobuf2ParserOPTIONAL, Protobuf2ParserREQUIRED, Protobuf2ParserGROUP, Protobuf2ParserONEOF, Protobuf2ParserMAP, Protobuf2ParserINT32, Protobuf2ParserINT64, Protobuf2ParserUINT32, Protobuf2ParserUINT64, Protobuf2ParserSINT32, Protobuf2ParserSINT64, Protobuf2ParserFIXED32, Protobuf2ParserFIXED64, Protobuf2ParserSFIXED32, Protobuf2ParserSFIXED64, Protobuf2ParserBOOL, Protobuf2ParserSTRING, Protobuf2ParserDOUBLE, Protobuf2ParserFLOAT, Protobuf2ParserBYTES, Protobuf2ParserRESERVED, Protobuf2ParserEXTENSIONS, Protobuf2ParserTO, Protobuf2ParserMAX, Protobuf2ParserENUM, Protobuf2ParserEXTEND, Protobuf2ParserMESSAGE, Protobuf2ParserSERVICE, Protobuf2ParserRPC, Protobuf2ParserSTREAM, Protobuf2ParserRETURNS, Protobuf2ParserBOOL_LIT, Protobuf2ParserIDENTIFIER:
			{
				p.SetState(154)
				p.Ident()
			}

		case Protobuf2ParserLP:
			{
				p.SetState(155)
				p.Match(Protobuf2ParserLP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(156)
				p.FullIdent()
			}
			{
				p.SetState(157)
				p.Match(Protobuf2ParserRP)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Protobuf2ParserDOT {
			{
				p.SetState(161)
				p.Match(Protobuf2ParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(162)
				p.FullIdent()
			}

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

// IFieldLabelContext is an interface to support dynamic dispatch.
type IFieldLabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REQUIRED() antlr.TerminalNode
	OPTIONAL() antlr.TerminalNode
	REPEATED() antlr.TerminalNode

	// IsFieldLabelContext differentiates from other interfaces.
	IsFieldLabelContext()
}

type FieldLabelContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldLabelContext() *FieldLabelContext {
	var p = new(FieldLabelContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_fieldLabel
	return p
}

func InitEmptyFieldLabelContext(p *FieldLabelContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_fieldLabel
}

func (*FieldLabelContext) IsFieldLabelContext() {}

func NewFieldLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldLabelContext {
	var p = new(FieldLabelContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_fieldLabel

	return p
}

func (s *FieldLabelContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldLabelContext) REQUIRED() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserREQUIRED, 0)
}

func (s *FieldLabelContext) OPTIONAL() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserOPTIONAL, 0)
}

func (s *FieldLabelContext) REPEATED() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserREPEATED, 0)
}

func (s *FieldLabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldLabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldLabelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitFieldLabel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) FieldLabel() (localctx IFieldLabelContext) {
	localctx = NewFieldLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, Protobuf2ParserRULE_fieldLabel)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&896) != 0) {
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

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FieldLabel() IFieldLabelContext
	Type_() IType_Context
	FieldName() IFieldNameContext
	EQ() antlr.TerminalNode
	FieldNumber() IFieldNumberContext
	SEMI() antlr.TerminalNode
	LB() antlr.TerminalNode
	FieldOptions() IFieldOptionsContext
	RB() antlr.TerminalNode

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_field
	return p
}

func InitEmptyFieldContext(p *FieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_field
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) FieldLabel() IFieldLabelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldLabelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldLabelContext)
}

func (s *FieldContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *FieldContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *FieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserEQ, 0)
}

func (s *FieldContext) FieldNumber() IFieldNumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNumberContext)
}

func (s *FieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSEMI, 0)
}

func (s *FieldContext) LB() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserLB, 0)
}

func (s *FieldContext) FieldOptions() IFieldOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *FieldContext) RB() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserRB, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, Protobuf2ParserRULE_field)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.FieldLabel()
	}
	{
		p.SetState(170)
		p.Type_()
	}
	{
		p.SetState(171)
		p.FieldName()
	}
	{
		p.SetState(172)
		p.Match(Protobuf2ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.FieldNumber()
	}
	p.SetState(178)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf2ParserLB {
		{
			p.SetState(174)
			p.Match(Protobuf2ParserLB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(175)
			p.FieldOptions()
		}
		{
			p.SetState(176)
			p.Match(Protobuf2ParserRB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(180)
		p.Match(Protobuf2ParserSEMI)
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

// IFieldOptionsContext is an interface to support dynamic dispatch.
type IFieldOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllFieldOption() []IFieldOptionContext
	FieldOption(i int) IFieldOptionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsFieldOptionsContext differentiates from other interfaces.
	IsFieldOptionsContext()
}

type FieldOptionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldOptionsContext() *FieldOptionsContext {
	var p = new(FieldOptionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_fieldOptions
	return p
}

func InitEmptyFieldOptionsContext(p *FieldOptionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_fieldOptions
}

func (*FieldOptionsContext) IsFieldOptionsContext() {}

func NewFieldOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptionsContext {
	var p = new(FieldOptionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_fieldOptions

	return p
}

func (s *FieldOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldOptionsContext) AllFieldOption() []IFieldOptionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFieldOptionContext); ok {
			len++
		}
	}

	tst := make([]IFieldOptionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFieldOptionContext); ok {
			tst[i] = t.(IFieldOptionContext)
			i++
		}
	}

	return tst
}

func (s *FieldOptionsContext) FieldOption(i int) IFieldOptionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldOptionContext); ok {
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

	return t.(IFieldOptionContext)
}

func (s *FieldOptionsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(Protobuf2ParserCOMMA)
}

func (s *FieldOptionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserCOMMA, i)
}

func (s *FieldOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldOptionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitFieldOptions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) FieldOptions() (localctx IFieldOptionsContext) {
	localctx = NewFieldOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, Protobuf2ParserRULE_fieldOptions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(182)
		p.FieldOption()
	}
	p.SetState(187)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf2ParserCOMMA {
		{
			p.SetState(183)
			p.Match(Protobuf2ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(184)
			p.FieldOption()
		}

		p.SetState(189)
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

// IFieldOptionContext is an interface to support dynamic dispatch.
type IFieldOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OptionName() IOptionNameContext
	EQ() antlr.TerminalNode
	Constant() IConstantContext

	// IsFieldOptionContext differentiates from other interfaces.
	IsFieldOptionContext()
}

type FieldOptionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldOptionContext() *FieldOptionContext {
	var p = new(FieldOptionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_fieldOption
	return p
}

func InitEmptyFieldOptionContext(p *FieldOptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_fieldOption
}

func (*FieldOptionContext) IsFieldOptionContext() {}

func NewFieldOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptionContext {
	var p = new(FieldOptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_fieldOption

	return p
}

func (s *FieldOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldOptionContext) OptionName() IOptionNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *FieldOptionContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserEQ, 0)
}

func (s *FieldOptionContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *FieldOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldOptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitFieldOption(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) FieldOption() (localctx IFieldOptionContext) {
	localctx = NewFieldOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, Protobuf2ParserRULE_fieldOption)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)
		p.OptionName()
	}
	{
		p.SetState(191)
		p.Match(Protobuf2ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(192)
		p.Constant()
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

// IFieldNumberContext is an interface to support dynamic dispatch.
type IFieldNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IntLit() IIntLitContext

	// IsFieldNumberContext differentiates from other interfaces.
	IsFieldNumberContext()
}

type FieldNumberContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNumberContext() *FieldNumberContext {
	var p = new(FieldNumberContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_fieldNumber
	return p
}

func InitEmptyFieldNumberContext(p *FieldNumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_fieldNumber
}

func (*FieldNumberContext) IsFieldNumberContext() {}

func NewFieldNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNumberContext {
	var p = new(FieldNumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_fieldNumber

	return p
}

func (s *FieldNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNumberContext) IntLit() IIntLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntLitContext)
}

func (s *FieldNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitFieldNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) FieldNumber() (localctx IFieldNumberContext) {
	localctx = NewFieldNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, Protobuf2ParserRULE_fieldNumber)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(194)
		p.IntLit()
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

// IGroupContext is an interface to support dynamic dispatch.
type IGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FieldLabel() IFieldLabelContext
	GROUP() antlr.TerminalNode
	GroupName() IGroupNameContext
	EQ() antlr.TerminalNode
	FieldNumber() IFieldNumberContext
	MessageBody() IMessageBodyContext

	// IsGroupContext differentiates from other interfaces.
	IsGroupContext()
}

type GroupContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupContext() *GroupContext {
	var p = new(GroupContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_group
	return p
}

func InitEmptyGroupContext(p *GroupContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_group
}

func (*GroupContext) IsGroupContext() {}

func NewGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupContext {
	var p = new(GroupContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_group

	return p
}

func (s *GroupContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupContext) FieldLabel() IFieldLabelContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldLabelContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldLabelContext)
}

func (s *GroupContext) GROUP() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserGROUP, 0)
}

func (s *GroupContext) GroupName() IGroupNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupNameContext)
}

func (s *GroupContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserEQ, 0)
}

func (s *GroupContext) FieldNumber() IFieldNumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNumberContext)
}

func (s *GroupContext) MessageBody() IMessageBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageBodyContext)
}

func (s *GroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) Group() (localctx IGroupContext) {
	localctx = NewGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, Protobuf2ParserRULE_group)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.FieldLabel()
	}
	{
		p.SetState(197)
		p.Match(Protobuf2ParserGROUP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(198)
		p.GroupName()
	}
	{
		p.SetState(199)
		p.Match(Protobuf2ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(200)
		p.FieldNumber()
	}
	{
		p.SetState(201)
		p.MessageBody()
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

// IOneofContext is an interface to support dynamic dispatch.
type IOneofContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ONEOF() antlr.TerminalNode
	OneofName() IOneofNameContext
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	AllOptionStatement() []IOptionStatementContext
	OptionStatement(i int) IOptionStatementContext
	AllOneofField() []IOneofFieldContext
	OneofField(i int) IOneofFieldContext
	AllEmptyStatement_() []IEmptyStatement_Context
	EmptyStatement_(i int) IEmptyStatement_Context

	// IsOneofContext differentiates from other interfaces.
	IsOneofContext()
}

type OneofContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneofContext() *OneofContext {
	var p = new(OneofContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_oneof
	return p
}

func InitEmptyOneofContext(p *OneofContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_oneof
}

func (*OneofContext) IsOneofContext() {}

func NewOneofContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofContext {
	var p = new(OneofContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_oneof

	return p
}

func (s *OneofContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofContext) ONEOF() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserONEOF, 0)
}

func (s *OneofContext) OneofName() IOneofNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOneofNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOneofNameContext)
}

func (s *OneofContext) LC() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserLC, 0)
}

func (s *OneofContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserRC, 0)
}

func (s *OneofContext) AllOptionStatement() []IOptionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOptionStatementContext); ok {
			len++
		}
	}

	tst := make([]IOptionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOptionStatementContext); ok {
			tst[i] = t.(IOptionStatementContext)
			i++
		}
	}

	return tst
}

func (s *OneofContext) OptionStatement(i int) IOptionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionStatementContext); ok {
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

	return t.(IOptionStatementContext)
}

func (s *OneofContext) AllOneofField() []IOneofFieldContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOneofFieldContext); ok {
			len++
		}
	}

	tst := make([]IOneofFieldContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOneofFieldContext); ok {
			tst[i] = t.(IOneofFieldContext)
			i++
		}
	}

	return tst
}

func (s *OneofContext) OneofField(i int) IOneofFieldContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOneofFieldContext); ok {
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

	return t.(IOneofFieldContext)
}

func (s *OneofContext) AllEmptyStatement_() []IEmptyStatement_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			len++
		}
	}

	tst := make([]IEmptyStatement_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEmptyStatement_Context); ok {
			tst[i] = t.(IEmptyStatement_Context)
			i++
		}
	}

	return tst
}

func (s *OneofContext) EmptyStatement_(i int) IEmptyStatement_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
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

	return t.(IEmptyStatement_Context)
}

func (s *OneofContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneofContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitOneof(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) Oneof() (localctx IOneofContext) {
	localctx = NewOneofContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, Protobuf2ParserRULE_oneof)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(203)
		p.Match(Protobuf2ParserONEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(204)
		p.OneofName()
	}
	{
		p.SetState(205)
		p.Match(Protobuf2ParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(211)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1299291241275457534) != 0 {
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(206)
				p.OptionStatement()
			}

		case 2:
			{
				p.SetState(207)
				p.OneofField()
			}

		case 3:
			{
				p.SetState(208)
				p.EmptyStatement_()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(214)
		p.Match(Protobuf2ParserRC)
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

// IOneofFieldContext is an interface to support dynamic dispatch.
type IOneofFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() IType_Context
	FieldName() IFieldNameContext
	EQ() antlr.TerminalNode
	FieldNumber() IFieldNumberContext
	SEMI() antlr.TerminalNode
	LB() antlr.TerminalNode
	FieldOptions() IFieldOptionsContext
	RB() antlr.TerminalNode

	// IsOneofFieldContext differentiates from other interfaces.
	IsOneofFieldContext()
}

type OneofFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneofFieldContext() *OneofFieldContext {
	var p = new(OneofFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_oneofField
	return p
}

func InitEmptyOneofFieldContext(p *OneofFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_oneofField
}

func (*OneofFieldContext) IsOneofFieldContext() {}

func NewOneofFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofFieldContext {
	var p = new(OneofFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_oneofField

	return p
}

func (s *OneofFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofFieldContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *OneofFieldContext) FieldName() IFieldNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNameContext)
}

func (s *OneofFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserEQ, 0)
}

func (s *OneofFieldContext) FieldNumber() IFieldNumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNumberContext)
}

func (s *OneofFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSEMI, 0)
}

func (s *OneofFieldContext) LB() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserLB, 0)
}

func (s *OneofFieldContext) FieldOptions() IFieldOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *OneofFieldContext) RB() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserRB, 0)
}

func (s *OneofFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneofFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitOneofField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) OneofField() (localctx IOneofFieldContext) {
	localctx = NewOneofFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, Protobuf2ParserRULE_oneofField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Type_()
	}
	{
		p.SetState(217)
		p.FieldName()
	}
	{
		p.SetState(218)
		p.Match(Protobuf2ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(219)
		p.FieldNumber()
	}
	p.SetState(224)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf2ParserLB {
		{
			p.SetState(220)
			p.Match(Protobuf2ParserLB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(221)
			p.FieldOptions()
		}
		{
			p.SetState(222)
			p.Match(Protobuf2ParserRB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(226)
		p.Match(Protobuf2ParserSEMI)
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

// IMapFieldContext is an interface to support dynamic dispatch.
type IMapFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MAP() antlr.TerminalNode
	LT() antlr.TerminalNode
	KeyType() IKeyTypeContext
	COMMA() antlr.TerminalNode
	Type_() IType_Context
	GT() antlr.TerminalNode
	MapName() IMapNameContext
	EQ() antlr.TerminalNode
	FieldNumber() IFieldNumberContext
	SEMI() antlr.TerminalNode
	LB() antlr.TerminalNode
	FieldOptions() IFieldOptionsContext
	RB() antlr.TerminalNode

	// IsMapFieldContext differentiates from other interfaces.
	IsMapFieldContext()
}

type MapFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapFieldContext() *MapFieldContext {
	var p = new(MapFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_mapField
	return p
}

func InitEmptyMapFieldContext(p *MapFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_mapField
}

func (*MapFieldContext) IsMapFieldContext() {}

func NewMapFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapFieldContext {
	var p = new(MapFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_mapField

	return p
}

func (s *MapFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *MapFieldContext) MAP() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserMAP, 0)
}

func (s *MapFieldContext) LT() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserLT, 0)
}

func (s *MapFieldContext) KeyType() IKeyTypeContext {
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

func (s *MapFieldContext) COMMA() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserCOMMA, 0)
}

func (s *MapFieldContext) Type_() IType_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IType_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *MapFieldContext) GT() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserGT, 0)
}

func (s *MapFieldContext) MapName() IMapNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapNameContext)
}

func (s *MapFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserEQ, 0)
}

func (s *MapFieldContext) FieldNumber() IFieldNumberContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldNumberContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldNumberContext)
}

func (s *MapFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSEMI, 0)
}

func (s *MapFieldContext) LB() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserLB, 0)
}

func (s *MapFieldContext) FieldOptions() IFieldOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldOptionsContext)
}

func (s *MapFieldContext) RB() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserRB, 0)
}

func (s *MapFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitMapField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) MapField() (localctx IMapFieldContext) {
	localctx = NewMapFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, Protobuf2ParserRULE_mapField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(228)
		p.Match(Protobuf2ParserMAP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(229)
		p.Match(Protobuf2ParserLT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(230)
		p.KeyType()
	}
	{
		p.SetState(231)
		p.Match(Protobuf2ParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(232)
		p.Type_()
	}
	{
		p.SetState(233)
		p.Match(Protobuf2ParserGT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(234)
		p.MapName()
	}
	{
		p.SetState(235)
		p.Match(Protobuf2ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(236)
		p.FieldNumber()
	}
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf2ParserLB {
		{
			p.SetState(237)
			p.Match(Protobuf2ParserLB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(238)
			p.FieldOptions()
		}
		{
			p.SetState(239)
			p.Match(Protobuf2ParserRB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(243)
		p.Match(Protobuf2ParserSEMI)
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

// IKeyTypeContext is an interface to support dynamic dispatch.
type IKeyTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT32() antlr.TerminalNode
	INT64() antlr.TerminalNode
	UINT32() antlr.TerminalNode
	UINT64() antlr.TerminalNode
	SINT32() antlr.TerminalNode
	SINT64() antlr.TerminalNode
	FIXED32() antlr.TerminalNode
	FIXED64() antlr.TerminalNode
	SFIXED32() antlr.TerminalNode
	SFIXED64() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	STRING() antlr.TerminalNode

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
	p.RuleIndex = Protobuf2ParserRULE_keyType
	return p
}

func InitEmptyKeyTypeContext(p *KeyTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_keyType
}

func (*KeyTypeContext) IsKeyTypeContext() {}

func NewKeyTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyTypeContext {
	var p = new(KeyTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_keyType

	return p
}

func (s *KeyTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyTypeContext) INT32() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserINT32, 0)
}

func (s *KeyTypeContext) INT64() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserINT64, 0)
}

func (s *KeyTypeContext) UINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserUINT32, 0)
}

func (s *KeyTypeContext) UINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserUINT64, 0)
}

func (s *KeyTypeContext) SINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSINT32, 0)
}

func (s *KeyTypeContext) SINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSINT64, 0)
}

func (s *KeyTypeContext) FIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserFIXED32, 0)
}

func (s *KeyTypeContext) FIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserFIXED64, 0)
}

func (s *KeyTypeContext) SFIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSFIXED32, 0)
}

func (s *KeyTypeContext) SFIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSFIXED64, 0)
}

func (s *KeyTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserBOOL, 0)
}

func (s *KeyTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSTRING, 0)
}

func (s *KeyTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitKeyType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) KeyType() (localctx IKeyTypeContext) {
	localctx = NewKeyTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, Protobuf2ParserRULE_keyType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(245)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&33546240) != 0) {
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

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DOUBLE() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	INT32() antlr.TerminalNode
	INT64() antlr.TerminalNode
	UINT32() antlr.TerminalNode
	UINT64() antlr.TerminalNode
	SINT32() antlr.TerminalNode
	SINT64() antlr.TerminalNode
	FIXED32() antlr.TerminalNode
	FIXED64() antlr.TerminalNode
	SFIXED32() antlr.TerminalNode
	SFIXED64() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	STRING() antlr.TerminalNode
	BYTES() antlr.TerminalNode
	MessageType() IMessageTypeContext
	EnumType() IEnumTypeContext

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_type_
	return p
}

func InitEmptyType_Context(p *Type_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_type_
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) DOUBLE() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserDOUBLE, 0)
}

func (s *Type_Context) FLOAT() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserFLOAT, 0)
}

func (s *Type_Context) INT32() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserINT32, 0)
}

func (s *Type_Context) INT64() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserINT64, 0)
}

func (s *Type_Context) UINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserUINT32, 0)
}

func (s *Type_Context) UINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserUINT64, 0)
}

func (s *Type_Context) SINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSINT32, 0)
}

func (s *Type_Context) SINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSINT64, 0)
}

func (s *Type_Context) FIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserFIXED32, 0)
}

func (s *Type_Context) FIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserFIXED64, 0)
}

func (s *Type_Context) SFIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSFIXED32, 0)
}

func (s *Type_Context) SFIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSFIXED64, 0)
}

func (s *Type_Context) BOOL() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserBOOL, 0)
}

func (s *Type_Context) STRING() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSTRING, 0)
}

func (s *Type_Context) BYTES() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserBYTES, 0)
}

func (s *Type_Context) MessageType() IMessageTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageTypeContext)
}

func (s *Type_Context) EnumType() IEnumTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumTypeContext)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitType_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) Type_() (localctx IType_Context) {
	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, Protobuf2ParserRULE_type_)
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(247)
			p.Match(Protobuf2ParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(248)
			p.Match(Protobuf2ParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(249)
			p.Match(Protobuf2ParserINT32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(250)
			p.Match(Protobuf2ParserINT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(251)
			p.Match(Protobuf2ParserUINT32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(252)
			p.Match(Protobuf2ParserUINT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(253)
			p.Match(Protobuf2ParserSINT32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(254)
			p.Match(Protobuf2ParserSINT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(255)
			p.Match(Protobuf2ParserFIXED32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(256)
			p.Match(Protobuf2ParserFIXED64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(257)
			p.Match(Protobuf2ParserSFIXED32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(258)
			p.Match(Protobuf2ParserSFIXED64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(259)
			p.Match(Protobuf2ParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(260)
			p.Match(Protobuf2ParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(261)
			p.Match(Protobuf2ParserBYTES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(262)
			p.MessageType()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(263)
			p.EnumType()
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

// IExtensionsContext is an interface to support dynamic dispatch.
type IExtensionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXTENSIONS() antlr.TerminalNode
	Ranges() IRangesContext
	SEMI() antlr.TerminalNode

	// IsExtensionsContext differentiates from other interfaces.
	IsExtensionsContext()
}

type ExtensionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtensionsContext() *ExtensionsContext {
	var p = new(ExtensionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_extensions
	return p
}

func InitEmptyExtensionsContext(p *ExtensionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_extensions
}

func (*ExtensionsContext) IsExtensionsContext() {}

func NewExtensionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtensionsContext {
	var p = new(ExtensionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_extensions

	return p
}

func (s *ExtensionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtensionsContext) EXTENSIONS() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserEXTENSIONS, 0)
}

func (s *ExtensionsContext) Ranges() IRangesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangesContext)
}

func (s *ExtensionsContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSEMI, 0)
}

func (s *ExtensionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtensionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtensionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitExtensions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) Extensions() (localctx IExtensionsContext) {
	localctx = NewExtensionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, Protobuf2ParserRULE_extensions)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.Match(Protobuf2ParserEXTENSIONS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(267)
		p.Ranges()
	}
	{
		p.SetState(268)
		p.Match(Protobuf2ParserSEMI)
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

// IReservedContext is an interface to support dynamic dispatch.
type IReservedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RESERVED() antlr.TerminalNode
	SEMI() antlr.TerminalNode
	Ranges() IRangesContext
	ReservedFieldNames() IReservedFieldNamesContext

	// IsReservedContext differentiates from other interfaces.
	IsReservedContext()
}

type ReservedContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReservedContext() *ReservedContext {
	var p = new(ReservedContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_reserved
	return p
}

func InitEmptyReservedContext(p *ReservedContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_reserved
}

func (*ReservedContext) IsReservedContext() {}

func NewReservedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReservedContext {
	var p = new(ReservedContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_reserved

	return p
}

func (s *ReservedContext) GetParser() antlr.Parser { return s.parser }

func (s *ReservedContext) RESERVED() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserRESERVED, 0)
}

func (s *ReservedContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSEMI, 0)
}

func (s *ReservedContext) Ranges() IRangesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRangesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRangesContext)
}

func (s *ReservedContext) ReservedFieldNames() IReservedFieldNamesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReservedFieldNamesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReservedFieldNamesContext)
}

func (s *ReservedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReservedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReservedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitReserved(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) Reserved() (localctx IReservedContext) {
	localctx = NewReservedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, Protobuf2ParserRULE_reserved)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(270)
		p.Match(Protobuf2ParserRESERVED)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(273)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Protobuf2ParserINT_LIT:
		{
			p.SetState(271)
			p.Ranges()
		}

	case Protobuf2ParserPROTO2_LIT_SINGLE, Protobuf2ParserPROTO2_LIT_DOBULE, Protobuf2ParserSTR_LIT:
		{
			p.SetState(272)
			p.ReservedFieldNames()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(275)
		p.Match(Protobuf2ParserSEMI)
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

// IRangesContext is an interface to support dynamic dispatch.
type IRangesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllRange_() []IRange_Context
	Range_(i int) IRange_Context
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsRangesContext differentiates from other interfaces.
	IsRangesContext()
}

type RangesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRangesContext() *RangesContext {
	var p = new(RangesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_ranges
	return p
}

func InitEmptyRangesContext(p *RangesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_ranges
}

func (*RangesContext) IsRangesContext() {}

func NewRangesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangesContext {
	var p = new(RangesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_ranges

	return p
}

func (s *RangesContext) GetParser() antlr.Parser { return s.parser }

func (s *RangesContext) AllRange_() []IRange_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IRange_Context); ok {
			len++
		}
	}

	tst := make([]IRange_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IRange_Context); ok {
			tst[i] = t.(IRange_Context)
			i++
		}
	}

	return tst
}

func (s *RangesContext) Range_(i int) IRange_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRange_Context); ok {
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

	return t.(IRange_Context)
}

func (s *RangesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(Protobuf2ParserCOMMA)
}

func (s *RangesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserCOMMA, i)
}

func (s *RangesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitRanges(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) Ranges() (localctx IRangesContext) {
	localctx = NewRangesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, Protobuf2ParserRULE_ranges)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(277)
		p.Range_()
	}
	p.SetState(282)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf2ParserCOMMA {
		{
			p.SetState(278)
			p.Match(Protobuf2ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(279)
			p.Range_()
		}

		p.SetState(284)
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

// IRange_Context is an interface to support dynamic dispatch.
type IRange_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIntLit() []IIntLitContext
	IntLit(i int) IIntLitContext
	TO() antlr.TerminalNode
	MAX() antlr.TerminalNode

	// IsRange_Context differentiates from other interfaces.
	IsRange_Context()
}

type Range_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRange_Context() *Range_Context {
	var p = new(Range_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_range_
	return p
}

func InitEmptyRange_Context(p *Range_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_range_
}

func (*Range_Context) IsRange_Context() {}

func NewRange_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Range_Context {
	var p = new(Range_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_range_

	return p
}

func (s *Range_Context) GetParser() antlr.Parser { return s.parser }

func (s *Range_Context) AllIntLit() []IIntLitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIntLitContext); ok {
			len++
		}
	}

	tst := make([]IIntLitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIntLitContext); ok {
			tst[i] = t.(IIntLitContext)
			i++
		}
	}

	return tst
}

func (s *Range_Context) IntLit(i int) IIntLitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntLitContext); ok {
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

	return t.(IIntLitContext)
}

func (s *Range_Context) TO() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserTO, 0)
}

func (s *Range_Context) MAX() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserMAX, 0)
}

func (s *Range_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Range_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Range_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitRange_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) Range_() (localctx IRange_Context) {
	localctx = NewRange_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, Protobuf2ParserRULE_range_)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(285)
		p.IntLit()
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf2ParserTO {
		{
			p.SetState(286)
			p.Match(Protobuf2ParserTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(289)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case Protobuf2ParserINT_LIT:
			{
				p.SetState(287)
				p.IntLit()
			}

		case Protobuf2ParserMAX:
			{
				p.SetState(288)
				p.Match(Protobuf2ParserMAX)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

// IReservedFieldNamesContext is an interface to support dynamic dispatch.
type IReservedFieldNamesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStrLit() []IStrLitContext
	StrLit(i int) IStrLitContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsReservedFieldNamesContext differentiates from other interfaces.
	IsReservedFieldNamesContext()
}

type ReservedFieldNamesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReservedFieldNamesContext() *ReservedFieldNamesContext {
	var p = new(ReservedFieldNamesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_reservedFieldNames
	return p
}

func InitEmptyReservedFieldNamesContext(p *ReservedFieldNamesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_reservedFieldNames
}

func (*ReservedFieldNamesContext) IsReservedFieldNamesContext() {}

func NewReservedFieldNamesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReservedFieldNamesContext {
	var p = new(ReservedFieldNamesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_reservedFieldNames

	return p
}

func (s *ReservedFieldNamesContext) GetParser() antlr.Parser { return s.parser }

func (s *ReservedFieldNamesContext) AllStrLit() []IStrLitContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStrLitContext); ok {
			len++
		}
	}

	tst := make([]IStrLitContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStrLitContext); ok {
			tst[i] = t.(IStrLitContext)
			i++
		}
	}

	return tst
}

func (s *ReservedFieldNamesContext) StrLit(i int) IStrLitContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrLitContext); ok {
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

	return t.(IStrLitContext)
}

func (s *ReservedFieldNamesContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(Protobuf2ParserCOMMA)
}

func (s *ReservedFieldNamesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserCOMMA, i)
}

func (s *ReservedFieldNamesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReservedFieldNamesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReservedFieldNamesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitReservedFieldNames(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) ReservedFieldNames() (localctx IReservedFieldNamesContext) {
	localctx = NewReservedFieldNamesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, Protobuf2ParserRULE_reservedFieldNames)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(293)
		p.StrLit()
	}
	p.SetState(298)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf2ParserCOMMA {
		{
			p.SetState(294)
			p.Match(Protobuf2ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(295)
			p.StrLit()
		}

		p.SetState(300)
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

// ITopLevelDefContext is an interface to support dynamic dispatch.
type ITopLevelDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MessageDef() IMessageDefContext
	EnumDef() IEnumDefContext
	ServiceDef() IServiceDefContext
	ExtendDef() IExtendDefContext

	// IsTopLevelDefContext differentiates from other interfaces.
	IsTopLevelDefContext()
}

type TopLevelDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTopLevelDefContext() *TopLevelDefContext {
	var p = new(TopLevelDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_topLevelDef
	return p
}

func InitEmptyTopLevelDefContext(p *TopLevelDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_topLevelDef
}

func (*TopLevelDefContext) IsTopLevelDefContext() {}

func NewTopLevelDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelDefContext {
	var p = new(TopLevelDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_topLevelDef

	return p
}

func (s *TopLevelDefContext) GetParser() antlr.Parser { return s.parser }

func (s *TopLevelDefContext) MessageDef() IMessageDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageDefContext)
}

func (s *TopLevelDefContext) EnumDef() IEnumDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumDefContext)
}

func (s *TopLevelDefContext) ServiceDef() IServiceDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceDefContext)
}

func (s *TopLevelDefContext) ExtendDef() IExtendDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtendDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtendDefContext)
}

func (s *TopLevelDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopLevelDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitTopLevelDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) TopLevelDef() (localctx ITopLevelDefContext) {
	localctx = NewTopLevelDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, Protobuf2ParserRULE_topLevelDef)
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Protobuf2ParserMESSAGE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(301)
			p.MessageDef()
		}

	case Protobuf2ParserENUM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(302)
			p.EnumDef()
		}

	case Protobuf2ParserSERVICE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(303)
			p.ServiceDef()
		}

	case Protobuf2ParserEXTEND:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(304)
			p.ExtendDef()
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

// IEnumDefContext is an interface to support dynamic dispatch.
type IEnumDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ENUM() antlr.TerminalNode
	EnumName() IEnumNameContext
	EnumBody() IEnumBodyContext

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
	p.RuleIndex = Protobuf2ParserRULE_enumDef
	return p
}

func InitEmptyEnumDefContext(p *EnumDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_enumDef
}

func (*EnumDefContext) IsEnumDefContext() {}

func NewEnumDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDefContext {
	var p = new(EnumDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_enumDef

	return p
}

func (s *EnumDefContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDefContext) ENUM() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserENUM, 0)
}

func (s *EnumDefContext) EnumName() IEnumNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumNameContext)
}

func (s *EnumDefContext) EnumBody() IEnumBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumBodyContext)
}

func (s *EnumDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitEnumDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) EnumDef() (localctx IEnumDefContext) {
	localctx = NewEnumDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, Protobuf2ParserRULE_enumDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		p.Match(Protobuf2ParserENUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(308)
		p.EnumName()
	}
	{
		p.SetState(309)
		p.EnumBody()
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

// IEnumBodyContext is an interface to support dynamic dispatch.
type IEnumBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	AllEnumElement() []IEnumElementContext
	EnumElement(i int) IEnumElementContext

	// IsEnumBodyContext differentiates from other interfaces.
	IsEnumBodyContext()
}

type EnumBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumBodyContext() *EnumBodyContext {
	var p = new(EnumBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_enumBody
	return p
}

func InitEmptyEnumBodyContext(p *EnumBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_enumBody
}

func (*EnumBodyContext) IsEnumBodyContext() {}

func NewEnumBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumBodyContext {
	var p = new(EnumBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_enumBody

	return p
}

func (s *EnumBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumBodyContext) LC() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserLC, 0)
}

func (s *EnumBodyContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserRC, 0)
}

func (s *EnumBodyContext) AllEnumElement() []IEnumElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnumElementContext); ok {
			len++
		}
	}

	tst := make([]IEnumElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnumElementContext); ok {
			tst[i] = t.(IEnumElementContext)
			i++
		}
	}

	return tst
}

func (s *EnumBodyContext) EnumElement(i int) IEnumElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumElementContext); ok {
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

	return t.(IEnumElementContext)
}

func (s *EnumBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitEnumBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) EnumBody() (localctx IEnumBodyContext) {
	localctx = NewEnumBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, Protobuf2ParserRULE_enumBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(311)
		p.Match(Protobuf2ParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(315)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1297039441461772286) != 0 {
		{
			p.SetState(312)
			p.EnumElement()
		}

		p.SetState(317)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(318)
		p.Match(Protobuf2ParserRC)
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

// IEnumElementContext is an interface to support dynamic dispatch.
type IEnumElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OptionStatement() IOptionStatementContext
	EnumField() IEnumFieldContext
	EmptyStatement_() IEmptyStatement_Context

	// IsEnumElementContext differentiates from other interfaces.
	IsEnumElementContext()
}

type EnumElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumElementContext() *EnumElementContext {
	var p = new(EnumElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_enumElement
	return p
}

func InitEmptyEnumElementContext(p *EnumElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_enumElement
}

func (*EnumElementContext) IsEnumElementContext() {}

func NewEnumElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumElementContext {
	var p = new(EnumElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_enumElement

	return p
}

func (s *EnumElementContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumElementContext) OptionStatement() IOptionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionStatementContext)
}

func (s *EnumElementContext) EnumField() IEnumFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumFieldContext)
}

func (s *EnumElementContext) EmptyStatement_() IEmptyStatement_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *EnumElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitEnumElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) EnumElement() (localctx IEnumElementContext) {
	localctx = NewEnumElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, Protobuf2ParserRULE_enumElement)
	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(320)
			p.OptionStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(321)
			p.EnumField()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(322)
			p.EmptyStatement_()
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

// IEnumFieldContext is an interface to support dynamic dispatch.
type IEnumFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext
	EQ() antlr.TerminalNode
	IntLit() IIntLitContext
	SEMI() antlr.TerminalNode
	MINUS() antlr.TerminalNode
	EnumValueOptions() IEnumValueOptionsContext

	// IsEnumFieldContext differentiates from other interfaces.
	IsEnumFieldContext()
}

type EnumFieldContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumFieldContext() *EnumFieldContext {
	var p = new(EnumFieldContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_enumField
	return p
}

func InitEmptyEnumFieldContext(p *EnumFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_enumField
}

func (*EnumFieldContext) IsEnumFieldContext() {}

func NewEnumFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumFieldContext {
	var p = new(EnumFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_enumField

	return p
}

func (s *EnumFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumFieldContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *EnumFieldContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserEQ, 0)
}

func (s *EnumFieldContext) IntLit() IIntLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntLitContext)
}

func (s *EnumFieldContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSEMI, 0)
}

func (s *EnumFieldContext) MINUS() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserMINUS, 0)
}

func (s *EnumFieldContext) EnumValueOptions() IEnumValueOptionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumValueOptionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumValueOptionsContext)
}

func (s *EnumFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitEnumField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) EnumField() (localctx IEnumFieldContext) {
	localctx = NewEnumFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, Protobuf2ParserRULE_enumField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(325)
		p.Ident()
	}
	{
		p.SetState(326)
		p.Match(Protobuf2ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(328)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf2ParserMINUS {
		{
			p.SetState(327)
			p.Match(Protobuf2ParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(330)
		p.IntLit()
	}
	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf2ParserLB {
		{
			p.SetState(331)
			p.EnumValueOptions()
		}

	}
	{
		p.SetState(334)
		p.Match(Protobuf2ParserSEMI)
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

// IEnumValueOptionsContext is an interface to support dynamic dispatch.
type IEnumValueOptionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LB() antlr.TerminalNode
	AllEnumValueOption() []IEnumValueOptionContext
	EnumValueOption(i int) IEnumValueOptionContext
	RB() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsEnumValueOptionsContext differentiates from other interfaces.
	IsEnumValueOptionsContext()
}

type EnumValueOptionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueOptionsContext() *EnumValueOptionsContext {
	var p = new(EnumValueOptionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_enumValueOptions
	return p
}

func InitEmptyEnumValueOptionsContext(p *EnumValueOptionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_enumValueOptions
}

func (*EnumValueOptionsContext) IsEnumValueOptionsContext() {}

func NewEnumValueOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueOptionsContext {
	var p = new(EnumValueOptionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_enumValueOptions

	return p
}

func (s *EnumValueOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueOptionsContext) LB() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserLB, 0)
}

func (s *EnumValueOptionsContext) AllEnumValueOption() []IEnumValueOptionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEnumValueOptionContext); ok {
			len++
		}
	}

	tst := make([]IEnumValueOptionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEnumValueOptionContext); ok {
			tst[i] = t.(IEnumValueOptionContext)
			i++
		}
	}

	return tst
}

func (s *EnumValueOptionsContext) EnumValueOption(i int) IEnumValueOptionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumValueOptionContext); ok {
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

	return t.(IEnumValueOptionContext)
}

func (s *EnumValueOptionsContext) RB() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserRB, 0)
}

func (s *EnumValueOptionsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(Protobuf2ParserCOMMA)
}

func (s *EnumValueOptionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserCOMMA, i)
}

func (s *EnumValueOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueOptionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitEnumValueOptions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) EnumValueOptions() (localctx IEnumValueOptionsContext) {
	localctx = NewEnumValueOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, Protobuf2ParserRULE_enumValueOptions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(336)
		p.Match(Protobuf2ParserLB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(337)
		p.EnumValueOption()
	}
	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf2ParserCOMMA {
		{
			p.SetState(338)
			p.Match(Protobuf2ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(339)
			p.EnumValueOption()
		}

		p.SetState(344)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(345)
		p.Match(Protobuf2ParserRB)
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

// IEnumValueOptionContext is an interface to support dynamic dispatch.
type IEnumValueOptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OptionName() IOptionNameContext
	EQ() antlr.TerminalNode
	Constant() IConstantContext

	// IsEnumValueOptionContext differentiates from other interfaces.
	IsEnumValueOptionContext()
}

type EnumValueOptionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueOptionContext() *EnumValueOptionContext {
	var p = new(EnumValueOptionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_enumValueOption
	return p
}

func InitEmptyEnumValueOptionContext(p *EnumValueOptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_enumValueOption
}

func (*EnumValueOptionContext) IsEnumValueOptionContext() {}

func NewEnumValueOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueOptionContext {
	var p = new(EnumValueOptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_enumValueOption

	return p
}

func (s *EnumValueOptionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueOptionContext) OptionName() IOptionNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionNameContext)
}

func (s *EnumValueOptionContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserEQ, 0)
}

func (s *EnumValueOptionContext) Constant() IConstantContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *EnumValueOptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueOptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueOptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitEnumValueOption(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) EnumValueOption() (localctx IEnumValueOptionContext) {
	localctx = NewEnumValueOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, Protobuf2ParserRULE_enumValueOption)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(347)
		p.OptionName()
	}
	{
		p.SetState(348)
		p.Match(Protobuf2ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(349)
		p.Constant()
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

// IMessageDefContext is an interface to support dynamic dispatch.
type IMessageDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MESSAGE() antlr.TerminalNode
	MessageName() IMessageNameContext
	MessageBody() IMessageBodyContext

	// IsMessageDefContext differentiates from other interfaces.
	IsMessageDefContext()
}

type MessageDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageDefContext() *MessageDefContext {
	var p = new(MessageDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_messageDef
	return p
}

func InitEmptyMessageDefContext(p *MessageDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_messageDef
}

func (*MessageDefContext) IsMessageDefContext() {}

func NewMessageDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageDefContext {
	var p = new(MessageDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_messageDef

	return p
}

func (s *MessageDefContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageDefContext) MESSAGE() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserMESSAGE, 0)
}

func (s *MessageDefContext) MessageName() IMessageNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageNameContext)
}

func (s *MessageDefContext) MessageBody() IMessageBodyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageBodyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageBodyContext)
}

func (s *MessageDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitMessageDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) MessageDef() (localctx IMessageDefContext) {
	localctx = NewMessageDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, Protobuf2ParserRULE_messageDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(351)
		p.Match(Protobuf2ParserMESSAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(352)
		p.MessageName()
	}
	{
		p.SetState(353)
		p.MessageBody()
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

// IMessageBodyContext is an interface to support dynamic dispatch.
type IMessageBodyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	AllMessageElement() []IMessageElementContext
	MessageElement(i int) IMessageElementContext

	// IsMessageBodyContext differentiates from other interfaces.
	IsMessageBodyContext()
}

type MessageBodyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageBodyContext() *MessageBodyContext {
	var p = new(MessageBodyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_messageBody
	return p
}

func InitEmptyMessageBodyContext(p *MessageBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_messageBody
}

func (*MessageBodyContext) IsMessageBodyContext() {}

func NewMessageBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageBodyContext {
	var p = new(MessageBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_messageBody

	return p
}

func (s *MessageBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageBodyContext) LC() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserLC, 0)
}

func (s *MessageBodyContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserRC, 0)
}

func (s *MessageBodyContext) AllMessageElement() []IMessageElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMessageElementContext); ok {
			len++
		}
	}

	tst := make([]IMessageElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMessageElementContext); ok {
			tst[i] = t.(IMessageElementContext)
			i++
		}
	}

	return tst
}

func (s *MessageBodyContext) MessageElement(i int) IMessageElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageElementContext); ok {
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

	return t.(IMessageElementContext)
}

func (s *MessageBodyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageBodyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageBodyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitMessageBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) MessageBody() (localctx IMessageBodyContext) {
	localctx = NewMessageBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, Protobuf2ParserRULE_messageBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(355)
		p.Match(Protobuf2ParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2229893340096) != 0 {
		{
			p.SetState(356)
			p.MessageElement()
		}

		p.SetState(361)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(362)
		p.Match(Protobuf2ParserRC)
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

// IMessageElementContext is an interface to support dynamic dispatch.
type IMessageElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Field() IFieldContext
	EnumDef() IEnumDefContext
	MessageDef() IMessageDefContext
	ExtendDef() IExtendDefContext
	OptionStatement() IOptionStatementContext
	Oneof() IOneofContext
	MapField() IMapFieldContext
	Extensions() IExtensionsContext
	Group() IGroupContext
	Reserved() IReservedContext
	EmptyStatement_() IEmptyStatement_Context

	// IsMessageElementContext differentiates from other interfaces.
	IsMessageElementContext()
}

type MessageElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageElementContext() *MessageElementContext {
	var p = new(MessageElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_messageElement
	return p
}

func InitEmptyMessageElementContext(p *MessageElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_messageElement
}

func (*MessageElementContext) IsMessageElementContext() {}

func NewMessageElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageElementContext {
	var p = new(MessageElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_messageElement

	return p
}

func (s *MessageElementContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageElementContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *MessageElementContext) EnumDef() IEnumDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumDefContext)
}

func (s *MessageElementContext) MessageDef() IMessageDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageDefContext)
}

func (s *MessageElementContext) ExtendDef() IExtendDefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtendDefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtendDefContext)
}

func (s *MessageElementContext) OptionStatement() IOptionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionStatementContext)
}

func (s *MessageElementContext) Oneof() IOneofContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOneofContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOneofContext)
}

func (s *MessageElementContext) MapField() IMapFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMapFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMapFieldContext)
}

func (s *MessageElementContext) Extensions() IExtensionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtensionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExtensionsContext)
}

func (s *MessageElementContext) Group() IGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupContext)
}

func (s *MessageElementContext) Reserved() IReservedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReservedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReservedContext)
}

func (s *MessageElementContext) EmptyStatement_() IEmptyStatement_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *MessageElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitMessageElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) MessageElement() (localctx IMessageElementContext) {
	localctx = NewMessageElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, Protobuf2ParserRULE_messageElement)
	p.SetState(375)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(364)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(365)
			p.EnumDef()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(366)
			p.MessageDef()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(367)
			p.ExtendDef()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(368)
			p.OptionStatement()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(369)
			p.Oneof()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(370)
			p.MapField()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(371)
			p.Extensions()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(372)
			p.Group()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(373)
			p.Reserved()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(374)
			p.EmptyStatement_()
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

// IExtendDefContext is an interface to support dynamic dispatch.
type IExtendDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EXTEND() antlr.TerminalNode
	MessageType() IMessageTypeContext
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	AllExtendElement() []IExtendElementContext
	ExtendElement(i int) IExtendElementContext

	// IsExtendDefContext differentiates from other interfaces.
	IsExtendDefContext()
}

type ExtendDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtendDefContext() *ExtendDefContext {
	var p = new(ExtendDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_extendDef
	return p
}

func InitEmptyExtendDefContext(p *ExtendDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_extendDef
}

func (*ExtendDefContext) IsExtendDefContext() {}

func NewExtendDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtendDefContext {
	var p = new(ExtendDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_extendDef

	return p
}

func (s *ExtendDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtendDefContext) EXTEND() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserEXTEND, 0)
}

func (s *ExtendDefContext) MessageType() IMessageTypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageTypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageTypeContext)
}

func (s *ExtendDefContext) LC() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserLC, 0)
}

func (s *ExtendDefContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserRC, 0)
}

func (s *ExtendDefContext) AllExtendElement() []IExtendElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExtendElementContext); ok {
			len++
		}
	}

	tst := make([]IExtendElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExtendElementContext); ok {
			tst[i] = t.(IExtendElementContext)
			i++
		}
	}

	return tst
}

func (s *ExtendDefContext) ExtendElement(i int) IExtendElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExtendElementContext); ok {
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

	return t.(IExtendElementContext)
}

func (s *ExtendDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtendDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtendDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitExtendDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) ExtendDef() (localctx IExtendDefContext) {
	localctx = NewExtendDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, Protobuf2ParserRULE_extendDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(377)
		p.Match(Protobuf2ParserEXTEND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(378)
		p.MessageType()
	}
	{
		p.SetState(379)
		p.Match(Protobuf2ParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(383)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2199023256448) != 0 {
		{
			p.SetState(380)
			p.ExtendElement()
		}

		p.SetState(385)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(386)
		p.Match(Protobuf2ParserRC)
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

// IExtendElementContext is an interface to support dynamic dispatch.
type IExtendElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Field() IFieldContext
	Group() IGroupContext
	EmptyStatement_() IEmptyStatement_Context

	// IsExtendElementContext differentiates from other interfaces.
	IsExtendElementContext()
}

type ExtendElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExtendElementContext() *ExtendElementContext {
	var p = new(ExtendElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_extendElement
	return p
}

func InitEmptyExtendElementContext(p *ExtendElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_extendElement
}

func (*ExtendElementContext) IsExtendElementContext() {}

func NewExtendElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtendElementContext {
	var p = new(ExtendElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_extendElement

	return p
}

func (s *ExtendElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtendElementContext) Field() IFieldContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFieldContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *ExtendElementContext) Group() IGroupContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupContext)
}

func (s *ExtendElementContext) EmptyStatement_() IEmptyStatement_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *ExtendElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtendElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtendElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitExtendElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) ExtendElement() (localctx IExtendElementContext) {
	localctx = NewExtendElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, Protobuf2ParserRULE_extendElement)
	p.SetState(391)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(388)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(389)
			p.Group()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(390)
			p.EmptyStatement_()
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

// IServiceDefContext is an interface to support dynamic dispatch.
type IServiceDefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SERVICE() antlr.TerminalNode
	ServiceName() IServiceNameContext
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	AllServiceElement() []IServiceElementContext
	ServiceElement(i int) IServiceElementContext

	// IsServiceDefContext differentiates from other interfaces.
	IsServiceDefContext()
}

type ServiceDefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceDefContext() *ServiceDefContext {
	var p = new(ServiceDefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_serviceDef
	return p
}

func InitEmptyServiceDefContext(p *ServiceDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_serviceDef
}

func (*ServiceDefContext) IsServiceDefContext() {}

func NewServiceDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceDefContext {
	var p = new(ServiceDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_serviceDef

	return p
}

func (s *ServiceDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceDefContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSERVICE, 0)
}

func (s *ServiceDefContext) ServiceName() IServiceNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IServiceNameContext)
}

func (s *ServiceDefContext) LC() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserLC, 0)
}

func (s *ServiceDefContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserRC, 0)
}

func (s *ServiceDefContext) AllServiceElement() []IServiceElementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IServiceElementContext); ok {
			len++
		}
	}

	tst := make([]IServiceElementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IServiceElementContext); ok {
			tst[i] = t.(IServiceElementContext)
			i++
		}
	}

	return tst
}

func (s *ServiceDefContext) ServiceElement(i int) IServiceElementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IServiceElementContext); ok {
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

	return t.(IServiceElementContext)
}

func (s *ServiceDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitServiceDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) ServiceDef() (localctx IServiceDefContext) {
	localctx = NewServiceDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, Protobuf2ParserRULE_serviceDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(393)
		p.Match(Protobuf2ParserSERVICE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(394)
		p.ServiceName()
	}
	{
		p.SetState(395)
		p.Match(Protobuf2ParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2405181685824) != 0 {
		{
			p.SetState(396)
			p.ServiceElement()
		}

		p.SetState(401)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(402)
		p.Match(Protobuf2ParserRC)
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

// IServiceElementContext is an interface to support dynamic dispatch.
type IServiceElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OptionStatement() IOptionStatementContext
	Rpc() IRpcContext
	Stream() IStreamContext
	EmptyStatement_() IEmptyStatement_Context

	// IsServiceElementContext differentiates from other interfaces.
	IsServiceElementContext()
}

type ServiceElementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceElementContext() *ServiceElementContext {
	var p = new(ServiceElementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_serviceElement
	return p
}

func InitEmptyServiceElementContext(p *ServiceElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_serviceElement
}

func (*ServiceElementContext) IsServiceElementContext() {}

func NewServiceElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceElementContext {
	var p = new(ServiceElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_serviceElement

	return p
}

func (s *ServiceElementContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceElementContext) OptionStatement() IOptionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionStatementContext)
}

func (s *ServiceElementContext) Rpc() IRpcContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRpcContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRpcContext)
}

func (s *ServiceElementContext) Stream() IStreamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStreamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStreamContext)
}

func (s *ServiceElementContext) EmptyStatement_() IEmptyStatement_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEmptyStatement_Context)
}

func (s *ServiceElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitServiceElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) ServiceElement() (localctx IServiceElementContext) {
	localctx = NewServiceElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, Protobuf2ParserRULE_serviceElement)
	p.SetState(408)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Protobuf2ParserOPTION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(404)
			p.OptionStatement()
		}

	case Protobuf2ParserRPC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(405)
			p.Rpc()
		}

	case Protobuf2ParserSTREAM:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(406)
			p.Stream()
		}

	case Protobuf2ParserSEMI:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(407)
			p.EmptyStatement_()
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

// IRpcContext is an interface to support dynamic dispatch.
type IRpcContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RPC() antlr.TerminalNode
	RpcName() IRpcNameContext
	AllLP() []antlr.TerminalNode
	LP(i int) antlr.TerminalNode
	AllMessageType() []IMessageTypeContext
	MessageType(i int) IMessageTypeContext
	AllRP() []antlr.TerminalNode
	RP(i int) antlr.TerminalNode
	RETURNS() antlr.TerminalNode
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	SEMI() antlr.TerminalNode
	AllSTREAM() []antlr.TerminalNode
	STREAM(i int) antlr.TerminalNode
	AllOptionStatement() []IOptionStatementContext
	OptionStatement(i int) IOptionStatementContext
	AllEmptyStatement_() []IEmptyStatement_Context
	EmptyStatement_(i int) IEmptyStatement_Context

	// IsRpcContext differentiates from other interfaces.
	IsRpcContext()
}

type RpcContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcContext() *RpcContext {
	var p = new(RpcContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_rpc
	return p
}

func InitEmptyRpcContext(p *RpcContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_rpc
}

func (*RpcContext) IsRpcContext() {}

func NewRpcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcContext {
	var p = new(RpcContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_rpc

	return p
}

func (s *RpcContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcContext) RPC() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserRPC, 0)
}

func (s *RpcContext) RpcName() IRpcNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRpcNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRpcNameContext)
}

func (s *RpcContext) AllLP() []antlr.TerminalNode {
	return s.GetTokens(Protobuf2ParserLP)
}

func (s *RpcContext) LP(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserLP, i)
}

func (s *RpcContext) AllMessageType() []IMessageTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMessageTypeContext); ok {
			len++
		}
	}

	tst := make([]IMessageTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMessageTypeContext); ok {
			tst[i] = t.(IMessageTypeContext)
			i++
		}
	}

	return tst
}

func (s *RpcContext) MessageType(i int) IMessageTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageTypeContext); ok {
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

	return t.(IMessageTypeContext)
}

func (s *RpcContext) AllRP() []antlr.TerminalNode {
	return s.GetTokens(Protobuf2ParserRP)
}

func (s *RpcContext) RP(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserRP, i)
}

func (s *RpcContext) RETURNS() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserRETURNS, 0)
}

func (s *RpcContext) LC() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserLC, 0)
}

func (s *RpcContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserRC, 0)
}

func (s *RpcContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSEMI, 0)
}

func (s *RpcContext) AllSTREAM() []antlr.TerminalNode {
	return s.GetTokens(Protobuf2ParserSTREAM)
}

func (s *RpcContext) STREAM(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSTREAM, i)
}

func (s *RpcContext) AllOptionStatement() []IOptionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOptionStatementContext); ok {
			len++
		}
	}

	tst := make([]IOptionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOptionStatementContext); ok {
			tst[i] = t.(IOptionStatementContext)
			i++
		}
	}

	return tst
}

func (s *RpcContext) OptionStatement(i int) IOptionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionStatementContext); ok {
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

	return t.(IOptionStatementContext)
}

func (s *RpcContext) AllEmptyStatement_() []IEmptyStatement_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			len++
		}
	}

	tst := make([]IEmptyStatement_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEmptyStatement_Context); ok {
			tst[i] = t.(IEmptyStatement_Context)
			i++
		}
	}

	return tst
}

func (s *RpcContext) EmptyStatement_(i int) IEmptyStatement_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
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

	return t.(IEmptyStatement_Context)
}

func (s *RpcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitRpc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) Rpc() (localctx IRpcContext) {
	localctx = NewRpcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, Protobuf2ParserRULE_rpc)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(410)
		p.Match(Protobuf2ParserRPC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(411)
		p.RpcName()
	}
	{
		p.SetState(412)
		p.Match(Protobuf2ParserLP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(414)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(413)
			p.Match(Protobuf2ParserSTREAM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(416)
		p.MessageType()
	}
	{
		p.SetState(417)
		p.Match(Protobuf2ParserRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(418)
		p.Match(Protobuf2ParserRETURNS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(419)
		p.Match(Protobuf2ParserLP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(421)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 31, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(420)
			p.Match(Protobuf2ParserSTREAM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(423)
		p.MessageType()
	}
	{
		p.SetState(424)
		p.Match(Protobuf2ParserRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(435)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Protobuf2ParserLC:
		{
			p.SetState(425)
			p.Match(Protobuf2ParserLC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(430)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == Protobuf2ParserOPTION || _la == Protobuf2ParserSEMI {
			p.SetState(428)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case Protobuf2ParserOPTION:
				{
					p.SetState(426)
					p.OptionStatement()
				}

			case Protobuf2ParserSEMI:
				{
					p.SetState(427)
					p.EmptyStatement_()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(432)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(433)
			p.Match(Protobuf2ParserRC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Protobuf2ParserSEMI:
		{
			p.SetState(434)
			p.Match(Protobuf2ParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IStreamContext is an interface to support dynamic dispatch.
type IStreamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STREAM() antlr.TerminalNode
	StreamName() IStreamNameContext
	LP() antlr.TerminalNode
	AllMessageType() []IMessageTypeContext
	MessageType(i int) IMessageTypeContext
	COMMA() antlr.TerminalNode
	RP() antlr.TerminalNode
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	SEMI() antlr.TerminalNode
	AllOptionStatement() []IOptionStatementContext
	OptionStatement(i int) IOptionStatementContext
	AllEmptyStatement_() []IEmptyStatement_Context
	EmptyStatement_(i int) IEmptyStatement_Context

	// IsStreamContext differentiates from other interfaces.
	IsStreamContext()
}

type StreamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStreamContext() *StreamContext {
	var p = new(StreamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_stream
	return p
}

func InitEmptyStreamContext(p *StreamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_stream
}

func (*StreamContext) IsStreamContext() {}

func NewStreamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StreamContext {
	var p = new(StreamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_stream

	return p
}

func (s *StreamContext) GetParser() antlr.Parser { return s.parser }

func (s *StreamContext) STREAM() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSTREAM, 0)
}

func (s *StreamContext) StreamName() IStreamNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStreamNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStreamNameContext)
}

func (s *StreamContext) LP() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserLP, 0)
}

func (s *StreamContext) AllMessageType() []IMessageTypeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMessageTypeContext); ok {
			len++
		}
	}

	tst := make([]IMessageTypeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMessageTypeContext); ok {
			tst[i] = t.(IMessageTypeContext)
			i++
		}
	}

	return tst
}

func (s *StreamContext) MessageType(i int) IMessageTypeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageTypeContext); ok {
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

	return t.(IMessageTypeContext)
}

func (s *StreamContext) COMMA() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserCOMMA, 0)
}

func (s *StreamContext) RP() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserRP, 0)
}

func (s *StreamContext) LC() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserLC, 0)
}

func (s *StreamContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserRC, 0)
}

func (s *StreamContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSEMI, 0)
}

func (s *StreamContext) AllOptionStatement() []IOptionStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOptionStatementContext); ok {
			len++
		}
	}

	tst := make([]IOptionStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOptionStatementContext); ok {
			tst[i] = t.(IOptionStatementContext)
			i++
		}
	}

	return tst
}

func (s *StreamContext) OptionStatement(i int) IOptionStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionStatementContext); ok {
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

	return t.(IOptionStatementContext)
}

func (s *StreamContext) AllEmptyStatement_() []IEmptyStatement_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
			len++
		}
	}

	tst := make([]IEmptyStatement_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEmptyStatement_Context); ok {
			tst[i] = t.(IEmptyStatement_Context)
			i++
		}
	}

	return tst
}

func (s *StreamContext) EmptyStatement_(i int) IEmptyStatement_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEmptyStatement_Context); ok {
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

	return t.(IEmptyStatement_Context)
}

func (s *StreamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StreamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StreamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitStream(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) Stream() (localctx IStreamContext) {
	localctx = NewStreamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, Protobuf2ParserRULE_stream)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(437)
		p.Match(Protobuf2ParserSTREAM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(438)
		p.StreamName()
	}
	{
		p.SetState(439)
		p.Match(Protobuf2ParserLP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(440)
		p.MessageType()
	}
	{
		p.SetState(441)
		p.Match(Protobuf2ParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(442)
		p.MessageType()
	}
	{
		p.SetState(443)
		p.Match(Protobuf2ParserRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(454)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Protobuf2ParserLC:
		{
			p.SetState(444)
			p.Match(Protobuf2ParserLC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(449)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == Protobuf2ParserOPTION || _la == Protobuf2ParserSEMI {
			p.SetState(447)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case Protobuf2ParserOPTION:
				{
					p.SetState(445)
					p.OptionStatement()
				}

			case Protobuf2ParserSEMI:
				{
					p.SetState(446)
					p.EmptyStatement_()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(451)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(452)
			p.Match(Protobuf2ParserRC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Protobuf2ParserSEMI:
		{
			p.SetState(453)
			p.Match(Protobuf2ParserSEMI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
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

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FullIdent() IFullIdentContext
	IntLit() IIntLitContext
	MINUS() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	FloatLit() IFloatLitContext
	StrLit() IStrLitContext
	BoolLit() IBoolLitContext
	BlockLit() IBlockLitContext

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_constant
	return p
}

func InitEmptyConstantContext(p *ConstantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_constant
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) FullIdent() IFullIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFullIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFullIdentContext)
}

func (s *ConstantContext) IntLit() IIntLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIntLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIntLitContext)
}

func (s *ConstantContext) MINUS() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserMINUS, 0)
}

func (s *ConstantContext) PLUS() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserPLUS, 0)
}

func (s *ConstantContext) FloatLit() IFloatLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFloatLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFloatLitContext)
}

func (s *ConstantContext) StrLit() IStrLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStrLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStrLitContext)
}

func (s *ConstantContext) BoolLit() IBoolLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolLitContext)
}

func (s *ConstantContext) BlockLit() IBlockLitContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockLitContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockLitContext)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, Protobuf2ParserRULE_constant)
	var _la int

	p.SetState(468)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(456)
			p.FullIdent()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(458)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Protobuf2ParserPLUS || _la == Protobuf2ParserMINUS {
			{
				p.SetState(457)
				_la = p.GetTokenStream().LA(1)

				if !(_la == Protobuf2ParserPLUS || _la == Protobuf2ParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(460)
			p.IntLit()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(462)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Protobuf2ParserPLUS || _la == Protobuf2ParserMINUS {
			{
				p.SetState(461)
				_la = p.GetTokenStream().LA(1)

				if !(_la == Protobuf2ParserPLUS || _la == Protobuf2ParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(464)
			p.FloatLit()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(465)
			p.StrLit()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(466)
			p.BoolLit()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(467)
			p.BlockLit()
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

// IBlockLitContext is an interface to support dynamic dispatch.
type IBlockLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LC() antlr.TerminalNode
	RC() antlr.TerminalNode
	AllIdent() []IIdentContext
	Ident(i int) IIdentContext
	AllCOLON() []antlr.TerminalNode
	COLON(i int) antlr.TerminalNode
	AllConstant() []IConstantContext
	Constant(i int) IConstantContext

	// IsBlockLitContext differentiates from other interfaces.
	IsBlockLitContext()
}

type BlockLitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockLitContext() *BlockLitContext {
	var p = new(BlockLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_blockLit
	return p
}

func InitEmptyBlockLitContext(p *BlockLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_blockLit
}

func (*BlockLitContext) IsBlockLitContext() {}

func NewBlockLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockLitContext {
	var p = new(BlockLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_blockLit

	return p
}

func (s *BlockLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockLitContext) LC() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserLC, 0)
}

func (s *BlockLitContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserRC, 0)
}

func (s *BlockLitContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *BlockLitContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
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

	return t.(IIdentContext)
}

func (s *BlockLitContext) AllCOLON() []antlr.TerminalNode {
	return s.GetTokens(Protobuf2ParserCOLON)
}

func (s *BlockLitContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserCOLON, i)
}

func (s *BlockLitContext) AllConstant() []IConstantContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConstantContext); ok {
			len++
		}
	}

	tst := make([]IConstantContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConstantContext); ok {
			tst[i] = t.(IConstantContext)
			i++
		}
	}

	return tst
}

func (s *BlockLitContext) Constant(i int) IConstantContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstantContext); ok {
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

	return t.(IConstantContext)
}

func (s *BlockLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitBlockLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) BlockLit() (localctx IBlockLitContext) {
	localctx = NewBlockLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, Protobuf2ParserRULE_blockLit)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(470)
		p.Match(Protobuf2ParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(477)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1297037242438516734) != 0 {
		{
			p.SetState(471)
			p.Ident()
		}
		{
			p.SetState(472)
			p.Match(Protobuf2ParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(473)
			p.Constant()
		}

		p.SetState(479)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(480)
		p.Match(Protobuf2ParserRC)
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

// IEmptyStatement_Context is an interface to support dynamic dispatch.
type IEmptyStatement_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SEMI() antlr.TerminalNode

	// IsEmptyStatement_Context differentiates from other interfaces.
	IsEmptyStatement_Context()
}

type EmptyStatement_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEmptyStatement_Context() *EmptyStatement_Context {
	var p = new(EmptyStatement_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_emptyStatement_
	return p
}

func InitEmptyEmptyStatement_Context(p *EmptyStatement_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_emptyStatement_
}

func (*EmptyStatement_Context) IsEmptyStatement_Context() {}

func NewEmptyStatement_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyStatement_Context {
	var p = new(EmptyStatement_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_emptyStatement_

	return p
}

func (s *EmptyStatement_Context) GetParser() antlr.Parser { return s.parser }

func (s *EmptyStatement_Context) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSEMI, 0)
}

func (s *EmptyStatement_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyStatement_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyStatement_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitEmptyStatement_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) EmptyStatement_() (localctx IEmptyStatement_Context) {
	localctx = NewEmptyStatement_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, Protobuf2ParserRULE_emptyStatement_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(482)
		p.Match(Protobuf2ParserSEMI)
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

// IIdentContext is an interface to support dynamic dispatch.
type IIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IDENTIFIER() antlr.TerminalNode
	Keywords() IKeywordsContext

	// IsIdentContext differentiates from other interfaces.
	IsIdentContext()
}

type IdentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIdentContext() *IdentContext {
	var p = new(IdentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_ident
	return p
}

func InitEmptyIdentContext(p *IdentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_ident
}

func (*IdentContext) IsIdentContext() {}

func NewIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentContext {
	var p = new(IdentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_ident

	return p
}

func (s *IdentContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserIDENTIFIER, 0)
}

func (s *IdentContext) Keywords() IKeywordsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IKeywordsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IKeywordsContext)
}

func (s *IdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) Ident() (localctx IIdentContext) {
	localctx = NewIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, Protobuf2ParserRULE_ident)
	p.SetState(486)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Protobuf2ParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(484)
			p.Match(Protobuf2ParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Protobuf2ParserSYNTAX, Protobuf2ParserIMPORT, Protobuf2ParserWEAK, Protobuf2ParserPUBLIC, Protobuf2ParserPACKAGE, Protobuf2ParserOPTION, Protobuf2ParserREPEATED, Protobuf2ParserOPTIONAL, Protobuf2ParserREQUIRED, Protobuf2ParserGROUP, Protobuf2ParserONEOF, Protobuf2ParserMAP, Protobuf2ParserINT32, Protobuf2ParserINT64, Protobuf2ParserUINT32, Protobuf2ParserUINT64, Protobuf2ParserSINT32, Protobuf2ParserSINT64, Protobuf2ParserFIXED32, Protobuf2ParserFIXED64, Protobuf2ParserSFIXED32, Protobuf2ParserSFIXED64, Protobuf2ParserBOOL, Protobuf2ParserSTRING, Protobuf2ParserDOUBLE, Protobuf2ParserFLOAT, Protobuf2ParserBYTES, Protobuf2ParserRESERVED, Protobuf2ParserEXTENSIONS, Protobuf2ParserTO, Protobuf2ParserMAX, Protobuf2ParserENUM, Protobuf2ParserEXTEND, Protobuf2ParserMESSAGE, Protobuf2ParserSERVICE, Protobuf2ParserRPC, Protobuf2ParserSTREAM, Protobuf2ParserRETURNS, Protobuf2ParserBOOL_LIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(485)
			p.Keywords()
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

// IFullIdentContext is an interface to support dynamic dispatch.
type IFullIdentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllIdent() []IIdentContext
	Ident(i int) IIdentContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsFullIdentContext differentiates from other interfaces.
	IsFullIdentContext()
}

type FullIdentContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFullIdentContext() *FullIdentContext {
	var p = new(FullIdentContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_fullIdent
	return p
}

func InitEmptyFullIdentContext(p *FullIdentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_fullIdent
}

func (*FullIdentContext) IsFullIdentContext() {}

func NewFullIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FullIdentContext {
	var p = new(FullIdentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_fullIdent

	return p
}

func (s *FullIdentContext) GetParser() antlr.Parser { return s.parser }

func (s *FullIdentContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *FullIdentContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
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

	return t.(IIdentContext)
}

func (s *FullIdentContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(Protobuf2ParserDOT)
}

func (s *FullIdentContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserDOT, i)
}

func (s *FullIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FullIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitFullIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) FullIdent() (localctx IFullIdentContext) {
	localctx = NewFullIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, Protobuf2ParserRULE_fullIdent)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(488)
		p.Ident()
	}
	p.SetState(493)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf2ParserDOT {
		{
			p.SetState(489)
			p.Match(Protobuf2ParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(490)
			p.Ident()
		}

		p.SetState(495)
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

// IMessageNameContext is an interface to support dynamic dispatch.
type IMessageNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsMessageNameContext differentiates from other interfaces.
	IsMessageNameContext()
}

type MessageNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageNameContext() *MessageNameContext {
	var p = new(MessageNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_messageName
	return p
}

func InitEmptyMessageNameContext(p *MessageNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_messageName
}

func (*MessageNameContext) IsMessageNameContext() {}

func NewMessageNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageNameContext {
	var p = new(MessageNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_messageName

	return p
}

func (s *MessageNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *MessageNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitMessageName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) MessageName() (localctx IMessageNameContext) {
	localctx = NewMessageNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, Protobuf2ParserRULE_messageName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(496)
		p.Ident()
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

// IEnumNameContext is an interface to support dynamic dispatch.
type IEnumNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsEnumNameContext differentiates from other interfaces.
	IsEnumNameContext()
}

type EnumNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumNameContext() *EnumNameContext {
	var p = new(EnumNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_enumName
	return p
}

func InitEmptyEnumNameContext(p *EnumNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_enumName
}

func (*EnumNameContext) IsEnumNameContext() {}

func NewEnumNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumNameContext {
	var p = new(EnumNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_enumName

	return p
}

func (s *EnumNameContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *EnumNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitEnumName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) EnumName() (localctx IEnumNameContext) {
	localctx = NewEnumNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, Protobuf2ParserRULE_enumName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(498)
		p.Ident()
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

// IFieldNameContext is an interface to support dynamic dispatch.
type IFieldNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsFieldNameContext differentiates from other interfaces.
	IsFieldNameContext()
}

type FieldNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldNameContext() *FieldNameContext {
	var p = new(FieldNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_fieldName
	return p
}

func InitEmptyFieldNameContext(p *FieldNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_fieldName
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_fieldName

	return p
}

func (s *FieldNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *FieldNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitFieldName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) FieldName() (localctx IFieldNameContext) {
	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, Protobuf2ParserRULE_fieldName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(500)
		p.Ident()
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

// IGroupNameContext is an interface to support dynamic dispatch.
type IGroupNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsGroupNameContext differentiates from other interfaces.
	IsGroupNameContext()
}

type GroupNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupNameContext() *GroupNameContext {
	var p = new(GroupNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_groupName
	return p
}

func InitEmptyGroupNameContext(p *GroupNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_groupName
}

func (*GroupNameContext) IsGroupNameContext() {}

func NewGroupNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupNameContext {
	var p = new(GroupNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_groupName

	return p
}

func (s *GroupNameContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *GroupNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitGroupName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) GroupName() (localctx IGroupNameContext) {
	localctx = NewGroupNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, Protobuf2ParserRULE_groupName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(502)
		p.Ident()
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

// IOneofNameContext is an interface to support dynamic dispatch.
type IOneofNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsOneofNameContext differentiates from other interfaces.
	IsOneofNameContext()
}

type OneofNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOneofNameContext() *OneofNameContext {
	var p = new(OneofNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_oneofName
	return p
}

func InitEmptyOneofNameContext(p *OneofNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_oneofName
}

func (*OneofNameContext) IsOneofNameContext() {}

func NewOneofNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofNameContext {
	var p = new(OneofNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_oneofName

	return p
}

func (s *OneofNameContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *OneofNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneofNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitOneofName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) OneofName() (localctx IOneofNameContext) {
	localctx = NewOneofNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, Protobuf2ParserRULE_oneofName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(504)
		p.Ident()
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

// IMapNameContext is an interface to support dynamic dispatch.
type IMapNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsMapNameContext differentiates from other interfaces.
	IsMapNameContext()
}

type MapNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMapNameContext() *MapNameContext {
	var p = new(MapNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_mapName
	return p
}

func InitEmptyMapNameContext(p *MapNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_mapName
}

func (*MapNameContext) IsMapNameContext() {}

func NewMapNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapNameContext {
	var p = new(MapNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_mapName

	return p
}

func (s *MapNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MapNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *MapNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitMapName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) MapName() (localctx IMapNameContext) {
	localctx = NewMapNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, Protobuf2ParserRULE_mapName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(506)
		p.Ident()
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

// IServiceNameContext is an interface to support dynamic dispatch.
type IServiceNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsServiceNameContext differentiates from other interfaces.
	IsServiceNameContext()
}

type ServiceNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyServiceNameContext() *ServiceNameContext {
	var p = new(ServiceNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_serviceName
	return p
}

func InitEmptyServiceNameContext(p *ServiceNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_serviceName
}

func (*ServiceNameContext) IsServiceNameContext() {}

func NewServiceNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceNameContext {
	var p = new(ServiceNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_serviceName

	return p
}

func (s *ServiceNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *ServiceNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ServiceNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ServiceNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitServiceName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) ServiceName() (localctx IServiceNameContext) {
	localctx = NewServiceNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, Protobuf2ParserRULE_serviceName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(508)
		p.Ident()
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

// IRpcNameContext is an interface to support dynamic dispatch.
type IRpcNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsRpcNameContext differentiates from other interfaces.
	IsRpcNameContext()
}

type RpcNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRpcNameContext() *RpcNameContext {
	var p = new(RpcNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_rpcName
	return p
}

func InitEmptyRpcNameContext(p *RpcNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_rpcName
}

func (*RpcNameContext) IsRpcNameContext() {}

func NewRpcNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcNameContext {
	var p = new(RpcNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_rpcName

	return p
}

func (s *RpcNameContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *RpcNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RpcNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RpcNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitRpcName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) RpcName() (localctx IRpcNameContext) {
	localctx = NewRpcNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, Protobuf2ParserRULE_rpcName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(510)
		p.Ident()
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

// IStreamNameContext is an interface to support dynamic dispatch.
type IStreamNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Ident() IIdentContext

	// IsStreamNameContext differentiates from other interfaces.
	IsStreamNameContext()
}

type StreamNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStreamNameContext() *StreamNameContext {
	var p = new(StreamNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_streamName
	return p
}

func InitEmptyStreamNameContext(p *StreamNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_streamName
}

func (*StreamNameContext) IsStreamNameContext() {}

func NewStreamNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StreamNameContext {
	var p = new(StreamNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_streamName

	return p
}

func (s *StreamNameContext) GetParser() antlr.Parser { return s.parser }

func (s *StreamNameContext) Ident() IIdentContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIdentContext)
}

func (s *StreamNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StreamNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StreamNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitStreamName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) StreamName() (localctx IStreamNameContext) {
	localctx = NewStreamNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, Protobuf2ParserRULE_streamName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(512)
		p.Ident()
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

// IMessageTypeContext is an interface to support dynamic dispatch.
type IMessageTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	MessageName() IMessageNameContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	AllIdent() []IIdentContext
	Ident(i int) IIdentContext

	// IsMessageTypeContext differentiates from other interfaces.
	IsMessageTypeContext()
}

type MessageTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMessageTypeContext() *MessageTypeContext {
	var p = new(MessageTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_messageType
	return p
}

func InitEmptyMessageTypeContext(p *MessageTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_messageType
}

func (*MessageTypeContext) IsMessageTypeContext() {}

func NewMessageTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageTypeContext {
	var p = new(MessageTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_messageType

	return p
}

func (s *MessageTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageTypeContext) MessageName() IMessageNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMessageNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMessageNameContext)
}

func (s *MessageTypeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(Protobuf2ParserDOT)
}

func (s *MessageTypeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserDOT, i)
}

func (s *MessageTypeContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *MessageTypeContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
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

	return t.(IIdentContext)
}

func (s *MessageTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MessageTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MessageTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitMessageType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) MessageType() (localctx IMessageTypeContext) {
	localctx = NewMessageTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, Protobuf2ParserRULE_messageType)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(515)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf2ParserDOT {
		{
			p.SetState(514)
			p.Match(Protobuf2ParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(522)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(517)
				p.Ident()
			}
			{
				p.SetState(518)
				p.Match(Protobuf2ParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(524)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 45, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(525)
		p.MessageName()
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

// IEnumTypeContext is an interface to support dynamic dispatch.
type IEnumTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EnumName() IEnumNameContext
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode
	AllIdent() []IIdentContext
	Ident(i int) IIdentContext

	// IsEnumTypeContext differentiates from other interfaces.
	IsEnumTypeContext()
}

type EnumTypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumTypeContext() *EnumTypeContext {
	var p = new(EnumTypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_enumType
	return p
}

func InitEmptyEnumTypeContext(p *EnumTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_enumType
}

func (*EnumTypeContext) IsEnumTypeContext() {}

func NewEnumTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumTypeContext {
	var p = new(EnumTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_enumType

	return p
}

func (s *EnumTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumTypeContext) EnumName() IEnumNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEnumNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEnumNameContext)
}

func (s *EnumTypeContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(Protobuf2ParserDOT)
}

func (s *EnumTypeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserDOT, i)
}

func (s *EnumTypeContext) AllIdent() []IIdentContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IIdentContext); ok {
			len++
		}
	}

	tst := make([]IIdentContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IIdentContext); ok {
			tst[i] = t.(IIdentContext)
			i++
		}
	}

	return tst
}

func (s *EnumTypeContext) Ident(i int) IIdentContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIdentContext); ok {
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

	return t.(IIdentContext)
}

func (s *EnumTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitEnumType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) EnumType() (localctx IEnumTypeContext) {
	localctx = NewEnumTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, Protobuf2ParserRULE_enumType)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(528)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf2ParserDOT {
		{
			p.SetState(527)
			p.Match(Protobuf2ParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(535)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(530)
				p.Ident()
			}
			{
				p.SetState(531)
				p.Match(Protobuf2ParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(537)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 47, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(538)
		p.EnumName()
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

// IIntLitContext is an interface to support dynamic dispatch.
type IIntLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT_LIT() antlr.TerminalNode

	// IsIntLitContext differentiates from other interfaces.
	IsIntLitContext()
}

type IntLitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntLitContext() *IntLitContext {
	var p = new(IntLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_intLit
	return p
}

func InitEmptyIntLitContext(p *IntLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_intLit
}

func (*IntLitContext) IsIntLitContext() {}

func NewIntLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntLitContext {
	var p = new(IntLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_intLit

	return p
}

func (s *IntLitContext) GetParser() antlr.Parser { return s.parser }

func (s *IntLitContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserINT_LIT, 0)
}

func (s *IntLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitIntLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) IntLit() (localctx IIntLitContext) {
	localctx = NewIntLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, Protobuf2ParserRULE_intLit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(540)
		p.Match(Protobuf2ParserINT_LIT)
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

// IStrLitContext is an interface to support dynamic dispatch.
type IStrLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STR_LIT() antlr.TerminalNode
	PROTO2_LIT_SINGLE() antlr.TerminalNode
	PROTO2_LIT_DOBULE() antlr.TerminalNode

	// IsStrLitContext differentiates from other interfaces.
	IsStrLitContext()
}

type StrLitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStrLitContext() *StrLitContext {
	var p = new(StrLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_strLit
	return p
}

func InitEmptyStrLitContext(p *StrLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_strLit
}

func (*StrLitContext) IsStrLitContext() {}

func NewStrLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrLitContext {
	var p = new(StrLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_strLit

	return p
}

func (s *StrLitContext) GetParser() antlr.Parser { return s.parser }

func (s *StrLitContext) STR_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSTR_LIT, 0)
}

func (s *StrLitContext) PROTO2_LIT_SINGLE() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserPROTO2_LIT_SINGLE, 0)
}

func (s *StrLitContext) PROTO2_LIT_DOBULE() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserPROTO2_LIT_DOBULE, 0)
}

func (s *StrLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitStrLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) StrLit() (localctx IStrLitContext) {
	localctx = NewStrLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, Protobuf2ParserRULE_strLit)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(542)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&72059243305369600) != 0) {
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

// IBoolLitContext is an interface to support dynamic dispatch.
type IBoolLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BOOL_LIT() antlr.TerminalNode

	// IsBoolLitContext differentiates from other interfaces.
	IsBoolLitContext()
}

type BoolLitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolLitContext() *BoolLitContext {
	var p = new(BoolLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_boolLit
	return p
}

func InitEmptyBoolLitContext(p *BoolLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_boolLit
}

func (*BoolLitContext) IsBoolLitContext() {}

func NewBoolLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolLitContext {
	var p = new(BoolLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_boolLit

	return p
}

func (s *BoolLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolLitContext) BOOL_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserBOOL_LIT, 0)
}

func (s *BoolLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitBoolLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) BoolLit() (localctx IBoolLitContext) {
	localctx = NewBoolLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, Protobuf2ParserRULE_boolLit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(544)
		p.Match(Protobuf2ParserBOOL_LIT)
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

// IFloatLitContext is an interface to support dynamic dispatch.
type IFloatLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FLOAT_LIT() antlr.TerminalNode

	// IsFloatLitContext differentiates from other interfaces.
	IsFloatLitContext()
}

type FloatLitContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatLitContext() *FloatLitContext {
	var p = new(FloatLitContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_floatLit
	return p
}

func InitEmptyFloatLitContext(p *FloatLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_floatLit
}

func (*FloatLitContext) IsFloatLitContext() {}

func NewFloatLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatLitContext {
	var p = new(FloatLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_floatLit

	return p
}

func (s *FloatLitContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatLitContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserFLOAT_LIT, 0)
}

func (s *FloatLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitFloatLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) FloatLit() (localctx IFloatLitContext) {
	localctx = NewFloatLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, Protobuf2ParserRULE_floatLit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(546)
		p.Match(Protobuf2ParserFLOAT_LIT)
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

// IKeywordsContext is an interface to support dynamic dispatch.
type IKeywordsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SYNTAX() antlr.TerminalNode
	IMPORT() antlr.TerminalNode
	WEAK() antlr.TerminalNode
	PUBLIC() antlr.TerminalNode
	PACKAGE() antlr.TerminalNode
	OPTION() antlr.TerminalNode
	REPEATED() antlr.TerminalNode
	OPTIONAL() antlr.TerminalNode
	REQUIRED() antlr.TerminalNode
	GROUP() antlr.TerminalNode
	ONEOF() antlr.TerminalNode
	MAP() antlr.TerminalNode
	INT32() antlr.TerminalNode
	INT64() antlr.TerminalNode
	UINT32() antlr.TerminalNode
	UINT64() antlr.TerminalNode
	SINT32() antlr.TerminalNode
	SINT64() antlr.TerminalNode
	FIXED32() antlr.TerminalNode
	FIXED64() antlr.TerminalNode
	SFIXED32() antlr.TerminalNode
	SFIXED64() antlr.TerminalNode
	BOOL() antlr.TerminalNode
	STRING() antlr.TerminalNode
	DOUBLE() antlr.TerminalNode
	FLOAT() antlr.TerminalNode
	BYTES() antlr.TerminalNode
	RESERVED() antlr.TerminalNode
	EXTENSIONS() antlr.TerminalNode
	TO() antlr.TerminalNode
	MAX() antlr.TerminalNode
	ENUM() antlr.TerminalNode
	MESSAGE() antlr.TerminalNode
	EXTEND() antlr.TerminalNode
	SERVICE() antlr.TerminalNode
	RPC() antlr.TerminalNode
	STREAM() antlr.TerminalNode
	RETURNS() antlr.TerminalNode
	BOOL_LIT() antlr.TerminalNode

	// IsKeywordsContext differentiates from other interfaces.
	IsKeywordsContext()
}

type KeywordsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyKeywordsContext() *KeywordsContext {
	var p = new(KeywordsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_keywords
	return p
}

func InitEmptyKeywordsContext(p *KeywordsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf2ParserRULE_keywords
}

func (*KeywordsContext) IsKeywordsContext() {}

func NewKeywordsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordsContext {
	var p = new(KeywordsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf2ParserRULE_keywords

	return p
}

func (s *KeywordsContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordsContext) SYNTAX() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSYNTAX, 0)
}

func (s *KeywordsContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserIMPORT, 0)
}

func (s *KeywordsContext) WEAK() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserWEAK, 0)
}

func (s *KeywordsContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserPUBLIC, 0)
}

func (s *KeywordsContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserPACKAGE, 0)
}

func (s *KeywordsContext) OPTION() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserOPTION, 0)
}

func (s *KeywordsContext) REPEATED() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserREPEATED, 0)
}

func (s *KeywordsContext) OPTIONAL() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserOPTIONAL, 0)
}

func (s *KeywordsContext) REQUIRED() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserREQUIRED, 0)
}

func (s *KeywordsContext) GROUP() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserGROUP, 0)
}

func (s *KeywordsContext) ONEOF() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserONEOF, 0)
}

func (s *KeywordsContext) MAP() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserMAP, 0)
}

func (s *KeywordsContext) INT32() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserINT32, 0)
}

func (s *KeywordsContext) INT64() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserINT64, 0)
}

func (s *KeywordsContext) UINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserUINT32, 0)
}

func (s *KeywordsContext) UINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserUINT64, 0)
}

func (s *KeywordsContext) SINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSINT32, 0)
}

func (s *KeywordsContext) SINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSINT64, 0)
}

func (s *KeywordsContext) FIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserFIXED32, 0)
}

func (s *KeywordsContext) FIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserFIXED64, 0)
}

func (s *KeywordsContext) SFIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSFIXED32, 0)
}

func (s *KeywordsContext) SFIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSFIXED64, 0)
}

func (s *KeywordsContext) BOOL() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserBOOL, 0)
}

func (s *KeywordsContext) STRING() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSTRING, 0)
}

func (s *KeywordsContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserDOUBLE, 0)
}

func (s *KeywordsContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserFLOAT, 0)
}

func (s *KeywordsContext) BYTES() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserBYTES, 0)
}

func (s *KeywordsContext) RESERVED() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserRESERVED, 0)
}

func (s *KeywordsContext) EXTENSIONS() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserEXTENSIONS, 0)
}

func (s *KeywordsContext) TO() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserTO, 0)
}

func (s *KeywordsContext) MAX() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserMAX, 0)
}

func (s *KeywordsContext) ENUM() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserENUM, 0)
}

func (s *KeywordsContext) MESSAGE() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserMESSAGE, 0)
}

func (s *KeywordsContext) EXTEND() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserEXTEND, 0)
}

func (s *KeywordsContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSERVICE, 0)
}

func (s *KeywordsContext) RPC() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserRPC, 0)
}

func (s *KeywordsContext) STREAM() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserSTREAM, 0)
}

func (s *KeywordsContext) RETURNS() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserRETURNS, 0)
}

func (s *KeywordsContext) BOOL_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf2ParserBOOL_LIT, 0)
}

func (s *KeywordsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf2Visitor:
		return t.VisitKeywords(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf2Parser) Keywords() (localctx IKeywordsContext) {
	localctx = NewKeywordsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, Protobuf2ParserRULE_keywords)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(548)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&144115737831669758) != 0) {
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
