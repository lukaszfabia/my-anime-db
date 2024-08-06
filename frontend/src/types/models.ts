interface Model {
    ID: number;
    CreatedAt: Date;
    UpdatedAt: Date;
    DeletedAt?: Date | null;
}

export interface User extends Model {
    username: string;
    email: string;
    password: string;
    picUrl: string;
    isVerified: boolean;
    isMod: boolean;
    bio: string;
    website: string;
    posts: Post[];
    friends: User[];
    userAnimes: UserAnime[];
}

export interface Post extends Model {
}

export interface UserAnime extends Model {
}