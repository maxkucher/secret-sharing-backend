package public

type CreateSecretDTO struct {
	PlainString string `json:"plain_string"`
}

type CreateSecretDTOResponse struct {
	Id string `json:"id"`
}

type GetSecretResponse struct {
	Data string `json:"data"`
}
