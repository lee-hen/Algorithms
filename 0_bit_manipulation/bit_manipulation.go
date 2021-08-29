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

// ModifyBit
// 00000110 x
// 00000101 6
// 00000001 state
// mask = 00000001 << 00000101 = 00100000
// ^mask = 11011111
// x & ^mask = 00000110 & 11011111 = 00000110  -> the same as clear bit
// -state = 11111111
// -state & mask = 11111111 & 00100000 = 00100000
// x & ^mask | (-state & mask) = 00000110 | 00100000 = 00100110
func ModifyBit(x, state int, position uint) int {
	mask := 1 << position
	return  x & ^mask | (-state & mask)
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
	return (x & x-1) == 0
}

func Abs(x int) int {
	// All major processors represent negative numbers using the two's-complement which is defined as:
	// for x ≥ 0 → x
	// for x < 0 → NOT(x) + 1
	//
	// On the lowest level, computers provide logical bit shifts and arithmetic bit shifts.
	//	Both shifts differ in handling how to fill the empty bits on the left side.
	//	Logical shifts insert zeros while arithmetic shifts replicate the formerly highest bit.
	//
	//	Whereas signed integers are arithmetically shifted in C, unsigned integers are logically shifted.
	//
	//	In our case x is shifted arithmetically 31 times to the right which basically erases its value
	// and spreads the highest bit. That means, line 3 evaluates either to 0x00000000 (→ 0) or
	// 0xFFFFFFFF (→ −1).
	// Note: 32 bit systems require a shift by 31, 64 bit systems require a shift by 63 accordingly.
	//
	// Consequently, line 4 turns out to be (x XOR 0) − 0 for positive values of x (including x=0).
	// x XOR 0 is still x and x − 0 remains x, too. So for positive x we get x ≥ 0 → x.
	//
	// We saw that for negative values of x, bit31 is set to 0xFFFFFFFF.
	// Line 4 is then (x XOR 0xFFFFFFFF) − 0xFFFFFFFF. The bracketed XOR is equivalent to NOT(x) and
	// the constant −0xFFFFFFFF turns out to be −(-1) = +1.
	// In the end, the whole term is NOT(x) + 1, exactly what we wanted: for x < 0 → NOT(x) + 1
	//
	// Note: Current C++ compilers (Microsoft, GCC and Intel) implemented a special assembler code sequence
	// for abs which runs faster than the shown algorithm on x86 CPUs (see below for its source code).
    // designed for 32 bits (simple modification for 64 bits possible)

	bit31 := x >> 31
	return (x ^ bit31) - bit31
}

// others
// https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetKernighan

