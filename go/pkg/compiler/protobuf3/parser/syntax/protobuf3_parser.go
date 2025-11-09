// Code generated from Protobuf3.g4 by ANTLR 4.13.0. DO NOT EDIT.

package syntax // Protobuf3
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

type Protobuf3Parser struct {
	*antlr.BaseParser
}

var Protobuf3ParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func protobuf3ParserInit() {
	staticData := &Protobuf3ParserStaticData
	staticData.LiteralNames = []string{
		"", "'syntax'", "'import'", "'weak'", "'public'", "'package'", "'option'",
		"'repeated'", "'oneof'", "'map'", "'int32'", "'int64'", "'uint32'",
		"'uint64'", "'sint32'", "'sint64'", "'fixed32'", "'fixed64'", "'sfixed32'",
		"'sfixed64'", "'bool'", "'string'", "'double'", "'float'", "'bytes'",
		"'reserved'", "'to'", "'max'", "'enum'", "'message'", "'service'", "'rpc'",
		"'stream'", "'returns'", "'\"proto3\"'", "''proto3''", "';'", "'='",
		"'('", "')'", "'['", "']'", "'{'", "'}'", "'<'", "'>'", "'.'", "','",
		"':'", "'+'", "'-'",
	}
	staticData.SymbolicNames = []string{
		"", "SYNTAX", "IMPORT", "WEAK", "PUBLIC", "PACKAGE", "OPTION", "REPEATED",
		"ONEOF", "MAP", "INT32", "INT64", "UINT32", "UINT64", "SINT32", "SINT64",
		"FIXED32", "FIXED64", "SFIXED32", "SFIXED64", "BOOL", "STRING", "DOUBLE",
		"FLOAT", "BYTES", "RESERVED", "TO", "MAX", "ENUM", "MESSAGE", "SERVICE",
		"RPC", "STREAM", "RETURNS", "PROTO3_LIT_DOBULE", "PROTO3_LIT_SINGLE",
		"SEMI", "EQ", "LP", "RP", "LB", "RB", "LC", "RC", "LT", "GT", "DOT",
		"COMMA", "COLON", "PLUS", "MINUS", "STR_LIT", "BOOL_LIT", "FLOAT_LIT",
		"INT_LIT", "IDENTIFIER", "WS", "LINE_COMMENT", "COMMENT",
	}
	staticData.RuleNames = []string{
		"proto", "syntax", "importStatement", "packageStatement", "optionStatement",
		"optionName", "field", "fieldOptions", "fieldOption", "fieldNumber",
		"oneof", "oneofField", "mapField", "keyType", "type_", "reserved", "ranges",
		"range_", "reservedFieldNames", "topLevelDef", "enumDef", "enumBody",
		"enumElement", "enumField", "enumValueOptions", "enumValueOption", "messageDef",
		"messageBody", "messageElement", "serviceDef", "serviceElement", "rpc",
		"requestStream", "responseStream", "constant", "blockLit", "emptyStatement_",
		"ident", "fullIdent", "messageName", "enumName", "fieldName", "oneofName",
		"mapName", "serviceName", "rpcName", "messageType", "enumType", "intLit",
		"strLit", "boolLit", "floatLit", "keywords",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 58, 484, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 5, 0, 113, 8, 0, 10, 0, 12, 0,
		116, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 3, 2,
		127, 8, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4,
		1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 148, 8, 5,
		3, 5, 150, 8, 5, 1, 6, 3, 6, 153, 8, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 3, 6, 163, 8, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 5, 7, 170,
		8, 7, 10, 7, 12, 7, 173, 9, 7, 1, 8, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10,
		1, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 187, 8, 10, 10, 10, 12, 10, 190,
		9, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1,
		11, 3, 11, 202, 8, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 3, 12, 219, 8,
		12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1,
		14, 3, 14, 242, 8, 14, 1, 15, 1, 15, 1, 15, 3, 15, 247, 8, 15, 1, 15, 1,
		15, 1, 16, 1, 16, 1, 16, 5, 16, 254, 8, 16, 10, 16, 12, 16, 257, 9, 16,
		1, 17, 1, 17, 1, 17, 1, 17, 3, 17, 263, 8, 17, 3, 17, 265, 8, 17, 1, 18,
		1, 18, 1, 18, 5, 18, 270, 8, 18, 10, 18, 12, 18, 273, 9, 18, 1, 19, 1,
		19, 1, 19, 3, 19, 278, 8, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 21, 1, 21,
		5, 21, 286, 8, 21, 10, 21, 12, 21, 289, 9, 21, 1, 21, 1, 21, 1, 22, 1,
		22, 1, 22, 3, 22, 296, 8, 22, 1, 23, 1, 23, 1, 23, 3, 23, 301, 8, 23, 1,
		23, 1, 23, 3, 23, 305, 8, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24,
		5, 24, 313, 8, 24, 10, 24, 12, 24, 316, 9, 24, 1, 24, 1, 24, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 5, 27, 330,
		8, 27, 10, 27, 12, 27, 333, 9, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 345, 8, 28, 1, 29, 1, 29, 1, 29,
		1, 29, 5, 29, 351, 8, 29, 10, 29, 12, 29, 354, 9, 29, 1, 29, 1, 29, 1,
		30, 1, 30, 1, 30, 3, 30, 361, 8, 30, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31,
		367, 8, 31, 1, 31, 1, 31, 1, 31, 1, 31, 1, 31, 3, 31, 374, 8, 31, 1, 31,
		1, 31, 1, 31, 1, 31, 1, 31, 5, 31, 381, 8, 31, 10, 31, 12, 31, 384, 9,
		31, 1, 31, 1, 31, 3, 31, 388, 8, 31, 1, 32, 1, 32, 1, 33, 1, 33, 1, 34,
		1, 34, 3, 34, 396, 8, 34, 1, 34, 1, 34, 3, 34, 400, 8, 34, 1, 34, 1, 34,
		1, 34, 1, 34, 3, 34, 406, 8, 34, 1, 35, 1, 35, 1, 35, 1, 35, 1, 35, 5,
		35, 413, 8, 35, 10, 35, 12, 35, 416, 9, 35, 1, 35, 1, 35, 1, 36, 1, 36,
		1, 37, 1, 37, 3, 37, 424, 8, 37, 1, 38, 1, 38, 1, 38, 5, 38, 429, 8, 38,
		10, 38, 12, 38, 432, 9, 38, 1, 39, 1, 39, 1, 40, 1, 40, 1, 41, 1, 41, 1,
		42, 1, 42, 1, 43, 1, 43, 1, 44, 1, 44, 1, 45, 1, 45, 1, 46, 3, 46, 449,
		8, 46, 1, 46, 1, 46, 1, 46, 5, 46, 454, 8, 46, 10, 46, 12, 46, 457, 9,
		46, 1, 46, 1, 46, 1, 47, 3, 47, 462, 8, 47, 1, 47, 1, 47, 1, 47, 5, 47,
		467, 8, 47, 10, 47, 12, 47, 470, 9, 47, 1, 47, 1, 47, 1, 48, 1, 48, 1,
		49, 1, 49, 1, 50, 1, 50, 1, 51, 1, 51, 1, 52, 1, 52, 1, 52, 0, 0, 53, 0,
		2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
		40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74,
		76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 0, 6, 1,
		0, 34, 35, 1, 0, 3, 4, 1, 0, 10, 21, 1, 0, 49, 50, 2, 0, 34, 35, 51, 51,
		2, 0, 1, 33, 52, 52, 505, 0, 106, 1, 0, 0, 0, 2, 119, 1, 0, 0, 0, 4, 124,
		1, 0, 0, 0, 6, 131, 1, 0, 0, 0, 8, 135, 1, 0, 0, 0, 10, 149, 1, 0, 0, 0,
		12, 152, 1, 0, 0, 0, 14, 166, 1, 0, 0, 0, 16, 174, 1, 0, 0, 0, 18, 178,
		1, 0, 0, 0, 20, 180, 1, 0, 0, 0, 22, 193, 1, 0, 0, 0, 24, 205, 1, 0, 0,
		0, 26, 222, 1, 0, 0, 0, 28, 241, 1, 0, 0, 0, 30, 243, 1, 0, 0, 0, 32, 250,
		1, 0, 0, 0, 34, 258, 1, 0, 0, 0, 36, 266, 1, 0, 0, 0, 38, 277, 1, 0, 0,
		0, 40, 279, 1, 0, 0, 0, 42, 283, 1, 0, 0, 0, 44, 295, 1, 0, 0, 0, 46, 297,
		1, 0, 0, 0, 48, 308, 1, 0, 0, 0, 50, 319, 1, 0, 0, 0, 52, 323, 1, 0, 0,
		0, 54, 327, 1, 0, 0, 0, 56, 344, 1, 0, 0, 0, 58, 346, 1, 0, 0, 0, 60, 360,
		1, 0, 0, 0, 62, 362, 1, 0, 0, 0, 64, 389, 1, 0, 0, 0, 66, 391, 1, 0, 0,
		0, 68, 405, 1, 0, 0, 0, 70, 407, 1, 0, 0, 0, 72, 419, 1, 0, 0, 0, 74, 423,
		1, 0, 0, 0, 76, 425, 1, 0, 0, 0, 78, 433, 1, 0, 0, 0, 80, 435, 1, 0, 0,
		0, 82, 437, 1, 0, 0, 0, 84, 439, 1, 0, 0, 0, 86, 441, 1, 0, 0, 0, 88, 443,
		1, 0, 0, 0, 90, 445, 1, 0, 0, 0, 92, 448, 1, 0, 0, 0, 94, 461, 1, 0, 0,
		0, 96, 473, 1, 0, 0, 0, 98, 475, 1, 0, 0, 0, 100, 477, 1, 0, 0, 0, 102,
		479, 1, 0, 0, 0, 104, 481, 1, 0, 0, 0, 106, 114, 3, 2, 1, 0, 107, 113,
		3, 4, 2, 0, 108, 113, 3, 6, 3, 0, 109, 113, 3, 8, 4, 0, 110, 113, 3, 38,
		19, 0, 111, 113, 3, 72, 36, 0, 112, 107, 1, 0, 0, 0, 112, 108, 1, 0, 0,
		0, 112, 109, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 111, 1, 0, 0, 0, 113,
		116, 1, 0, 0, 0, 114, 112, 1, 0, 0, 0, 114, 115, 1, 0, 0, 0, 115, 117,
		1, 0, 0, 0, 116, 114, 1, 0, 0, 0, 117, 118, 5, 0, 0, 1, 118, 1, 1, 0, 0,
		0, 119, 120, 5, 1, 0, 0, 120, 121, 5, 37, 0, 0, 121, 122, 7, 0, 0, 0, 122,
		123, 5, 36, 0, 0, 123, 3, 1, 0, 0, 0, 124, 126, 5, 2, 0, 0, 125, 127, 7,
		1, 0, 0, 126, 125, 1, 0, 0, 0, 126, 127, 1, 0, 0, 0, 127, 128, 1, 0, 0,
		0, 128, 129, 3, 98, 49, 0, 129, 130, 5, 36, 0, 0, 130, 5, 1, 0, 0, 0, 131,
		132, 5, 5, 0, 0, 132, 133, 3, 76, 38, 0, 133, 134, 5, 36, 0, 0, 134, 7,
		1, 0, 0, 0, 135, 136, 5, 6, 0, 0, 136, 137, 3, 10, 5, 0, 137, 138, 5, 37,
		0, 0, 138, 139, 3, 68, 34, 0, 139, 140, 5, 36, 0, 0, 140, 9, 1, 0, 0, 0,
		141, 150, 3, 76, 38, 0, 142, 143, 5, 38, 0, 0, 143, 144, 3, 76, 38, 0,
		144, 147, 5, 39, 0, 0, 145, 146, 5, 46, 0, 0, 146, 148, 3, 76, 38, 0, 147,
		145, 1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 150, 1, 0, 0, 0, 149, 141,
		1, 0, 0, 0, 149, 142, 1, 0, 0, 0, 150, 11, 1, 0, 0, 0, 151, 153, 5, 7,
		0, 0, 152, 151, 1, 0, 0, 0, 152, 153, 1, 0, 0, 0, 153, 154, 1, 0, 0, 0,
		154, 155, 3, 28, 14, 0, 155, 156, 3, 82, 41, 0, 156, 157, 5, 37, 0, 0,
		157, 162, 3, 18, 9, 0, 158, 159, 5, 40, 0, 0, 159, 160, 3, 14, 7, 0, 160,
		161, 5, 41, 0, 0, 161, 163, 1, 0, 0, 0, 162, 158, 1, 0, 0, 0, 162, 163,
		1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 165, 5, 36, 0, 0, 165, 13, 1, 0,
		0, 0, 166, 171, 3, 16, 8, 0, 167, 168, 5, 47, 0, 0, 168, 170, 3, 16, 8,
		0, 169, 167, 1, 0, 0, 0, 170, 173, 1, 0, 0, 0, 171, 169, 1, 0, 0, 0, 171,
		172, 1, 0, 0, 0, 172, 15, 1, 0, 0, 0, 173, 171, 1, 0, 0, 0, 174, 175, 3,
		10, 5, 0, 175, 176, 5, 37, 0, 0, 176, 177, 3, 68, 34, 0, 177, 17, 1, 0,
		0, 0, 178, 179, 3, 96, 48, 0, 179, 19, 1, 0, 0, 0, 180, 181, 5, 8, 0, 0,
		181, 182, 3, 84, 42, 0, 182, 188, 5, 42, 0, 0, 183, 187, 3, 8, 4, 0, 184,
		187, 3, 22, 11, 0, 185, 187, 3, 72, 36, 0, 186, 183, 1, 0, 0, 0, 186, 184,
		1, 0, 0, 0, 186, 185, 1, 0, 0, 0, 187, 190, 1, 0, 0, 0, 188, 186, 1, 0,
		0, 0, 188, 189, 1, 0, 0, 0, 189, 191, 1, 0, 0, 0, 190, 188, 1, 0, 0, 0,
		191, 192, 5, 43, 0, 0, 192, 21, 1, 0, 0, 0, 193, 194, 3, 28, 14, 0, 194,
		195, 3, 82, 41, 0, 195, 196, 5, 37, 0, 0, 196, 201, 3, 18, 9, 0, 197, 198,
		5, 40, 0, 0, 198, 199, 3, 14, 7, 0, 199, 200, 5, 41, 0, 0, 200, 202, 1,
		0, 0, 0, 201, 197, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 203, 1, 0, 0,
		0, 203, 204, 5, 36, 0, 0, 204, 23, 1, 0, 0, 0, 205, 206, 5, 9, 0, 0, 206,
		207, 5, 44, 0, 0, 207, 208, 3, 26, 13, 0, 208, 209, 5, 47, 0, 0, 209, 210,
		3, 28, 14, 0, 210, 211, 5, 45, 0, 0, 211, 212, 3, 86, 43, 0, 212, 213,
		5, 37, 0, 0, 213, 218, 3, 18, 9, 0, 214, 215, 5, 40, 0, 0, 215, 216, 3,
		14, 7, 0, 216, 217, 5, 41, 0, 0, 217, 219, 1, 0, 0, 0, 218, 214, 1, 0,
		0, 0, 218, 219, 1, 0, 0, 0, 219, 220, 1, 0, 0, 0, 220, 221, 5, 36, 0, 0,
		221, 25, 1, 0, 0, 0, 222, 223, 7, 2, 0, 0, 223, 27, 1, 0, 0, 0, 224, 242,
		5, 22, 0, 0, 225, 242, 5, 23, 0, 0, 226, 242, 5, 10, 0, 0, 227, 242, 5,
		11, 0, 0, 228, 242, 5, 12, 0, 0, 229, 242, 5, 13, 0, 0, 230, 242, 5, 14,
		0, 0, 231, 242, 5, 15, 0, 0, 232, 242, 5, 16, 0, 0, 233, 242, 5, 17, 0,
		0, 234, 242, 5, 18, 0, 0, 235, 242, 5, 19, 0, 0, 236, 242, 5, 20, 0, 0,
		237, 242, 5, 21, 0, 0, 238, 242, 5, 24, 0, 0, 239, 242, 3, 92, 46, 0, 240,
		242, 3, 94, 47, 0, 241, 224, 1, 0, 0, 0, 241, 225, 1, 0, 0, 0, 241, 226,
		1, 0, 0, 0, 241, 227, 1, 0, 0, 0, 241, 228, 1, 0, 0, 0, 241, 229, 1, 0,
		0, 0, 241, 230, 1, 0, 0, 0, 241, 231, 1, 0, 0, 0, 241, 232, 1, 0, 0, 0,
		241, 233, 1, 0, 0, 0, 241, 234, 1, 0, 0, 0, 241, 235, 1, 0, 0, 0, 241,
		236, 1, 0, 0, 0, 241, 237, 1, 0, 0, 0, 241, 238, 1, 0, 0, 0, 241, 239,
		1, 0, 0, 0, 241, 240, 1, 0, 0, 0, 242, 29, 1, 0, 0, 0, 243, 246, 5, 25,
		0, 0, 244, 247, 3, 32, 16, 0, 245, 247, 3, 36, 18, 0, 246, 244, 1, 0, 0,
		0, 246, 245, 1, 0, 0, 0, 247, 248, 1, 0, 0, 0, 248, 249, 5, 36, 0, 0, 249,
		31, 1, 0, 0, 0, 250, 255, 3, 34, 17, 0, 251, 252, 5, 47, 0, 0, 252, 254,
		3, 34, 17, 0, 253, 251, 1, 0, 0, 0, 254, 257, 1, 0, 0, 0, 255, 253, 1,
		0, 0, 0, 255, 256, 1, 0, 0, 0, 256, 33, 1, 0, 0, 0, 257, 255, 1, 0, 0,
		0, 258, 264, 3, 96, 48, 0, 259, 262, 5, 26, 0, 0, 260, 263, 3, 96, 48,
		0, 261, 263, 5, 27, 0, 0, 262, 260, 1, 0, 0, 0, 262, 261, 1, 0, 0, 0, 263,
		265, 1, 0, 0, 0, 264, 259, 1, 0, 0, 0, 264, 265, 1, 0, 0, 0, 265, 35, 1,
		0, 0, 0, 266, 271, 3, 98, 49, 0, 267, 268, 5, 47, 0, 0, 268, 270, 3, 98,
		49, 0, 269, 267, 1, 0, 0, 0, 270, 273, 1, 0, 0, 0, 271, 269, 1, 0, 0, 0,
		271, 272, 1, 0, 0, 0, 272, 37, 1, 0, 0, 0, 273, 271, 1, 0, 0, 0, 274, 278,
		3, 52, 26, 0, 275, 278, 3, 40, 20, 0, 276, 278, 3, 58, 29, 0, 277, 274,
		1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 277, 276, 1, 0, 0, 0, 278, 39, 1, 0,
		0, 0, 279, 280, 5, 28, 0, 0, 280, 281, 3, 80, 40, 0, 281, 282, 3, 42, 21,
		0, 282, 41, 1, 0, 0, 0, 283, 287, 5, 42, 0, 0, 284, 286, 3, 44, 22, 0,
		285, 284, 1, 0, 0, 0, 286, 289, 1, 0, 0, 0, 287, 285, 1, 0, 0, 0, 287,
		288, 1, 0, 0, 0, 288, 290, 1, 0, 0, 0, 289, 287, 1, 0, 0, 0, 290, 291,
		5, 43, 0, 0, 291, 43, 1, 0, 0, 0, 292, 296, 3, 8, 4, 0, 293, 296, 3, 46,
		23, 0, 294, 296, 3, 72, 36, 0, 295, 292, 1, 0, 0, 0, 295, 293, 1, 0, 0,
		0, 295, 294, 1, 0, 0, 0, 296, 45, 1, 0, 0, 0, 297, 298, 3, 74, 37, 0, 298,
		300, 5, 37, 0, 0, 299, 301, 5, 50, 0, 0, 300, 299, 1, 0, 0, 0, 300, 301,
		1, 0, 0, 0, 301, 302, 1, 0, 0, 0, 302, 304, 3, 96, 48, 0, 303, 305, 3,
		48, 24, 0, 304, 303, 1, 0, 0, 0, 304, 305, 1, 0, 0, 0, 305, 306, 1, 0,
		0, 0, 306, 307, 5, 36, 0, 0, 307, 47, 1, 0, 0, 0, 308, 309, 5, 40, 0, 0,
		309, 314, 3, 50, 25, 0, 310, 311, 5, 47, 0, 0, 311, 313, 3, 50, 25, 0,
		312, 310, 1, 0, 0, 0, 313, 316, 1, 0, 0, 0, 314, 312, 1, 0, 0, 0, 314,
		315, 1, 0, 0, 0, 315, 317, 1, 0, 0, 0, 316, 314, 1, 0, 0, 0, 317, 318,
		5, 41, 0, 0, 318, 49, 1, 0, 0, 0, 319, 320, 3, 10, 5, 0, 320, 321, 5, 37,
		0, 0, 321, 322, 3, 68, 34, 0, 322, 51, 1, 0, 0, 0, 323, 324, 5, 29, 0,
		0, 324, 325, 3, 78, 39, 0, 325, 326, 3, 54, 27, 0, 326, 53, 1, 0, 0, 0,
		327, 331, 5, 42, 0, 0, 328, 330, 3, 56, 28, 0, 329, 328, 1, 0, 0, 0, 330,
		333, 1, 0, 0, 0, 331, 329, 1, 0, 0, 0, 331, 332, 1, 0, 0, 0, 332, 334,
		1, 0, 0, 0, 333, 331, 1, 0, 0, 0, 334, 335, 5, 43, 0, 0, 335, 55, 1, 0,
		0, 0, 336, 345, 3, 12, 6, 0, 337, 345, 3, 40, 20, 0, 338, 345, 3, 52, 26,
		0, 339, 345, 3, 8, 4, 0, 340, 345, 3, 20, 10, 0, 341, 345, 3, 24, 12, 0,
		342, 345, 3, 30, 15, 0, 343, 345, 3, 72, 36, 0, 344, 336, 1, 0, 0, 0, 344,
		337, 1, 0, 0, 0, 344, 338, 1, 0, 0, 0, 344, 339, 1, 0, 0, 0, 344, 340,
		1, 0, 0, 0, 344, 341, 1, 0, 0, 0, 344, 342, 1, 0, 0, 0, 344, 343, 1, 0,
		0, 0, 345, 57, 1, 0, 0, 0, 346, 347, 5, 30, 0, 0, 347, 348, 3, 88, 44,
		0, 348, 352, 5, 42, 0, 0, 349, 351, 3, 60, 30, 0, 350, 349, 1, 0, 0, 0,
		351, 354, 1, 0, 0, 0, 352, 350, 1, 0, 0, 0, 352, 353, 1, 0, 0, 0, 353,
		355, 1, 0, 0, 0, 354, 352, 1, 0, 0, 0, 355, 356, 5, 43, 0, 0, 356, 59,
		1, 0, 0, 0, 357, 361, 3, 8, 4, 0, 358, 361, 3, 62, 31, 0, 359, 361, 3,
		72, 36, 0, 360, 357, 1, 0, 0, 0, 360, 358, 1, 0, 0, 0, 360, 359, 1, 0,
		0, 0, 361, 61, 1, 0, 0, 0, 362, 363, 5, 31, 0, 0, 363, 364, 3, 90, 45,
		0, 364, 366, 5, 38, 0, 0, 365, 367, 3, 64, 32, 0, 366, 365, 1, 0, 0, 0,
		366, 367, 1, 0, 0, 0, 367, 368, 1, 0, 0, 0, 368, 369, 3, 92, 46, 0, 369,
		370, 5, 39, 0, 0, 370, 371, 5, 33, 0, 0, 371, 373, 5, 38, 0, 0, 372, 374,
		3, 66, 33, 0, 373, 372, 1, 0, 0, 0, 373, 374, 1, 0, 0, 0, 374, 375, 1,
		0, 0, 0, 375, 376, 3, 92, 46, 0, 376, 387, 5, 39, 0, 0, 377, 382, 5, 42,
		0, 0, 378, 381, 3, 8, 4, 0, 379, 381, 3, 72, 36, 0, 380, 378, 1, 0, 0,
		0, 380, 379, 1, 0, 0, 0, 381, 384, 1, 0, 0, 0, 382, 380, 1, 0, 0, 0, 382,
		383, 1, 0, 0, 0, 383, 385, 1, 0, 0, 0, 384, 382, 1, 0, 0, 0, 385, 388,
		5, 43, 0, 0, 386, 388, 5, 36, 0, 0, 387, 377, 1, 0, 0, 0, 387, 386, 1,
		0, 0, 0, 388, 63, 1, 0, 0, 0, 389, 390, 5, 32, 0, 0, 390, 65, 1, 0, 0,
		0, 391, 392, 5, 32, 0, 0, 392, 67, 1, 0, 0, 0, 393, 406, 3, 76, 38, 0,
		394, 396, 7, 3, 0, 0, 395, 394, 1, 0, 0, 0, 395, 396, 1, 0, 0, 0, 396,
		397, 1, 0, 0, 0, 397, 406, 3, 96, 48, 0, 398, 400, 7, 3, 0, 0, 399, 398,
		1, 0, 0, 0, 399, 400, 1, 0, 0, 0, 400, 401, 1, 0, 0, 0, 401, 406, 3, 102,
		51, 0, 402, 406, 3, 98, 49, 0, 403, 406, 3, 100, 50, 0, 404, 406, 3, 70,
		35, 0, 405, 393, 1, 0, 0, 0, 405, 395, 1, 0, 0, 0, 405, 399, 1, 0, 0, 0,
		405, 402, 1, 0, 0, 0, 405, 403, 1, 0, 0, 0, 405, 404, 1, 0, 0, 0, 406,
		69, 1, 0, 0, 0, 407, 414, 5, 42, 0, 0, 408, 409, 3, 74, 37, 0, 409, 410,
		5, 48, 0, 0, 410, 411, 3, 68, 34, 0, 411, 413, 1, 0, 0, 0, 412, 408, 1,
		0, 0, 0, 413, 416, 1, 0, 0, 0, 414, 412, 1, 0, 0, 0, 414, 415, 1, 0, 0,
		0, 415, 417, 1, 0, 0, 0, 416, 414, 1, 0, 0, 0, 417, 418, 5, 43, 0, 0, 418,
		71, 1, 0, 0, 0, 419, 420, 5, 36, 0, 0, 420, 73, 1, 0, 0, 0, 421, 424, 5,
		55, 0, 0, 422, 424, 3, 104, 52, 0, 423, 421, 1, 0, 0, 0, 423, 422, 1, 0,
		0, 0, 424, 75, 1, 0, 0, 0, 425, 430, 3, 74, 37, 0, 426, 427, 5, 46, 0,
		0, 427, 429, 3, 74, 37, 0, 428, 426, 1, 0, 0, 0, 429, 432, 1, 0, 0, 0,
		430, 428, 1, 0, 0, 0, 430, 431, 1, 0, 0, 0, 431, 77, 1, 0, 0, 0, 432, 430,
		1, 0, 0, 0, 433, 434, 3, 74, 37, 0, 434, 79, 1, 0, 0, 0, 435, 436, 3, 74,
		37, 0, 436, 81, 1, 0, 0, 0, 437, 438, 3, 74, 37, 0, 438, 83, 1, 0, 0, 0,
		439, 440, 3, 74, 37, 0, 440, 85, 1, 0, 0, 0, 441, 442, 3, 74, 37, 0, 442,
		87, 1, 0, 0, 0, 443, 444, 3, 74, 37, 0, 444, 89, 1, 0, 0, 0, 445, 446,
		3, 74, 37, 0, 446, 91, 1, 0, 0, 0, 447, 449, 5, 46, 0, 0, 448, 447, 1,
		0, 0, 0, 448, 449, 1, 0, 0, 0, 449, 455, 1, 0, 0, 0, 450, 451, 3, 74, 37,
		0, 451, 452, 5, 46, 0, 0, 452, 454, 1, 0, 0, 0, 453, 450, 1, 0, 0, 0, 454,
		457, 1, 0, 0, 0, 455, 453, 1, 0, 0, 0, 455, 456, 1, 0, 0, 0, 456, 458,
		1, 0, 0, 0, 457, 455, 1, 0, 0, 0, 458, 459, 3, 78, 39, 0, 459, 93, 1, 0,
		0, 0, 460, 462, 5, 46, 0, 0, 461, 460, 1, 0, 0, 0, 461, 462, 1, 0, 0, 0,
		462, 468, 1, 0, 0, 0, 463, 464, 3, 74, 37, 0, 464, 465, 5, 46, 0, 0, 465,
		467, 1, 0, 0, 0, 466, 463, 1, 0, 0, 0, 467, 470, 1, 0, 0, 0, 468, 466,
		1, 0, 0, 0, 468, 469, 1, 0, 0, 0, 469, 471, 1, 0, 0, 0, 470, 468, 1, 0,
		0, 0, 471, 472, 3, 80, 40, 0, 472, 95, 1, 0, 0, 0, 473, 474, 5, 54, 0,
		0, 474, 97, 1, 0, 0, 0, 475, 476, 7, 4, 0, 0, 476, 99, 1, 0, 0, 0, 477,
		478, 5, 52, 0, 0, 478, 101, 1, 0, 0, 0, 479, 480, 5, 53, 0, 0, 480, 103,
		1, 0, 0, 0, 481, 482, 7, 5, 0, 0, 482, 105, 1, 0, 0, 0, 43, 112, 114, 126,
		147, 149, 152, 162, 171, 186, 188, 201, 218, 241, 246, 255, 262, 264, 271,
		277, 287, 295, 300, 304, 314, 331, 344, 352, 360, 366, 373, 380, 382, 387,
		395, 399, 405, 414, 423, 430, 448, 455, 461, 468,
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

// Protobuf3ParserInit initializes any static state used to implement Protobuf3Parser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewProtobuf3Parser(). You can call this function if you wish to initialize the static state ahead
// of time.
func Protobuf3ParserInit() {
	staticData := &Protobuf3ParserStaticData
	staticData.once.Do(protobuf3ParserInit)
}

// NewProtobuf3Parser produces a new parser instance for the optional input antlr.TokenStream.
func NewProtobuf3Parser(input antlr.TokenStream) *Protobuf3Parser {
	Protobuf3ParserInit()
	this := new(Protobuf3Parser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &Protobuf3ParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Protobuf3.g4"

	return this
}

// Protobuf3Parser tokens.
const (
	Protobuf3ParserEOF               = antlr.TokenEOF
	Protobuf3ParserSYNTAX            = 1
	Protobuf3ParserIMPORT            = 2
	Protobuf3ParserWEAK              = 3
	Protobuf3ParserPUBLIC            = 4
	Protobuf3ParserPACKAGE           = 5
	Protobuf3ParserOPTION            = 6
	Protobuf3ParserREPEATED          = 7
	Protobuf3ParserONEOF             = 8
	Protobuf3ParserMAP               = 9
	Protobuf3ParserINT32             = 10
	Protobuf3ParserINT64             = 11
	Protobuf3ParserUINT32            = 12
	Protobuf3ParserUINT64            = 13
	Protobuf3ParserSINT32            = 14
	Protobuf3ParserSINT64            = 15
	Protobuf3ParserFIXED32           = 16
	Protobuf3ParserFIXED64           = 17
	Protobuf3ParserSFIXED32          = 18
	Protobuf3ParserSFIXED64          = 19
	Protobuf3ParserBOOL              = 20
	Protobuf3ParserSTRING            = 21
	Protobuf3ParserDOUBLE            = 22
	Protobuf3ParserFLOAT             = 23
	Protobuf3ParserBYTES             = 24
	Protobuf3ParserRESERVED          = 25
	Protobuf3ParserTO                = 26
	Protobuf3ParserMAX               = 27
	Protobuf3ParserENUM              = 28
	Protobuf3ParserMESSAGE           = 29
	Protobuf3ParserSERVICE           = 30
	Protobuf3ParserRPC               = 31
	Protobuf3ParserSTREAM            = 32
	Protobuf3ParserRETURNS           = 33
	Protobuf3ParserPROTO3_LIT_DOBULE = 34
	Protobuf3ParserPROTO3_LIT_SINGLE = 35
	Protobuf3ParserSEMI              = 36
	Protobuf3ParserEQ                = 37
	Protobuf3ParserLP                = 38
	Protobuf3ParserRP                = 39
	Protobuf3ParserLB                = 40
	Protobuf3ParserRB                = 41
	Protobuf3ParserLC                = 42
	Protobuf3ParserRC                = 43
	Protobuf3ParserLT                = 44
	Protobuf3ParserGT                = 45
	Protobuf3ParserDOT               = 46
	Protobuf3ParserCOMMA             = 47
	Protobuf3ParserCOLON             = 48
	Protobuf3ParserPLUS              = 49
	Protobuf3ParserMINUS             = 50
	Protobuf3ParserSTR_LIT           = 51
	Protobuf3ParserBOOL_LIT          = 52
	Protobuf3ParserFLOAT_LIT         = 53
	Protobuf3ParserINT_LIT           = 54
	Protobuf3ParserIDENTIFIER        = 55
	Protobuf3ParserWS                = 56
	Protobuf3ParserLINE_COMMENT      = 57
	Protobuf3ParserCOMMENT           = 58
)

// Protobuf3Parser rules.
const (
	Protobuf3ParserRULE_proto              = 0
	Protobuf3ParserRULE_syntax             = 1
	Protobuf3ParserRULE_importStatement    = 2
	Protobuf3ParserRULE_packageStatement   = 3
	Protobuf3ParserRULE_optionStatement    = 4
	Protobuf3ParserRULE_optionName         = 5
	Protobuf3ParserRULE_field              = 6
	Protobuf3ParserRULE_fieldOptions       = 7
	Protobuf3ParserRULE_fieldOption        = 8
	Protobuf3ParserRULE_fieldNumber        = 9
	Protobuf3ParserRULE_oneof              = 10
	Protobuf3ParserRULE_oneofField         = 11
	Protobuf3ParserRULE_mapField           = 12
	Protobuf3ParserRULE_keyType            = 13
	Protobuf3ParserRULE_type_              = 14
	Protobuf3ParserRULE_reserved           = 15
	Protobuf3ParserRULE_ranges             = 16
	Protobuf3ParserRULE_range_             = 17
	Protobuf3ParserRULE_reservedFieldNames = 18
	Protobuf3ParserRULE_topLevelDef        = 19
	Protobuf3ParserRULE_enumDef            = 20
	Protobuf3ParserRULE_enumBody           = 21
	Protobuf3ParserRULE_enumElement        = 22
	Protobuf3ParserRULE_enumField          = 23
	Protobuf3ParserRULE_enumValueOptions   = 24
	Protobuf3ParserRULE_enumValueOption    = 25
	Protobuf3ParserRULE_messageDef         = 26
	Protobuf3ParserRULE_messageBody        = 27
	Protobuf3ParserRULE_messageElement     = 28
	Protobuf3ParserRULE_serviceDef         = 29
	Protobuf3ParserRULE_serviceElement     = 30
	Protobuf3ParserRULE_rpc                = 31
	Protobuf3ParserRULE_requestStream      = 32
	Protobuf3ParserRULE_responseStream     = 33
	Protobuf3ParserRULE_constant           = 34
	Protobuf3ParserRULE_blockLit           = 35
	Protobuf3ParserRULE_emptyStatement_    = 36
	Protobuf3ParserRULE_ident              = 37
	Protobuf3ParserRULE_fullIdent          = 38
	Protobuf3ParserRULE_messageName        = 39
	Protobuf3ParserRULE_enumName           = 40
	Protobuf3ParserRULE_fieldName          = 41
	Protobuf3ParserRULE_oneofName          = 42
	Protobuf3ParserRULE_mapName            = 43
	Protobuf3ParserRULE_serviceName        = 44
	Protobuf3ParserRULE_rpcName            = 45
	Protobuf3ParserRULE_messageType        = 46
	Protobuf3ParserRULE_enumType           = 47
	Protobuf3ParserRULE_intLit             = 48
	Protobuf3ParserRULE_strLit             = 49
	Protobuf3ParserRULE_boolLit            = 50
	Protobuf3ParserRULE_floatLit           = 51
	Protobuf3ParserRULE_keywords           = 52
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
	p.RuleIndex = Protobuf3ParserRULE_proto
	return p
}

func InitEmptyProtoContext(p *ProtoContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_proto
}

func (*ProtoContext) IsProtoContext() {}

func NewProtoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProtoContext {
	var p = new(ProtoContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_proto

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
	return s.GetToken(Protobuf3ParserEOF, 0)
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
	case Protobuf3Visitor:
		return t.VisitProto(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Proto() (localctx IProtoContext) {
	localctx = NewProtoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Protobuf3ParserRULE_proto)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Syntax()
	}
	p.SetState(114)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&70598525028) != 0 {
		p.SetState(112)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case Protobuf3ParserIMPORT:
			{
				p.SetState(107)
				p.ImportStatement()
			}

		case Protobuf3ParserPACKAGE:
			{
				p.SetState(108)
				p.PackageStatement()
			}

		case Protobuf3ParserOPTION:
			{
				p.SetState(109)
				p.OptionStatement()
			}

		case Protobuf3ParserENUM, Protobuf3ParserMESSAGE, Protobuf3ParserSERVICE:
			{
				p.SetState(110)
				p.TopLevelDef()
			}

		case Protobuf3ParserSEMI:
			{
				p.SetState(111)
				p.EmptyStatement_()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}

		p.SetState(116)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(117)
		p.Match(Protobuf3ParserEOF)
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
	PROTO3_LIT_SINGLE() antlr.TerminalNode
	PROTO3_LIT_DOBULE() antlr.TerminalNode

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
	p.RuleIndex = Protobuf3ParserRULE_syntax
	return p
}

func InitEmptySyntaxContext(p *SyntaxContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_syntax
}

func (*SyntaxContext) IsSyntaxContext() {}

func NewSyntaxContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SyntaxContext {
	var p = new(SyntaxContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_syntax

	return p
}

func (s *SyntaxContext) GetParser() antlr.Parser { return s.parser }

func (s *SyntaxContext) SYNTAX() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSYNTAX, 0)
}

func (s *SyntaxContext) EQ() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserEQ, 0)
}

func (s *SyntaxContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *SyntaxContext) PROTO3_LIT_SINGLE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPROTO3_LIT_SINGLE, 0)
}

