"use client";

import { useAuth } from "@/components/providers/auth";
import { Spinner } from "@/components/ui/spinner";
import React from "react";
import { ButtonWithBackgroundPicProps, Menu } from "@/components/manage/menu";

export default function Manage() {
    const { user, loading } = useAuth();

    if (!user || loading) return <Spinner />;

    if (!user.isMod) return null;

    const props: ButtonWithBackgroundPicProps[] = [
        { imageUrl: "/images/anya.jpg", link: "/manage/post-info", title: "Post info", content: "Create an annoucment on welcome page!", },
        { imageUrl: "/images/aot.webp", link: "/manage/add-anime", title: "Create new Anime", content: "Add new Anime to our db.", btnText: "Add anime" },
        { imageUrl: "/images/mono.jpg", link: "/manage/add-others", title: "Add others", content: "Create voice actors or characters", btnText: "Add others" },
    ]

    return (
        <Menu cards={props}>
            <h1 className="text-center text-5xl font-extrabold">Moderator zone</h1>
            <p className="text-sm text-center text-gray-500 py-4">You can change here some data. Please choose from options below.</p>
        </Menu>
    );
}
