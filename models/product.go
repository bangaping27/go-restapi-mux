package models

type Product struct {
	ID    int64   `gorm:"primary_key;auto_increment" json:"id"`
	Nama  string  `gorm:"type:varchar(255)" json:"nama"`
	Stok  int32   `gorm:"type:int(5)" json:"stok"`
	Harga float64 `gorm:"type:decimal(14,2);not null" json:"harga"`
}
