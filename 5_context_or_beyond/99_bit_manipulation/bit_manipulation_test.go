package bit_manipulation

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAnd(t *testing.T) {
	expected := 1
	output := And(5,3)
	require.Equal(t, expected, output)
}

func TestOr(t *testing.T) {
	expected := 7
	output := Or(5,3)
	require.Equal(t, expected, output)
}

func TestXor(t *testing.T) {
	expected := 6
	output := Xor(5,3)
	require.Equal(t, expected, output)
}

func TestSetBit(t *testing.T) {
	expected := 38
	output := SetBit(6,5)
	require.Equal(t, expected, output)
}

func TestClearBit(t *testing.T) {
	expected := 2
	output := ClearBit(6,2)
	require.Equal(t, expected, output)
}

func TestFlipBit(t *testing.T) {
	expected := 98
	output := FlipBit(102,2)
	require.Equal(t, expected, output)
}

func TestIsBitSet(t *testing.T) {
	output := IsBitSet(98,5)
	require.True(t, output)
}

func TestModifyBit(t *testing.T) {
	expected := 38
	output := ModifyBit(6,1, 5)
	require.Equal(t, expected, output)
}

func TestCheckEven(t *testing.T) {
	output := CheckEven(6)
	require.True(t, output)
}

func TestCheckPowerOfTwo(t *testing.T) {
	output := CheckPowerOfTwo(8)
	require.True(t, output)
}

func TestAbs(t *testing.T) {
	expected := 10
	output := Abs(-10)
	require.Equal(t, expected, output)
}

func TestSwap(t *testing.T) {
	x, y := 3, 4
	Swap(&x, &y)
	require.Equal(t, x, 4)
	require.Equal(t, y, 3)
}

func TestPopulationCount(t *testing.T) {
	expected := 3
	output := PopulationCount(416)
	require.Equal(t, expected, output)
}

func TestPopCount(t *testing.T) {
	expected := 3
	output := PopCount(416)
	require.Equal(t, expected, output)
}

func TestPopCount2(t *testing.T) {
	expected := 3
	output := PopCount2(416)
	require.Equal(t, expected, output)
}
