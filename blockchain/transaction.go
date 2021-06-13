package blockchain

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"go-blockchain/errors"
	"go-blockchain/wallet"
	"log"
	"math/big"
	"strings"
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
		ID:        []byte{},
		Out:       -1,
		Signature: nil,
		PubKey:    []byte(data),
	}

	txout := NewTXOutput(100, to)

	tx := Transaction{
		Inputs:  []TxInput{txin},
		Outputs: []TxOutput{*txout},
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

	wallets, err := wallet.CreateWallets()
	errors.HandleErr(err)
	w := wallets.GetWallet(from)
	pubKeyHash := wallet.PublicKeyHash(w.PublicKey)

	acc, validOutputs := chain.FindSpendableOutputs(pubKeyHash, amount)

	if acc < amount {
		log.Panic("Error, not enough funds... :(")
	}

	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		errors.HandleErr(err)

		for _, out := range outs {
			input := TxInput{
				ID:        txID,
				Out:       out,
				Signature: nil,
				PubKey:    w.PublicKey,
			}
			inputs = append(inputs, input)
		}
	}

	output := *NewTXOutput(amount, to)
	outputs = append(outputs, output)

	if acc > amount {
		output = *NewTXOutput(acc-amount, from)
		outputs = append(outputs, output)
	}

	tx = &Transaction{
		Inputs:  inputs,
		Outputs: outputs,
	}

	tx.ID = tx.Hash()
	chain.SignTransaction(tx, w.PrivateKey)

	return tx
}

// TrimmedCopy returns a copy of the transaction without the signature and the public key
func (tx *Transaction) TrimmedCopy() Transaction {
	var inputs []TxInput
	var outputs []TxOutput

	for _, input := range tx.Inputs {
		inputCopy := TxInput{
			ID:        input.ID,
			Out:       input.Out,
			Signature: nil,
			PubKey:    nil,
		}
		inputs = append(inputs, inputCopy)
	}

	for _, output := range tx.Outputs {
		outputCopy := TxOutput{
			Value:      output.Value,
			PubKeyHash: output.PubKeyHash,
		}
		outputs = append(outputs, outputCopy)
	}

	txCopy := Transaction{
		ID:      tx.ID,
		Inputs:  inputs,
		Outputs: outputs,
	}

	return txCopy
}

// Sign signs the transaction with the passed in private key
func (tx *Transaction) Sign(privKey ecdsa.PrivateKey, prevTXs map[string]Transaction) {
	if tx.isCoinbase() {
		return
	}

	for _, input := range tx.Inputs {
		prevTx := prevTXs[hex.EncodeToString(input.ID)]
		if prevTx.ID == nil {
			log.Panic("ERROR: Previous transaction does not exist")
		}
	}

	txCopy := tx.TrimmedCopy()

	for inputID, input := range txCopy.Inputs {
		prevTx := prevTXs[hex.EncodeToString(input.ID)]
		txCopy.Inputs[inputID].Signature = nil
		txCopy.Inputs[inputID].PubKey = prevTx.Outputs[input.Out].PubKeyHash
		txCopy.ID = txCopy.Hash()
		txCopy.Inputs[inputID].PubKey = nil

		r, s, err := ecdsa.Sign(rand.Reader, &privKey, txCopy.ID)
		errors.HandleErr(err)

		signature := append(r.Bytes(), s.Bytes()...)

		tx.Inputs[inputID].Signature = signature
	}
}

// Verify verifies the transaction
func (tx *Transaction) Verify(prevTXs map[string]Transaction) bool {
	if tx.isCoinbase() {
		return true
	}

	for _, input := range tx.Inputs {
		prevTx := prevTXs[hex.EncodeToString(input.ID)]
		if prevTx.ID == nil {
			log.Panic("ERROR: Previous transaction does not exist")
		}
	}

	txCopy := tx.TrimmedCopy()
	curve := elliptic.P256()

	for inputID, input := range tx.Inputs {
		prevTx := prevTXs[hex.EncodeToString(input.ID)]
		txCopy.Inputs[inputID].Signature = nil
		txCopy.Inputs[inputID].PubKey = prevTx.Outputs[input.Out].PubKeyHash
		txCopy.ID = txCopy.Hash()
		txCopy.Inputs[inputID].PubKey = nil

		r := big.Int{}
		s := big.Int{}
		signatureLen := len(input.Signature)
		r.SetBytes(input.Signature[:(signatureLen / 2)])
		s.SetBytes(input.Signature[(signatureLen / 2):])

		x := big.Int{}
		y := big.Int{}
		keyLen := len(input.PubKey)
		x.SetBytes(input.PubKey[:(keyLen / 2)])
		y.SetBytes(input.PubKey[(keyLen / 2):])

		rawPubKey := ecdsa.PublicKey{
			Curve: curve,
			X:     &x,
			Y:     &y,
		}
		if ecdsa.Verify(&rawPubKey, txCopy.ID, &r, &s) == false {
			return false
		}
	}

	return true
}

func (tx Transaction) String() string {
	var lines []string

	lines = append(lines, fmt.Sprintf("Transaction %x:", tx.ID))

	for i, input := range tx.Inputs {
		lines = append(lines, fmt.Sprintf("\tInput %d:", i))
		lines = append(lines, fmt.Sprintf("\tTXID:     %x", input.ID))
		lines = append(lines, fmt.Sprintf("\tOut:       %d", input.Out))
		lines = append(lines, fmt.Sprintf("\tSignature: %x", input.Signature))
		lines = append(lines, fmt.Sprintf("\tPubKey:    %x", input.PubKey))
	}

	for i, output := range tx.Outputs {
		lines = append(lines, fmt.Sprintf("\tOutput %d:", i))
		lines = append(lines, fmt.Sprintf("\tValue:  %d", output.Value))
		lines = append(lines, fmt.Sprintf("\tScript: %x", output.PubKeyHash))
	}

	return strings.Join(lines, "\n")
}
