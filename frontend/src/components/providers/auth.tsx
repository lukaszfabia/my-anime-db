"use client";

import { createContext, useContext, useState, useEffect, ReactNode, FormEvent, Dispatch, SetStateAction } from 'react';
import api from '@/lib/api';
import { ACCESS_TOKEN } from '@/lib/constants';
import { redirect } from 'next/navigation';
import { User } from '@/types/models';
import { toast } from 'react-toastify';
import { AxiosError, AxiosResponse } from 'axios';

interface LoginProps {
    username: string;
    password: string;
}

type AuthContextProps = {
    user: User | null;
    login: (setError: (error: string) => void, remember: boolean, e?: FormEvent<HTMLFormElement>, loginData?: LoginProps) => void;
    logout: (callback?: () => void) => void;
    removeAccount: () => void;
    refreshUser: () => void;
    createAccount: (e: FormEvent<HTMLFormElement>, setError: (error: string) => void) => void;
    loading: boolean;
}

const AuthContext = createContext<AuthContextProps>({
    user: null,
    login: async () => "",
    logout: () => { },
    removeAccount: () => { },
    refreshUser: () => { },
    createAccount: () => { },
    loading: true
});

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
        }
        setLoading(false);
    }, []);

    const login = (setError: (error: string) => void, remember: boolean, e?: FormEvent<HTMLFormElement>, loginData?: LoginProps) => {

        const getToken = async (formData: FormData) => await api.post<GoTokenResponse>('/login/', formData).then((response: AxiosResponse<GoTokenResponse>) => {
            const token: string = response.data.token!;
            remember ? localStorage.setItem(ACCESS_TOKEN, token) : sessionStorage.setItem(ACCESS_TOKEN, token);
            getUser().then((user: User | null) => {
                setUser(user);
            })
            toast.success('Successfully logged in 👌');
        }).catch((error: AxiosError<GoTokenResponse>) => {
            const message = error.response?.data.error!;
            setError(message);
        });

        if (e) {
            getToken(new FormData(e.currentTarget));
        }
    };

    const logout = (callback?: () => void) => {
        localStorage.removeItem(ACCESS_TOKEN);
        sessionStorage.removeItem(ACCESS_TOKEN);
        setUser(null);

        callback && callback();

        redirect("/login");
    };

    const removeAccount = () => {
        api.delete<GoResponse>("/auth/account/me/").then((response: AxiosResponse<GoResponse>) => {
            localStorage.removeItem(ACCESS_TOKEN);
            sessionStorage.removeItem(ACCESS_TOKEN);
            setUser(null);
            toast.info(response.data.message!)
        }).catch((error: AxiosError<GoResponse>) => {
            const message: string = error.response?.data.error!;
            toast.error(message);
        });
    };

    const createAccount = (e: FormEvent<HTMLFormElement>, setError: (error: string) => void) => {

        const signup = (formData: FormData) => {
            api.post("/sign-up/", formData).then((response: AxiosResponse<GoResponse>) => {
                login(setError, false, undefined, {
                    username: formData.get("username")?.toString()!,
                    password: formData.get("password")?.toString()!,
                });
                toast.success(response.data.message);
            }).catch((error: AxiosError<GoResponse>) => {
                const message = error.response?.data.error!
                setError(message);
                toast.error('Something went wrong!');
            });
        };

        signup(new FormData(e.currentTarget));
    }

    // use it for updating user data when you want to see immediate changes
    const refreshUser = () => {
        getUser().then((user: User | null) => {
            setUser(user);
        });
    }

    return (
        <AuthContext.Provider value={{ user, login, logout, removeAccount, refreshUser, createAccount, loading }}>
            {children}
        </AuthContext.Provider>
    );
}

const getUser = async (): Promise<User | null> => {
    return await api.get<User>("/auth/account/me/")
        .then((response: AxiosResponse<User>) => response.data)
        .catch(() => null);
}

export function useAuth() {
    return useContext(AuthContext);
}
