"use client";

import { Dispatch, FC, FormEvent, SetStateAction, useState } from "react";
import { redirect } from "next/navigation";
import { useAuth } from "@/components/providers/auth";

import React from "react";
import api from "@/lib/api";
import { AxiosError, AxiosResponse } from "axios";
import { toast } from "react-toastify";
import { Overview, Statistics, FavAnime, RecentPosts } from "@/components/person";
import { Spinner } from "@/components/ui/spinner";
import { Post } from "@/types/models";
import { PostForm } from "@/components/ui/forms/postForm";



const CreatePost: FC<{ setPosts: Dispatch<SetStateAction<Post[]>>, posts: Post[] }> = ({ setPosts, posts }) => {

    const submitPost = (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        console.log("submitting post");
        const formData = new FormData(e.currentTarget);

        formData.get("isPublic") === "on" ? formData.set("isPublic", "true") : formData.set("isPublic", "false");

        api.post<Post>("/auth/post/", formData).then((response: AxiosResponse<Post>) => {
            setPosts([response.data, ...posts]);
            toast.success("Post created successfully");
        }).catch((err: AxiosError<GoResponse>) => {
            toast.error(err.response?.data.error)
        });
    }


    return <PostForm submitFunc={submitPost} />
};


const Profile: FC = () => {
    const { user, loading } = useAuth();
    const [posts, setPosts] = useState<Post[]>(user?.posts ? user.posts : []);

    if (!user) {
        redirect("/login");
    }

    if (loading) return <Spinner />


    return (
        <div>
            <div className="lg:flex">
                <div className="lg:w-1/3">
                    <Overview apiUser={user} />
                </div>
                <div className="lg:w-2/3">
                    <Statistics />
                    <FavAnime />
                </div>
            </div>
            <div className="lg:px-12 mt-10">
                <div className="divider"></div>
                <CreatePost setPosts={setPosts} posts={posts} />
                <RecentPosts user={user} posts={posts} setPosts={setPosts} />
            </div>
        </div>
    )
}

export default Profile;