package types

const (
	mojoInChia int64 = 1000000000000
)

// Mojo is a special type for Mojos, to keep track of what unit an amount is
type Mojo int64

// XCH is a special type for Chia, to keep track of what unit an amount is
type XCH float64

// ToMojo converts chia to mojos
func (c XCH) ToMojo() Mojo {
	return Mojo(c * XCH(mojoInChia))
}

// ToChia converts mojo to chia
func (m Mojo) ToChia() XCH {
	return XCH(m) / XCH(mojoInChia)
}

// Coin is a coin
type Coin struct {
	Amount         Mojo   `json:"amount"`
	ParentCoinInfo string `json:"parent_coin_info"`
	PuzzleHash     string `json:"puzzle_hash"`
}

// CoinSolution solution to a coin
type CoinSolution struct {
	Coin         *Coin              `json:"coin"`
	PuzzleReveal *SerializedProgram `json:"puzzle_reveal"`
	Solution     *SerializedProgram `json:"solution"`
}
