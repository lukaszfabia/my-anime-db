"use client";

import { AnimeRecord } from "@/components/anime/record";
import { useAuth } from "@/components/providers/auth"
import { Spinner } from "@/components/ui/spinner"
import api from "@/lib/api";
import { Genre, Review, Studio } from "@/types/models"
import { GoResponse } from "@/types/responses";
import { faSearch } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { Dispatch, FC, forwardRef, ReactNode, SetStateAction, useEffect, useRef, useState } from "react"
import { toast } from "react-toastify";


interface SearchMenuProps {
    types: string[];
    scores: string[];
    statuses: string[];
    genres: Genre[];
    studios: Studio[];
}


const Chip = forwardRef<HTMLInputElement, { text: string }>(({ text }, ref) => {
    const [checked, setChecked] = useState(false);
    const handleChange = () => {
        setChecked(!checked);
    };

    return (
        <div className="m-2">
            <input
                type="checkbox"
                ref={ref}
                id={text.toLowerCase()}
                name={text.toLowerCase()}
                checked={checked}
                value={text.toLowerCase()}
                onChange={handleChange}
                className="peer hidden"
            />
            <label
                htmlFor={text.toLowerCase()}
                className={`flex items-center justify-between w-full p-2 transition-all ease-in-out duration-200 text-gray-500 border-1 rounded-3xl cursor-pointer hover:bg-base-300 border border-base-100 hover:border hover:border-gray-500 ${checked ? 'bg-indigo-600 text-white border-indigo-700 hover:bg-indigo-700 hover:text-gray-300' : 'hover:bg-base-300'
                    }`}
            >
                <div className="block">
                    <div className="text-center text-base">{text}</div>
                </div>
            </label>
        </div>
    );
});


const ChipWrapper: FC<{ text: string, odd: boolean, children?: ReactNode }> = ({ text, odd, children }) => {
    return (
        <div className="collapse collapse-plus bg-base-200">
            <input type="radio" name="accordion" defaultChecked />
            <h1 className={`collapse-title text-xl font-medium dark:text-white text-black ${odd && "md:text-right"}`}>{text}</h1>
            <div className="collapse-content font-light flex flex-wrap">
                {children}
            </div>
        </div>
    )
}

