// File provided by the K Framework Go backend. Timestamp: 2019-07-15 11:19:23.686

package ieletestingmodel

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

// IntFromByte converts a byte to an integer in the model
func (ms *ModelState) IntFromByte(x byte) KReference {
	return ms.FromInt(int(x))
}

// IntFromBytes converts a byte array to an integer in the model
func (ms *ModelState) IntFromBytes(bytes []byte) KReference {
	z := big.NewInt(0)
	z.SetBytes(bytes)
	return ms.FromBigInt(z)
}

// ParseInt creates K int from string representation
func (ms *ModelState) ParseInt(str string) (KReference, error) {
	if str == "0" {
		return IntZero, nil
	}
	if len(str) < maxSmallIntStringLength {
		i, err := strconv.Atoi(str)
		if err != nil {
			return NullReference, &parseIntError{parseVal: str}
		}
		return createKrefSmallInt(int64(i)), nil
	}

	b := big.NewInt(0)
	b.UnmarshalText([]byte(str))
	if b.Sign() == 0 {
		return IntZero, &parseIntError{parseVal: str}
	}
	return ms.FromBigInt(b), nil
}

// ParseIntFromBase creates K int from string representation in a given base
func (ms *ModelState) ParseIntFromBase(str string, base int) (KReference, error) {
	if base == 10 {
		return ms.ParseInt(str)
	}
	if str == "0" {
		return IntZero, nil
	}
	b := big.NewInt(0)
	_, ok := b.SetString(str, base)
	if !ok {
		return IntZero, &parseIntError{parseVal: str}
	}
	return ms.FromBigInt(b), nil
}

// IntFromString does the same as ParseInt but panics instead of returning an error
func (ms *ModelState) IntFromString(s string) KReference {
	i, err := ms.ParseInt(s)
	if err != nil {
		panic(err)
	}
	return i
}

// GetBigIntUnsafe yields a big.Int cast from any K integer object, if possible.
// Can retrieve objects from the model.
// Only use if you are absolutely certain that the retrieved object will not be changed!!!
func (ms *ModelState) GetBigIntUnsafe(ref KReference) (*big.Int, bool) {
	bigFromSmall, isSmall := convertSmallIntRefToBigInt(ref)
	if isSmall {
		return bigFromSmall, true
	}

	bi, isBigInt := ms.getBigIntObject(ref)
	if isBigInt {
		return bi.bigValue, true
	}
	return nil, false
}

// GetBigInt yields a big.Int cast from any K integer object, if possible.
// Does not provide any big.Int object from the model, only copies,
// so it is safe to use anywhere.
func (ms *ModelState) GetBigInt(ref KReference) (*big.Int, bool) {
	small, isSmall := parseKrefSmallInt(ref)
	if isSmall {
		return big.NewInt(int64(small)), true
	}

	bi, isBigInt := ms.getBigIntObject(ref)
	if isBigInt {
		return big.NewInt(0).Set(bi.bigValue), true
	}
	return nil, false
}

// IsZero returns true if an item represents number 0
func (ms *ModelState) IsZero(ref KReference) bool {
	small, isSmall := parseKrefSmallInt(ref)
	if isSmall {
		return small == 0
	}

	bi, isBigInt := ms.getBigIntObject(ref)
	if isBigInt {
		return bi.bigValue.Sign() == 0
	}

	return false
}

// GetUint converts to uint if possible, returns (0, false) if not
func (ms *ModelState) GetUint(ref KReference) (uint, bool) {
	small, isSmall := parseKrefSmallInt(ref)
	if isSmall {
		if small < 0 {
			return 0, false
		}
		return uint(small), true
	}

	bi, isBigInt := ms.getBigIntObject(ref)
	if isBigInt {
		if !bi.bigValue.IsUint64() {
			return 0, false
		}
		u64 := bi.bigValue.Uint64()
		if u64 > math.MaxUint32 {
			return 0, false
		}
		return uint(u64), true
	}

	return 0, false
}

// GetUint64 converts to uint64 if possible, returns (0, false) if not
func (ms *ModelState) GetUint64(ref KReference) (uint64, bool) {
	small, isSmall := parseKrefSmallInt(ref)
	if isSmall {
		if small < 0 {
			return 0, false
		}
		return uint64(small), true
	}

	bi, isBigInt := ms.getBigIntObject(ref)
	if isBigInt {
		if !bi.bigValue.IsUint64() {
			return 0, false
		}
		return bi.bigValue.Uint64(), true
	}

	return 0, false
}

// GetInt converts to int if possible, returns (0, false) if not
func (ms *ModelState) GetInt(ref KReference) (int, bool) {
	small, isSmall := parseKrefSmallInt(ref)
	if isSmall {
		return int(small), true
	}

	bi, isBigInt := ms.getBigIntObject(ref)
	if isBigInt {
		if !bi.bigValue.IsInt64() {
			return 0, false
		}
		i64 := bi.bigValue.Int64()
		if i64 >= math.MinInt32 && i64 <= math.MaxInt32 {
			return int(i64), true
		}
	}

	return 0, false
}

// GetPositiveInt converts to int32 if possible, returns (0, false) if not.
// Also rejects negative numbers, so we don't have to test for that again.
func (ms *ModelState) GetPositiveInt(ref KReference) (int, bool) {
	small, isSmall := parseKrefSmallInt(ref)
	if isSmall {
		if small < 0 {
			return 0, false
		}
		return int(small), true
	}

	bi, isBigInt := ms.getBigIntObject(ref)
	if isBigInt {
		if !bi.bigValue.IsUint64() {
			return 0, false
		}
		u64 := bi.bigValue.Uint64()
		if u64 > math.MaxUint32 {
			return 0, false
		}
		return int(u64), true
	}

	return 0, false
}

// GetByte converts to 1 byte if possible, returns (0, false) if not
func (ms *ModelState) GetByte(ref KReference) (byte, bool) {
	small, isSmall := parseKrefSmallInt(ref)
	if isSmall {
		if small < 0 || small > 255 {
			return 0, false
		}
		return byte(small), true
	}

	bi, isBigInt := ms.getBigIntObject(ref)
	if isBigInt {
		if !bi.bigValue.IsUint64() {
			return 0, false
		}
		u64 := bi.bigValue.Uint64()
		if u64 > 255 {
			return 0, false
		}
		return byte(u64), true
	}

	return 0, false
}

// GetIntAsDecimalString converts a K integer to a decimal string representation, decimal, if possible.
func (ms *ModelState) GetIntAsDecimalString(ref KReference) (string, bool) {
	small, isSmall := parseKrefSmallInt(ref)
	if isSmall {
		return fmt.Sprintf("%d", small), true
	}

	bigI, isBigInt := ms.getBigIntObject(ref)
	if isBigInt {
		return bigI.bigValue.String(), true
	}

	return "", false
}

// GetIntToString converts a K integer to a string representation in given base, if possible.
func (ms *ModelState) GetIntToString(ref KReference, base int) (string, bool) {
	small, isSmall := parseKrefSmallInt(ref)
	if isSmall {
		return strconv.FormatInt(int64(small), base), true
	}

	bigI, isBigInt := ms.getBigIntObject(ref)
	if isBigInt {
		return bigI.bigValue.Text(base), true
	}

	return "", false
}
