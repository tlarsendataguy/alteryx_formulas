// Code generated from /Users/tlarsen/Documents/Repositories/alteryx_formulas/grammar/AlteryxFormulas.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // AlteryxFormulas

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 85, 460,
	4, 2, 9, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 19, 10, 2, 12, 2, 14, 2, 22, 11, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 69, 10, 2, 12, 2,
	14, 2, 72, 11, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 5, 2, 98, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 135, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 6, 2,
	200, 10, 2, 13, 2, 14, 2, 201, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	7, 2, 211, 10, 2, 12, 2, 14, 2, 214, 11, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 6, 2, 223, 10, 2, 13, 2, 14, 2, 224, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 7, 2, 234, 10, 2, 12, 2, 14, 2, 237, 11, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 284, 10, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 295, 10, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 308,
	10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 6, 2, 360, 10, 2, 13, 2, 14, 2, 361, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 388, 10, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 432, 10, 2, 12, 2, 14, 2, 435,
	11, 2, 5, 2, 437, 10, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	7, 2, 447, 10, 2, 12, 2, 14, 2, 450, 11, 2, 5, 2, 452, 10, 2, 3, 2, 7,
	2, 455, 10, 2, 12, 2, 14, 2, 458, 11, 2, 3, 2, 2, 3, 2, 3, 2, 2, 4, 4,
	2, 16, 16, 70, 70, 4, 2, 17, 17, 71, 71, 2, 548, 2, 387, 3, 2, 2, 2, 4,
	5, 8, 2, 1, 2, 5, 6, 7, 3, 2, 2, 6, 7, 5, 2, 2, 2, 7, 8, 7, 4, 2, 2, 8,
	388, 3, 2, 2, 2, 9, 10, 7, 72, 2, 2, 10, 11, 5, 2, 2, 2, 11, 12, 7, 73,
	2, 2, 12, 20, 5, 2, 2, 2, 13, 14, 7, 75, 2, 2, 14, 15, 5, 2, 2, 2, 15,
	16, 7, 73, 2, 2, 16, 17, 5, 2, 2, 2, 17, 19, 3, 2, 2, 2, 18, 13, 3, 2,
	2, 2, 19, 22, 3, 2, 2, 2, 20, 18, 3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 23,
	3, 2, 2, 2, 22, 20, 3, 2, 2, 2, 23, 24, 7, 74, 2, 2, 24, 25, 5, 2, 2, 2,
	25, 26, 7, 76, 2, 2, 26, 388, 3, 2, 2, 2, 27, 28, 7, 67, 2, 2, 28, 29,
	7, 3, 2, 2, 29, 30, 5, 2, 2, 2, 30, 31, 7, 15, 2, 2, 31, 32, 5, 2, 2, 2,
	32, 33, 7, 15, 2, 2, 33, 34, 5, 2, 2, 2, 34, 35, 7, 4, 2, 2, 35, 388, 3,
	2, 2, 2, 36, 37, 7, 19, 2, 2, 37, 38, 7, 3, 2, 2, 38, 39, 5, 2, 2, 2, 39,
	40, 7, 4, 2, 2, 40, 388, 3, 2, 2, 2, 41, 42, 7, 20, 2, 2, 42, 43, 7, 3,
	2, 2, 43, 44, 5, 2, 2, 2, 44, 45, 7, 4, 2, 2, 45, 388, 3, 2, 2, 2, 46,
	47, 7, 21, 2, 2, 47, 48, 7, 3, 2, 2, 48, 49, 5, 2, 2, 2, 49, 50, 7, 4,
	2, 2, 50, 388, 3, 2, 2, 2, 51, 52, 7, 22, 2, 2, 52, 53, 7, 3, 2, 2, 53,
	54, 5, 2, 2, 2, 54, 55, 7, 4, 2, 2, 55, 388, 3, 2, 2, 2, 56, 57, 7, 23,
	2, 2, 57, 58, 7, 3, 2, 2, 58, 59, 5, 2, 2, 2, 59, 60, 7, 15, 2, 2, 60,
	61, 5, 2, 2, 2, 61, 62, 7, 4, 2, 2, 62, 388, 3, 2, 2, 2, 63, 64, 7, 24,
	2, 2, 64, 65, 7, 3, 2, 2, 65, 70, 5, 2, 2, 2, 66, 67, 7, 15, 2, 2, 67,
	69, 5, 2, 2, 2, 68, 66, 3, 2, 2, 2, 69, 72, 3, 2, 2, 2, 70, 68, 3, 2, 2,
	2, 70, 71, 3, 2, 2, 2, 71, 73, 3, 2, 2, 2, 72, 70, 3, 2, 2, 2, 73, 74,
	7, 4, 2, 2, 74, 388, 3, 2, 2, 2, 75, 76, 7, 25, 2, 2, 76, 77, 7, 3, 2,
	2, 77, 78, 5, 2, 2, 2, 78, 79, 7, 4, 2, 2, 79, 388, 3, 2, 2, 2, 80, 81,
	7, 26, 2, 2, 81, 82, 7, 3, 2, 2, 82, 83, 5, 2, 2, 2, 83, 84, 7, 4, 2, 2,
	84, 388, 3, 2, 2, 2, 85, 86, 7, 27, 2, 2, 86, 87, 7, 3, 2, 2, 87, 88, 5,
	2, 2, 2, 88, 89, 7, 4, 2, 2, 89, 388, 3, 2, 2, 2, 90, 91, 7, 28, 2, 2,
	91, 92, 7, 3, 2, 2, 92, 93, 5, 2, 2, 2, 93, 94, 7, 15, 2, 2, 94, 97, 5,
	2, 2, 2, 95, 96, 7, 15, 2, 2, 96, 98, 5, 2, 2, 2, 97, 95, 3, 2, 2, 2, 97,
	98, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 100, 7, 4, 2, 2, 100, 388, 3, 2,
	2, 2, 101, 102, 7, 29, 2, 2, 102, 103, 7, 3, 2, 2, 103, 104, 5, 2, 2, 2,
	104, 105, 7, 4, 2, 2, 105, 388, 3, 2, 2, 2, 106, 107, 7, 30, 2, 2, 107,
	108, 7, 3, 2, 2, 108, 109, 5, 2, 2, 2, 109, 110, 7, 4, 2, 2, 110, 388,
	3, 2, 2, 2, 111, 112, 7, 31, 2, 2, 112, 113, 7, 3, 2, 2, 113, 114, 5, 2,
	2, 2, 114, 115, 7, 4, 2, 2, 115, 388, 3, 2, 2, 2, 116, 117, 7, 32, 2, 2,
	117, 118, 7, 3, 2, 2, 118, 119, 5, 2, 2, 2, 119, 120, 7, 15, 2, 2, 120,
	121, 5, 2, 2, 2, 121, 122, 7, 15, 2, 2, 122, 123, 5, 2, 2, 2, 123, 124,
	7, 15, 2, 2, 124, 125, 5, 2, 2, 2, 125, 126, 7, 4, 2, 2, 126, 388, 3, 2,
	2, 2, 127, 128, 7, 33, 2, 2, 128, 129, 7, 3, 2, 2, 129, 130, 5, 2, 2, 2,
	130, 131, 7, 15, 2, 2, 131, 134, 5, 2, 2, 2, 132, 133, 7, 15, 2, 2, 133,
	135, 5, 2, 2, 2, 134, 132, 3, 2, 2, 2, 134, 135, 3, 2, 2, 2, 135, 136,
	3, 2, 2, 2, 136, 137, 7, 4, 2, 2, 137, 388, 3, 2, 2, 2, 138, 139, 7, 35,
	2, 2, 139, 140, 7, 3, 2, 2, 140, 141, 5, 2, 2, 2, 141, 142, 7, 4, 2, 2,
	142, 388, 3, 2, 2, 2, 143, 144, 7, 34, 2, 2, 144, 145, 7, 3, 2, 2, 145,
	146, 5, 2, 2, 2, 146, 147, 7, 15, 2, 2, 147, 148, 5, 2, 2, 2, 148, 149,
	7, 4, 2, 2, 149, 388, 3, 2, 2, 2, 150, 151, 7, 36, 2, 2, 151, 152, 7, 3,
	2, 2, 152, 153, 5, 2, 2, 2, 153, 154, 7, 4, 2, 2, 154, 388, 3, 2, 2, 2,
	155, 156, 7, 37, 2, 2, 156, 157, 7, 3, 2, 2, 157, 158, 5, 2, 2, 2, 158,
	159, 7, 15, 2, 2, 159, 160, 5, 2, 2, 2, 160, 161, 7, 4, 2, 2, 161, 388,
	3, 2, 2, 2, 162, 163, 7, 38, 2, 2, 163, 164, 7, 3, 2, 2, 164, 165, 5, 2,
	2, 2, 165, 166, 7, 4, 2, 2, 166, 388, 3, 2, 2, 2, 167, 168, 7, 39, 2, 2,
	168, 169, 7, 3, 2, 2, 169, 170, 5, 2, 2, 2, 170, 171, 7, 15, 2, 2, 171,
	172, 5, 2, 2, 2, 172, 173, 7, 4, 2, 2, 173, 388, 3, 2, 2, 2, 174, 175,
	7, 40, 2, 2, 175, 176, 7, 3, 2, 2, 176, 177, 5, 2, 2, 2, 177, 178, 7, 4,
	2, 2, 178, 388, 3, 2, 2, 2, 179, 180, 7, 41, 2, 2, 180, 181, 7, 3, 2, 2,
	181, 182, 5, 2, 2, 2, 182, 183, 7, 4, 2, 2, 183, 388, 3, 2, 2, 2, 184,
	185, 7, 42, 2, 2, 185, 186, 7, 3, 2, 2, 186, 187, 5, 2, 2, 2, 187, 188,
	7, 4, 2, 2, 188, 388, 3, 2, 2, 2, 189, 190, 7, 43, 2, 2, 190, 191, 7, 3,
	2, 2, 191, 192, 5, 2, 2, 2, 192, 193, 7, 4, 2, 2, 193, 388, 3, 2, 2, 2,
	194, 195, 7, 44, 2, 2, 195, 196, 7, 3, 2, 2, 196, 199, 5, 2, 2, 2, 197,
	198, 7, 15, 2, 2, 198, 200, 5, 2, 2, 2, 199, 197, 3, 2, 2, 2, 200, 201,
	3, 2, 2, 2, 201, 199, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 203, 3, 2,
	2, 2, 203, 204, 7, 4, 2, 2, 204, 388, 3, 2, 2, 2, 205, 206, 7, 45, 2, 2,
	206, 207, 7, 3, 2, 2, 207, 212, 5, 2, 2, 2, 208, 209, 7, 15, 2, 2, 209,
	211, 5, 2, 2, 2, 210, 208, 3, 2, 2, 2, 211, 214, 3, 2, 2, 2, 212, 210,
	3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 215, 3, 2, 2, 2, 214, 212, 3, 2,
	2, 2, 215, 216, 7, 4, 2, 2, 216, 388, 3, 2, 2, 2, 217, 218, 7, 46, 2, 2,
	218, 219, 7, 3, 2, 2, 219, 222, 5, 2, 2, 2, 220, 221, 7, 15, 2, 2, 221,
	223, 5, 2, 2, 2, 222, 220, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 222,
	3, 2, 2, 2, 224, 225, 3, 2, 2, 2, 225, 226, 3, 2, 2, 2, 226, 227, 7, 4,
	2, 2, 227, 388, 3, 2, 2, 2, 228, 229, 7, 47, 2, 2, 229, 230, 7, 3, 2, 2,
	230, 235, 5, 2, 2, 2, 231, 232, 7, 15, 2, 2, 232, 234, 5, 2, 2, 2, 233,
	231, 3, 2, 2, 2, 234, 237, 3, 2, 2, 2, 235, 233, 3, 2, 2, 2, 235, 236,
	3, 2, 2, 2, 236, 238, 3, 2, 2, 2, 237, 235, 3, 2, 2, 2, 238, 239, 7, 4,
	2, 2, 239, 388, 3, 2, 2, 2, 240, 241, 7, 48, 2, 2, 241, 388, 7, 18, 2,
	2, 242, 243, 7, 49, 2, 2, 243, 244, 7, 3, 2, 2, 244, 245, 5, 2, 2, 2, 245,
	246, 7, 15, 2, 2, 246, 247, 5, 2, 2, 2, 247, 248, 7, 15, 2, 2, 248, 249,
	5, 2, 2, 2, 249, 250, 7, 4, 2, 2, 250, 388, 3, 2, 2, 2, 251, 252, 7, 50,
	2, 2, 252, 253, 7, 3, 2, 2, 253, 254, 5, 2, 2, 2, 254, 255, 7, 15, 2, 2,
	255, 256, 5, 2, 2, 2, 256, 257, 7, 15, 2, 2, 257, 258, 5, 2, 2, 2, 258,
	259, 7, 4, 2, 2, 259, 388, 3, 2, 2, 2, 260, 261, 7, 51, 2, 2, 261, 388,
	7, 18, 2, 2, 262, 263, 7, 52, 2, 2, 263, 264, 7, 3, 2, 2, 264, 265, 5,
	2, 2, 2, 265, 266, 7, 15, 2, 2, 266, 267, 5, 2, 2, 2, 267, 268, 7, 4, 2,
	2, 268, 388, 3, 2, 2, 2, 269, 270, 7, 53, 2, 2, 270, 388, 7, 18, 2, 2,
	271, 272, 7, 54, 2, 2, 272, 273, 7, 3, 2, 2, 273, 274, 5, 2, 2, 2, 274,
	275, 7, 4, 2, 2, 275, 388, 3, 2, 2, 2, 276, 277, 7, 55, 2, 2, 277, 278,
	7, 3, 2, 2, 278, 279, 5, 2, 2, 2, 279, 280, 7, 15, 2, 2, 280, 283, 5, 2,
	2, 2, 281, 282, 7, 15, 2, 2, 282, 284, 5, 2, 2, 2, 283, 281, 3, 2, 2, 2,
	283, 284, 3, 2, 2, 2, 284, 285, 3, 2, 2, 2, 285, 286, 7, 4, 2, 2, 286,
	388, 3, 2, 2, 2, 287, 288, 7, 56, 2, 2, 288, 289, 7, 3, 2, 2, 289, 290,
	5, 2, 2, 2, 290, 291, 7, 15, 2, 2, 291, 294, 5, 2, 2, 2, 292, 293, 7, 15,
	2, 2, 293, 295, 5, 2, 2, 2, 294, 292, 3, 2, 2, 2, 294, 295, 3, 2, 2, 2,
	295, 296, 3, 2, 2, 2, 296, 297, 7, 4, 2, 2, 297, 388, 3, 2, 2, 2, 298,
	299, 7, 57, 2, 2, 299, 300, 7, 3, 2, 2, 300, 301, 5, 2, 2, 2, 301, 302,
	7, 15, 2, 2, 302, 303, 5, 2, 2, 2, 303, 304, 7, 15, 2, 2, 304, 307, 5,
	2, 2, 2, 305, 306, 7, 15, 2, 2, 306, 308, 5, 2, 2, 2, 307, 305, 3, 2, 2,
	2, 307, 308, 3, 2, 2, 2, 308, 309, 3, 2, 2, 2, 309, 310, 7, 4, 2, 2, 310,
	388, 3, 2, 2, 2, 311, 312, 7, 58, 2, 2, 312, 313, 7, 3, 2, 2, 313, 314,
	5, 2, 2, 2, 314, 315, 7, 15, 2, 2, 315, 316, 5, 2, 2, 2, 316, 317, 7, 15,
	2, 2, 317, 318, 5, 2, 2, 2, 318, 319, 7, 4, 2, 2, 319, 388, 3, 2, 2, 2,
	320, 321, 7, 59, 2, 2, 321, 322, 7, 3, 2, 2, 322, 323, 5, 2, 2, 2, 323,
	324, 7, 15, 2, 2, 324, 325, 5, 2, 2, 2, 325, 326, 7, 4, 2, 2, 326, 388,
	3, 2, 2, 2, 327, 328, 7, 60, 2, 2, 328, 329, 7, 3, 2, 2, 329, 330, 5, 2,
	2, 2, 330, 331, 7, 15, 2, 2, 331, 332, 5, 2, 2, 2, 332, 333, 7, 4, 2, 2,
	333, 388, 3, 2, 2, 2, 334, 335, 7, 61, 2, 2, 335, 336, 7, 3, 2, 2, 336,
	337, 5, 2, 2, 2, 337, 338, 7, 4, 2, 2, 338, 388, 3, 2, 2, 2, 339, 340,
	7, 62, 2, 2, 340, 341, 7, 3, 2, 2, 341, 342, 5, 2, 2, 2, 342, 343, 7, 4,
	2, 2, 343, 388, 3, 2, 2, 2, 344, 345, 7, 63, 2, 2, 345, 346, 7, 3, 2, 2,
	346, 347, 5, 2, 2, 2, 347, 348, 7, 4, 2, 2, 348, 388, 3, 2, 2, 2, 349,
	350, 7, 64, 2, 2, 350, 351, 7, 3, 2, 2, 351, 352, 5, 2, 2, 2, 352, 353,
	7, 15, 2, 2, 353, 359, 5, 2, 2, 2, 354, 355, 7, 15, 2, 2, 355, 356, 5,
	2, 2, 2, 356, 357, 7, 15, 2, 2, 357, 358, 5, 2, 2, 2, 358, 360, 3, 2, 2,
	2, 359, 354, 3, 2, 2, 2, 360, 361, 3, 2, 2, 2, 361, 359, 3, 2, 2, 2, 361,
	362, 3, 2, 2, 2, 362, 363, 3, 2, 2, 2, 363, 364, 7, 4, 2, 2, 364, 388,
	3, 2, 2, 2, 365, 366, 7, 65, 2, 2, 366, 367, 7, 3, 2, 2, 367, 368, 5, 2,
	2, 2, 368, 369, 7, 4, 2, 2, 369, 388, 3, 2, 2, 2, 370, 371, 7, 66, 2, 2,
	371, 372, 7, 3, 2, 2, 372, 373, 5, 2, 2, 2, 373, 374, 7, 4, 2, 2, 374,
	388, 3, 2, 2, 2, 375, 388, 7, 78, 2, 2, 376, 377, 7, 8, 2, 2, 377, 388,
	7, 78, 2, 2, 378, 388, 7, 79, 2, 2, 379, 380, 7, 8, 2, 2, 380, 388, 7,
	79, 2, 2, 381, 388, 7, 83, 2, 2, 382, 388, 7, 84, 2, 2, 383, 388, 7, 81,
	2, 2, 384, 388, 7, 80, 2, 2, 385, 388, 7, 77, 2, 2, 386, 388, 7, 82, 2,
	2, 387, 4, 3, 2, 2, 2, 387, 9, 3, 2, 2, 2, 387, 27, 3, 2, 2, 2, 387, 36,
	3, 2, 2, 2, 387, 41, 3, 2, 2, 2, 387, 46, 3, 2, 2, 2, 387, 51, 3, 2, 2,
	2, 387, 56, 3, 2, 2, 2, 387, 63, 3, 2, 2, 2, 387, 75, 3, 2, 2, 2, 387,
	80, 3, 2, 2, 2, 387, 85, 3, 2, 2, 2, 387, 90, 3, 2, 2, 2, 387, 101, 3,
	2, 2, 2, 387, 106, 3, 2, 2, 2, 387, 111, 3, 2, 2, 2, 387, 116, 3, 2, 2,
	2, 387, 127, 3, 2, 2, 2, 387, 138, 3, 2, 2, 2, 387, 143, 3, 2, 2, 2, 387,
	150, 3, 2, 2, 2, 387, 155, 3, 2, 2, 2, 387, 162, 3, 2, 2, 2, 387, 167,
	3, 2, 2, 2, 387, 174, 3, 2, 2, 2, 387, 179, 3, 2, 2, 2, 387, 184, 3, 2,
	2, 2, 387, 189, 3, 2, 2, 2, 387, 194, 3, 2, 2, 2, 387, 205, 3, 2, 2, 2,
	387, 217, 3, 2, 2, 2, 387, 228, 3, 2, 2, 2, 387, 240, 3, 2, 2, 2, 387,
	242, 3, 2, 2, 2, 387, 251, 3, 2, 2, 2, 387, 260, 3, 2, 2, 2, 387, 262,
	3, 2, 2, 2, 387, 269, 3, 2, 2, 2, 387, 271, 3, 2, 2, 2, 387, 276, 3, 2,
	2, 2, 387, 287, 3, 2, 2, 2, 387, 298, 3, 2, 2, 2, 387, 311, 3, 2, 2, 2,
	387, 320, 3, 2, 2, 2, 387, 327, 3, 2, 2, 2, 387, 334, 3, 2, 2, 2, 387,
	339, 3, 2, 2, 2, 387, 344, 3, 2, 2, 2, 387, 349, 3, 2, 2, 2, 387, 365,
	3, 2, 2, 2, 387, 370, 3, 2, 2, 2, 387, 375, 3, 2, 2, 2, 387, 376, 3, 2,
	2, 2, 387, 378, 3, 2, 2, 2, 387, 379, 3, 2, 2, 2, 387, 381, 3, 2, 2, 2,
	387, 382, 3, 2, 2, 2, 387, 383, 3, 2, 2, 2, 387, 384, 3, 2, 2, 2, 387,
	385, 3, 2, 2, 2, 387, 386, 3, 2, 2, 2, 388, 456, 3, 2, 2, 2, 389, 390,
	12, 76, 2, 2, 390, 391, 7, 5, 2, 2, 391, 455, 5, 2, 2, 77, 392, 393, 12,
	75, 2, 2, 393, 394, 7, 6, 2, 2, 394, 455, 5, 2, 2, 76, 395, 396, 12, 74,
	2, 2, 396, 397, 7, 7, 2, 2, 397, 455, 5, 2, 2, 75, 398, 399, 12, 73, 2,
	2, 399, 400, 7, 8, 2, 2, 400, 455, 5, 2, 2, 74, 401, 402, 12, 72, 2, 2,
	402, 403, 7, 9, 2, 2, 403, 455, 5, 2, 2, 73, 404, 405, 12, 71, 2, 2, 405,
	406, 7, 10, 2, 2, 406, 455, 5, 2, 2, 72, 407, 408, 12, 70, 2, 2, 408, 409,
	7, 11, 2, 2, 409, 455, 5, 2, 2, 71, 410, 411, 12, 69, 2, 2, 411, 412, 7,
	12, 2, 2, 412, 455, 5, 2, 2, 70, 413, 414, 12, 68, 2, 2, 414, 415, 7, 13,
	2, 2, 415, 455, 5, 2, 2, 69, 416, 417, 12, 67, 2, 2, 417, 418, 7, 14, 2,
	2, 418, 455, 5, 2, 2, 68, 419, 420, 12, 64, 2, 2, 420, 421, 9, 2, 2, 2,
	421, 455, 5, 2, 2, 65, 422, 423, 12, 63, 2, 2, 423, 424, 9, 3, 2, 2, 424,
	455, 5, 2, 2, 64, 425, 426, 12, 66, 2, 2, 426, 427, 7, 68, 2, 2, 427, 436,
	7, 3, 2, 2, 428, 433, 5, 2, 2, 2, 429, 430, 7, 15, 2, 2, 430, 432, 5, 2,
	2, 2, 431, 429, 3, 2, 2, 2, 432, 435, 3, 2, 2, 2, 433, 431, 3, 2, 2, 2,
	433, 434, 3, 2, 2, 2, 434, 437, 3, 2, 2, 2, 435, 433, 3, 2, 2, 2, 436,
	428, 3, 2, 2, 2, 436, 437, 3, 2, 2, 2, 437, 438, 3, 2, 2, 2, 438, 455,
	7, 4, 2, 2, 439, 440, 12, 65, 2, 2, 440, 441, 7, 69, 2, 2, 441, 442, 7,
	68, 2, 2, 442, 451, 7, 3, 2, 2, 443, 448, 5, 2, 2, 2, 444, 445, 7, 15,
	2, 2, 445, 447, 5, 2, 2, 2, 446, 444, 3, 2, 2, 2, 447, 450, 3, 2, 2, 2,
	448, 446, 3, 2, 2, 2, 448, 449, 3, 2, 2, 2, 449, 452, 3, 2, 2, 2, 450,
	448, 3, 2, 2, 2, 451, 443, 3, 2, 2, 2, 451, 452, 3, 2, 2, 2, 452, 453,
	3, 2, 2, 2, 453, 455, 7, 4, 2, 2, 454, 389, 3, 2, 2, 2, 454, 392, 3, 2,
	2, 2, 454, 395, 3, 2, 2, 2, 454, 398, 3, 2, 2, 2, 454, 401, 3, 2, 2, 2,
	454, 404, 3, 2, 2, 2, 454, 407, 3, 2, 2, 2, 454, 410, 3, 2, 2, 2, 454,
	413, 3, 2, 2, 2, 454, 416, 3, 2, 2, 2, 454, 419, 3, 2, 2, 2, 454, 422,
	3, 2, 2, 2, 454, 425, 3, 2, 2, 2, 454, 439, 3, 2, 2, 2, 455, 458, 3, 2,
	2, 2, 456, 454, 3, 2, 2, 2, 456, 457, 3, 2, 2, 2, 457, 3, 3, 2, 2, 2, 458,
	456, 3, 2, 2, 2, 21, 20, 70, 97, 134, 201, 212, 224, 235, 283, 294, 307,
	361, 387, 433, 436, 448, 451, 454, 456,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "')'", "'*'", "'/'", "'+'", "'-'", "'='", "'>'", "'>='", "'<'",
	"'<='", "'!='", "','", "'&&'", "'||'", "'()'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "Abs",
	"Acos", "Asin", "Atan", "Atan2", "Average", "Ceil", "CharFromInt", "CharToInt",
	"Contains", "Cos", "Cosh", "CountWords", "Distance", "EndsWith", "FindString",
	"Exp", "Floor", "GetWord", "HexToNumber", "Left", "Length", "Log", "Log10",
	"Lowercase", "Max", "Median", "Min", "Mod", "Null", "PadLeft", "PadRight",
	"Pi", "Pow", "Rand", "RandInt", "Regex_CountMatches", "Regex_Match", "Regex_Replace",
	"Replace", "Right", "Round", "Sin", "Sinh", "Sqrt", "Switch", "Tan", "Tanh",
	"Iif", "In", "Not", "And", "Or", "If", "Then", "Else", "Elseif", "Endif",
	"Bool", "Integer", "Decimal", "Date", "Datetime", "Field", "SingleQuoteString",
	"DoubleQuoteString", "Whitespace",
}

