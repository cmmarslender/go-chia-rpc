package types

// Coin is a coin
type Coin struct {
	Amount         uint64 `json:"amount"`
	ParentCoinInfo string `json:"parent_coin_info"`
	PuzzleHash     string `json:"puzzle_hash"`
}

// CoinSolution solution to a coin
type CoinSolution struct {
	Coin         *Coin              `json:"coin"`
	PuzzleReveal *SerializedProgram `json:"puzzle_reveal"`
	Solution     *SerializedProgram `json:"solution"`
}
