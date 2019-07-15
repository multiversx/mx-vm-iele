// File provided by the K Framework Go backend. Timestamp: 2019-07-15 14:09:18.513

package ieletestingmodel

import (
	"fmt"
	"strings"
)

// KObject defines a K item object that is managed by the model
type KObject interface {
	referenceType() kreferenceType
	equals(ms *ModelState, other KObject) bool
	deepCopy(ms *ModelState) KObject
	prettyPrint(ms *ModelState, sb *strings.Builder, indent int)
	kprint(ms *ModelState, sb *strings.Builder)
	collectionsToK(ms *ModelState) KReference
	increaseUsage(ms *ModelState)
	decreaseUsage(ms *ModelState)
	recycleUnused(ms *ModelState)
	preserve(ms *ModelState)
}

type objectReuseStatus int

const (
	active objectReuseStatus = iota
	inRecycleBin
	preserved
)

// ModelState holds the state of the executor at a certain moment
type ModelState struct {
	// allKs keeps all K sequence elements into one large structure
	// all K sequence element references point to this structure
	allKsElements []ksequenceElem

	// contains all KApply args, concatenated into one slice
	// KApply references contain the start position and arity,
	// so enough data to find their args in this slice
	allKApplyArgs []KReference

	// keeps big int objects, big int references point here
	bigInts []*bigInt

	// recycle bin for big ints
	// works as a stack
	bigIntRecycleBin []KReference

	// contains all bytes from types String, Bytes and KToken
	allBytes []byte

	// keeps object types mixed together
	allObjects []KObject

	// memoTables is a structure containing all memoization maps.
	// Memoization tables are implemented as maps of maps of maps of ...
	memoTables map[MemoTable]interface{}
}

// constantsModel is another instance of the model, but which only contains a few constants.
var constantsModel = NewModel()

func (ms *ModelState) getReferencedObject(index uint64, constant bool) KObject {
	if constant {
		return constantsModel.allObjects[index]
	}
	if index >= uint64(len(ms.allObjects)) {
		panic("trying to reference object beyond allocated objects")
	}
	return ms.allObjects[index]
}

func (ms *ModelState) addObject(obj KObject) KReference {
	newIndex := len(ms.allObjects)
	ms.allObjects = append(ms.allObjects, obj)
	return createKrefBasic(obj.referenceType(), false, uint64(newIndex))
}

func addConstantObject(obj KObject) KReference {
	newIndex := len(constantsModel.allObjects)
	constantsModel.allObjects = append(constantsModel.allObjects, obj)
	return createKrefBasic(obj.referenceType(), false, uint64(newIndex))
}

// NewModel creates a new blank model.
func NewModel() *ModelState {
	ms := &ModelState{}
	ms.allKsElements = make([]ksequenceElem, 0, 100000)
	ms.allKApplyArgs = make([]KReference, 0, 1000000)
	ms.allBytes = make([]byte, 0, 1<<20)
	ms.allObjects = make([]KObject, 0, 10000)
	ms.memoTables = nil
	return ms
}

// Clear resets the model as if it were new,
// but does not free the memory allocated by previous execution.
func (ms *ModelState) Clear() {
	ms.allKsElements = ms.allKsElements[:0]
	ms.allKApplyArgs = ms.allKApplyArgs[:0]
	ms.allObjects = ms.allObjects[:0]
	ms.allBytes = ms.allBytes[:0]
	ms.recycleAllInts()
	ms.memoTables = nil
}

// PrintStats simply prints some statistics to the console.
// Useful for checking the size of the model data.
func (ms *ModelState) PrintStats() {
	fmt.Printf("Nr. BigInt objects: %d\n", len(ms.bigInts))
	fmt.Printf("Nr. K sequence elements: %d\n", len(ms.allKsElements))
	fmt.Printf("Nr. KApply args: %d\n", len(ms.allKApplyArgs))
	fmt.Printf("Nr. bytes (strings, byte arrays, KTokens): %d\n", len(ms.allBytes))
	fmt.Printf("Nr. other objects: %d\n", len(ms.allObjects))
	fmt.Printf("Recycle bin\n")
	fmt.Printf("     BigInt    %d\n", len(ms.bigIntRecycleBin))
}
