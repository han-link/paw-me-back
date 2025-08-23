package types

import (
	"time"

	"github.com/google/uuid"
)

type Response[T any] struct {
	Data    T      `json:"data"`
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
} // @name Response

type Base struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserBrief struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
} // @name UserBrief

type Group struct {
	Base
	Name  string    `json:"name"`
	Owner UserBrief `json:"owner"`
} // @name Group

type GroupWithMembers struct {
	Group
	Members []UserBrief `json:"members"`
} // @name GroupWithMembers

type GroupListResponse = Response[[]Group] // @name GroupListResponse
