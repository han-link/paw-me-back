import { ReactNode } from "react";

export const ComponentWrapper = ({ children }: { children: ReactNode }) => {
    return <div className="mb-18">{children}</div>;
};
