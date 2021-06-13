package blockchain

import (
	"bytes"
	"go-blockchain/wallet"
)

// TxInput is the transaction input
type TxInput struct {
	ID        []byte // transaction ID
	Out       int    // index of the output
	Signature []byte // digital signature
	PubKey    []byte // unhashed public key
}

// TxOutput is the transaction output
type TxOutput struct {
	Value      int    // value in tokens
	PubKeyHash []byte // hashed public key
}

// NewTXOutput creates a new transaction output and locks it
func NewTXOutput(value int, address string) *TxOutput {
	txo := &TxOutput{Value: value, PubKeyHash: nil}
	txo.Lock([]byte(address))

	return txo
}

// UsesKey checks if the transaction input uses the passed in public key hash passed in to lock it
func (in *TxInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := wallet.PublicKeyHash(in.PubKey)

	return bytes.Compare(pubKeyHash, lockingHash) == 0
}

// Lock will lock the output ensuing that the output can only be unlocked by the passed in address
func (out *TxOutput) Lock(address []byte) {
	decodedAddress := wallet.Base58Decode(address)
	pubKeyHash := decodedAddress[1 : len(decodedAddress)-wallet.ChecksumLength]
	out.PubKeyHash = pubKeyHash
}

// IsLockedWithKey checks if the output is locked with the passed in public key hash
func (out *TxOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

// // CanUnlock means the account(data) owns the information referenced by the input
// func (txin *TxInput) CanUnlock(data string) bool {
// 	return txin.Sig == data
// }

// // CanBeUnlockedBy means the account(data) owns the information referenced by the output
// func (txout *TxOutput) CanBeUnlockedBy(data string) bool {
// 	return txout.PubKey == data
// }