const SearchMenu: FC<{ params: SearchMenuProps, animes: Review[], setAnimes: Dispatch<SetStateAction<Review[]>> }> = ({ params, animes, setAnimes }) => {
    const searchRef = useRef<HTMLInputElement>(null);

    const genresRefs = useRef<{ [key: string]: HTMLInputElement | null }>({});
    const studiosRefs = useRef<{ [key: string]: HTMLInputElement | null }>({});
    const typesRefs = useRef<{ [key: string]: HTMLInputElement | null }>({});
    const scoresRefs = useRef<{ [key: string]: HTMLInputElement | null }>({});
    const statusesRefs = useRef<{ [key: string]: HTMLInputElement | null }>({});

    const filter = () => {
        let filteredAnimes = [...animes];

        const search = searchRef.current?.value.toLowerCase();
        if (search) {
            filteredAnimes = filteredAnimes.filter((r: Review) =>
                r.userAnime.anime.title.toLowerCase().includes(search) ||
                r.userAnime.anime.alternativeTitles?.join(" ").toLowerCase().includes(search));
        }

        const selectedGenres = Object.values(genresRefs.current).filter(ref => ref?.checked).map(ref => ref?.value);
        if (selectedGenres.length > 0) {
            filteredAnimes = filteredAnimes.filter((r: Review) =>
                r.userAnime.anime.genres.some((g: Genre) => selectedGenres.includes(g.name.toLowerCase())));
        }

        const selectedStudios = Object.values(studiosRefs.current).filter(ref => ref?.checked).map(ref => ref?.value);
        if (selectedStudios.length > 0) {
            filteredAnimes = filteredAnimes.filter((r: Review) =>
                selectedStudios.includes(r.userAnime.anime.studio?.name.toLowerCase() || ""));
        }

        const selectedTypes = Object.values(typesRefs.current).filter(ref => ref?.checked).map(ref => ref?.value);
        if (selectedTypes.length > 0) {
            filteredAnimes = filteredAnimes.filter((r: Review) =>
                selectedTypes.includes(r.userAnime.anime.animeType.toLowerCase()));
        }

        const selectedStatuses = Object.values(statusesRefs.current).filter(ref => ref?.checked).map(ref => ref?.value);
        if (selectedStatuses.length > 0) {
            filteredAnimes = filteredAnimes.filter((r: Review) =>
                selectedStatuses.includes(r.userAnime.watchStatus.toLowerCase()));
        }

        const selectedScores = Object.values(scoresRefs.current).filter(ref => ref?.checked).map(ref => ref?.value);
        if (selectedScores.length > 0) {
            filteredAnimes = filteredAnimes.filter((r: Review) =>
                selectedScores.includes(r.userAnime.score.toString()));
        }

        setAnimes(filteredAnimes);
    };

    return (
        <div className="bg-base-200 px-5 py-4 rounded-xl shadow-xl border border-gray-800 dark:border-gray-100" onChange={filter}>
            <h1 className="text-xl font-bold pt-5">
                <FontAwesomeIcon icon={faSearch} /> Filter options
            </h1>
            <p className="text-sm text-gray-500 pt-2">
                You can also use <b>sort</b> methods above <b>anime</b> record by clicking on the <u>label</u>.
            </p>
            <div className="flex items-center justify-center py-4">
                <input placeholder="Search..." type="text" className="input w-full" ref={searchRef} />
            </div>

            <div className="text-xl font-bold md:text-left text-center space-y-6">
                <ChipWrapper text="Types" odd={false}>
                    {params.types.map((type) => (
                        <Chip key={type} text={type} ref={(ref) => { if (ref) (typesRefs.current[type] = ref) }} />
                    ))}
                </ChipWrapper>

                <ChipWrapper text="Scores" odd={true}>
                    {params.scores.map((score) => (
                        <Chip key={score} text={score} ref={(ref) => { if (ref) (scoresRefs.current[score] = ref) }} />
                    ))}
                </ChipWrapper>

                <ChipWrapper text="Statuses" odd={false}>
                    {params.statuses.map((status: string) => (
                        <Chip key={status} text={status} ref={(ref) => { if (ref) (statusesRefs.current[status] = ref) }} />
                    ))}
                </ChipWrapper>

                <ChipWrapper text="Genres" odd={true}>
                    {params.genres.map((genre: Genre) => (
                        <Chip key={genre.name} text={genre.name} ref={(ref) => { if (ref) (genresRefs.current[genre.name] = ref) }} />))}
                </ChipWrapper>

                <ChipWrapper text="Studios" odd={false}>
                    {params.studios.map((studio: Studio) => (
                        <Chip key={studio.name} text={studio.name} ref={(ref) => { if (ref) (studiosRefs.current[studio.name] = ref) }} />
                    ))}
                </ChipWrapper>
            </div>
        </div>
    )
}


export default function Collection() {
    const { user, loading } = useAuth();
    const [animes, setAnimes] = useState<Review[]>([]);
    const [filteredAnimes, setFilteredAnimes] = useState<Review[]>([]);
    const [params, setParams] = useState<SearchMenuProps | null>(null);

    useEffect(() => {
        api.get<GoResponse>("/categories/")
            .then(res => {
                const updatedParams = {
                    types: res.data.data.anime_type || [],
                    scores: res.data.data.score || [],
                    statuses: res.data.data.anime_status || [],
                    genres: res.data.data.genre || [],
                    studios: res.data.data.studio || [],
                };
                setParams(updatedParams);
                if (user?.reviews) {
                    setAnimes(user.reviews)
                    setFilteredAnimes(user.reviews)
                }
            })
            .catch(() => toast.error("Failed to fetch categories"));
    }, [user, loading]);

    if (loading || !params) {
        return <Spinner />;
    }

    if (!user) {
        return null;
    }

    return (
        <div className="flex flex-col-reverse lg:flex-row">
            <div className="lg:w-3/4">
                <AnimeRecord rs={filteredAnimes} setAnimes={setFilteredAnimes} />
            </div>
            <div className="divider divider-horizontal"></div>
            <div className="lg:w-1/4">
                <SearchMenu params={params} animes={animes} setAnimes={setFilteredAnimes} />
            </div>
        </div>
    );
}
