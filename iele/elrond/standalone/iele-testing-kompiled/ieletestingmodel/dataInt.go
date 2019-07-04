// File provided by the K Framework Go backend. Timestamp: 2019-07-04 13:14:15.638

package ieletestingmodel

import (
	"fmt"
	"math"
	"math/big"
)

const maxSmallInt = math.MaxInt32
const minSmallInt = math.MinInt32

var maxSmallIntAsBigInt = big.NewInt(maxSmallInt)
var minSmallIntAsBigInt = big.NewInt(minSmallInt)

// only attempt to multiply as small int numbers less than the sqrt of this max, by a safety margin
// otherwise play it safe and perform big.Int multiplication
var maxSmallMultiplicationInt = int32(math.Sqrt(float64(math.MaxInt32))) - 100
var minSmallMultiplicationInt = -maxSmallMultiplicationInt

// only attempt to parse as small int strings shorter than this
var maxSmallIntStringLength = len(fmt.Sprintf("%d", maxSmallIntAsBigInt)) - 2

// contains a big.Int corresponding to every small int constant
var smallToBigIntConstants map[int32]*big.Int

// bigInt is a KObject representing a big int in K
type bigInt struct {
	referenceCount int
	recycleCount   uint32
	reuseStatus    objectReuseStatus
	bigValue       *big.Int
}

func fitsInSmallIntReference(i int64) bool {
	return i >= minSmallInt && i <= maxSmallInt
}

func smallMultiplicationSafe(a, b int32) bool {
	return a >= minSmallMultiplicationInt && a <= maxSmallMultiplicationInt &&
		b >= minSmallMultiplicationInt && b <= maxSmallMultiplicationInt
}

func smallIntReference(i int32) KReference {
	if i < 0 {
		return KReference{refType: smallNegativeIntRef, value1: uint32(-i)}
	}
	return KReference{refType: smallPositiveIntRef, value1: uint32(i)}
}

func getSmallInt(ref KReference) (int32, bool) {
	if ref.refType == smallPositiveIntRef {
		return int32(ref.value1), true
	}
	if ref.refType == smallNegativeIntRef {
		return -int32(ref.value1), true
	}
	return 0, false
}

// recycle big Int, or create a new one
func (ms *ModelState) newBigIntObject() (KReference, *bigInt) {
	recycleBinSize := len(ms.bigIntRecycleBin)
	if len(ms.bigIntRecycleBin) > 0 {
		// pop
		recycled := ms.bigIntRecycleBin[recycleBinSize-1]
		ms.bigIntRecycleBin = ms.bigIntRecycleBin[:recycleBinSize-1]
		if recycled.constantObject {
            panic("constant ended up in bigInt recycle bin")
        }

		// update object
		bigObj, isBigObj := ms.getBigIntObject(recycled)
		if !isBigObj {
			panic("recycled bigInt is in fact not a big int reference")
		}
        if bigObj.reuseStatus != inRecycleBin {
			panic("recycled bigInt does not have status inRecycleBin")
		}
		bigObj.reuseStatus = active

		bigObj.recycleCount++
		recycled.value2++ // we match value2 with the recycleCount

		return recycled, bigObj
	}

	return ms.newBigIntObjectNoRecycle()
}

func (ms *ModelState) newBigIntObjectNoRecycle() (KReference, *bigInt) {
	newIndex := len(ms.bigInts)
	bigObj := &bigInt{referenceCount: 0, recycleCount: 0, reuseStatus: active, bigValue: big.NewInt(0)}
	ms.bigInts = append(ms.bigInts, bigObj)
	newRef := KReference{refType: bigIntRef, constantObject: false, value1: uint32(newIndex), value2: 0}
	return newRef, bigObj
}

func (ms *ModelState) getBigIntObject(ref KReference) (*bigInt, bool) {
	if ref.refType != bigIntRef {
		return nil, false
	}
	index := int(ref.value1)
	if ref.constantObject {
		return constantsModel.bigInts[index], true
	}
	if index >= len(ms.bigInts) {
		panic("trying to reference object beyond allocated objects")
	}
	obj := ms.bigInts[index]
	if ref.value2 != obj.recycleCount {
	    panic("reference points to bigInt that was recycled in the mean time and can no longer be used in this context")
	}
	return obj, true
}

func convertSmallIntRefToBigInt(ref KReference) (*big.Int, bool) {
	small, isSmall := getSmallInt(ref)
	if isSmall {
		if smallToBigIntConstants != nil {
			bigIntConstant, found := smallToBigIntConstants[small]
			if found {
				return bigIntConstant, true
			}
		}
		return big.NewInt(int64(small)), true
	}
	return nil, false
}

// IsInt returns true if reference points to an integer
func IsInt(ref KReference) bool {
	return ref.refType == smallPositiveIntRef || ref.refType == smallNegativeIntRef || ref.refType == bigIntRef
}

// IntZero is a reference to the constant integer 0
var IntZero = smallIntReference(0)

// IntOne is a reference to the constant integer 1
var IntOne = smallIntReference(1)

// IntMinusOne is a reference to the constant integer -1
var IntMinusOne = smallIntReference(-1)

// FromBigInt provides a reference to an integer (big or small)
func (ms *ModelState) FromBigInt(bi *big.Int) KReference {
	// attempt to make it small
	if bi.IsInt64() {
		biInt64 := bi.Int64()
		if biInt64 >= minSmallInt && biInt64 <= maxSmallInt {
			return smallIntReference(int32(biInt64))
		}
	}
	// make it big
	ref, obj := ms.newBigIntObject()
	obj.bigValue.Set(bi)
	return ref
}

// NewIntConstant creates a new integer constant, which is saved statically.
// Do not use for anything other than constants, since these never get cleaned up.
func NewIntConstant(stringRepresentation string) KReference {
	ref := constantsModel.IntFromString(stringRepresentation)
	ref.constantObject = true

	// if a small constant, also create a big.Int constant
	// if we don't create them now as constants, they will keep getting created at runtime
	small, isSmall := getSmallInt(ref)
	if isSmall {
		if smallToBigIntConstants == nil {
			smallToBigIntConstants = make(map[int32]*big.Int)
		}
		smallToBigIntConstants[small] = big.NewInt(int64(small))
	}

	return ref
}

// FromInt converts a Go integer to an integer in the model
func (ms *ModelState) FromInt(x int) KReference {
	if x >= minSmallInt && x <= maxSmallInt {
		return smallIntReference(int32(x))
	}
	ref, obj := ms.newBigIntObject()
	obj.bigValue.SetInt64(int64(x))
	return ref
}

// FromInt64 converts a int64 to an integer in the model
func (ms *ModelState) FromInt64(x int64) KReference {
	if x >= minSmallInt && x <= maxSmallInt {
		return smallIntReference(int32(x))
	}
	ref, obj := ms.newBigIntObject()
	obj.bigValue.SetInt64(x)
	return ref
}

// FromUint64 converts a uint64 to an integer in the model
func (ms *ModelState) FromUint64(x uint64) KReference {
	if x <= maxSmallInt {
		return smallIntReference(int32(x))
	}
	ref, obj := ms.newBigIntObject()
	obj.bigValue.SetUint64(x)
	return ref
}
