import { getImageUrl } from "@/lib/getImageUrl";
import { Anime, OtherTitles } from "@/types/models"
import { FC, ReactNode } from "react"
import Image from "next/image";
import { format } from "date-fns";
import Link from "next/link";

export const OverviewForAnime: FC<{ anime: Anime, children?: ReactNode }> = ({ anime, children }) => {
    return (
        <div className="md:p-5 flex flex-col max-sm:items-center max-sm:justify-center">
            <h1 className="text-5xl font-light dark:text-black text-white md:text-left text-center">
                {anime.title}
            </h1>
            <p className="text-gray-500 pb-4">{anime.alternativeTitles?.map((o: OtherTitles) => o.title).join(", ")}</p>
            <div className="text-neutral-content w-2/3">
                <div className="divider divider-vertical"></div>
                {anime.picUrl && (
                    <Image src={getImageUrl(anime.picUrl)} alt={anime.title} className="shadow-xl rounded-xl" width={500} height={500} />
                )}
                {/* place for buttons or sth  */}
                {children}

                <div className="w-full">
                    <div className="pt-5 flex flex-col dark:text-black text-white">
                        <h2 className="text-center text-2xl font-light">Informations</h2>
                        <div className="divider divider-vertical"></div>
                        {anime.studio && (<span><b className="text-gray-500">Studio: </b><Link href={`/studio/${anime.studio.id}`} className="text-blue-500 link link-hover">{anime.studio.name}</Link></span>)}
                        <span><b className="text-gray-500">Type: </b>{anime.animeType}</span>
                        <span><b className="text-gray-500">Pegi: </b>{anime.pegi}</span>
                        <span><b className="text-gray-500">Episodes: </b>{anime.episodes} / {anime.episodeLength} min.</span>
                        <div className="py-3"></div>
                        <h3 className="text-center text-lg font-light">Emitted time</h3>
                        <div className="divider divider-vertical"></div>
                        <span><b className="text-gray-500">Status: </b>{anime.status}</span>
                        {(anime.finishDate === anime.startDate && anime.startDate) ? (
                            <span><b className="text-gray-500">Aired: </b>{format(anime.startDate, "PPP")}</span>
                        ) : (
                            <>
                                {anime.startDate && (<span><b className="text-gray-500">Start: </b>{format(anime.startDate, "PPP")}</span>)}
                                {anime.finishDate && (<span><b className="text-gray-500">Finished: </b>{format(anime.finishDate, "PPP")}</span>)}
                            </>
                        )}
                    </div>
                </div>
            </div>
        </div>
    )
}