package endpoint

import (
	"errors"
	"fmt"

	vmi "github.com/ElrondNetwork/elrond-vm/iele/vm-interface"
)

func validateInput(input *vmi.VMInput) error {
	if input.BlockHeader == nil {
		return errors.New("block header required")
	}

	var validationErr error
	validationErr = validateFrom(input.CallerAddr)
	if validationErr != nil {
		return validationErr
	}

	validationErr = validateTo(input.RecipientAddr)
	if validationErr != nil {
		return validationErr
	}

	return nil
}

func validateFrom(address []byte) error {
	if len(address) == addressLength {
		return nil
	}
	return fmt.Errorf("caller address is not %d bytes in length", addressLength)
}

func validateTo(address []byte) error {
	if len(address) == addressLength {
		return nil
	}

	if len(address) == 0 {
		return nil
	}

	if len(address) == 1 && address[0] == 0 {
		return nil
	}

	return fmt.Errorf("recipient address is not %d bytes in length", addressLength)
}
