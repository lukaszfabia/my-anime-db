"use client";

import { OverviewForAnime } from "@/components/anime/animeOverview";
import { useAuth } from "@/components/providers/auth"
import { Spinner } from "@/components/ui/spinner";
import api from "@/lib/api";

import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Link from "next/link";
import { useParams } from "next/navigation";
import { Dispatch, FC, SetStateAction, useEffect, useState } from "react";
import { Info } from "@/components/anime/info";
import { CreateReview, DisplayReviews } from "@/components/anime/review";
import { GoResponse } from "@/types/responses";
import { Anime, Review } from "@/types/models";
import { faAdd, faArrowsRotate, faEdit, faHeart, faHeartBroken, } from "@fortawesome/free-solid-svg-icons";
import { Selector } from "@/components/ui/selector";
import { toast } from "react-toastify";


export default function AnimePage() {
    const { user, loading, refreshUser } = useAuth();
    const { id } = useParams();

    const [anime, setAnime] = useState<Anime | null>(null);
    const [hasCompleted, setHasCompleted] = useState<boolean>(false);

    const [userAnimeState, setUserAnimeState] = useState({
        isFavorite: false,
        watchStatus: "",
        score: "",
        btnText: "Add to list"
    });

    const [watchStatuses, setWatchStatuses] = useState<string[]>([]);
    const [scores, setScores] = useState<string[]>([]);
    const [fixedReviews, setFixedReviews] = useState<Review[]>([]);

    useEffect(() => {
        const params = user ? { userId: String(user.id), id: String(id) } : { id: String(id) };

        api.get<GoResponse>(`/anime/`, { params: params }).then((res) => {
            if (res.data.code === 200) {
                const anime: Anime = res.data.data!;
                setAnime(anime);
                setHasCompleted(anime.reviews.some((r: Review) => r.userAnime.watchStatus === "completed"));
                const reviews = anime.reviews.filter((r: Review) => r.userAnime.watchStatus === "completed" && r.content !== "");

                setFixedReviews(reviews);

                if (user) {
                    const elem = user.reviews.find((r: Review) => r.userAnime.animeId === anime.id);
                    if (elem) {
                        setUserAnimeState(prevState => ({
                            ...prevState,
                            btnText: "Update",
                            isFavorite: elem.userAnime.isFav,
                            watchStatus: elem.userAnime.watchStatus,
                            score: elem.userAnime.score
                        }));
                    }
                }
            }
        }).catch((_: any) => {
        });

        api.get<GoResponse>("/categories/?category=watch_status&category=score").then((res) => {
            if (res.data.code === 200) {
                const data = res.data.data!;
                setWatchStatuses(data.watch_status);
                setScores(data.score);
            }
        }).catch((_: any) => { });
    }, [loading, id, user]);

    if (loading) {
        return <Spinner />;
    }

    return anime && (
        <div className="bg-base-200 rounded-2xl p-5">
            <div className="flex flex-col lg:flex-row lg:space-x-4">
                <div className="lg:w-1/3 w-full">
                    <OverviewForAnime anime={anime}>
                        <div className="flex items-center justify-center flex-row pt-3">
                            {user && (
                                <button type="submit" form="addToListForm" className="btn btn-info">
                                    {userAnimeState.btnText}
                                    <FontAwesomeIcon className={`mb-1`} icon={userAnimeState.btnText === "Add to list" ? faAdd : faArrowsRotate} width={15} />
                                </button>
                            )}
                            {user?.isMod && (
                                <>
                                    <div className="divider divider-horizontal">or</div>
                                    <Link href="/manage/add-anime" className="btn btn-outline">
                                        Edit <FontAwesomeIcon icon={faEdit} className="mb-1" width={15} />
                                    </Link>
                                </>
                            )}
                        </div>
                    </OverviewForAnime>
                </div>
                <div className="lg:w-2/3 w-full">
                    <Info anime={anime} scoreRange={scores.length}>
                        {user && (
                            <AddToList
                                userAnimeState={userAnimeState}
                                setUserAnimeState={setUserAnimeState}
                                scores={scores}
                                watchStatuses={watchStatuses} anime={anime} />
                        )}
                    </Info>
                    {user && (
                        <CreateReview hasCompleted={hasCompleted} refreshUser={refreshUser} user={user} anime={anime} />
                    )}
                    <DisplayReviews fixedReviews={fixedReviews} user={user} />
                </div>
            </div>
        </div>
    );
}


