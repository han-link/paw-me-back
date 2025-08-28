import { signOut } from "supertokens-auth-react/recipe/session";

import {useNavigate} from "react-router-dom";

export default function Profile() {
    const navigate = useNavigate();

    async function logoutClicked() {
        await signOut();
        navigate("/auth");
    }

    return (
        <>
            <button className="btn btn-soft" onClick={logoutClicked}>Logout</button>
        </>
    );
}
