package bit_manipulation

// And
//     0101 (decimal 5)
// AND 0011 (decimal 3)
//   = 0001 (decimal 1)
func And(x, y int) int {
	return x & y
}

// Or
//    0101 (decimal 5)
// OR 0011 (decimal 3)
//  = 0111 (decimal 7)
func Or(x, y int) int {
	return x | y
}

// Xor
//     0101 (decimal 5)
// XOR 0011 (decimal 3)
//   = 0110 (decimal 6)
func Xor (x, y int) int {
	return x ^ y
}

// How to convert binary to decimal
// The decimal number is equal to the sum of binary digits (dn) times their power of 2 (2n):
// decimal = d0×2^0 + d1×2^1 + d2×2^2 + ...

// 111001 = 1⋅2^5+1⋅2^4+1⋅2^3+0⋅2^2+0⋅2^1+1⋅2^0 = 57
// 100011 = 1⋅2^5+0⋅2^4+0⋅2^3+0⋅2^2+1⋅2^1+1⋅2^0 = 35

// unsigned value
// 10010110
// 128 + 16 + 4 + 2 = 150

// signed value
// 10010110
// -128 + 16 + 4 + 2 = -106

// two's component
// 00000000 = 0
// 11111111 = -1  0xff
// -2^7 + 2^6 + 2^5 + 2^4 + 2^3 + 2^2 + 2^1 + 1 = -1

// 00000100 + ? = -0
// 00000100  4
//           +
// 11111100 -4 -> answer
// 10000000 -0
// -2^7 + 2^6 + 2^5 + 2^4 + 2^3 + 2^2 + 0^1 + 0 = -4

// how to find -integer to binary
// 1. convert integer to binary
// 2. invert bits
// 3. Add 1

// -123
// 0 1 1 1 1 0 1 1
//   0 0 0 0 1 0 0
//   0 0 0 0 1 0 1
// 1 0 0 0 0 1 0 1

//  x = 01111011
// -x = ?
// ^x = 10000100
// +1 = 00000001
// -x = 10000101 = ^x + 1
// -x = ^x + 1

// SetBit
// 00000110
// 00000101 5
// mask 00000001 << 00000101 = 00100000
// 00100000 | 00000110 = 00100110
func SetBit(x int, position uint) int {
	mask := 1 << position
	return x | mask
}

// ClearBit
// 00000110
// 00000010 2
// mask 00000001 << 00000010 = 00000100
// XOR ^00000100 = 11111011
// 11111011 & 00000110 = 00000010
func ClearBit(x int, position uint) int {
	mask := 1 << position
	return x &^ mask
}

// FlipBit
// 01100110
// 00000010 2
// mask 00000001 << 00000010 = 00000100
// XOR 00000100 ^ 01100110 = 01100010
func FlipBit(x int, position uint) int {
	mask := 1 << position
	return x ^ mask
}

// IsBitSet
// 01100110
// 00000101 5
// shift 01100110 >> 00000101 = 00000011
// 00000001 & 00000011 = 00000001
func IsBitSet(x int, position uint) bool {
	shift := x >> position
	return shift & 1 == 1
}

// CheckEven
// 0110 6
// 0001
// 0000
func CheckEven(x int) bool {
	return x & 1 == 0
}

// CheckPowerOfTwo
// 1000     8
// 0111     7 x-1
// 0000
func CheckPowerOfTwo(x int) bool {
	return (x & (x-1)) == 0
}

// Swap
// *x = 10111101
// *y = 00101110
// *x = 10010011 = *x^*y
// *y = 10111101 = *x^*y
// *x = 00101110 = *x^*y
func Swap(x, y *int) {
	*x = *x ^ *y // flip *x with *y
	*y = *x ^ *y // flip back to origin x
	*x = *x ^ *y // flop back to origin y
}

