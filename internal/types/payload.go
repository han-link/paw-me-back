package types

import "github.com/google/uuid"

type AddMembersPayload struct {
	UserIDs []uuid.UUID `json:"user_ids" validate:"required,dive,uuid4"`
} // @name AddMembersPayload

type CreateGroupPayload struct {
	// Name of the group
	Name string `json:"name" validate:"required,max=255"`

	// Optional list of member IDs
	// Example: ["d5075280-2f7c-4967-8526-aaaca282de36"]
	Members []uuid.UUID `json:"members,omitempty" validate:"dive,uuid4"`
} // @name CreateGroupPayload
