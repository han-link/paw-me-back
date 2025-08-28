import { GroupDetail } from "../../types";
import { useCallback, useEffect, useMemo, useState } from "react";
import { getApiDomain, reUUID } from "../../config";
import { Outlet, useNavigate, useOutletContext, useParams } from "react-router-dom";
import { LayoutOutletContext } from "../../Layout";

export type GroupLayoutOutletCtx = {
    group: GroupDetail;
    refresh: () => Promise<void>;
    rootLayoutCtx: LayoutOutletContext;
};

export default function GroupDetailLayout() {
    const { id } = useParams<{ id: string }>();
    const navigate = useNavigate();

    const [group, setGroup] = useState<GroupDetail | null>(null);
    const [loading, setLoading] = useState(true);
    const [error, setError]   = useState<string | null>(null);

    const rootLayoutCtx = useOutletContext<LayoutOutletContext>();

    const fetchGroup = useCallback(async () => {
        if (!id) return;
        setLoading(true);
        try {
            const res = await fetch(`${getApiDomain()}/api/v1/groups/${id}`, { credentials: "include" });
            const json = await res.json();
            if (json.success) {
                setGroup(json.data as GroupDetail);
                setError(null);
            } else {
                setError(json.error || "Failed to load group");
                setGroup(null);
            }
        } catch {
            setError("Network error while fetching group");
            setGroup(null);
        } finally {
            setLoading(false);
        }
    }, [id]);

    useEffect(() => {
        if (!id || !reUUID.test(id)) {
            navigate("/groups", { replace: true });
            return;
        }
        void fetchGroup();
    }, [id, navigate, fetchGroup]);

    const groupCtx: GroupLayoutOutletCtx = useMemo(
        () => ({ group: group as GroupDetail, refresh: fetchGroup, rootLayoutCtx }),
        [group, fetchGroup, rootLayoutCtx]
    );

    if (loading) return <p>Loading group…</p>;
    if (error)   return <p className="text-error">{error}</p>;
    if (!group)  return null;

    return <Outlet context={groupCtx} />;
}
