"use client";

import { CustomInput } from "@/components/ui/forms/accountForm";
import { Selector } from "@/components/ui/selector";
import api from "@/lib/api";
import { convertTime } from "@/lib/computeTime";
import { ACCEPTED_IMAGE_TYPES } from "@/lib/config";
import { getImageUrl } from "@/lib/getImageUrl";
import { createName } from "@/lib/name";
import { Anime, Character, Genre, OtherTitles, Role, Studio, VoiceActor, } from "@/types/models";
import { GoResponse } from "@/types/responses";
import { faAdd, faArrowRight, faSave, faXmark } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Image from "next/image";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { FC, ReactNode, useEffect, useRef, useState } from "react";
import { toast } from "react-toastify";


export default function AddAnime() {
    const [alternativeTitles, setAlternativeTitles] = useState<string[]>([]);
    const altRef = useRef<HTMLInputElement | null>(null);

    // edit mode
    const [existiingAnimes, setExistingAnimes] = useState<Anime[]>([]);
    const [animeToUpdate, setAnimeToUpdate] = useState<Anime | null>(null);


    const [status, setStatus] = useState<string>("");
    const [animeType, setAnimeType] = useState<string>("");
    const [pegi, setPegi] = useState<string>("");

    const [animeTypes, setAnimeTypes] = useState<string[]>([]);
    const [pegis, setPegis] = useState<string[]>([]);
    const [statuses, setStatuses] = useState<string[]>([]);
    const [avaiableGenres, setAvaiableGenres] = useState<Genre[]>([]);
    const [genres, setGenres] = useState<Genre[]>([]);
    const [studios, setStudios] = useState<Studio[]>([]);
    const [animes, setAnimes] = useState<Anime[]>([]);

    const [sequel, setSequel] = useState<Anime | null>(null);
    const [prequel, setPrequel] = useState<Anime | null>(null);

    const [studio, setStudio] = useState<Studio | null>(null);
    const [allCastRoles, setAllCastRoles] = useState<string[]>([]);

    const [allCharacters, setAllCharacters] = useState<Character[]>([]);
    const [allVoiceActors, setAllVoiceActors] = useState<VoiceActor[]>([]);

    const [character, setCharacter] = useState<Character | null>(null);
    const [voiceActor, setVoiceActor] = useState<VoiceActor | null>(null);
    const [castRole, setCastRole] = useState<string>("");

    const [roles, setRoles] = useState<Role[]>([]);

    useEffect(() => {
        api.get<GoResponse>("/all-anime/")
            .then((res) => {
                if (res.status === 200) {
                    const existingAnimes: Anime[] = res.data.data || [];
                    console.log(existingAnimes)
                    setExistingAnimes(existingAnimes);
                } else {
                    toast.warning("there is no anime to update!");
                }
            }).catch((_: any) => toast.error("failed to fetch anime!"));


        api
            .get<GoResponse>("/categories/?category=cast_role&category=voice_actor&category=character&category=anime_type&category=pegi&category=anime_status&category=genre&category=studio&category=anime")
            .then((res) => {
                if (res.status === 200) {
                    setAnimeTypes(res.data.data.anime_type);
                    setPegis(res.data.data.pegi);
                    setStatuses(res.data.data.anime_status);
                    setAvaiableGenres(res.data.data.genre);
                    setStudios(res.data.data.studio);
                    setAnimes(res.data.data.anime);
                    setAllCharacters(res.data.data.character);
                    setAllVoiceActors(res.data.data.voice_actor);
                    setAllCastRoles(res.data.data.cast_role);
                }
            })
            .catch((_: any) => toast.error("failed to fetch anime types, pegi and status!"));
    }, []);

    const AlternativeTitles: FC = () => {
        const addAltTitle = () => {
            if (altRef.current?.value) {
                const v: string = (altRef.current.value).trim();
                setAlternativeTitles(alternativeTitles.find((title: string) => title.toLowerCase() === v.toLowerCase()) ? alternativeTitles : [...alternativeTitles, v]);
                altRef.current.value = "";
            }
        };

        const removeAltTitle = (title: string) => {
            setAlternativeTitles(alternativeTitles.filter((t: string) => t !== title));
        };
        return (
            <div>
                <p className="text-gray-500">
                    Here you can add other titles to help in finding anime!
                </p>
                <div className="flex flex-row justify-center md:justify-start">
                    <CustomInput placeholder="Other titles..." type="text" name="altTitles" inputRef={altRef} required={false} />
                    <button type="button" className="btn btn-primary btn-circle ml-3 mt-4" onClick={addAltTitle}><FontAwesomeIcon icon={faAdd} width={15} /></button>
                </div>
                <div className="flex flex-row flex-wrap justify-center md:justify-start">
                    {alternativeTitles.map((title: string) => (
                        <div key={title} className="badge badge-lg badge-secondary transition-all ease-in-out duration-200 hover:badge-error mr-2 mt-2 gap-1 rounded-lg">
                            <button type="button" onClick={() => removeAltTitle(title)}>
                                <FontAwesomeIcon icon={faXmark} width={15} />
                            </button>{title}
                        </div>
                    ))}
                </div>
            </div>
        );
    };

    const router = useRouter();

    const addAnime = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const form = new FormData(e.currentTarget);
        form.delete("altTitles");
        form.delete("genres");
        form.delete("roles");
        form.delete("voice-actor");
        form.delete("character");
        form.delete("cast-role");
        form.set("prequel", prequel?.id.toString() || "");
        form.set("sequel", sequel?.id.toString() || "");

        alternativeTitles.forEach((title: string) => form.append("altTitles", title));
        genres.forEach((genre: Genre) => form.append("genres", genre.id.toString()));
        form.append("roles", JSON.stringify(roles));
        form.set("studio", studio?.id.toString() || "");


        const method = animeToUpdate ?
            api.put<GoResponse>(`/auth/manage/anime/${animeToUpdate.id}`, form) :
            api.post<GoResponse>("/auth/manage/anime/", form);

        method.then((res) => {
            if (res.status === 201 || res.status === 200) {
                toast.success(`${form.get("title")} has been ${animeToUpdate ? "updated" : "added"}!`);
                const newAnime: Anime = res.data.data;
                router.push(`/anime/${newAnime.id}`);
            }
        }).catch((err: any) => {
            console.log(err);
            toast.error("something went wrong!");
        });
    };

    const addGenre = (e: React.ChangeEvent<HTMLSelectElement>) => {
        const v = e.currentTarget.value;

        const genre = avaiableGenres.find((genre: Genre) => genre.name === v);

        if (genre && !genres.find((g: Genre) => g.name === genre.name)) {
            setGenres((prevGenres) => [...prevGenres, genre]);
        }
    };


    const handleSeqPreq = (e: React.ChangeEvent<HTMLSelectElement>, type: "sequel" | "prequel") => {
        e.preventDefault();
        const v = e.currentTarget.value;
        const anime = animes.find((a: Anime) => a.title === v)!;
        type === "sequel" ? setSequel(anime) : setPrequel(anime);
    };

    const handleStudio = (e: React.ChangeEvent<HTMLSelectElement>) => {
        e.preventDefault();
        const v = e.currentTarget.value;
        const studio = studios.find((s: Studio) => s.name === v);
        if (studio) {
            setStudio(studio);
        }
    };

    const handleCharOrActor = (e: React.ChangeEvent<HTMLSelectElement>, type: "actor" | "character") => {
        e.preventDefault();
        const v = e.currentTarget.value;
        if (type === "character") {
            const character = allCharacters.find((c: Character) => (c.lastname + " " + c.name) === v)!;
            setCharacter(character);
        } else {
            const voiceActor = allVoiceActors.find((va: VoiceActor) => (va.lastname + " " + va.name) === v)!;
            setVoiceActor(voiceActor);
        }
    };

    const setRole = () => {
        if (character && voiceActor && castRole) {
            const role: Role = {
                actorId: voiceActor.id,
                characterId: character.id,
                role: castRole,
                voiceActor: voiceActor,
                character: character,
                animeId: 0,
                anime: {} as Anime,
            };

            // adding only unique roles
            if (!roles.find((r: Role) => r.characterId === role.characterId || r.actorId === role.actorId)) {
                setRoles((prevRoles) => [...prevRoles, role]);
            }
        };
    }

    const removeRole = (role: Role) => {
        setRoles(roles.filter((r: Role) => r.characterId !== role.characterId && r.actorId !== role.actorId));
    }

    const handleAnimeUpdate = (e: React.ChangeEvent<HTMLSelectElement>) => {
        const v = e.currentTarget.value;

        const anime = existiingAnimes.find((a: Anime) => a.title === v);
        if (anime) {
            setAnimeToUpdate(anime);

            api.get<GoResponse>(`/anime/?animeId=${anime.id}`).then((res) => {
                if (res.status === 200) {
                    const roles: Role[] | null = res.data.data.roles;
                    if (!roles) {
                        return;
                    }
                    setRoles(roles);
                }
            }).catch((_: any) => toast.error("something went wrong!"));

            // fill form with existing anime data
            setAlternativeTitles(anime.alternativeTitles ? anime.alternativeTitles?.map((t: OtherTitles) => t.title) : []);
            setStatus(anime.status);
            setAnimeType(anime.animeType);
            setPegi(anime.pegi);
            setGenres(anime.genres);
            setStudio(anime.studio ? anime.studio : null);
            setPrequel(anime.prequel ? anime.prequel : null);
            setSequel(anime.sequel ? anime.sequel : null);
        } else {
            clearFields();
        }
    }

    const clearFields = () => {
        setAnimeToUpdate(null);
        setAlternativeTitles([]);
        setStatus("");
        setAnimeType("");
        setPegi("");
        setGenres([]);
        setStudio(null);
        setRoles([]);
        setVoiceActor(null);
        setCharacter(null);
        setCastRole("");
    }


    return (
        <form encType="multipart/form-data" onSubmit={(e) => addAnime(e)} className="p-1 max-w-full overflow-x-hidden text-center md:text-left">
            <h1 className="font-extrabold text-4xl">Add <span className="text-blue-400">Anime</span> or update existing</h1>
            <p className="text-gray-500 pt-2 pb-3">Down below you can choose anime to update or if you want to create new just fill required fields.</p>
            <div className="flex flex-row">
                <Selector collection={existiingAnimes.map((a: Anime) => a.title)} text="Choose anime to update" name="animeToUpdate" lastElem={animeToUpdate?.title} handler={handleAnimeUpdate} />
                <button type="button" className="btn btn-error ml-3 mt-3 btn-circle" onClick={clearFields}><FontAwesomeIcon icon={faXmark} width={15} /></button>
            </div>
            <div className="flex flex-col md:flex-row pt-10 justify-center md:justify-between">
                <div className="w-full md:w-1/3">
                    <p className="text-warning text-left">Attention: Title must be <b>unique</b>.</p>
                    <CustomInput placeholder="" defaultValue={animeToUpdate ? animeToUpdate.title : ""} type="text" name="title" required>
                        <b>Title</b>
                    </CustomInput>
                </div>
                <div className="w-full md:w-1/2">
                    <AlternativeTitles />
                </div>
            </div>

            <div className="flex flex-col md:flex-row pt-10 justify-center items-center md:items-start md:justify-between">
                <div className="md:w-1/2 w-full">
                    <h1 className="text-2xl font-semibold">Select parameters</h1>
                    <p className="text-gray-500">Contains important technical information.</p>
                    <div className="my-5"></div>
                    <div className="max-w-md">
                        <div className="text-gray">
                            <h1 className="text-lg font-semibold">Difference between ONA & OVA.</h1>
                            <ul className="text-gray-500 space-y-2 md:list-disc list-inside">
                                <li><b>ONA</b> (<i>Original Net Animation</i>) - relased for Internet community.</li>
                                <li><b>OVA</b> (<i>Original Video Animation</i>) - kind of specials or alternative story.</li>
                            </ul>
                        </div>
                    </div>
                    <div className="flex flex-col justify-center items-center md:items-start md:justify-start">
                        <Selector collection={animeTypes} text="Choose anime type" name="animeType" lastElem={animeType} handler={(e) => {
                            e.preventDefault();
                            setAnimeType(e.currentTarget.value);
                        }} />
                        <Selector collection={pegis} text="What is target group?" name="pegi" lastElem={pegi} handler={(e) => {
                            e.preventDefault();
                            setPegi(e.currentTarget.value);
                        }} />
                        <Selector collection={statuses} text="Set anime status" name="status" lastElem={status} handler={(e) => {
                            e.preventDefault();
                            setStatus(e.currentTarget.value);
                        }} />
                    </div>
                </div>
                <div className="md:w-1/2 w-full">
                    <h1 className="text-2xl font-semibold">Some numbers</h1>
                    <p className="text-gray-500">
                        Set number of episodes and duration of one episode.
                    </p>
                    <div className="w-72 md:w-fit mx-auto md:mx-0">
                        <CustomInput defaultValue={animeToUpdate ? animeToUpdate.episodes : animeType === "movie" ? 1 : 12} placeholder="e.g 12" type="number" name="episodes" required={status === "currently-airing" || status === "finished"}>
                            <b className="text-sm md:text-base">Episodes</b>
                        </CustomInput>
                        <CustomInput defaultValue={animeToUpdate ? animeToUpdate.episodeLength : 24} placeholder="e.g 24" type="number" name="episodeLength" required={status === "currently-airing" || status === "finished"}>
                            <b className="text-sm md:text-base">Duration</b>
                        </CustomInput>
                        <div className="divider divider-vertical"></div>
                        <h1 className="text-2xl font-semibold">Dates</h1>
                        <p className="text-gray-500">You can set <b>optionally</b> start and finish date.</p>
                        <CustomInput defaultValue={animeToUpdate ? convertTime(animeToUpdate.startDate) : ""} placeholder="" type="date" name="startDate" disabled={status === "unknown"}>
                            <b>Start</b>
                        </CustomInput>

                        <CustomInput placeholder="" defaultValue={animeToUpdate ? convertTime(animeToUpdate.finishDate) : ""} type="date" name="finishDate" disabled={status === "unknown" || status === "planned"}>
                            <b>Finish</b>
                        </CustomInput>
                    </div>
                </div>
            </div>

            <div className="flex flex-col md:flex-row pt-10 justify-center md:justify-between">
                <div className="md:w-2/5 w-full max-sm:py-5">
                    <h1 className="text-2xl font-semibold">Description</h1>
                    <p className="text-gray-500">Write something about anime.</p>
                    <textarea defaultValue={animeToUpdate ? animeToUpdate.description : ""} className="textarea textarea-bordered w-full mt-2" placeholder="Description..." required name="description"></textarea>
                </div>

                <div className="md:w-1/2 w-full">
                    <h1 className="text-2xl font-semibold">Upload cover</h1>
                    <p className="text-gray-500">You can <span>{animeToUpdate?.picUrl ? "update" : "upload new"}</span> image.</p>
                    <input type="file" name="pic" id="pic" className="file file-input mt-5 w-full" accept={ACCEPTED_IMAGE_TYPES} />
                </div>
            </div>
            <div className="flex flex-col md:flex-row pt-10 justify-center items-center md:justify-between">
                <div className="py-5 md:w-1/2 w-full">
                    <h1 className="text-2xl font-semibold">Assign some <span className="text-green-600">Genres</span></h1>
                    <p className="text-gray-500">Choose <b>genres</b> which describes anime.</p>
                    <div className="flex items-center justify-center md:justify-start">
                        <Selector collection={avaiableGenres.map((e: Genre) => e.name)} lastElem={genres.length > 0 ? genres[genres.length - 1].name : ""} text="Select genre" name="genres" handler={addGenre} />
                    </div>

                    {genres.length > 0 && genres.map((genre: Genre) => (
                        <div key={genre.name} className="badge badge-lg badge-primary transition-all ease-in-out duration-200 hover:badge-error mr-2 mt-2 gap-1 rounded-lg">
                            <button type="button" onClick={() => setGenres(genres.filter((g: Genre) => g.name !== genre.name))}>
                                <FontAwesomeIcon icon={faXmark} width={15} />
                            </button>{genre.name}
                        </div>
                    ))}
                </div>

                <div className="py-5 md:w-1/2 w-full">
                    <h1 className="text-2xl font-semibold">What <span className="text-red-400">Studio</span> has been created this Anime?</h1>
                    <p className="text-gray-500">Just take a look what we have got and select!</p>
                    <div className="flex justify-center items-center md:justify-start">
                        <Selector collection={studios.map((e: Studio) => e.name)} text="Select studio" name="studio" handler={handleStudio}
                            lastElem={studio ? studio.name : ""} />
                    </div>
                </div>
            </div>

            {animes &&
                <div className="text-center md:text-left">
                    <h1 className="text-2xl font-semibold">
                        You can set <span className="text-teal-700">Prequel</span> and <span className="text-violet-800">Sequel</span>
                    </h1>
                    <p className="text-gray-500 md:w-3/4">
                        While these fields are <b>optional</b>, adding them can significantly enhance navigation and context for your anime. Linking related titles makes it easier for viewers to follow the storyline across different seasons or related series.
                    </p>
                    <div className="flex flex-col md:flex-row justify-center items-center py-3">
                        <Selector disablePrompt collection={animes.map((e: Anime) => e.title)} disableValue={sequel?.title} text="Select prequel" name="prequel" lastElem={prequel ? prequel.title : ""} handler={(e) => handleSeqPreq(e, "prequel")} />
                        <div className="divider divider-horizontal"></div>
                        <Selector disablePrompt collection={animes.map((e: Anime) => e.title)} disableValue={prequel?.title} text="Select sequel" name="sequel" lastElem={sequel ? sequel.title : ""} handler={(e) => handleSeqPreq(e, "sequel")} />
                    </div>
                </div>
            }

            <div className="flex flex-col md:flex-row justify-center md:justify-between" id="roles">
                <div className="py-5 md:w-1/3">
                    <h1 className="text-2xl font-semibold">Create <span className="text-orange-400">Roles</span></h1>
                    <p className="text-gray-500">Add characters with their voice actors! Maybe you are missing an <Link href="/manage/add-others/voice_actor" className="link link-info">actor</Link> or <Link href="/manage/add-others/character" className="link link-info">character</Link> - you can add them in any time!</p>
                    <div className="flex flex-col md:justify-start justify-center items-center">
                        <Selector handler={(e) => handleCharOrActor(e, "character")} lastElem={createName<Character>(character)} collection={allCharacters.map((e: Character) => e.lastname + " " + e.name)} text="Select character" name="character" />
                        <Selector handler={(e) => handleCharOrActor(e, "actor")} lastElem={createName<VoiceActor>(voiceActor)} collection={allVoiceActors.map((e: VoiceActor) => e.lastname + " " + e.name)} text="Select voice actor" name="voice-actor" />
                        <Selector handler={(e) => {
                            e.preventDefault();
                            setCastRole(e.currentTarget.value);
                        }} lastElem={castRole} collection={allCastRoles} text="Select role" name="cast-role" />
                        <div className="flex items-center justify-center w-full max-w-xs">
                            <button type="button" className="btn btn-ghost btn-outline mt-3" onClick={setRole}><FontAwesomeIcon icon={faArrowRight} width={15} /><span>Append</span></button>
                        </div>
                    </div>
                </div>
                {roles && roles.length > 0 && (
                    <>
                        <div className="divider divider-horizontal"></div>
                        <div className="py-5 md:w-2/3 grid grid-cols-1 md:grid-cols-2 gap-4">
                            {roles.map((role: Role) => (
                                <CharacterCard key={role.character.name} role={role} >
                                    <button type="button" className="ml-auto btn btn-sm btn-error" onClick={() => removeRole(role)}>
                                        <FontAwesomeIcon icon={faXmark} />
                                    </button>
                                </CharacterCard>
                            ))}
                        </div>
                    </>
                )}
            </div>
            <div className="flex justify-center md:justify-end w-full md:w-5/6 mt-5">
                <button type="submit" className="btn btn-secondary">
                    <FontAwesomeIcon icon={animeToUpdate ? faSave : faAdd} width={15} />
                    <span>{animeToUpdate ? "Save anime" : "Add anime"}</span>
                </button>
            </div>
        </form>
    );
}

