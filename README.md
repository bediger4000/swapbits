# Daily Coding Problem: Problem #693 [Medium]

This problem was asked by Cisco.

Given an unsigned 8-bit integer,
swap its even and odd bits.
The 1st and 2nd bit should be swapped,
the 3rd and 4th bit should be swapped, and so on.

For example, 10101010 should be 01010101. 11100010 should be 11010001.

Bonus: Can you do this in one line?

## Analysis

I wrote two programs:

* [brute force](b2.go)
* [a one liner](b3.go)

There's probably a fair number of brute force variants.
Mine finds an even-numbered position bit,
and an odd-numbered postition bit.
It moves the even-numbered-position bit one place left,
the odd-numbered-position bit one place right,
then bitwise-ors them together.
This happens 4 times, for even-numbered bits 0, 2, 4, 6,
and odd-umbered bits 1, 3, 5, 7.
This is fairly fiddly.
I pre-calculated 8 values, each with only 1 bit set,
to be able to sort out the even- or odd-number-position bits
from the original value.
Lots of work.

The one-liner is actually not that hard once you've done the brute force version.
You can sort out even-numbered-position bits by bitwise-and with 0b01010101.
You can sort out odd-numbered-position bits by bitwise-and with 0b10101010.
Bitwise-and the original with each of those two magic values
give you all the even-numbered-position bits,
and all the odd-numbered-position bits.
Right-shift the odd-numbered-position bits,
left-shift the even-numbered-position bits,
and bitwise-or those two values together.
This is in effect unrolling the loop of the brute-force version.
The program puts together even-numbered-position bits,
and puts together odd-numbered-position bits.
Because bitwise operations can be composed and are commutative,
this lets you work everything into 2 bitwise-ors,
shift them appropriately,
the or-together the shifted values.

This might be a good problem for Cisco,
a company doing Internet Protocol routing and firewalling hardware.
Cisco software and firmware has to do a lot of bitwise operations.
This problem emphasizes bitwise operations,
and some clever-trousers obfuscated programming ("do this in one line").

I have to agree with the "medium" rating in this case.
It's not trivial, it's difficult to get correct.
For Cisco, it might be a way of getting at a candidate's
bitwise-operation knowledge.
If the job in question does not do a lot of bitwise operations,
this is not a great question.
