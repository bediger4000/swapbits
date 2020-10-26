package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

/*
Given an unsigned 8-bit integer,
swap its even and odd bits.
The 1st and 2nd bit should be swapped,
the 3rd and 4th bit should be swapped, and so on.

For example, 10101010 should be 01010101. 11100010 should be 11010001.

Bonus: Can you do this in one line?
*/

func main() {

	ui, err := strconv.ParseUint(os.Args[1], 0x10, 8)
	if err != nil {
		log.Fatal(err)
	}
	b := uint8(ui)
	fmt.Printf("Original value:    %08b  0x%02x\n", b, b)
	swapped := swapBits(b)
	fmt.Printf("Bit-swapped value: %08b  0x%02x\n", swapped, swapped)
}

var evenBits = []uint8{0b00000001, 0b00000100, 0b00010000, 0b01000000}
var oddBits = []uint8{0b00000010, 0b00001000, 0b00100000, 0b10000000}

func swapBits(b uint8) uint8 {
	var r uint8
	for i := range evenBits {
		evenbit := b & evenBits[i]
		oddbit := b & oddBits[i]
		r |= (evenbit << 1)
		r |= (oddbit >> 1)
	}
	return r
}
