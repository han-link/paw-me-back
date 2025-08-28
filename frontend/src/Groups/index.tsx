import {useEffect, useState} from "react";
import {getApiDomain} from "../config.tsx";
import {Group} from "../types";
import {useNavigate, useOutletContext} from "react-router-dom";
import {PlusIcon} from '@heroicons/react/24/outline'
import {LayoutOutletContext} from "../Layout.tsx";

export default function Groups() {
    const navigate = useNavigate();
    const [groups, setGroups] = useState<Group[]>([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState<string | null>(null);

    const {setHeaderAction, setHeaderIcon} = useOutletContext<LayoutOutletContext>();

    useEffect(() => {
        setHeaderAction("/groups/new");
        setHeaderIcon(PlusIcon)

        return () => {
            setHeaderIcon(null)
            setHeaderAction(null)
        }
    }, [setHeaderIcon, setHeaderAction]);

    useEffect(() => {
        async function fetchGroups() {
            try {
                const response = await fetch(getApiDomain() + "/api/v1/groups");
                const result = await response.json();

                if (result.success) {
                    setGroups(result.data);
                } else {
                    setError(result.error || "Failed to load groups");
                }
            } catch (err) {
                setError("Network error while fetching groups");
            } finally {
                setLoading(false);
            }
        }

        void fetchGroups();
    }, []);

    if (loading) {
        return <p>Loading groups...</p>;
    }

    if (error) {
        return <p className="text-red-500">{error}</p>;
    }

    return (
        <ul className="list bg-base-100 rounded-box shadow-md">
            <li className="p-4 pb-2 text-xs opacity-60 tracking-wide">Your Groups</li>

            {groups.map((group) => (
                <li onClick={() => navigate(`/groups/${group.id}`)} key={group.id} className="list-row cursor-pointer">
                    <div>
                        {/* Just a placeholder avatar for now */}
                        <img
                            className="size-10 rounded-box"
                            src={`https://api.dicebear.com/8.x/identicon/svg?seed=${group.id}`}
                            alt={group.name}
                        />
                    </div>
                    <div>
                        <div>{group.name}</div>
                    </div>
                </li>
            ))}
        </ul>
    );
}