// ModifyBit
// 00000110 x
// 00000101 5
// 00000001 state
// mask = 00000001 << 00000101 = 00100000
// ^mask = 11011111
// x & ^mask = 00000110 & 11011111 = 00000110  -> the same as clear bit
// -state = 11111111
// -state & mask = 11111111 & 00100000 = 00100000
// x & ^mask | (-state & mask) = 00000110 | 00100000 = 00100110
// when state is 4
// 00000100 state
// 11111100 -state
// 00100000  mask
// 00100000  -state & mask
func ModifyBit(x, state int, position uint) int {
	mask := 1 << position
	return  x &^ mask | (-state & mask)
}

// Abs
// All major processors represent negative numbers using the two's-complement which is defined as:
// for x ≥ 0 → x
// for x < 0 → NOT(x) + 1
// On the lowest level, computers provide logical bit shifts and arithmetic bit shifts.
// Both shifts differ in handling how to fill the empty bits on the left side.
// Logical shifts insert zeros while arithmetic shifts replicate the formerly highest bit.
// Whereas signed integers are arithmetically shifted in C, unsigned integers are logically shifted.
// In our case x is shifted arithmetically 31 times to the right which basically erases its value
// and spreads the highest bit. That means, line 3 evaluates either to 0x00000000 (→ 0) or 0xFFFFFFFF (→ −1).
// Note: 32 bit systems require a shift by 31, 64 bit systems require a shift by 63 accordingly.
// Consequently, line 2 turns out to be (x XOR 0) − 0 for positive values of x (including x=0).
// x XOR 0 is still x and x − 0 remains x, too. So for positive x we get x ≥ 0 → x.
// We saw that for negative values of x, bit31 is set to 0xFFFFFFFF.
// Line 4 is then (x XOR 0xFFFFFFFF) − 0xFFFFFFFF. The bracketed XOR is equivalent to NOT(x) and
// the constant −0xFFFFFFFF turns out to be −(-1) = +1.
// In the end, the whole term is NOT(x) + 1, exactly what we wanted: for x < 0 → NOT(x) + 1

// Abs32
// 01111011  x -> 123
// 00000000  bit7 -> x >> 7 -> 123/128 -> 0 // logical shifts: insert zeros
// 01111011  x^bit7
// 01111011 - 00000000 = 01111011 = x^bit7 - bit7
func Abs32(x int32) int32 {
	bit31 := x >> 31
	return (x ^ bit31) - bit31
}

// Abs8
// 10000101  x -> -123
// 11111111  bit7 -> x >> 7 -> -123/128 -> -1 // arithmetic shifts: replicate the formerly highest bit
// 01111010  x^bit7
// 01111010 - 11111111 = 01111010 + 00000001 = 01111011 = x^bit7 - bit7
func Abs8(x int8) int8 {
	 bit7 := x >> 7
	 return (x ^ bit7) - bit7
}

// PopulationCount
// count the number of 1 bits in a word x
// 11010000 x
// 11001111 x-1
// 11000000 x &= x-1   1st
// 10111111 x-1
// 10000000 x &= x-1   2nd
// 00000000 x-1
// 00000000 x &= x-1   3rd
func PopulationCount(x int) int {
	var r int
	for x != 0 {
		x &= x-1  // the same as check power of two
		r++
	}
	return r
}

// PopCount
// 101010 42
// i & 1 = the last bit of i
// count(101010) = count(10101) + 0
// count(10101)  = count(1010)  + 1
// count(1010)   = count(101)   + 0
// count(101)    = count(10)    + 1
// count(10)     = count(1)     + 0
// count(1)      = count(0)     + 1
func PopCount(x uint64) int {
	// build up table
	// store each eight bit words of 1s.
	var count [256]byte
	for i := range count {
		count[i] = count[i>>1] + byte(i&1) // count[i>>1] = count[i/2]
	}

	var r int
	for x != 0 {
		r += int(count[x & 0xff]) // 0xff -> 11111111 -> -1
		x >>= 8
	}

	return r
}

func PopCount2(x uint64) int {
	var pc [256]byte
	for i := range pc {
		pc[i] = pc[i/2] + byte(i & 1) // pc[i/2] = pc[i>>1]
	}

	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// others
// https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetKernighan

