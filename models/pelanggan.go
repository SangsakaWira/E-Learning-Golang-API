package model

type Pelanggan struct {
	ID     string `json:"id"`
	Nama   string `json:"nama"`
	Alamat string `json:"alamat"`
	Email  string `json:"email"`
	Dompet string `json:"dompet"`
}
