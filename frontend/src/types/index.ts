import { Mutable } from "next/dist/client/components/router-reducer/router-reducer-types";
import { Dispatch, MutableRefObject, ReactNode, SetStateAction } from "react";

export type StrengthLevel = "weak" | "good" | "strong";

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
    type: "email" | "password" | "text";
    name: string;
    placeholder: string;
    inputRef?: MutableRefObject<HTMLInputElement | null>;
    children?: ReactNode;
}
