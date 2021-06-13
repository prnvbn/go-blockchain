package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"go-blockchain/errors"
	"log"
)

// Transaction struct
// No sensitive info should be added to this
type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

// Serialize serializes the transaction struct into bytes
func (tx Transaction) Serialize() []byte {
	var encoded bytes.Buffer

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	errors.HandleErr(err)

	return encoded.Bytes()
}

// Hash hashes the transaction struct into bytes
func (tx *Transaction) Hash() []byte {
	var hash [32]byte

	txCopy := *tx
	txCopy.ID = []byte{}

	hash = sha256.Sum256(txCopy.Serialize())

	return hash[:]
}

// SetID calculates and sets the
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer

	encoder := gob.NewEncoder(&encoded)
	err := encoder.Encode(tx)
	errors.HandleErr(err)

	hash := sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
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

// NewTransaction creates and returns a new transaction
func NewTransaction(from, to string, amount int, chain *BlockChain) (tx *Transaction) {
	var inputs []TxInput
	var outputs []TxOutput

	acc, validOutputs := chain.FindSpendableOutputs(from, amount)

	if acc < amount {
		log.Panic("Error, not enough funds... :(")
	}

	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		errors.HandleErr(err)

		for _, out := range outs {
			input := TxInput{
				ID:  txID,
				Out: out,
				Sig: from,
			}
			inputs = append(inputs, input)
		}
	}

	output := TxOutput{
		Value:  amount,
		PubKey: to,
	}
	outputs = append(outputs, output)

	if acc > amount {
		output = TxOutput{
			Value:  acc - amount,
			PubKey: from,
		}
		outputs = append(outputs, output)
	}

	tx = &Transaction{
		Inputs:  inputs,
		Outputs: outputs,
	}
	tx.SetID()

	return tx
}
