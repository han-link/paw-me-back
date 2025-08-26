package main

import (
	"paw-me-back/internal/env"
	"paw-me-back/internal/model"
	"paw-me-back/internal/store"
	"regexp"

	"github.com/supertokens/supertokens-golang/recipe/dashboard"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/emailpassword/epmodels"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/userroles"
	"github.com/supertokens/supertokens-golang/supertokens"
)

func generateSuperTokenConfig(
	authBasePath string,
	apiUrl string,
	frontendUrl string,
	userStore store.Users,
) supertokens.TypeInput {
	return supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			ConnectionURI: "http://localhost:3567",
			APIKey:        env.GetString("SUPER_TOKEN_API_KEY", ""),
		},
		AppInfo: supertokens.AppInfo{
			AppName:         "SuperTokens Demo App",
			APIDomain:       apiUrl,
			WebsiteDomain:   frontendUrl,
			APIBasePath:     &authBasePath,
			WebsiteBasePath: &authBasePath,
		},
		RecipeList: []supertokens.Recipe{
			emailpassword.Init(&epmodels.TypeInput{
				SignUpFeature: &epmodels.TypeInputSignUp{
					FormFields: []epmodels.TypeInputFormField{
						{
							ID: "username",
							Validate: func(v interface{}, _ string) *string {
								s, _ := v.(string)
								if len(s) < 3 || len(s) > 32 {
									msg := "username must be 3–32 characters"
									return &msg
								}
								matched, _ := regexp.MatchString(`^[a-z0-9_-]+$`, s)
								if !matched {
									msg := "username can only contain lowercase letters, numbers, underscores, and hyphens"
									return &msg
								}
								user, err := isUsernameTaken(userStore, s)
								if err != nil {
									msg := "temporary error validating username"
									return &msg
								}
								if user {
									msg := "username already taken"
									return &msg
								}
								return nil
							},
						},
					},
				},
				Override: &epmodels.OverrideStruct{
					APIs: func(impl epmodels.APIInterface) epmodels.APIInterface {
						// After successful sign-up, persist username in your DB mapped to ST userId.
						if impl.SignUpPOST != nil {
							og := *impl.SignUpPOST
							*impl.SignUpPOST = func(fields []epmodels.TypeFormField, tenantId string, opt epmodels.APIOptions, uc supertokens.UserContext) (epmodels.SignUpPOSTResponse, error) {
								resp, err := og(fields, tenantId, opt, uc)
								if err != nil {
									return epmodels.SignUpPOSTResponse{}, err
								}
								if resp.OK != nil {
									var username, email string
									for _, f := range fields {
										if f.ID == "username" {
											username, _ = f.Value.(string)
										}
										if f.ID == "email" {
											email, _ = f.Value.(string)
										}
									}
									if err := createUser(userStore, username, email, resp.OK.User.ID); err != nil {
										return epmodels.SignUpPOSTResponse{}, err
									}
								}
								return resp, nil
							}
						}
						return impl
					},
				},
			}),
			session.Init(nil),
			dashboard.Init(nil),
			userroles.Init(nil),
		},
	}
}

func isUsernameTaken(userStore store.Users, username string) (bool, error) {
	return userStore.UsernameExists(username)
}

func createUser(userStore store.Users, username string, email string, superTokenId string) error {
	user := model.User{
		Username:     username,
		Email:        email,
		SuperTokenID: superTokenId,
	}
	err := userStore.Create(&user)
	return err
}
