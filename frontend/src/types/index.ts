import { FormEvent, MutableRefObject, ReactNode } from "react";
import { Post } from "./models";

export type StrengthLevel = "weak" | "good" | "strong";

export interface Stat {
    title: string;
    value: string | number | string[];
    desc?: string;
    icon: ReactNode;
}

export const strengthStyles: Record<StrengthLevel, string> = {
    "weak": "input-error",
    "good": "input-warning",
    "strong": "input-success"
};

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
    type: "email" | "password" | "text" | "date" | "number";
    name: string;
    placeholder: string;
    inputRef?: MutableRefObject<HTMLInputElement | null>;
    children?: ReactNode;
    required?: boolean;
    disabled?: boolean;
    defaultValue?: string | number;
}


export interface PostFormProps {
    submitFunc: (e: FormEvent<HTMLFormElement>) => void;
    defaultValues?: Post;
}

export interface Anchor {
    title: string;
    id: string;
}