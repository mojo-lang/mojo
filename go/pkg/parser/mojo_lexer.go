// Code generated from MojoLexer.g4 by ANTLR 4.8. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 86, 799,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 4, 69, 9, 69, 4, 70, 9,
	70, 4, 71, 9, 71, 4, 72, 9, 72, 4, 73, 9, 73, 4, 74, 9, 74, 4, 75, 9, 75,
	4, 76, 9, 76, 4, 77, 9, 77, 4, 78, 9, 78, 4, 79, 9, 79, 4, 80, 9, 80, 4,
	81, 9, 81, 4, 82, 9, 82, 4, 83, 9, 83, 4, 84, 9, 84, 4, 85, 9, 85, 4, 86,
	9, 86, 4, 87, 9, 87, 4, 88, 9, 88, 4, 89, 9, 89, 4, 90, 9, 90, 4, 91, 9,
	91, 4, 92, 9, 92, 4, 93, 9, 93, 4, 94, 9, 94, 4, 95, 9, 95, 4, 96, 9, 96,
	4, 97, 9, 97, 4, 98, 9, 98, 4, 99, 9, 99, 4, 100, 9, 100, 4, 101, 9, 101,
	4, 102, 9, 102, 4, 103, 9, 103, 4, 104, 9, 104, 4, 105, 9, 105, 4, 106,
	9, 106, 4, 107, 9, 107, 4, 108, 9, 108, 4, 109, 9, 109, 4, 110, 9, 110,
	4, 111, 9, 111, 4, 112, 9, 112, 4, 113, 9, 113, 4, 114, 9, 114, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3,
	25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27,
	3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 33,
	3, 33, 3, 34, 3, 34, 3, 35, 3, 35, 3, 36, 3, 36, 3, 37, 3, 37, 3, 38, 3,
	38, 3, 39, 3, 39, 3, 40, 3, 40, 3, 41, 3, 41, 3, 42, 3, 42, 3, 43, 3, 43,
	3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46, 3, 47, 3, 47, 3, 48, 3, 48, 3,
	49, 3, 49, 3, 50, 3, 50, 3, 51, 3, 51, 3, 52, 3, 52, 3, 53, 3, 53, 3, 54,
	3, 54, 3, 55, 3, 55, 3, 56, 3, 56, 3, 57, 3, 57, 3, 58, 3, 58, 3, 58, 3,
	59, 3, 59, 3, 59, 3, 60, 3, 60, 3, 60, 3, 61, 3, 61, 3, 61, 3, 62, 3, 62,
	3, 62, 3, 62, 3, 63, 3, 63, 3, 63, 3, 63, 3, 64, 3, 64, 3, 64, 3, 64, 3,
	65, 3, 65, 7, 65, 470, 10, 65, 12, 65, 14, 65, 473, 11, 65, 3, 66, 3, 66,
	5, 66, 477, 10, 66, 3, 67, 5, 67, 480, 10, 67, 3, 68, 3, 68, 5, 68, 484,
	10, 68, 3, 69, 6, 69, 487, 10, 69, 13, 69, 14, 69, 488, 3, 70, 5, 70, 492,
	10, 70, 3, 71, 5, 71, 495, 10, 71, 3, 72, 3, 72, 3, 72, 3, 73, 3, 73, 5,
	73, 502, 10, 73, 3, 74, 3, 74, 3, 75, 3, 75, 3, 75, 3, 75, 3, 75, 5, 75,
	511, 10, 75, 3, 76, 3, 76, 3, 77, 3, 77, 5, 77, 517, 10, 77, 3, 78, 6,
	78, 520, 10, 78, 13, 78, 14, 78, 521, 3, 79, 3, 79, 3, 79, 3, 79, 3, 79,
	5, 79, 529, 10, 79, 3, 80, 3, 80, 3, 81, 3, 81, 5, 81, 535, 10, 81, 3,
	82, 6, 82, 538, 10, 82, 13, 82, 14, 82, 539, 3, 83, 3, 83, 7, 83, 544,
	10, 83, 12, 83, 14, 83, 547, 11, 83, 3, 84, 6, 84, 550, 10, 84, 13, 84,
	14, 84, 551, 3, 85, 3, 85, 3, 86, 3, 86, 5, 86, 558, 10, 86, 3, 87, 6,
	87, 561, 10, 87, 13, 87, 14, 87, 562, 3, 88, 3, 88, 3, 88, 3, 88, 3, 88,
	5, 88, 570, 10, 88, 3, 89, 3, 89, 3, 90, 3, 90, 5, 90, 576, 10, 90, 3,
	91, 6, 91, 579, 10, 91, 13, 91, 14, 91, 580, 3, 92, 3, 92, 5, 92, 585,
	10, 92, 3, 92, 5, 92, 588, 10, 92, 3, 92, 3, 92, 5, 92, 592, 10, 92, 3,
	92, 3, 92, 5, 92, 596, 10, 92, 3, 93, 3, 93, 3, 93, 3, 94, 3, 94, 5, 94,
	603, 10, 94, 3, 94, 3, 94, 3, 95, 3, 95, 3, 95, 5, 95, 610, 10, 95, 3,
	96, 3, 96, 5, 96, 614, 10, 96, 3, 96, 3, 96, 3, 97, 3, 97, 3, 98, 3, 98,
	3, 99, 3, 99, 3, 100, 3, 100, 5, 100, 626, 10, 100, 3, 100, 3, 100, 3,
	100, 5, 100, 631, 10, 100, 3, 100, 5, 100, 634, 10, 100, 3, 101, 6, 101,
	637, 10, 101, 13, 101, 14, 101, 638, 3, 102, 6, 102, 642, 10, 102, 13,
	102, 14, 102, 643, 3, 103, 3, 103, 5, 103, 648, 10, 103, 3, 104, 3, 104,
	5, 104, 652, 10, 104, 3, 105, 3, 105, 3, 105, 3, 105, 3, 105, 3, 105, 3,
	105, 3, 105, 3, 105, 3, 105, 3, 105, 3, 105, 3, 105, 3, 105, 3, 105, 3,
	105, 3, 105, 3, 105, 3, 105, 3, 105, 3, 105, 3, 105, 3, 105, 3, 105, 3,
	105, 3, 105, 3, 105, 3, 105, 3, 105, 3, 105, 3, 105, 3, 105, 3, 105, 3,
	105, 3, 105, 3, 105, 3, 105, 3, 105, 3, 105, 3, 105, 5, 105, 694, 10, 105,
	3, 106, 3, 106, 7, 106, 698, 10, 106, 12, 106, 14, 106, 701, 11, 106, 3,
	106, 3, 106, 3, 106, 7, 106, 706, 10, 106, 12, 106, 14, 106, 709, 11, 106,
	3, 106, 5, 106, 712, 10, 106, 3, 107, 3, 107, 3, 107, 3, 107, 3, 107, 6,
	107, 719, 10, 107, 13, 107, 14, 107, 720, 3, 107, 3, 107, 3, 107, 5, 107,
	726, 10, 107, 3, 108, 3, 108, 3, 108, 3, 108, 3, 108, 6, 108, 733, 10,
	108, 13, 108, 14, 108, 734, 3, 108, 3, 108, 3, 108, 5, 108, 740, 10, 108,
	3, 109, 3, 109, 3, 109, 3, 109, 3, 110, 3, 110, 3, 110, 3, 110, 3, 110,
	7, 110, 751, 10, 110, 12, 110, 14, 110, 754, 11, 110, 3, 110, 3, 110, 3,
	110, 3, 110, 3, 110, 3, 111, 3, 111, 3, 111, 3, 111, 3, 111, 7, 111, 766,
	10, 111, 12, 111, 14, 111, 769, 11, 111, 5, 111, 771, 10, 111, 3, 111,
	3, 111, 3, 112, 3, 112, 3, 112, 5, 112, 778, 10, 112, 3, 113, 3, 113, 3,
	113, 3, 113, 3, 113, 7, 113, 785, 10, 113, 12, 113, 14, 113, 788, 11, 113,
	3, 114, 3, 114, 3, 114, 3, 114, 3, 114, 7, 114, 795, 10, 114, 12, 114,
	14, 114, 798, 11, 114, 3, 752, 2, 115, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13,
	8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26,
	51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35,
	69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44,
	87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 97, 50, 99, 51, 101, 52, 103, 53,
	105, 54, 107, 55, 109, 56, 111, 57, 113, 58, 115, 59, 117, 60, 119, 61,
	121, 62, 123, 63, 125, 64, 127, 65, 129, 66, 131, 67, 133, 2, 135, 2, 137,
	2, 139, 68, 141, 69, 143, 70, 145, 71, 147, 72, 149, 73, 151, 2, 153, 2,
	155, 2, 157, 74, 159, 2, 161, 2, 163, 2, 165, 75, 167, 76, 169, 2, 171,
	2, 173, 2, 175, 77, 177, 2, 179, 2, 181, 2, 183, 78, 185, 2, 187, 2, 189,
	2, 191, 2, 193, 2, 195, 2, 197, 2, 199, 79, 201, 2, 203, 2, 205, 2, 207,
	2, 209, 2, 211, 80, 213, 2, 215, 2, 217, 81, 219, 82, 221, 83, 223, 84,
	225, 85, 227, 86, 3, 2, 22, 3, 2, 67, 92, 5, 2, 50, 59, 67, 92, 99, 124,
	35, 2, 99, 124, 170, 170, 172, 172, 175, 175, 177, 177, 180, 183, 185,
	188, 190, 192, 194, 216, 218, 248, 250, 769, 882, 5761, 5763, 6159, 6161,
	7617, 7682, 8193, 8205, 8207, 8236, 8240, 8257, 8258, 8278, 8278, 8290,
	8401, 8450, 8593, 9314, 9473, 10104, 10133, 11266, 11777, 11906, 12289,
	12294, 12297, 12323, 12337, 12339, 55297, 63746, 64831, 64834, 64977, 65010,
	65057, 65074, 65094, 65097, 65535, 9, 2, 50, 59, 67, 92, 97, 97, 770, 881,
	7618, 7681, 8402, 8449, 65058, 65073, 23, 2, 163, 169, 171, 171, 173, 174,
	176, 176, 178, 179, 184, 184, 189, 189, 193, 193, 217, 217, 249, 249, 8216,
	8217, 8226, 8233, 8242, 8256, 8259, 8277, 8279, 8288, 8594, 9217, 9474,
	10103, 10134, 11265, 11778, 11905, 12291, 12293, 12298, 12338, 7, 2, 770,
	881, 7618, 7681, 8402, 8449, 65026, 65041, 65058, 65073, 3, 2, 50, 51,
	3, 2, 50, 57, 3, 2, 50, 59, 4, 2, 50, 59, 97, 97, 5, 2, 50, 59, 67, 72,
	99, 104, 4, 2, 71, 71, 103, 103, 4, 2, 82, 82, 114, 114, 4, 2, 45, 45,
	47, 47, 6, 2, 12, 12, 15, 15, 36, 36, 94, 94, 6, 2, 12, 12, 15, 15, 41,
	41, 94, 94, 9, 2, 36, 36, 41, 41, 50, 50, 94, 94, 112, 112, 116, 116, 118,
	118, 6, 2, 2, 2, 11, 11, 13, 14, 34, 34, 6, 2, 12, 12, 15, 15, 49, 49,
	62, 62, 4, 2, 12, 12, 15, 15, 2, 821, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2,
	2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2,
	2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3,
	2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29,
	3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2,
	37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2,
	2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2,
	2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2,
	2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3,
	2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75,
	3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2,
	83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2,
	2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2,
	2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3,
	2, 2, 2, 2, 107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2,
	113, 3, 2, 2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119, 3, 2,
	2, 2, 2, 121, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2, 2, 127,
	3, 2, 2, 2, 2, 129, 3, 2, 2, 2, 2, 131, 3, 2, 2, 2, 2, 139, 3, 2, 2, 2,
	2, 141, 3, 2, 2, 2, 2, 143, 3, 2, 2, 2, 2, 145, 3, 2, 2, 2, 2, 147, 3,
	2, 2, 2, 2, 149, 3, 2, 2, 2, 2, 157, 3, 2, 2, 2, 2, 165, 3, 2, 2, 2, 2,
	167, 3, 2, 2, 2, 2, 175, 3, 2, 2, 2, 2, 183, 3, 2, 2, 2, 2, 199, 3, 2,
	2, 2, 2, 211, 3, 2, 2, 2, 2, 217, 3, 2, 2, 2, 2, 219, 3, 2, 2, 2, 2, 221,
	3, 2, 2, 2, 2, 223, 3, 2, 2, 2, 2, 225, 3, 2, 2, 2, 2, 227, 3, 2, 2, 2,
	3, 229, 3, 2, 2, 2, 5, 233, 3, 2, 2, 2, 7, 236, 3, 2, 2, 2, 9, 246, 3,
	2, 2, 2, 11, 252, 3, 2, 2, 2, 13, 258, 3, 2, 2, 2, 15, 267, 3, 2, 2, 2,
	17, 272, 3, 2, 2, 2, 19, 277, 3, 2, 2, 2, 21, 283, 3, 2, 2, 2, 23, 287,
	3, 2, 2, 2, 25, 292, 3, 2, 2, 2, 27, 295, 3, 2, 2, 2, 29, 302, 3, 2, 2,
	2, 31, 305, 3, 2, 2, 2, 33, 315, 3, 2, 2, 2, 35, 318, 3, 2, 2, 2, 37, 324,
	3, 2, 2, 2, 39, 328, 3, 2, 2, 2, 41, 333, 3, 2, 2, 2, 43, 336, 3, 2, 2,
	2, 45, 344, 3, 2, 2, 2, 47, 351, 3, 2, 2, 2, 49, 358, 3, 2, 2, 2, 51, 365,
	3, 2, 2, 2, 53, 370, 3, 2, 2, 2, 55, 375, 3, 2, 2, 2, 57, 379, 3, 2, 2,
	2, 59, 385, 3, 2, 2, 2, 61, 389, 3, 2, 2, 2, 63, 391, 3, 2, 2, 2, 65, 393,
	3, 2, 2, 2, 67, 395, 3, 2, 2, 2, 69, 397, 3, 2, 2, 2, 71, 399, 3, 2, 2,
	2, 73, 401, 3, 2, 2, 2, 75, 403, 3, 2, 2, 2, 77, 405, 3, 2, 2, 2, 79, 407,
	3, 2, 2, 2, 81, 409, 3, 2, 2, 2, 83, 411, 3, 2, 2, 2, 85, 413, 3, 2, 2,
	2, 87, 415, 3, 2, 2, 2, 89, 417, 3, 2, 2, 2, 91, 419, 3, 2, 2, 2, 93, 421,
	3, 2, 2, 2, 95, 423, 3, 2, 2, 2, 97, 425, 3, 2, 2, 2, 99, 427, 3, 2, 2,
	2, 101, 429, 3, 2, 2, 2, 103, 431, 3, 2, 2, 2, 105, 433, 3, 2, 2, 2, 107,
	435, 3, 2, 2, 2, 109, 437, 3, 2, 2, 2, 111, 439, 3, 2, 2, 2, 113, 441,
	3, 2, 2, 2, 115, 443, 3, 2, 2, 2, 117, 446, 3, 2, 2, 2, 119, 449, 3, 2,
	2, 2, 121, 452, 3, 2, 2, 2, 123, 455, 3, 2, 2, 2, 125, 459, 3, 2, 2, 2,
	127, 463, 3, 2, 2, 2, 129, 467, 3, 2, 2, 2, 131, 474, 3, 2, 2, 2, 133,
	479, 3, 2, 2, 2, 135, 483, 3, 2, 2, 2, 137, 486, 3, 2, 2, 2, 139, 491,
	3, 2, 2, 2, 141, 494, 3, 2, 2, 2, 143, 496, 3, 2, 2, 2, 145, 501, 3, 2,
	2, 2, 147, 503, 3, 2, 2, 2, 149, 505, 3, 2, 2, 2, 151, 512, 3, 2, 2, 2,
	153, 516, 3, 2, 2, 2, 155, 519, 3, 2, 2, 2, 157, 523, 3, 2, 2, 2, 159,
	530, 3, 2, 2, 2, 161, 534, 3, 2, 2, 2, 163, 537, 3, 2, 2, 2, 165, 541,
	3, 2, 2, 2, 167, 549, 3, 2, 2, 2, 169, 553, 3, 2, 2, 2, 171, 557, 3, 2,
	2, 2, 173, 560, 3, 2, 2, 2, 175, 564, 3, 2, 2, 2, 177, 571, 3, 2, 2, 2,
	179, 575, 3, 2, 2, 2, 181, 578, 3, 2, 2, 2, 183, 595, 3, 2, 2, 2, 185,
	597, 3, 2, 2, 2, 187, 600, 3, 2, 2, 2, 189, 606, 3, 2, 2, 2, 191, 611,
	3, 2, 2, 2, 193, 617, 3, 2, 2, 2, 195, 619, 3, 2, 2, 2, 197, 621, 3, 2,
	2, 2, 199, 633, 3, 2, 2, 2, 201, 636, 3, 2, 2, 2, 203, 641, 3, 2, 2, 2,
	205, 647, 3, 2, 2, 2, 207, 651, 3, 2, 2, 2, 209, 693, 3, 2, 2, 2, 211,
	711, 3, 2, 2, 2, 213, 725, 3, 2, 2, 2, 215, 739, 3, 2, 2, 2, 217, 741,
	3, 2, 2, 2, 219, 745, 3, 2, 2, 2, 221, 760, 3, 2, 2, 2, 223, 777, 3, 2,
	2, 2, 225, 779, 3, 2, 2, 2, 227, 789, 3, 2, 2, 2, 229, 230, 7, 99, 2, 2,
	230, 231, 7, 112, 2, 2, 231, 232, 7, 102, 2, 2, 232, 4, 3, 2, 2, 2, 233,
	234, 7, 99, 2, 2, 234, 235, 7, 117, 2, 2, 235, 6, 3, 2, 2, 2, 236, 237,
	7, 99, 2, 2, 237, 238, 7, 118, 2, 2, 238, 239, 7, 118, 2, 2, 239, 240,
	7, 116, 2, 2, 240, 241, 7, 107, 2, 2, 241, 242, 7, 100, 2, 2, 242, 243,
	7, 119, 2, 2, 243, 244, 7, 118, 2, 2, 244, 245, 7, 103, 2, 2, 245, 8, 3,
	2, 2, 2, 246, 247, 7, 100, 2, 2, 247, 248, 7, 116, 2, 2, 248, 249, 7, 103,
	2, 2, 249, 250, 7, 99, 2, 2, 250, 251, 7, 109, 2, 2, 251, 10, 3, 2, 2,
	2, 252, 253, 7, 101, 2, 2, 253, 254, 7, 113, 2, 2, 254, 255, 7, 112, 2,
	2, 255, 256, 7, 117, 2, 2, 256, 257, 7, 118, 2, 2, 257, 12, 3, 2, 2, 2,
	258, 259, 7, 101, 2, 2, 259, 260, 7, 113, 2, 2, 260, 261, 7, 112, 2, 2,
	261, 262, 7, 118, 2, 2, 262, 263, 7, 107, 2, 2, 263, 264, 7, 112, 2, 2,
	264, 265, 7, 119, 2, 2, 265, 266, 7, 103, 2, 2, 266, 14, 3, 2, 2, 2, 267,
	268, 7, 103, 2, 2, 268, 269, 7, 110, 2, 2, 269, 270, 7, 117, 2, 2, 270,
	271, 7, 103, 2, 2, 271, 16, 3, 2, 2, 2, 272, 273, 7, 103, 2, 2, 273, 274,
	7, 112, 2, 2, 274, 275, 7, 119, 2, 2, 275, 276, 7, 111, 2, 2, 276, 18,
	3, 2, 2, 2, 277, 278, 7, 104, 2, 2, 278, 279, 7, 99, 2, 2, 279, 280, 7,
	110, 2, 2, 280, 281, 7, 117, 2, 2, 281, 282, 7, 103, 2, 2, 282, 20, 3,
	2, 2, 2, 283, 284, 7, 104, 2, 2, 284, 285, 7, 113, 2, 2, 285, 286, 7, 116,
	2, 2, 286, 22, 3, 2, 2, 2, 287, 288, 7, 104, 2, 2, 288, 289, 7, 119, 2,
	2, 289, 290, 7, 112, 2, 2, 290, 291, 7, 101, 2, 2, 291, 24, 3, 2, 2, 2,
	292, 293, 7, 107, 2, 2, 293, 294, 7, 104, 2, 2, 294, 26, 3, 2, 2, 2, 295,
	296, 7, 107, 2, 2, 296, 297, 7, 111, 2, 2, 297, 298, 7, 114, 2, 2, 298,
	299, 7, 113, 2, 2, 299, 300, 7, 116, 2, 2, 300, 301, 7, 118, 2, 2, 301,
	28, 3, 2, 2, 2, 302, 303, 7, 107, 2, 2, 303, 304, 7, 112, 2, 2, 304, 30,
	3, 2, 2, 2, 305, 306, 7, 107, 2, 2, 306, 307, 7, 112, 2, 2, 307, 308, 7,
	118, 2, 2, 308, 309, 7, 103, 2, 2, 309, 310, 7, 116, 2, 2, 310, 311, 7,
	104, 2, 2, 311, 312, 7, 99, 2, 2, 312, 313, 7, 101, 2, 2, 313, 314, 7,
	103, 2, 2, 314, 32, 3, 2, 2, 2, 315, 316, 7, 107, 2, 2, 316, 317, 7, 117,
	2, 2, 317, 34, 3, 2, 2, 2, 318, 319, 7, 111, 2, 2, 319, 320, 7, 99, 2,
	2, 320, 321, 7, 118, 2, 2, 321, 322, 7, 101, 2, 2, 322, 323, 7, 106, 2,
	2, 323, 36, 3, 2, 2, 2, 324, 325, 7, 112, 2, 2, 325, 326, 7, 113, 2, 2,
	326, 327, 7, 118, 2, 2, 327, 38, 3, 2, 2, 2, 328, 329, 7, 112, 2, 2, 329,
	330, 7, 119, 2, 2, 330, 331, 7, 110, 2, 2, 331, 332, 7, 110, 2, 2, 332,
	40, 3, 2, 2, 2, 333, 334, 7, 113, 2, 2, 334, 335, 7, 116, 2, 2, 335, 42,
	3, 2, 2, 2, 336, 337, 7, 114, 2, 2, 337, 338, 7, 99, 2, 2, 338, 339, 7,
	101, 2, 2, 339, 340, 7, 109, 2, 2, 340, 341, 7, 99, 2, 2, 341, 342, 7,
	105, 2, 2, 342, 343, 7, 103, 2, 2, 343, 44, 3, 2, 2, 2, 344, 345, 7, 116,
	2, 2, 345, 346, 7, 103, 2, 2, 346, 347, 7, 114, 2, 2, 347, 348, 7, 103,
	2, 2, 348, 349, 7, 99, 2, 2, 349, 350, 7, 118, 2, 2, 350, 46, 3, 2, 2,
	2, 351, 352, 7, 116, 2, 2, 352, 353, 7, 103, 2, 2, 353, 354, 7, 118, 2,
	2, 354, 355, 7, 119, 2, 2, 355, 356, 7, 116, 2, 2, 356, 357, 7, 112, 2,
	2, 357, 48, 3, 2, 2, 2, 358, 359, 7, 117, 2, 2, 359, 360, 7, 118, 2, 2,
	360, 361, 7, 116, 2, 2, 361, 362, 7, 119, 2, 2, 362, 363, 7, 101, 2, 2,
	363, 364, 7, 118, 2, 2, 364, 50, 3, 2, 2, 2, 365, 366, 7, 118, 2, 2, 366,
	367, 7, 116, 2, 2, 367, 368, 7, 119, 2, 2, 368, 369, 7, 103, 2, 2, 369,
	52, 3, 2, 2, 2, 370, 371, 7, 118, 2, 2, 371, 372, 7, 123, 2, 2, 372, 373,
	7, 114, 2, 2, 373, 374, 7, 103, 2, 2, 374, 54, 3, 2, 2, 2, 375, 376, 7,
	120, 2, 2, 376, 377, 7, 99, 2, 2, 377, 378, 7, 116, 2, 2, 378, 56, 3, 2,
	2, 2, 379, 380, 7, 121, 2, 2, 380, 381, 7, 106, 2, 2, 381, 382, 7, 107,
	2, 2, 382, 383, 7, 110, 2, 2, 383, 384, 7, 103, 2, 2, 384, 58, 3, 2, 2,
	2, 385, 386, 7, 122, 2, 2, 386, 387, 7, 113, 2, 2, 387, 388, 7, 116, 2,
	2, 388, 60, 3, 2, 2, 2, 389, 390, 7, 48, 2, 2, 390, 62, 3, 2, 2, 2, 391,
	392, 7, 125, 2, 2, 392, 64, 3, 2, 2, 2, 393, 394, 7, 42, 2, 2, 394, 66,
	3, 2, 2, 2, 395, 396, 7, 93, 2, 2, 396, 68, 3, 2, 2, 2, 397, 398, 7, 127,
	2, 2, 398, 70, 3, 2, 2, 2, 399, 400, 7, 43, 2, 2, 400, 72, 3, 2, 2, 2,
	401, 402, 7, 95, 2, 2, 402, 74, 3, 2, 2, 2, 403, 404, 7, 46, 2, 2, 404,
	76, 3, 2, 2, 2, 405, 406, 7, 60, 2, 2, 406, 78, 3, 2, 2, 2, 407, 408, 7,
	61, 2, 2, 408, 80, 3, 2, 2, 2, 409, 410, 7, 62, 2, 2, 410, 82, 3, 2, 2,
	2, 411, 412, 7, 64, 2, 2, 412, 84, 3, 2, 2, 2, 413, 414, 7, 97, 2, 2, 414,
	86, 3, 2, 2, 2, 415, 416, 7, 35, 2, 2, 416, 88, 3, 2, 2, 2, 417, 418, 7,
	65, 2, 2, 418, 90, 3, 2, 2, 2, 419, 420, 7, 66, 2, 2, 420, 92, 3, 2, 2,
	2, 421, 422, 7, 40, 2, 2, 422, 94, 3, 2, 2, 2, 423, 424, 7, 47, 2, 2, 424,
	96, 3, 2, 2, 2, 425, 426, 7, 63, 2, 2, 426, 98, 3, 2, 2, 2, 427, 428, 7,
	126, 2, 2, 428, 100, 3, 2, 2, 2, 429, 430, 7, 49, 2, 2, 430, 102, 3, 2,
	2, 2, 431, 432, 7, 45, 2, 2, 432, 104, 3, 2, 2, 2, 433, 434, 7, 44, 2,
	2, 434, 106, 3, 2, 2, 2, 435, 436, 7, 39, 2, 2, 436, 108, 3, 2, 2, 2, 437,
	438, 7, 96, 2, 2, 438, 110, 3, 2, 2, 2, 439, 440, 7, 128, 2, 2, 440, 112,
	3, 2, 2, 2, 441, 442, 7, 38, 2, 2, 442, 114, 3, 2, 2, 2, 443, 444, 7, 60,
	2, 2, 444, 445, 7, 63, 2, 2, 445, 116, 3, 2, 2, 2, 446, 447, 7, 63, 2,
	2, 447, 448, 7, 64, 2, 2, 448, 118, 3, 2, 2, 2, 449, 450, 7, 47, 2, 2,
	450, 451, 7, 64, 2, 2, 451, 120, 3, 2, 2, 2, 452, 453, 7, 48, 2, 2, 453,
	454, 7, 48, 2, 2, 454, 122, 3, 2, 2, 2, 455, 456, 7, 62, 2, 2, 456, 457,
	7, 48, 2, 2, 457, 458, 7, 48, 2, 2, 458, 124, 3, 2, 2, 2, 459, 460, 7,
	48, 2, 2, 460, 461, 7, 48, 2, 2, 461, 462, 7, 62, 2, 2, 462, 126, 3, 2,
	2, 2, 463, 464, 7, 48, 2, 2, 464, 465, 7, 48, 2, 2, 465, 466, 7, 48, 2,
	2, 466, 128, 3, 2, 2, 2, 467, 471, 9, 2, 2, 2, 468, 470, 9, 3, 2, 2, 469,
	468, 3, 2, 2, 2, 470, 473, 3, 2, 2, 2, 471, 469, 3, 2, 2, 2, 471, 472,
	3, 2, 2, 2, 472, 130, 3, 2, 2, 2, 473, 471, 3, 2, 2, 2, 474, 476, 5, 133,
	67, 2, 475, 477, 5, 137, 69, 2, 476, 475, 3, 2, 2, 2, 476, 477, 3, 2, 2,
	2, 477, 132, 3, 2, 2, 2, 478, 480, 9, 4, 2, 2, 479, 478, 3, 2, 2, 2, 480,
	134, 3, 2, 2, 2, 481, 484, 9, 5, 2, 2, 482, 484, 5, 133, 67, 2, 483, 481,
	3, 2, 2, 2, 483, 482, 3, 2, 2, 2, 484, 136, 3, 2, 2, 2, 485, 487, 5, 135,
	68, 2, 486, 485, 3, 2, 2, 2, 487, 488, 3, 2, 2, 2, 488, 486, 3, 2, 2, 2,
	488, 489, 3, 2, 2, 2, 489, 138, 3, 2, 2, 2, 490, 492, 9, 6, 2, 2, 491,
	490, 3, 2, 2, 2, 492, 140, 3, 2, 2, 2, 493, 495, 9, 7, 2, 2, 494, 493,
	3, 2, 2, 2, 495, 142, 3, 2, 2, 2, 496, 497, 5, 113, 57, 2, 497, 498, 5,
	167, 84, 2, 498, 144, 3, 2, 2, 2, 499, 502, 5, 51, 26, 2, 500, 502, 5,
	19, 10, 2, 501, 499, 3, 2, 2, 2, 501, 500, 3, 2, 2, 2, 502, 146, 3, 2,
	2, 2, 503, 504, 5, 39, 20, 2, 504, 148, 3, 2, 2, 2, 505, 506, 7, 50, 2,
	2, 506, 507, 7, 100, 2, 2, 507, 508, 3, 2, 2, 2, 508, 510, 5, 151, 76,
	2, 509, 511, 5, 155, 78, 2, 510, 509, 3, 2, 2, 2, 510, 511, 3, 2, 2, 2,
	511, 150, 3, 2, 2, 2, 512, 513, 9, 8, 2, 2, 513, 152, 3, 2, 2, 2, 514,
	517, 5, 151, 76, 2, 515, 517, 7, 97, 2, 2, 516, 514, 3, 2, 2, 2, 516, 515,
	3, 2, 2, 2, 517, 154, 3, 2, 2, 2, 518, 520, 5, 153, 77, 2, 519, 518, 3,
	2, 2, 2, 520, 521, 3, 2, 2, 2, 521, 519, 3, 2, 2, 2, 521, 522, 3, 2, 2,
	2, 522, 156, 3, 2, 2, 2, 523, 524, 7, 50, 2, 2, 524, 525, 7, 113, 2, 2,
	525, 526, 3, 2, 2, 2, 526, 528, 5, 159, 80, 2, 527, 529, 5, 163, 82, 2,
	528, 527, 3, 2, 2, 2, 528, 529, 3, 2, 2, 2, 529, 158, 3, 2, 2, 2, 530,
	531, 9, 9, 2, 2, 531, 160, 3, 2, 2, 2, 532, 535, 5, 159, 80, 2, 533, 535,
	7, 97, 2, 2, 534, 532, 3, 2, 2, 2, 534, 533, 3, 2, 2, 2, 535, 162, 3, 2,
	2, 2, 536, 538, 5, 161, 81, 2, 537, 536, 3, 2, 2, 2, 538, 539, 3, 2, 2,
	2, 539, 537, 3, 2, 2, 2, 539, 540, 3, 2, 2, 2, 540, 164, 3, 2, 2, 2, 541,
	545, 9, 10, 2, 2, 542, 544, 9, 11, 2, 2, 543, 542, 3, 2, 2, 2, 544, 547,
	3, 2, 2, 2, 545, 543, 3, 2, 2, 2, 545, 546, 3, 2, 2, 2, 546, 166, 3, 2,
	2, 2, 547, 545, 3, 2, 2, 2, 548, 550, 9, 10, 2, 2, 549, 548, 3, 2, 2, 2,
	550, 551, 3, 2, 2, 2, 551, 549, 3, 2, 2, 2, 551, 552, 3, 2, 2, 2, 552,
	168, 3, 2, 2, 2, 553, 554, 9, 10, 2, 2, 554, 170, 3, 2, 2, 2, 555, 558,
	5, 169, 85, 2, 556, 558, 7, 97, 2, 2, 557, 555, 3, 2, 2, 2, 557, 556, 3,
	2, 2, 2, 558, 172, 3, 2, 2, 2, 559, 561, 5, 171, 86, 2, 560, 559, 3, 2,
	2, 2, 561, 562, 3, 2, 2, 2, 562, 560, 3, 2, 2, 2, 562, 563, 3, 2, 2, 2,
	563, 174, 3, 2, 2, 2, 564, 565, 7, 50, 2, 2, 565, 566, 7, 122, 2, 2, 566,
	567, 3, 2, 2, 2, 567, 569, 5, 177, 89, 2, 568, 570, 5, 181, 91, 2, 569,
	568, 3, 2, 2, 2, 569, 570, 3, 2, 2, 2, 570, 176, 3, 2, 2, 2, 571, 572,
	9, 12, 2, 2, 572, 178, 3, 2, 2, 2, 573, 576, 5, 177, 89, 2, 574, 576, 7,
	97, 2, 2, 575, 573, 3, 2, 2, 2, 575, 574, 3, 2, 2, 2, 576, 180, 3, 2, 2,
	2, 577, 579, 5, 179, 90, 2, 578, 577, 3, 2, 2, 2, 579, 580, 3, 2, 2, 2,
	580, 578, 3, 2, 2, 2, 580, 581, 3, 2, 2, 2, 581, 182, 3, 2, 2, 2, 582,
	584, 5, 165, 83, 2, 583, 585, 5, 185, 93, 2, 584, 583, 3, 2, 2, 2, 584,
	585, 3, 2, 2, 2, 585, 587, 3, 2, 2, 2, 586, 588, 5, 187, 94, 2, 587, 586,
	3, 2, 2, 2, 587, 588, 3, 2, 2, 2, 588, 596, 3, 2, 2, 2, 589, 591, 5, 175,
	88, 2, 590, 592, 5, 189, 95, 2, 591, 590, 3, 2, 2, 2, 591, 592, 3, 2, 2,
	2, 592, 593, 3, 2, 2, 2, 593, 594, 5, 191, 96, 2, 594, 596, 3, 2, 2, 2,
	595, 582, 3, 2, 2, 2, 595, 589, 3, 2, 2, 2, 596, 184, 3, 2, 2, 2, 597,
	598, 7, 48, 2, 2, 598, 599, 5, 165, 83, 2, 599, 186, 3, 2, 2, 2, 600, 602,
	5, 193, 97, 2, 601, 603, 5, 197, 99, 2, 602, 601, 3, 2, 2, 2, 602, 603,
	3, 2, 2, 2, 603, 604, 3, 2, 2, 2, 604, 605, 5, 165, 83, 2, 605, 188, 3,
	2, 2, 2, 606, 607, 7, 48, 2, 2, 607, 609, 5, 177, 89, 2, 608, 610, 5, 181,
	91, 2, 609, 608, 3, 2, 2, 2, 609, 610, 3, 2, 2, 2, 610, 190, 3, 2, 2, 2,
	611, 613, 5, 195, 98, 2, 612, 614, 5, 197, 99, 2, 613, 612, 3, 2, 2, 2,
	613, 614, 3, 2, 2, 2, 614, 615, 3, 2, 2, 2, 615, 616, 5, 165, 83, 2, 616,
	192, 3, 2, 2, 2, 617, 618, 9, 13, 2, 2, 618, 194, 3, 2, 2, 2, 619, 620,
	9, 14, 2, 2, 620, 196, 3, 2, 2, 2, 621, 622, 9, 15, 2, 2, 622, 198, 3,
	2, 2, 2, 623, 625, 7, 36, 2, 2, 624, 626, 5, 201, 101, 2, 625, 624, 3,
	2, 2, 2, 625, 626, 3, 2, 2, 2, 626, 627, 3, 2, 2, 2, 627, 634, 7, 36, 2,
	2, 628, 630, 7, 41, 2, 2, 629, 631, 5, 203, 102, 2, 630, 629, 3, 2, 2,
	2, 630, 631, 3, 2, 2, 2, 631, 632, 3, 2, 2, 2, 632, 634, 7, 41, 2, 2, 633,
	623, 3, 2, 2, 2, 633, 628, 3, 2, 2, 2, 634, 200, 3, 2, 2, 2, 635, 637,
	5, 205, 103, 2, 636, 635, 3, 2, 2, 2, 637, 638, 3, 2, 2, 2, 638, 636, 3,
	2, 2, 2, 638, 639, 3, 2, 2, 2, 639, 202, 3, 2, 2, 2, 640, 642, 5, 207,
	104, 2, 641, 640, 3, 2, 2, 2, 642, 643, 3, 2, 2, 2, 643, 641, 3, 2, 2,
	2, 643, 644, 3, 2, 2, 2, 644, 204, 3, 2, 2, 2, 645, 648, 5, 209, 105, 2,
	646, 648, 10, 16, 2, 2, 647, 645, 3, 2, 2, 2, 647, 646, 3, 2, 2, 2, 648,
	206, 3, 2, 2, 2, 649, 652, 5, 209, 105, 2, 650, 652, 10, 17, 2, 2, 651,
	649, 3, 2, 2, 2, 651, 650, 3, 2, 2, 2, 652, 208, 3, 2, 2, 2, 653, 654,
	7, 94, 2, 2, 654, 694, 9, 18, 2, 2, 655, 656, 7, 94, 2, 2, 656, 657, 7,
	122, 2, 2, 657, 658, 3, 2, 2, 2, 658, 659, 5, 177, 89, 2, 659, 660, 5,
	177, 89, 2, 660, 694, 3, 2, 2, 2, 661, 662, 7, 94, 2, 2, 662, 663, 7, 119,
	2, 2, 663, 664, 3, 2, 2, 2, 664, 665, 5, 177, 89, 2, 665, 666, 5, 177,
	89, 2, 666, 667, 5, 177, 89, 2, 667, 668, 5, 177, 89, 2, 668, 694, 3, 2,
	2, 2, 669, 670, 7, 94, 2, 2, 670, 671, 7, 119, 2, 2, 671, 672, 3, 2, 2,
	2, 672, 673, 7, 125, 2, 2, 673, 674, 5, 177, 89, 2, 674, 675, 5, 177, 89,
	2, 675, 676, 5, 177, 89, 2, 676, 677, 5, 177, 89, 2, 677, 678, 7, 127,
	2, 2, 678, 694, 3, 2, 2, 2, 679, 680, 7, 94, 2, 2, 680, 681, 7, 119, 2,
	2, 681, 682, 3, 2, 2, 2, 682, 683, 7, 125, 2, 2, 683, 684, 5, 177, 89,
	2, 684, 685, 5, 177, 89, 2, 685, 686, 5, 177, 89, 2, 686, 687, 5, 177,
	89, 2, 687, 688, 5, 177, 89, 2, 688, 689, 5, 177, 89, 2, 689, 690, 5, 177,
	89, 2, 690, 691, 5, 177, 89, 2, 691, 692, 7, 127, 2, 2, 692, 694, 3, 2,
	2, 2, 693, 653, 3, 2, 2, 2, 693, 655, 3, 2, 2, 2, 693, 661, 3, 2, 2, 2,
	693, 669, 3, 2, 2, 2, 693, 679, 3, 2, 2, 2, 694, 210, 3, 2, 2, 2, 695,
	699, 7, 36, 2, 2, 696, 698, 5, 213, 107, 2, 697, 696, 3, 2, 2, 2, 698,
	701, 3, 2, 2, 2, 699, 697, 3, 2, 2, 2, 699, 700, 3, 2, 2, 2, 700, 702,
	3, 2, 2, 2, 701, 699, 3, 2, 2, 2, 702, 712, 7, 36, 2, 2, 703, 707, 7, 41,
	2, 2, 704, 706, 5, 215, 108, 2, 705, 704, 3, 2, 2, 2, 706, 709, 3, 2, 2,
	2, 707, 705, 3, 2, 2, 2, 707, 708, 3, 2, 2, 2, 708, 710, 3, 2, 2, 2, 709,
	707, 3, 2, 2, 2, 710, 712, 7, 41, 2, 2, 711, 695, 3, 2, 2, 2, 711, 703,
	3, 2, 2, 2, 712, 212, 3, 2, 2, 2, 713, 714, 7, 94, 2, 2, 714, 715, 7, 125,
	2, 2, 715, 718, 3, 2, 2, 2, 716, 719, 5, 211, 106, 2, 717, 719, 5, 213,
	107, 2, 718, 716, 3, 2, 2, 2, 718, 717, 3, 2, 2, 2, 719, 720, 3, 2, 2,
	2, 720, 718, 3, 2, 2, 2, 720, 721, 3, 2, 2, 2, 721, 722, 3, 2, 2, 2, 722,
	723, 7, 127, 2, 2, 723, 726, 3, 2, 2, 2, 724, 726, 5, 205, 103, 2, 725,
	713, 3, 2, 2, 2, 725, 724, 3, 2, 2, 2, 726, 214, 3, 2, 2, 2, 727, 728,
	7, 94, 2, 2, 728, 729, 7, 125, 2, 2, 729, 732, 3, 2, 2, 2, 730, 733, 5,
	211, 106, 2, 731, 733, 5, 215, 108, 2, 732, 730, 3, 2, 2, 2, 732, 731,
	3, 2, 2, 2, 733, 734, 3, 2, 2, 2, 734, 732, 3, 2, 2, 2, 734, 735, 3, 2,
	2, 2, 735, 736, 3, 2, 2, 2, 736, 737, 7, 127, 2, 2, 737, 740, 3, 2, 2,
	2, 738, 740, 5, 207, 104, 2, 739, 727, 3, 2, 2, 2, 739, 738, 3, 2, 2, 2,
	740, 216, 3, 2, 2, 2, 741, 742, 9, 19, 2, 2, 742, 743, 3, 2, 2, 2, 743,
	744, 8, 109, 2, 2, 744, 218, 3, 2, 2, 2, 745, 746, 7, 49, 2, 2, 746, 747,
	7, 44, 2, 2, 747, 752, 3, 2, 2, 2, 748, 751, 5, 219, 110, 2, 749, 751,
	11, 2, 2, 2, 750, 748, 3, 2, 2, 2, 750, 749, 3, 2, 2, 2, 751, 754, 3, 2,
	2, 2, 752, 753, 3, 2, 2, 2, 752, 750, 3, 2, 2, 2, 753, 755, 3, 2, 2, 2,
	754, 752, 3, 2, 2, 2, 755, 756, 7, 44, 2, 2, 756, 757, 7, 49, 2, 2, 757,
	758, 3, 2, 2, 2, 758, 759, 8, 110, 3, 2, 759, 220, 3, 2, 2, 2, 760, 761,
	7, 49, 2, 2, 761, 762, 7, 49, 2, 2, 762, 770, 3, 2, 2, 2, 763, 767, 10,
	20, 2, 2, 764, 766, 10, 21, 2, 2, 765, 764, 3, 2, 2, 2, 766, 769, 3, 2,
	2, 2, 767, 765, 3, 2, 2, 2, 767, 768, 3, 2, 2, 2, 768, 771, 3, 2, 2, 2,
	769, 767, 3, 2, 2, 2, 770, 763, 3, 2, 2, 2, 770, 771, 3, 2, 2, 2, 771,
	772, 3, 2, 2, 2, 772, 773, 8, 111, 3, 2, 773, 222, 3, 2, 2, 2, 774, 778,
	7, 12, 2, 2, 775, 776, 7, 15, 2, 2, 776, 778, 7, 12, 2, 2, 777, 774, 3,
	2, 2, 2, 777, 775, 3, 2, 2, 2, 778, 224, 3, 2, 2, 2, 779, 780, 7, 49, 2,
	2, 780, 781, 7, 49, 2, 2, 781, 782, 7, 49, 2, 2, 782, 786, 3, 2, 2, 2,
	783, 785, 10, 21, 2, 2, 784, 783, 3, 2, 2, 2, 785, 788, 3, 2, 2, 2, 786,
	784, 3, 2, 2, 2, 786, 787, 3, 2, 2, 2, 787, 226, 3, 2, 2, 2, 788, 786,
	3, 2, 2, 2, 789, 790, 7, 49, 2, 2, 790, 791, 7, 49, 2, 2, 791, 792, 7,
	62, 2, 2, 792, 796, 3, 2, 2, 2, 793, 795, 10, 21, 2, 2, 794, 793, 3, 2,
	2, 2, 795, 798, 3, 2, 2, 2, 796, 794, 3, 2, 2, 2, 796, 797, 3, 2, 2, 2,
	797, 228, 3, 2, 2, 2, 798, 796, 3, 2, 2, 2, 55, 2, 471, 476, 479, 483,
	488, 491, 494, 501, 510, 516, 521, 528, 534, 539, 545, 551, 557, 562, 569,
	575, 580, 584, 587, 591, 595, 602, 609, 613, 625, 630, 633, 638, 643, 647,
	651, 693, 699, 707, 711, 718, 720, 725, 732, 734, 739, 750, 752, 767, 770,
	777, 786, 796, 4, 8, 2, 2, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'and'", "'as'", "'attribute'", "'break'", "'const'", "'continue'",
	"'else'", "'enum'", "'false'", "'for'", "'func'", "'if'", "'import'", "'in'",
	"'interface'", "'is'", "'match'", "'not'", "'null'", "'or'", "'package'",
	"'repeat'", "'return'", "'struct'", "'true'", "'type'", "'var'", "'while'",
	"'xor'", "'.'", "'{'", "'('", "'['", "'}'", "')'", "']'", "','", "':'",
	"';'", "'<'", "'>'", "'_'", "'!'", "'?'", "'@'", "'&'", "'-'", "'='", "'|'",
	"'/'", "'+'", "'*'", "'%'", "'^'", "'~'", "'$'",
}