func (s *SyntaxContext) PROTO3_LIT_DOBULE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPROTO3_LIT_DOBULE, 0)
}

func (s *SyntaxContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SyntaxContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SyntaxContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitSyntax(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Syntax() (localctx ISyntaxContext) {
	localctx = NewSyntaxContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Protobuf3ParserRULE_syntax)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(119)
		p.Match(Protobuf3ParserSYNTAX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(120)
		p.Match(Protobuf3ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(121)
		_la = p.GetTokenStream().LA(1)

		if !(_la == Protobuf3ParserPROTO3_LIT_DOBULE || _la == Protobuf3ParserPROTO3_LIT_SINGLE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(122)
		p.Match(Protobuf3ParserSEMI)
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
	p.RuleIndex = Protobuf3ParserRULE_importStatement
	return p
}

func InitEmptyImportStatementContext(p *ImportStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_importStatement
}

func (*ImportStatementContext) IsImportStatementContext() {}

func NewImportStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStatementContext {
	var p = new(ImportStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_importStatement

	return p
}

func (s *ImportStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStatementContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIMPORT, 0)
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
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *ImportStatementContext) WEAK() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserWEAK, 0)
}

func (s *ImportStatementContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPUBLIC, 0)
}

func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitImportStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) ImportStatement() (localctx IImportStatementContext) {
	localctx = NewImportStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Protobuf3ParserRULE_importStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(Protobuf3ParserIMPORT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserWEAK || _la == Protobuf3ParserPUBLIC {
		{
			p.SetState(125)
			_la = p.GetTokenStream().LA(1)

			if !(_la == Protobuf3ParserWEAK || _la == Protobuf3ParserPUBLIC) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(128)
		p.StrLit()
	}
	{
		p.SetState(129)
		p.Match(Protobuf3ParserSEMI)
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
	p.RuleIndex = Protobuf3ParserRULE_packageStatement
	return p
}

func InitEmptyPackageStatementContext(p *PackageStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_packageStatement
}

func (*PackageStatementContext) IsPackageStatementContext() {}

func NewPackageStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PackageStatementContext {
	var p = new(PackageStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_packageStatement

	return p
}

func (s *PackageStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *PackageStatementContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPACKAGE, 0)
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
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *PackageStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PackageStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PackageStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitPackageStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) PackageStatement() (localctx IPackageStatementContext) {
	localctx = NewPackageStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Protobuf3ParserRULE_packageStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(131)
		p.Match(Protobuf3ParserPACKAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(132)
		p.FullIdent()
	}
	{
		p.SetState(133)
		p.Match(Protobuf3ParserSEMI)
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
	p.RuleIndex = Protobuf3ParserRULE_optionStatement
	return p
}

func InitEmptyOptionStatementContext(p *OptionStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_optionStatement
}

func (*OptionStatementContext) IsOptionStatementContext() {}

func NewOptionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionStatementContext {
	var p = new(OptionStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_optionStatement

	return p
}

func (s *OptionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionStatementContext) OPTION() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserOPTION, 0)
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
	return s.GetToken(Protobuf3ParserEQ, 0)
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
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *OptionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitOptionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) OptionStatement() (localctx IOptionStatementContext) {
	localctx = NewOptionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Protobuf3ParserRULE_optionStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(Protobuf3ParserOPTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(136)
		p.OptionName()
	}
	{
		p.SetState(137)
		p.Match(Protobuf3ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(138)
		p.Constant()
	}
	{
		p.SetState(139)
		p.Match(Protobuf3ParserSEMI)
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
	p.RuleIndex = Protobuf3ParserRULE_optionName
	return p
}

func InitEmptyOptionNameContext(p *OptionNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_optionName
}

func (*OptionNameContext) IsOptionNameContext() {}

func NewOptionNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionNameContext {
	var p = new(OptionNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_optionName

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

func (s *OptionNameContext) LP() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLP, 0)
}

func (s *OptionNameContext) RP() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRP, 0)
}

func (s *OptionNameContext) DOT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserDOT, 0)
}

func (s *OptionNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitOptionName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) OptionName() (localctx IOptionNameContext) {
	localctx = NewOptionNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, Protobuf3ParserRULE_optionName)
	var _la int

	p.SetState(149)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserSYNTAX, Protobuf3ParserIMPORT, Protobuf3ParserWEAK, Protobuf3ParserPUBLIC, Protobuf3ParserPACKAGE, Protobuf3ParserOPTION, Protobuf3ParserREPEATED, Protobuf3ParserONEOF, Protobuf3ParserMAP, Protobuf3ParserINT32, Protobuf3ParserINT64, Protobuf3ParserUINT32, Protobuf3ParserUINT64, Protobuf3ParserSINT32, Protobuf3ParserSINT64, Protobuf3ParserFIXED32, Protobuf3ParserFIXED64, Protobuf3ParserSFIXED32, Protobuf3ParserSFIXED64, Protobuf3ParserBOOL, Protobuf3ParserSTRING, Protobuf3ParserDOUBLE, Protobuf3ParserFLOAT, Protobuf3ParserBYTES, Protobuf3ParserRESERVED, Protobuf3ParserTO, Protobuf3ParserMAX, Protobuf3ParserENUM, Protobuf3ParserMESSAGE, Protobuf3ParserSERVICE, Protobuf3ParserRPC, Protobuf3ParserSTREAM, Protobuf3ParserRETURNS, Protobuf3ParserBOOL_LIT, Protobuf3ParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(141)
			p.FullIdent()
		}

	case Protobuf3ParserLP:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(142)
			p.Match(Protobuf3ParserLP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(143)
			p.FullIdent()
		}
		{
			p.SetState(144)
			p.Match(Protobuf3ParserRP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(147)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Protobuf3ParserDOT {
			{
				p.SetState(145)
				p.Match(Protobuf3ParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(146)
				p.FullIdent()
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

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Type_() IType_Context
	FieldName() IFieldNameContext
	EQ() antlr.TerminalNode
	FieldNumber() IFieldNumberContext
	SEMI() antlr.TerminalNode
	REPEATED() antlr.TerminalNode
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
	p.RuleIndex = Protobuf3ParserRULE_field
	return p
}

func InitEmptyFieldContext(p *FieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_field
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

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
	return s.GetToken(Protobuf3ParserEQ, 0)
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
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *FieldContext) REPEATED() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserREPEATED, 0)
}

func (s *FieldContext) LB() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLB, 0)
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
	return s.GetToken(Protobuf3ParserRB, 0)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, Protobuf3ParserRULE_field)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(152)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(151)
			p.Match(Protobuf3ParserREPEATED)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(154)
		p.Type_()
	}
	{
		p.SetState(155)
		p.FieldName()
	}
	{
		p.SetState(156)
		p.Match(Protobuf3ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(157)
		p.FieldNumber()
	}
	p.SetState(162)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserLB {
		{
			p.SetState(158)
			p.Match(Protobuf3ParserLB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(159)
			p.FieldOptions()
		}
		{
			p.SetState(160)
			p.Match(Protobuf3ParserRB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(164)
		p.Match(Protobuf3ParserSEMI)
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
	p.RuleIndex = Protobuf3ParserRULE_fieldOptions
	return p
}

func InitEmptyFieldOptionsContext(p *FieldOptionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldOptions
}

func (*FieldOptionsContext) IsFieldOptionsContext() {}

func NewFieldOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptionsContext {
	var p = new(FieldOptionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fieldOptions

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
	return s.GetTokens(Protobuf3ParserCOMMA)
}

func (s *FieldOptionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserCOMMA, i)
}

func (s *FieldOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldOptionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitFieldOptions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) FieldOptions() (localctx IFieldOptionsContext) {
	localctx = NewFieldOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, Protobuf3ParserRULE_fieldOptions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.FieldOption()
	}
	p.SetState(171)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserCOMMA {
		{
			p.SetState(167)
			p.Match(Protobuf3ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(168)
			p.FieldOption()
		}

		p.SetState(173)
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
	p.RuleIndex = Protobuf3ParserRULE_fieldOption
	return p
}

func InitEmptyFieldOptionContext(p *FieldOptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldOption
}

func (*FieldOptionContext) IsFieldOptionContext() {}

func NewFieldOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldOptionContext {
	var p = new(FieldOptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fieldOption

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
	return s.GetToken(Protobuf3ParserEQ, 0)
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
	case Protobuf3Visitor:
		return t.VisitFieldOption(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) FieldOption() (localctx IFieldOptionContext) {
	localctx = NewFieldOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, Protobuf3ParserRULE_fieldOption)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.OptionName()
	}
	{
		p.SetState(175)
		p.Match(Protobuf3ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
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
	p.RuleIndex = Protobuf3ParserRULE_fieldNumber
	return p
}

func InitEmptyFieldNumberContext(p *FieldNumberContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldNumber
}

func (*FieldNumberContext) IsFieldNumberContext() {}

func NewFieldNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNumberContext {
	var p = new(FieldNumberContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fieldNumber

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
	case Protobuf3Visitor:
		return t.VisitFieldNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) FieldNumber() (localctx IFieldNumberContext) {
	localctx = NewFieldNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, Protobuf3ParserRULE_fieldNumber)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(178)
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
	p.RuleIndex = Protobuf3ParserRULE_oneof
	return p
}

func InitEmptyOneofContext(p *OneofContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_oneof
}

func (*OneofContext) IsOneofContext() {}

func NewOneofContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofContext {
	var p = new(OneofContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_oneof

	return p
}

func (s *OneofContext) GetParser() antlr.Parser { return s.parser }

func (s *OneofContext) ONEOF() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserONEOF, 0)
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
	return s.GetToken(Protobuf3ParserLC, 0)
}

func (s *OneofContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRC, 0)
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
	case Protobuf3Visitor:
		return t.VisitOneof(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Oneof() (localctx IOneofContext) {
	localctx = NewOneofContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, Protobuf3ParserRULE_oneof)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(180)
		p.Match(Protobuf3ParserONEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(181)
		p.OneofName()
	}
	{
		p.SetState(182)
		p.Match(Protobuf3ParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&40602851289858046) != 0 {
		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(183)
				p.OptionStatement()
			}

		case 2:
			{
				p.SetState(184)
				p.OneofField()
			}

		case 3:
			{
				p.SetState(185)
				p.EmptyStatement_()
			}

		case antlr.ATNInvalidAltNumber:
			goto errorExit
		}

		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(191)
		p.Match(Protobuf3ParserRC)
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
	p.RuleIndex = Protobuf3ParserRULE_oneofField
	return p
}

func InitEmptyOneofFieldContext(p *OneofFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_oneofField
}

func (*OneofFieldContext) IsOneofFieldContext() {}

func NewOneofFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofFieldContext {
	var p = new(OneofFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_oneofField

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
	return s.GetToken(Protobuf3ParserEQ, 0)
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
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *OneofFieldContext) LB() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLB, 0)
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
	return s.GetToken(Protobuf3ParserRB, 0)
}

func (s *OneofFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OneofFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OneofFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitOneofField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) OneofField() (localctx IOneofFieldContext) {
	localctx = NewOneofFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, Protobuf3ParserRULE_oneofField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(193)
		p.Type_()
	}
	{
		p.SetState(194)
		p.FieldName()
	}
	{
		p.SetState(195)
		p.Match(Protobuf3ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(196)
		p.FieldNumber()
	}
	p.SetState(201)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserLB {
		{
			p.SetState(197)
			p.Match(Protobuf3ParserLB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(198)
			p.FieldOptions()
		}
		{
			p.SetState(199)
			p.Match(Protobuf3ParserRB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(203)
		p.Match(Protobuf3ParserSEMI)
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
	p.RuleIndex = Protobuf3ParserRULE_mapField
	return p
}

func InitEmptyMapFieldContext(p *MapFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_mapField
}

func (*MapFieldContext) IsMapFieldContext() {}

func NewMapFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapFieldContext {
	var p = new(MapFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_mapField

	return p
}

func (s *MapFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *MapFieldContext) MAP() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMAP, 0)
}

func (s *MapFieldContext) LT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLT, 0)
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
	return s.GetToken(Protobuf3ParserCOMMA, 0)
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
	return s.GetToken(Protobuf3ParserGT, 0)
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
	return s.GetToken(Protobuf3ParserEQ, 0)
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
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *MapFieldContext) LB() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLB, 0)
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
	return s.GetToken(Protobuf3ParserRB, 0)
}

func (s *MapFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MapFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MapFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitMapField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) MapField() (localctx IMapFieldContext) {
	localctx = NewMapFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, Protobuf3ParserRULE_mapField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(205)
		p.Match(Protobuf3ParserMAP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(206)
		p.Match(Protobuf3ParserLT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(207)
		p.KeyType()
	}
	{
		p.SetState(208)
		p.Match(Protobuf3ParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(209)
		p.Type_()
	}
	{
		p.SetState(210)
		p.Match(Protobuf3ParserGT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(211)
		p.MapName()
	}
	{
		p.SetState(212)
		p.Match(Protobuf3ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.FieldNumber()
	}
	p.SetState(218)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserLB {
		{
			p.SetState(214)
			p.Match(Protobuf3ParserLB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(215)
			p.FieldOptions()
		}
		{
			p.SetState(216)
			p.Match(Protobuf3ParserRB)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(220)
		p.Match(Protobuf3ParserSEMI)
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
	p.RuleIndex = Protobuf3ParserRULE_keyType
	return p
}

func InitEmptyKeyTypeContext(p *KeyTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_keyType
}

func (*KeyTypeContext) IsKeyTypeContext() {}

func NewKeyTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeyTypeContext {
	var p = new(KeyTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_keyType

	return p
}

func (s *KeyTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *KeyTypeContext) INT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT32, 0)
}

func (s *KeyTypeContext) INT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT64, 0)
}

func (s *KeyTypeContext) UINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserUINT32, 0)
}

func (s *KeyTypeContext) UINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserUINT64, 0)
}

func (s *KeyTypeContext) SINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSINT32, 0)
}

func (s *KeyTypeContext) SINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSINT64, 0)
}

