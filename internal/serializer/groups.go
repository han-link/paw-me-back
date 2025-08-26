package mapper

import (
	"paw-me-back/internal/model"
	"paw-me-back/internal/types"
)

func toUserBrief(u *model.User) types.UserBrief {
	if u == nil {
		return types.UserBrief{}
	}
	return types.UserBrief{
		ID:       u.ID,
		Username: u.Username,
	}
}

func SanitizeSingleGroup(g *model.Group) types.GroupWithMembers {
	if g == nil {
		return types.GroupWithMembers{}
	}

	members := make([]types.UserBrief, len(g.Members))
	for i := range g.Members {
		members[i] = toUserBrief(g.Members[i])
	}

	return types.GroupWithMembers{
		Group: types.Group{
			Base: types.Base{
				ID:        g.ID,
				CreatedAt: g.CreatedAt,
				UpdatedAt: g.UpdatedAt,
			},
			Name:  g.Name,
			Owner: toUserBrief(g.Owner),
		},
		Members: members,
	}
}

func SanitizeGroupList(gs []model.Group) []types.Group {
	out := make([]types.Group, len(gs))
	for i := range gs {
		g := &gs[i]
		out[i] = types.Group{
			Base: types.Base{
				ID:        g.ID,
				CreatedAt: g.CreatedAt,
				UpdatedAt: g.UpdatedAt,
			},
			Name:  g.Name,
			Owner: toUserBrief(g.Owner),
		}
	}
	return out
}
