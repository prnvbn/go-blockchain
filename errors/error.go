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

var errorTypes = []string{"TransactionNotFoundError"}

const (
	transactionNotFoundErr = iota + 1
)

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

// NewTransactionNotFound returns
// TransactionNotFoundError: No transaction found with ID: ID
func NewTransactionNotFound(ID []byte) error {
	return newError(transactionNotFoundErr, "No transaction found with ID %s", ID)
}
