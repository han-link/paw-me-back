import SuperTokens, {SuperTokensWrapper} from "supertokens-auth-react";
import {getSuperTokensRoutesForReactRouterDom} from "supertokens-auth-react/ui";
import * as ReactRouter from "react-router-dom";
import {BrowserRouter, Navigate, Route, Routes} from "react-router-dom";
import {PreBuiltUIList, SuperTokensConfig} from "./config";
import Friends from "./Friends";
import Groups from "./Groups";
import Profile from "./Profile";
import Activities from "./Activities";
import AddGroup from "./Groups/AddGroup";
import GroupDetail from "./Groups/Detail";
import Settings from "./Groups/Detail/Settings";
import GroupDetailLayout from "./Groups/Detail/Layout.tsx";
import Layout from "./Layout.tsx";

SuperTokens.init(SuperTokensConfig);

export default function App() {
    return (
        <SuperTokensWrapper>
            <BrowserRouter>
                <main className="App app-container">
                    <Routes>
                        {/* SuperTokens auth routes */}
                        {getSuperTokensRoutesForReactRouterDom(ReactRouter, PreBuiltUIList)}

                        <Route path="/" element={<Navigate to="/groups" replace/>}/>

                        <Route element={<Layout/>}>
                            <Route path="groups">
                                <Route index element={<Groups/>}/>
                                <Route path="new" element={<AddGroup/>}/>
                                <Route path=":id" element={<GroupDetailLayout/>}>
                                    <Route index element={<GroupDetail/>}/>
                                    <Route path="settings" element={<Settings/>}/>
                                </Route>
                            </Route>
                            <Route path="friends" element={<Friends/>}/>
                            <Route path="activities" element={<Activities/>}/>
                            <Route path="profile" element={<Profile/>}/>
                        </Route>
                    </Routes>
                </main>
            </BrowserRouter>
        </SuperTokensWrapper>
    );
}