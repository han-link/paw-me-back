import {useLocation, useNavigate} from "react-router-dom";
import {navItems} from "../constants.ts";

export function Dock() {
    const navigate = useNavigate();
    const location = useLocation();

    const isActive = (path: string) => location.pathname === path;

    return (
        <div className="dock">
            {navItems.map(({ path, label, Icon }) => (
                <button
                    key={path}
                    className={isActive(path) ? "dock-active" : ""}
                    onClick={() => navigate(path)}
                >
                    <Icon className="size-6" />
                    <span className={isActive(path) ? "dock-active" : "dock-label"}>
                        {label}
                    </span>
                </button>
            ))}
        </div>
    );
}