package api

type AuthCreateOrUpdateRequest struct {
	Login      string `json:"login"`
	Passphrase string `json:"passphrase"`
}

func MapAuthRequestFrom(u *UserCreateOrUpdateRequest) *AuthCreateOrUpdateRequest {
	return &AuthCreateOrUpdateRequest{
		Login:      u.Login,
		Passphrase: u.Passphrase,
	}
}
