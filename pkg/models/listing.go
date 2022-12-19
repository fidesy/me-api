package models

type Listing struct {
	PdaAddress   string  `json:"pdaAddress"`
	AuctionHouse string  `json:"auctionHouse"`
	TokenAddress string  `json:"tokenAddress"`
	TokenMint    string  `json:"tokenMint"`
	Seller       string  `json:"seller"`
	TokenSize    int     `json:"tokenSize"`
	Price        float64 `json:"price"`
}
