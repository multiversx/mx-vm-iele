// File provided by the K Framework Go backend. Timestamp: 2019-06-24 20:04:33.113

package ieletestingmodel

import (
	"math"
	"math/big"
)

// IntZero ... K Int with value zero
var IntZero = &Int{Value: big.NewInt(0)}

// IntOne ... K Int with value 1
var IntOne = &Int{Value: big.NewInt(1)}

// IntMinusOne ... K Int with value -1
var IntMinusOne = &Int{Value: big.NewInt(-1)}

// NewInt ... provides new Int instance
func NewInt(bi *big.Int) *Int {
	return &Int{Value: bi}
}

// NewIntFromInt ... provides new Int instance
func NewIntFromInt(x int) *Int {
	return NewIntFromInt64(int64(x))
}

// NewIntFromInt64 ... provides new Int instance
func NewIntFromInt64(x int64) *Int {
	return &Int{Value: big.NewInt(x)}
}

// NewIntFromUint64 ... provides new Int instance
func NewIntFromUint64(x uint64) *Int {
	var z big.Int
	z.SetUint64(x)
	return &Int{Value: &z}
}

// NewIntFromByte ... provides new Int instance
func NewIntFromByte(x byte) *Int {
	var z big.Int
	z.SetUint64(uint64(x))
	return &Int{Value: &z}
}

// NewIntFromBytes ... provides new Int instance from byte array
func NewIntFromBytes(bytes []byte) *Int {
	z := big.NewInt(0)
	z.SetBytes(bytes)
	return &Int{Value: z}
}

// ParseInt ... creates K int from string representation
func ParseInt(s string) (*Int, error) {
	b := big.NewInt(0)
	if s != "0" {
		b.UnmarshalText([]byte(s))
		if b.Cmp(IntZero.Value) == 0 {
			return IntZero, &parseIntError{parseVal: s}
		}
	}
	return NewInt(b), nil
}

// NewIntFromString ... same as ParseInt but panics instead of error
func NewIntFromString(s string) *Int {
	i, err := ParseInt(s)
	if err != nil {
		panic(err)
	}
	return i
}

// IsZero ... true if item represents number 0
func (k *Int) IsZero() bool {
	return k.Value.Sign() == 0
}

// IsPositive ... true if represented number is >= 0
func (k *Int) IsPositive() bool {
	return k.Value.Sign() >= 0
}

// IsNegative ... true if represented number is < 0
func (k *Int) IsNegative() bool {
	return k.Value.Sign() < 0
}

// ToUint32 ... converts to uint if possible, returns (0, false) if not
func (k *Int) ToUint32() (uint, bool) {
	if !k.Value.IsUint64() {
		return 0, false
	}

	u64 := k.Value.Uint64()
	if u64 > math.MaxUint32 {
		return 0, false
	}

	return uint(u64), true
}

// ToInt32 ... converts to int if possible, returns (0, false) if not
func (k *Int) ToInt32() (int, bool) {
	if !k.Value.IsInt64() {
		return 0, false
	}

	i64 := k.Value.Int64()
	if i64 < math.MinInt32 || i64 > math.MaxInt32 {
		return 0, false
	}

	return int(i64), true
}

// ToPositiveInt32 ... converts to int32 if possible, returns (0, false) if not
// also rejects negative numbers, so we don't have to test for that again
func (k *Int) ToPositiveInt32() (int, bool) {
	if !k.Value.IsInt64() {
		return 0, false
	}

	i64 := k.Value.Int64()
	if i64 < 0 || i64 > math.MaxInt32 {
		return 0, false
	}

	return int(i64), true
}

// ToByte ... converts to 1 byte if possible, returns (0, false) if not
func (k *Int) ToByte() (byte, bool) {
	if !k.Value.IsUint64() {
		return 0, false
	}

	u64 := k.Value.Uint64()
	if u64 > 255 {
		return 0, false
	}

	return byte(u64), true
}

// BigIntToTwosComplementBytes ... returns a byte array representation, 2's complement if number is negative
// big endian
func BigIntToTwosComplementBytes(i *big.Int, bytesLength int) []byte {
	var resultBytes []byte
	switch i.Sign() {
	case -1:
		// compute 2's complement
		plus1 := big.NewInt(0)
		plus1.Add(i, IntOne.Value) // add 1
		plus1Bytes := plus1.Bytes()
		offset := len(plus1Bytes) - bytesLength
		resultBytes = make([]byte, bytesLength)
		for i := 0; i < bytesLength; i++ {
			j := offset + i
			if j < 0 {
				resultBytes[i] = 255 // pad left with 11111111
			} else {
				resultBytes[i] = ^plus1Bytes[j] // also negate every bit
			}
		}
		break
	case 0:
		// just zeroes
		resultBytes = make([]byte, bytesLength)
		break
	case 1:
		originalBytes := i.Bytes()
		resultBytes = make([]byte, bytesLength)
		offset := len(originalBytes) - bytesLength
		for i := 0; i < bytesLength; i++ {
			j := offset + i
			if j < 0 {
				resultBytes[i] = 0 // pad left with 00000000
			} else {
				resultBytes[i] = originalBytes[j]
			}
		}
		break
	}

	return resultBytes
}

// TwosComplementBytesToBigInt ... convert a byte array to a number
// interprets input as a 2's complement representation if the first bit (most significant) is 1
// big endian
func TwosComplementBytesToBigInt(twosBytes []byte) *big.Int {

	testBit := twosBytes[0] >> 7
	result := new(big.Int)
	if testBit == 0 {
		// positive number, no further processing required
		result.SetBytes(twosBytes)
	} else {
		// convert to negative number
		notBytes := make([]byte, len(twosBytes))
		for i, b := range twosBytes {
			notBytes[i] = ^b // negate every bit
		}
		result.SetBytes(notBytes)
		result.Neg(result)
		result.Sub(result, IntOne.Value) // -1
	}

	return result
}
