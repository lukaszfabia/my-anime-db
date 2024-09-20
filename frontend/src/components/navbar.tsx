"use client";

import React, { FC } from "react";
import Link from "next/link";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faBars, faCompass, faUser, faXmark } from "@fortawesome/free-solid-svg-icons";

import { ThemeSwitcher } from "./buttons/themeSwitcher";
import { SearchBar } from "./ui/navbar/searchBar";
import { Logged, LoggedMoblie } from "./ui/navbar/logged";
import { NotLogged, NotLoggedMobile } from "./ui/navbar/notlogged";
import { useAuth } from "./providers/auth";
import { ExplorerButton } from "./buttons/explorer";


export const Navbar: FC = () => {
    const { user, loading } = useAuth();


    if (loading) return null;

    return !loading && (
        <div className="drawer z-10 backdrop-blur-lg fixed">
            <input id="my-drawer-3" type="checkbox" className="drawer-toggle" />
            <div className="drawer-content flex flex-col">
                {/* Navbar */}
                <div className="navbar w-full">
                    <div className="flex-none lg:hidden">
                        <label htmlFor="my-drawer-3" aria-label="open sidebar" className="btn btn-square btn-ghost">
                            <FontAwesomeIcon icon={faBars} className="inline-block h-5 w-5 stroke-current" />
                        </label>
                    </div>
                    <div className="mx-2 flex-1 px-2">
                        <Link href="/" className="btn btn-ghost md:text-2xl text-xl">
                            <div>
                                myanime<span className="text-indigo-600">.db</span>
                            </div>
                        </Link>
                    </div>
                    <div className="hidden flex-none lg:block">
                        <ul className="menu menu-horizontal">
                            {/* Navbar menu content here */}
                            <SearchBar />
                            <ExplorerButton />
                            {user ? <Logged user={user} /> : <NotLogged />}
                        </ul>
                    </div>
                    <ThemeSwitcher />
                </div>
            </div>
            <div className="drawer-side">
                <label htmlFor="my-drawer-3" aria-label="close sidebar" className="drawer-overlay"></label>
                <ul className="menu min-h-full bg-base-200 w-72 p-4">
                    {/* Sidebar content here */}
                    <div className="flex justify-between pb-5">
                        <div className="flex flex-col justify-start items-start">
                            <h1 className="text-2xl font-extrabold pt-2">
                                Menu<FontAwesomeIcon icon={faCompass} className="text-rose-300 ml-2" />
                            </h1>
                        </div>
                        <div className="flex flex-col justify-end items-end">
                            <label className="btn btn-circle btn-outline btn-sm drawer-button" aria-label="close sidebar" htmlFor="my-drawer-3">
                                <FontAwesomeIcon icon={faXmark} width={10} />
                            </label>
                        </div>
                    </div>
                    <SearchBar />
                    <ExplorerButton />
                    <div className="divider font-semibold">Account <FontAwesomeIcon icon={faUser} width={10} /></div>
                    {user ? <LoggedMoblie user={user} /> : <NotLoggedMobile />}
                </ul>
            </div>
        </div >
    )
}
