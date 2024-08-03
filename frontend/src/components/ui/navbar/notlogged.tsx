import Link from "next/link";
import { FC } from "react";

import { faSignIn, faUserPlus, faUser } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

export const NotLogged: FC = () => {
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
