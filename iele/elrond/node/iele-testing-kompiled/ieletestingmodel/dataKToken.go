// File provided by the K Framework Go backend. Timestamp: 2019-08-28 14:13:50.189

package ieletestingmodel

// KToken is a KObject representing a KToken item in K
type KToken struct {
	Value string
	Sort  Sort
}

// GetKTokenObject a struct containing KToken data, if possible.
func (ms *ModelState) GetKTokenObject(ref KReference) (KToken, bool) {
	isKToken, dataRef, sort, length, index := parseKrefKToken(ref)
	if !isKToken {
		return KToken{}, false
	}
	value := ""
	if length > 0 {
		value = string(ms.getData(dataRef).allBytes[index : index+length])
	}
	return KToken{
		Sort:  Sort(sort),
		Value: value,
	}, true
}

// KTokenValue yields the value of a KToken object.
func (ms *ModelState) KTokenValue(ref KReference) string {
	isKToken, dataRef, _, length, index := parseKrefKToken(ref)
	if !isKToken {
		panic("KTokenValue called for reference to item other than KToken")
	}
	if length == 0 {
		return ""
	}
	return string(ms.getData(dataRef).allBytes[index : index+length])
}

func (md *ModelData) newKToken(sortInt uint64, valueBytes []byte) KReference {
	length := uint64(len(valueBytes))
	if length == 0 {
		return createKrefKToken(noDataRef, sortInt, length, 0)
	}
	startIndex := uint64(len(md.allBytes))
	md.allBytes = append(md.allBytes, valueBytes...)
	return createKrefKToken(md.selfRef, sortInt, length, startIndex)
}

// NewKToken creates a new object and returns the reference.
func (ms *ModelState) NewKToken(sort Sort, value string) KReference {
	return ms.mainData.newKToken(uint64(sort), []byte(value))
}

// NewKTokenConstant creates a new KToken constant, which is saved statically.
// Do not use for anything other than constants, since these never get cleaned up.
func NewKTokenConstant(sort Sort, value string) KReference {
	return constantsData.newKToken(uint64(sort), []byte(value))
}