export const CharacterCard: FC<{ role: Role, children?: ReactNode }> = ({ role, children }) => {
    const [isHovered, setIsHovered] = useState<boolean>(false);
    //TODO check it 
    return (
        <div className="bg-base-200 p-5 rounded-xl shadow-xl w-full h-min">
            <div className="flex items-center">
                <div>
                    <div
                        className="placeholder relative w-24 h-24 rounded-full overflow-hidden ring-offset-base-100 ring-2 ring-offset-2"
                        onMouseEnter={() => setIsHovered(true)}
                        onMouseLeave={() => setIsHovered(false)}
                    >
                        {role.character.picUrl ? (
                            <Image
                                src={getImageUrl(role.character.picUrl)}
                                alt={role.character.name}
                                layout="responsive"
                                width={100}
                                height={100}
                                objectFit="cover"
                                className={`transition-opacity duration-300 ${isHovered ? 'opacity-0' : 'opacity-100'}`}
                            />
                        ) : <span className="text-3xl">{role.character.name.charAt(0)}</span>}
                        {role.voiceActor.picUrl ? (
                            <Image
                                src={getImageUrl(role.voiceActor.picUrl)}
                                alt={role.voiceActor.name}
                                layout="responsive"
                                width={100}
                                height={100}
                                objectFit="cover"
                                className={`transition-opacity duration-300 absolute top-0 left-0 ${isHovered ? 'opacity-100' : 'opacity-0'}`}
                            />
                        ) : <span className="text-3xl">{role.voiceActor.name.charAt(0)}</span>}
                    </div>
                </div>
                <div className="ml-4 flex flex-col max-w-xs">
                    <p className="text-md font-bold hover:underline"><Link href={`/character/${role.characterId}`}>{createName<Character>(role.character)}</Link></p>
                    <p className="text-sm font-semibold text-blue-500 hover:underline" ><Link href={`/voice-actors/${role.actorId}`}>{createName<VoiceActor>(role.voiceActor)}</Link></p>
                    <p className="text-sm">{role.role}</p>
                </div>
            </div>
            <div className="flex items-center justify-center">
                {children}
            </div>
        </div>
    )
}