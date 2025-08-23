package types

import "github.com/google/uuid"

type AddMembersPayload struct {
	UserIDs []uuid.UUID `json:"user_ids" validate:"required,dive,uuid4"`
}
