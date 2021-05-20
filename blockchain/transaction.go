package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
)

// Transaction struct
// No sensitive info should be added to this
type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

// SetID calculates and sets the
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer

	encoder := gob.NewEncoder(&encoded)
	err := encoder.Encode(tx)
	HandleErr(err)

	hash := sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

// TxInput is the transaction input
type TxInput struct {
	ID  []byte // transaction ID
	Out int    // index of the output
	Sig string // digital signature
}

// TxOutput is the transaction output
type TxOutput struct {
	Value  int    // value in tokens
	PubKey string // public key
}

// CoinBaseTx is the first transaction in the block
func CoinBaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", to)
	}

	txin := TxInput{
		ID:  []byte{},
		Out: -1,
		Sig: data,
	}

	txout := TxOutput{
		Value:  100,
		PubKey: to,
	}

	tx := Transaction{
		Inputs:  []TxInput{txin},
		Outputs: []TxOutput{txout},
	}
	tx.SetID()

	return &tx
}

func (tx *Transaction) isCoinbase() bool {
	return len(tx.Inputs) == 1 &&
		len(tx.Inputs[0].ID) == 0 &&
		tx.Inputs[0].Out == -1
}

// CanUnlock means the account(data) owns the information referenced by the input
func (txin *TxInput) CanUnlock(data string) bool {
	return txin.Sig == data
}

// CanUnlock means the account(data) owns the information referenced by the output
func (txout *TxOutput) CanUnlock(data string) bool {
	return txout.PubKey == data
}
