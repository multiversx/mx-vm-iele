// File provided by the K Framework Go backend. Timestamp: 2019-08-28 14:13:50.189

package ieletestingmodel

import (
	"math/big"
)

// IntBitRange is modelled on K rule:
// bitRangeInt(I::Int, IDX::Int, LEN::Int) => (I >>Int IDX) modInt (1 <<Int LEN)
func (ms *ModelState) IntBitRange(refI, refOffset, refLen KReference) (KReference, bool) {
	if ms.IsZero(refI) {
		return IntZero, true // any operation on zero will result in zero
	}
	if ms.IsZero(refLen) {
		return IntZero, true // length = 0 means the result is zero
	}

	length, lengthOk := ms.GetPositiveInt(refLen)
	if !lengthOk {
		return NullReference, false
	}

	bigI, iOk := ms.GetBigInt(refI)
	if !iOk {
		return NullReference, false
	}

	offset, offsetOk := ms.GetPositiveInt(refOffset)
	if !offsetOk {
		if bigI.Sign() > 0 {
			// means it doesn't fit in an int32, so a huge number
			// huge offset means that certainly no 1 bits will be caught
			// scenario occurs in tests/VMTests/vmIOandFlowOperations/byte1/byte1.iele.json
			// but only if the number is positive, otherwise the result would be a ridiculously large number of 1's
			return IntZero, true
		}
		return NullReference, false
	}

	if offset&7 != 0 || length&7 != 0 {
		// this is a quick check that they are both divisible by 8
		// as long as they are divisible by 8, we can operate on whole bytes
		// if they are not, things get more complicated, will only implement when necessary
		return NullReference, false
	}
	offsetBytes := offset >> 3 // divide by 8 to get number of bytes
	lengthBytes := length >> 3 // divide by 8 to get number of bytes

	resultBytes := BigIntToTwosComplementBytes(bigI, lengthBytes+offsetBytes)
	if offsetBytes != 0 {
		resultBytes = resultBytes[0:lengthBytes]
	}

	result := new(big.Int)
	result.SetBytes(resultBytes)
	return ms.FromBigInt(result), true
}

// IntSignExtendBitRange is modelled on K rule:
// signExtendBitRangeInt(I::Int, IDX::Int, LEN::Int) => (bitRangeInt(I, IDX, LEN) +Int (1 <<Int (LEN -Int 1))) modInt (1 <<Int LEN) -Int (1 <<Int (LEN -Int 1))
func (ms *ModelState) IntSignExtendBitRange(refI, refOffset, refLen KReference) (KReference, bool) {
	if ms.IsZero(refI) {
		return IntZero, true // any operation on zero will result in zero
	}
	if ms.IsZero(refLen) {
		return IntZero, true // length = 0 means the result is zero
	}

	length, lengthOk := ms.GetPositiveInt(refLen)
	if !lengthOk {
		return NullReference, false
	}

	bigI, iOk := ms.GetBigInt(refI)
	if !iOk {
		return NullReference, false
	}

	offset, offsetOk := ms.GetPositiveInt(refOffset)
	if !offsetOk {
		if bigI.Sign() > 0 {
			// means it doesn't fit in an int32, so a huge number
			// huge offset means that certainly no 1 bits will be caught
			// scenario occurs in tests/VMTests/vmIOandFlowOperations/byte1/byte1.iele.json
			// but only if the number is positive, otherwise the result would be a ridiculously large number of 1's
			return IntZero, true
		}
		return NullReference, false
	}

	if offset&7 != 0 || length&7 != 0 {
		// this is a quick check that they are both divisible by 8
		// as long as they are divisible by 8, we can operate on whole bytes
		// if they are not, things get more complicated, will only implement when necessary
		return NullReference, false
	}
	offsetBytes := offset >> 3 // divide by 8 to get number of bytes
	lengthBytes := length >> 3 // divide by 8 to get number of bytes

	resultBytes := BigIntToTwosComplementBytes(bigI, lengthBytes+offsetBytes)
	if offsetBytes != 0 {
		resultBytes = resultBytes[0:lengthBytes]
	}

	result := TwosComplementBytesToBigInt(resultBytes)
	return ms.FromBigInt(result), true
}

// BigIntToTwosComplementBytes returns a byte array representation, 2's complement if number is negative
// big endian
func BigIntToTwosComplementBytes(i *big.Int, bytesLength int) []byte {
	var resultBytes []byte
	switch i.Sign() {
	case -1:
		// compute 2's complement
		plus1 := big.NewInt(0)
		plus1.Add(i, big.NewInt(1)) // add 1
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

// TwosComplementBytesToBigInt convert a byte array to a number
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
		result.Sub(result, bigOne) // -1
	}

	return result
}