func (s *KeyTypeContext) FIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFIXED32, 0)
}

func (s *KeyTypeContext) FIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFIXED64, 0)
}

func (s *KeyTypeContext) SFIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSFIXED32, 0)
}

func (s *KeyTypeContext) SFIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSFIXED64, 0)
}

func (s *KeyTypeContext) BOOL() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBOOL, 0)
}

func (s *KeyTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTRING, 0)
}

func (s *KeyTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeyTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitKeyType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) KeyType() (localctx IKeyTypeContext) {
	localctx = NewKeyTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, Protobuf3ParserRULE_keyType)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(222)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4193280) != 0) {
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
	p.RuleIndex = Protobuf3ParserRULE_type_
	return p
}

func InitEmptyType_Context(p *Type_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_type_
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) DOUBLE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserDOUBLE, 0)
}

func (s *Type_Context) FLOAT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFLOAT, 0)
}

func (s *Type_Context) INT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT32, 0)
}

func (s *Type_Context) INT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT64, 0)
}

func (s *Type_Context) UINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserUINT32, 0)
}

func (s *Type_Context) UINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserUINT64, 0)
}

func (s *Type_Context) SINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSINT32, 0)
}

func (s *Type_Context) SINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSINT64, 0)
}

func (s *Type_Context) FIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFIXED32, 0)
}

