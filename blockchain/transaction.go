package blockchain

// Transaction struct
// No sensitive info should be added to this
type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
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
