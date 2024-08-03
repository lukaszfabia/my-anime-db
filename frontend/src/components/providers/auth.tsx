"use client";

import { createContext, useContext, useState, useEffect, ReactNode, useCallback } from 'react';
import api from '@/lib/api';
import { ACCESS_TOKEN } from '@/lib/constants';
import { redirect } from 'next/navigation';
import { User } from '@/types';
import readError from '@/lib/handleErrors';
import { AxiosError } from 'axios';

interface LoginProps {
    username: string;
    password: string;
    remember?: boolean | null;
}

interface SignupProps extends LoginProps {
    email: string;
    picUrl: File;
}

type AuthContextProps = {
    user?: User | null;
    login: (props: LoginProps) => Promise<string | null>;
    signup: (props: SignupProps) => Promise<string | null>;
    logout: (callback?: () => void) => void
    loading: boolean;
}

const AuthContext = createContext<AuthContextProps>({} as AuthContextProps);

export function AuthProvider({ children }: { children: ReactNode }) {
    const [user, setUser] = useState<User | null>(null);
    const [loading, setLoading] = useState<boolean>(true);

    useEffect(() => {
        const token = localStorage.getItem(ACCESS_TOKEN) || sessionStorage.getItem(ACCESS_TOKEN);
        if (token) {
            getUser().then((user: User | null) => {
                if (user) {
                    setUser(user);
                } else {
                    localStorage.removeItem(ACCESS_TOKEN);
                    sessionStorage.removeItem(ACCESS_TOKEN);
                }
                setLoading(false);
            }).catch(() => {
                localStorage.removeItem(ACCESS_TOKEN);
                sessionStorage.removeItem(ACCESS_TOKEN);
                setLoading(false);
            });
        } else {
            setLoading(false);
        }
    }, []);

    const login = async (props: LoginProps) => {
        try {
            const formData = new FormData();
            formData.append("username", props.username);
            formData.append("password", props.password);

            const response = await api.post('/login/', formData);

            const token: string = response.data.token;

            props.remember ? localStorage.setItem(ACCESS_TOKEN, token) : sessionStorage.setItem(ACCESS_TOKEN, token);

            const user = await getUser();
            setUser(user);

            return null;
        } catch {
            return "Login failed, username or password are wrong"
        }

    };

    const logout = (callback?: () => void) => {
        localStorage.removeItem(ACCESS_TOKEN);
        sessionStorage.removeItem(ACCESS_TOKEN);
        setUser(null);

        callback && callback();

        redirect("/login");
    }

    const signup = async (props: SignupProps) => {
        try {
            const signupData = new FormData();
            signupData.append("username", props.username);
            signupData.append("password", props.password);
            signupData.append("email", props.email);
            signupData.append("picUrl", props.picUrl);

            await api.post("/sign-up/", signupData)

            await login({ username: props.username, password: props.password });

            return null;
        } catch (error: any) {
            return readError(error)
        }
    }

    return (
        <AuthContext.Provider value={{ user, login, logout, signup, loading }}>
            {children}
        </AuthContext.Provider>
    );
}

const getUser = async (): Promise<User | null> => {
    try {
        const response = await api.get("/auth/account/me/");
        return response.data;
    } catch {
        return null;
    }
}

export function useAuth() {
    return useContext(AuthContext);
}
