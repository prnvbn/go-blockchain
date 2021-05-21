package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

// Proof Of Work Algorithm

// 1. Take data from the block
// 2. Create a counter (nonce) starting at 0
// 3. Create a hash of the data plus the nonce
// 4. Check the hash to see if it meets the requirements

// Requirements:
// i) First few bytes must contain 0s

// Difficulty affects the requirements
// Hence affecting the time taken to do the proofofwork work
const Difficulty = 18

// ProofOfWork structure
type ProofOfWork struct {
	Block  *Block
	Target *big.Int // represents the requirement(s)
}

// NewProof does the proof-of-work work
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	return &ProofOfWork{b, target}
}

// InitData initialises the data
func (pow ProofOfWork) InitData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.HashTransactions(),
			toHex(int64(nonce)),
			toHex(int64(Difficulty)),
		},
		[]byte{},
	)

	return data
}

// Run does the actual work for the proof
func (pow *ProofOfWork) Run() (int, []byte) {
	nonce := 0

	var hash [32]byte
	var intHash big.Int

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}

	fmt.Println()

	return nonce, hash[:]
}

// Validate validates the calculated hash
// computationally easier than the actual proof done in Run function
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce) //TODO: Law of demeter?

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])
	return intHash.Cmp(pow.Target) == -1
}

// ====================== UTILITIES ======================

func toHex(no int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, no)
	if err != nil {
		log.Panic(err)

	}

	return buff.Bytes()
}
