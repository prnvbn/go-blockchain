package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
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

// CanBeUnlockedBy means the account(data) owns the information referenced by the output
func (txout *TxOutput) CanBeUnlockedBy(data string) bool {
	return txout.PubKey == data
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
		HandleErr(err)

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
