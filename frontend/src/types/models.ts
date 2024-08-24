export interface Model {
    id: number;
    createdAt: Date;
    updatedAt: Date;
    deletedAt?: Date | null;
}

export interface User extends Model {
    username: string;
    email: string;
    picUrl: string;
    isVerified: boolean;
    isMod: boolean;
    bio: string;
    website?: string;
    posts: Post[];
    friends: User[];
    userAnimes: UserAnime[];
}

export interface Post extends Model {
    title: string;
    content: string;
    image?: string;
    isPublic: boolean;
    userId: number;
    user: User;
}

export interface UserAnime extends Model {
}

export type RequestStatus = "accepted" | "rejected" | "pending" | "respond" | "cancel";

export interface FriendRequest extends Model {
    senderId: number;
    receiverId: number;
    status: RequestStatus;
    sender: User;
    receiver: User;
}
