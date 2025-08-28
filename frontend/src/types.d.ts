type UserBrief = {
    id: string;
    username: string;
}

export type Group = {
    id: string;
    name: string;
    created_at: string;
    updated_at: string;
    owner: UserBrief
};

export type GroupDetail = Group & {
    members: UserBrief[];
}