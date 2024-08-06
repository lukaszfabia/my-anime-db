import React, { FC, useRef } from "react";
import Link from "next/link";

import { faUser, faGear, faFilm, faArrowRightFromBracket, faTrash, faX } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Image from "next/image";
import { NavbarItem } from "@/types";
import { User } from "@/types/models";
import { useAuth } from "@/components/providers/auth";
import { toast } from "react-toastify";

const DiaglogBeforeRemoving: FC<{ removeAccount: () => void }> = ({ removeAccount }) => {
    const modalRef = useRef<HTMLDialogElement>(null);

    return (
        <div className="preview">
            <button onClick={() => modalRef.current?.showModal()} className="text-red-500">
                <FontAwesomeIcon icon={faTrash} width={10} height={10} className="mr-2" />
                <span>Delete account</span>
            </button>

            <dialog className="modal" ref={modalRef}>
                <div className="modal-box w-5/6 max-w-xl">
                    <h3 className="font-bold text-lg">Attention!</h3>
                    <p className="py-4">
                        Click the <b>Close</b> button to cancel the operation or click <b>Remove</b> to delete the account.
                        Please note that this action is <b>irreversible</b>.
                    </p>
                    <div className="modal-action flex items-center justify-center">
                        <form method="dialog">
                            <button className="btn btn-error mr-2" onClick={removeAccount}><FontAwesomeIcon icon={faTrash} />Remove</button>
                            <button className="btn btn-success ml-2" onClick={() => modalRef.current?.close()}><FontAwesomeIcon icon={faX} />Close</button>
                        </form>
                    </div>
                </div>
            </dialog>
        </div>
    );
}



export const Logged: FC<{ user: User }> = ({ user }) => {
    const { logout, removeAccount } = useAuth();

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
            <div className={`rounded-full flex items-center justify-center ${!user.isVerified && "indicator"}`}>
                <div tabIndex={0} role="button" className="btn btn-ghost btn-circle avatar">
                    {!user.isVerified && <span className="indicator-item badge badge-secondary w-5 pt-0.5 mt-2">!</span>}
                    <Image
                        src={user.picUrl}
                        alt={`${user.username}'s profile picture`}
                        width={50}
                        height={50}
                        className="rounded-full shadow-lg transition-opacity duration-300 ease-in-out group-hover:opacity-75"
                    />
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
                                    {elem.icon} {
                                        elem.name === "Settings" && !user.isVerified ? (
                                            <div className="indicator"><span className="mr-3">{elem.name}</span> <span className="indicator-item badge badge-secondary items-end w-5">!</span></div>
                                        ) : elem.name
                                    }
                                </Link>
                            </li>
                        </React.Fragment>
                    ))
                }
                <div className="divider"></div>
                <li>
                    <DiaglogBeforeRemoving removeAccount={removeAccount} />
                </li>
                <li className="text-warning">
                    <button onClick={handleLogout}>
                        <FontAwesomeIcon icon={faArrowRightFromBracket} width={10} height={10} />
                        <span>Logout</span>
                    </button>
                </li>
            </ul>
        </div >
    )
}