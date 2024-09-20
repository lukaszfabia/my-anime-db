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
    reviews: Review[];
    stats?: UserStat;
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
    userId: number;
    animeId: number;
    score: string;
    watchStatus: string;
    isFav: boolean;
    review: Review;

    user: User;
    anime: Anime;
}

export interface Review extends Model {
    userId: number;
    animeId: number;
    content: string;

    userAnime: UserAnime;
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
    logoUrl?: string;
    website?: string;
}

export interface Role {
    actorId: number;
    characterId: number;
    animeId: number;
    role: string;

    voiceActor: VoiceActor;
    character: Character;
    anime: Anime;
}

export interface Character extends Model {
    name: string;
    lastname: string;
    picUrl?: string;
    information: string;
    roles: Role[];
}

export interface VoiceActor extends Model {
    name: string;
    lastname: string;
    picUrl?: string;
    birthdate?: Date;
    roles: Role[];
}

export interface OtherTitles extends Model {
    title: string;
    animeId: number;
}

export interface Anime extends Model {
    title: string;
    alternativeTitles?: OtherTitles[];
    animeType: string;
    episodes: number;
    description: string;
    episodeLength: number;
    startDate?: Date;
    finishDate?: Date;
    pegi: string;
    status: string;
    picUrl: string;
    stats: AnimeStat;
    genres: Genre[];
    studio?: Studio;
    roles: Role[];
    prequel?: Anime;
    sequel?: Anime;
    reviews: Review[];
}

export interface UserStat extends Model {
    userId: number;
    favGenres: string[];
    watchedHours: number;
    postedReviews: number;
}


export interface AnimeStat extends Model {
    animeId: number;
    popularity: number;
    score: number;
    mostPopularGrade: string;
}