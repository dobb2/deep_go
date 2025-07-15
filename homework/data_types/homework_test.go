package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v homework_test.go

func ToLittleEndian(number uint32) uint32 {
	return number<<(8*3) | ((number>>8)&0xFF)<<(8*2) | ((number>>(8*2))&0xFF)<<8 | number>>(8*3)
}

/*
 Задание: Релизовать функцию по конвертации числа из прямого порядка следования байт (Big Endian) в обратный порядок следования байт (Little Endian). С uint32
 Решение:
 Нужно получить каждый байт числа в нужном порядке следующим образом:
 1) сдвинуть число на i-1 байт вправо, где i позиция нужного байта, начинается с 1
 2) из полученного числа в п.1 получить 1 байт ( & 0xFF и старшие байты(биты) занулятся)
 3) число из п.3 сдвинуть на b-i байт влево. b - из скольки байт представлено исходное число(uint32 - 4, uint16 - 2)

 Байтовый сдвиг эквивалентен сдвигу на 8 бит
 4) полученные числа сложить побитовой операцией ИЛИ

 Пример: число 0x11121314
   11 12 13 14
  1 байт - 04, 1 позиция, сдвиг вправо на 1-1=0 байт => 0x11121314,
	оперцию & 0xFF можно и не делать, так как сдвигаем старший байт,
	сдвиг числа 0x11121314(или 0x00000014) на 4-1=3 байта(24 бита) влево  => 0x14000000
  2 байт - 03, 1 позиция, сдвиг вправо на 2-1=1 байт => 0x111213 & 0xFF = 0x13,
		сдвиг влево на 4-2=2 байта(16 бит) => 0x13000
  3 байт - 02, 2 позиция, сдвиг вправо на 3-1=2 байта => 0x1112 & 0xFF = 0x12,
 		сдвиг влево на 4-3=1 байт => 0x1200
  4 байт - 01, 3 позиция, сдвиг вправо на 4-1=3 байта => 0x11, сдвиг влево на 4-4 байт => 0x11
  0x14000000 | 0x13000 | 0x1200 | 0x11 = 0x14131211
*/

func TestСonversion(t *testing.T) {
	tests := map[string]struct {
		number uint32
		result uint32
	}{
		"test case #1": {
			number: 0x00000000,
			result: 0x00000000,
		},
		"test case #2": {
			number: 0xFFFFFFFF,
			result: 0xFFFFFFFF,
		},
		"test case #3": {
			number: 0x00FF00FF,
			result: 0xFF00FF00,
		},
		"test case #4": {
			number: 0x0000FFFF,
			result: 0xFFFF0000,
		},
		"test case #5": {
			number: 0x01020304,
			result: 0x04030201,
		},
		"test case #6": {
			number: 0x11121314,
			result: 0x14131211,
		},
		"test case #7": {
			number: 0x90090014,
			result: 0x14000990,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ToLittleEndian(test.number)
			assert.Equal(t, test.result, result)
		})
	}
}
