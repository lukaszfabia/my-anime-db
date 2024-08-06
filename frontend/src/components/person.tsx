import Link from "next/link";
import { FC } from "react";
import Image from "next/image"

export const Avatar: FC<{ picUrl: string, name: string, bio?: string }> = ({ picUrl, name, bio }) => (
    <div className="avatar flex justify-center items-center flex-col">
        <div className="ring-slate-500 ring-offset-base-100 w-3/4 rounded-full ring ring-offset-2">
            <Link href="/settings">
                <Image
                    src={picUrl}
                    alt={`${name}'s profile picture`}
                    width={250}
                    height={250}
                    priority={true}
                    className="tooltip tooltip-top"
                    data-tip="Change avatar"
                />
            </Link>
            <div className="absolute inset-0 bg-black opacity-0 group-hover:opacity-50 rounded-full transition-opacity duration-300 ease-in-out"></div>
        </div>
        <p className="text-center py-5">{bio}</p>
    </div>
)