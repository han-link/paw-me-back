package main

import (
	"net/http"
	"paw-me-back/internal/model"
	"paw-me-back/internal/serializer"
	"paw-me-back/internal/types"
)

type groupKey string

const groupCtx groupKey = "group"

// Retrieve all groups
//
//	@Summary		Retrieve groups
//	@Description	Retrieve all groups
//	@Tags			groups
//	@Produce		json
//	@Success		200	{object}	types.GroupListResponse
//	@Failure		401	{object}	error
//	@Failure		500	{object}	error
//	@Router			/groups [get]
func (app *application) getGroupsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user := getUserFromContext(r)

	groups, err := app.store.Groups.GetAll(ctx, user.ID)

	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	res := mapper.SanitizeGroupList(groups)

	if err := app.jsonResponse(w, http.StatusOK, res); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// Get singular group
//
//	@Summary		Retrieve group
//	@Description	Retrieve a group
//	@Tags			groups
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"Group ID"	Format(uuid)
//	@Success		200	{object}	types.GroupWithMembers
//	@Failure		401	{object}	error
//	@Failure		500	{object}	error
//	@Router			/groups/{id} [get]
func (app *application) getGroupHandler(w http.ResponseWriter, r *http.Request) {
	group := getGroupFromContext(r)

	res := mapper.SanitizeSingleGroup(group)

	if err := app.jsonResponse(w, http.StatusOK, res); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

// Add Members
//
//	@Summary	Update a group and modify its members
//	@Tags		groups
//	@Param		id	path	string	true	"Group ID"	Format(uuid)
//	@Accept		json
//	@Produce	json
//	@Param		body	body		types.AddMembersPayload	true	"Add members"
//	@Success	200		{object}	types.GroupWithMembers
//	@Failure	400		{object}	error
//	@Failure	404		{object}	error
//	@Router		/groups/{id}/members [put]
func (app *application) addMembersToGroupHandler(w http.ResponseWriter, r *http.Request) {
	var payload types.AddMembersPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	group := getGroupFromContext(r)

	ctx := r.Context()

	if err := app.store.Groups.AddMembers(ctx, group, payload.UserIDs); err != nil {
		app.internalServerError(w, r, err)
	}

	res := mapper.SanitizeSingleGroup(group)

	if err := app.jsonResponse(w, http.StatusOK, res); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func getGroupFromContext(r *http.Request) *model.Group {
	group, _ := r.Context().Value(groupCtx).(*model.Group)
	return group
}
