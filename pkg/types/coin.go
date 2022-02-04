package types

import (
	"fmt"
	"strconv"
)

const (
	mojoInChia int64 = 1000000000000
)

// Mojo is a special type for Mojos, to keep track of what unit an amount is
type Mojo int64

func (m Mojo) MarshalJSON() ([]byte, error) {
	s := strconv.FormatInt(int64(m), 10)
	return []byte(s), nil
}

func (m *Mojo) UnmarshalJSON(data []byte) error {
	mojo, err := strconv.ParseUint(string(data), 10, 64)
	if err != nil {
		return err
	}
	*m = Mojo(mojo)
	return nil
}

// XCH is a special type for Chia, to keep track of what unit an amount is
type XCH float64

func (xch XCH) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("%f", xch)
	return []byte(s), nil
}

func (xch *XCH) UnmarshalJSON(data []byte) error {
	x, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return err
	}
	*xch = XCH(x)
	return nil
}

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
