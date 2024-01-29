package dto

//get important data without getting all the data
type LocationDTO struct {
	Status bool `json:"status,omitempty"`
}
