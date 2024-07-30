import Link from "next/link";
import { FC } from "react";
import { ThemeSwitcher } from "./buttons/themeSwitcher";
import { NavbarItem } from "@/types";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faUser, faGear, faFilm, faRightFromBracket, faLightbulb, faMagnifyingGlass, faSignIn, faUserPlus } from "@fortawesome/free-solid-svg-icons";
import React from "react";


export const loggedLinks: NavbarItem[] = [
    { name: "Profile", href: "/profile", icon: <FontAwesomeIcon icon={faUser} width={10} height={10} /> },
    { name: "Settings", href: "/settings", icon: <FontAwesomeIcon icon={faGear} width={10} height={10} /> },
    { name: "Collection", href: "/collection", icon: <FontAwesomeIcon icon={faFilm} width={10} height={10} /> },
    { name: "Logout", href: "/logout", icon: <FontAwesomeIcon icon={faRightFromBracket} width={10} height={10} /> },
];

const Logged: FC = () => {
    return (
        <>
            {
                loggedLinks.map((elem: NavbarItem, i: number) => (
                    <React.Fragment key={i}>
                        {elem.name === "Logout" ? <div className="divider"></div> : null}
                        <li>
                            <Link href={elem.href}>
                                {elem.icon} {elem.name}
                            </Link>
                        </li>
                    </React.Fragment>
                ))
            }
        </>
    )
}

const NotLogged: FC = () => {
    return (
        <>
            <div className="hidden md:visible md:flex md:w-full">
                <div className="pt-3">
                    <button className="btn btn-outline btn-sm flex items-center gap-2">
                        <Link href="/login" className="flex items-center gap-2">
                            <FontAwesomeIcon icon={faSignIn} width={15} height={15} />
                            Login
                        </Link>
                    </button>
                </div>
                <div className="divider divider-horizontal">OR</div>
                <div className="pt-3">
                    <button className="btn btn-accent btn-sm flex items-center gap-2">
                        <Link href="/signup" className="flex items-center gap-2">
                            <FontAwesomeIcon icon={faUserPlus} width={15} height={15} />
                            Sign up
                        </Link>
                    </button>
                </div>
            </div>

            <div className="md:hidden">
                <div className="dropdown dropdown-end">
                    <div tabIndex={0} role="button" className="btn m-1"><FontAwesomeIcon icon={faUser} width={15} height={15} /></div>
                    <ul tabIndex={0} className="dropdown-content menu bg-base-100 rounded-box z-[1] w-52 p-2 shadow">
                        <button className="btn btn-outline btn-sm flex items-center gap-2">
                            <Link href="/login" className="flex items-center gap-2">
                                <FontAwesomeIcon icon={faSignIn} width={15} height={15} />
                                Login
                            </Link>
                        </button>
                        <div className="divider"></div>
                        <button className="btn btn-accent btn-sm flex items-center gap-2">
                            <Link href="/signup" className="flex items-center gap-2">
                                <FontAwesomeIcon icon={faUserPlus} width={15} height={15} />
                                Sign up
                            </Link>
                        </button>
                    </ul>
                </div>
            </div>
        </>
    );
};

const SearchBar: FC = () => {
    return (
        <label className="hidden md:visible input input-bordered md:input-md md:flex items-center gap-2">
            <FontAwesomeIcon icon={faMagnifyingGlass} width={15} height={15} />
            <input type="text" className="grow" placeholder="Search..." autoComplete="off" />
        </label>
    )
}

const MenuProfile: FC = () => {
    return (
        <div className="dropdown dropdown-end">
            <div tabIndex={0} role="button" className="btn btn-ghost btn-circle avatar">
                <div className="w-7 rounded-full">
                    <img
                        alt="Tailwind CSS Navbar component"
                        src="https://img.daisyui.com/images/stock/photo-1534528741775-53994a69daeb.webp" />
                </div>
            </div>
            <ul
                tabIndex={0}
                className="menu menu-sm dropdown-content bg-base-100 rounded-box z-[1] mt-3 w-52 p-2 shadow">
                <Logged />
            </ul>
        </div>
    )
}

export const Navbar: FC = () => {
    return (
        <div className="navbar fixed z-10">
            <div className="flex-1">
                <Link href="/" className="btn btn-ghost md:text-2xl text-xl">
                    <div>
                        myanime<span className="text-indigo-600">.db</span>
                    </div>
                </Link>
            </div>
            <div className="flex-none gap-2">
                <SearchBar />
                <Link className="btn btn-ghost" href="/explorer"><FontAwesomeIcon icon={faLightbulb} width={15} height={15} /> Explorer</Link>

                {/* profile */}
                <MenuProfile />
                {/* <NotLogged /> */}
            </div>
            <div className="px-5">
                <ThemeSwitcher />
            </div>
        </div>
    );
};
