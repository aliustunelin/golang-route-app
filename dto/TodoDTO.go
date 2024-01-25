package dto

//tüm  verileri almadan bazılaırnı almamızı sağlayan
type TodoDTO struct {
	Status bool `json:"status,omitempty"`
}
