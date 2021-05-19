package blockchain

// Proof Of Work Algorithm

// 1. Take data from the block
// 2. Create a counter (nonce) starting at 0
// 3. Create a hash of the data plaus the nonce
// 4. Check the hash to see if it meets the requirements

// Requirements:
// i) First few bytes must contain 0s

const Difficulty = 12

