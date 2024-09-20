import { getImageUrl } from "@/lib/getImageUrl";
import { Review } from "@/types/models";
import { faXmark } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Link from "next/link";
import { Dispatch, FC, SetStateAction, useState } from "react";
import Image from "next/image";

type SortKey = "title" | "status" | "score" | "nr";

interface SortButtonProps {
    sort: (key: SortKey) => void;
    text: string;
    currKey: SortKey;
}

const SortButton: FC<SortButtonProps> = ({ sort, text, currKey }) => {
    return (
        <button type="button" onClick={() => sort("title")} className={`btn btn-sm btn-ghost ${currKey === text.toLowerCase().replace(".", "") && "btn-active"}`}>{text}</button>
    )
}

export const AnimeRecord: FC<{ rs: Review[], setAnimes: Dispatch<SetStateAction<Review[]>> }> = ({ rs, setAnimes }) => {
    const [_, setOnSort] = useState<number>(1);
    const [currKey, setCurrKey] = useState<SortKey>("nr");

    const sort = (key: SortKey) => {
        setOnSort((prevOnSort) => {
            const newOnSort = prevOnSort * -1;
            let sorted = [...rs];

            switch (key) {
                case "title":
                    sorted.sort((a, b) => newOnSort * a.userAnime.anime.title.localeCompare(b.userAnime.anime.title));
                    break;
                case "nr":
                    sorted.sort((a, b) => newOnSort * (a.userAnime.createdAt.getTime > b.userAnime.createdAt.getTime ? 1 : -1));
                    break;
                case "status":
                    sorted.sort((a, b) => a.userAnime.watchStatus.localeCompare(b.userAnime.watchStatus) * newOnSort);
                    break;
                case "score":
                    sorted.sort((a, b) => a.userAnime.score.localeCompare(b.userAnime.score) * newOnSort);
                    break;
                default:
                    break;
            }

            setAnimes(sorted);
            setCurrKey(key);
            return newOnSort;
        });
    };


    return (
        <div className="grid grid-cols-1 border border-gray-800 dark:border-gray-100 divide-y divide-gray-700 dark:divide-gray-200 bg-base-200 rounded-xl shadow-xl">
            <div className="grid grid-cols-6 px-5 py-2 font-bold text-center">
                <SortButton sort={() => sort("nr")} text="Nr." currKey={currKey} />
                <h1 className="btn btn-sm">Image</h1>
                <SortButton sort={() => sort("title")} text="Title" currKey={currKey} />
                <SortButton sort={() => sort("status")} text="Status" currKey={currKey} />
                <SortButton sort={() => sort("score")} text="Score" currKey={currKey} />
                <h1 className="btn btn-sm">Remove</h1>
            </div>

            {rs.length > 0 ? rs.map((r: Review, i: number) => (
                <div key={r.userAnime.anime.title} className={`grid grid-cols-6 items-center px-5 py-2 font-light text-center ${i % 2 == 1 && "bg-base-300"} ${i === rs.length - 1 && "rounded-b-xl"}`}>
                    <h1 className="text-indigo-500 font-extrabold text-3xl">#{i + 1}</h1>

                    <div className="flex items-center justify-center">
                        <Image
                            src={getImageUrl(r.userAnime.anime.picUrl)}
                            alt={r.userAnime.anime.title}
                            objectFit="cover"
                            width={150}
                            height={150}
                            className="rounded-lg shadow-xl w-20"
                        />
                    </div>

                    <h1 className="font-bold text-blue-500 link link-hover text-wrap"><Link href={`/anime/${r.userAnime.animeId}`}>{r.userAnime.anime.title}</Link>
                        {r.userAnime.anime.alternativeTitles && (<p className="text-sm text-gray-500">{r.userAnime.anime.alternativeTitles.join(", ")}</p>)}</h1>
                    <h1>{r.userAnime.watchStatus}</h1>
                    <h1 className="font-semibold text-black dark:text-white">{r.userAnime.score === "" ? "-" : r.userAnime.score}</h1>

                    <div className="flex items-center justify-center">
                        <button type="button" className="btn btn-sm btn-error btn-circle">
                            <FontAwesomeIcon icon={faXmark} width={10} />
                        </button>
                    </div>
                </div>

            )) : (
                <h1 className="text-center font-extrabold text-5xl py-5">
                    No <span className="text-blue-500">animes</span> found.
                </h1>
            )}
        </div>
    );
};