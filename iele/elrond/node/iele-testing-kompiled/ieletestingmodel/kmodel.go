// File provided by the K Framework Go backend. Timestamp: 2019-05-20 22:38:10.632

package ieletestingmodel

import (
	"math/big"
	"strings"
)

// K ... Defines a K entity
type K interface {
	Equals(K) bool
	DeepCopy() K
	structPrint(sb *strings.Builder, indent int)
	prettyPrint(sb *strings.Builder, indent int)
	kprint(sb *strings.Builder)
	collectionsToK() K
}

// KSequence ... a sequence of K items
type KSequence struct {
	Ks []K
}

// EmptyKSequence ... the KSequence with no elements
var EmptyKSequence = &KSequence{Ks: nil}

// KItem ...
type KItem interface {
}

// KApply ... a type of KItem, TODO: document
type KApply struct {
	Label KLabel
	List  []K
}

// InjectedKLabel ... a type of KItem, TODO: document
type InjectedKLabel struct {
	Label KLabel
}

// KToken ... a type of KItem, TODO: document
type KToken struct {
	Value string
	Sort  Sort
}

// KVariable ... a type of KItem, TODO: document
type KVariable struct {
	Name string
}

// Map ... a type of KItem, TODO: document
type Map struct {
	Sort  Sort
	Label KLabel
	Data  map[KMapKey]K
}

// Set ... a type of KItem, TODO: document
type Set struct {
	Sort  Sort
	Label KLabel
	Data  map[KMapKey]bool
}

// List ... a type of KItem, TODO: document
type List struct {
	Sort  Sort
	Label KLabel
	Data  []K
}

// Array ... array of K Items of fixed size
type Array struct {
	Sort Sort
	Data *DynamicArray
}

// Int ... integer type, implemented via a big int
type Int struct {
	Value *big.Int
}

// MInt ... machine integer
type MInt struct {
	Value int32
}

// Float ... float type
type Float struct {
	Value float32
}

// String ... string type
type String struct {
	Value string
}

// StringBuffer ... a string builder, in which strings can be appended
type StringBuffer struct {
	Value strings.Builder
}

// Bytes ... a type of KItem, TODO: document
type Bytes struct {
	Value []byte
}

// Bool ... K boolean value
type Bool struct {
	Value bool
}

// Bottom ... a K item that contains no data
type Bottom struct {
}
