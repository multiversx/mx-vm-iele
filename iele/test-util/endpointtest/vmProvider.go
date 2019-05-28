package endpointtest

import (
	vmi "github.com/ElrondNetwork/elrond-vm-common"
)

// VMProvider is an interface used by the tests to get the right VM version for the required schedule
type VMProvider interface {
	GetVM(scheduleName string) (vmi.VMExecutionHandler, error)
}
