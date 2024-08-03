import { ReactNode } from "react";

export type StrengthLevel = "weak" | "good" | "strong";

export type NavbarItem = {
    icon?: ReactNode;
    name: string;
    href: string;
}

export type RequireAuthProps = {
    children: ReactNode;
    others?: ReactNode;
}

export type FormProps = {
    type: "login" | "signup";
}

export type CustomInputProps = {
    type: "email" | "password" | "text";
    placeholder: string;
    inputRef: any;
    children?: ReactNode;
    isPassword?: boolean;
}

interface Model {
    id: number;
    createdAt: Date;
    updatedAt: Date;
    deletedAt?: Date | null;
}

export interface User extends Model {
    username: string;
    email: string;
    password: string;
    picUrl?: string | null;
    isVerified?: boolean | null;
    isMod?: boolean | null;
    bio?: string | null;
    website?: string | null;
    posts?: Post[] | null;
    friends?: User[] | null;
    userAnimes?: UserAnime[] | null;
}

export interface Post extends Model {
}

export interface UserAnime extends Model {
}
