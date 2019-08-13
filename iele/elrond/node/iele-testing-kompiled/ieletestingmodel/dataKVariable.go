// File provided by the K Framework Go backend. Timestamp: 2019-08-13 18:10:37.856

package ieletestingmodel

// KVariable is a KObject representing a KVariable item in K
type KVariable struct {
	Name string
}

func (*KVariable) referenceType() kreferenceType {
	return kvariableRef
}

// NewKVariable creates a new object and returns the reference.
func (ms *ModelState) NewKVariable(name string) KReference {
	return ms.mainData.addObject(&KVariable{Name: name})
}
