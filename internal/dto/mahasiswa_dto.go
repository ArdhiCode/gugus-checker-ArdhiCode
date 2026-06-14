package dto

type GetByNRPRequest struct {
	NRP string `json:"nrp" binding:"required"`
}

type GetByNRPResponse struct {
	Name   string `json:"name"`
	NRP    string `json:"nrp"`
	Gugus  string `json:"gugus"`
	Region string `json:"region"`
}
