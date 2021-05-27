package types

// WalletType types of wallets
type WalletType uint8

const (
	// WalletTypeStandard Standard Wallet
	WalletTypeStandard = WalletType(0)

	// WalletTypeRateLimited Rate Limited Wallet
	WalletTypeRateLimited = WalletType(1)

	// WalletTypeAtomicSwap Atomic Swap
	WalletTypeAtomicSwap = WalletType(2)

	// WalletTypeAuthorizedPayee Authorized Payee
	WalletTypeAuthorizedPayee = WalletType(3)

	// WalletTypeMultiSig Multi Sig
	WalletTypeMultiSig = WalletType(4)

	// WalletTypeCustody Custody
	WalletTypeCustody = WalletType(5)

	// WalletTypeColouredCoin Coloured Coin
	WalletTypeColouredCoin = WalletType(6)

	// WalletTypeRecoverable Recoverable Wallet
	WalletTypeRecoverable = WalletType(7)

	// WalletTypeDistributedID DID Wallet
	WalletTypeDistributedID = WalletType(8)
)

// Wallet single wallet record
type Wallet struct {
	Data string      `json:"data"`
	ID   uint32      `json:"id"`
	Name string      `json:"name"`
	Type *WalletType `json:"type"`
}

// WalletBalance specific wallet balance information
type WalletBalance struct {
	ConfirmedWalletBalance   int64 `json:"confirmed_wallet_balance"` // @TODO uint128
	MaxSendAmount            int64 `json:"max_send_amount"`
	PendingChange            int64 `json:"pending_change"`
	SpendableBalance         int64 `json:"spendable_balance"`          // @TODO uint128
	UnconfirmedWalletBalance int64 `json:"unconfirmed_wallet_balance"` // @TODO uint128
	UnspentCoinCount         int64 `json:"unspent_coin_count"`
	PendingCoinRemovalCount  int64 `json:"pending_coin_removal_count"`
	WalletID                 int32 `json:"wallet_id"`
}