var ruleNames = []string{
	"expr",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type AlteryxFormulasParser struct {
	*antlr.BaseParser
}

func NewAlteryxFormulasParser(input antlr.TokenStream) *AlteryxFormulasParser {
	this := new(AlteryxFormulasParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "AlteryxFormulas.g4"

	return this
}

// AlteryxFormulasParser tokens.
const (
	AlteryxFormulasParserEOF                = antlr.TokenEOF
	AlteryxFormulasParserT__0               = 1
	AlteryxFormulasParserT__1               = 2
	AlteryxFormulasParserT__2               = 3
	AlteryxFormulasParserT__3               = 4
	AlteryxFormulasParserT__4               = 5
	AlteryxFormulasParserT__5               = 6
	AlteryxFormulasParserT__6               = 7
	AlteryxFormulasParserT__7               = 8
	AlteryxFormulasParserT__8               = 9
	AlteryxFormulasParserT__9               = 10
	AlteryxFormulasParserT__10              = 11
	AlteryxFormulasParserT__11              = 12
	AlteryxFormulasParserT__12              = 13
	AlteryxFormulasParserT__13              = 14
	AlteryxFormulasParserT__14              = 15
	AlteryxFormulasParserT__15              = 16
	AlteryxFormulasParserAbs                = 17
	AlteryxFormulasParserAcos               = 18
	AlteryxFormulasParserAsin               = 19
	AlteryxFormulasParserAtan               = 20
	AlteryxFormulasParserAtan2              = 21
	AlteryxFormulasParserAverage            = 22
	AlteryxFormulasParserCeil               = 23
	AlteryxFormulasParserCharFromInt        = 24
	AlteryxFormulasParserCharToInt          = 25
	AlteryxFormulasParserContains           = 26
	AlteryxFormulasParserCos                = 27
	AlteryxFormulasParserCosh               = 28
	AlteryxFormulasParserCountWords         = 29
	AlteryxFormulasParserDistance           = 30
	AlteryxFormulasParserEndsWith           = 31
	AlteryxFormulasParserFindString         = 32
	AlteryxFormulasParserExp                = 33
	AlteryxFormulasParserFloor              = 34
	AlteryxFormulasParserGetWord            = 35
	AlteryxFormulasParserHexToNumber        = 36
	AlteryxFormulasParserLeft               = 37
	AlteryxFormulasParserLength             = 38
	AlteryxFormulasParserLog                = 39
	AlteryxFormulasParserLog10              = 40
	AlteryxFormulasParserLowercase          = 41
	AlteryxFormulasParserMax                = 42
	AlteryxFormulasParserMedian             = 43
	AlteryxFormulasParserMin                = 44
	AlteryxFormulasParserMod                = 45
	AlteryxFormulasParserNull               = 46
	AlteryxFormulasParserPadLeft            = 47
	AlteryxFormulasParserPadRight           = 48
	AlteryxFormulasParserPi                 = 49
	AlteryxFormulasParserPow                = 50
	AlteryxFormulasParserRand               = 51
	AlteryxFormulasParserRandInt            = 52
	AlteryxFormulasParserRegex_CountMatches = 53
	AlteryxFormulasParserRegex_Match        = 54
	AlteryxFormulasParserRegex_Replace      = 55
	AlteryxFormulasParserReplace            = 56
	AlteryxFormulasParserRight              = 57
	AlteryxFormulasParserRound              = 58
	AlteryxFormulasParserSin                = 59
	AlteryxFormulasParserSinh               = 60
	AlteryxFormulasParserSqrt               = 61
	AlteryxFormulasParserSwitch             = 62
	AlteryxFormulasParserTan                = 63
	AlteryxFormulasParserTanh               = 64
	AlteryxFormulasParserIif                = 65
	AlteryxFormulasParserIn                 = 66
	AlteryxFormulasParserNot                = 67
	AlteryxFormulasParserAnd                = 68
	AlteryxFormulasParserOr                 = 69
	AlteryxFormulasParserIf                 = 70
	AlteryxFormulasParserThen               = 71
	AlteryxFormulasParserElse               = 72
	AlteryxFormulasParserElseif             = 73
	AlteryxFormulasParserEndif              = 74
	AlteryxFormulasParserBool               = 75
	AlteryxFormulasParserInteger            = 76
	AlteryxFormulasParserDecimal            = 77
	AlteryxFormulasParserDate               = 78
	AlteryxFormulasParserDatetime           = 79
	AlteryxFormulasParserField              = 80
	AlteryxFormulasParserSingleQuoteString  = 81
	AlteryxFormulasParserDoubleQuoteString  = 82
	AlteryxFormulasParserWhitespace         = 83
)

// AlteryxFormulasParserRULE_expr is the AlteryxFormulasParser rule.
const AlteryxFormulasParserRULE_expr = 0

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AlteryxFormulasParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AlteryxFormulasParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MaxFuncContext struct {
	*ExprContext
}

func NewMaxFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MaxFuncContext {
	var p = new(MaxFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MaxFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MaxFuncContext) Max() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserMax, 0)
}

func (s *MaxFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MaxFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MaxFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterMaxFunc(s)
	}
}