func (s *Type_Context) FIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFIXED64, 0)
}

func (s *Type_Context) SFIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSFIXED32, 0)
}

func (s *Type_Context) SFIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSFIXED64, 0)
}

func (s *Type_Context) BOOL() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBOOL, 0)
}

func (s *Type_Context) STRING() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTRING, 0)
}

func (s *Type_Context) BYTES() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBYTES, 0)
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
	case Protobuf3Visitor:
		return t.VisitType_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Type_() (localctx IType_Context) {
	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, Protobuf3ParserRULE_type_)
	p.SetState(241)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(224)
			p.Match(Protobuf3ParserDOUBLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(225)
			p.Match(Protobuf3ParserFLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(226)
			p.Match(Protobuf3ParserINT32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(227)
			p.Match(Protobuf3ParserINT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(228)
			p.Match(Protobuf3ParserUINT32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(229)
			p.Match(Protobuf3ParserUINT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(230)
			p.Match(Protobuf3ParserSINT32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(231)
			p.Match(Protobuf3ParserSINT64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(232)
			p.Match(Protobuf3ParserFIXED32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(233)
			p.Match(Protobuf3ParserFIXED64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(234)
			p.Match(Protobuf3ParserSFIXED32)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(235)
			p.Match(Protobuf3ParserSFIXED64)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(236)
			p.Match(Protobuf3ParserBOOL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(237)
			p.Match(Protobuf3ParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(238)
			p.Match(Protobuf3ParserBYTES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 16:
		p.EnterOuterAlt(localctx, 16)
		{
			p.SetState(239)
			p.MessageType()
		}

	case 17:
		p.EnterOuterAlt(localctx, 17)
		{
			p.SetState(240)
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
	p.RuleIndex = Protobuf3ParserRULE_reserved
	return p
}

func InitEmptyReservedContext(p *ReservedContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_reserved
}

func (*ReservedContext) IsReservedContext() {}

func NewReservedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReservedContext {
	var p = new(ReservedContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_reserved

	return p
}

func (s *ReservedContext) GetParser() antlr.Parser { return s.parser }

func (s *ReservedContext) RESERVED() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRESERVED, 0)
}

func (s *ReservedContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
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
	case Protobuf3Visitor:
		return t.VisitReserved(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Reserved() (localctx IReservedContext) {
	localctx = NewReservedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, Protobuf3ParserRULE_reserved)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(243)
		p.Match(Protobuf3ParserRESERVED)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(246)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserINT_LIT:
		{
			p.SetState(244)
			p.Ranges()
		}

	case Protobuf3ParserPROTO3_LIT_DOBULE, Protobuf3ParserPROTO3_LIT_SINGLE, Protobuf3ParserSTR_LIT:
		{
			p.SetState(245)
			p.ReservedFieldNames()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(248)
		p.Match(Protobuf3ParserSEMI)
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
	p.RuleIndex = Protobuf3ParserRULE_ranges
	return p
}

func InitEmptyRangesContext(p *RangesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_ranges
}

func (*RangesContext) IsRangesContext() {}

func NewRangesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RangesContext {
	var p = new(RangesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_ranges

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
	return s.GetTokens(Protobuf3ParserCOMMA)
}

func (s *RangesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserCOMMA, i)
}

func (s *RangesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RangesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RangesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitRanges(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Ranges() (localctx IRangesContext) {
	localctx = NewRangesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, Protobuf3ParserRULE_ranges)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(250)
		p.Range_()
	}
	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserCOMMA {
		{
			p.SetState(251)
			p.Match(Protobuf3ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(252)
			p.Range_()
		}

		p.SetState(257)
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
	p.RuleIndex = Protobuf3ParserRULE_range_
	return p
}

func InitEmptyRange_Context(p *Range_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_range_
}

func (*Range_Context) IsRange_Context() {}

func NewRange_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Range_Context {
	var p = new(Range_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_range_

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
	return s.GetToken(Protobuf3ParserTO, 0)
}

func (s *Range_Context) MAX() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMAX, 0)
}

func (s *Range_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Range_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Range_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitRange_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Range_() (localctx IRange_Context) {
	localctx = NewRange_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, Protobuf3ParserRULE_range_)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(258)
		p.IntLit()
	}
	p.SetState(264)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserTO {
		{
			p.SetState(259)
			p.Match(Protobuf3ParserTO)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(262)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case Protobuf3ParserINT_LIT:
			{
				p.SetState(260)
				p.IntLit()
			}

		case Protobuf3ParserMAX:
			{
				p.SetState(261)
				p.Match(Protobuf3ParserMAX)
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
	p.RuleIndex = Protobuf3ParserRULE_reservedFieldNames
	return p
}

func InitEmptyReservedFieldNamesContext(p *ReservedFieldNamesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_reservedFieldNames
}

func (*ReservedFieldNamesContext) IsReservedFieldNamesContext() {}

func NewReservedFieldNamesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReservedFieldNamesContext {
	var p = new(ReservedFieldNamesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_reservedFieldNames

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
	return s.GetTokens(Protobuf3ParserCOMMA)
}

func (s *ReservedFieldNamesContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserCOMMA, i)
}

func (s *ReservedFieldNamesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReservedFieldNamesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReservedFieldNamesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitReservedFieldNames(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) ReservedFieldNames() (localctx IReservedFieldNamesContext) {
	localctx = NewReservedFieldNamesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, Protobuf3ParserRULE_reservedFieldNames)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.StrLit()
	}
	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserCOMMA {
		{
			p.SetState(267)
			p.Match(Protobuf3ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(268)
			p.StrLit()
		}

		p.SetState(273)
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
	p.RuleIndex = Protobuf3ParserRULE_topLevelDef
	return p
}

func InitEmptyTopLevelDefContext(p *TopLevelDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_topLevelDef
}

func (*TopLevelDefContext) IsTopLevelDefContext() {}

func NewTopLevelDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TopLevelDefContext {
	var p = new(TopLevelDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_topLevelDef

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

func (s *TopLevelDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TopLevelDefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TopLevelDefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitTopLevelDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) TopLevelDef() (localctx ITopLevelDefContext) {
	localctx = NewTopLevelDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, Protobuf3ParserRULE_topLevelDef)
	p.SetState(277)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserMESSAGE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(274)
			p.MessageDef()
		}

	case Protobuf3ParserENUM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(275)
			p.EnumDef()
		}

	case Protobuf3ParserSERVICE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(276)
			p.ServiceDef()
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
	p.RuleIndex = Protobuf3ParserRULE_enumDef
	return p
}

func InitEmptyEnumDefContext(p *EnumDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumDef
}

func (*EnumDefContext) IsEnumDefContext() {}

func NewEnumDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumDefContext {
	var p = new(EnumDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumDef

	return p
}

func (s *EnumDefContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumDefContext) ENUM() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserENUM, 0)
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
	case Protobuf3Visitor:
		return t.VisitEnumDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) EnumDef() (localctx IEnumDefContext) {
	localctx = NewEnumDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, Protobuf3ParserRULE_enumDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(279)
		p.Match(Protobuf3ParserENUM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(280)
		p.EnumName()
	}
	{
		p.SetState(281)
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
	p.RuleIndex = Protobuf3ParserRULE_enumBody
	return p
}

func InitEmptyEnumBodyContext(p *EnumBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumBody
}

func (*EnumBodyContext) IsEnumBodyContext() {}

func NewEnumBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumBodyContext {
	var p = new(EnumBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumBody

	return p
}

func (s *EnumBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumBodyContext) LC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLC, 0)
}

func (s *EnumBodyContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRC, 0)
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
	case Protobuf3Visitor:
		return t.VisitEnumBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) EnumBody() (localctx IEnumBodyContext) {
	localctx = NewEnumBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, Protobuf3ParserRULE_enumBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)
		p.Match(Protobuf3ParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&40532482545680382) != 0 {
		{
			p.SetState(284)
			p.EnumElement()
		}

		p.SetState(289)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(290)
		p.Match(Protobuf3ParserRC)
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
	p.RuleIndex = Protobuf3ParserRULE_enumElement
	return p
}

func InitEmptyEnumElementContext(p *EnumElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumElement
}

func (*EnumElementContext) IsEnumElementContext() {}

func NewEnumElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumElementContext {
	var p = new(EnumElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumElement

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
	case Protobuf3Visitor:
		return t.VisitEnumElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) EnumElement() (localctx IEnumElementContext) {
	localctx = NewEnumElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, Protobuf3ParserRULE_enumElement)
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(292)
			p.OptionStatement()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(293)
			p.EnumField()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(294)
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
	p.RuleIndex = Protobuf3ParserRULE_enumField
	return p
}

func InitEmptyEnumFieldContext(p *EnumFieldContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumField
}

func (*EnumFieldContext) IsEnumFieldContext() {}

func NewEnumFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumFieldContext {
	var p = new(EnumFieldContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumField

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
	return s.GetToken(Protobuf3ParserEQ, 0)
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
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *EnumFieldContext) MINUS() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMINUS, 0)
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
	case Protobuf3Visitor:
		return t.VisitEnumField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) EnumField() (localctx IEnumFieldContext) {
	localctx = NewEnumFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, Protobuf3ParserRULE_enumField)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		p.Ident()
	}
	{
		p.SetState(298)
		p.Match(Protobuf3ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(300)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserMINUS {
		{
			p.SetState(299)
			p.Match(Protobuf3ParserMINUS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(302)
		p.IntLit()
	}
	p.SetState(304)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserLB {
		{
			p.SetState(303)
			p.EnumValueOptions()
		}

	}
	{
		p.SetState(306)
		p.Match(Protobuf3ParserSEMI)
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
	p.RuleIndex = Protobuf3ParserRULE_enumValueOptions
	return p
}

func InitEmptyEnumValueOptionsContext(p *EnumValueOptionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumValueOptions
}

func (*EnumValueOptionsContext) IsEnumValueOptionsContext() {}

func NewEnumValueOptionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueOptionsContext {
	var p = new(EnumValueOptionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumValueOptions

	return p
}

func (s *EnumValueOptionsContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueOptionsContext) LB() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLB, 0)
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
	return s.GetToken(Protobuf3ParserRB, 0)
}

func (s *EnumValueOptionsContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(Protobuf3ParserCOMMA)
}

func (s *EnumValueOptionsContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserCOMMA, i)
}

func (s *EnumValueOptionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueOptionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueOptionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitEnumValueOptions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) EnumValueOptions() (localctx IEnumValueOptionsContext) {
	localctx = NewEnumValueOptionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, Protobuf3ParserRULE_enumValueOptions)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(308)
		p.Match(Protobuf3ParserLB)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(309)
		p.EnumValueOption()
	}
	p.SetState(314)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserCOMMA {
		{
			p.SetState(310)
			p.Match(Protobuf3ParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(311)
			p.EnumValueOption()
		}

		p.SetState(316)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(317)
		p.Match(Protobuf3ParserRB)
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
	p.RuleIndex = Protobuf3ParserRULE_enumValueOption
	return p
}

func InitEmptyEnumValueOptionContext(p *EnumValueOptionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumValueOption
}

func (*EnumValueOptionContext) IsEnumValueOptionContext() {}

func NewEnumValueOptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueOptionContext {
	var p = new(EnumValueOptionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumValueOption

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
	return s.GetToken(Protobuf3ParserEQ, 0)
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
	case Protobuf3Visitor:
		return t.VisitEnumValueOption(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) EnumValueOption() (localctx IEnumValueOptionContext) {
	localctx = NewEnumValueOptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, Protobuf3ParserRULE_enumValueOption)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(319)
		p.OptionName()
	}
	{
		p.SetState(320)
		p.Match(Protobuf3ParserEQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(321)
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
	p.RuleIndex = Protobuf3ParserRULE_messageDef
	return p
}

func InitEmptyMessageDefContext(p *MessageDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageDef
}

func (*MessageDefContext) IsMessageDefContext() {}

func NewMessageDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageDefContext {
	var p = new(MessageDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageDef

	return p
}

func (s *MessageDefContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageDefContext) MESSAGE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMESSAGE, 0)
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
	case Protobuf3Visitor:
		return t.VisitMessageDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) MessageDef() (localctx IMessageDefContext) {
	localctx = NewMessageDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, Protobuf3ParserRULE_messageDef)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(323)
		p.Match(Protobuf3ParserMESSAGE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(324)
		p.MessageName()
	}
	{
		p.SetState(325)
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
	p.RuleIndex = Protobuf3ParserRULE_messageBody
	return p
}

func InitEmptyMessageBodyContext(p *MessageBodyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageBody
}

func (*MessageBodyContext) IsMessageBodyContext() {}

func NewMessageBodyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageBodyContext {
	var p = new(MessageBodyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageBody

	return p
}

func (s *MessageBodyContext) GetParser() antlr.Parser { return s.parser }

func (s *MessageBodyContext) LC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLC, 0)
}

func (s *MessageBodyContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRC, 0)
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
	case Protobuf3Visitor:
		return t.VisitMessageBody(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) MessageBody() (localctx IMessageBodyContext) {
	localctx = NewMessageBodyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, Protobuf3ParserRULE_messageBody)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(327)
		p.Match(Protobuf3ParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(331)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&40602851289858046) != 0 {
		{
			p.SetState(328)
			p.MessageElement()
		}

		p.SetState(333)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(334)
		p.Match(Protobuf3ParserRC)
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
	OptionStatement() IOptionStatementContext
	Oneof() IOneofContext
	MapField() IMapFieldContext
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
	p.RuleIndex = Protobuf3ParserRULE_messageElement
	return p
}

func InitEmptyMessageElementContext(p *MessageElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageElement
}

func (*MessageElementContext) IsMessageElementContext() {}

func NewMessageElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageElementContext {
	var p = new(MessageElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageElement

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
	case Protobuf3Visitor:
		return t.VisitMessageElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) MessageElement() (localctx IMessageElementContext) {
	localctx = NewMessageElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, Protobuf3ParserRULE_messageElement)
	p.SetState(344)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(336)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(337)
			p.EnumDef()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(338)
			p.MessageDef()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(339)
			p.OptionStatement()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(340)
			p.Oneof()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(341)
			p.MapField()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(342)
			p.Reserved()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(343)
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
	p.RuleIndex = Protobuf3ParserRULE_serviceDef
	return p
}

func InitEmptyServiceDefContext(p *ServiceDefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_serviceDef
}

func (*ServiceDefContext) IsServiceDefContext() {}

func NewServiceDefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceDefContext {
	var p = new(ServiceDefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_serviceDef

	return p
}

func (s *ServiceDefContext) GetParser() antlr.Parser { return s.parser }

func (s *ServiceDefContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSERVICE, 0)
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
	return s.GetToken(Protobuf3ParserLC, 0)
}

func (s *ServiceDefContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRC, 0)
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
	case Protobuf3Visitor:
		return t.VisitServiceDef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) ServiceDef() (localctx IServiceDefContext) {
	localctx = NewServiceDefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, Protobuf3ParserRULE_serviceDef)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(346)
		p.Match(Protobuf3ParserSERVICE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(347)
		p.ServiceName()
	}
	{
		p.SetState(348)
		p.Match(Protobuf3ParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&70866960448) != 0 {
		{
			p.SetState(349)
			p.ServiceElement()
		}

		p.SetState(354)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(355)
		p.Match(Protobuf3ParserRC)
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
	p.RuleIndex = Protobuf3ParserRULE_serviceElement
	return p
}

func InitEmptyServiceElementContext(p *ServiceElementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_serviceElement
}

func (*ServiceElementContext) IsServiceElementContext() {}

func NewServiceElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceElementContext {
	var p = new(ServiceElementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_serviceElement

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
	case Protobuf3Visitor:
		return t.VisitServiceElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) ServiceElement() (localctx IServiceElementContext) {
	localctx = NewServiceElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, Protobuf3ParserRULE_serviceElement)
	p.SetState(360)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserOPTION:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(357)
			p.OptionStatement()
		}

	case Protobuf3ParserRPC:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(358)
			p.Rpc()
		}

	case Protobuf3ParserSEMI:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(359)
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
	RequestStream() IRequestStreamContext
	ResponseStream() IResponseStreamContext
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
	p.RuleIndex = Protobuf3ParserRULE_rpc
	return p
}

func InitEmptyRpcContext(p *RpcContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_rpc
}

func (*RpcContext) IsRpcContext() {}

func NewRpcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcContext {
	var p = new(RpcContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_rpc

	return p
}

func (s *RpcContext) GetParser() antlr.Parser { return s.parser }

func (s *RpcContext) RPC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRPC, 0)
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
	return s.GetTokens(Protobuf3ParserLP)
}

func (s *RpcContext) LP(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLP, i)
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
	return s.GetTokens(Protobuf3ParserRP)
}

func (s *RpcContext) RP(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRP, i)
}

func (s *RpcContext) RETURNS() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRETURNS, 0)
}

func (s *RpcContext) LC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLC, 0)
}

func (s *RpcContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRC, 0)
}

func (s *RpcContext) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *RpcContext) RequestStream() IRequestStreamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRequestStreamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRequestStreamContext)
}

func (s *RpcContext) ResponseStream() IResponseStreamContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IResponseStreamContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IResponseStreamContext)
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
	case Protobuf3Visitor:
		return t.VisitRpc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Rpc() (localctx IRpcContext) {
	localctx = NewRpcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, Protobuf3ParserRULE_rpc)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(362)
		p.Match(Protobuf3ParserRPC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(363)
		p.RpcName()
	}
	{
		p.SetState(364)
		p.Match(Protobuf3ParserLP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(366)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(365)
			p.RequestStream()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(368)
		p.MessageType()
	}
	{
		p.SetState(369)
		p.Match(Protobuf3ParserRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(370)
		p.Match(Protobuf3ParserRETURNS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(371)
		p.Match(Protobuf3ParserLP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(373)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(372)
			p.ResponseStream()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	{
		p.SetState(375)
		p.MessageType()
	}
	{
		p.SetState(376)
		p.Match(Protobuf3ParserRP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserLC:
		{
			p.SetState(377)
			p.Match(Protobuf3ParserLC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(382)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == Protobuf3ParserOPTION || _la == Protobuf3ParserSEMI {
			p.SetState(380)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case Protobuf3ParserOPTION:
				{
					p.SetState(378)
					p.OptionStatement()
				}

			case Protobuf3ParserSEMI:
				{
					p.SetState(379)
					p.EmptyStatement_()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

			p.SetState(384)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(385)
			p.Match(Protobuf3ParserRC)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Protobuf3ParserSEMI:
		{
			p.SetState(386)
			p.Match(Protobuf3ParserSEMI)
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

// IRequestStreamContext is an interface to support dynamic dispatch.
type IRequestStreamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STREAM() antlr.TerminalNode

	// IsRequestStreamContext differentiates from other interfaces.
	IsRequestStreamContext()
}

type RequestStreamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRequestStreamContext() *RequestStreamContext {
	var p = new(RequestStreamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_requestStream
	return p
}

func InitEmptyRequestStreamContext(p *RequestStreamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_requestStream
}

func (*RequestStreamContext) IsRequestStreamContext() {}

func NewRequestStreamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RequestStreamContext {
	var p = new(RequestStreamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_requestStream

	return p
}

func (s *RequestStreamContext) GetParser() antlr.Parser { return s.parser }

func (s *RequestStreamContext) STREAM() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTREAM, 0)
}

func (s *RequestStreamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RequestStreamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RequestStreamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitRequestStream(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) RequestStream() (localctx IRequestStreamContext) {
	localctx = NewRequestStreamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, Protobuf3ParserRULE_requestStream)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(389)
		p.Match(Protobuf3ParserSTREAM)
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

// IResponseStreamContext is an interface to support dynamic dispatch.
type IResponseStreamContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STREAM() antlr.TerminalNode

	// IsResponseStreamContext differentiates from other interfaces.
	IsResponseStreamContext()
}

type ResponseStreamContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyResponseStreamContext() *ResponseStreamContext {
	var p = new(ResponseStreamContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_responseStream
	return p
}

func InitEmptyResponseStreamContext(p *ResponseStreamContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_responseStream
}

func (*ResponseStreamContext) IsResponseStreamContext() {}

func NewResponseStreamContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ResponseStreamContext {
	var p = new(ResponseStreamContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_responseStream

	return p
}

func (s *ResponseStreamContext) GetParser() antlr.Parser { return s.parser }

func (s *ResponseStreamContext) STREAM() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTREAM, 0)
}

func (s *ResponseStreamContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ResponseStreamContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ResponseStreamContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitResponseStream(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) ResponseStream() (localctx IResponseStreamContext) {
	localctx = NewResponseStreamContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, Protobuf3ParserRULE_responseStream)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(391)
		p.Match(Protobuf3ParserSTREAM)
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
	p.RuleIndex = Protobuf3ParserRULE_constant
	return p
}

func InitEmptyConstantContext(p *ConstantContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_constant
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_constant

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
	return s.GetToken(Protobuf3ParserMINUS, 0)
}

func (s *ConstantContext) PLUS() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPLUS, 0)
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
	case Protobuf3Visitor:
		return t.VisitConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, Protobuf3ParserRULE_constant)
	var _la int

	p.SetState(405)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(393)
			p.FullIdent()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(395)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Protobuf3ParserPLUS || _la == Protobuf3ParserMINUS {
			{
				p.SetState(394)
				_la = p.GetTokenStream().LA(1)

				if !(_la == Protobuf3ParserPLUS || _la == Protobuf3ParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(397)
			p.IntLit()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(399)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == Protobuf3ParserPLUS || _la == Protobuf3ParserMINUS {
			{
				p.SetState(398)
				_la = p.GetTokenStream().LA(1)

				if !(_la == Protobuf3ParserPLUS || _la == Protobuf3ParserMINUS) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}

		}
		{
			p.SetState(401)
			p.FloatLit()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(402)
			p.StrLit()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(403)
			p.BoolLit()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(404)
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
	p.RuleIndex = Protobuf3ParserRULE_blockLit
	return p
}

func InitEmptyBlockLitContext(p *BlockLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_blockLit
}

func (*BlockLitContext) IsBlockLitContext() {}

func NewBlockLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockLitContext {
	var p = new(BlockLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_blockLit

	return p
}

func (s *BlockLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockLitContext) LC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserLC, 0)
}

func (s *BlockLitContext) RC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRC, 0)
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
	return s.GetTokens(Protobuf3ParserCOLON)
}

func (s *BlockLitContext) COLON(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserCOLON, i)
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
	case Protobuf3Visitor:
		return t.VisitBlockLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) BlockLit() (localctx IBlockLitContext) {
	localctx = NewBlockLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, Protobuf3ParserRULE_blockLit)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(407)
		p.Match(Protobuf3ParserLC)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&40532413826203646) != 0 {
		{
			p.SetState(408)
			p.Ident()
		}
		{
			p.SetState(409)
			p.Match(Protobuf3ParserCOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(410)
			p.Constant()
		}

		p.SetState(416)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(417)
		p.Match(Protobuf3ParserRC)
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
	p.RuleIndex = Protobuf3ParserRULE_emptyStatement_
	return p
}

func InitEmptyEmptyStatement_Context(p *EmptyStatement_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_emptyStatement_
}

func (*EmptyStatement_Context) IsEmptyStatement_Context() {}

func NewEmptyStatement_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EmptyStatement_Context {
	var p = new(EmptyStatement_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_emptyStatement_

	return p
}

func (s *EmptyStatement_Context) GetParser() antlr.Parser { return s.parser }

func (s *EmptyStatement_Context) SEMI() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSEMI, 0)
}

func (s *EmptyStatement_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EmptyStatement_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EmptyStatement_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitEmptyStatement_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) EmptyStatement_() (localctx IEmptyStatement_Context) {
	localctx = NewEmptyStatement_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, Protobuf3ParserRULE_emptyStatement_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(419)
		p.Match(Protobuf3ParserSEMI)
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
	p.RuleIndex = Protobuf3ParserRULE_ident
	return p
}

func InitEmptyIdentContext(p *IdentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_ident
}

func (*IdentContext) IsIdentContext() {}

func NewIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IdentContext {
	var p = new(IdentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_ident

	return p
}

func (s *IdentContext) GetParser() antlr.Parser { return s.parser }

func (s *IdentContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIDENTIFIER, 0)
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
	case Protobuf3Visitor:
		return t.VisitIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Ident() (localctx IIdentContext) {
	localctx = NewIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, Protobuf3ParserRULE_ident)
	p.SetState(423)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case Protobuf3ParserIDENTIFIER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(421)
			p.Match(Protobuf3ParserIDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case Protobuf3ParserSYNTAX, Protobuf3ParserIMPORT, Protobuf3ParserWEAK, Protobuf3ParserPUBLIC, Protobuf3ParserPACKAGE, Protobuf3ParserOPTION, Protobuf3ParserREPEATED, Protobuf3ParserONEOF, Protobuf3ParserMAP, Protobuf3ParserINT32, Protobuf3ParserINT64, Protobuf3ParserUINT32, Protobuf3ParserUINT64, Protobuf3ParserSINT32, Protobuf3ParserSINT64, Protobuf3ParserFIXED32, Protobuf3ParserFIXED64, Protobuf3ParserSFIXED32, Protobuf3ParserSFIXED64, Protobuf3ParserBOOL, Protobuf3ParserSTRING, Protobuf3ParserDOUBLE, Protobuf3ParserFLOAT, Protobuf3ParserBYTES, Protobuf3ParserRESERVED, Protobuf3ParserTO, Protobuf3ParserMAX, Protobuf3ParserENUM, Protobuf3ParserMESSAGE, Protobuf3ParserSERVICE, Protobuf3ParserRPC, Protobuf3ParserSTREAM, Protobuf3ParserRETURNS, Protobuf3ParserBOOL_LIT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(422)
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
	p.RuleIndex = Protobuf3ParserRULE_fullIdent
	return p
}

func InitEmptyFullIdentContext(p *FullIdentContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fullIdent
}

func (*FullIdentContext) IsFullIdentContext() {}

func NewFullIdentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FullIdentContext {
	var p = new(FullIdentContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fullIdent

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
	return s.GetTokens(Protobuf3ParserDOT)
}

func (s *FullIdentContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserDOT, i)
}

func (s *FullIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FullIdentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FullIdentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitFullIdent(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) FullIdent() (localctx IFullIdentContext) {
	localctx = NewFullIdentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, Protobuf3ParserRULE_fullIdent)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(425)
		p.Ident()
	}
	p.SetState(430)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == Protobuf3ParserDOT {
		{
			p.SetState(426)
			p.Match(Protobuf3ParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(427)
			p.Ident()
		}

		p.SetState(432)
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
	p.RuleIndex = Protobuf3ParserRULE_messageName
	return p
}

func InitEmptyMessageNameContext(p *MessageNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageName
}

func (*MessageNameContext) IsMessageNameContext() {}

func NewMessageNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageNameContext {
	var p = new(MessageNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageName

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
	case Protobuf3Visitor:
		return t.VisitMessageName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) MessageName() (localctx IMessageNameContext) {
	localctx = NewMessageNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, Protobuf3ParserRULE_messageName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(433)
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
	p.RuleIndex = Protobuf3ParserRULE_enumName
	return p
}

func InitEmptyEnumNameContext(p *EnumNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumName
}

func (*EnumNameContext) IsEnumNameContext() {}

func NewEnumNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumNameContext {
	var p = new(EnumNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumName

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
	case Protobuf3Visitor:
		return t.VisitEnumName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) EnumName() (localctx IEnumNameContext) {
	localctx = NewEnumNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, Protobuf3ParserRULE_enumName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(435)
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
	p.RuleIndex = Protobuf3ParserRULE_fieldName
	return p
}

func InitEmptyFieldNameContext(p *FieldNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_fieldName
}

func (*FieldNameContext) IsFieldNameContext() {}

func NewFieldNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldNameContext {
	var p = new(FieldNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_fieldName

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
	case Protobuf3Visitor:
		return t.VisitFieldName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) FieldName() (localctx IFieldNameContext) {
	localctx = NewFieldNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, Protobuf3ParserRULE_fieldName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(437)
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
	p.RuleIndex = Protobuf3ParserRULE_oneofName
	return p
}

func InitEmptyOneofNameContext(p *OneofNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_oneofName
}

func (*OneofNameContext) IsOneofNameContext() {}

func NewOneofNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OneofNameContext {
	var p = new(OneofNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_oneofName

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
	case Protobuf3Visitor:
		return t.VisitOneofName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) OneofName() (localctx IOneofNameContext) {
	localctx = NewOneofNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, Protobuf3ParserRULE_oneofName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(439)
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
	p.RuleIndex = Protobuf3ParserRULE_mapName
	return p
}

func InitEmptyMapNameContext(p *MapNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_mapName
}

func (*MapNameContext) IsMapNameContext() {}

func NewMapNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MapNameContext {
	var p = new(MapNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_mapName

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
	case Protobuf3Visitor:
		return t.VisitMapName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) MapName() (localctx IMapNameContext) {
	localctx = NewMapNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, Protobuf3ParserRULE_mapName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(441)
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
	p.RuleIndex = Protobuf3ParserRULE_serviceName
	return p
}

func InitEmptyServiceNameContext(p *ServiceNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_serviceName
}

func (*ServiceNameContext) IsServiceNameContext() {}

func NewServiceNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ServiceNameContext {
	var p = new(ServiceNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_serviceName

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
	case Protobuf3Visitor:
		return t.VisitServiceName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) ServiceName() (localctx IServiceNameContext) {
	localctx = NewServiceNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, Protobuf3ParserRULE_serviceName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(443)
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
	p.RuleIndex = Protobuf3ParserRULE_rpcName
	return p
}

func InitEmptyRpcNameContext(p *RpcNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_rpcName
}

func (*RpcNameContext) IsRpcNameContext() {}

func NewRpcNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RpcNameContext {
	var p = new(RpcNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_rpcName

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
	case Protobuf3Visitor:
		return t.VisitRpcName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) RpcName() (localctx IRpcNameContext) {
	localctx = NewRpcNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, Protobuf3ParserRULE_rpcName)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(445)
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
	p.RuleIndex = Protobuf3ParserRULE_messageType
	return p
}

func InitEmptyMessageTypeContext(p *MessageTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_messageType
}

func (*MessageTypeContext) IsMessageTypeContext() {}

func NewMessageTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MessageTypeContext {
	var p = new(MessageTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_messageType

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
	return s.GetTokens(Protobuf3ParserDOT)
}

func (s *MessageTypeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserDOT, i)
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
	case Protobuf3Visitor:
		return t.VisitMessageType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) MessageType() (localctx IMessageTypeContext) {
	localctx = NewMessageTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, Protobuf3ParserRULE_messageType)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(448)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserDOT {
		{
			p.SetState(447)
			p.Match(Protobuf3ParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(455)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(450)
				p.Ident()
			}
			{
				p.SetState(451)
				p.Match(Protobuf3ParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(457)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 40, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(458)
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
	p.RuleIndex = Protobuf3ParserRULE_enumType
	return p
}

func InitEmptyEnumTypeContext(p *EnumTypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_enumType
}

func (*EnumTypeContext) IsEnumTypeContext() {}

func NewEnumTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumTypeContext {
	var p = new(EnumTypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_enumType

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
	return s.GetTokens(Protobuf3ParserDOT)
}

func (s *EnumTypeContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserDOT, i)
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
	case Protobuf3Visitor:
		return t.VisitEnumType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) EnumType() (localctx IEnumTypeContext) {
	localctx = NewEnumTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, Protobuf3ParserRULE_enumType)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(461)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == Protobuf3ParserDOT {
		{
			p.SetState(460)
			p.Match(Protobuf3ParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(468)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(463)
				p.Ident()
			}
			{
				p.SetState(464)
				p.Match(Protobuf3ParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(470)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 42, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}
	{
		p.SetState(471)
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
	p.RuleIndex = Protobuf3ParserRULE_intLit
	return p
}

func InitEmptyIntLitContext(p *IntLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_intLit
}

func (*IntLitContext) IsIntLitContext() {}

func NewIntLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntLitContext {
	var p = new(IntLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_intLit

	return p
}

func (s *IntLitContext) GetParser() antlr.Parser { return s.parser }

func (s *IntLitContext) INT_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT_LIT, 0)
}

func (s *IntLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitIntLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) IntLit() (localctx IIntLitContext) {
	localctx = NewIntLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, Protobuf3ParserRULE_intLit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(473)
		p.Match(Protobuf3ParserINT_LIT)
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
	PROTO3_LIT_SINGLE() antlr.TerminalNode
	PROTO3_LIT_DOBULE() antlr.TerminalNode

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
	p.RuleIndex = Protobuf3ParserRULE_strLit
	return p
}

func InitEmptyStrLitContext(p *StrLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_strLit
}

func (*StrLitContext) IsStrLitContext() {}

func NewStrLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StrLitContext {
	var p = new(StrLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_strLit

	return p
}

func (s *StrLitContext) GetParser() antlr.Parser { return s.parser }

func (s *StrLitContext) STR_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTR_LIT, 0)
}

func (s *StrLitContext) PROTO3_LIT_SINGLE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPROTO3_LIT_SINGLE, 0)
}

func (s *StrLitContext) PROTO3_LIT_DOBULE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPROTO3_LIT_DOBULE, 0)
}

func (s *StrLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StrLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitStrLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) StrLit() (localctx IStrLitContext) {
	localctx = NewStrLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, Protobuf3ParserRULE_strLit)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(475)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2251851353292800) != 0) {
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
	p.RuleIndex = Protobuf3ParserRULE_boolLit
	return p
}

func InitEmptyBoolLitContext(p *BoolLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_boolLit
}

func (*BoolLitContext) IsBoolLitContext() {}

func NewBoolLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolLitContext {
	var p = new(BoolLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_boolLit

	return p
}

func (s *BoolLitContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolLitContext) BOOL_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBOOL_LIT, 0)
}

func (s *BoolLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoolLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitBoolLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) BoolLit() (localctx IBoolLitContext) {
	localctx = NewBoolLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, Protobuf3ParserRULE_boolLit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(477)
		p.Match(Protobuf3ParserBOOL_LIT)
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
	p.RuleIndex = Protobuf3ParserRULE_floatLit
	return p
}

func InitEmptyFloatLitContext(p *FloatLitContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_floatLit
}

func (*FloatLitContext) IsFloatLitContext() {}

func NewFloatLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatLitContext {
	var p = new(FloatLitContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_floatLit

	return p
}

func (s *FloatLitContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatLitContext) FLOAT_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFLOAT_LIT, 0)
}

func (s *FloatLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatLitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitFloatLit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) FloatLit() (localctx IFloatLitContext) {
	localctx = NewFloatLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, Protobuf3ParserRULE_floatLit)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(479)
		p.Match(Protobuf3ParserFLOAT_LIT)
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
	TO() antlr.TerminalNode
	MAX() antlr.TerminalNode
	ENUM() antlr.TerminalNode
	MESSAGE() antlr.TerminalNode
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
	p.RuleIndex = Protobuf3ParserRULE_keywords
	return p
}

func InitEmptyKeywordsContext(p *KeywordsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = Protobuf3ParserRULE_keywords
}

func (*KeywordsContext) IsKeywordsContext() {}

func NewKeywordsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *KeywordsContext {
	var p = new(KeywordsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = Protobuf3ParserRULE_keywords

	return p
}

func (s *KeywordsContext) GetParser() antlr.Parser { return s.parser }

func (s *KeywordsContext) SYNTAX() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSYNTAX, 0)
}

func (s *KeywordsContext) IMPORT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserIMPORT, 0)
}

func (s *KeywordsContext) WEAK() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserWEAK, 0)
}

func (s *KeywordsContext) PUBLIC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPUBLIC, 0)
}

func (s *KeywordsContext) PACKAGE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserPACKAGE, 0)
}

func (s *KeywordsContext) OPTION() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserOPTION, 0)
}

