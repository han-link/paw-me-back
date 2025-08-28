"use client";

import EmailPassword from "supertokens-auth-react/recipe/emailpassword";
import {EmailPasswordPreBuiltUI} from "supertokens-auth-react/recipe/emailpassword/prebuiltui";
import Session from "supertokens-auth-react/recipe/session";


export function getApiDomain() {
    const apiPort = 3000;
    return `http://localhost:${apiPort}`;
}

export function getWebsiteDomain() {
    const websitePort = 5173;
    return `http://localhost:${websitePort}`;
}

const reUsername = /^[a-z0-9_-]{3,32}$/;
export const reUUID = /^[0-9a-f]{8}-[0-9a-f]{4}-[1-5][0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$/i

export const SuperTokensConfig = {
    appInfo: {
        appName: "Paw Me Back",
        apiDomain: getApiDomain(),
        websiteDomain: getWebsiteDomain(),
        apiBasePath: "/api/auth",
        websiteBasePath: "/auth",
    },
    
    recipeList: [
        EmailPassword.init({
            signInAndUpFeature: {
                signUpForm: {
                    formFields: [
                        {
                            id: "username",
                            label: "Username",
                            placeholder: "john.doe",
                            validate: async (input) => {
                                return reUsername.test(input) ? undefined : "3–32 chars, lowercase letters, numbers, _ or -"
                            },
                        },
                    ]
                }
            }
        }),
        Session.init()
    ],
    getRedirectionURL: async (context: any) => {
        if (context.action === "SUCCESS") {
            return "/groups";
        }
        return undefined;
    },
};

export const PreBuiltUIList = [EmailPasswordPreBuiltUI];