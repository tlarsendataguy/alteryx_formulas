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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 62, 546,
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
	91, 4, 92, 9, 92, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15,
	3, 15, 3, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3,
	22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3,
	26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 27,
	3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32,
	3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36,
	3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 39, 3,
	39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40,
	3, 40, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 41, 3, 42, 3, 42, 3, 42, 3,
	42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 45, 3, 45, 3, 45,
	3, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3,
	48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50,
	3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 51, 3, 52, 3, 52, 3, 52, 3,
	52, 3, 52, 3, 52, 3, 53, 3, 53, 5, 53, 400, 10, 53, 3, 54, 6, 54, 403,
	10, 54, 13, 54, 14, 54, 404, 3, 55, 7, 55, 408, 10, 55, 12, 55, 14, 55,
	411, 11, 55, 3, 55, 3, 55, 6, 55, 415, 10, 55, 13, 55, 14, 55, 416, 3,
	56, 3, 56, 3, 56, 3, 56, 3, 57, 3, 57, 3, 57, 3, 57, 3, 58, 3, 58, 6, 58,
	429, 10, 58, 13, 58, 14, 58, 430, 3, 58, 3, 58, 3, 59, 3, 59, 7, 59, 437,
	10, 59, 12, 59, 14, 59, 440, 11, 59, 3, 59, 3, 59, 3, 60, 3, 60, 7, 60,
	446, 10, 60, 12, 60, 14, 60, 449, 11, 60, 3, 60, 3, 60, 3, 61, 6, 61, 454,
	10, 61, 13, 61, 14, 61, 455, 3, 61, 3, 61, 3, 62, 3, 62, 3, 63, 3, 63,
	3, 64, 3, 64, 3, 65, 3, 65, 3, 66, 3, 66, 3, 67, 3, 67, 3, 68, 3, 68, 3,
	69, 3, 69, 3, 70, 3, 70, 3, 71, 3, 71, 3, 72, 3, 72, 3, 73, 3, 73, 3, 74,
	3, 74, 3, 75, 3, 75, 3, 76, 3, 76, 3, 77, 3, 77, 3, 78, 3, 78, 3, 79, 3,
	79, 3, 80, 3, 80, 3, 81, 3, 81, 3, 82, 3, 82, 3, 83, 3, 83, 3, 84, 3, 84,
	3, 85, 3, 85, 3, 86, 3, 86, 3, 87, 3, 87, 3, 88, 3, 88, 3, 89, 3, 89, 3,
	89, 3, 89, 3, 89, 3, 89, 3, 89, 3, 89, 3, 89, 3, 89, 3, 89, 3, 90, 3, 90,
	3, 90, 3, 90, 3, 90, 3, 90, 3, 90, 3, 90, 3, 90, 3, 90, 3, 90, 3, 91, 3,
	91, 3, 91, 3, 91, 3, 91, 3, 92, 3, 92, 3, 92, 3, 92, 3, 92, 3, 92, 2, 2,
	93, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12,
	23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21,
	41, 22, 43, 23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30,
	59, 31, 61, 32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39,
	77, 40, 79, 41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 93, 48,
	95, 49, 97, 50, 99, 51, 101, 52, 103, 53, 105, 54, 107, 55, 109, 56, 111,
	57, 113, 58, 115, 59, 117, 60, 119, 61, 121, 62, 123, 2, 125, 2, 127, 2,
	129, 2, 131, 2, 133, 2, 135, 2, 137, 2, 139, 2, 141, 2, 143, 2, 145, 2,
	147, 2, 149, 2, 151, 2, 153, 2, 155, 2, 157, 2, 159, 2, 161, 2, 163, 2,
	165, 2, 167, 2, 169, 2, 171, 2, 173, 2, 175, 2, 177, 2, 179, 2, 181, 2,
	183, 2, 3, 2, 34, 3, 2, 98, 98, 3, 2, 95, 95, 3, 2, 41, 41, 3, 2, 36, 36,
	5, 2, 11, 12, 15, 15, 34, 34, 4, 2, 67, 67, 99, 99, 4, 2, 68, 68, 100,
	100, 4, 2, 69, 69, 101, 101, 4, 2, 70, 70, 102, 102, 4, 2, 71, 71, 103,
	103, 4, 2, 72, 72, 104, 104, 4, 2, 73, 73, 105, 105, 4, 2, 74, 74, 106,
	106, 4, 2, 75, 75, 107, 107, 4, 2, 76, 76, 108, 108, 4, 2, 77, 77, 109,
	109, 4, 2, 78, 78, 110, 110, 4, 2, 79, 79, 111, 111, 4, 2, 80, 80, 112,
	112, 4, 2, 81, 81, 113, 113, 4, 2, 82, 82, 114, 114, 4, 2, 83, 83, 115,
	115, 4, 2, 84, 84, 116, 116, 4, 2, 85, 85, 117, 117, 4, 2, 86, 86, 118,
	118, 4, 2, 87, 87, 119, 119, 4, 2, 88, 88, 120, 120, 4, 2, 89, 89, 121,
	121, 4, 2, 90, 90, 122, 122, 4, 2, 91, 91, 123, 123, 4, 2, 92, 92, 124,
	124, 3, 2, 50, 59, 2, 522, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3,
	2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15,
	3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2,
	23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2,
	2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2,
	2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2,
	2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3,
	2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61,
	3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2,
	69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2, 73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2,
	2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2, 2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2,
	2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2, 2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2,
	2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3,
	2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 103, 3, 2, 2, 2, 2, 105, 3, 2, 2, 2, 2,
	107, 3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2, 113, 3, 2,
	2, 2, 2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 119, 3, 2, 2, 2, 2, 121,
	3, 2, 2, 2, 3, 185, 3, 2, 2, 2, 5, 187, 3, 2, 2, 2, 7, 189, 3, 2, 2, 2,
	9, 191, 3, 2, 2, 2, 11, 193, 3, 2, 2, 2, 13, 195, 3, 2, 2, 2, 15, 197,
	3, 2, 2, 2, 17, 199, 3, 2, 2, 2, 19, 201, 3, 2, 2, 2, 21, 204, 3, 2, 2,
	2, 23, 206, 3, 2, 2, 2, 25, 209, 3, 2, 2, 2, 27, 212, 3, 2, 2, 2, 29, 214,
	3, 2, 2, 2, 31, 217, 3, 2, 2, 2, 33, 220, 3, 2, 2, 2, 35, 223, 3, 2, 2,
	2, 37, 227, 3, 2, 2, 2, 39, 232, 3, 2, 2, 2, 41, 237, 3, 2, 2, 2, 43, 242,
	3, 2, 2, 2, 45, 248, 3, 2, 2, 2, 47, 256, 3, 2, 2, 2, 49, 261, 3, 2, 2,
	2, 51, 265, 3, 2, 2, 2, 53, 270, 3, 2, 2, 2, 55, 279, 3, 2, 2, 2, 57, 283,
	3, 2, 2, 2, 59, 289, 3, 2, 2, 2, 61, 293, 3, 2, 2, 2, 63, 299, 3, 2, 2,
	2, 65, 303, 3, 2, 2, 2, 67, 310, 3, 2, 2, 2, 69, 314, 3, 2, 2, 2, 71, 318,
	3, 2, 2, 2, 73, 323, 3, 2, 2, 2, 75, 326, 3, 2, 2, 2, 77, 330, 3, 2, 2,
	2, 79, 335, 3, 2, 2, 2, 81, 343, 3, 2, 2, 2, 83, 349, 3, 2, 2, 2, 85, 353,
	3, 2, 2, 2, 87, 357, 3, 2, 2, 2, 89, 360, 3, 2, 2, 2, 91, 364, 3, 2, 2,
	2, 93, 368, 3, 2, 2, 2, 95, 371, 3, 2, 2, 2, 97, 374, 3, 2, 2, 2, 99, 379,
	3, 2, 2, 2, 101, 384, 3, 2, 2, 2, 103, 391, 3, 2, 2, 2, 105, 399, 3, 2,
	2, 2, 107, 402, 3, 2, 2, 2, 109, 409, 3, 2, 2, 2, 111, 418, 3, 2, 2, 2,
	113, 422, 3, 2, 2, 2, 115, 426, 3, 2, 2, 2, 117, 434, 3, 2, 2, 2, 119,
	443, 3, 2, 2, 2, 121, 453, 3, 2, 2, 2, 123, 459, 3, 2, 2, 2, 125, 461,
	3, 2, 2, 2, 127, 463, 3, 2, 2, 2, 129, 465, 3, 2, 2, 2, 131, 467, 3, 2,
	2, 2, 133, 469, 3, 2, 2, 2, 135, 471, 3, 2, 2, 2, 137, 473, 3, 2, 2, 2,
	139, 475, 3, 2, 2, 2, 141, 477, 3, 2, 2, 2, 143, 479, 3, 2, 2, 2, 145,
	481, 3, 2, 2, 2, 147, 483, 3, 2, 2, 2, 149, 485, 3, 2, 2, 2, 151, 487,
	3, 2, 2, 2, 153, 489, 3, 2, 2, 2, 155, 491, 3, 2, 2, 2, 157, 493, 3, 2,
	2, 2, 159, 495, 3, 2, 2, 2, 161, 497, 3, 2, 2, 2, 163, 499, 3, 2, 2, 2,
	165, 501, 3, 2, 2, 2, 167, 503, 3, 2, 2, 2, 169, 505, 3, 2, 2, 2, 171,
	507, 3, 2, 2, 2, 173, 509, 3, 2, 2, 2, 175, 511, 3, 2, 2, 2, 177, 513,
	3, 2, 2, 2, 179, 524, 3, 2, 2, 2, 181, 535, 3, 2, 2, 2, 183, 540, 3, 2,
	2, 2, 185, 186, 7, 42, 2, 2, 186, 4, 3, 2, 2, 2, 187, 188, 7, 43, 2, 2,
	188, 6, 3, 2, 2, 2, 189, 190, 7, 44, 2, 2, 190, 8, 3, 2, 2, 2, 191, 192,
	7, 49, 2, 2, 192, 10, 3, 2, 2, 2, 193, 194, 7, 45, 2, 2, 194, 12, 3, 2,
	2, 2, 195, 196, 7, 47, 2, 2, 196, 14, 3, 2, 2, 2, 197, 198, 7, 63, 2, 2,
	198, 16, 3, 2, 2, 2, 199, 200, 7, 64, 2, 2, 200, 18, 3, 2, 2, 2, 201, 202,
	7, 64, 2, 2, 202, 203, 7, 63, 2, 2, 203, 20, 3, 2, 2, 2, 204, 205, 7, 62,
	2, 2, 205, 22, 3, 2, 2, 2, 206, 207, 7, 62, 2, 2, 207, 208, 7, 63, 2, 2,
	208, 24, 3, 2, 2, 2, 209, 210, 7, 35, 2, 2, 210, 211, 7, 63, 2, 2, 211,
	26, 3, 2, 2, 2, 212, 213, 7, 46, 2, 2, 213, 28, 3, 2, 2, 2, 214, 215, 7,
	40, 2, 2, 215, 216, 7, 40, 2, 2, 216, 30, 3, 2, 2, 2, 217, 218, 7, 126,
	2, 2, 218, 219, 7, 126, 2, 2, 219, 32, 3, 2, 2, 2, 220, 221, 7, 42, 2,
	2, 221, 222, 7, 43, 2, 2, 222, 34, 3, 2, 2, 2, 223, 224, 5, 123, 62, 2,
	224, 225, 5, 125, 63, 2, 225, 226, 5, 159, 80, 2, 226, 36, 3, 2, 2, 2,
	227, 228, 5, 123, 62, 2, 228, 229, 5, 127, 64, 2, 229, 230, 5, 151, 76,
	2, 230, 231, 5, 159, 80, 2, 231, 38, 3, 2, 2, 2, 232, 233, 5, 123, 62,
	2, 233, 234, 5, 159, 80, 2, 234, 235, 5, 139, 70, 2, 235, 236, 5, 149,
	75, 2, 236, 40, 3, 2, 2, 2, 237, 238, 5, 123, 62, 2, 238, 239, 5, 161,
	81, 2, 239, 240, 5, 123, 62, 2, 240, 241, 5, 149, 75, 2, 241, 42, 3, 2,
	2, 2, 242, 243, 5, 123, 62, 2, 243, 244, 5, 161, 81, 2, 244, 245, 5, 123,
	62, 2, 245, 246, 5, 149, 75, 2, 246, 247, 7, 52, 2, 2, 247, 44, 3, 2, 2,
	2, 248, 249, 5, 123, 62, 2, 249, 250, 5, 165, 83, 2, 250, 251, 5, 131,
	66, 2, 251, 252, 5, 157, 79, 2, 252, 253, 5, 123, 62, 2, 253, 254, 5, 135,
	68, 2, 254, 255, 5, 131, 66, 2, 255, 46, 3, 2, 2, 2, 256, 257, 5, 127,
	64, 2, 257, 258, 5, 131, 66, 2, 258, 259, 5, 139, 70, 2, 259, 260, 5, 145,
	73, 2, 260, 48, 3, 2, 2, 2, 261, 262, 5, 127, 64, 2, 262, 263, 5, 151,
	76, 2, 263, 264, 5, 159, 80, 2, 264, 50, 3, 2, 2, 2, 265, 266, 5, 127,
	64, 2, 266, 267, 5, 151, 76, 2, 267, 268, 5, 159, 80, 2, 268, 269, 5, 137,
	69, 2, 269, 52, 3, 2, 2, 2, 270, 271, 5, 129, 65, 2, 271, 272, 5, 139,
	70, 2, 272, 273, 5, 159, 80, 2, 273, 274, 5, 161, 81, 2, 274, 275, 5, 123,
	62, 2, 275, 276, 5, 149, 75, 2, 276, 277, 5, 127, 64, 2, 277, 278, 5, 131,
	66, 2, 278, 54, 3, 2, 2, 2, 279, 280, 5, 131, 66, 2, 280, 281, 5, 169,
	85, 2, 281, 282, 5, 153, 77, 2, 282, 56, 3, 2, 2, 2, 283, 284, 5, 133,
	67, 2, 284, 285, 5, 145, 73, 2, 285, 286, 5, 151, 76, 2, 286, 287, 5, 151,
	76, 2, 287, 288, 5, 157, 79, 2, 288, 58, 3, 2, 2, 2, 289, 290, 5, 145,
	73, 2, 290, 291, 5, 151, 76, 2, 291, 292, 5, 135, 68, 2, 292, 60, 3, 2,
	2, 2, 293, 294, 5, 145, 73, 2, 294, 295, 5, 151, 76, 2, 295, 296, 5, 135,
	68, 2, 296, 297, 7, 51, 2, 2, 297, 298, 7, 50, 2, 2, 298, 62, 3, 2, 2,
	2, 299, 300, 5, 147, 74, 2, 300, 301, 5, 123, 62, 2, 301, 302, 5, 169,
	85, 2, 302, 64, 3, 2, 2, 2, 303, 304, 5, 147, 74, 2, 304, 305, 5, 131,
	66, 2, 305, 306, 5, 129, 65, 2, 306, 307, 5, 139, 70, 2, 307, 308, 5, 123,
	62, 2, 308, 309, 5, 149, 75, 2, 309, 66, 3, 2, 2, 2, 310, 311, 5, 147,
	74, 2, 311, 312, 5, 139, 70, 2, 312, 313, 5, 149, 75, 2, 313, 68, 3, 2,
	2, 2, 314, 315, 5, 147, 74, 2, 315, 316, 5, 151, 76, 2, 316, 317, 5, 129,
	65, 2, 317, 70, 3, 2, 2, 2, 318, 319, 5, 149, 75, 2, 319, 320, 5, 163,
	82, 2, 320, 321, 5, 145, 73, 2, 321, 322, 5, 145, 73, 2, 322, 72, 3, 2,
	2, 2, 323, 324, 5, 153, 77, 2, 324, 325, 5, 139, 70, 2, 325, 74, 3, 2,
	2, 2, 326, 327, 5, 153, 77, 2, 327, 328, 5, 151, 76, 2, 328, 329, 5, 167,
	84, 2, 329, 76, 3, 2, 2, 2, 330, 331, 5, 157, 79, 2, 331, 332, 5, 123,
	62, 2, 332, 333, 5, 149, 75, 2, 333, 334, 5, 129, 65, 2, 334, 78, 3, 2,
	2, 2, 335, 336, 5, 157, 79, 2, 336, 337, 5, 123, 62, 2, 337, 338, 5, 149,
	75, 2, 338, 339, 5, 129, 65, 2, 339, 340, 5, 139, 70, 2, 340, 341, 5, 149,
	75, 2, 341, 342, 5, 161, 81, 2, 342, 80, 3, 2, 2, 2, 343, 344, 5, 157,
	79, 2, 344, 345, 5, 151, 76, 2, 345, 346, 5, 163, 82, 2, 346, 347, 5, 149,
	75, 2, 347, 348, 5, 129, 65, 2, 348, 82, 3, 2, 2, 2, 349, 350, 5, 159,
	80, 2, 350, 351, 5, 139, 70, 2, 351, 352, 5, 149, 75, 2, 352, 84, 3, 2,
	2, 2, 353, 354, 5, 139, 70, 2, 354, 355, 5, 139, 70, 2, 355, 356, 5, 133,
	67, 2, 356, 86, 3, 2, 2, 2, 357, 358, 5, 139, 70, 2, 358, 359, 5, 149,
	75, 2, 359, 88, 3, 2, 2, 2, 360, 361, 5, 149, 75, 2, 361, 362, 5, 151,
	76, 2, 362, 363, 5, 161, 81, 2, 363, 90, 3, 2, 2, 2, 364, 365, 5, 123,
	62, 2, 365, 366, 5, 149, 75, 2, 366, 367, 5, 129, 65, 2, 367, 92, 3, 2,
	2, 2, 368, 369, 5, 151, 76, 2, 369, 370, 5, 157, 79, 2, 370, 94, 3, 2,
	2, 2, 371, 372, 5, 139, 70, 2, 372, 373, 5, 133, 67, 2, 373, 96, 3, 2,
	2, 2, 374, 375, 5, 161, 81, 2, 375, 376, 5, 137, 69, 2, 376, 377, 5, 131,
	66, 2, 377, 378, 5, 149, 75, 2, 378, 98, 3, 2, 2, 2, 379, 380, 5, 131,
	66, 2, 380, 381, 5, 145, 73, 2, 381, 382, 5, 159, 80, 2, 382, 383, 5, 131,
	66, 2, 383, 100, 3, 2, 2, 2, 384, 385, 5, 131, 66, 2, 385, 386, 5, 145,
	73, 2, 386, 387, 5, 159, 80, 2, 387, 388, 5, 131, 66, 2, 388, 389, 5, 139,
	70, 2, 389, 390, 5, 133, 67, 2, 390, 102, 3, 2, 2, 2, 391, 392, 5, 131,
	66, 2, 392, 393, 5, 149, 75, 2, 393, 394, 5, 129, 65, 2, 394, 395, 5, 139,
	70, 2, 395, 396, 5, 133, 67, 2, 396, 104, 3, 2, 2, 2, 397, 400, 5, 181,
	91, 2, 398, 400, 5, 183, 92, 2, 399, 397, 3, 2, 2, 2, 399, 398, 3, 2, 2,
	2, 400, 106, 3, 2, 2, 2, 401, 403, 5, 175, 88, 2, 402, 401, 3, 2, 2, 2,
	403, 404, 3, 2, 2, 2, 404, 402, 3, 2, 2, 2, 404, 405, 3, 2, 2, 2, 405,
	108, 3, 2, 2, 2, 406, 408, 5, 175, 88, 2, 407, 406, 3, 2, 2, 2, 408, 411,
	3, 2, 2, 2, 409, 407, 3, 2, 2, 2, 409, 410, 3, 2, 2, 2, 410, 412, 3, 2,
	2, 2, 411, 409, 3, 2, 2, 2, 412, 414, 7, 48, 2, 2, 413, 415, 5, 175, 88,
	2, 414, 413, 3, 2, 2, 2, 415, 416, 3, 2, 2, 2, 416, 414, 3, 2, 2, 2, 416,
	417, 3, 2, 2, 2, 417, 110, 3, 2, 2, 2, 418, 419, 9, 2, 2, 2, 419, 420,
	5, 177, 89, 2, 420, 421, 9, 2, 2, 2, 421, 112, 3, 2, 2, 2, 422, 423, 9,
	2, 2, 2, 423, 424, 5, 179, 90, 2, 424, 425, 9, 2, 2, 2, 425, 114, 3, 2,
	2, 2, 426, 428, 7, 93, 2, 2, 427, 429, 10, 3, 2, 2, 428, 427, 3, 2, 2,
	2, 429, 430, 3, 2, 2, 2, 430, 428, 3, 2, 2, 2, 430, 431, 3, 2, 2, 2, 431,
	432, 3, 2, 2, 2, 432, 433, 7, 95, 2, 2, 433, 116, 3, 2, 2, 2, 434, 438,
	9, 4, 2, 2, 435, 437, 10, 4, 2, 2, 436, 435, 3, 2, 2, 2, 437, 440, 3, 2,
	2, 2, 438, 436, 3, 2, 2, 2, 438, 439, 3, 2, 2, 2, 439, 441, 3, 2, 2, 2,
	440, 438, 3, 2, 2, 2, 441, 442, 9, 4, 2, 2, 442, 118, 3, 2, 2, 2, 443,
	447, 7, 36, 2, 2, 444, 446, 10, 5, 2, 2, 445, 444, 3, 2, 2, 2, 446, 449,
	3, 2, 2, 2, 447, 445, 3, 2, 2, 2, 447, 448, 3, 2, 2, 2, 448, 450, 3, 2,
	2, 2, 449, 447, 3, 2, 2, 2, 450, 451, 7, 36, 2, 2, 451, 120, 3, 2, 2, 2,
	452, 454, 9, 6, 2, 2, 453, 452, 3, 2, 2, 2, 454, 455, 3, 2, 2, 2, 455,
	453, 3, 2, 2, 2, 455, 456, 3, 2, 2, 2, 456, 457, 3, 2, 2, 2, 457, 458,
	8, 61, 2, 2, 458, 122, 3, 2, 2, 2, 459, 460, 9, 7, 2, 2, 460, 124, 3, 2,
	2, 2, 461, 462, 9, 8, 2, 2, 462, 126, 3, 2, 2, 2, 463, 464, 9, 9, 2, 2,
	464, 128, 3, 2, 2, 2, 465, 466, 9, 10, 2, 2, 466, 130, 3, 2, 2, 2, 467,
	468, 9, 11, 2, 2, 468, 132, 3, 2, 2, 2, 469, 470, 9, 12, 2, 2, 470, 134,
	3, 2, 2, 2, 471, 472, 9, 13, 2, 2, 472, 136, 3, 2, 2, 2, 473, 474, 9, 14,
	2, 2, 474, 138, 3, 2, 2, 2, 475, 476, 9, 15, 2, 2, 476, 140, 3, 2, 2, 2,
	477, 478, 9, 16, 2, 2, 478, 142, 3, 2, 2, 2, 479, 480, 9, 17, 2, 2, 480,
	144, 3, 2, 2, 2, 481, 482, 9, 18, 2, 2, 482, 146, 3, 2, 2, 2, 483, 484,
	9, 19, 2, 2, 484, 148, 3, 2, 2, 2, 485, 486, 9, 20, 2, 2, 486, 150, 3,
	2, 2, 2, 487, 488, 9, 21, 2, 2, 488, 152, 3, 2, 2, 2, 489, 490, 9, 22,
	2, 2, 490, 154, 3, 2, 2, 2, 491, 492, 9, 23, 2, 2, 492, 156, 3, 2, 2, 2,
	493, 494, 9, 24, 2, 2, 494, 158, 3, 2, 2, 2, 495, 496, 9, 25, 2, 2, 496,
	160, 3, 2, 2, 2, 497, 498, 9, 26, 2, 2, 498, 162, 3, 2, 2, 2, 499, 500,
	9, 27, 2, 2, 500, 164, 3, 2, 2, 2, 501, 502, 9, 28, 2, 2, 502, 166, 3,
	2, 2, 2, 503, 504, 9, 29, 2, 2, 504, 168, 3, 2, 2, 2, 505, 506, 9, 30,
	2, 2, 506, 170, 3, 2, 2, 2, 507, 508, 9, 31, 2, 2, 508, 172, 3, 2, 2, 2,
	509, 510, 9, 32, 2, 2, 510, 174, 3, 2, 2, 2, 511, 512, 9, 33, 2, 2, 512,
	176, 3, 2, 2, 2, 513, 514, 5, 175, 88, 2, 514, 515, 5, 175, 88, 2, 515,
	516, 5, 175, 88, 2, 516, 517, 5, 175, 88, 2, 517, 518, 7, 47, 2, 2, 518,
	519, 5, 175, 88, 2, 519, 520, 5, 175, 88, 2, 520, 521, 7, 47, 2, 2, 521,
	522, 5, 175, 88, 2, 522, 523, 5, 175, 88, 2, 523, 178, 3, 2, 2, 2, 524,
	525, 5, 177, 89, 2, 525, 526, 7, 34, 2, 2, 526, 527, 5, 175, 88, 2, 527,
	528, 5, 175, 88, 2, 528, 529, 7, 60, 2, 2, 529, 530, 5, 175, 88, 2, 530,
	531, 5, 175, 88, 2, 531, 532, 7, 60, 2, 2, 532, 533, 5, 175, 88, 2, 533,
	534, 5, 175, 88, 2, 534, 180, 3, 2, 2, 2, 535, 536, 5, 161, 81, 2, 536,
	537, 5, 157, 79, 2, 537, 538, 5, 163, 82, 2, 538, 539, 5, 131, 66, 2, 539,
	182, 3, 2, 2, 2, 540, 541, 5, 133, 67, 2, 541, 542, 5, 123, 62, 2, 542,
	543, 5, 145, 73, 2, 543, 544, 5, 159, 80, 2, 544, 545, 5, 131, 66, 2, 545,
	184, 3, 2, 2, 2, 11, 2, 399, 404, 409, 416, 430, 438, 447, 455, 3, 8, 2,
	2,
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
	"Acos", "Asin", "Atan", "Atan2", "Average", "Ceil", "Cos", "Cosh", "Distance",
	"Exp", "Floor", "Log", "Log10", "Max", "Median", "Min", "Mod", "Null",
	"Pi", "Pow", "Rand", "RandInt", "Round", "Sin", "Iif", "In", "Not", "And",
	"Or", "If", "Then", "Else", "Elseif", "Endif", "Bool", "Integer", "Decimal",
	"Date", "Datetime", "Field", "SingleQuoteString", "DoubleQuoteString",
	"Whitespace",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "Abs", "Acos",
	"Asin", "Atan", "Atan2", "Average", "Ceil", "Cos", "Cosh", "Distance",
	"Exp", "Floor", "Log", "Log10", "Max", "Median", "Min", "Mod", "Null",
	"Pi", "Pow", "Rand", "RandInt", "Round", "Sin", "Iif", "In", "Not", "And",
	"Or", "If", "Then", "Else", "Elseif", "Endif", "Bool", "Integer", "Decimal",
	"Date", "Datetime", "Field", "SingleQuoteString", "DoubleQuoteString",
	"Whitespace", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L",
	"M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "Digit",
	"DateLiteral", "DateTimeLiteral", "TrueLiteral", "FalseLiteral",
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
	AlteryxFormulasLexerCos               = 24
	AlteryxFormulasLexerCosh              = 25
	AlteryxFormulasLexerDistance          = 26
	AlteryxFormulasLexerExp               = 27
	AlteryxFormulasLexerFloor             = 28
	AlteryxFormulasLexerLog               = 29
	AlteryxFormulasLexerLog10             = 30
	AlteryxFormulasLexerMax               = 31
	AlteryxFormulasLexerMedian            = 32
	AlteryxFormulasLexerMin               = 33
	AlteryxFormulasLexerMod               = 34
	AlteryxFormulasLexerNull              = 35
	AlteryxFormulasLexerPi                = 36
	AlteryxFormulasLexerPow               = 37
	AlteryxFormulasLexerRand              = 38
	AlteryxFormulasLexerRandInt           = 39
	AlteryxFormulasLexerRound             = 40
	AlteryxFormulasLexerSin               = 41
	AlteryxFormulasLexerIif               = 42
	AlteryxFormulasLexerIn                = 43
	AlteryxFormulasLexerNot               = 44
	AlteryxFormulasLexerAnd               = 45
	AlteryxFormulasLexerOr                = 46
	AlteryxFormulasLexerIf                = 47
	AlteryxFormulasLexerThen              = 48
	AlteryxFormulasLexerElse              = 49
	AlteryxFormulasLexerElseif            = 50
	AlteryxFormulasLexerEndif             = 51
	AlteryxFormulasLexerBool              = 52
	AlteryxFormulasLexerInteger           = 53
	AlteryxFormulasLexerDecimal           = 54
	AlteryxFormulasLexerDate              = 55
	AlteryxFormulasLexerDatetime          = 56
	AlteryxFormulasLexerField             = 57
	AlteryxFormulasLexerSingleQuoteString = 58
	AlteryxFormulasLexerDoubleQuoteString = 59
	AlteryxFormulasLexerWhitespace        = 60
)
