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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 41, 391,
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
	70, 4, 71, 9, 71, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21,
	3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3,
	24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26,
	3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 5, 32, 245, 10, 32, 3,
	33, 6, 33, 248, 10, 33, 13, 33, 14, 33, 249, 3, 34, 7, 34, 253, 10, 34,
	12, 34, 14, 34, 256, 11, 34, 3, 34, 3, 34, 6, 34, 260, 10, 34, 13, 34,
	14, 34, 261, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3,
	37, 3, 37, 6, 37, 274, 10, 37, 13, 37, 14, 37, 275, 3, 37, 3, 37, 3, 38,
	3, 38, 7, 38, 282, 10, 38, 12, 38, 14, 38, 285, 11, 38, 3, 38, 3, 38, 3,
	39, 3, 39, 7, 39, 291, 10, 39, 12, 39, 14, 39, 294, 11, 39, 3, 39, 3, 39,
	3, 40, 6, 40, 299, 10, 40, 13, 40, 14, 40, 300, 3, 40, 3, 40, 3, 41, 3,
	41, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 45, 3, 45, 3, 46, 3, 46,
	3, 47, 3, 47, 3, 48, 3, 48, 3, 49, 3, 49, 3, 50, 3, 50, 3, 51, 3, 51, 3,
	52, 3, 52, 3, 53, 3, 53, 3, 54, 3, 54, 3, 55, 3, 55, 3, 56, 3, 56, 3, 57,
	3, 57, 3, 58, 3, 58, 3, 59, 3, 59, 3, 60, 3, 60, 3, 61, 3, 61, 3, 62, 3,
	62, 3, 63, 3, 63, 3, 64, 3, 64, 3, 65, 3, 65, 3, 66, 3, 66, 3, 67, 3, 67,
	3, 68, 3, 68, 3, 68, 3, 68, 3, 68, 3, 68, 3, 68, 3, 68, 3, 68, 3, 68, 3,
	68, 3, 69, 3, 69, 3, 69, 3, 69, 3, 69, 3, 69, 3, 69, 3, 69, 3, 69, 3, 69,
	3, 69, 3, 70, 3, 70, 3, 70, 3, 70, 3, 70, 3, 71, 3, 71, 3, 71, 3, 71, 3,
	71, 3, 71, 2, 2, 72, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10,
	19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19,
	37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28,
	55, 29, 57, 30, 59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37,
	73, 38, 75, 39, 77, 40, 79, 41, 81, 2, 83, 2, 85, 2, 87, 2, 89, 2, 91,
	2, 93, 2, 95, 2, 97, 2, 99, 2, 101, 2, 103, 2, 105, 2, 107, 2, 109, 2,
	111, 2, 113, 2, 115, 2, 117, 2, 119, 2, 121, 2, 123, 2, 125, 2, 127, 2,
	129, 2, 131, 2, 133, 2, 135, 2, 137, 2, 139, 2, 141, 2, 3, 2, 34, 3, 2,
	98, 98, 3, 2, 95, 95, 3, 2, 41, 41, 3, 2, 36, 36, 5, 2, 11, 12, 15, 15,
	34, 34, 4, 2, 67, 67, 99, 99, 4, 2, 68, 68, 100, 100, 4, 2, 69, 69, 101,
	101, 4, 2, 70, 70, 102, 102, 4, 2, 71, 71, 103, 103, 4, 2, 72, 72, 104,
	104, 4, 2, 73, 73, 105, 105, 4, 2, 74, 74, 106, 106, 4, 2, 75, 75, 107,
	107, 4, 2, 76, 76, 108, 108, 4, 2, 77, 77, 109, 109, 4, 2, 78, 78, 110,
	110, 4, 2, 79, 79, 111, 111, 4, 2, 80, 80, 112, 112, 4, 2, 81, 81, 113,
	113, 4, 2, 82, 82, 114, 114, 4, 2, 83, 83, 115, 115, 4, 2, 84, 84, 116,
	116, 4, 2, 85, 85, 117, 117, 4, 2, 86, 86, 118, 118, 4, 2, 87, 87, 119,
	119, 4, 2, 88, 88, 120, 120, 4, 2, 89, 89, 121, 121, 4, 2, 90, 90, 122,
	122, 4, 2, 91, 91, 123, 123, 4, 2, 92, 92, 124, 124, 3, 2, 50, 59, 2, 367,
	2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2,
	2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2,
	2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2,
	2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3,
	2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41,
	3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2,
	49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2,
	2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2,
	2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2,
	2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3,
	2, 2, 2, 3, 143, 3, 2, 2, 2, 5, 145, 3, 2, 2, 2, 7, 147, 3, 2, 2, 2, 9,
	149, 3, 2, 2, 2, 11, 151, 3, 2, 2, 2, 13, 153, 3, 2, 2, 2, 15, 155, 3,
	2, 2, 2, 17, 157, 3, 2, 2, 2, 19, 159, 3, 2, 2, 2, 21, 162, 3, 2, 2, 2,
	23, 164, 3, 2, 2, 2, 25, 167, 3, 2, 2, 2, 27, 170, 3, 2, 2, 2, 29, 172,
	3, 2, 2, 2, 31, 175, 3, 2, 2, 2, 33, 178, 3, 2, 2, 2, 35, 181, 3, 2, 2,
	2, 37, 185, 3, 2, 2, 2, 39, 189, 3, 2, 2, 2, 41, 193, 3, 2, 2, 2, 43, 198,
	3, 2, 2, 2, 45, 202, 3, 2, 2, 2, 47, 205, 3, 2, 2, 2, 49, 209, 3, 2, 2,
	2, 51, 213, 3, 2, 2, 2, 53, 216, 3, 2, 2, 2, 55, 219, 3, 2, 2, 2, 57, 224,
	3, 2, 2, 2, 59, 229, 3, 2, 2, 2, 61, 236, 3, 2, 2, 2, 63, 244, 3, 2, 2,
	2, 65, 247, 3, 2, 2, 2, 67, 254, 3, 2, 2, 2, 69, 263, 3, 2, 2, 2, 71, 267,
	3, 2, 2, 2, 73, 271, 3, 2, 2, 2, 75, 279, 3, 2, 2, 2, 77, 288, 3, 2, 2,
	2, 79, 298, 3, 2, 2, 2, 81, 304, 3, 2, 2, 2, 83, 306, 3, 2, 2, 2, 85, 308,
	3, 2, 2, 2, 87, 310, 3, 2, 2, 2, 89, 312, 3, 2, 2, 2, 91, 314, 3, 2, 2,
	2, 93, 316, 3, 2, 2, 2, 95, 318, 3, 2, 2, 2, 97, 320, 3, 2, 2, 2, 99, 322,
	3, 2, 2, 2, 101, 324, 3, 2, 2, 2, 103, 326, 3, 2, 2, 2, 105, 328, 3, 2,
	2, 2, 107, 330, 3, 2, 2, 2, 109, 332, 3, 2, 2, 2, 111, 334, 3, 2, 2, 2,
	113, 336, 3, 2, 2, 2, 115, 338, 3, 2, 2, 2, 117, 340, 3, 2, 2, 2, 119,
	342, 3, 2, 2, 2, 121, 344, 3, 2, 2, 2, 123, 346, 3, 2, 2, 2, 125, 348,
	3, 2, 2, 2, 127, 350, 3, 2, 2, 2, 129, 352, 3, 2, 2, 2, 131, 354, 3, 2,
	2, 2, 133, 356, 3, 2, 2, 2, 135, 358, 3, 2, 2, 2, 137, 369, 3, 2, 2, 2,
	139, 380, 3, 2, 2, 2, 141, 385, 3, 2, 2, 2, 143, 144, 7, 42, 2, 2, 144,
	4, 3, 2, 2, 2, 145, 146, 7, 43, 2, 2, 146, 6, 3, 2, 2, 2, 147, 148, 7,
	44, 2, 2, 148, 8, 3, 2, 2, 2, 149, 150, 7, 49, 2, 2, 150, 10, 3, 2, 2,
	2, 151, 152, 7, 45, 2, 2, 152, 12, 3, 2, 2, 2, 153, 154, 7, 47, 2, 2, 154,
	14, 3, 2, 2, 2, 155, 156, 7, 63, 2, 2, 156, 16, 3, 2, 2, 2, 157, 158, 7,
	64, 2, 2, 158, 18, 3, 2, 2, 2, 159, 160, 7, 64, 2, 2, 160, 161, 7, 63,
	2, 2, 161, 20, 3, 2, 2, 2, 162, 163, 7, 62, 2, 2, 163, 22, 3, 2, 2, 2,
	164, 165, 7, 62, 2, 2, 165, 166, 7, 63, 2, 2, 166, 24, 3, 2, 2, 2, 167,
	168, 7, 35, 2, 2, 168, 169, 7, 63, 2, 2, 169, 26, 3, 2, 2, 2, 170, 171,
	7, 46, 2, 2, 171, 28, 3, 2, 2, 2, 172, 173, 7, 40, 2, 2, 173, 174, 7, 40,
	2, 2, 174, 30, 3, 2, 2, 2, 175, 176, 7, 126, 2, 2, 176, 177, 7, 126, 2,
	2, 177, 32, 3, 2, 2, 2, 178, 179, 7, 42, 2, 2, 179, 180, 7, 43, 2, 2, 180,
	34, 3, 2, 2, 2, 181, 182, 5, 111, 56, 2, 182, 183, 5, 109, 55, 2, 183,
	184, 5, 125, 63, 2, 184, 36, 3, 2, 2, 2, 185, 186, 5, 105, 53, 2, 186,
	187, 5, 97, 49, 2, 187, 188, 5, 107, 54, 2, 188, 38, 3, 2, 2, 2, 189, 190,
	5, 105, 53, 2, 190, 191, 5, 81, 41, 2, 191, 192, 5, 127, 64, 2, 192, 40,
	3, 2, 2, 2, 193, 194, 5, 107, 54, 2, 194, 195, 5, 121, 61, 2, 195, 196,
	5, 103, 52, 2, 196, 197, 5, 103, 52, 2, 197, 42, 3, 2, 2, 2, 198, 199,
	5, 97, 49, 2, 199, 200, 5, 97, 49, 2, 200, 201, 5, 91, 46, 2, 201, 44,
	3, 2, 2, 2, 202, 203, 5, 97, 49, 2, 203, 204, 5, 107, 54, 2, 204, 46, 3,
	2, 2, 2, 205, 206, 5, 107, 54, 2, 206, 207, 5, 109, 55, 2, 207, 208, 5,
	119, 60, 2, 208, 48, 3, 2, 2, 2, 209, 210, 5, 81, 41, 2, 210, 211, 5, 107,
	54, 2, 211, 212, 5, 87, 44, 2, 212, 50, 3, 2, 2, 2, 213, 214, 5, 109, 55,
	2, 214, 215, 5, 115, 58, 2, 215, 52, 3, 2, 2, 2, 216, 217, 5, 97, 49, 2,
	217, 218, 5, 91, 46, 2, 218, 54, 3, 2, 2, 2, 219, 220, 5, 119, 60, 2, 220,
	221, 5, 95, 48, 2, 221, 222, 5, 89, 45, 2, 222, 223, 5, 107, 54, 2, 223,
	56, 3, 2, 2, 2, 224, 225, 5, 89, 45, 2, 225, 226, 5, 103, 52, 2, 226, 227,
	5, 117, 59, 2, 227, 228, 5, 89, 45, 2, 228, 58, 3, 2, 2, 2, 229, 230, 5,
	89, 45, 2, 230, 231, 5, 103, 52, 2, 231, 232, 5, 117, 59, 2, 232, 233,
	5, 89, 45, 2, 233, 234, 5, 97, 49, 2, 234, 235, 5, 91, 46, 2, 235, 60,
	3, 2, 2, 2, 236, 237, 5, 89, 45, 2, 237, 238, 5, 107, 54, 2, 238, 239,
	5, 87, 44, 2, 239, 240, 5, 97, 49, 2, 240, 241, 5, 91, 46, 2, 241, 62,
	3, 2, 2, 2, 242, 245, 5, 139, 70, 2, 243, 245, 5, 141, 71, 2, 244, 242,
	3, 2, 2, 2, 244, 243, 3, 2, 2, 2, 245, 64, 3, 2, 2, 2, 246, 248, 5, 133,
	67, 2, 247, 246, 3, 2, 2, 2, 248, 249, 3, 2, 2, 2, 249, 247, 3, 2, 2, 2,
	249, 250, 3, 2, 2, 2, 250, 66, 3, 2, 2, 2, 251, 253, 5, 133, 67, 2, 252,
	251, 3, 2, 2, 2, 253, 256, 3, 2, 2, 2, 254, 252, 3, 2, 2, 2, 254, 255,
	3, 2, 2, 2, 255, 257, 3, 2, 2, 2, 256, 254, 3, 2, 2, 2, 257, 259, 7, 48,
	2, 2, 258, 260, 5, 133, 67, 2, 259, 258, 3, 2, 2, 2, 260, 261, 3, 2, 2,
	2, 261, 259, 3, 2, 2, 2, 261, 262, 3, 2, 2, 2, 262, 68, 3, 2, 2, 2, 263,
	264, 9, 2, 2, 2, 264, 265, 5, 135, 68, 2, 265, 266, 9, 2, 2, 2, 266, 70,
	3, 2, 2, 2, 267, 268, 9, 2, 2, 2, 268, 269, 5, 137, 69, 2, 269, 270, 9,
	2, 2, 2, 270, 72, 3, 2, 2, 2, 271, 273, 7, 93, 2, 2, 272, 274, 10, 3, 2,
	2, 273, 272, 3, 2, 2, 2, 274, 275, 3, 2, 2, 2, 275, 273, 3, 2, 2, 2, 275,
	276, 3, 2, 2, 2, 276, 277, 3, 2, 2, 2, 277, 278, 7, 95, 2, 2, 278, 74,
	3, 2, 2, 2, 279, 283, 9, 4, 2, 2, 280, 282, 10, 4, 2, 2, 281, 280, 3, 2,
	2, 2, 282, 285, 3, 2, 2, 2, 283, 281, 3, 2, 2, 2, 283, 284, 3, 2, 2, 2,
	284, 286, 3, 2, 2, 2, 285, 283, 3, 2, 2, 2, 286, 287, 9, 4, 2, 2, 287,
	76, 3, 2, 2, 2, 288, 292, 7, 36, 2, 2, 289, 291, 10, 5, 2, 2, 290, 289,
	3, 2, 2, 2, 291, 294, 3, 2, 2, 2, 292, 290, 3, 2, 2, 2, 292, 293, 3, 2,
	2, 2, 293, 295, 3, 2, 2, 2, 294, 292, 3, 2, 2, 2, 295, 296, 7, 36, 2, 2,
	296, 78, 3, 2, 2, 2, 297, 299, 9, 6, 2, 2, 298, 297, 3, 2, 2, 2, 299, 300,
	3, 2, 2, 2, 300, 298, 3, 2, 2, 2, 300, 301, 3, 2, 2, 2, 301, 302, 3, 2,
	2, 2, 302, 303, 8, 40, 2, 2, 303, 80, 3, 2, 2, 2, 304, 305, 9, 7, 2, 2,
	305, 82, 3, 2, 2, 2, 306, 307, 9, 8, 2, 2, 307, 84, 3, 2, 2, 2, 308, 309,
	9, 9, 2, 2, 309, 86, 3, 2, 2, 2, 310, 311, 9, 10, 2, 2, 311, 88, 3, 2,
	2, 2, 312, 313, 9, 11, 2, 2, 313, 90, 3, 2, 2, 2, 314, 315, 9, 12, 2, 2,
	315, 92, 3, 2, 2, 2, 316, 317, 9, 13, 2, 2, 317, 94, 3, 2, 2, 2, 318, 319,
	9, 14, 2, 2, 319, 96, 3, 2, 2, 2, 320, 321, 9, 15, 2, 2, 321, 98, 3, 2,
	2, 2, 322, 323, 9, 16, 2, 2, 323, 100, 3, 2, 2, 2, 324, 325, 9, 17, 2,
	2, 325, 102, 3, 2, 2, 2, 326, 327, 9, 18, 2, 2, 327, 104, 3, 2, 2, 2, 328,
	329, 9, 19, 2, 2, 329, 106, 3, 2, 2, 2, 330, 331, 9, 20, 2, 2, 331, 108,
	3, 2, 2, 2, 332, 333, 9, 21, 2, 2, 333, 110, 3, 2, 2, 2, 334, 335, 9, 22,
	2, 2, 335, 112, 3, 2, 2, 2, 336, 337, 9, 23, 2, 2, 337, 114, 3, 2, 2, 2,
	338, 339, 9, 24, 2, 2, 339, 116, 3, 2, 2, 2, 340, 341, 9, 25, 2, 2, 341,
	118, 3, 2, 2, 2, 342, 343, 9, 26, 2, 2, 343, 120, 3, 2, 2, 2, 344, 345,
	9, 27, 2, 2, 345, 122, 3, 2, 2, 2, 346, 347, 9, 28, 2, 2, 347, 124, 3,
	2, 2, 2, 348, 349, 9, 29, 2, 2, 349, 126, 3, 2, 2, 2, 350, 351, 9, 30,
	2, 2, 351, 128, 3, 2, 2, 2, 352, 353, 9, 31, 2, 2, 353, 130, 3, 2, 2, 2,
	354, 355, 9, 32, 2, 2, 355, 132, 3, 2, 2, 2, 356, 357, 9, 33, 2, 2, 357,
	134, 3, 2, 2, 2, 358, 359, 5, 133, 67, 2, 359, 360, 5, 133, 67, 2, 360,
	361, 5, 133, 67, 2, 361, 362, 5, 133, 67, 2, 362, 363, 7, 47, 2, 2, 363,
	364, 5, 133, 67, 2, 364, 365, 5, 133, 67, 2, 365, 366, 7, 47, 2, 2, 366,
	367, 5, 133, 67, 2, 367, 368, 5, 133, 67, 2, 368, 136, 3, 2, 2, 2, 369,
	370, 5, 135, 68, 2, 370, 371, 7, 34, 2, 2, 371, 372, 5, 133, 67, 2, 372,
	373, 5, 133, 67, 2, 373, 374, 7, 60, 2, 2, 374, 375, 5, 133, 67, 2, 375,
	376, 5, 133, 67, 2, 376, 377, 7, 60, 2, 2, 377, 378, 5, 133, 67, 2, 378,
	379, 5, 133, 67, 2, 379, 138, 3, 2, 2, 2, 380, 381, 5, 119, 60, 2, 381,
	382, 5, 115, 58, 2, 382, 383, 5, 121, 61, 2, 383, 384, 5, 89, 45, 2, 384,
	140, 3, 2, 2, 2, 385, 386, 5, 91, 46, 2, 386, 387, 5, 81, 41, 2, 387, 388,
	5, 103, 52, 2, 388, 389, 5, 117, 59, 2, 389, 390, 5, 89, 45, 2, 390, 142,
	3, 2, 2, 2, 11, 2, 244, 249, 254, 261, 275, 283, 292, 300, 3, 8, 2, 2,
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
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Pow",
	"Min", "Max", "Null", "Iif", "In", "Not", "And", "Or", "If", "Then", "Else",
	"Elseif", "Endif", "Bool", "Integer", "Decimal", "Date", "Datetime", "Field",
	"SingleQuoteString", "DoubleQuoteString", "Whitespace",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "Pow", "Min",
	"Max", "Null", "Iif", "In", "Not", "And", "Or", "If", "Then", "Else", "Elseif",
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
	AlteryxFormulasLexerPow               = 17
	AlteryxFormulasLexerMin               = 18
	AlteryxFormulasLexerMax               = 19
	AlteryxFormulasLexerNull              = 20
	AlteryxFormulasLexerIif               = 21
	AlteryxFormulasLexerIn                = 22
	AlteryxFormulasLexerNot               = 23
	AlteryxFormulasLexerAnd               = 24
	AlteryxFormulasLexerOr                = 25
	AlteryxFormulasLexerIf                = 26
	AlteryxFormulasLexerThen              = 27
	AlteryxFormulasLexerElse              = 28
	AlteryxFormulasLexerElseif            = 29
	AlteryxFormulasLexerEndif             = 30
	AlteryxFormulasLexerBool              = 31
	AlteryxFormulasLexerInteger           = 32
	AlteryxFormulasLexerDecimal           = 33
	AlteryxFormulasLexerDate              = 34
	AlteryxFormulasLexerDatetime          = 35
	AlteryxFormulasLexerField             = 36
	AlteryxFormulasLexerSingleQuoteString = 37
	AlteryxFormulasLexerDoubleQuoteString = 38
	AlteryxFormulasLexerWhitespace        = 39
)
