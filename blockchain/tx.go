package blockchain

// TxInput is the transaction input
type TxInput struct {
	ID  []byte // transaction ID
	Out int    // index of the output
	Sig string // digital signature
}

// CanUnlock means the account(data) owns the information referenced by the input
func (txin *TxInput) CanUnlock(data string) bool {
	return txin.Sig == data
}

// TxOutput is the transaction output
type TxOutput struct {
	Value  int    // value in tokens
	PubKey string // public key
}

// CanBeUnlockedBy means the account(data) owns the information referenced by the output
func (txout *TxOutput) CanBeUnlockedBy(data string) bool {
	return txout.PubKey == data
}
