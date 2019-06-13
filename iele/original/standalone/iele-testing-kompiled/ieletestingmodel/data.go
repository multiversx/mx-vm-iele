// File provided by the K Framework Go backend. Timestamp: 2019-06-13 13:34:21.386

package ieletestingmodel

// ModelState holds the state of the executor at a certain moment
type ModelState struct {
    initialized bool

    // allKs keeps all KSequences into one large 2d structure
    // all KSequences point to this structure
    // the first element of allKs should be empty sequence
    allKs [][]K
}

// Init prepares model for execution
func NewModel() *ModelState {
    ms := &ModelState{}
    ms.Init()
    return ms
}

// Init prepares model for execution
func (ms *ModelState) Init() {
    if ms.initialized {
        return
    }
    ms.initialized = true
    ms.allKs = [][]K{[]K{}}
}

// ClearModel ... clean up any data left from previous executions, to save memory
func (ms *ModelState) ClearModel() {
    ms.initialized = false
    ms.Init()
}