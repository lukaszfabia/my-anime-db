"use client";

import { FavAnime, Overview, RecentPosts, Statistics } from "@/components/person";
import { useAuth } from "@/components/providers/auth";
import { Spinner } from "@/components/ui/spinner";
import api from "@/lib/api";
import { User } from "@/types/models";
import { AxiosError, AxiosResponse } from "axios";
import { redirect, useParams } from "next/navigation";
import { FC, useEffect, useState } from "react";

const ReadOnlyUser: FC = () => {
    const { id } = useParams();
    const [readOnlyUser, setReadOnlyUser] = useState<User | null>(null);
    const [loading, setLoading] = useState<boolean>(true);
    const { user } = useAuth();

    useEffect(() => {
        api.get<User>(`/user/${id}`).then((res: AxiosResponse<User>) => {
            setLoading(false);
            setReadOnlyUser(res.data);
        }).catch((err: AxiosError<GoResponse>) => {
            setLoading(false);
            console.error(err);
        });
    }, []);

    if (user && readOnlyUser && user.id === readOnlyUser.id) {
        redirect("/profile");
    }

    if (loading) {
        return <Spinner />;
    }


    return (
        <>
            <div className="lg:flex">
                <div className="lg:w-1/3">
                    {readOnlyUser && <Overview apiUser={readOnlyUser} isReadOnly />}
                </div>
                <div className="px-5"></div>
                <div className="lg:w-2/3">
                    <Statistics />
                    <FavAnime />
                </div>
            </div>
            <div className="lg:px-12 mt-10">
                <div className="divider"></div>
                {readOnlyUser && <RecentPosts user={readOnlyUser} isReadOnly />}
            </div>
        </>
    )
}

export default ReadOnlyUser;