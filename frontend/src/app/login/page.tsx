"use client";

import { FC } from "react";

import AbstractForm from "@/components/accountForm";
import { useAuth } from "@/components/providers/auth";
import { Spinner } from "@/components/ui/spinner";

const Login: FC = () => {

    const { loading } = useAuth();

    if (loading) {
        return (
            <div className="flex items-center justify-center">
                <Spinner />
            </div>
        )
    }

    return (
        <div className="flex flex-col md:flex-row items-center justify-center bg-base-200 rounded-3xl">
            <div className="flex-1 md:mb-0 mb-4 md:pb-20 pb-10 md:text-left text-center p-5">
                <h1 className="md:text-5xl text-3xl font-extrabold py-4">
                    <span className="text-yellow-300">Welcome</span> back 👋
                </h1>
                <p className="px-4 md:px-0 md:pt-4 md:text-lg">
                    Log in to access your account and connect with your favorite anime community. Stay updated, share your thoughts, and enjoy exclusive member content.
                </p>
                <p className="md:text-xl text-lg py-3 flex justify-end">
                    Your<b className="bg-gradient-to-r from-violet-600 via-pink-500 to-amber-500 inline-block text-transparent bg-clip-text px-1">adventure</b>is waiting!
                </p>
            </div>
            <div className="md:px-4 px-0"></div>
            <div className="flex-1 w-full">
                <AbstractForm type="login" />
            </div>
        </div>
    );
};


export default Login;