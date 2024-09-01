"use client";

import { CustomInput } from "@/components/ui/forms/accountForm";
import api from "@/lib/api";
import { loadImage } from "@/lib/loadImage";
import { Character, Genre, Studio, VoiceActor } from "@/types/models";
import { GoResponse } from "@/types/responses";
import { faPlus, faSave, faXmark } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { AxiosResponse } from "axios";
import Image from "next/image";
import Link from "next/link";
import { usePathname } from "next/navigation";
import React from "react";
import { FC, useEffect, useState } from "react";
import { toast } from "react-toastify";
import { createObj, deleteObj } from "./manage";
import { ACCEPTED_IMAGE_TYPES } from "@/lib/config";

const availableTypes: CreatorProps[] = [
    { entity: "voice-actor" },
    { entity: "character" },
    { entity: "studio" },
    { entity: "genre" },
];

export default function OtherCreator() {
    const slug = usePathname();
    const slugArray = slug.split("/");
    const addedEntity = slugArray[slugArray.length - 1];
    const mappedType: CreatorProps | undefined = availableTypes.find((e: CreatorProps) => e.entity === addedEntity);

    return <>{mappedType ? <Creator entity={mappedType.entity} /> : null}</>;
}

export interface CreatorProps {
    entity: "voice-actor" | "character" | "studio" | "genre";
}

const Creator: FC<CreatorProps> = ({ entity }) => {
    switch (entity) {
        case "voice-actor":
            return <VoiceActorCreator />;
        case "character":
            return <CharacterCreator />;
        case "studio":
            return <StudioCreator />;
        case "genre":
            return <GenreCreator />;
        default:
            return null;
    }
};

const CharacterCreator: FC = () => {
    const [characters, setCharacters] = useState<Character[]>([]);
    useEffect(() => {
        api.get<GoResponse>("/categories/?category=character")
            .then((res) => {
                setCharacters(res.data.data.character);
            })
            .catch((_: any) => toast.error("failed to fetch characters"));
    }, []);

    const createCharacter = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        createObj(e, { entity: "character" }).then((res) => {
            if (res && res.data) {
                const newCharacter: Character = res.data
                setCharacters((prev) => [...prev, newCharacter]);
            }
        });
    }

    const deleteCharacter = (id: number) => {
        deleteObj(id, { entity: "character" }, () => {
            setCharacters((prev) => prev.filter((e) => e.id !== id));
        })
    };

    return <CreatorForm
        createObj={createCharacter}
        deleteObj={deleteCharacter}
        objectName="Character"
        objects={characters} />;
}

const StudioCreator: FC = () => {
    const [studios, setStudios] = useState<Studio[]>([]);
    useEffect(() => {
        api.get<GoResponse>("/categories/?category=studio")
            .then((res) => {
                setStudios(res.data.data.studio);
            })
            .catch((_: any) => toast.error("failed to fetch studio"));
    }, []);

    const createStudio = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        createObj(e, { entity: "studio" }).then((res) => {
            if (res && res.data) {
                const newStudio: Studio = res.data;
                setStudios((prev) => [...prev, newStudio]);
            }
        });
    }


    const deleteStudio = (id: number) => {
        deleteObj(id, { entity: "studio" }, () => {
            setStudios((prev) => prev.filter((e) => e.id !== id));
        })
    }

    return (
        <CreatorForm<Studio>
            createObj={createStudio}
            deleteObj={deleteStudio}
            objectName="Studio"
            objects={studios}
        />
    )
}

