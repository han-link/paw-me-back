import {useOutletContext} from "react-router-dom";
import {useEffect} from "react";
import {Cog6ToothIcon} from '@heroicons/react/24/outline'
import {GroupLayoutOutletCtx} from "./Layout.tsx";

export default function GroupDetail() {
    const { group, rootLayoutCtx: {setHeaderAction, setHeaderIcon} } = useOutletContext<GroupLayoutOutletCtx>();

    useEffect(() => {
        setHeaderAction(`/groups/${group.id}/settings`);
        setHeaderIcon(Cog6ToothIcon)

        return () => {
            setHeaderIcon(null)
            setHeaderAction(null)
        }
    }, [setHeaderIcon, setHeaderAction]);

    return (
        <div>
            <h1 className="text-xl font-bold">{group.name}</h1>
            <p className="text-sm text-gray-500">
                Owner: {group.owner.username}
            </p>
            <p className="text-xs opacity-70">ID: {group.id}</p>
        </div>
    );
}