"use client";

import React, { FC } from "react";
import Link from "next/link";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faLightbulb } from "@fortawesome/free-solid-svg-icons";

import { ThemeSwitcher } from "./buttons/themeSwitcher";
import { SearchBar } from "./ui/navbar/searchBar";
import { Logged } from "./ui/navbar/logged";
import { NotLogged } from "./ui/navbar/notlogged";
import { useAuth } from "./providers/auth";

export const Navbar: FC = () => {
    const { user, loading } = useAuth();


    if (loading) return;

    return (
        <div className="navbar fixed z-10 backdrop-blur-lg">
            <div className="flex-1">
                <Link href="/" className="btn btn-ghost md:text-2xl text-xl">
                    <div>
                        myanime<span className="text-indigo-600">.db</span>
                    </div>
                </Link>
            </div>
            <div className="flex-none md:gap-2">
                <SearchBar />
                <Link className="btn btn-ghost" href="/explorer"><FontAwesomeIcon icon={faLightbulb} width={15} height={15} /> Explorer</Link>

                {user ? <Logged user={user} /> : <NotLogged />}
            </div>
            <div className="px-2">
                <ThemeSwitcher />
            </div>
        </div>
    );
};
