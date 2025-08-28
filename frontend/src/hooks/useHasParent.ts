import { useLocation } from "react-router-dom";

export function useHasParent() {
    const location = useLocation();
    const segments = location.pathname.split('/').filter(Boolean);
    return segments.length > 1;
}