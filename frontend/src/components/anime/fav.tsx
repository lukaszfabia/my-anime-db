import { Review } from "@/types/models"
import { faHeart } from "@fortawesome/free-solid-svg-icons"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { FC } from "react"
import { AnimeShowcase } from "./animeOverview"

export const FavAnime: FC<{ userAnimes: Review[] }> = ({ userAnimes }) => {

    return (
        <div className="bg-base-300 p-5 text-center md:text-left rounded-xl">
            <h1 className="text-4xl pt-5 font-extrabold">
                Fav <span className="text-rose-600">anime <FontAwesomeIcon icon={faHeart} width={30} /></span>
            </h1>
            <p className="text-gray-500 text-sm pb-5">Here you have your favourite <u>anime</u>.</p>
            <div className="flex flex-row p-1 max-sm:justify-center max-sm:items-center">
                {
                    userAnimes && userAnimes
                        .filter((elem: Review) => elem.userAnime.isFav)
                        .map((elem: Review) =>
                            <AnimeShowcase a={elem.userAnime.anime} key={elem.animeId} />
                        )}
            </div>
        </div>
    )
}
