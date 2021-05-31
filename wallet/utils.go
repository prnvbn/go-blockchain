package wallet

import (
	"go-blockchain/errors"

	"github.com/mr-tron/base58"
)

// Base58Encode encodes the input
// 0 O l I + / are missing from the encoding algorithm
func Base58Encode(input []byte) []byte {
	encode := base58.Encode(input)

	return []byte(encode)
}

// Base58Decode decodes the input
func Base58Decode(input []byte) []byte {
	decode, err := base58.Decode(string(input))
	errors.HandleErr(err)

	return decode
}
