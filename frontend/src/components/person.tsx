import Link from "next/link";
import { Dispatch, FC, FormEvent, SetStateAction, useEffect, useRef, useState } from "react";
import Image from "next/image"
import { User, Post, FriendRequest, RequestStatus } from "@/types/models";
import { faBook, faChartSimple, faCheck, faClock, faEdit, faEllipsis, faEnvelope, faFilm, faGlobe, faHeart, faPen, faPlus, faScrewdriverWrench, faUser, faXmark } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import React from "react";
import { PostWrapper } from "./post/wrapper";
import { Stat } from "@/types";
import api from "@/lib/api";
import { AxiosResponse } from "axios";
import { toast } from "react-toastify";
import { useAuth } from "./providers/auth";
import { DialogWindow } from "./ui/dialog";
import transformTime from "@/lib/computeTime";
import { PostForm } from "./ui/forms/postForm";
import { removeFriend, respondToFriendRequest } from "@/app/friends/manager";
import { getImageUrl } from "@/lib/getImageUrl";

export const Avatar: FC<{ picUrl?: string, name: string, bio?: string }> = ({ picUrl, name, bio }) => (
    <div className="avatar flex justify-center items-center flex-col">
        <div className="ring-slate-500 ring-offset-base-100 w-3/4 rounded-full ring ring-offset-2 avatar placeholder">
            {picUrl ? (
                <>
                    <Image
                        src={picUrl}
                        alt={`${name}'s profile picture`}
                        sizes="300px"
                        width={100}
                        priority
                        height={100}
                        key={name}
                    />
                </>
            ) :
                <div className="avatar placeholder">
                    <div className="bg-neutral text-neutral-content w-[300px] rounded-full">
                        <span className="text-9xl">{name[0]}</span>
                    </div>
                </div>
            }
        </div>
        <p className="text-center py-5">{bio}</p>
    </div>
)

const ManagePost: FC<{ id: number, posts: Post[], setPosts: Dispatch<SetStateAction<Post[]>> }> = ({ id, posts, setPosts }) => {

    const handleDelete = () => {
        api.delete(`/auth/post/${id}`)
            .then(() => {
                setPosts(posts.filter((post: Post) => post.id !== id));
            }).catch((_: any) => toast.error("Failed to delete post"));
    }

    const handleEdit = (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const formData = new FormData(e.currentTarget);
        formData.get("isPublic") === "on" ? formData.set("isPublic", "true") : formData.set("isPublic", "false");

        api.put<GoResponse>(`/auth/post/${id}`, formData)
            .then((res: AxiosResponse<GoResponse>) => {
                const editedPost: Post = res.data.data;
                setPosts(posts.map((post: Post) => post.id === id ? editedPost : post));
                ref.current?.close();
            }).catch((_: any) => toast.error("Failed to edit post") && ref.current?.close());
    }

    const ref = useRef<HTMLDialogElement>(null);

    return (
        <div className="absolute top-4 right-4">
            <div className="dropdown dropdown-end">
                <label tabIndex={0} className="btn btn-ghost rounded-btn">
                    <FontAwesomeIcon icon={faEllipsis} className="text-xl" />
                </label>
                <ul tabIndex={0} className="dropdown-content menu p-2 shadow-lg bg-base-100 rounded-box w-52">
                    <li>
                        <DialogWindow wantButton modalRef={ref} title="Edit post" icon={<FontAwesomeIcon icon={faEdit} className="mr-2" width={12} />}>
                            <PostForm defaultValues={posts.find((post: Post) => post.id === id)} submitFunc={handleEdit} />
                        </DialogWindow>
                    </li>
                    <li>
                        <DialogWindow actionOrClose short title="Delete post" handler={handleDelete} errorColor wantButton >
                            <span>
                                Are you sure you want to delete this post?
                            </span>
                        </DialogWindow>
                    </li>
                </ul>
            </div>
        </div>
    )
}

export const RecentPosts: FC<{ user: User, posts?: Post[], setPosts?: Dispatch<SetStateAction<Post[]>>, isReadOnly?: boolean }> = ({ user, posts = user.posts, setPosts, isReadOnly = false }) => {
    return (
        posts && (
            <div>
                <h1 className="text-4xl py-5 text-center md:text-left font-extrabold">
                    Recent <span className="text-fuchsia-400">posts <FontAwesomeIcon icon={faPen} width={30} /></span>
                </h1>
                <>
                    {posts.map((post: Post) => (
                        <React.Fragment key={post.id}>
                            {isReadOnly && post.isPublic ?
                                <PostWrapper post={post} user={user} isReadOnly={isReadOnly} /> :
                                !isReadOnly &&
                                (
                                    <PostWrapper post={post} user={user} >
                                        {/* when we are on profile we know that we can modify posts */}
                                        <ManagePost id={post.id} posts={posts} setPosts={!isReadOnly && setPosts!} />
                                    </PostWrapper>
                                )
                            }
                        </React.Fragment>
                    ))}
                </>
            </div>
        )
    )
}

