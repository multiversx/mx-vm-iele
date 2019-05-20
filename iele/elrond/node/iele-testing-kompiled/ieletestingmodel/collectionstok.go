// File provided by the K Framework Go backend. Timestamp: 2019-05-20 22:38:10.632

package ieletestingmodel

import (
	"sort"
)

// CollectionsToK ... convert all collections to standard K items, like KApply, KToken, etc.
func CollectionsToK(k K) K {
	return k.collectionsToK()
}

func (k *Map) collectionsToK() K {
	if len(k.Data) == 0 {
		return &KApply{Label: UnitFor(k.Label), List: nil}
	}

	// sort entries
	orderedKVPairs := k.ToOrderedKeyValuePairs()

	// process
	elemLabel := ElementFor(k.Label)
	var result K
	for i, pair := range orderedKVPairs {
		elemK := &KApply{Label: elemLabel, List: []K{pair.Key, pair.Value.collectionsToK()}}
		if i == 0 {
			result = elemK
		} else {
			newResult := &KApply{Label: k.Label, List: []K{result, elemK}}
			result = newResult
		}
	}

	return result
}

func (k *List) collectionsToK() K {
	if len(k.Data) == 0 {
		return &KApply{Label: UnitFor(k.Label), List: nil}
	}

	elemLabel := ElementFor(k.Label)
	var result K
	for i, elem := range k.Data {
		elemK := &KApply{Label: elemLabel, List: []K{elem}}
		if i == 0 {
			result = elemK
		} else {
			newResult := &KApply{Label: k.Label, List: []K{result, elemK}}
			result = newResult
		}
	}

	return result
}

func (k *Set) collectionsToK() K {
	if len(k.Data) == 0 {
		return &KApply{Label: UnitFor(k.Label), List: nil}
	}

	// sort keys
	var keysAsString []string
	keyStrToK := make(map[string]K)
	for key := range k.Data {
		keyAsStr := key.String()
		keysAsString = append(keysAsString, keyAsStr)
		kkey, err := key.ToKItem()
		if err != nil {
			panic(err)
		}
		keyStrToK[keyAsStr] = kkey
	}
	sort.Strings(keysAsString)

	// process
	elemLabel := ElementFor(k.Label)
	var result K
	for i, keyAsStr := range keysAsString {
		key := keyStrToK[keyAsStr]
		elemK := &KApply{Label: elemLabel, List: []K{key}}
		if i == 0 {
			result = elemK
		} else {
			newResult := &KApply{Label: k.Label, List: []K{result, elemK}}
			result = newResult
		}
	}

	return result
}

func (k *Array) collectionsToK() K {
	result := &KApply{Label: UnitForArray(k.Sort), List: []K{
		NewString("uid"),
		NewIntFromUint64(k.Data.MaxSize),
	}}

	slice := k.Data.ToSlice()
	for i, item := range slice {
		newResult := &KApply{Label: ElementForArray(k.Sort), List: []K{
			result,
			NewIntFromInt(i),
			item,
		}}
		result = newResult
	}

	return result

}

func (k *KApply) collectionsToK() K {
	newList := make([]K, len(k.List))
	for i, child := range k.List {
		newList[i] = child.collectionsToK()
	}
	return &KApply{Label: k.Label, List: newList}
}

func (k *KSequence) collectionsToK() K {
	newKs := make([]K, len(k.Ks))
	for i, child := range k.Ks {
		newKs[i] = child.collectionsToK()
	}
	return &KSequence{Ks: newKs}
}

func (k *InjectedKLabel) collectionsToK() K {
	return k
}

func (k *KToken) collectionsToK() K {
	return k
}

func (k *KVariable) collectionsToK() K {
	return k
}

func (k *Int) collectionsToK() K {
	return k
}

func (k *MInt) collectionsToK() K {
	return k
}

func (k *Float) collectionsToK() K {
	return k
}

func (k *String) collectionsToK() K {
	return k
}

func (k *StringBuffer) collectionsToK() K {
	return k
}

func (k *Bytes) collectionsToK() K {
	return k
}

func (k *Bool) collectionsToK() K {
	return k
}

func (k *Bottom) collectionsToK() K {
	return k
}