const VoiceActorCreator: FC = () => {
    const [voiceActors, setVoiceActors] = useState<VoiceActor[]>([]);

    useEffect(() => {
        api.get<GoResponse>("/categories/?category=voice-actor")
            .then((res) => {
                setVoiceActors(res.data.data["voice-actor"]);
            })
            .catch((_: any) => toast.error("failed to fetch voice actors"));
    }, []);

    const createVoiceActor = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        createObj(e, { entity: "voice-actor" }).then((res) => {
            if (res && res.data) {
                const newActor: VoiceActor = res.data;
                setVoiceActors((prev) => [...prev, newActor]);
            }
        });
    }

    const deleteVoiceActor = (id: number) => {
        deleteObj(id, { entity: "voice-actor" }, () => {
            setVoiceActors((prev) => prev.filter((e) => e.id !== id));
        })
    };

    return (
        <CreatorForm<VoiceActor>
            createObj={createVoiceActor}
            deleteObj={deleteVoiceActor}
            objectName="Voice Actor"
            objects={voiceActors}
        />
    );
}

const GenreCreator: FC = () => {
    const [genres, setGenres] = useState<Genre[]>([]);

    useEffect(() => {
        api.get<GoResponse>("/categories/?category=genre").
            then((res: AxiosResponse<GoResponse>) => {
                if (res.data.code === 200 && res.data.data) {
                    setGenres(res.data.data.genre);
                }
            }).
            catch((_: any) => toast.error("failed to fetch genres"));
    }, [])

    const addGenre = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const formData = new FormData(e.currentTarget);

        api.post<GoResponse>("/auth/manage/genre/", formData)
            .then((res: AxiosResponse<GoResponse>) => {
                if (res.data.code === 200) {
                    const newGenre: Genre = res.data.data;
                    setGenres((prev) => [...prev, newGenre]);
                }
            })
            .catch((_: any) => toast.error("failed to add genre"));
    }

    const deleteGenre = (id: number) => {
        api.delete<GoResponse>(`/auth/manage/genre/${id}`).
            then((res: AxiosResponse<GoResponse>) => {
                if (res.data.code === 200) {
                    setGenres((prev) => prev.filter((e) => e.id !== id));
                }
            }).
            catch((_: any) => toast.error("genre already exists"));
    }

    return (
        <div>
            <h1 className="text-4xl font-extrabold pb-1">New <span className="text-rose-500">Genre</span></h1>
            <p className="text-gray-500 pb-5">Here you can manage genres. It provides you creating and deleting. Also you cannot create already exists genre.</p>

            <div>
                {genres.map((genre: Genre) => (
                    <React.Fragment key={genre.name}>
                        <div className="cursor-default flex-row badge badge-info gap-2 shadow-lg mr-2 mb-2 rounded-3xl transition-transform ease-in-out duration-300 hover:scale-105">
                            <button onClick={() => deleteGenre(genre.id)}>
                                <FontAwesomeIcon icon={faXmark} width={10} />
                            </button>
                            <p className="font-semibold">{genre.name}</p>
                        </div>
                    </React.Fragment>
                ))}
            </div>
            <form encType="multipart/form-data" onSubmit={(e) => addGenre(e)} className="md:w-1/3 flex">
                <div className="flex-row">
                    <CustomInput type="text" name="genre" placeholder="Write new genre..." required />
                </div>
                <button className="flex-row btn btn-circle btn-success mt-4 ml-3" type="submit"><FontAwesomeIcon icon={faPlus} width={15} /></button>
            </form>
        </div>
    )
}

type CreatableObject = Studio | Character | VoiceActor;

interface CreatorFormProps<T extends CreatableObject> {
    createObj: (e: React.FormEvent<HTMLFormElement>) => void;
    deleteObj: (id: number) => void;
    objectName: string;
    objects?: (T & { lastname?: string })[];
}

