"use client";

import EmailPassword from "supertokens-auth-react/recipe/emailpassword";
import {EmailPasswordPreBuiltUI} from "supertokens-auth-react/recipe/emailpassword/prebuiltui";
import Session from "supertokens-auth-react/recipe/session";


export function getApiDomain() {
    const apiPort = 3000;
    const apiUrl = `http://localhost:${apiPort}`;
    return apiUrl;
}

export function getWebsiteDomain() {
    const websitePort = 5173;
    const websiteUrl = `http://localhost:${websitePort}`;
    return websiteUrl;
}

const reUsername = /^[a-z0-9_-]{3,32}$/;

export const SuperTokensConfig = {
    appInfo: {
        appName: "Paw Me Back",
        apiDomain: getApiDomain(),
        websiteDomain: getWebsiteDomain(),
        apiBasePath: "/auth",
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
            return "/dashboard";
        }
        return undefined;
    },
};

export const recipeDetails = {
    docsLink: "https://supertokens.com/docs/quickstart/introduction",
};

export const PreBuiltUIList = [EmailPasswordPreBuiltUI];



export const ComponentWrapper = (props: { children: JSX.Element }): JSX.Element => {
    let childrenToRender = props.children;

    
    return childrenToRender;
}