func (s *MaxFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitMaxFunc(s)
	}
}

type CosFuncContext struct {
	*ExprContext
}

func NewCosFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CosFuncContext {
	var p = new(CosFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CosFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CosFuncContext) Cos() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserCos, 0)
}

func (s *CosFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CosFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterCosFunc(s)
	}
}

func (s *CosFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitCosFunc(s)
	}
}

type ReplaceFuncContext struct {
	*ExprContext
}

func NewReplaceFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReplaceFuncContext {
	var p = new(ReplaceFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ReplaceFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReplaceFuncContext) Replace() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserReplace, 0)
}

func (s *ReplaceFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ReplaceFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ReplaceFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterReplaceFunc(s)
	}
}

func (s *ReplaceFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitReplaceFunc(s)
	}
}

type RandIntFuncContext struct {
	*ExprContext
}

func NewRandIntFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RandIntFuncContext {
	var p = new(RandIntFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *RandIntFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RandIntFuncContext) RandInt() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserRandInt, 0)
}

func (s *RandIntFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RandIntFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterRandIntFunc(s)
	}
}

func (s *RandIntFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitRandIntFunc(s)
	}
}

type CharToIntFuncContext struct {
	*ExprContext
}

func NewCharToIntFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CharToIntFuncContext {
	var p = new(CharToIntFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CharToIntFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharToIntFuncContext) CharToInt() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserCharToInt, 0)
}

func (s *CharToIntFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CharToIntFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterCharToIntFunc(s)
	}
}

func (s *CharToIntFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitCharToIntFunc(s)
	}
}

type LeftFuncContext struct {
	*ExprContext
}

func NewLeftFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LeftFuncContext {
	var p = new(LeftFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LeftFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LeftFuncContext) Left() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserLeft, 0)
}

func (s *LeftFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *LeftFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LeftFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterLeftFunc(s)
	}
}

func (s *LeftFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitLeftFunc(s)
	}
}

type TanFuncContext struct {
	*ExprContext
}

func NewTanFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TanFuncContext {
	var p = new(TanFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *TanFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TanFuncContext) Tan() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserTan, 0)
}

func (s *TanFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TanFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterTanFunc(s)
	}
}