const CreatorForm = <T extends CreatableObject>({
    createObj,
    deleteObj,
    objectName,
    objects,
}: CreatorFormProps<T>) => {
    const [pic, setPic] = useState<string | undefined>(undefined);

    return (
        <div>
            {/* manage Zone */}
            <form
                className="md:w-3/4 flex flex-col md:flex-row justify-between items-start"
                id="actorForm"
                encType="multipart/form-data"
                onSubmit={(e) => createObj(e)}
            >
                <div className="w-full md:w-3/5">
                    <h1 className="text-3xl font-extrabold">
                        Add new <span className="text-lime-400">{objectName}</span>
                    </h1>
                    <p className="text-gray-500 pb-5">
                        Here you can create a new {objectName}. If you want to delete, just unroll the element with <b>See all your {objectName}s</b>. To update an {objectName}, just click on the chosen one, and you will be redirected to their page, where you can make changes.
                    </p>
                    <CustomInput type="text" name="name" placeholder="Name..." required />
                    {objectName === "Voice Actor" && (
                        <>
                            <CustomInput type="text" name="lastname" placeholder="Last name..." required />
                            <div className="w-fit pb-10">
                                <div className="label">
                                    <span className="label-text"></span>
                                    <span className="label-text-alt">Birth date</span>
                                </div>
                                <input
                                    type="date"
                                    name="birthdate"
                                    className="input input-bordered max-w-xs"
                                    required
                                />
                            </div>
                        </>
                    )}
                    {objectName === "Character" && (
                        <>
                            <CustomInput type="text" name="lastname" placeholder="Last name..." required />
                            <div className="pb-10">
                                <div className="label">
                                    <span className="label-text"></span>
                                    <span className="label-text-alt">Character information</span>
                                </div>
                                <textarea
                                    name="information"
                                    className="textarea textarea-bordered w-full"
                                    placeholder="Information about the character..."
                                    required
                                />
                            </div>
                        </>
                    )}
                    {objectName === "Studio" && (
                        <>
                            <div className="w-fit pb-10">
                                <div className="label">
                                    <span className="label-text"></span>
                                    <span className="label-text-alt">Established Date</span>
                                </div>
                                <input
                                    type="date"
                                    name="establishedDate"
                                    className="input input-bordered max-w-xs"
                                    required
                                />
                            </div>
                            <CustomInput type="text" name="website" placeholder="Website..." required />
                        </>
                    )}
                    <div className="w-fit">
                        <div className="label">
                            <span className="label-text"></span>
                            <span className="label-text-alt">Add image</span>
                        </div>
                        <input
                            type="file"
                            name="pic"
                            className="file-input w-full max-w-xs"
                            accept={ACCEPTED_IMAGE_TYPES}
                            onChange={(e) => loadImage(e, setPic)}
                        />
                    </div>
                </div>
                {pic && (
                    <div className="w-full md:w-1/3 md:pl-6 mt-6 md:mt-0 items-center justify-center">
                        <h1 className="text-xl font-semibold mb-2 text-center">Uploaded image</h1>
                        <Image
                            src={pic}
                            alt="Uploaded image"
                            className="w-full h-auto rounded-lg shadow-lg"
                            width={300}
                            objectFit="cover"
                            layout="responsive"
                            height={300}
                        />
                    </div>
                )}
            </form>

            <div className="pt-5 flex justify-end items-end">
                <button type="submit" form="actorForm" className="btn btn-outline btn-ghost">
                    <FontAwesomeIcon icon={faSave} width={20} />
                    <span>Save</span>
                </button>
            </div>

            {!objects && (
                <h1 className="text-4xl font-extrabold pt-10">There is no {objectName} yet...</h1>
            )}

            {/* menu with all objects */}
            {objects && (
                <div className="collapse bg-base-200 mt-10">
                    <input type="checkbox" />
                    <div className="collapse-title text-xl font-medium">
                        See all your {objectName}s
                    </div>
                    <div className="collapse-content">
                        {objects.map((object) => (
                            <React.Fragment key={object.id}>
                                <div className="text-lg flex pb-5">
                                    <button onClick={() => deleteObj(object.id)} className="btn btn-sm btn-circle mr-1">
                                        <FontAwesomeIcon icon={faXmark} width={20} className="text-red-600" />
                                    </button>
                                    {/* TODO: Add redirect link */}
                                    <Link href="#" className="text-blue-500 font-semibold hover:underline mt-0.5">
                                        {object.id}. {object.name} {object.lastname && object.lastname}
                                    </Link>
                                </div>
                            </React.Fragment>
                        ))}
                    </div>
                </div>
            )}
        </div>
    );
};
