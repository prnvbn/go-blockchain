package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
)

// Proof Of Work Algorithm

// 1. Take data from the block
// 2. Create a counter (nonce) starting at 0
// 3. Create a hash of the data plaus the nonce
// 4. Check the hash to see if it meets the requirements

// Requirements:
// i) First few bytes must contain 0s

// Difficulty affects the requirements
const Difficulty = 12

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
			pow.Block.Data,
			toHex(int64(nonce)),
			toHex(int64(Difficulty)),
		},
		[]byte{},
	)

	return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
	nonce := 0

	var hash [32]byte
	var intHash big.Int

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}

	fmt.Println()

	return nonce, hash[:]
}

// ====================== UTILITIES ======================

func toHex(no int64) []byte {
	buff := new(bytes.Buffer)

	err := binary.Write(buff, binary.BigEndian, no)
	if err != nil {
		panic(err)
	}

	return buff.Bytes()
}
