import api from "@/lib/api";
import { transformTime } from "@/lib/computeTime";
import { getImageUrl } from "@/lib/getImageUrl";
import { Anime, User, UserAnime } from "@/types/models";
import { GoResponse } from "@/types/responses";
import { faBook, faHeart, faPaperPlane } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Image from "next/image";
import Link from "next/link";
import { FC } from "react";
import { toast } from "react-toastify";

const Review: FC<{ r: UserAnime, my: boolean }> = ({ r, my }) => {
    return (
        <div>
            {my && (
                <div className="label">
                    <span className="label-text">Your review</span>
                </div>
            )}
            <div className={`flex bg-base-200 rounded-2xl md:p-5 p-4 border ${my ? "border-primary" : "border-gray-600"}`}>
                <div>
                    <div className="flex flex-row w-full">
                        <div className="avatar placeholder">
                            <div className="ring-gray-500 ring-offset-base-100 md:w-20 md:h-20 w-14 h-14 rounded-full ring ring-offset-2">
                                {r.user.picUrl ? (
                                    <Image alt={`${r.user.username}'s avatar`} src={getImageUrl(r.user.picUrl)} width={150} height={150} />
                                ) : <span className="text-3xl">{r.user.username.charAt(0).toUpperCase()}</span>}
                            </div>
                        </div>
                        <div className="md:divider md:divider-horizontal max-sm:px-2"></div>
                        <div className="md:flex-col w-full">
                            <div className="flex md:flex-row flex-col md:divide-x md:divide-gray-400">
                                <h1 className="md:pr-2 font-bold text-blue-500 hover:underline">
                                    <Link href={`/user/${r.userId}`}>{r.user.username}</Link>
                                </h1>

                                <p className="md:px-2">{r.score} {r.isFav && (
                                    <FontAwesomeIcon icon={faHeart} className="text-red-500" />)}
                                </p>
                            </div>

                            <p className="text-sm text-gray-500">{transformTime(r.updatedAt)}</p>
                            <div className="max-sm:hidden">
                                {r.review}
                            </div>
                        </div>
                    </div>
                    <div className="md:hidden md:mt-2">
                        {r.review}
                    </div>
                </div>
            </div>
        </div>
    )
}

export const DisplayReviews: FC<{ fixedReviews: UserAnime[], user: User | null }> = ({ fixedReviews, user }) => {

    return fixedReviews.length > 0 && (
        <div className="p-5 bg-base-300 shadow-xl rounded-xl">
            <h1 className="text-4xl font-extrabold">Reviews <FontAwesomeIcon icon={faBook} /></h1>
            <div className="grid grid-cols-1 gap-4 mt-5">
                {fixedReviews.map((r: UserAnime) => (
                    <Review key={r.id} r={r} my={user ? r.user.username === user.username : false} />
                ))}
            </div>
        </div>
    )
}

interface ReviewProps {
    user: User;
    anime: Anime;
    hasCompleted: boolean;
    refreshUser: () => void;
}

export const CreateReview: FC<ReviewProps> = ({ anime, hasCompleted, refreshUser }) => {

    const addReview = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const form = new FormData(e.currentTarget);

        api.put<GoResponse>(`/auth/anime/${anime.id}/review/`, form).then((res) => {
            if (res.data.code === 200) {
                refreshUser();
            } else {
                toast.error("Failed to add review");
            }
        }).catch((_: any) => toast.error("Something went wrong"));

        e.currentTarget.reset();
    }

    return (
        <form className="p-5 shadow" encType="multipart/form-data" onSubmit={addReview}>
            <label className="form-control">
                <div className="label">
                    <span className="label-text">Your review</span>
                </div>
                <textarea name="review" id="review" className="textarea textarea-bordered h-24" placeholder={`Write here...`} disabled={!hasCompleted}></textarea>
            </label>

            <div className="flex justify-end items-end mt-5"><button type="submit" className="btn btn-success sm:w-fit w-full" disabled={!hasCompleted}>
                <FontAwesomeIcon icon={faPaperPlane} />
                Post
            </button></div>
        </form>
    )
}