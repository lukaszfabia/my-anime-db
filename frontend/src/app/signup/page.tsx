"use client";

import { FC } from "react";

import { useAuth } from "@/components/providers/auth";
import { Spinner } from "@/components/ui/spinner";
import AbstractForm from "@/components/ui/forms/accountForm";


const Signup: FC = () => {
    const { loading } = useAuth();

    if (loading) return <Spinner />


    return (
        <div className="flex flex-col-reverse lg:flex-row items-center justify-center bg-base-200 rounded-3xl">
            <div className="flex-1 w-full">
                <AbstractForm type="signup" />
            </div>
            <div className="md:px-4 px-0"></div>
            <div className="flex-1 md:mb-0 mb-4 md:pb-20 pb-10 md:text-left text-center p-7">
                <h1 className="lg:text-5xl md:text-4xl sm:text-3xl text-2xl font-extrabold py-4">
                    Create an <span className="text-violet-600">account</span>!
                </h1>
                <p className="px-4 md:px-0 md:pt-4 md:text-lg text-center text-base">
                    Join us and become a part of the <span className="font-semibold">men of culture</span> community. Share your passion, connect with fellow fans, and explore the latest in the <b>anime</b> world.
                </p>
                <p className="text-xl py-3 flex justify-end">Your<b className="bg-gradient-to-r from-sky-400 via-red-500 to-violet-900 inline-block bg-clip-text text-transparent px-1">journey</b>begins <b className="px-1">here</b>!</p>
            </div>
        </div>
    )
}

export default Signup;