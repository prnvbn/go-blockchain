package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"go-blockchain/errors"
)

// Wallet is a wallet
type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

// NewKeyPair generate a pair of of public and private key
// can generate upto 10 ^ 77 different keys
func NewKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()

	private, err := ecdsa.GenerateKey(curve, rand.Reader)
	errors.HandleErr(err)

	public := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, public
}

// CreateWallet creates a wallet with a new key pair
func CreateWallet() *Wallet {
	private, public := NewKeyPair()
	return &Wallet{
		PrivateKey: private,
		PublicKey:  public,
	}
}
