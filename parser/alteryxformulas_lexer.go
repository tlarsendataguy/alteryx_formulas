// Code generated from /Users/tlarsen/Documents/Repositories/alteryx_formulas/grammar/AlteryxFormulas.g4 by ANTLR 4.8. DO NOT EDIT.

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 80, 729,
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
	3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7,
	3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 12, 3, 12,
	3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 16, 3,
	16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27,
	3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3,
	34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36,
	3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3,
	37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38,
	3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3,
	40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3, 42,
	3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3,
	44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 45,
	3, 46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3,
	48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49,
	3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 3,
	51, 3, 51, 3, 52, 3, 52, 3, 52, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 3, 53,
	3, 53, 3, 53, 3, 53, 3, 53, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3, 54, 3,
	55, 3, 55, 3, 55, 3, 55, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57,
	3, 57, 3, 57, 3, 57, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3, 58, 3,
	59, 3, 59, 3, 59, 3, 59, 3, 60, 3, 60, 3, 60, 3, 60, 3, 60, 3, 61, 3, 61,
	3, 61, 3, 61, 3, 62, 3, 62, 3, 62, 3, 63, 3, 63, 3, 63, 3, 63, 3, 64, 3,
	64, 3, 64, 3, 64, 3, 65, 3, 65, 3, 65, 3, 66, 3, 66, 3, 66, 3, 67, 3, 67,
	3, 67, 3, 67, 3, 67, 3, 68, 3, 68, 3, 68, 3, 68, 3, 68, 3, 69, 3, 69, 3,
	69, 3, 69, 3, 69, 3, 69, 3, 69, 3, 70, 3, 70, 3, 70, 3, 70, 3, 70, 3, 70,
	3, 71, 3, 71, 5, 71, 583, 10, 71, 3, 72, 6, 72, 586, 10, 72, 13, 72, 14,
	72, 587, 3, 73, 7, 73, 591, 10, 73, 12, 73, 14, 73, 594, 11, 73, 3, 73,
	3, 73, 6, 73, 598, 10, 73, 13, 73, 14, 73, 599, 3, 74, 3, 74, 3, 74, 3,
	74, 3, 75, 3, 75, 3, 75, 3, 75, 3, 76, 3, 76, 6, 76, 612, 10, 76, 13, 76,
	14, 76, 613, 3, 76, 3, 76, 3, 77, 3, 77, 7, 77, 620, 10, 77, 12, 77, 14,
	77, 623, 11, 77, 3, 77, 3, 77, 3, 78, 3, 78, 7, 78, 629, 10, 78, 12, 78,
	14, 78, 632, 11, 78, 3, 78, 3, 78, 3, 79, 6, 79, 637, 10, 79, 13, 79, 14,
	79, 638, 3, 79, 3, 79, 3, 80, 3, 80, 3, 81, 3, 81, 3, 82, 3, 82, 3, 83,
	3, 83, 3, 84, 3, 84, 3, 85, 3, 85, 3, 86, 3, 86, 3, 87, 3, 87, 3, 88, 3,
	88, 3, 89, 3, 89, 3, 90, 3, 90, 3, 91, 3, 91, 3, 92, 3, 92, 3, 93, 3, 93,
	3, 94, 3, 94, 3, 95, 3, 95, 3, 96, 3, 96, 3, 97, 3, 97, 3, 98, 3, 98, 3,
	99, 3, 99, 3, 100, 3, 100, 3, 101, 3, 101, 3, 102, 3, 102, 3, 103, 3, 103,
	3, 104, 3, 104, 3, 105, 3, 105, 3, 106, 3, 106, 3, 107, 3, 107, 3, 107,
	3, 107, 3, 107, 3, 107, 3, 107, 3, 107, 3, 107, 3, 107, 3, 107, 3, 108,
	3, 108, 3, 108, 3, 108, 3, 108, 3, 108, 3, 108, 3, 108, 3, 108, 3, 108,
	3, 108, 3, 109, 3, 109, 3, 109, 3, 109, 3, 109, 3, 110, 3, 110, 3, 110,
	3, 110, 3, 110, 3, 110, 2, 2, 111, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8,
	15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17,
	33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26,
	51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35,
	69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79, 41, 81, 42, 83, 43, 85, 44,
	87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 97, 50, 99, 51, 101, 52, 103, 53,
	105, 54, 107, 55, 109, 56, 111, 57, 113, 58, 115, 59, 117, 60, 119, 61,
	121, 62, 123, 63, 125, 64, 127, 65, 129, 66, 131, 67, 133, 68, 135, 69,
	137, 70, 139, 71, 141, 72, 143, 73, 145, 74, 147, 75, 149, 76, 151, 77,
	153, 78, 155, 79, 157, 80, 159, 2, 161, 2, 163, 2, 165, 2, 167, 2, 169,
	2, 171, 2, 173, 2, 175, 2, 177, 2, 179, 2, 181, 2, 183, 2, 185, 2, 187,
	2, 189, 2, 191, 2, 193, 2, 195, 2, 197, 2, 199, 2, 201, 2, 203, 2, 205,
	2, 207, 2, 209, 2, 211, 2, 213, 2, 215, 2, 217, 2, 219, 2, 3, 2, 34, 3,
	2, 98, 98, 3, 2, 95, 95, 3, 2, 41, 41, 3, 2, 36, 36, 5, 2, 11, 12, 15,
	15, 34, 34, 4, 2, 67, 67, 99, 99, 4, 2, 68, 68, 100, 100, 4, 2, 69, 69,
	101, 101, 4, 2, 70, 70, 102, 102, 4, 2, 71, 71, 103, 103, 4, 2, 72, 72,
	104, 104, 4, 2, 73, 73, 105, 105, 4, 2, 74, 74, 106, 106, 4, 2, 75, 75,
	107, 107, 4, 2, 76, 76, 108, 108, 4, 2, 77, 77, 109, 109, 4, 2, 78, 78,
	110, 110, 4, 2, 79, 79, 111, 111, 4, 2, 80, 80, 112, 112, 4, 2, 81, 81,
	113, 113, 4, 2, 82, 82, 114, 114, 4, 2, 83, 83, 115, 115, 4, 2, 84, 84,
	116, 116, 4, 2, 85, 85, 117, 117, 4, 2, 86, 86, 118, 118, 4, 2, 87, 87,
	119, 119, 4, 2, 88, 88, 120, 120, 4, 2, 89, 89, 121, 121, 4, 2, 90, 90,
	122, 122, 4, 2, 91, 91, 123, 123, 4, 2, 92, 92, 124, 124, 3, 2, 50, 59,
	2, 705, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3,
	2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17,
	3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2,
	25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2,
	2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2,
	2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2,
	2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3,
	2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63,
	3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2,
	71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2,
	2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2,
	2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2,
	2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101,
	3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3, 2, 2, 2, 2, 107, 3, 2, 2, 2,
	2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2, 113, 3, 2, 2, 2, 2, 115, 3,
	2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119, 3, 2, 2, 2, 2, 121, 3, 2, 2, 2, 2,
	123, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2, 2, 127, 3, 2, 2, 2, 2, 129, 3, 2,
	2, 2, 2, 131, 3, 2, 2, 2, 2, 133, 3, 2, 2, 2, 2, 135, 3, 2, 2, 2, 2, 137,
	3, 2, 2, 2, 2, 139, 3, 2, 2, 2, 2, 141, 3, 2, 2, 2, 2, 143, 3, 2, 2, 2,
	2, 145, 3, 2, 2, 2, 2, 147, 3, 2, 2, 2, 2, 149, 3, 2, 2, 2, 2, 151, 3,
	2, 2, 2, 2, 153, 3, 2, 2, 2, 2, 155, 3, 2, 2, 2, 2, 157, 3, 2, 2, 2, 3,
	221, 3, 2, 2, 2, 5, 223, 3, 2, 2, 2, 7, 225, 3, 2, 2, 2, 9, 227, 3, 2,
	2, 2, 11, 229, 3, 2, 2, 2, 13, 231, 3, 2, 2, 2, 15, 233, 3, 2, 2, 2, 17,
	235, 3, 2, 2, 2, 19, 237, 3, 2, 2, 2, 21, 240, 3, 2, 2, 2, 23, 242, 3,
	2, 2, 2, 25, 245, 3, 2, 2, 2, 27, 248, 3, 2, 2, 2, 29, 250, 3, 2, 2, 2,
	31, 253, 3, 2, 2, 2, 33, 256, 3, 2, 2, 2, 35, 259, 3, 2, 2, 2, 37, 263,
	3, 2, 2, 2, 39, 268, 3, 2, 2, 2, 41, 273, 3, 2, 2, 2, 43, 278, 3, 2, 2,
	2, 45, 284, 3, 2, 2, 2, 47, 292, 3, 2, 2, 2, 49, 297, 3, 2, 2, 2, 51, 309,
	3, 2, 2, 2, 53, 319, 3, 2, 2, 2, 55, 328, 3, 2, 2, 2, 57, 332, 3, 2, 2,
	2, 59, 337, 3, 2, 2, 2, 61, 348, 3, 2, 2, 2, 63, 357, 3, 2, 2, 2, 65, 366,
	3, 2, 2, 2, 67, 377, 3, 2, 2, 2, 69, 381, 3, 2, 2, 2, 71, 387, 3, 2, 2,
	2, 73, 395, 3, 2, 2, 2, 75, 407, 3, 2, 2, 2, 77, 412, 3, 2, 2, 2, 79, 419,
	3, 2, 2, 2, 81, 423, 3, 2, 2, 2, 83, 429, 3, 2, 2, 2, 85, 439, 3, 2, 2,
	2, 87, 443, 3, 2, 2, 2, 89, 450, 3, 2, 2, 2, 91, 454, 3, 2, 2, 2, 93, 458,
	3, 2, 2, 2, 95, 463, 3, 2, 2, 2, 97, 471, 3, 2, 2, 2, 99, 480, 3, 2, 2,
	2, 101, 483, 3, 2, 2, 2, 103, 487, 3, 2, 2, 2, 105, 492, 3, 2, 2, 2, 107,
	500, 3, 2, 2, 2, 109, 506, 3, 2, 2, 2, 111, 510, 3, 2, 2, 2, 113, 515,
	3, 2, 2, 2, 115, 520, 3, 2, 2, 2, 117, 527, 3, 2, 2, 2, 119, 531, 3, 2,
	2, 2, 121, 536, 3, 2, 2, 2, 123, 540, 3, 2, 2, 2, 125, 543, 3, 2, 2, 2,
	127, 547, 3, 2, 2, 2, 129, 551, 3, 2, 2, 2, 131, 554, 3, 2, 2, 2, 133,
	557, 3, 2, 2, 2, 135, 562, 3, 2, 2, 2, 137, 567, 3, 2, 2, 2, 139, 574,
	3, 2, 2, 2, 141, 582, 3, 2, 2, 2, 143, 585, 3, 2, 2, 2, 145, 592, 3, 2,
	2, 2, 147, 601, 3, 2, 2, 2, 149, 605, 3, 2, 2, 2, 151, 609, 3, 2, 2, 2,
	153, 617, 3, 2, 2, 2, 155, 626, 3, 2, 2, 2, 157, 636, 3, 2, 2, 2, 159,
	642, 3, 2, 2, 2, 161, 644, 3, 2, 2, 2, 163, 646, 3, 2, 2, 2, 165, 648,
	3, 2, 2, 2, 167, 650, 3, 2, 2, 2, 169, 652, 3, 2, 2, 2, 171, 654, 3, 2,
	2, 2, 173, 656, 3, 2, 2, 2, 175, 658, 3, 2, 2, 2, 177, 660, 3, 2, 2, 2,
	179, 662, 3, 2, 2, 2, 181, 664, 3, 2, 2, 2, 183, 666, 3, 2, 2, 2, 185,
	668, 3, 2, 2, 2, 187, 670, 3, 2, 2, 2, 189, 672, 3, 2, 2, 2, 191, 674,
	3, 2, 2, 2, 193, 676, 3, 2, 2, 2, 195, 678, 3, 2, 2, 2, 197, 680, 3, 2,
	2, 2, 199, 682, 3, 2, 2, 2, 201, 684, 3, 2, 2, 2, 203, 686, 3, 2, 2, 2,
	205, 688, 3, 2, 2, 2, 207, 690, 3, 2, 2, 2, 209, 692, 3, 2, 2, 2, 211,
	694, 3, 2, 2, 2, 213, 696, 3, 2, 2, 2, 215, 707, 3, 2, 2, 2, 217, 718,
	3, 2, 2, 2, 219, 723, 3, 2, 2, 2, 221, 222, 7, 42, 2, 2, 222, 4, 3, 2,
	2, 2, 223, 224, 7, 43, 2, 2, 224, 6, 3, 2, 2, 2, 225, 226, 7, 44, 2, 2,
	226, 8, 3, 2, 2, 2, 227, 228, 7, 49, 2, 2, 228, 10, 3, 2, 2, 2, 229, 230,
	7, 45, 2, 2, 230, 12, 3, 2, 2, 2, 231, 232, 7, 47, 2, 2, 232, 14, 3, 2,
	2, 2, 233, 234, 7, 63, 2, 2, 234, 16, 3, 2, 2, 2, 235, 236, 7, 64, 2, 2,
	236, 18, 3, 2, 2, 2, 237, 238, 7, 64, 2, 2, 238, 239, 7, 63, 2, 2, 239,
	20, 3, 2, 2, 2, 240, 241, 7, 62, 2, 2, 241, 22, 3, 2, 2, 2, 242, 243, 7,
	62, 2, 2, 243, 244, 7, 63, 2, 2, 244, 24, 3, 2, 2, 2, 245, 246, 7, 35,
	2, 2, 246, 247, 7, 63, 2, 2, 247, 26, 3, 2, 2, 2, 248, 249, 7, 46, 2, 2,
	249, 28, 3, 2, 2, 2, 250, 251, 7, 40, 2, 2, 251, 252, 7, 40, 2, 2, 252,
	30, 3, 2, 2, 2, 253, 254, 7, 126, 2, 2, 254, 255, 7, 126, 2, 2, 255, 32,
	3, 2, 2, 2, 256, 257, 7, 42, 2, 2, 257, 258, 7, 43, 2, 2, 258, 34, 3, 2,
	2, 2, 259, 260, 5, 159, 80, 2, 260, 261, 5, 161, 81, 2, 261, 262, 5, 195,
	98, 2, 262, 36, 3, 2, 2, 2, 263, 264, 5, 159, 80, 2, 264, 265, 5, 163,
	82, 2, 265, 266, 5, 187, 94, 2, 266, 267, 5, 195, 98, 2, 267, 38, 3, 2,
	2, 2, 268, 269, 5, 159, 80, 2, 269, 270, 5, 195, 98, 2, 270, 271, 5, 175,
	88, 2, 271, 272, 5, 185, 93, 2, 272, 40, 3, 2, 2, 2, 273, 274, 5, 159,
	80, 2, 274, 275, 5, 197, 99, 2, 275, 276, 5, 159, 80, 2, 276, 277, 5, 185,
	93, 2, 277, 42, 3, 2, 2, 2, 278, 279, 5, 159, 80, 2, 279, 280, 5, 197,
	99, 2, 280, 281, 5, 159, 80, 2, 281, 282, 5, 185, 93, 2, 282, 283, 7, 52,
	2, 2, 283, 44, 3, 2, 2, 2, 284, 285, 5, 159, 80, 2, 285, 286, 5, 201, 101,
	2, 286, 287, 5, 167, 84, 2, 287, 288, 5, 193, 97, 2, 288, 289, 5, 159,
	80, 2, 289, 290, 5, 171, 86, 2, 290, 291, 5, 167, 84, 2, 291, 46, 3, 2,
	2, 2, 292, 293, 5, 163, 82, 2, 293, 294, 5, 167, 84, 2, 294, 295, 5, 175,
	88, 2, 295, 296, 5, 181, 91, 2, 296, 48, 3, 2, 2, 2, 297, 298, 5, 163,
	82, 2, 298, 299, 5, 173, 87, 2, 299, 300, 5, 159, 80, 2, 300, 301, 5, 193,
	97, 2, 301, 302, 5, 169, 85, 2, 302, 303, 5, 193, 97, 2, 303, 304, 5, 187,
	94, 2, 304, 305, 5, 183, 92, 2, 305, 306, 5, 175, 88, 2, 306, 307, 5, 185,
	93, 2, 307, 308, 5, 197, 99, 2, 308, 50, 3, 2, 2, 2, 309, 310, 5, 163,
	82, 2, 310, 311, 5, 173, 87, 2, 311, 312, 5, 159, 80, 2, 312, 313, 5, 193,
	97, 2, 313, 314, 5, 197, 99, 2, 314, 315, 5, 187, 94, 2, 315, 316, 5, 175,
	88, 2, 316, 317, 5, 185, 93, 2, 317, 318, 5, 197, 99, 2, 318, 52, 3, 2,
	2, 2, 319, 320, 5, 163, 82, 2, 320, 321, 5, 187, 94, 2, 321, 322, 5, 185,
	93, 2, 322, 323, 5, 197, 99, 2, 323, 324, 5, 159, 80, 2, 324, 325, 5, 175,
	88, 2, 325, 326, 5, 185, 93, 2, 326, 327, 5, 195, 98, 2, 327, 54, 3, 2,
	2, 2, 328, 329, 5, 163, 82, 2, 329, 330, 5, 187, 94, 2, 330, 331, 5, 195,
	98, 2, 331, 56, 3, 2, 2, 2, 332, 333, 5, 163, 82, 2, 333, 334, 5, 187,
	94, 2, 334, 335, 5, 195, 98, 2, 335, 336, 5, 173, 87, 2, 336, 58, 3, 2,
	2, 2, 337, 338, 5, 163, 82, 2, 338, 339, 5, 187, 94, 2, 339, 340, 5, 199,
	100, 2, 340, 341, 5, 185, 93, 2, 341, 342, 5, 197, 99, 2, 342, 343, 5,
	203, 102, 2, 343, 344, 5, 187, 94, 2, 344, 345, 5, 193, 97, 2, 345, 346,
	5, 165, 83, 2, 346, 347, 5, 195, 98, 2, 347, 60, 3, 2, 2, 2, 348, 349,
	5, 165, 83, 2, 349, 350, 5, 175, 88, 2, 350, 351, 5, 195, 98, 2, 351, 352,
	5, 197, 99, 2, 352, 353, 5, 159, 80, 2, 353, 354, 5, 185, 93, 2, 354, 355,
	5, 163, 82, 2, 355, 356, 5, 167, 84, 2, 356, 62, 3, 2, 2, 2, 357, 358,
	5, 167, 84, 2, 358, 359, 5, 185, 93, 2, 359, 360, 5, 165, 83, 2, 360, 361,
	5, 195, 98, 2, 361, 362, 5, 203, 102, 2, 362, 363, 5, 175, 88, 2, 363,
	364, 5, 197, 99, 2, 364, 365, 5, 173, 87, 2, 365, 64, 3, 2, 2, 2, 366,
	367, 5, 169, 85, 2, 367, 368, 5, 175, 88, 2, 368, 369, 5, 185, 93, 2, 369,
	370, 5, 165, 83, 2, 370, 371, 5, 195, 98, 2, 371, 372, 5, 197, 99, 2, 372,
	373, 5, 193, 97, 2, 373, 374, 5, 175, 88, 2, 374, 375, 5, 185, 93, 2, 375,
	376, 5, 171, 86, 2, 376, 66, 3, 2, 2, 2, 377, 378, 5, 167, 84, 2, 378,
	379, 5, 205, 103, 2, 379, 380, 5, 189, 95, 2, 380, 68, 3, 2, 2, 2, 381,
	382, 5, 169, 85, 2, 382, 383, 5, 181, 91, 2, 383, 384, 5, 187, 94, 2, 384,
	385, 5, 187, 94, 2, 385, 386, 5, 193, 97, 2, 386, 70, 3, 2, 2, 2, 387,
	388, 5, 171, 86, 2, 388, 389, 5, 167, 84, 2, 389, 390, 5, 197, 99, 2, 390,
	391, 5, 203, 102, 2, 391, 392, 5, 187, 94, 2, 392, 393, 5, 193, 97, 2,
	393, 394, 5, 165, 83, 2, 394, 72, 3, 2, 2, 2, 395, 396, 5, 173, 87, 2,
	396, 397, 5, 167, 84, 2, 397, 398, 5, 205, 103, 2, 398, 399, 5, 197, 99,
	2, 399, 400, 5, 187, 94, 2, 400, 401, 5, 185, 93, 2, 401, 402, 5, 199,
	100, 2, 402, 403, 5, 183, 92, 2, 403, 404, 5, 161, 81, 2, 404, 405, 5,
	167, 84, 2, 405, 406, 5, 193, 97, 2, 406, 74, 3, 2, 2, 2, 407, 408, 5,
	181, 91, 2, 408, 409, 5, 167, 84, 2, 409, 410, 5, 169, 85, 2, 410, 411,
	5, 197, 99, 2, 411, 76, 3, 2, 2, 2, 412, 413, 5, 181, 91, 2, 413, 414,
	5, 167, 84, 2, 414, 415, 5, 185, 93, 2, 415, 416, 5, 171, 86, 2, 416, 417,
	5, 197, 99, 2, 417, 418, 5, 173, 87, 2, 418, 78, 3, 2, 2, 2, 419, 420,
	5, 181, 91, 2, 420, 421, 5, 187, 94, 2, 421, 422, 5, 171, 86, 2, 422, 80,
	3, 2, 2, 2, 423, 424, 5, 181, 91, 2, 424, 425, 5, 187, 94, 2, 425, 426,
	5, 171, 86, 2, 426, 427, 7, 51, 2, 2, 427, 428, 7, 50, 2, 2, 428, 82, 3,
	2, 2, 2, 429, 430, 5, 181, 91, 2, 430, 431, 5, 187, 94, 2, 431, 432, 5,
	203, 102, 2, 432, 433, 5, 167, 84, 2, 433, 434, 5, 193, 97, 2, 434, 435,
	5, 163, 82, 2, 435, 436, 5, 159, 80, 2, 436, 437, 5, 195, 98, 2, 437, 438,
	5, 167, 84, 2, 438, 84, 3, 2, 2, 2, 439, 440, 5, 183, 92, 2, 440, 441,
	5, 159, 80, 2, 441, 442, 5, 205, 103, 2, 442, 86, 3, 2, 2, 2, 443, 444,
	5, 183, 92, 2, 444, 445, 5, 167, 84, 2, 445, 446, 5, 165, 83, 2, 446, 447,
	5, 175, 88, 2, 447, 448, 5, 159, 80, 2, 448, 449, 5, 185, 93, 2, 449, 88,
	3, 2, 2, 2, 450, 451, 5, 183, 92, 2, 451, 452, 5, 175, 88, 2, 452, 453,
	5, 185, 93, 2, 453, 90, 3, 2, 2, 2, 454, 455, 5, 183, 92, 2, 455, 456,
	5, 187, 94, 2, 456, 457, 5, 165, 83, 2, 457, 92, 3, 2, 2, 2, 458, 459,
	5, 185, 93, 2, 459, 460, 5, 199, 100, 2, 460, 461, 5, 181, 91, 2, 461,
	462, 5, 181, 91, 2, 462, 94, 3, 2, 2, 2, 463, 464, 5, 189, 95, 2, 464,
	465, 5, 159, 80, 2, 465, 466, 5, 165, 83, 2, 466, 467, 5, 181, 91, 2, 467,
	468, 5, 167, 84, 2, 468, 469, 5, 169, 85, 2, 469, 470, 5, 197, 99, 2, 470,
	96, 3, 2, 2, 2, 471, 472, 5, 189, 95, 2, 472, 473, 5, 159, 80, 2, 473,
	474, 5, 165, 83, 2, 474, 475, 5, 193, 97, 2, 475, 476, 5, 175, 88, 2, 476,
	477, 5, 171, 86, 2, 477, 478, 5, 173, 87, 2, 478, 479, 5, 197, 99, 2, 479,
	98, 3, 2, 2, 2, 480, 481, 5, 189, 95, 2, 481, 482, 5, 175, 88, 2, 482,
	100, 3, 2, 2, 2, 483, 484, 5, 189, 95, 2, 484, 485, 5, 187, 94, 2, 485,
	486, 5, 203, 102, 2, 486, 102, 3, 2, 2, 2, 487, 488, 5, 193, 97, 2, 488,
	489, 5, 159, 80, 2, 489, 490, 5, 185, 93, 2, 490, 491, 5, 165, 83, 2, 491,
	104, 3, 2, 2, 2, 492, 493, 5, 193, 97, 2, 493, 494, 5, 159, 80, 2, 494,
	495, 5, 185, 93, 2, 495, 496, 5, 165, 83, 2, 496, 497, 5, 175, 88, 2, 497,
	498, 5, 185, 93, 2, 498, 499, 5, 197, 99, 2, 499, 106, 3, 2, 2, 2, 500,
	501, 5, 193, 97, 2, 501, 502, 5, 187, 94, 2, 502, 503, 5, 199, 100, 2,
	503, 504, 5, 185, 93, 2, 504, 505, 5, 165, 83, 2, 505, 108, 3, 2, 2, 2,
	506, 507, 5, 195, 98, 2, 507, 508, 5, 175, 88, 2, 508, 509, 5, 185, 93,
	2, 509, 110, 3, 2, 2, 2, 510, 511, 5, 195, 98, 2, 511, 512, 5, 175, 88,
	2, 512, 513, 5, 185, 93, 2, 513, 514, 5, 173, 87, 2, 514, 112, 3, 2, 2,
	2, 515, 516, 5, 195, 98, 2, 516, 517, 5, 191, 96, 2, 517, 518, 5, 193,
	97, 2, 518, 519, 5, 197, 99, 2, 519, 114, 3, 2, 2, 2, 520, 521, 5, 195,
	98, 2, 521, 522, 5, 203, 102, 2, 522, 523, 5, 175, 88, 2, 523, 524, 5,
	197, 99, 2, 524, 525, 5, 163, 82, 2, 525, 526, 5, 173, 87, 2, 526, 116,
	3, 2, 2, 2, 527, 528, 5, 197, 99, 2, 528, 529, 5, 159, 80, 2, 529, 530,
	5, 185, 93, 2, 530, 118, 3, 2, 2, 2, 531, 532, 5, 197, 99, 2, 532, 533,
	5, 159, 80, 2, 533, 534, 5, 185, 93, 2, 534, 535, 5, 173, 87, 2, 535, 120,
	3, 2, 2, 2, 536, 537, 5, 175, 88, 2, 537, 538, 5, 175, 88, 2, 538, 539,
	5, 169, 85, 2, 539, 122, 3, 2, 2, 2, 540, 541, 5, 175, 88, 2, 541, 542,
	5, 185, 93, 2, 542, 124, 3, 2, 2, 2, 543, 544, 5, 185, 93, 2, 544, 545,
	5, 187, 94, 2, 545, 546, 5, 197, 99, 2, 546, 126, 3, 2, 2, 2, 547, 548,
	5, 159, 80, 2, 548, 549, 5, 185, 93, 2, 549, 550, 5, 165, 83, 2, 550, 128,
	3, 2, 2, 2, 551, 552, 5, 187, 94, 2, 552, 553, 5, 193, 97, 2, 553, 130,
	3, 2, 2, 2, 554, 555, 5, 175, 88, 2, 555, 556, 5, 169, 85, 2, 556, 132,
	3, 2, 2, 2, 557, 558, 5, 197, 99, 2, 558, 559, 5, 173, 87, 2, 559, 560,
	5, 167, 84, 2, 560, 561, 5, 185, 93, 2, 561, 134, 3, 2, 2, 2, 562, 563,
	5, 167, 84, 2, 563, 564, 5, 181, 91, 2, 564, 565, 5, 195, 98, 2, 565, 566,
	5, 167, 84, 2, 566, 136, 3, 2, 2, 2, 567, 568, 5, 167, 84, 2, 568, 569,
	5, 181, 91, 2, 569, 570, 5, 195, 98, 2, 570, 571, 5, 167, 84, 2, 571, 572,
	5, 175, 88, 2, 572, 573, 5, 169, 85, 2, 573, 138, 3, 2, 2, 2, 574, 575,
	5, 167, 84, 2, 575, 576, 5, 185, 93, 2, 576, 577, 5, 165, 83, 2, 577, 578,
	5, 175, 88, 2, 578, 579, 5, 169, 85, 2, 579, 140, 3, 2, 2, 2, 580, 583,
	5, 217, 109, 2, 581, 583, 5, 219, 110, 2, 582, 580, 3, 2, 2, 2, 582, 581,
	3, 2, 2, 2, 583, 142, 3, 2, 2, 2, 584, 586, 5, 211, 106, 2, 585, 584, 3,
	2, 2, 2, 586, 587, 3, 2, 2, 2, 587, 585, 3, 2, 2, 2, 587, 588, 3, 2, 2,
	2, 588, 144, 3, 2, 2, 2, 589, 591, 5, 211, 106, 2, 590, 589, 3, 2, 2, 2,
	591, 594, 3, 2, 2, 2, 592, 590, 3, 2, 2, 2, 592, 593, 3, 2, 2, 2, 593,
	595, 3, 2, 2, 2, 594, 592, 3, 2, 2, 2, 595, 597, 7, 48, 2, 2, 596, 598,
	5, 211, 106, 2, 597, 596, 3, 2, 2, 2, 598, 599, 3, 2, 2, 2, 599, 597, 3,
	2, 2, 2, 599, 600, 3, 2, 2, 2, 600, 146, 3, 2, 2, 2, 601, 602, 9, 2, 2,
	2, 602, 603, 5, 213, 107, 2, 603, 604, 9, 2, 2, 2, 604, 148, 3, 2, 2, 2,
	605, 606, 9, 2, 2, 2, 606, 607, 5, 215, 108, 2, 607, 608, 9, 2, 2, 2, 608,
	150, 3, 2, 2, 2, 609, 611, 7, 93, 2, 2, 610, 612, 10, 3, 2, 2, 611, 610,
	3, 2, 2, 2, 612, 613, 3, 2, 2, 2, 613, 611, 3, 2, 2, 2, 613, 614, 3, 2,
	2, 2, 614, 615, 3, 2, 2, 2, 615, 616, 7, 95, 2, 2, 616, 152, 3, 2, 2, 2,
	617, 621, 9, 4, 2, 2, 618, 620, 10, 4, 2, 2, 619, 618, 3, 2, 2, 2, 620,
	623, 3, 2, 2, 2, 621, 619, 3, 2, 2, 2, 621, 622, 3, 2, 2, 2, 622, 624,
	3, 2, 2, 2, 623, 621, 3, 2, 2, 2, 624, 625, 9, 4, 2, 2, 625, 154, 3, 2,
	2, 2, 626, 630, 7, 36, 2, 2, 627, 629, 10, 5, 2, 2, 628, 627, 3, 2, 2,
	2, 629, 632, 3, 2, 2, 2, 630, 628, 3, 2, 2, 2, 630, 631, 3, 2, 2, 2, 631,
	633, 3, 2, 2, 2, 632, 630, 3, 2, 2, 2, 633, 634, 7, 36, 2, 2, 634, 156,
	3, 2, 2, 2, 635, 637, 9, 6, 2, 2, 636, 635, 3, 2, 2, 2, 637, 638, 3, 2,
	2, 2, 638, 636, 3, 2, 2, 2, 638, 639, 3, 2, 2, 2, 639, 640, 3, 2, 2, 2,
	640, 641, 8, 79, 2, 2, 641, 158, 3, 2, 2, 2, 642, 643, 9, 7, 2, 2, 643,
	160, 3, 2, 2, 2, 644, 645, 9, 8, 2, 2, 645, 162, 3, 2, 2, 2, 646, 647,
	9, 9, 2, 2, 647, 164, 3, 2, 2, 2, 648, 649, 9, 10, 2, 2, 649, 166, 3, 2,
	2, 2, 650, 651, 9, 11, 2, 2, 651, 168, 3, 2, 2, 2, 652, 653, 9, 12, 2,
	2, 653, 170, 3, 2, 2, 2, 654, 655, 9, 13, 2, 2, 655, 172, 3, 2, 2, 2, 656,
	657, 9, 14, 2, 2, 657, 174, 3, 2, 2, 2, 658, 659, 9, 15, 2, 2, 659, 176,
	3, 2, 2, 2, 660, 661, 9, 16, 2, 2, 661, 178, 3, 2, 2, 2, 662, 663, 9, 17,
	2, 2, 663, 180, 3, 2, 2, 2, 664, 665, 9, 18, 2, 2, 665, 182, 3, 2, 2, 2,
	666, 667, 9, 19, 2, 2, 667, 184, 3, 2, 2, 2, 668, 669, 9, 20, 2, 2, 669,
	186, 3, 2, 2, 2, 670, 671, 9, 21, 2, 2, 671, 188, 3, 2, 2, 2, 672, 673,
	9, 22, 2, 2, 673, 190, 3, 2, 2, 2, 674, 675, 9, 23, 2, 2, 675, 192, 3,
	2, 2, 2, 676, 677, 9, 24, 2, 2, 677, 194, 3, 2, 2, 2, 678, 679, 9, 25,
	2, 2, 679, 196, 3, 2, 2, 2, 680, 681, 9, 26, 2, 2, 681, 198, 3, 2, 2, 2,
	682, 683, 9, 27, 2, 2, 683, 200, 3, 2, 2, 2, 684, 685, 9, 28, 2, 2, 685,
	202, 3, 2, 2, 2, 686, 687, 9, 29, 2, 2, 687, 204, 3, 2, 2, 2, 688, 689,
	9, 30, 2, 2, 689, 206, 3, 2, 2, 2, 690, 691, 9, 31, 2, 2, 691, 208, 3,
	2, 2, 2, 692, 693, 9, 32, 2, 2, 693, 210, 3, 2, 2, 2, 694, 695, 9, 33,
	2, 2, 695, 212, 3, 2, 2, 2, 696, 697, 5, 211, 106, 2, 697, 698, 5, 211,
	106, 2, 698, 699, 5, 211, 106, 2, 699, 700, 5, 211, 106, 2, 700, 701, 7,
	47, 2, 2, 701, 702, 5, 211, 106, 2, 702, 703, 5, 211, 106, 2, 703, 704,
	7, 47, 2, 2, 704, 705, 5, 211, 106, 2, 705, 706, 5, 211, 106, 2, 706, 214,
	3, 2, 2, 2, 707, 708, 5, 213, 107, 2, 708, 709, 7, 34, 2, 2, 709, 710,
	5, 211, 106, 2, 710, 711, 5, 211, 106, 2, 711, 712, 7, 60, 2, 2, 712, 713,
	5, 211, 106, 2, 713, 714, 5, 211, 106, 2, 714, 715, 7, 60, 2, 2, 715, 716,
	5, 211, 106, 2, 716, 717, 5, 211, 106, 2, 717, 216, 3, 2, 2, 2, 718, 719,
	5, 197, 99, 2, 719, 720, 5, 193, 97, 2, 720, 721, 5, 199, 100, 2, 721,
	722, 5, 167, 84, 2, 722, 218, 3, 2, 2, 2, 723, 724, 5, 169, 85, 2, 724,
	725, 5, 159, 80, 2, 725, 726, 5, 181, 91, 2, 726, 727, 5, 195, 98, 2, 727,
	728, 5, 167, 84, 2, 728, 220, 3, 2, 2, 2, 11, 2, 582, 587, 592, 599, 613,
	621, 630, 638, 3, 8, 2, 2,
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
	"", "'('", "')'", "'*'", "'/'", "'+'", "'-'", "'='", "'>'", "'>='", "'<'",
	"'<='", "'!='", "','", "'&&'", "'||'", "'()'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Abs",
	"Acos", "Asin", "Atan", "Atan2", "Average", "Ceil", "CharFromInt", "CharToInt",
	"Contains", "Cos", "Cosh", "CountWords", "Distance", "EndsWith", "FindString",
	"Exp", "Floor", "GetWord", "HexToNumber", "Left", "Length", "Log", "Log10",
	"Lowercase", "Max", "Median", "Min", "Mod", "Null", "PadLeft", "PadRight",
	"Pi", "Pow", "Rand", "RandInt", "Round", "Sin", "Sinh", "Sqrt", "Switch",
	"Tan", "Tanh", "Iif", "In", "Not", "And", "Or", "If", "Then", "Else", "Elseif",
	"Endif", "Bool", "Integer", "Decimal", "Date", "Datetime", "Field", "SingleQuoteString",
	"DoubleQuoteString", "Whitespace",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "Abs", "Acos",
	"Asin", "Atan", "Atan2", "Average", "Ceil", "CharFromInt", "CharToInt",
	"Contains", "Cos", "Cosh", "CountWords", "Distance", "EndsWith", "FindString",
	"Exp", "Floor", "GetWord", "HexToNumber", "Left", "Length", "Log", "Log10",
	"Lowercase", "Max", "Median", "Min", "Mod", "Null", "PadLeft", "PadRight",
	"Pi", "Pow", "Rand", "RandInt", "Round", "Sin", "Sinh", "Sqrt", "Switch",
	"Tan", "Tanh", "Iif", "In", "Not", "And", "Or", "If", "Then", "Else", "Elseif",
	"Endif", "Bool", "Integer", "Decimal", "Date", "Datetime", "Field", "SingleQuoteString",
	"DoubleQuoteString", "Whitespace", "A", "B", "C", "D", "E", "F", "G", "H",
	"I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W",
	"X", "Y", "Z", "Digit", "DateLiteral", "DateTimeLiteral", "TrueLiteral",
	"FalseLiteral",
}