var lexerSymbolicNames = []string{
	"", "KEYWORD_AND", "KEYWORD_AS", "KEYWORD_ATTRIBUTE", "KEYWORD_BREAK",
	"KEYWORD_CONST", "KEYWORD_CONTINUE", "KEYWORD_ELSE", "KEYWORD_ENUM", "KEYWORD_FALSE",
	"KEYWORD_FOR", "KEYWORD_FUNC", "KEYWORD_IF", "KEYWORD_IMPORT", "KEYWORD_IN",
	"KEYWORD_INTERFACE", "KEYWORD_IS", "KEYWORD_MATCH", "KEYWORD_NOT", "KEYWORD_NULL",
	"KEYWORD_OR", "KEYWORD_PACKAGE", "KEYWORD_REPEATE", "KEYWORD_RETURN", "KEYWORD_STRUCT",
	"KEYWORD_TRUE", "KEYWORD_TYPE", "KEYWORD_VAR", "KEYWORD_WHILE", "KEYWORD_XOR",
	"DOT", "LCURLY", "LPAREN", "LBRACK", "RCURLY", "RPAREN", "RBRACK", "COMMA",
	"COLON", "SEMI", "LT", "GT", "UNDERSCORE", "BANG", "QUESTION", "AT", "AND",
	"SUB", "EQUAL", "OR", "DIV", "ADD", "MUL", "MOD", "CARET", "TILDE", "DOLLER",
	"COLON_EQUAL", "RIGHT_RIGHT_ARROWS", "RIGHT_ARROW", "DOT_DOT", "LT_DOT_DOT",
	"DOT_DOT_LT", "ELLIPSIS", "TypeName", "Identifier", "Operator_head_other",
	"Operator_following_character", "Implicit_parameter_name", "BoolLiteral",
	"NullLiteral", "BinaryLiteral", "OctalLiteral", "DecimalLiteral", "PureDecimalDigits",
	"HexadecimalLiteral", "FloatLiteral", "StaticStringLiteral", "InterpolatedStringLiteral",
	"WS", "DelimitedComment", "LineComment", "EOL", "LineDocument", "FollowingLineDocument",
}