func (s *TanFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitTanFunc(s)
	}
}

type RegexReplaceFuncContext struct {
	*ExprContext
}

func NewRegexReplaceFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RegexReplaceFuncContext {
	var p = new(RegexReplaceFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *RegexReplaceFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexReplaceFuncContext) Regex_Replace() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserRegex_Replace, 0)
}

func (s *RegexReplaceFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *RegexReplaceFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RegexReplaceFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterRegexReplaceFunc(s)
	}
}

func (s *RegexReplaceFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitRegexReplaceFunc(s)
	}
}

type FloorFuncContext struct {
	*ExprContext
}

func NewFloorFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FloorFuncContext {
	var p = new(FloorFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FloorFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloorFuncContext) Floor() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserFloor, 0)
}

func (s *FloorFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FloorFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterFloorFunc(s)
	}
}

func (s *FloorFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitFloorFunc(s)
	}
}

type NotInContext struct {
	*ExprContext
}

func NewNotInContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotInContext {
	var p = new(NotInContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NotInContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotInContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *NotInContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotInContext) Not() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserNot, 0)
}

func (s *NotInContext) In() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserIn, 0)
}

func (s *NotInContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNotIn(s)
	}
}

func (s *NotInContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNotIn(s)
	}
}

type TanhFuncContext struct {
	*ExprContext
}

func NewTanhFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TanhFuncContext {
	var p = new(TanhFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *TanhFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TanhFuncContext) Tanh() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserTanh, 0)
}

func (s *TanhFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *TanhFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterTanhFunc(s)
	}
}

func (s *TanhFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitTanhFunc(s)
	}
}

type GreaterThanContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewGreaterThanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterThanContext {
	var p = new(GreaterThanContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *GreaterThanContext) GetLeft() IExprContext { return s.left }

func (s *GreaterThanContext) GetRight() IExprContext { return s.right }

func (s *GreaterThanContext) SetLeft(v IExprContext) { s.left = v }

func (s *GreaterThanContext) SetRight(v IExprContext) { s.right = v }

func (s *GreaterThanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterThanContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *GreaterThanContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GreaterThanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterGreaterThan(s)
	}
}

func (s *GreaterThanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitGreaterThan(s)
	}
}

type AsinFuncContext struct {
	*ExprContext
}

func NewAsinFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AsinFuncContext {
	var p = new(AsinFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AsinFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsinFuncContext) Asin() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserAsin, 0)
}

func (s *AsinFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AsinFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterAsinFunc(s)
	}
}

func (s *AsinFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitAsinFunc(s)
	}
}

type AddContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewAddContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddContext {
	var p = new(AddContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AddContext) GetLeft() IExprContext { return s.left }

func (s *AddContext) GetRight() IExprContext { return s.right }

func (s *AddContext) SetLeft(v IExprContext) { s.left = v }

func (s *AddContext) SetRight(v IExprContext) { s.right = v }

func (s *AddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AddContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterAdd(s)
	}
}

func (s *AddContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitAdd(s)
	}
}

type CoshFuncContext struct {
	*ExprContext
}

func NewCoshFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CoshFuncContext {
	var p = new(CoshFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CoshFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CoshFuncContext) Cosh() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserCosh, 0)
}

func (s *CoshFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CoshFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterCoshFunc(s)
	}
}

func (s *CoshFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitCoshFunc(s)
	}
}

type InContext struct {
	*ExprContext
}

func NewInContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InContext {
	var p = new(InContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *InContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *InContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *InContext) In() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserIn, 0)
}

func (s *InContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterIn(s)
	}
}

func (s *InContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitIn(s)
	}
}

type SubtractContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewSubtractContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SubtractContext {
	var p = new(SubtractContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SubtractContext) GetLeft() IExprContext { return s.left }

func (s *SubtractContext) GetRight() IExprContext { return s.right }

func (s *SubtractContext) SetLeft(v IExprContext) { s.left = v }

func (s *SubtractContext) SetRight(v IExprContext) { s.right = v }

func (s *SubtractContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SubtractContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *SubtractContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SubtractContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterSubtract(s)
	}
}

func (s *SubtractContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitSubtract(s)
	}
}

type NotEqualContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewNotEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotEqualContext {
	var p = new(NotEqualContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NotEqualContext) GetLeft() IExprContext { return s.left }

func (s *NotEqualContext) GetRight() IExprContext { return s.right }

func (s *NotEqualContext) SetLeft(v IExprContext) { s.left = v }

func (s *NotEqualContext) SetRight(v IExprContext) { s.right = v }

func (s *NotEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotEqualContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *NotEqualContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *NotEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNotEqual(s)
	}
}

func (s *NotEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNotEqual(s)
	}
}

type ParenthesisContext struct {
	*ExprContext
}

func NewParenthesisContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesisContext {
	var p = new(ParenthesisContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesisContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterParenthesis(s)
	}
}

func (s *ParenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitParenthesis(s)
	}
}

type HexToNumberFuncContext struct {
	*ExprContext
}

func NewHexToNumberFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *HexToNumberFuncContext {
	var p = new(HexToNumberFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *HexToNumberFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HexToNumberFuncContext) HexToNumber() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserHexToNumber, 0)
}

func (s *HexToNumberFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *HexToNumberFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterHexToNumberFunc(s)
	}
}

func (s *HexToNumberFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitHexToNumberFunc(s)
	}
}

type MedianFuncContext struct {
	*ExprContext
}

func NewMedianFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MedianFuncContext {
	var p = new(MedianFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MedianFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MedianFuncContext) Median() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserMedian, 0)
}

func (s *MedianFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MedianFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MedianFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterMedianFunc(s)
	}
}

func (s *MedianFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitMedianFunc(s)
	}
}

type NullFuncContext struct {
	*ExprContext
}

func NewNullFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullFuncContext {
	var p = new(NullFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NullFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullFuncContext) Null() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserNull, 0)
}

func (s *NullFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNullFunc(s)
	}
}

func (s *NullFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNullFunc(s)
	}
}

type DatetimeLiteralContext struct {
	*ExprContext
}

func NewDatetimeLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DatetimeLiteralContext {
	var p = new(DatetimeLiteralContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DatetimeLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatetimeLiteralContext) Datetime() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserDatetime, 0)
}

func (s *DatetimeLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDatetimeLiteral(s)
	}
}

func (s *DatetimeLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDatetimeLiteral(s)
	}
}

type StringLiteralContext struct {
	*ExprContext
}

func NewStringLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) SingleQuoteString() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserSingleQuoteString, 0)
}

func (s *StringLiteralContext) DoubleQuoteString() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserDoubleQuoteString, 0)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

type RegexMatchFuncContext struct {
	*ExprContext
}

func NewRegexMatchFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RegexMatchFuncContext {
	var p = new(RegexMatchFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *RegexMatchFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexMatchFuncContext) Regex_Match() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserRegex_Match, 0)
}

func (s *RegexMatchFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *RegexMatchFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RegexMatchFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterRegexMatchFunc(s)
	}
}

func (s *RegexMatchFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitRegexMatchFunc(s)
	}
}

type DateLiteralContext struct {
	*ExprContext
}

func NewDateLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DateLiteralContext {
	var p = new(DateLiteralContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DateLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DateLiteralContext) Date() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserDate, 0)
}

func (s *DateLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDateLiteral(s)
	}
}

func (s *DateLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDateLiteral(s)
	}
}

type AbsFuncContext struct {
	*ExprContext
}

func NewAbsFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AbsFuncContext {
	var p = new(AbsFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AbsFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AbsFuncContext) Abs() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserAbs, 0)
}

func (s *AbsFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AbsFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterAbsFunc(s)
	}
}

func (s *AbsFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitAbsFunc(s)
	}
}

type Atan2FuncContext struct {
	*ExprContext
}

func NewAtan2FuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Atan2FuncContext {
	var p = new(Atan2FuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *Atan2FuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Atan2FuncContext) Atan2() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserAtan2, 0)
}

func (s *Atan2FuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *Atan2FuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Atan2FuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterAtan2Func(s)
	}
}

func (s *Atan2FuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitAtan2Func(s)
	}
}

type PadLeftFuncContext struct {
	*ExprContext
}

func NewPadLeftFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PadLeftFuncContext {
	var p = new(PadLeftFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PadLeftFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PadLeftFuncContext) PadLeft() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserPadLeft, 0)
}

func (s *PadLeftFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *PadLeftFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PadLeftFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterPadLeftFunc(s)
	}
}

func (s *PadLeftFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitPadLeftFunc(s)
	}
}

type FindStringFuncContext struct {
	*ExprContext
}

func NewFindStringFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FindStringFuncContext {
	var p = new(FindStringFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *FindStringFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FindStringFuncContext) FindString() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserFindString, 0)
}

func (s *FindStringFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *FindStringFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *FindStringFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterFindStringFunc(s)
	}
}

func (s *FindStringFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitFindStringFunc(s)
	}
}

type Log10FuncContext struct {
	*ExprContext
}

func NewLog10FuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *Log10FuncContext {
	var p = new(Log10FuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *Log10FuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Log10FuncContext) Log10() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserLog10, 0)
}

func (s *Log10FuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *Log10FuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterLog10Func(s)
	}
}

func (s *Log10FuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitLog10Func(s)
	}
}

type CeilFuncContext struct {
	*ExprContext
}

func NewCeilFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CeilFuncContext {
	var p = new(CeilFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CeilFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CeilFuncContext) Ceil() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserCeil, 0)
}

func (s *CeilFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CeilFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterCeilFunc(s)
	}
}

func (s *CeilFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitCeilFunc(s)
	}
}

type EndsWithFuncContext struct {
	*ExprContext
}

func NewEndsWithFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EndsWithFuncContext {
	var p = new(EndsWithFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EndsWithFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndsWithFuncContext) EndsWith() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserEndsWith, 0)
}

func (s *EndsWithFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *EndsWithFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EndsWithFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterEndsWithFunc(s)
	}
}

func (s *EndsWithFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitEndsWithFunc(s)
	}
}

type DistanceFuncContext struct {
	*ExprContext
}

func NewDistanceFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DistanceFuncContext {
	var p = new(DistanceFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DistanceFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DistanceFuncContext) Distance() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserDistance, 0)
}

func (s *DistanceFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *DistanceFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DistanceFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDistanceFunc(s)
	}
}

func (s *DistanceFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDistanceFunc(s)
	}
}

type CharFromIntFuncContext struct {
	*ExprContext
}

func NewCharFromIntFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CharFromIntFuncContext {
	var p = new(CharFromIntFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CharFromIntFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CharFromIntFuncContext) CharFromInt() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserCharFromInt, 0)
}

func (s *CharFromIntFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CharFromIntFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterCharFromIntFunc(s)
	}
}

func (s *CharFromIntFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitCharFromIntFunc(s)
	}
}

type BoolLiteralContext struct {
	*ExprContext
}

func NewBoolLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolLiteralContext {
	var p = new(BoolLiteralContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *BoolLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolLiteralContext) Bool() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserBool, 0)
}

func (s *BoolLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterBoolLiteral(s)
	}
}

func (s *BoolLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitBoolLiteral(s)
	}
}

type AcosFuncContext struct {
	*ExprContext
}

func NewAcosFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AcosFuncContext {
	var p = new(AcosFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AcosFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AcosFuncContext) Acos() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserAcos, 0)
}

func (s *AcosFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AcosFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterAcosFunc(s)
	}
}

func (s *AcosFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitAcosFunc(s)
	}
}

type SqrtFuncContext struct {
	*ExprContext
}

func NewSqrtFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SqrtFuncContext {
	var p = new(SqrtFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SqrtFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SqrtFuncContext) Sqrt() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserSqrt, 0)
}

func (s *SqrtFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SqrtFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterSqrtFunc(s)
	}
}

func (s *SqrtFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitSqrtFunc(s)
	}
}

type SinhFuncContext struct {
	*ExprContext
}

func NewSinhFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SinhFuncContext {
	var p = new(SinhFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SinhFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SinhFuncContext) Sinh() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserSinh, 0)
}

func (s *SinhFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SinhFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterSinhFunc(s)
	}
}

func (s *SinhFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitSinhFunc(s)
	}
}

type RandFuncContext struct {
	*ExprContext
}

func NewRandFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RandFuncContext {
	var p = new(RandFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *RandFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RandFuncContext) Rand() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserRand, 0)
}

func (s *RandFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterRandFunc(s)
	}
}

func (s *RandFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitRandFunc(s)
	}
}

type MinFuncContext struct {
	*ExprContext
}

func NewMinFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MinFuncContext {
	var p = new(MinFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MinFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MinFuncContext) Min() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserMin, 0)
}

func (s *MinFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MinFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MinFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterMinFunc(s)
	}
}

func (s *MinFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitMinFunc(s)
	}
}

type PiFuncContext struct {
	*ExprContext
}

func NewPiFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PiFuncContext {
	var p = new(PiFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PiFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PiFuncContext) Pi() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserPi, 0)
}

func (s *PiFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterPiFunc(s)
	}
}

func (s *PiFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitPiFunc(s)
	}
}

type AndContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndContext {
	var p = new(AndContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AndContext) GetLeft() IExprContext { return s.left }

func (s *AndContext) GetRight() IExprContext { return s.right }

func (s *AndContext) SetLeft(v IExprContext) { s.left = v }

func (s *AndContext) SetRight(v IExprContext) { s.right = v }

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AndContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AndContext) And() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserAnd, 0)
}

func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitAnd(s)
	}
}

type LessThanContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewLessThanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessThanContext {
	var p = new(LessThanContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LessThanContext) GetLeft() IExprContext { return s.left }

func (s *LessThanContext) GetRight() IExprContext { return s.right }

func (s *LessThanContext) SetLeft(v IExprContext) { s.left = v }

func (s *LessThanContext) SetRight(v IExprContext) { s.right = v }

func (s *LessThanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessThanContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *LessThanContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LessThanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterLessThan(s)
	}
}

func (s *LessThanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitLessThan(s)
	}
}

type PadRightFuncContext struct {
	*ExprContext
}

func NewPadRightFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PadRightFuncContext {
	var p = new(PadRightFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PadRightFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PadRightFuncContext) PadRight() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserPadRight, 0)
}

func (s *PadRightFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *PadRightFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PadRightFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterPadRightFunc(s)
	}
}

func (s *PadRightFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitPadRightFunc(s)
	}
}

type RegexCountMatchesFuncContext struct {
	*ExprContext
}

func NewRegexCountMatchesFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RegexCountMatchesFuncContext {
	var p = new(RegexCountMatchesFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *RegexCountMatchesFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexCountMatchesFuncContext) Regex_CountMatches() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserRegex_CountMatches, 0)
}

func (s *RegexCountMatchesFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *RegexCountMatchesFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RegexCountMatchesFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterRegexCountMatchesFunc(s)
	}
}

func (s *RegexCountMatchesFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitRegexCountMatchesFunc(s)
	}
}

type DivideContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewDivideContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DivideContext {
	var p = new(DivideContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DivideContext) GetLeft() IExprContext { return s.left }

func (s *DivideContext) GetRight() IExprContext { return s.right }

func (s *DivideContext) SetLeft(v IExprContext) { s.left = v }

func (s *DivideContext) SetRight(v IExprContext) { s.right = v }

func (s *DivideContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DivideContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *DivideContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DivideContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterDivide(s)
	}
}

func (s *DivideContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitDivide(s)
	}
}

type ExprIfContext struct {
	*ExprContext
	ifStmnt   IExprContext
	thenStmnt IExprContext
	elseStmnt IExprContext
}

func NewExprIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprIfContext {
	var p = new(ExprIfContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExprIfContext) GetIfStmnt() IExprContext { return s.ifStmnt }

func (s *ExprIfContext) GetThenStmnt() IExprContext { return s.thenStmnt }

func (s *ExprIfContext) GetElseStmnt() IExprContext { return s.elseStmnt }

func (s *ExprIfContext) SetIfStmnt(v IExprContext) { s.ifStmnt = v }

func (s *ExprIfContext) SetThenStmnt(v IExprContext) { s.thenStmnt = v }

func (s *ExprIfContext) SetElseStmnt(v IExprContext) { s.elseStmnt = v }

func (s *ExprIfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprIfContext) If() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserIf, 0)
}

func (s *ExprIfContext) AllThen() []antlr.TerminalNode {
	return s.GetTokens(AlteryxFormulasParserThen)
}

func (s *ExprIfContext) Then(i int) antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserThen, i)
}

func (s *ExprIfContext) Else() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserElse, 0)
}

func (s *ExprIfContext) Endif() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserEndif, 0)
}

func (s *ExprIfContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ExprIfContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExprIfContext) AllElseif() []antlr.TerminalNode {
	return s.GetTokens(AlteryxFormulasParserElseif)
}

func (s *ExprIfContext) Elseif(i int) antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserElseif, i)
}

func (s *ExprIfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterExprIf(s)
	}
}

func (s *ExprIfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitExprIf(s)
	}
}

type ContainsFuncContext struct {
	*ExprContext
}

func NewContainsFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContainsFuncContext {
	var p = new(ContainsFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ContainsFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContainsFuncContext) Contains() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserContains, 0)
}

func (s *ContainsFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ContainsFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ContainsFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterContainsFunc(s)
	}
}

func (s *ContainsFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitContainsFunc(s)
	}
}

type LowercaseFuncContext struct {
	*ExprContext
}

func NewLowercaseFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LowercaseFuncContext {
	var p = new(LowercaseFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LowercaseFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LowercaseFuncContext) Lowercase() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserLowercase, 0)
}

func (s *LowercaseFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LowercaseFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterLowercaseFunc(s)
	}
}

func (s *LowercaseFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitLowercaseFunc(s)
	}
}

type MultiplyContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewMultiplyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiplyContext {
	var p = new(MultiplyContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *MultiplyContext) GetLeft() IExprContext { return s.left }

func (s *MultiplyContext) GetRight() IExprContext { return s.right }

func (s *MultiplyContext) SetLeft(v IExprContext) { s.left = v }

func (s *MultiplyContext) SetRight(v IExprContext) { s.right = v }

func (s *MultiplyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplyContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *MultiplyContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MultiplyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterMultiply(s)
	}
}

func (s *MultiplyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitMultiply(s)
	}
}

type AtanFuncContext struct {
	*ExprContext
}

func NewAtanFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AtanFuncContext {
	var p = new(AtanFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AtanFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtanFuncContext) Atan() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserAtan, 0)
}

func (s *AtanFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AtanFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterAtanFunc(s)
	}
}

func (s *AtanFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitAtanFunc(s)
	}
}

type RoundFuncContext struct {
	*ExprContext
}

func NewRoundFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RoundFuncContext {
	var p = new(RoundFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *RoundFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RoundFuncContext) Round() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserRound, 0)
}

func (s *RoundFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *RoundFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RoundFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterRoundFunc(s)
	}
}

func (s *RoundFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitRoundFunc(s)
	}
}

type SwitchFuncContext struct {
	*ExprContext
}

func NewSwitchFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchFuncContext {
	var p = new(SwitchFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SwitchFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchFuncContext) Switch() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserSwitch, 0)
}

func (s *SwitchFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *SwitchFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SwitchFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterSwitchFunc(s)
	}
}

func (s *SwitchFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitSwitchFunc(s)
	}
}

type ExprFieldContext struct {
	*ExprContext
}

func NewExprFieldContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprFieldContext {
	var p = new(ExprFieldContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExprFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprFieldContext) Field() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserField, 0)
}

func (s *ExprFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterExprField(s)
	}
}

func (s *ExprFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitExprField(s)
	}
}

type GetWordFuncContext struct {
	*ExprContext
}

func NewGetWordFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GetWordFuncContext {
	var p = new(GetWordFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *GetWordFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GetWordFuncContext) GetWord() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserGetWord, 0)
}

func (s *GetWordFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *GetWordFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GetWordFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterGetWordFunc(s)
	}
}

func (s *GetWordFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitGetWordFunc(s)
	}
}

type LengthFuncContext struct {
	*ExprContext
}

func NewLengthFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LengthFuncContext {
	var p = new(LengthFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LengthFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LengthFuncContext) Length() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserLength, 0)
}

func (s *LengthFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LengthFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterLengthFunc(s)
	}
}

func (s *LengthFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitLengthFunc(s)
	}
}

type OrContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OrContext {
	var p = new(OrContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *OrContext) GetLeft() IExprContext { return s.left }

func (s *OrContext) GetRight() IExprContext { return s.right }

func (s *OrContext) SetLeft(v IExprContext) { s.left = v }

func (s *OrContext) SetRight(v IExprContext) { s.right = v }

func (s *OrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *OrContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OrContext) Or() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserOr, 0)
}

func (s *OrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterOr(s)
	}
}

func (s *OrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitOr(s)
	}
}

type ExpFuncContext struct {
	*ExprContext
}

func NewExpFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpFuncContext {
	var p = new(ExpFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ExpFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpFuncContext) Exp() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserExp, 0)
}

func (s *ExpFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ExpFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterExpFunc(s)
	}
}

func (s *ExpFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitExpFunc(s)
	}
}

type CountWordsFuncContext struct {
	*ExprContext
}

func NewCountWordsFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CountWordsFuncContext {
	var p = new(CountWordsFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *CountWordsFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CountWordsFuncContext) CountWords() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserCountWords, 0)
}

func (s *CountWordsFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *CountWordsFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterCountWordsFunc(s)
	}
}

func (s *CountWordsFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitCountWordsFunc(s)
	}
}

type PowFuncContext struct {
	*ExprContext
}

func NewPowFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PowFuncContext {
	var p = new(PowFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *PowFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PowFuncContext) Pow() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserPow, 0)
}

func (s *PowFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *PowFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *PowFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterPowFunc(s)
	}
}

func (s *PowFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitPowFunc(s)
	}
}

type IifFuncContext struct {
	*ExprContext
}

func NewIifFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IifFuncContext {
	var p = new(IifFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *IifFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IifFuncContext) Iif() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserIif, 0)
}

func (s *IifFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *IifFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *IifFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterIifFunc(s)
	}
}

func (s *IifFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitIifFunc(s)
	}
}

type EqualContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EqualContext {
	var p = new(EqualContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *EqualContext) GetLeft() IExprContext { return s.left }

func (s *EqualContext) GetRight() IExprContext { return s.right }

func (s *EqualContext) SetLeft(v IExprContext) { s.left = v }

func (s *EqualContext) SetRight(v IExprContext) { s.right = v }

func (s *EqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EqualContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *EqualContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *EqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterEqual(s)
	}
}

