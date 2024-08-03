"use client";

import { useAuth } from "@/components/providers/auth";
import { redirect } from "next/navigation";
import { FC } from "react";

const Profile: FC = () => {
    const { user } = useAuth();

    if (!user) {
        redirect("/login");
    }

    return (
        <div>
            <h1 className="text-3xl font-extrabold">
                Hello {user.username}
            </h1>
            <p>{user.bio}</p>
        </div>
    )
}

export default Profile;