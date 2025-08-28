import {useOutletContext} from "react-router-dom";
import {GroupLayoutOutletCtx} from "../Layout.tsx";
import {useEffect} from "react";

export default function Settings() {
    const { group, rootLayoutCtx: { setHeaderParent } } = useOutletContext<GroupLayoutOutletCtx>();

    useEffect(() => {
        setHeaderParent(`/groups/${group.id}`);

        return () => {
            setHeaderParent(null)
        }
    }, [setHeaderParent]);

    return (
        <>
            <div className="flex flex-row">
                <div className="basis-1/4">
                    <div className="skeleton h-16 w-16"></div>
                </div>
                <div className="basis-3/4">{group.name}</div>
            </div>
            <ul className="list bg-base-100 rounded-box shadow-md">
                <li className="p-4 pb-2 text-xs opacity-60 tracking-wide">Group members</li>

                {group.members.map((member) => (
                    <li key={member.id} className="list-row cursor-pointer">
                        <div>
                            <img
                                className="size-10 rounded-box"
                                src={`https://api.dicebear.com/8.x/identicon/svg?seed=${member.username}`}
                                alt={member.username}
                            />
                        </div>
                        <div>
                            <div>{member.username}</div>
                        </div>
                    </li>
                ))}
            </ul>
        </>
    )
}