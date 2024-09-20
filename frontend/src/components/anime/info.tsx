import { CharacterCard } from "@/app/manage/add-anime/page";
import { Stat } from "@/types";
import { Anime, Role } from "@/types/models";
import { faFire, faPeopleGroup, faChartSimple, faMask } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { FC, ReactNode } from "react";
import { faStar } from "@fortawesome/free-regular-svg-icons";
import { AnimeShowcase } from "./animeOverview";

export const Info: FC<{ scoreRange: number, anime: Anime, children?: ReactNode }> = ({ scoreRange, anime, children }) => {
    const mostPopularGrade = anime.stats.mostPopularGrade !== "DEFAULT" ? anime.stats.mostPopularGrade : "-";

    const stats: Stat[] = [
        {
            title: "Popularity", value: anime.stats.popularity, icon: <FontAwesomeIcon icon={faFire} className="text-orange-600" />
        },
        { title: "Score", value: anime.stats.score, desc: "average score", icon: <FontAwesomeIcon icon={faStar} className="text-yellow-400" /> },
        { title: "Most popular grade", value: mostPopularGrade, desc: "most popular grade", icon: <FontAwesomeIcon icon={faPeopleGroup} /> },
    ]


    return (
        <div className="p-5 shadow">
            <div className="flex flex-col">
                <h1 className="text-4xl text-center md:text-left font-extrabold py-5">
                    Statistics <FontAwesomeIcon icon={faChartSimple} width={30} />
                </h1>
                <div className="flex flex-col items-center lg:flex-row lg:items-start lg:justify-between shadow p-3">
                    {stats.map((stat) => (
                        <div className="stat flex flex-col items-center text-center lg:text-left" key={stat.title}>
                            <div className="stat-title font-semibold">{stat.title}</div>
                            <div className="stat-value py-2 text-black dark:text-white max-md:text-3xl">
                                <span className="mr-2">
                                    {stat.title === "Popularity" && (<span className="text-gray-500 text-2xl mr-1">#</span>)}
                                    {stat.value}
                                    {stat.title === "Score" && (<span className="text-gray-500 text-sm">/ {scoreRange}</span>)}
                                </span>{stat.icon}
                            </div>
                            <div className="stat-desc text-wrap text-center">{stat.desc}</div>
                        </div>
                    ))}
                </div>

                <div>
                    {children}
                </div>

                <div className="text-center md:text-left">
                    <h1 className="text-4xl font-extrabold pt-3">Genres</h1>
                    {anime.genres.map((genre) => (<div key={genre.name} className="badge rounded-2xl badge-lg badge-neutral m-2 hover:badge-ghost hover:cursor-pointer">{genre.name}</div>))}
                </div>

                <div className="text-center md:text-left">
                    <h1 className="text-4xl font-extrabold pt-3">Synopsis</h1>
                    <p className="pt-1">{anime.description}</p>
                </div>

                <div className="divider divider-vertical"></div>

                <DisplayPrequelAndSequel anime={anime} />

                <div>
                    <h1 className="text-4xl font-extrabold pt-3">Characters <FontAwesomeIcon icon={faMask} /></h1>
                    <div className="py-5 grid grid-cols-1 md:grid-cols-2 gap-4">
                        {anime.roles.sort((a: Role, b: Role) => a.character.lastname.localeCompare(b.character.lastname)).map((role) => <CharacterCard key={role.character.name} role={role} />)}
                    </div>
                </div>
            </div>
        </div>
    )
}

const DisplayPrequelAndSequel: FC<{ anime: Anime }> = ({ anime }) => {
    return (
        <div className="w-full">
            <div className="flex justify-between px-8">
                {anime.prequel && <h1 className="text-4xl font-extrabold pt-3">Prequel</h1>}
                {anime.sequel && <h1 className="text-4xl font-extrabold pt-3">Sequel</h1>}
            </div>
            <div className="flex md:flex-row flex-col p-1 max-sm:justify-center max-sm:items-center py-4">
                {anime.prequel && <AnimeShowcase a={anime.prequel} />}
                {anime.sequel && <AnimeShowcase a={anime.sequel} />}
            </div>
        </div>
    )
}