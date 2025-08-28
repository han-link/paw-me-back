import {Dock} from "./Components/Dock.tsx";
import {ComponentWrapper} from "./Components/ComponentWrapper.tsx";
import {Header} from "./Components/Header.tsx";
import {ComponentType, Dispatch, SetStateAction, SVGProps, useState} from "react";
import {SessionAuth} from "supertokens-auth-react/recipe/session";
import {Outlet} from "react-router-dom";

export interface LayoutOutletContext {
    setHeaderAction: Dispatch<SetStateAction<string | null>>;
    setHeaderIcon: Dispatch<SetStateAction<ComponentType<SVGProps<SVGSVGElement>> | null>>;
    setHeaderParent: Dispatch<SetStateAction<string | null>>;
}

export default function Layout() {
    const [headerAction, setHeaderAction] = useState<string | null>(null);
    const [headerIcon, setHeaderIcon] = useState<ComponentType<SVGProps<SVGSVGElement>> | null>(null);
    const [headerParent, setHeaderParent] = useState<string | null>(null);
    return (
        <SessionAuth>
            <Header action={headerAction} icon={headerIcon} parent={headerParent}/>
            <div className="fill" id="home-container">
                <ComponentWrapper>
                    <Outlet context={{setHeaderAction, setHeaderIcon, setHeaderParent} as LayoutOutletContext}/>
                </ComponentWrapper>
            </div>
            <Dock/>
        </SessionAuth>
    );
}