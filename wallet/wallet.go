package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"go-blockchain/errors"

	"golang.org/x/crypto/ripemd160"
)

const (
	checksumLength = 4
	version        = byte(0x00)
)

// Wallet is a wallet
type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
}

func (w Wallet) Address() []byte {
	pubHash := PublicKeyHash(w.PublicKey)
	versionedHash := append([]byte{version}, pubHash...)
	checksum := Checksum(versionedHash)
	fullHash := append(versionedHash, checksum...)
	address := Base58Encode(fullHash)

	fmt.Printf("pub key: %x\n", w.PublicKey)
	fmt.Printf("pub hash: %x\n", pubHash)
	fmt.Printf("address: %x\n", address)
	return address
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

// PublicKeyHash returns the publickeyhash of the p
func PublicKeyHash(pubKey []byte) []byte {
	pubHashed := sha256.Sum256(pubKey)

	hasher := ripemd160.New()
	_, err := hasher.Write(pubHashed[:])
	errors.HandleErr(err)

	return hasher.Sum(nil)
}

// Checksum return the first 4 bytes of the payload after it has been hashed twice
func Checksum(payload []byte) (checksum []byte) {
	firstHash := sha256.Sum256(payload)
	secondHash := sha256.Sum256(firstHash[:])

	checksum = secondHash[:checksumLength]
	return checksum
}