func (s *EqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitEqual(s)
	}
}

type RightFuncContext struct {
	*ExprContext
}

func NewRightFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RightFuncContext {
	var p = new(RightFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *RightFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RightFuncContext) Right() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserRight, 0)
}

func (s *RightFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *RightFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *RightFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterRightFunc(s)
	}
}

func (s *RightFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitRightFunc(s)
	}
}

type ModFuncContext struct {
	*ExprContext
}

func NewModFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ModFuncContext {
	var p = new(ModFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ModFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModFuncContext) Mod() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserMod, 0)
}

func (s *ModFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *ModFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ModFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterModFunc(s)
	}
}

func (s *ModFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitModFunc(s)
	}
}

type SinFuncContext struct {
	*ExprContext
}

func NewSinFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SinFuncContext {
	var p = new(SinFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *SinFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SinFuncContext) Sin() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserSin, 0)
}

func (s *SinFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *SinFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterSinFunc(s)
	}
}

func (s *SinFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitSinFunc(s)
	}
}

type AverageFuncContext struct {
	*ExprContext
}

func NewAverageFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AverageFuncContext {
	var p = new(AverageFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *AverageFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AverageFuncContext) Average() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserAverage, 0)
}

func (s *AverageFuncContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *AverageFuncContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AverageFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterAverageFunc(s)
	}
}

func (s *AverageFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitAverageFunc(s)
	}
}

type GreaterEqualContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewGreaterEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *GreaterEqualContext {
	var p = new(GreaterEqualContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *GreaterEqualContext) GetLeft() IExprContext { return s.left }

func (s *GreaterEqualContext) GetRight() IExprContext { return s.right }

func (s *GreaterEqualContext) SetLeft(v IExprContext) { s.left = v }

func (s *GreaterEqualContext) SetRight(v IExprContext) { s.right = v }

func (s *GreaterEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GreaterEqualContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *GreaterEqualContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *GreaterEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterGreaterEqual(s)
	}
}

func (s *GreaterEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitGreaterEqual(s)
	}
}

type LessEqualContext struct {
	*ExprContext
	left  IExprContext
	right IExprContext
}

func NewLessEqualContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LessEqualContext {
	var p = new(LessEqualContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LessEqualContext) GetLeft() IExprContext { return s.left }

func (s *LessEqualContext) GetRight() IExprContext { return s.right }

func (s *LessEqualContext) SetLeft(v IExprContext) { s.left = v }

func (s *LessEqualContext) SetRight(v IExprContext) { s.right = v }

func (s *LessEqualContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LessEqualContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *LessEqualContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LessEqualContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterLessEqual(s)
	}
}

func (s *LessEqualContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitLessEqual(s)
	}
}

type LogFuncContext struct {
	*ExprContext
}

func NewLogFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogFuncContext {
	var p = new(LogFuncContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *LogFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogFuncContext) Log() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserLog, 0)
}

func (s *LogFuncContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *LogFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterLogFunc(s)
	}
}

func (s *LogFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitLogFunc(s)
	}
}

type NumberLiteralContext struct {
	*ExprContext
}

func NewNumberLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberLiteralContext {
	var p = new(NumberLiteralContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *NumberLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberLiteralContext) Integer() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserInteger, 0)
}

func (s *NumberLiteralContext) Decimal() antlr.TerminalNode {
	return s.GetToken(AlteryxFormulasParserDecimal, 0)
}

func (s *NumberLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.EnterNumberLiteral(s)
	}
}

func (s *NumberLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(AlteryxFormulasListener); ok {
		listenerT.ExitNumberLiteral(s)
	}
}

