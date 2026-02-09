package main

type ValidationRequest struct {
	NIK string `json:"nik" binding:"required"`
	KK  string `json:"kk" binding:"required"`
}

type ValidationResponse struct {
	IsValid  bool   `json:"isValid"`
	Message  string `json:"message"`
	FullName string `json:"fullName,omitempty"`
}