type AlteryxFormulasLexer struct {
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

func NewAlteryxFormulasLexer(input antlr.CharStream) *AlteryxFormulasLexer {

	l := new(AlteryxFormulasLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "AlteryxFormulas.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// AlteryxFormulasLexer tokens.
const (
	AlteryxFormulasLexerT__0              = 1
	AlteryxFormulasLexerT__1              = 2
	AlteryxFormulasLexerT__2              = 3
	AlteryxFormulasLexerT__3              = 4
	AlteryxFormulasLexerT__4              = 5
	AlteryxFormulasLexerT__5              = 6
	AlteryxFormulasLexerT__6              = 7
	AlteryxFormulasLexerT__7              = 8
	AlteryxFormulasLexerT__8              = 9
	AlteryxFormulasLexerT__9              = 10
	AlteryxFormulasLexerT__10             = 11
	AlteryxFormulasLexerT__11             = 12
	AlteryxFormulasLexerT__12             = 13
	AlteryxFormulasLexerT__13             = 14
	AlteryxFormulasLexerT__14             = 15
	AlteryxFormulasLexerT__15             = 16
	AlteryxFormulasLexerAbs               = 17
	AlteryxFormulasLexerAcos              = 18
	AlteryxFormulasLexerAsin              = 19
	AlteryxFormulasLexerAtan              = 20
	AlteryxFormulasLexerAtan2             = 21
	AlteryxFormulasLexerAverage           = 22
	AlteryxFormulasLexerCeil              = 23
	AlteryxFormulasLexerCharFromInt       = 24
	AlteryxFormulasLexerCharToInt         = 25
	AlteryxFormulasLexerContains          = 26
	AlteryxFormulasLexerCos               = 27
	AlteryxFormulasLexerCosh              = 28
	AlteryxFormulasLexerCountWords        = 29
	AlteryxFormulasLexerDistance          = 30
	AlteryxFormulasLexerEndsWith          = 31
	AlteryxFormulasLexerFindString        = 32
	AlteryxFormulasLexerExp               = 33
	AlteryxFormulasLexerFloor             = 34
	AlteryxFormulasLexerGetWord           = 35
	AlteryxFormulasLexerHexToNumber       = 36
	AlteryxFormulasLexerLeft              = 37
	AlteryxFormulasLexerLength            = 38
	AlteryxFormulasLexerLog               = 39
	AlteryxFormulasLexerLog10             = 40
	AlteryxFormulasLexerLowercase         = 41
	AlteryxFormulasLexerMax               = 42
	AlteryxFormulasLexerMedian            = 43
	AlteryxFormulasLexerMin               = 44
	AlteryxFormulasLexerMod               = 45
	AlteryxFormulasLexerNull              = 46
	AlteryxFormulasLexerPadLeft           = 47
	AlteryxFormulasLexerPadRight          = 48
	AlteryxFormulasLexerPi                = 49
	AlteryxFormulasLexerPow               = 50
	AlteryxFormulasLexerRand              = 51
	AlteryxFormulasLexerRandInt           = 52
	AlteryxFormulasLexerRound             = 53
	AlteryxFormulasLexerSin               = 54
	AlteryxFormulasLexerSinh              = 55
	AlteryxFormulasLexerSqrt              = 56
	AlteryxFormulasLexerSwitch            = 57
	AlteryxFormulasLexerTan               = 58
	AlteryxFormulasLexerTanh              = 59
	AlteryxFormulasLexerIif               = 60
	AlteryxFormulasLexerIn                = 61
	AlteryxFormulasLexerNot               = 62
	AlteryxFormulasLexerAnd               = 63
	AlteryxFormulasLexerOr                = 64
	AlteryxFormulasLexerIf                = 65
	AlteryxFormulasLexerThen              = 66
	AlteryxFormulasLexerElse              = 67
	AlteryxFormulasLexerElseif            = 68
	AlteryxFormulasLexerEndif             = 69
	AlteryxFormulasLexerBool              = 70
	AlteryxFormulasLexerInteger           = 71
	AlteryxFormulasLexerDecimal           = 72
	AlteryxFormulasLexerDate              = 73
	AlteryxFormulasLexerDatetime          = 74
	AlteryxFormulasLexerField             = 75
	AlteryxFormulasLexerSingleQuoteString = 76
	AlteryxFormulasLexerDoubleQuoteString = 77
	AlteryxFormulasLexerWhitespace        = 78
)
