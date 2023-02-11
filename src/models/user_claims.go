package models

import "github.com/online.scheduling-api/constants"

func MapUserClaimsFrom(u *User) map[string]any {
	return map[string]any{
		constants.CLAIM_USER_ID:   &u.Id,
		constants.CLAIM_USER_NAME: &u.Name,
		constants.CLAIM_USER_ROLE: &u.Role,
	}
}
