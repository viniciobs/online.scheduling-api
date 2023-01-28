package api

type ModalityCreateOrUpdateRequest struct {
	Name        string `json:"name"`
	Description string `string:"description"`
}