var lexerRuleNames = []string{
	"KEYWORD_AND", "KEYWORD_AS", "KEYWORD_ATTRIBUTE", "KEYWORD_BREAK", "KEYWORD_CONST",
	"KEYWORD_CONTINUE", "KEYWORD_ELSE", "KEYWORD_ENUM", "KEYWORD_FALSE", "KEYWORD_FOR",
	"KEYWORD_FUNC", "KEYWORD_IF", "KEYWORD_IMPORT", "KEYWORD_IN", "KEYWORD_INTERFACE",
	"KEYWORD_IS", "KEYWORD_MATCH", "KEYWORD_NOT", "KEYWORD_NULL", "KEYWORD_OR",
	"KEYWORD_PACKAGE", "KEYWORD_REPEATE", "KEYWORD_RETURN", "KEYWORD_STRUCT",
	"KEYWORD_TRUE", "KEYWORD_TYPE", "KEYWORD_VAR", "KEYWORD_WHILE", "KEYWORD_XOR",
	"DOT", "LCURLY", "LPAREN", "LBRACK", "RCURLY", "RPAREN", "RBRACK", "COMMA",
	"COLON", "SEMI", "LT", "GT", "UNDERSCORE", "BANG", "QUESTION", "AT", "AND",
	"SUB", "EQUAL", "OR", "DIV", "ADD", "MUL", "MOD", "CARET", "TILDE", "DOLLER",
	"COLON_EQUAL", "RIGHT_RIGHT_ARROWS", "RIGHT_ARROW", "DOT_DOT", "LT_DOT_DOT",
	"DOT_DOT_LT", "ELLIPSIS", "TypeName", "Identifier", "Identifier_head",
	"Identifier_character", "Identifier_characters", "Operator_head_other",
	"Operator_following_character", "Implicit_parameter_name", "BoolLiteral",
	"NullLiteral", "BinaryLiteral", "Binary_digit", "BinaryLiteral_character",
	"BinaryLiteral_characters", "OctalLiteral", "Octal_digit", "OctalLiteral_character",
	"OctalLiteral_characters", "DecimalLiteral", "PureDecimalDigits", "Decimal_digit",
	"Decimal_literal_character", "Decimal_literal_characters", "HexadecimalLiteral",
	"Hexadecimal_digit", "Hexadecimal_literal_character", "Hexadecimal_literal_characters",
	"FloatLiteral", "Decimal_fraction", "Decimal_exponent", "Hexadecimal_fraction",
	"Hexadecimal_exponent", "Floating_point_e", "Floating_point_p", "Sign",
	"StaticStringLiteral", "Double_quoted_text", "Single_quoted_text", "Double_quoted_text_item",
	"Single_quoted_text_item", "Escaped_character", "InterpolatedStringLiteral",
	"Double_interpolated_text_item", "Single_interpolated_text_item", "WS",
	"DelimitedComment", "LineComment", "EOL", "LineDocument", "FollowingLineDocument",
}

type MojoLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewMojoLexer(input antlr.CharStream) *MojoLexer {

	l := new(MojoLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "MojoLexer.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// MojoLexer tokens.
const (
	MojoLexerKEYWORD_AND                  = 1
	MojoLexerKEYWORD_AS                   = 2
	MojoLexerKEYWORD_ATTRIBUTE            = 3
	MojoLexerKEYWORD_BREAK                = 4
	MojoLexerKEYWORD_CONST                = 5
	MojoLexerKEYWORD_CONTINUE             = 6
	MojoLexerKEYWORD_ELSE                 = 7
	MojoLexerKEYWORD_ENUM                 = 8
	MojoLexerKEYWORD_FALSE                = 9
	MojoLexerKEYWORD_FOR                  = 10
	MojoLexerKEYWORD_FUNC                 = 11
	MojoLexerKEYWORD_IF                   = 12
	MojoLexerKEYWORD_IMPORT               = 13
	MojoLexerKEYWORD_IN                   = 14
	MojoLexerKEYWORD_INTERFACE            = 15
	MojoLexerKEYWORD_IS                   = 16
	MojoLexerKEYWORD_MATCH                = 17
	MojoLexerKEYWORD_NOT                  = 18
	MojoLexerKEYWORD_NULL                 = 19
	MojoLexerKEYWORD_OR                   = 20
	MojoLexerKEYWORD_PACKAGE              = 21
	MojoLexerKEYWORD_REPEATE              = 22
	MojoLexerKEYWORD_RETURN               = 23
	MojoLexerKEYWORD_STRUCT               = 24
	MojoLexerKEYWORD_TRUE                 = 25
	MojoLexerKEYWORD_TYPE                 = 26
	MojoLexerKEYWORD_VAR                  = 27
	MojoLexerKEYWORD_WHILE                = 28
	MojoLexerKEYWORD_XOR                  = 29
	MojoLexerDOT                          = 30
	MojoLexerLCURLY                       = 31
	MojoLexerLPAREN                       = 32
	MojoLexerLBRACK                       = 33
	MojoLexerRCURLY                       = 34
	MojoLexerRPAREN                       = 35
	MojoLexerRBRACK                       = 36
	MojoLexerCOMMA                        = 37
	MojoLexerCOLON                        = 38
	MojoLexerSEMI                         = 39
	MojoLexerLT                           = 40
	MojoLexerGT                           = 41
	MojoLexerUNDERSCORE                   = 42
	MojoLexerBANG                         = 43
	MojoLexerQUESTION                     = 44
	MojoLexerAT                           = 45
	MojoLexerAND                          = 46
	MojoLexerSUB                          = 47
	MojoLexerEQUAL                        = 48
	MojoLexerOR                           = 49
	MojoLexerDIV                          = 50
	MojoLexerADD                          = 51
	MojoLexerMUL                          = 52
	MojoLexerMOD                          = 53
	MojoLexerCARET                        = 54
	MojoLexerTILDE                        = 55
	MojoLexerDOLLER                       = 56
	MojoLexerCOLON_EQUAL                  = 57
	MojoLexerRIGHT_RIGHT_ARROWS           = 58
	MojoLexerRIGHT_ARROW                  = 59
	MojoLexerDOT_DOT                      = 60
	MojoLexerLT_DOT_DOT                   = 61
	MojoLexerDOT_DOT_LT                   = 62
	MojoLexerELLIPSIS                     = 63
	MojoLexerTypeName                     = 64
	MojoLexerIdentifier                   = 65
	MojoLexerOperator_head_other          = 66
	MojoLexerOperator_following_character = 67
	MojoLexerImplicit_parameter_name      = 68
	MojoLexerBoolLiteral                  = 69
	MojoLexerNullLiteral                  = 70
	MojoLexerBinaryLiteral                = 71
	MojoLexerOctalLiteral                 = 72
	MojoLexerDecimalLiteral               = 73
	MojoLexerPureDecimalDigits            = 74
	MojoLexerHexadecimalLiteral           = 75
	MojoLexerFloatLiteral                 = 76
	MojoLexerStaticStringLiteral          = 77
	MojoLexerInterpolatedStringLiteral    = 78
	MojoLexerWS                           = 79
	MojoLexerDelimitedComment             = 80
	MojoLexerLineComment                  = 81
	MojoLexerEOL                          = 82
	MojoLexerLineDocument                 = 83
	MojoLexerFollowingLineDocument        = 84
)
