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

	// One-liner
	swapped := ((b & 0b01010101) << 1) | ((b & 0b10101010) >> 1)

	fmt.Printf("Bit-swapped value: %08b  0x%02x\n", swapped, swapped)
}
