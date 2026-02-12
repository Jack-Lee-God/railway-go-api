package models

type Bioskop struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Nama   string  `json:"nama"`
	Lokasi string  `json:"lokasi"`
	Rating float32 `json:"rating"`
}
