package main

import keyboard "github.com/percyjw-2/tinygo-keyboard"

const noCols = 14
const noRows = 11

func posToCoordY(col uint8) uint8 {
	return 224 / (noCols - 1) * col
}

func posToCoordX(row uint8) uint8 {
	return 64 / (noRows - 1) * row
}

var rgbLedMap = []keyboard.LedMatrixPosition{
	{
		PhysicalX:   posToCoordX(9),
		PhysicalY:   posToCoordY(3),
		KbIndex:     0,
		MatrixIndex: 19,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT | keyboard.LED_FLAG_MODIFIER,
	}, {
		PhysicalX:   posToCoordX(9),
		PhysicalY:   posToCoordY(4),
		KbIndex:     0,
		MatrixIndex: 20,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT | keyboard.LED_FLAG_MODIFIER,
	}, {
		PhysicalX:   posToCoordX(10),
		PhysicalY:   posToCoordY(5),
		KbIndex:     0,
		MatrixIndex: 21,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT | keyboard.LED_FLAG_MODIFIER,
	}, {
		PhysicalX:   posToCoordX(9),
		PhysicalY:   posToCoordY(6),
		KbIndex:     0,
		MatrixIndex: 22,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT | keyboard.LED_FLAG_MODIFIER,
	}, {
		PhysicalX:   posToCoordX(8),
		PhysicalY:   posToCoordY(5),
		KbIndex:     0,
		MatrixIndex: 17,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(7),
		PhysicalY:   posToCoordY(4),
		KbIndex:     0,
		MatrixIndex: 16,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(6),
		PhysicalY:   posToCoordY(3),
		KbIndex:     0,
		MatrixIndex: 15,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(7),
		PhysicalY:   posToCoordY(2),
		KbIndex:     0,
		MatrixIndex: 14,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(8),
		PhysicalY:   posToCoordY(1),
		KbIndex:     0,
		MatrixIndex: 13,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(8),
		PhysicalY:   posToCoordY(0),
		KbIndex:     0,
		MatrixIndex: 12,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT | keyboard.LED_FLAG_MODIFIER,
	}, {
		PhysicalX:   posToCoordX(5),
		PhysicalY:   posToCoordY(0),
		KbIndex:     0,
		MatrixIndex: 6,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT | keyboard.LED_FLAG_MODIFIER,
	}, {
		PhysicalX:   posToCoordX(5),
		PhysicalY:   posToCoordY(1),
		KbIndex:     0,
		MatrixIndex: 7,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(4),
		PhysicalY:   posToCoordY(2),
		KbIndex:     0,
		MatrixIndex: 8,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(3),
		PhysicalY:   posToCoordY(3),
		KbIndex:     0,
		MatrixIndex: 9,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(4),
		PhysicalY:   posToCoordY(4),
		KbIndex:     0,
		MatrixIndex: 10,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(5),
		PhysicalY:   posToCoordY(5),
		KbIndex:     0,
		MatrixIndex: 11,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(2),
		PhysicalY:   posToCoordY(5),
		KbIndex:     0,
		MatrixIndex: 5,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(1),
		PhysicalY:   posToCoordY(4),
		KbIndex:     0,
		MatrixIndex: 4,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(0),
		PhysicalY:   posToCoordY(3),
		KbIndex:     0,
		MatrixIndex: 3,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(1),
		PhysicalY:   posToCoordY(2),
		KbIndex:     0,
		MatrixIndex: 2,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(2),
		PhysicalY:   posToCoordY(1),
		KbIndex:     0,
		MatrixIndex: 1,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(2),
		PhysicalY:   posToCoordY(9),
		KbIndex:     0,
		MatrixIndex: 0,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT | keyboard.LED_FLAG_MODIFIER,
	}, {
		PhysicalX:   posToCoordX(2),
		PhysicalY:   posToCoordY(8),
		KbIndex:     1,
		MatrixIndex: 0,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(1),
		PhysicalY:   posToCoordY(9),
		KbIndex:     1,
		MatrixIndex: 1,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(0),
		PhysicalY:   posToCoordY(10),
		KbIndex:     1,
		MatrixIndex: 2,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(1),
		PhysicalY:   posToCoordY(11),
		KbIndex:     1,
		MatrixIndex: 3,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(2),
		PhysicalY:   posToCoordY(12),
		KbIndex:     1,
		MatrixIndex: 4,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(2),
		PhysicalY:   posToCoordY(13),
		KbIndex:     1,
		MatrixIndex: 5,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT | keyboard.LED_FLAG_MODIFIER,
	}, {
		PhysicalX:   posToCoordX(5),
		PhysicalY:   posToCoordY(13),
		KbIndex:     1,
		MatrixIndex: 11,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT | keyboard.LED_FLAG_MODIFIER,
	}, {
		PhysicalX:   posToCoordX(5),
		PhysicalY:   posToCoordY(12),
		KbIndex:     1,
		MatrixIndex: 10,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(4),
		PhysicalY:   posToCoordY(11),
		KbIndex:     1,
		MatrixIndex: 9,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(3),
		PhysicalY:   posToCoordY(10),
		KbIndex:     1,
		MatrixIndex: 8,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(4),
		PhysicalY:   posToCoordY(9),
		KbIndex:     1,
		MatrixIndex: 7,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(5),
		PhysicalY:   posToCoordY(8),
		KbIndex:     6,
		MatrixIndex: keyboard.LED_FLAG_KEYLIGHT,
		LedFlags:    0,
	}, {
		PhysicalX:   posToCoordX(8),
		PhysicalY:   posToCoordY(8),
		KbIndex:     1,
		MatrixIndex: 12,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(7),
		PhysicalY:   posToCoordY(9),
		KbIndex:     1,
		MatrixIndex: 13,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(6),
		PhysicalY:   posToCoordY(10),
		KbIndex:     1,
		MatrixIndex: 14,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(7),
		PhysicalY:   posToCoordY(11),
		KbIndex:     1,
		MatrixIndex: 15,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(8),
		PhysicalY:   posToCoordY(12),
		KbIndex:     1,
		MatrixIndex: 16,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT,
	}, {
		PhysicalX:   posToCoordX(8),
		PhysicalY:   posToCoordY(13),
		KbIndex:     1,
		MatrixIndex: 17,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT | keyboard.LED_FLAG_MODIFIER,
	}, {
		PhysicalX:   posToCoordX(9),
		PhysicalY:   posToCoordY(10),
		KbIndex:     1,
		MatrixIndex: 21,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT | keyboard.LED_FLAG_MODIFIER,
	}, {
		PhysicalX:   posToCoordX(9),
		PhysicalY:   posToCoordY(9),
		KbIndex:     1,
		MatrixIndex: 20,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT | keyboard.LED_FLAG_MODIFIER,
	}, {
		PhysicalX:   posToCoordX(10),
		PhysicalY:   posToCoordY(8),
		KbIndex:     1,
		MatrixIndex: 19,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT | keyboard.LED_FLAG_MODIFIER,
	}, {
		PhysicalX:   posToCoordX(9),
		PhysicalY:   posToCoordY(7),
		KbIndex:     1,
		MatrixIndex: 18,
		LedFlags:    keyboard.LED_FLAG_KEYLIGHT | keyboard.LED_FLAG_MODIFIER,
	},
}
