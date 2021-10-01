package computer_science

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSqrt(t *testing.T) {
	expected := 1000.0
	output := Sqrt(1000000)
	require.Equal(t, expected, output)
}

func TestPowersOfTwo(t *testing.T) {
	expected := 512
	output := PowersOfTwo(8)
	require.Equal(t, expected, output)
}

func TestHarmonic(t *testing.T) {
	expected := 2.9289682539682538
	output := Harmonic(10)
	require.Equal(t, expected, output)
}

func TestBinary(t *testing.T) {
	expected := "1101010"
	output := Binary(106)
	require.Equal(t, expected, output)
}

func TestIntegerToBinary(t *testing.T) {
	expected := "101101110"
	output := IntegerToBinary(366)
	require.Equal(t, expected, output)
}

func TestFactors(t *testing.T) {
	expected := []int64{2, 2, 11, 41, 271, 9091}
	output := Factors(4444444444)
	require.Equal(t, expected, output)
}
