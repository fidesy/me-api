package models

type Activity struct {
	Signature      string  `json:"signature"`
	Type           string  `json:"type"`
	Source         string  `json:"source"`
	TokenMint      string  `json:"tokenMint"`
	Collection     string  `json:"collection"`
	Slot           int     `json:"slot"`
	BlockTime      int64   `json:"blockTime"`
	Buyer          string  `json:"buyer"`
	BuyerReferral  string  `json:"buyerReferral"`
	Seller         string  `json:"seller"`
	SellerReferral string  `json:"sellerReferral"`
	Price          float64 `json:"price"`
}
