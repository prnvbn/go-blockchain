package errors

import (
	"fmt"
	"log"
)

// HandleErr panics and logs the error
func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

type errorType int

const (
	transactionNotFoundErr = iota + 1
	invalidAddressErr
)

var errorTypes = []string{"TransactionNotFoundError", "InvalidAddressError"}

func (e errorType) String() string {
	return red(errorTypes[e-1])
}

func newError(errType errorType, template string, args ...interface{}) error {
	msg := fmt.Sprintf(template, args...)
	err := fmt.Errorf("%s: %s\t👎", errType.String(), msg)
	return err
}

func red(s string) string {
	return "\u001b[1m\u001b[31m" + s + "\u001b[0m"
}

/************************************ FACTORY METHODS ************************************/

// NewTransactionNotFoundError returns
// TransactionNotFoundError: No transaction found with ID: ID
func NewTransactionNotFoundError(ID []byte) error {
	return newError(transactionNotFoundErr, "No transaction found with ID %s", ID)
}

// NewInvalidAddressError returns
// InvalidAddressError: ADDRESS is not a valid address
func NewInvalidAddressError(address string) error {
	return newError(invalidAddressErr, "%s is not a valid address", address)
}
