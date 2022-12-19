package models

type Stats struct {
	Symbol      string  `json:"symbol"`
	FloorPrice  float64 `json:"floorPrice"`
	ListedCount float64 `json:"listedCount"`
	VolumeAll   float64 `json:"volumeAll"`
}
