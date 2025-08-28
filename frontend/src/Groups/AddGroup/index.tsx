import {FormEvent, useState} from "react";
import {getApiDomain} from "../../config.tsx";
import {useNavigate} from "react-router-dom";

export default function AddGroup() {
    const [name, setName] = useState<string>('');
    const [error, setError] = useState<string | null>(null);
    const [loading, setLoading] = useState<boolean>(false);
    const navigate = useNavigate();

    async function handleSubmit(e: FormEvent) {
        e.preventDefault();
        setLoading(true);
        setError(null);

        try {
            const response = await fetch(getApiDomain() + '/api/v1/groups', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({name})
            });

            if (response.status === 204) {
                navigate("/groups")
            } else {
                const result = await response.json();
                setError(result.error);
            }
        } catch (error) {
            console.error(error);
            setError("Network error while creating group");
        } finally {
            setLoading(false);
        }
    }

    return (
        // TOdo: Implement Validator
        <form onSubmit={handleSubmit}>
            <fieldset className="fieldset">
                <legend className="fieldset-legend">Group name</legend>
                <input
                    type="text"
                    className="input"
                    placeholder="My awesome group"
                    value={name}
                    onChange={(e) => {
                        setName(e.target.value);
                    }}
                    required
                />
            </fieldset>

            {error && <div className="text-red-500 text-sm">{error}</div>}

            <button className="btn btn-primary" type="submit" disabled={loading}>Add Group
                {loading ? "Creating Group..." : "Create Group"}
            </button>
        </form>
    )
}