import React, { FC } from "react";
import Link from "next/link";

import { faUser, faGear, faFilm, faArrowRightFromBracket } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

import { NavbarItem, User } from "@/types";
import { useAuth } from "@/components/providers/auth";
import { toast, Bounce } from "react-toastify";



export const Logged: FC<{ user: User }> = ({ user }) => {
    const { logout } = useAuth();

    const loggedLinks: NavbarItem[] = [
        { name: user.username, href: "/profile", icon: <FontAwesomeIcon icon={faUser} width={10} height={10} /> },
        { name: "Settings", href: "/settings", icon: <FontAwesomeIcon icon={faGear} width={10} height={10} /> },
        { name: "Collection", href: "/collection", icon: <FontAwesomeIcon icon={faFilm} width={10} height={10} /> },
    ];

    const handleLogout = () => {
        logout(() => toast.info('Logged out!'));
    }

    return (
        <div className="dropdown dropdown-end">
            <div tabIndex={0} role="button" className="btn btn-ghost btn-circle avatar">
                <div className="w-10 rounded-full flex items-center justify-center">
                    <img
                        alt="Tailwind CSS Navbar component"
                        src={user.picUrl!} />
                </div>
            </div>
            <ul
                tabIndex={0}
                className="menu menu-sm dropdown-content bg-base-100 rounded-box z-[1] mt-3 w-52 p-2 shadow">
                {
                    loggedLinks.map((elem: NavbarItem) => (
                        <React.Fragment key={elem.name}>
                            <li>
                                <Link href={elem.href}>
                                    {elem.icon} {elem.name}
                                </Link>
                            </li>
                        </React.Fragment>
                    ))
                }
                <div className="divider"></div>
                <li className="text-warning">
                    <button onClick={handleLogout}>
                        <FontAwesomeIcon icon={faArrowRightFromBracket} width={10} height={10} />
                        <span>Logout</span>
                    </button>
                </li>
            </ul>
        </div>
    )
}