func (s *KeywordsContext) REPEATED() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserREPEATED, 0)
}

func (s *KeywordsContext) ONEOF() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserONEOF, 0)
}

func (s *KeywordsContext) MAP() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMAP, 0)
}

func (s *KeywordsContext) INT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT32, 0)
}

func (s *KeywordsContext) INT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserINT64, 0)
}

func (s *KeywordsContext) UINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserUINT32, 0)
}

func (s *KeywordsContext) UINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserUINT64, 0)
}

func (s *KeywordsContext) SINT32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSINT32, 0)
}

func (s *KeywordsContext) SINT64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSINT64, 0)
}

func (s *KeywordsContext) FIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFIXED32, 0)
}

func (s *KeywordsContext) FIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFIXED64, 0)
}

func (s *KeywordsContext) SFIXED32() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSFIXED32, 0)
}

func (s *KeywordsContext) SFIXED64() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSFIXED64, 0)
}

func (s *KeywordsContext) BOOL() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBOOL, 0)
}

func (s *KeywordsContext) STRING() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTRING, 0)
}

func (s *KeywordsContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserDOUBLE, 0)
}

func (s *KeywordsContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserFLOAT, 0)
}

func (s *KeywordsContext) BYTES() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBYTES, 0)
}

func (s *KeywordsContext) RESERVED() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRESERVED, 0)
}

func (s *KeywordsContext) TO() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserTO, 0)
}

func (s *KeywordsContext) MAX() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMAX, 0)
}

func (s *KeywordsContext) ENUM() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserENUM, 0)
}

func (s *KeywordsContext) MESSAGE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserMESSAGE, 0)
}

func (s *KeywordsContext) SERVICE() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSERVICE, 0)
}

func (s *KeywordsContext) RPC() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRPC, 0)
}

func (s *KeywordsContext) STREAM() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserSTREAM, 0)
}

func (s *KeywordsContext) RETURNS() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserRETURNS, 0)
}

func (s *KeywordsContext) BOOL_LIT() antlr.TerminalNode {
	return s.GetToken(Protobuf3ParserBOOL_LIT, 0)
}

func (s *KeywordsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeywordsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *KeywordsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case Protobuf3Visitor:
		return t.VisitKeywords(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *Protobuf3Parser) Keywords() (localctx IKeywordsContext) {
	localctx = NewKeywordsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, Protobuf3ParserRULE_keywords)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(481)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&4503616807239678) != 0) {
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