func (p *AlteryxFormulasParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *AlteryxFormulasParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, AlteryxFormulasParserRULE_expr, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewParenthesisContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(3)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(4)
			p.expr(0)
		}
		{
			p.SetState(5)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 2:
		localctx = NewExprIfContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(7)
			p.Match(AlteryxFormulasParserIf)
		}
		{
			p.SetState(8)

			var _x = p.expr(0)

			localctx.(*ExprIfContext).ifStmnt = _x
		}
		{
			p.SetState(9)
			p.Match(AlteryxFormulasParserThen)
		}
		{
			p.SetState(10)

			var _x = p.expr(0)

			localctx.(*ExprIfContext).thenStmnt = _x
		}
		p.SetState(18)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AlteryxFormulasParserElseif {
			{
				p.SetState(11)
				p.Match(AlteryxFormulasParserElseif)
			}
			{
				p.SetState(12)

				var _x = p.expr(0)

				localctx.(*ExprIfContext).ifStmnt = _x
			}
			{
				p.SetState(13)
				p.Match(AlteryxFormulasParserThen)
			}
			{
				p.SetState(14)

				var _x = p.expr(0)

				localctx.(*ExprIfContext).thenStmnt = _x
			}

			p.SetState(20)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(21)
			p.Match(AlteryxFormulasParserElse)
		}
		{
			p.SetState(22)

			var _x = p.expr(0)

			localctx.(*ExprIfContext).elseStmnt = _x
		}
		{
			p.SetState(23)
			p.Match(AlteryxFormulasParserEndif)
		}

	case 3:
		localctx = NewIifFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(25)
			p.Match(AlteryxFormulasParserIif)
		}
		{
			p.SetState(26)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(27)
			p.expr(0)
		}
		{
			p.SetState(28)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(29)
			p.expr(0)
		}
		{
			p.SetState(30)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(31)
			p.expr(0)
		}
		{
			p.SetState(32)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 4:
		localctx = NewAbsFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(34)
			p.Match(AlteryxFormulasParserAbs)
		}
		{
			p.SetState(35)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(36)
			p.expr(0)
		}
		{
			p.SetState(37)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 5:
		localctx = NewAcosFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(39)
			p.Match(AlteryxFormulasParserAcos)
		}
		{
			p.SetState(40)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(41)
			p.expr(0)
		}
		{
			p.SetState(42)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 6:
		localctx = NewAsinFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(44)
			p.Match(AlteryxFormulasParserAsin)
		}
		{
			p.SetState(45)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(46)
			p.expr(0)
		}
		{
			p.SetState(47)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 7:
		localctx = NewAtanFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(49)
			p.Match(AlteryxFormulasParserAtan)
		}
		{
			p.SetState(50)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(51)
			p.expr(0)
		}
		{
			p.SetState(52)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 8:
		localctx = NewAtan2FuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(54)
			p.Match(AlteryxFormulasParserAtan2)
		}
		{
			p.SetState(55)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(56)
			p.expr(0)
		}
		{
			p.SetState(57)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(58)
			p.expr(0)
		}
		{
			p.SetState(59)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 9:
		localctx = NewAverageFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(61)
			p.Match(AlteryxFormulasParserAverage)
		}
		{
			p.SetState(62)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(63)
			p.expr(0)
		}
		p.SetState(68)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(64)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(65)
				p.expr(0)
			}

			p.SetState(70)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(71)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 10:
		localctx = NewCeilFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(73)
			p.Match(AlteryxFormulasParserCeil)
		}
		{
			p.SetState(74)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(75)
			p.expr(0)
		}
		{
			p.SetState(76)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 11:
		localctx = NewCharFromIntFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(78)
			p.Match(AlteryxFormulasParserCharFromInt)
		}
		{
			p.SetState(79)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(80)
			p.expr(0)
		}
		{
			p.SetState(81)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 12:
		localctx = NewCharToIntFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(83)
			p.Match(AlteryxFormulasParserCharToInt)
		}
		{
			p.SetState(84)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(85)
			p.expr(0)
		}
		{
			p.SetState(86)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 13:
		localctx = NewContainsFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(88)
			p.Match(AlteryxFormulasParserContains)
		}
		{
			p.SetState(89)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(90)
			p.expr(0)
		}
		{
			p.SetState(91)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(92)
			p.expr(0)
		}
		p.SetState(95)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(93)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(94)
				p.expr(0)
			}

		}
		{
			p.SetState(97)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 14:
		localctx = NewCosFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(99)
			p.Match(AlteryxFormulasParserCos)
		}
		{
			p.SetState(100)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(101)
			p.expr(0)
		}
		{
			p.SetState(102)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 15:
		localctx = NewCoshFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(104)
			p.Match(AlteryxFormulasParserCosh)
		}
		{
			p.SetState(105)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(106)
			p.expr(0)
		}
		{
			p.SetState(107)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 16:
		localctx = NewCountWordsFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(109)
			p.Match(AlteryxFormulasParserCountWords)
		}
		{
			p.SetState(110)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(111)
			p.expr(0)
		}
		{
			p.SetState(112)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 17:
		localctx = NewDistanceFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(114)
			p.Match(AlteryxFormulasParserDistance)
		}
		{
			p.SetState(115)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(116)
			p.expr(0)
		}
		{
			p.SetState(117)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(118)
			p.expr(0)
		}
		{
			p.SetState(119)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(120)
			p.expr(0)
		}
		{
			p.SetState(121)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(122)
			p.expr(0)
		}
		{
			p.SetState(123)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 18:
		localctx = NewEndsWithFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(125)
			p.Match(AlteryxFormulasParserEndsWith)
		}
		{
			p.SetState(126)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(127)
			p.expr(0)
		}
		{
			p.SetState(128)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(129)
			p.expr(0)
		}
		p.SetState(132)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(130)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(131)
				p.expr(0)
			}

		}
		{
			p.SetState(134)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 19:
		localctx = NewExpFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(136)
			p.Match(AlteryxFormulasParserExp)
		}
		{
			p.SetState(137)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(138)
			p.expr(0)
		}
		{
			p.SetState(139)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 20:
		localctx = NewFindStringFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(141)
			p.Match(AlteryxFormulasParserFindString)
		}
		{
			p.SetState(142)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(143)
			p.expr(0)
		}
		{
			p.SetState(144)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(145)
			p.expr(0)
		}
		{
			p.SetState(146)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 21:
		localctx = NewFloorFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(148)
			p.Match(AlteryxFormulasParserFloor)
		}
		{
			p.SetState(149)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(150)
			p.expr(0)
		}
		{
			p.SetState(151)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 22:
		localctx = NewGetWordFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(153)
			p.Match(AlteryxFormulasParserGetWord)
		}
		{
			p.SetState(154)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(155)
			p.expr(0)
		}
		{
			p.SetState(156)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(157)
			p.expr(0)
		}
		{
			p.SetState(158)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 23:
		localctx = NewHexToNumberFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(160)
			p.Match(AlteryxFormulasParserHexToNumber)
		}
		{
			p.SetState(161)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(162)
			p.expr(0)
		}
		{
			p.SetState(163)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 24:
		localctx = NewLeftFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(165)
			p.Match(AlteryxFormulasParserLeft)
		}
		{
			p.SetState(166)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(167)
			p.expr(0)
		}
		{
			p.SetState(168)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(169)
			p.expr(0)
		}
		{
			p.SetState(170)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 25:
		localctx = NewLengthFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(172)
			p.Match(AlteryxFormulasParserLength)
		}
		{
			p.SetState(173)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(174)
			p.expr(0)
		}
		{
			p.SetState(175)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 26:
		localctx = NewLogFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(177)
			p.Match(AlteryxFormulasParserLog)
		}
		{
			p.SetState(178)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(179)
			p.expr(0)
		}
		{
			p.SetState(180)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 27:
		localctx = NewLog10FuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(182)
			p.Match(AlteryxFormulasParserLog10)
		}
		{
			p.SetState(183)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(184)
			p.expr(0)
		}
		{
			p.SetState(185)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 28:
		localctx = NewLowercaseFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(187)
			p.Match(AlteryxFormulasParserLowercase)
		}
		{
			p.SetState(188)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(189)
			p.expr(0)
		}
		{
			p.SetState(190)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 29:
		localctx = NewMaxFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(192)
			p.Match(AlteryxFormulasParserMax)
		}
		{
			p.SetState(193)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(194)
			p.expr(0)
		}
		p.SetState(197)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(195)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(196)
				p.expr(0)
			}

			p.SetState(199)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(201)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 30:
		localctx = NewMedianFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(203)
			p.Match(AlteryxFormulasParserMedian)
		}
		{
			p.SetState(204)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(205)
			p.expr(0)
		}
		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(206)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(207)
				p.expr(0)
			}

			p.SetState(212)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(213)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 31:
		localctx = NewMinFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(215)
			p.Match(AlteryxFormulasParserMin)
		}
		{
			p.SetState(216)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(217)
			p.expr(0)
		}
		p.SetState(220)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(218)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(219)
				p.expr(0)
			}

			p.SetState(222)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(224)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 32:
		localctx = NewModFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(226)
			p.Match(AlteryxFormulasParserMod)
		}
		{
			p.SetState(227)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(228)
			p.expr(0)
		}
		p.SetState(233)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(229)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(230)
				p.expr(0)
			}

			p.SetState(235)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(236)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 33:
		localctx = NewNullFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(238)
			p.Match(AlteryxFormulasParserNull)
		}
		{
			p.SetState(239)
			p.Match(AlteryxFormulasParserT__15)
		}

	case 34:
		localctx = NewPadLeftFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(240)
			p.Match(AlteryxFormulasParserPadLeft)
		}
		{
			p.SetState(241)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(242)
			p.expr(0)
		}
		{
			p.SetState(243)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(244)
			p.expr(0)
		}
		{
			p.SetState(245)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(246)
			p.expr(0)
		}
		{
			p.SetState(247)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 35:
		localctx = NewPadRightFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(249)
			p.Match(AlteryxFormulasParserPadRight)
		}
		{
			p.SetState(250)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(251)
			p.expr(0)
		}
		{
			p.SetState(252)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(253)
			p.expr(0)
		}
		{
			p.SetState(254)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(255)
			p.expr(0)
		}
		{
			p.SetState(256)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 36:
		localctx = NewPiFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(258)
			p.Match(AlteryxFormulasParserPi)
		}
		{
			p.SetState(259)
			p.Match(AlteryxFormulasParserT__15)
		}

	case 37:
		localctx = NewPowFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(260)
			p.Match(AlteryxFormulasParserPow)
		}
		{
			p.SetState(261)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(262)
			p.expr(0)
		}
		{
			p.SetState(263)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(264)
			p.expr(0)
		}
		{
			p.SetState(265)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 38:
		localctx = NewRandFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(267)
			p.Match(AlteryxFormulasParserRand)
		}
		{
			p.SetState(268)
			p.Match(AlteryxFormulasParserT__15)
		}

	case 39:
		localctx = NewRandIntFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(269)
			p.Match(AlteryxFormulasParserRandInt)
		}
		{
			p.SetState(270)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(271)
			p.expr(0)
		}
		{
			p.SetState(272)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 40:
		localctx = NewRegexCountMatchesFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(274)
			p.Match(AlteryxFormulasParserRegex_CountMatches)
		}
		{
			p.SetState(275)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(276)
			p.expr(0)
		}
		{
			p.SetState(277)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(278)
			p.expr(0)
		}
		p.SetState(281)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(279)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(280)
				p.expr(0)
			}

		}
		{
			p.SetState(283)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 41:
		localctx = NewRegexMatchFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(285)
			p.Match(AlteryxFormulasParserRegex_Match)
		}
		{
			p.SetState(286)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(287)
			p.expr(0)
		}
		{
			p.SetState(288)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(289)
			p.expr(0)
		}
		p.SetState(292)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(290)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(291)
				p.expr(0)
			}

		}
		{
			p.SetState(294)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 42:
		localctx = NewRegexReplaceFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(296)
			p.Match(AlteryxFormulasParserRegex_Replace)
		}
		{
			p.SetState(297)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(298)
			p.expr(0)
		}
		{
			p.SetState(299)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(300)
			p.expr(0)
		}
		{
			p.SetState(301)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(302)
			p.expr(0)
		}
		p.SetState(305)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(303)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(304)
				p.expr(0)
			}

		}
		{
			p.SetState(307)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 43:
		localctx = NewReplaceFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(309)
			p.Match(AlteryxFormulasParserReplace)
		}
		{
			p.SetState(310)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(311)
			p.expr(0)
		}
		{
			p.SetState(312)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(313)
			p.expr(0)
		}
		{
			p.SetState(314)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(315)
			p.expr(0)
		}
		{
			p.SetState(316)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 44:
		localctx = NewRightFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(318)
			p.Match(AlteryxFormulasParserRight)
		}
		{
			p.SetState(319)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(320)
			p.expr(0)
		}
		{
			p.SetState(321)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(322)
			p.expr(0)
		}
		{
			p.SetState(323)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 45:
		localctx = NewRoundFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(325)
			p.Match(AlteryxFormulasParserRound)
		}
		{
			p.SetState(326)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(327)
			p.expr(0)
		}
		{
			p.SetState(328)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(329)
			p.expr(0)
		}
		{
			p.SetState(330)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 46:
		localctx = NewSinFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(332)
			p.Match(AlteryxFormulasParserSin)
		}
		{
			p.SetState(333)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(334)
			p.expr(0)
		}
		{
			p.SetState(335)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 47:
		localctx = NewSinhFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(337)
			p.Match(AlteryxFormulasParserSinh)
		}
		{
			p.SetState(338)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(339)
			p.expr(0)
		}
		{
			p.SetState(340)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 48:
		localctx = NewSqrtFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(342)
			p.Match(AlteryxFormulasParserSqrt)
		}
		{
			p.SetState(343)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(344)
			p.expr(0)
		}
		{
			p.SetState(345)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 49:
		localctx = NewSwitchFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(347)
			p.Match(AlteryxFormulasParserSwitch)
		}
		{
			p.SetState(348)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(349)
			p.expr(0)
		}
		{
			p.SetState(350)
			p.Match(AlteryxFormulasParserT__12)
		}
		{
			p.SetState(351)
			p.expr(0)
		}
		p.SetState(357)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == AlteryxFormulasParserT__12 {
			{
				p.SetState(352)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(353)
				p.expr(0)
			}
			{
				p.SetState(354)
				p.Match(AlteryxFormulasParserT__12)
			}
			{
				p.SetState(355)
				p.expr(0)
			}

			p.SetState(359)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(361)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 50:
		localctx = NewTanFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(363)
			p.Match(AlteryxFormulasParserTan)
		}
		{
			p.SetState(364)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(365)
			p.expr(0)
		}
		{
			p.SetState(366)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 51:
		localctx = NewTanhFuncContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(368)
			p.Match(AlteryxFormulasParserTanh)
		}
		{
			p.SetState(369)
			p.Match(AlteryxFormulasParserT__0)
		}
		{
			p.SetState(370)
			p.expr(0)
		}
		{
			p.SetState(371)
			p.Match(AlteryxFormulasParserT__1)
		}

	case 52:
		localctx = NewNumberLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(373)
			p.Match(AlteryxFormulasParserInteger)
		}

	case 53:
		localctx = NewNumberLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(374)
			p.Match(AlteryxFormulasParserT__5)
		}
		{
			p.SetState(375)
			p.Match(AlteryxFormulasParserInteger)
		}

	case 54:
		localctx = NewNumberLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(376)
			p.Match(AlteryxFormulasParserDecimal)
		}

	case 55:
		localctx = NewNumberLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(377)
			p.Match(AlteryxFormulasParserT__5)
		}
		{
			p.SetState(378)
			p.Match(AlteryxFormulasParserDecimal)
		}

	case 56:
		localctx = NewStringLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(379)
			p.Match(AlteryxFormulasParserSingleQuoteString)
		}

	case 57:
		localctx = NewStringLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(380)
			p.Match(AlteryxFormulasParserDoubleQuoteString)
		}

	case 58:
		localctx = NewDatetimeLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(381)
			p.Match(AlteryxFormulasParserDatetime)
		}

	case 59:
		localctx = NewDateLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(382)
			p.Match(AlteryxFormulasParserDate)
		}

	case 60:
		localctx = NewBoolLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(383)
			p.Match(AlteryxFormulasParserBool)
		}

	case 61:
		localctx = NewExprFieldContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(384)
			p.Match(AlteryxFormulasParserField)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(454)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(452)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMultiplyContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*MultiplyContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(387)

				if !(p.Precpred(p.GetParserRuleContext(), 74)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 74)", ""))
				}
				{
					p.SetState(388)
					p.Match(AlteryxFormulasParserT__2)
				}
				{
					p.SetState(389)

					var _x = p.expr(75)

					localctx.(*MultiplyContext).right = _x
				}

			case 2:
				localctx = NewDivideContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*DivideContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(390)

				if !(p.Precpred(p.GetParserRuleContext(), 73)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 73)", ""))
				}
				{
					p.SetState(391)
					p.Match(AlteryxFormulasParserT__3)
				}
				{
					p.SetState(392)

					var _x = p.expr(74)

					localctx.(*DivideContext).right = _x
				}

			case 3:
				localctx = NewAddContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AddContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(393)

				if !(p.Precpred(p.GetParserRuleContext(), 72)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 72)", ""))
				}
				{
					p.SetState(394)
					p.Match(AlteryxFormulasParserT__4)
				}
				{
					p.SetState(395)

					var _x = p.expr(73)

					localctx.(*AddContext).right = _x
				}

			case 4:
				localctx = NewSubtractContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*SubtractContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(396)

				if !(p.Precpred(p.GetParserRuleContext(), 71)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 71)", ""))
				}
				{
					p.SetState(397)
					p.Match(AlteryxFormulasParserT__5)
				}
				{
					p.SetState(398)

					var _x = p.expr(72)

					localctx.(*SubtractContext).right = _x
				}

			case 5:
				localctx = NewEqualContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*EqualContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(399)

				if !(p.Precpred(p.GetParserRuleContext(), 70)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 70)", ""))
				}
				{
					p.SetState(400)
					p.Match(AlteryxFormulasParserT__6)
				}
				{
					p.SetState(401)

					var _x = p.expr(71)

					localctx.(*EqualContext).right = _x
				}

			case 6:
				localctx = NewGreaterThanContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*GreaterThanContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(402)

				if !(p.Precpred(p.GetParserRuleContext(), 69)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 69)", ""))
				}
				{
					p.SetState(403)
					p.Match(AlteryxFormulasParserT__7)
				}
				{
					p.SetState(404)

					var _x = p.expr(70)

					localctx.(*GreaterThanContext).right = _x
				}

			case 7:
				localctx = NewGreaterEqualContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*GreaterEqualContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(405)

				if !(p.Precpred(p.GetParserRuleContext(), 68)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 68)", ""))
				}
				{
					p.SetState(406)
					p.Match(AlteryxFormulasParserT__8)
				}
				{
					p.SetState(407)

					var _x = p.expr(69)

					localctx.(*GreaterEqualContext).right = _x
				}

			case 8:
				localctx = NewLessThanContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LessThanContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(408)

				if !(p.Precpred(p.GetParserRuleContext(), 67)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 67)", ""))
				}
				{
					p.SetState(409)
					p.Match(AlteryxFormulasParserT__9)
				}
				{
					p.SetState(410)

					var _x = p.expr(68)

					localctx.(*LessThanContext).right = _x
				}

			case 9:
				localctx = NewLessEqualContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*LessEqualContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(411)

				if !(p.Precpred(p.GetParserRuleContext(), 66)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 66)", ""))
				}
				{
					p.SetState(412)
					p.Match(AlteryxFormulasParserT__10)
				}
				{
					p.SetState(413)

					var _x = p.expr(67)

					localctx.(*LessEqualContext).right = _x
				}

			case 10:
				localctx = NewNotEqualContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*NotEqualContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(414)

				if !(p.Precpred(p.GetParserRuleContext(), 65)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 65)", ""))
				}
				{
					p.SetState(415)
					p.Match(AlteryxFormulasParserT__11)
				}
				{
					p.SetState(416)

					var _x = p.expr(66)

					localctx.(*NotEqualContext).right = _x
				}

			case 11:
				localctx = NewAndContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AndContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(417)

				if !(p.Precpred(p.GetParserRuleContext(), 62)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 62)", ""))
				}
				{
					p.SetState(418)
					_la = p.GetTokenStream().LA(1)

					if !(_la == AlteryxFormulasParserT__13 || _la == AlteryxFormulasParserAnd) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(419)

					var _x = p.expr(63)

					localctx.(*AndContext).right = _x
				}

			case 12:
				localctx = NewOrContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OrContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(420)

				if !(p.Precpred(p.GetParserRuleContext(), 61)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 61)", ""))
				}
				{
					p.SetState(421)
					_la = p.GetTokenStream().LA(1)

					if !(_la == AlteryxFormulasParserT__14 || _la == AlteryxFormulasParserOr) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(422)

					var _x = p.expr(62)

					localctx.(*OrContext).right = _x
				}

			case 13:
				localctx = NewInContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(423)

				if !(p.Precpred(p.GetParserRuleContext(), 64)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 64)", ""))
				}
				{
					p.SetState(424)
					p.Match(AlteryxFormulasParserIn)
				}
				{
					p.SetState(425)
					p.Match(AlteryxFormulasParserT__0)
				}
				p.SetState(434)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AlteryxFormulasParserT__0)|(1<<AlteryxFormulasParserT__5)|(1<<AlteryxFormulasParserAbs)|(1<<AlteryxFormulasParserAcos)|(1<<AlteryxFormulasParserAsin)|(1<<AlteryxFormulasParserAtan)|(1<<AlteryxFormulasParserAtan2)|(1<<AlteryxFormulasParserAverage)|(1<<AlteryxFormulasParserCeil)|(1<<AlteryxFormulasParserCharFromInt)|(1<<AlteryxFormulasParserCharToInt)|(1<<AlteryxFormulasParserContains)|(1<<AlteryxFormulasParserCos)|(1<<AlteryxFormulasParserCosh)|(1<<AlteryxFormulasParserCountWords)|(1<<AlteryxFormulasParserDistance)|(1<<AlteryxFormulasParserEndsWith))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AlteryxFormulasParserFindString-32))|(1<<(AlteryxFormulasParserExp-32))|(1<<(AlteryxFormulasParserFloor-32))|(1<<(AlteryxFormulasParserGetWord-32))|(1<<(AlteryxFormulasParserHexToNumber-32))|(1<<(AlteryxFormulasParserLeft-32))|(1<<(AlteryxFormulasParserLength-32))|(1<<(AlteryxFormulasParserLog-32))|(1<<(AlteryxFormulasParserLog10-32))|(1<<(AlteryxFormulasParserLowercase-32))|(1<<(AlteryxFormulasParserMax-32))|(1<<(AlteryxFormulasParserMedian-32))|(1<<(AlteryxFormulasParserMin-32))|(1<<(AlteryxFormulasParserMod-32))|(1<<(AlteryxFormulasParserNull-32))|(1<<(AlteryxFormulasParserPadLeft-32))|(1<<(AlteryxFormulasParserPadRight-32))|(1<<(AlteryxFormulasParserPi-32))|(1<<(AlteryxFormulasParserPow-32))|(1<<(AlteryxFormulasParserRand-32))|(1<<(AlteryxFormulasParserRandInt-32))|(1<<(AlteryxFormulasParserRegex_CountMatches-32))|(1<<(AlteryxFormulasParserRegex_Match-32))|(1<<(AlteryxFormulasParserRegex_Replace-32))|(1<<(AlteryxFormulasParserReplace-32))|(1<<(AlteryxFormulasParserRight-32))|(1<<(AlteryxFormulasParserRound-32))|(1<<(AlteryxFormulasParserSin-32))|(1<<(AlteryxFormulasParserSinh-32))|(1<<(AlteryxFormulasParserSqrt-32))|(1<<(AlteryxFormulasParserSwitch-32))|(1<<(AlteryxFormulasParserTan-32)))) != 0) || (((_la-64)&-(0x1f+1)) == 0 && ((1<<uint((_la-64)))&((1<<(AlteryxFormulasParserTanh-64))|(1<<(AlteryxFormulasParserIif-64))|(1<<(AlteryxFormulasParserIf-64))|(1<<(AlteryxFormulasParserBool-64))|(1<<(AlteryxFormulasParserInteger-64))|(1<<(AlteryxFormulasParserDecimal-64))|(1<<(AlteryxFormulasParserDate-64))|(1<<(AlteryxFormulasParserDatetime-64))|(1<<(AlteryxFormulasParserField-64))|(1<<(AlteryxFormulasParserSingleQuoteString-64))|(1<<(AlteryxFormulasParserDoubleQuoteString-64)))) != 0) {
					{
						p.SetState(426)
						p.expr(0)
					}
					p.SetState(431)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == AlteryxFormulasParserT__12 {
						{
							p.SetState(427)
							p.Match(AlteryxFormulasParserT__12)
						}
						{
							p.SetState(428)
							p.expr(0)
						}

						p.SetState(433)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(436)
					p.Match(AlteryxFormulasParserT__1)
				}

			case 14:
				localctx = NewNotInContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, AlteryxFormulasParserRULE_expr)
				p.SetState(437)

				if !(p.Precpred(p.GetParserRuleContext(), 63)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 63)", ""))
				}
				{
					p.SetState(438)
					p.Match(AlteryxFormulasParserNot)
				}
				{
					p.SetState(439)
					p.Match(AlteryxFormulasParserIn)
				}
				{
					p.SetState(440)
					p.Match(AlteryxFormulasParserT__0)
				}
				p.SetState(449)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AlteryxFormulasParserT__0)|(1<<AlteryxFormulasParserT__5)|(1<<AlteryxFormulasParserAbs)|(1<<AlteryxFormulasParserAcos)|(1<<AlteryxFormulasParserAsin)|(1<<AlteryxFormulasParserAtan)|(1<<AlteryxFormulasParserAtan2)|(1<<AlteryxFormulasParserAverage)|(1<<AlteryxFormulasParserCeil)|(1<<AlteryxFormulasParserCharFromInt)|(1<<AlteryxFormulasParserCharToInt)|(1<<AlteryxFormulasParserContains)|(1<<AlteryxFormulasParserCos)|(1<<AlteryxFormulasParserCosh)|(1<<AlteryxFormulasParserCountWords)|(1<<AlteryxFormulasParserDistance)|(1<<AlteryxFormulasParserEndsWith))) != 0) || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(AlteryxFormulasParserFindString-32))|(1<<(AlteryxFormulasParserExp-32))|(1<<(AlteryxFormulasParserFloor-32))|(1<<(AlteryxFormulasParserGetWord-32))|(1<<(AlteryxFormulasParserHexToNumber-32))|(1<<(AlteryxFormulasParserLeft-32))|(1<<(AlteryxFormulasParserLength-32))|(1<<(AlteryxFormulasParserLog-32))|(1<<(AlteryxFormulasParserLog10-32))|(1<<(AlteryxFormulasParserLowercase-32))|(1<<(AlteryxFormulasParserMax-32))|(1<<(AlteryxFormulasParserMedian-32))|(1<<(AlteryxFormulasParserMin-32))|(1<<(AlteryxFormulasParserMod-32))|(1<<(AlteryxFormulasParserNull-32))|(1<<(AlteryxFormulasParserPadLeft-32))|(1<<(AlteryxFormulasParserPadRight-32))|(1<<(AlteryxFormulasParserPi-32))|(1<<(AlteryxFormulasParserPow-32))|(1<<(AlteryxFormulasParserRand-32))|(1<<(AlteryxFormulasParserRandInt-32))|(1<<(AlteryxFormulasParserRegex_CountMatches-32))|(1<<(AlteryxFormulasParserRegex_Match-32))|(1<<(AlteryxFormulasParserRegex_Replace-32))|(1<<(AlteryxFormulasParserReplace-32))|(1<<(AlteryxFormulasParserRight-32))|(1<<(AlteryxFormulasParserRound-32))|(1<<(AlteryxFormulasParserSin-32))|(1<<(AlteryxFormulasParserSinh-32))|(1<<(AlteryxFormulasParserSqrt-32))|(1<<(AlteryxFormulasParserSwitch-32))|(1<<(AlteryxFormulasParserTan-32)))) != 0) || (((_la-64)&-(0x1f+1)) == 0 && ((1<<uint((_la-64)))&((1<<(AlteryxFormulasParserTanh-64))|(1<<(AlteryxFormulasParserIif-64))|(1<<(AlteryxFormulasParserIf-64))|(1<<(AlteryxFormulasParserBool-64))|(1<<(AlteryxFormulasParserInteger-64))|(1<<(AlteryxFormulasParserDecimal-64))|(1<<(AlteryxFormulasParserDate-64))|(1<<(AlteryxFormulasParserDatetime-64))|(1<<(AlteryxFormulasParserField-64))|(1<<(AlteryxFormulasParserSingleQuoteString-64))|(1<<(AlteryxFormulasParserDoubleQuoteString-64)))) != 0) {
					{
						p.SetState(441)
						p.expr(0)
					}
					p.SetState(446)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == AlteryxFormulasParserT__12 {
						{
							p.SetState(442)
							p.Match(AlteryxFormulasParserT__12)
						}
						{
							p.SetState(443)
							p.expr(0)
						}

						p.SetState(448)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(451)
					p.Match(AlteryxFormulasParserT__1)
				}

			}

		}
		p.SetState(456)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
	}

	return localctx
}

func (p *AlteryxFormulasParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *AlteryxFormulasParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 74)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 73)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 72)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 71)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 70)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 69)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 68)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 67)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 66)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 65)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 62)

	case 11:
		return p.Precpred(p.GetParserRuleContext(), 61)

	case 12:
		return p.Precpred(p.GetParserRuleContext(), 64)

	case 13:
		return p.Precpred(p.GetParserRuleContext(), 63)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
