"use client";

import { useAuth } from "@/components/providers/auth"
import { Spinner } from "@/components/ui/spinner";
import api from "@/lib/api";
import { FriendRequest, RequestStatus, User } from "@/types/models";
import { faCheck, faClock, faEnvelope, faUserGroup, faXmark } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { AxiosResponse } from "axios";
import { redirect } from "next/navigation";
import { FC, ReactNode, useEffect, useState } from "react";
import { toast } from "react-toastify";
import Image from "next/image";
import transformTime from "@/lib/computeTime";
import Link from "next/link";
import { DialogWindow } from "@/components/ui/dialog";
import { respondToFriendRequest, removeFriend } from "./manager";
import { getImageUrl } from "@/lib/getImageUrl";


const Invitation: FC<{ user: User, createdAt?: string, children?: ReactNode }> = ({ createdAt, user, children }) => {
    return (
        <div className="max-w-max md:p-4 md:mr-2 max-sm:mb-2">
            <div className="flex flex-col sm:flex-row flex-grow bg-base-200 p-6 rounded-2xl shadow-lg">
                <div className="avatar flex items-center justify-center mb-4 md:mb-0">
                    <div className="ring-primary ring-offset-base-100 w-20 rounded-full ring ring-offset-2">
                        <Link href={`/user/${user.id}`}>
                            {user.picUrl ? (
                                <Image
                                    src={getImageUrl(user.picUrl)}
                                    alt={user.username}
                                    width={100}
                                    height={100}
                                    className="rounded-full"
                                    key={user.username}
                                />
                            ) : (
                                <div className="placeholder avatar">
                                    <div className="bg-neutral text-neutral-content w-20 rounded-full">
                                        <span className="text-3xl">{user.username[0]}</span>
                                    </div>
                                </div>
                            )}
                        </Link>
                    </div>
                </div>
                <div className="sm:divider sm:divider-horizontal"></div>
                <div className="flex flex-col items-center md:items-start">
                    <h2 className="text-2xl font-bold text-center md:text-left text-blue-500 hover:underline">
                        <Link href={`/user/${user.id}`}>{user.username}</Link>
                    </h2>
                    {createdAt && <p className="text-center md:text-left">sent {createdAt}</p>}
                    <div className="flex flex-col md:flex-row items-center justify-center md:pt-2 pt-4 w-full md:w-auto">
                        {children}
                    </div>
                </div>
            </div>
        </div>
    )
}



export default function Friends() {
    const { user, loading, refreshUser } = useAuth();

    if (!user) {
        redirect("/login");
    }

    if (loading) return <Spinner />

    const [friendsRequests, setFriendsRequests] = useState<FriendRequest[]>([]);
    const [friends, setFriends] = useState<User[]>(user.friends ? user.friends : []);
    const [yourPendingRequests, setYourPendingRequests] = useState<FriendRequest[]>([]);

    const afterRespond = (curr: FriendRequest, action: RequestStatus) => {
        if (action === "accepted") {
            setFriends(prevFriends => [
                friendsRequests.find((request: FriendRequest) => request.id === curr.id)!.sender,
                ...prevFriends
            ]);
        }

        setYourPendingRequests(prevRequests => prevRequests.filter((request: FriendRequest) => request.id !== curr.id));
        setFriendsRequests(prevRequests =>
            prevRequests.filter((request: FriendRequest) => request.id !== curr.id)
        );

        refreshUser();
    }

    const handleRemove = (friendId: number) => {
        removeFriend(friendId, () => {
            setFriends(friends.filter((friend: User) => friend.id !== friendId));
            refreshUser();
        });
    }

    useEffect(() => {
        refreshUser();
        api.get<GoResponse>("/auth/friend/invitations/").
            then((res: AxiosResponse<GoResponse>) => {
                setFriendsRequests(res.data.data.pending || []);
                setYourPendingRequests(res.data.data.sent || []);
            }).
            catch((_: any) => toast.error("Failed to fetch friend requests"));
    }, []);

    const DisplaySentequests: FC = () => {
        return (
            <>
                {
                    yourPendingRequests.length > 0 && (
                        <div>
                            <h1 className="lg:text-4xl md:text-3xl text-2xl font-extrabold py-5">Sent <span className="text-rose-700">invites <FontAwesomeIcon icon={faEnvelope} /></span></h1>
                            {yourPendingRequests.map((request: FriendRequest) => (
                                <Invitation key={request.receiver.username} user={request.receiver} createdAt={transformTime(request.createdAt)}>
                                    <button className="btn btn-error btn-sm" onClick={() => respondToFriendRequest(request, "cancel", () => afterRespond(request, "cancel"))}><FontAwesomeIcon icon={faXmark} />Cancel</button>
                                </Invitation>
                            ))}
                        </div>
                    )
                }
            </>
        )
    }

    const DisplayPendingRequests: FC = () => {
        return (
            <>
                {
                    friendsRequests.length > 0 ? (
                        <div>
                            <h1 className="lg:text-4xl md:text-3xl text-2xl font-extrabold py-5">Pending <span className="text-orange-600">requests <FontAwesomeIcon icon={faClock} /></span></h1>
                            {friendsRequests.map((request: FriendRequest) => (
                                <Invitation user={request.sender} createdAt={transformTime(request.createdAt)} key={request.sender.username}>
                                    <button className="btn btn-success btn-sm" onClick={() => respondToFriendRequest(request, "accepted", () => afterRespond(request, "accepted"))}><FontAwesomeIcon icon={faCheck} />Accept</button>
                                    <div className="divider md:divider-horizontal">OR</div>
                                    <button className="btn btn-error btn-sm" onClick={() => respondToFriendRequest(request, "rejected", () => afterRespond(request, "rejected"))}><FontAwesomeIcon icon={faXmark} width={10} />Reject</button>
                                </Invitation>
                            ))}
                        </div>
                    ) :
                        <h1 className="mt-5 lg:text-4xl md:text-3xl text-2xl font-extrabold">No pending requests...</h1>
                }
            </>
        )
    }

    const YourFriendList: FC = () => {
        return (
            <div>
                {friends.length > 0 ? (
                    <h1 className="lg:text-4xl md:text-3xl text-2xl font-extrabold py-5">
                        Your{" "}<span className="text-green-600">friends <FontAwesomeIcon icon={faUserGroup} width={30} /></span>
                    </h1>
                ) : (
                    <h1 className="lg:text-4xl md:text-3xl text-2xl font-extrabold py-5">No <span className="text-lime-300">friends</span> yet 😓</h1>
                )}
                <div className="flex flex-row">
                    {friends.map((friend: User) => (
                        <Invitation user={friend} key={friend.username}>
                            <DialogWindow title=" Remove friend" handler={() => handleRemove(friend.id)} icon={<FontAwesomeIcon icon={faXmark} width={10} />}>
                                <span>Are you sure you want to remove <b>{friend.username}</b> from your friends list?</span>
                            </DialogWindow>
                        </Invitation>
                    ))}
                </div>
            </div>
        )
    }

    return (
        <div>
            <DisplaySentequests />
            <DisplayPendingRequests />
            <YourFriendList />
        </div>
    )
}
