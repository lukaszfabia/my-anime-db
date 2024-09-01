export interface Model {
    id: number;
    createdAt: Date;
    updatedAt: Date;
    deletedAt?: Date | null;
}

export interface User extends Model {
    username: string;
    email: string;
    picUrl?: string;
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


export interface Genre extends Model {
    name: string;
}

export interface Studio extends Model {
    name: string;
    establishedDate: Date;
    logoUrl: string;
    website: string;
}

export interface Character extends Model {
    name: string;
    lastname: string;
    picUrl: string;
    information: string;
    // roles : Role[];
}

export interface VoiceActor extends Model {
    name: string;
    lastname: string;
    picUrl: string;
    birthdate: Date;
    // roles : Role[];
}