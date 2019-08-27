// File provided by the K Framework Go backend. Timestamp: 2019-08-27 09:22:42.803

package ieletestingmodel

import (
	"errors"
	"fmt"
	"strconv"
)

// KMapKey is a compact representation of a K item to be used as key in a map.
type KMapKey interface {
	toKItem(ms *ModelState) (KReference, error)
	String() string
}

// MapKey converts a K item to a map key, if possible
func (ms *ModelState) MapKey(ref KReference) (KMapKey, bool) {
	if b, t := CastToBool(ref); t {
		return kmapKeyBasic{typeName: "Bool", value: fmt.Sprintf("%t", b)}, true
	}
	if IsBottom(ref) {
		return kmapBottom{}, true
	}
	if str, t := ms.GetString(ref); t {
		return kmapKeyBasic{typeName: "String", value: str}, true
	}
	if iStr, t := ms.GetIntToString(ref, 16); t {
		return kmapKeyBasic{typeName: "Int", value: iStr}, true
	}
	if ktoken, t := ms.GetKTokenObject(ref); t {
		return ktoken, true
	}
	if isKApply, dataRef, label, arity, index := parseKrefKApply(ref); isKApply {
		switch arity {
		case 0:
			return kmapKeyKApply0{label: KLabel(label)}, true
		case 1:
			argAsKey, argOk := ms.MapKey(ms.getData(dataRef).allKApplyArgs[index+0]) // first argument
			if !argOk {
				return kmapBottom{}, false
			}
			return kmapKeyKApply1{label: KLabel(label), arg1: argAsKey}, true
		default:
			return kmapBottom{}, false
		}
	}

	return kmapBottom{}, false
}

// ToKItem converts a map key back to a regular K item
func (ms *ModelState) ToKItem(mapKey KMapKey) (KReference, error) {
	return mapKey.toKItem(ms)
}

// kmapKeyBasic ... representation of basic types: Int, String, Bool
type kmapKeyBasic struct {
	typeName string
	value    string
}

type kmapKeyKApply0 struct {
	label KLabel
}

type kmapKeyKApply1 struct {
	label KLabel
	arg1  KMapKey
}

type kmapBottom struct {
}

func (mapKey kmapKeyBasic) toKItem(ms *ModelState) (KReference, error) {
	switch mapKey.typeName {
	case "Int":
		return ms.ParseIntFromBase(mapKey.value, 16)
	case "Bool":
		b, err := strconv.ParseBool(mapKey.value)
		if err != nil {
			return NoResult, err
		}
		return ToKBool(b), nil
	case "String":
		return ms.NewString(mapKey.value), nil
	default:
		return NoResult, errors.New("unable to convert KMapKey to K. Unknown type")
	}

}

func (k KToken) toKItem(ms *ModelState) (KReference, error) {
	return ms.NewKToken(k.Sort, k.Value), nil
}

func (mapKey kmapKeyKApply0) toKItem(ms *ModelState) (KReference, error) {
	return ms.NewKApply(mapKey.label), nil
}

func (mapKey kmapKeyKApply1) toKItem(ms *ModelState) (KReference, error) {
	argKItem, err := mapKey.arg1.toKItem(ms)
	if err != nil {
		return NoResult, err
	}
	return ms.NewKApply(mapKey.label, argKItem), nil
}

func (mapKey kmapBottom) toKItem(ms *ModelState) (KReference, error) {
	return InternedBottom, nil
}

// String ... string representation of the key
func (mapKey kmapKeyBasic) String() string {
	return fmt.Sprintf("%s_%s", mapKey.typeName, mapKey.value)
}

// String yields string representation of the key
func (k KToken) String() string {
	return fmt.Sprintf("KToken(%s)_%s", k.Sort.Name(), k.Value)
}

// String yields string representation of the key
func (mapKey kmapKeyKApply0) String() string {
	return fmt.Sprintf("KApply(%s)", mapKey.label.Name())
}

// String yields string representation of the key
func (mapKey kmapKeyKApply1) String() string {
	return fmt.Sprintf("KApply(%s)_%s", mapKey.label.Name(), mapKey.arg1.String())
}

// String yields string representation of the key
func (mapKey kmapBottom) String() string {
	return "Bottom"
}
