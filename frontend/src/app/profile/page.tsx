"use client";

import { FC, ReactNode } from "react";
import { redirect } from "next/navigation";
import { useAuth } from "@/components/providers/auth";
import Link from "next/link";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faBook, faChartSimple, faCircleUser, faClock, faEnvelope, faGlobe, faPen, faScrewdriverWrench, faUser, faHeart } from "@fortawesome/free-solid-svg-icons";

import { User } from "@/types/models";
import { format } from "date-fns";
import React from "react";
import { Spinner } from "@/components/ui/spinner";
import { Avatar } from "@/components/person";

interface Stat {
    title: string;
    value: string;
    desc?: string;
    icon: ReactNode;
}


const Overview: FC<{ user: User }> = ({ user }) => {
    const formattedDate = format(new Date(user.CreatedAt), 'dd.MM.yyyy');
    return (
        <div>
            <div className="flex items-center justify-center">
                <Avatar picUrl={user.picUrl} name={user.username} bio={user.bio} />
            </div>
            <div className="md:px-14 max-md:text-center rounded-lg">
                <h1 className="md:text-4xl text-5xl font-bold dark:text-black text-white">
                    {user.username}
                    {!user.isMod && (
                        <div className="tooltip tooltip-right font-normal" data-tip="Moderator">
                            <FontAwesomeIcon icon={faScrewdriverWrench} className="ml-2 w-2/3" />
                        </div>
                    )}
                </h1>
                <h1 className="md:text-lg lg:my-1 my-3">
                    <FontAwesomeIcon icon={faEnvelope} className="mr-1" />
                    {user.email}
                </h1>
                {user.website !== "" && (
                    <h1 className="md:text-lg lg:my-1 my-3">
                        <FontAwesomeIcon icon={faGlobe} className="mr-1" />
                        <Link href={user.website} target="_blank" className="transition-all ease-in-out duration-200 text-blue-500 hover:text-blue-600">
                            {user.website.split("https://")[1]}
                        </Link>
                    </h1>
                )}
                <h1 className="md:text-lg my-1">
                    <FontAwesomeIcon icon={faUser} className="mr-1" />
                    Joined in{" "}
                    <span className="font-semibold">{formattedDate}</span>
                </h1>
                <Link href="/settings" className="btn btn-outline lg:my-1 my-3 lg:w-full h-1/6">Edit profile</Link>
            </div>
        </div>
    );
}


const Statistics: FC = () => {
    const stats: Stat[] = [
        { title: "Watched hours", value: "1337 H", desc: "your total time spent watching", icon: <FontAwesomeIcon icon={faClock} /> },
        { title: "Posts", value: "43", desc: "number of published posts", icon: <FontAwesomeIcon icon={faPen} /> },
        { title: "Reviews", value: "13", desc: "number of left reviews <3", icon: <FontAwesomeIcon icon={faBook} /> },
        { title: "Account age", value: "2 years", icon: <FontAwesomeIcon icon={faCircleUser} /> },
    ]

    return (
        <div className="flex flex-col">
            <h1 className="text-4xl text-center md:text-left font-extrabold py-5">
                Your <span className="text-violet-500">statistics <FontAwesomeIcon icon={faChartSimple} width={30} /></span>
            </h1>
            <div className="flex flex-col items-center lg:flex-row lg:items-start lg:justify-between shadow p-4">
                {stats.map((stat) => (
                    <div className="stat flex flex-col items-center text-center lg:text-left" key={stat.title}>
                        <div className="stat-title">{stat.title}</div>
                        <div className="stat-value py-2 dark:text-black text-white">
                            <span className="mr-2">{stat.value}</span>{stat.icon}
                        </div>
                        <div className="stat-desc">{stat.desc}</div>
                    </div>
                ))}
            </div>
        </div>
    )
}


const FavAnime: FC = () => {
    return (
        <div className="">
            {/* display 3 posts */}
            <h1 className="text-4xl py-5 text-center md:text-left font-extrabold">
                Fav <span className="text-rose-600">anime <FontAwesomeIcon icon={faHeart} width={30} /></span>
            </h1>
            <div>
                Lorem ipsum dolor sit amet consectetur adipisicing elit. Ex laudantium aperiam itaque saepe neque iure eligendi, earum, maxime consequatur tenetur veritatis magni voluptatibus nam aliquam laborum maiores impedit quam? Consequuntur.
            </div>
        </div>
    )
}

const RecentPosts: FC = () => {
    return (
        <div className="">
            {/* display 3 posts */}
            <h1 className="text-4xl py-5 text-center md:text-right font-extrabold">
                Recent <span className="text-fuchsia-400">posts <FontAwesomeIcon icon={faPen} width={30} /></span>
            </h1>
            <div>
                Lorem ipsum dolor sit, amet consectetur adipisicing elit. Optio expedita perspiciatis aut sit repellat blanditiis eum excepturi natus? Dicta modi autem neque. Vero sapiente voluptatum distinctio itaque aliquid quod consectetur.
            </div>
        </div>
    )
}



const Profile: FC = () => {
    const { user, loading } = useAuth();

    if (!user) {
        redirect("/login");
    }

    if (loading) return <Spinner />


    return (
        <div className="md:flex">
            <div className="lg:w-1/3">
                <Overview user={user} />
            </div>
            <div className="px-5"></div>
            <div className="lg:w-2/3">
                <Statistics />
                <RecentPosts />
                <FavAnime />
            </div>
        </div>
    )
}

export default Profile;