export const Statistics: FC = () => {
    const stats: Stat[] = [
        { title: "Watched hours", value: "1337 H", desc: "your total time spent watching", icon: <FontAwesomeIcon icon={faClock} /> },
        { title: "Reviews", value: "13", desc: "number of left reviews <3", icon: <FontAwesomeIcon icon={faBook} /> },
        { title: "Fav genre", value: "Drama", desc: "based on your completed anime", icon: <FontAwesomeIcon icon={faFilm} /> },
    ]

    return (
        <div className="flex flex-col">
            <h1 className="text-4xl text-center md:text-left font-extrabold py-5">
                Your <span className="text-violet-500">statistics <FontAwesomeIcon icon={faChartSimple} width={30} /></span>
            </h1>
            <div className="flex flex-col items-center lg:flex-row lg:items-start lg:justify-between shadow p-3">
                {stats.map((stat) => (
                    <div className="stat flex flex-col items-center text-center lg:text-left" key={stat.title}>
                        <div className="stat-title font-semibold">{stat.title}</div>
                        <div className="stat-value py-2 dark:text-black text-white max-md:text-3xl">
                            <span className="mr-2">{stat.value}</span>{stat.icon}
                        </div>
                        <div className="stat-desc">{stat.desc}</div>
                    </div>
                ))}
            </div>
        </div>
    )
}


export const FavAnime: FC = () => {
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


export const Overview: FC<{ apiUser: User, isReadOnly?: boolean }> = ({ apiUser, isReadOnly = false }) => {
    const { user, refreshUser } = useAuth();

    const [buttonText, setButtonText] = useState<RequestStatus>("rejected");

    const [friendRequest, setFriendsRequest] = useState<FriendRequest | null>(null);

    useEffect(() => {
        if (user && isReadOnly) {
            api.get<GoResponse>(`/auth/friend/state/?sender=${user.id}&receiver=${apiUser.id}`).
                then((res: AxiosResponse<GoResponse>) => {
                    setFriendsRequest(res.data.data);
                    user.id === res.data.data.receiverId && res.data.data.status === "pending" ? setButtonText("respond") : setButtonText(res.data.data.status);
                }).
                catch((_: any) => setButtonText("rejected"));
        }
    }, [buttonText]);


    const BuildButton: FC = () => {
        if (!user) return null; // if am not logged in, don't show any button


        const sendFriendRequest = () => {
            api.post<GoResponse>(`/auth/friend/${apiUser.id}`).then(() => {
                setButtonText("pending");
            }).catch((_: any) => toast.error("Something went wrong"));
        };

        if (buttonText === "pending") {
            return (
                <button onClick={() => respondToFriendRequest(friendRequest, "cancel", () => setButtonText("cancel"))} className="btn btn-outline btn-info lg:my-1 my-3 lg:w-full h-1/6">
                    <span>Pending</span>
                    <FontAwesomeIcon icon={faClock} />
                </button>
            )
        }

        else if (buttonText === "accepted") {
            return (
                <button onClick={() => removeFriend(apiUser.id, () => { setButtonText("rejected"); refreshUser(); })} className="btn btn-outline btn-error lg:my-1 my-3 lg:w-full h-1/6">
                    <span>Remove from friends</span>
                    <FontAwesomeIcon icon={faXmark} />
                </button>
            )
        }

        else if (buttonText === "respond") {
            return (
                <>
                    <h1 className="text-center my-2 text-info"><b>{apiUser.username}</b> sent invitation to you!</h1>

                    <div className="flex flex-row items-center justify-center">
                        <button className="btn btn-sm btn-outline  btn-success" onClick={() => respondToFriendRequest(friendRequest, "accepted", () => { setButtonText("accepted"); refreshUser() })}>
                            <FontAwesomeIcon icon={faCheck} width={10} />
                            <span>Accept</span>
                        </button>
                        <div className="divider divider-horizontal"></div>
                        <button className="btn btn-outline btn-sm btn-error" onClick={() => respondToFriendRequest(friendRequest, "rejected", () => { setButtonText("rejected"); refreshUser(); })}>
                            <FontAwesomeIcon icon={faXmark} width={10} />
                            <span>Reject</span>
                        </button>
                    </div>
                </>
            )
        }

        else {
            return (
                <button onClick={sendFriendRequest} className="btn btn-outline btn-info lg:my-1 my-3 lg:w-full h-1/6">
                    <span>Add to friends</span>
                    <FontAwesomeIcon icon={faPlus} />
                </button>
            )
        }
    }

    return (
        <div>
            <div className="flex items-center justify-center">
                <Avatar picUrl={apiUser.picUrl && getImageUrl(apiUser.picUrl)} name={apiUser.username} bio={apiUser.bio} />
            </div>
            <div className="md:px-12 max-lg:text-center rounded-lg">
                <h1 className="text-4xl md:text-5xl font-bold dark:text-black text-white">
                    {apiUser.username}
                    {apiUser.isMod && (
                        <FontAwesomeIcon icon={faScrewdriverWrench} className="ml-2" width={35} />
                    )}
                </h1>
                <h1 className="md:text-lg lg:my-1 my-3">
                    <FontAwesomeIcon icon={faEnvelope} className="mr-1" />
                    {apiUser.email}
                </h1>
                {apiUser.website && (
                    <h1 className="md:text-lg lg:my-1 my-3">
                        <FontAwesomeIcon icon={faGlobe} className="mr-1" />
                        <Link href={apiUser.website} target="_blank" className="transition-all ease-in-out duration-200 text-blue-500 hover:text-blue-600">
                            {apiUser.website.split("https://")[1]}
                        </Link>
                    </h1>
                )}
                <h1 className="md:text-lg my-1">
                    <FontAwesomeIcon icon={faUser} className="mr-1" />
                    Joined {" "}
                    <span className="font-semibold">{transformTime(apiUser.createdAt)}</span>
                </h1>
                {!isReadOnly && <Link href="/settings" className="btn btn-outline lg:my-1 my-3 lg:w-full h-1/6">Edit profile</Link>}
                {isReadOnly && <BuildButton />}
            </div>
        </div>
    );
}