import Link from "next/link";
import { FC } from "react";

import { faSignIn, faUserPlus } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";

export const NotLoggedMobile: FC = () => {
    return (
        <div className="flex items-center justify-center p-3">
            <div className="divider"></div>
            <button className="btn btn-outline btn-sm flex items-center gap-2">
                <Link href="/login" className="flex items-center gap-2">
                    <FontAwesomeIcon icon={faSignIn} width={15} height={15} />
                    Login
                </Link>
            </button>
            <div className="divider divider-horizontal">OR</div>
            <button className="btn btn-accent btn-sm flex items-center gap-2">
                <Link href="/signup" className="flex items-center gap-2">
                    <FontAwesomeIcon icon={faUserPlus} width={15} height={15} />
                    Sign up
                </Link>
            </button>
        </div>
    )
}

export const NotLogged: FC = () => {
    return (
        <div className="flex justify-center items-center px-2">
            <button className="btn btn-outline btn-sm flex items-center gap-2">
                <Link href="/login" className="flex items-center gap-2">
                    <FontAwesomeIcon icon={faSignIn} width={15} height={15} />
                    Login
                </Link>
            </button>
            <div className="divider divider-horizontal">OR</div>
            <button className="btn btn-accent btn-sm flex items-center gap-2">
                <Link href="/signup" className="flex items-center gap-2">
                    <FontAwesomeIcon icon={faUserPlus} width={15} height={15} />
                    Sign up
                </Link>
            </button>
        </div>
    );
};
