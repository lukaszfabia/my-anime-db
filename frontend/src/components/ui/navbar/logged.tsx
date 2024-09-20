import React, { FC } from "react";
import Link from "next/link";

import { faUser, faGear, faFilm, faArrowRightFromBracket, faUserGroup, faPlus, faDatabase, faScrewdriverWrench, faDoorOpen } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Image from "next/image";
import { NavbarItem } from "@/types";
import { User } from "@/types/models";
import { useAuth } from "@/components/providers/auth";
import { toast } from "react-toastify";
import { DialogWindow } from "../dialog";
import { getImageUrl } from "@/lib/getImageUrl";

const routes = (username: string): NavbarItem[] => [
    { name: username, href: "/profile", icon: <FontAwesomeIcon icon={faUser} width={10} /> },
    { name: "Settings", href: "/settings", icon: <FontAwesomeIcon icon={faGear} width={10} /> },
    { name: "Collection", href: "/profile/collection", icon: <FontAwesomeIcon icon={faFilm} width={10} /> },
    { name: "Friends", href: "/friends", icon: <FontAwesomeIcon icon={faUserGroup} width={10} /> },
]


export const LoggedMoblie: FC<{ user: User }> = ({ user }) => {
    const loggedLinks = routes(user.username);

    return (
        <div>
            <div className="flex justify-between">
                <div className="flex flex-row pl-5">
                    <Link href="/profile">
                        <MiniAvatar user={user} />
                    </Link>
                    <div className="px-2"></div>
                    <div className="flex flex-col">
                        <h1 className="text-lg font-semibold btn btn-sm animate-pulse text-black dark:text-white"><Link href="\profile">{user.username}</Link></h1>
                        <p className="text-sm">{user.email}</p>
                    </div>
                </div>
            </div>
            <LoggedLinks user={user} elemToSkip={loggedLinks[0]} />
            <ModeratorZone user={user} />
            <LoggedActions />
        </div>
    )
}

const MiniAvatar: FC<{ user: User }> = ({ user }) => {
    return (
        <div className={`rounded-full flex items-center justify-center ${!user.isVerified && "indicator"}`}>
            <div tabIndex={0} role="button" className="btn btn-ghost btn-circle avatar placeholder">
                {!user.isVerified && <span className="indicator-item badge badge-secondary w-5 pt-0.5 mt-2">!</span>}
                {user.picUrl ?
                    <>
                        <Image
                            src={getImageUrl(user.picUrl)}
                            key={user.username}
                            alt={`${user.username}'s profile picture`}
                            width={50}
                            height={50}
                            className="rounded-full shadow-lg transition-opacity duration-300 ease-in-out group-hover:opacity-75"
                        />
                    </> :
                    <div className="avatar placeholder">
                        <div className="w-12 bg-neutral text-neutral-content rounded-full shadow-lg transition-opacity duration-300 ease-in-out group-hover:opacity-75">
                            <span className="text-lg">{user.username[0]}</span>
                        </div>
                    </div>
                }
            </div>
        </div>
    )
}

export const Logged: FC<{ user: User }> = ({ user }) => {

    return (
        <div className="dropdown dropdown-end">
            <MiniAvatar user={user} />
            <ul
                tabIndex={0}
                className="menu menu-sm dropdown-content bg-base-100 rounded-box z-[1] mt-3 w-52 p-2 shadow">
                <LoggedLinks user={user} />
                <ModeratorZone user={user} />
                <LoggedActions />
            </ul>
        </div>
    )
}

const LoggedLinks: FC<{ user: User, elemToSkip?: NavbarItem }> = ({ user, elemToSkip }) => {
    const loggedLinks = routes(user.username);

    return (
        <div className="lg:py-0 py-3">
            {loggedLinks.map((elem: NavbarItem) => (
                (!elemToSkip || elemToSkip.name !== elem.name) && (
                    <React.Fragment key={elem.name}>
                        <li>
                            <Link href={elem.href}>
                                {elem.icon} {elem.name}
                            </Link>
                        </li>
                    </React.Fragment>
                )
            ))}
        </div>
    )
}

const LoggedActions: FC = () => {
    const { logout, removeAccount } = useAuth();
    const handleLogout = () => {
        logout(() => toast.info('Logged out!'));
    }

    return (
        <>
            <div className="divider">
                <span>Actions <FontAwesomeIcon icon={faDoorOpen} width={15} className="pt-1 pl-1" /></span>
            </div>
            <li>
                <DialogWindow actionOrClose title="Delete account" handler={removeAccount} short wantButton errorColor>
                    Click the <b>Close</b> button to cancel the operation or click <b>Remove</b> to delete the account.
                    Please note that this action is <b>irreversible</b>.
                </DialogWindow>
            </li>
            <li className="text-warning">
                <button onClick={handleLogout}>
                    <FontAwesomeIcon icon={faArrowRightFromBracket} width={10} />
                    <span>Logout</span>
                </button>
            </li>
        </>
    )
}

const ModeratorZone: FC<{ user: User }> = ({ user }) => {
    if (!user.isMod) return null;
    else {
        const moderatorLinks: NavbarItem[] = [
            { name: "Manage content", href: "/manage", icon: <FontAwesomeIcon icon={faDatabase} width={10} /> },
            { name: "Post global info", href: "/manage/post-info", icon: <FontAwesomeIcon icon={faPlus} width={10} /> },
        ]

        return (
            <>
                <div className="divider">
                    <span>Moderator zone <FontAwesomeIcon icon={faScrewdriverWrench} width={15} className="pt-1 pl-1" /></span>
                </div>
                {moderatorLinks.map((elem: NavbarItem) => (
                    <React.Fragment key={elem.name}>
                        <li>
                            <Link href={elem.href}>
                                {elem.icon} {elem.name}
                            </Link>
                        </li>
                    </React.Fragment>
                ))}
            </>
        )
    }
}