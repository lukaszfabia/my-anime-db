"use client";

import { createContext, useContext, useState, useEffect, ReactNode, FormEvent } from 'react';
import api from '@/lib/api';
import { ACCESS_TOKEN } from '@/lib/constants';
import { redirect } from 'next/navigation';
import { User } from '@/types/models';
import { toast } from 'react-toastify';
import { AxiosResponse } from 'axios';

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
            }).catch((_: any) => {
                localStorage.removeItem(ACCESS_TOKEN);
                sessionStorage.removeItem(ACCESS_TOKEN);
                setLoading(false);
            });
        }
        setLoading(false);
    }, []);

    const login = (setError: (error: string) => void, remember: boolean, e?: FormEvent<HTMLFormElement>, loginData?: LoginProps) => {

        const getToken = async (formData: FormData) => await api.post<GoResponse>('/login/', formData).then((response: AxiosResponse<GoResponse>) => {
            const token: string = response.data.data!;
            remember ? localStorage.setItem(ACCESS_TOKEN, token) : sessionStorage.setItem(ACCESS_TOKEN, token);
            getUser().then((user: User | null) => {
                setUser(user);
            });
            toast.success('Successfully logged in 👌');
        }).catch((_: any) => {
            setError("Invalid credentials");
        });

        if (e) {
            getToken(new FormData(e.currentTarget));
        } else if (loginData) {
            const formData = new FormData();
            formData.append("username", loginData?.username);
            formData.append("password", loginData?.password);
            getToken(formData);
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
        api.delete<GoResponse>("/auth/account/me/").then((_: AxiosResponse<GoResponse>) => {
            localStorage.removeItem(ACCESS_TOKEN);
            sessionStorage.removeItem(ACCESS_TOKEN);
            setUser(null);
            toast.info("Account deleted successfully");
        }).catch((_: any) => {
            toast.error("Something went wrong!");
        });
    };

    const createAccount = (e: FormEvent<HTMLFormElement>, setError: (error: string) => void) => {

        const signup = (formData: FormData) => {
            api.post("/sign-up/", formData).then((_: AxiosResponse<GoResponse>) => {
                login(setError, false, undefined, {
                    username: formData.get("username")?.toString()!,
                    password: formData.get("password")?.toString()!,
                });
                toast.success('Account created successfully');
            }).catch((_: any) => {
                toast.error('Something went wrong!');
            });
        };

        signup(new FormData(e.currentTarget));
    }

    // use it for updating user data when you want to see immediate changes
    const refreshUser = () => {
        getUser().then((user: User | null) => {
            setUser(user);
        }).catch((_: any) => { });
    }

    return (
        <AuthContext.Provider value={{ user, login, logout, removeAccount, refreshUser, createAccount, loading }}>
            {children}
        </AuthContext.Provider>
    );
}

const getUser = async (): Promise<User | null> => {
    return await api.get<GoResponse>("/auth/account/me/")
        .then((response: AxiosResponse<GoResponse>) => response.data.data)
        .catch((_: any) => null);
}

export function useAuth() {
    return useContext(AuthContext);
}