interface AddToListProps {
    userAnimeState: {
        btnText: string;
        isFavorite: boolean;
        watchStatus: string;
        score: string;
    };
    setUserAnimeState: Dispatch<SetStateAction<{
        btnText: string;
        isFavorite: boolean;
        watchStatus: string;
        score: string;
    }>>;
    scores: string[];
    watchStatuses: string[];
}
interface UserAnimeState {
    btnText: string;
    isFavorite: boolean;
    score: string;
    watchStatus: string;
}

interface AddToListProps {
    userAnimeState: UserAnimeState;
    setUserAnimeState: Dispatch<SetStateAction<UserAnimeState>>;
    anime: Anime;
    scores: string[];
    watchStatuses: string[];
}

const AddToList: FC<AddToListProps> = ({
    userAnimeState,
    setUserAnimeState,
    anime,
    scores,
    watchStatuses,
}) => {

    const { btnText, isFavorite, score, watchStatus } = userAnimeState;

    const handleAddToList = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        const formData = new FormData(e.currentTarget);
        formData.set("isFav", String(formData.get("isFav") === "on"));
        if (formData.get("watchStatus") === "plan-to-watch") {
            formData.set("score", "");
        }

        api.put<GoResponse>(`/auth/anime/${anime.id}/add-to-list/`, formData).then((res) => {
            if (res.data.code === 200) {
                setUserAnimeState((prevState) => ({
                    ...prevState,
                    btnText: "Update",
                }));
                toast.success(`${anime.title} has been updated!`);
            } else {
                toast.error("Could not update your anime list!");
            }
        }).catch(() => toast.error("You need to select values!"));
    };

    const handleStatusOrScore = (e: React.ChangeEvent<HTMLSelectElement>, isScore: boolean = false) => {
        e.preventDefault();
        const { value } = e.target;

        setUserAnimeState((prevState) => ({
            ...prevState,
            [isScore ? "score" : "watchStatus"]: value,
        }));
    };

    const toggleFavorite = () => {
        setUserAnimeState((prevState) => ({
            ...prevState,
            isFavorite: !prevState.isFavorite,
        }));
    };

    return (
        <form
            className="flex flex-col-reverse items-center justify-center md:flex-row lg:items-start lg:justify-between shadow p-3"
            encType="multipart/form-data"
            onSubmit={handleAddToList}
            id="addToListForm"
        >
            <div className="pt-3 flex justify-center">
                <label className="flex items-center justify-center cursor-pointer">
                    <input
                        type="checkbox"
                        name="isFav"
                        className="hidden"
                        checked={isFavorite}
                        onChange={toggleFavorite}
                    />
                    <span className="inline-flex items-center px-3 py-2 rounded">
                        <FontAwesomeIcon
                            icon={isFavorite ? faHeart : faHeartBroken}
                            className={`${isFavorite ? 'text-red-600' : 'text-gray-600'} text-4xl`}
                        />
                    </span>
                </label>
            </div>

            <div className="divider divider-horizontal"></div>
            <div>
                <Selector
                    lastElem={watchStatus}
                    text={"Current status"}
                    collection={watchStatuses}
                    name={"watchStatus"}
                    handler={(e) => handleStatusOrScore(e)}
                />
            </div>

            <div className="divider divider-horizontal"></div>
            <div>
                <Selector
                    lastElem={score}
                    text={"You liked it?"}
                    collection={scores}
                    name={"score"}
                    handler={(e) => handleStatusOrScore(e, true)}
                />
            </div>

            <div className="divider divider-horizontal"></div>
            <div className="flex flex-col items-center justify-center">
                <h1 className="md:text-xl text-2xl font-light">
                    Track your <span className="text-rose-400 font-extrabold animate-pulse">anime</span>!
                </h1>
                <button
                    form="addToListForm"
                    type="submit"
                    className="btn btn-success md:btn-sm max-sm:my-2 mt-2 rounded-2xl"
                >
                    <FontAwesomeIcon
                        icon={btnText === "Add to list" ? faAdd : faArrowsRotate}
                        className="mb-1"
                    />
                    {btnText}
                </button>
            </div>
        </form>
    );
};
