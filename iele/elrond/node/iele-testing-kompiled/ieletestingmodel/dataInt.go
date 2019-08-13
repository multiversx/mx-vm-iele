// File provided by the K Framework Go backend. Timestamp: 2019-08-13 18:10:37.856

package ieletestingmodel

import (
	"fmt"
	"math"
	"math/big"
)

const maxSmallInt int64 = (1 << 59) - 1
const minSmallInt int64 = -maxSmallInt

var maxSmallIntAsBigInt = big.NewInt(maxSmallInt)
var minSmallIntAsBigInt = big.NewInt(minSmallInt)

// only attempt to multiply as small int numbers less than the sqrt of this max, by a safety margin
// otherwise play it safe and perform big.Int multiplication
var maxSmallMultiplicationInt = int64(math.Sqrt(float64(maxSmallInt))) - 100
var minSmallMultiplicationInt = -maxSmallMultiplicationInt

// only attempt to parse as small int strings shorter than this
var maxSmallIntStringLength = len(fmt.Sprintf("%d", maxSmallIntAsBigInt)) - 2

// contains a big.Int corresponding to every small int constant
var smallToBigIntConstants map[int64]*big.Int

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

func smallMultiplicationSafe(a, b int64) bool {
	return a >= minSmallMultiplicationInt && a <= maxSmallMultiplicationInt &&
		b >= minSmallMultiplicationInt && b <= maxSmallMultiplicationInt
}

// recycle big Int, or create a new one
func (ms *ModelState) newBigIntObject() (KReference, *bigInt) {
	return ms.mainData.newBigIntObject()
}

// recycle big Int, or create a new one
func (md *ModelData) newBigIntObject() (KReference, *bigInt) {
	recycleBinSize := len(md.bigIntRecycleBin)
	if len(md.bigIntRecycleBin) > 0 {
		// pop
		recycled := md.bigIntRecycleBin[recycleBinSize-1]
		md.bigIntRecycleBin = md.bigIntRecycleBin[:recycleBinSize-1]
		_, modelRef, refRecycleCount, index := parseKrefBigInt(recycled)
		if modelRef != md.selfRef {
			panic("recycled big Int ended up in the wrong data container")
		}

		// update object
		bigObj, isBigObj := md.getBigIntObject(recycled)
		if !isBigObj {
			panic("recycled bigInt is in fact not a big int reference")
		}
		if bigObj.reuseStatus != inRecycleBin {
			panic("recycled bigInt does not have status inRecycleBin")
		}
		bigObj.reuseStatus = active

		bigObj.recycleCount++
		refRecycleCount++ // we match value2 with the recycleCount
		return createKrefBigInt(md.selfRef, refRecycleCount, index), bigObj
	}

	return md.newBigIntObjectNoRecycle()
}

func (md *ModelData) newBigIntObjectNoRecycle() (KReference, *bigInt) {
	newIndex := uint64(len(md.bigInts))
	bigObj := &bigInt{referenceCount: 0, recycleCount: 0, reuseStatus: active, bigValue: big.NewInt(0)}
	md.bigInts = append(md.bigInts, bigObj)
	newRef := createKrefBigInt(md.selfRef, 0, newIndex)
	return newRef, bigObj
}

func (ms *ModelState) getBigIntObject(ref KReference) (*bigInt, bool) {
	isBigInt, dataRef, _, _ := parseKrefBigInt(ref)
	if !isBigInt {
		return nil, false
	}
	return ms.getData(dataRef).getBigIntObject(ref)
}

func (md *ModelData) getBigIntObject(ref KReference) (*bigInt, bool) {
	isBigInt, dataRef, refRecycleCount, index := parseKrefBigInt(ref)
	if !isBigInt {
		return nil, false
	}
	if dataRef != md.selfRef {
		panic("trying to retrieve big Int from the wrong data container")
	}
	if index >= uint64(len(md.bigInts)) {
		panic("trying to reference object beyond allocated objects")
	}
	obj := md.bigInts[index]
	if refRecycleCount != uint64(obj.recycleCount) {
		panic("reference points to bigInt that was recycled in the mean time and can no longer be used in this context")
	}
	return obj, true
}

func (md *ModelData) recycleAllInts() {
	if cap(md.bigIntRecycleBin) < len(md.bigInts) {
		md.bigIntRecycleBin = make([]KReference, len(md.bigInts))
	} else {
		md.bigIntRecycleBin = md.bigIntRecycleBin[:len(md.bigInts)]
	}
	for i, bo := range md.bigInts {
		bo.referenceCount = 0
		bo.reuseStatus = inRecycleBin
		md.bigIntRecycleBin[i] = createKrefBigInt(md.selfRef, uint64(bo.recycleCount), uint64(i))
	}
}

func convertSmallIntRefToBigInt(ref KReference) (*big.Int, bool) {
	small, isSmall := parseKrefSmallInt(ref)
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
	refType := kreferenceType(ref >> refTypeShift)
	return refType == smallPositiveIntRef || refType == smallNegativeIntRef || refType == bigIntRef
}

// IntZero is a reference to the constant integer 0
var IntZero = createKrefSmallInt(0)

// IntOne is a reference to the constant integer 1
var IntOne = createKrefSmallInt(1)

// IntMinusOne is a reference to the constant integer -1
var IntMinusOne = createKrefSmallInt(-1)

// FromBigInt provides a reference to an integer (big or small)
func (ms *ModelState) FromBigInt(bi *big.Int) KReference {
	return ms.mainData.fromBigInt(bi)
}

func (md *ModelData) fromBigInt(bi *big.Int) KReference {
	// attempt to make it small
	if bi.IsInt64() {
		biInt64 := bi.Int64()
		if biInt64 >= minSmallInt && biInt64 <= maxSmallInt {
			return createKrefSmallInt(biInt64)
		}
	}
	// make it big
	ref, obj := md.newBigIntObject()
	obj.bigValue.Set(bi)
	return ref
}

// NewIntConstant creates a new integer constant, which is saved statically.
// Do not use for anything other than constants, since these never get cleaned up.
func NewIntConstant(stringRepresentation string) KReference {
	ref := constantsData.intFromString(stringRepresentation)
	assertModelDataFlag(ref, constDataRef)

	// if the result is a small int constant, also create a big.Int constant
	// if we don't create them now as constants, they will keep getting created at runtime
	small, isSmall := parseKrefSmallInt(ref)
	if isSmall {
		if smallToBigIntConstants == nil {
			smallToBigIntConstants = make(map[int64]*big.Int)
		}
		smallToBigIntConstants[small] = big.NewInt(int64(small))
	}

	return ref
}

// FromInt converts a Go integer to an integer in the model
func (ms *ModelState) FromInt(x int) KReference {
	if int64(x) >= minSmallInt && int64(x) <= maxSmallInt {
		return createKrefSmallInt(int64(x))
	}
	ref, obj := ms.newBigIntObject()
	obj.bigValue.SetInt64(int64(x))
	return ref
}

// FromInt64 converts a int64 to an integer in the model
func (ms *ModelState) FromInt64(x int64) KReference {
	if x >= minSmallInt && x <= maxSmallInt {
		return createKrefSmallInt(x)
	}
	ref, obj := ms.newBigIntObject()
	obj.bigValue.SetInt64(x)
	return ref
}

// FromUint64 converts a uint64 to an integer in the model
func (ms *ModelState) FromUint64(x uint64) KReference {
	if x <= uint64(maxSmallInt) {
		return createKrefSmallInt(int64(x))
	}
	ref, obj := ms.newBigIntObject()
	obj.bigValue.SetUint64(x)
	return ref
}
