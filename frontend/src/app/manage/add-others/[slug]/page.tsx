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
import { Selector } from "@/components/ui/selector";
import { createName } from "@/lib/name";
import { convertTime } from "@/lib/computeTime";
import { getImageUrl } from "@/lib/getImageUrl";

const availableTypes: CreatorProps[] = [
    { entity: "voice_actor" },
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
    entity: "voice_actor" | "character" | "studio" | "genre";
}

const Creator: FC<CreatorProps> = ({ entity }) => {
    switch (entity) {
        case "voice_actor":
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
    const [characterToUpdate, setCharacterToupdate] = useState<Character | null>(null);
    useEffect(() => {
        api.get<GoResponse>("/categories/?category=character")
            .then((res) => {
                setCharacters(res.data.data.character);
            })
            .catch((_: any) => toast.error("failed to fetch characters"));
    }, []);

    const createCharacter = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        createObj(e, { entity: "character" }, characterToUpdate ? characterToUpdate.id : null).then((res) => {
            if (res && res.data) {
                const newCharacter: Character = res.data
                setCharacters((prev) => [...prev, newCharacter]);
                toast.success(`${createName<Character>(newCharacter)} created successfully`);
            }
        });
    }

    const deleteCharacter = (id: number) => {
        deleteObj(id, { entity: "character" }, () => {
            setCharacters((prev) => prev.filter((e) => e.id !== id));
        })
    };

    const handleCharacterToUpdate = (e: React.ChangeEvent<HTMLSelectElement>) => {
        const v = e.currentTarget.value;
        const character = characters.find((e) => createName<Character>(e) === v);

        if (character) {
            const id = character.id;
            api.get<GoResponse>(`/character/${id}`)
                .then((res) => setCharacterToupdate(res.data.data))
                .catch((_: any) => toast.error("failed to fetch character"));

            console.log(character.picUrl);
        } else {
            setCharacterToupdate(null);
        }
    }

    return <CreatorForm
        createObj={createCharacter}
        deleteObj={deleteCharacter}
        objectName="Character"
        objects={characters}
        reset={() => setCharacterToupdate(null)}
        defaultValues={characterToUpdate}>
        <Selector text="Select character to update" collection={characters.map((e) => createName<Character>(e))} name="character" lastElem={createName<Character>(characterToUpdate)} handler={handleCharacterToUpdate} />
    </CreatorForm>;
}

const StudioCreator: FC = () => {
    const [studios, setStudios] = useState<Studio[]>([]);
    const [studioToUpdate, setStudioToUpdate] = useState<Studio | null>(null);
    useEffect(() => {
        api.get<GoResponse>("/categories/?category=studio")
            .then((res) => {
                setStudios(res.data.data.studio);
            })
            .catch((_: any) => toast.error("failed to fetch studio"));
    }, []);

    const createStudio = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        createObj(e, { entity: "studio" }, studioToUpdate ? studioToUpdate.id : null).then((res) => {
            if (res && res.data) {
                const newStudio: Studio = res.data;
                setStudios((prev) => [...prev, newStudio]);
                toast.success(`${newStudio.name} ${studioToUpdate ? "updated" : "created"} successfully`);
            }
        });
    }


    const deleteStudio = (id: number) => {
        deleteObj(id, { entity: "studio" }, () => {
            setStudios((prev) => prev.filter((e) => e.id !== id));
        })
    }

    const handleSelectStudio = (e: React.ChangeEvent<HTMLSelectElement>) => {
        const v = e.currentTarget.value;
        const studio = studios.find((e) => e.name === v);

        if (studio) {
            setStudioToUpdate(studio);
            api.get<GoResponse>(`/studio/${studio.id}`).then((res) => {
                setStudioToUpdate(res.data.data);
            }).catch((_: any) => toast.error("failed to fetch studio"));
        } else {
            setStudioToUpdate(null);
        }
    }

    return (
        <CreatorForm<Studio>
            createObj={createStudio}
            deleteObj={deleteStudio}
            objectName="Studio"
            objects={studios}
            reset={() => setStudioToUpdate(null)}
            defaultValues={studioToUpdate}
        >
            <Selector text="Select studio to update" collection={studios.map((e) => e.name)} name="studio" handler={handleSelectStudio} lastElem={studioToUpdate?.name} />
        </CreatorForm>
    )
}

const VoiceActorCreator: FC = () => {
    const [voiceActors, setVoiceActors] = useState<VoiceActor[]>([]);
    const [voiceActorToUpdate, setVoiceActorToUpdate] = useState<VoiceActor | null>(null);

    useEffect(() => {
        api.get<GoResponse>("/categories/?category=voice_actor")
            .then((res) => {
                setVoiceActors(res.data.data.voice_actor);
            })
            .catch((_: any) => toast.error("failed to fetch voice actors"));
    }, []);

    const createVoiceActor = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        createObj(e, { entity: "voice_actor" }, voiceActorToUpdate ? voiceActorToUpdate.id : null).then((res) => {
            if (res && res.data) {
                const newActor: VoiceActor = res.data;
                setVoiceActors((prev) => [...prev, newActor]);
                toast.success(`${newActor.lastname} ${newActor.name} created successfully`);
            }
        });
    }

    const deleteVoiceActor = (id: number) => {
        deleteObj(id, { entity: "voice_actor" }, () => {
            setVoiceActors((prev) => prev.filter((e) => e.id !== id));
        })
    };

    const handleSelectVoiceActor = (e: React.ChangeEvent<HTMLSelectElement>) => {
        const v = e.currentTarget.value;
        const actor = voiceActors.find((e) => e.lastname + " " + e.name === v);

        if (actor) {
            setVoiceActorToUpdate(actor);

            api.get<GoResponse>(`/voice_actor/${actor.id}`).then((res) => {
                setVoiceActorToUpdate(res.data.data);
            }).catch((_: any) => toast.error("failed to fetch voice actor"));
        } else {
            setVoiceActorToUpdate(null);
        }
    }

    return (
        <CreatorForm<VoiceActor>
            createObj={createVoiceActor}
            deleteObj={deleteVoiceActor}
            objectName="Voice Actor"
            objects={voiceActors}
            reset={() => setVoiceActorToUpdate(null)}
            defaultValues={voiceActorToUpdate}
        >
            <Selector text="Select voice actor to update" collection={voiceActors.map((e) => createName<VoiceActor>(e))} name="voiceActor" lastElem={createName<VoiceActor>(voiceActorToUpdate)} handler={handleSelectVoiceActor} />
        </CreatorForm>
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
    defaultValues?: (T & { lastname?: string } & { picUrl?: string } & { logoUrl?: string } & { birthdate?: Date } & { information?: string } & { establishedDate?: Date } & { website?: string }) | null;
    children?: React.ReactNode;
    reset?: () => void;
}

const CreatorForm = <T extends CreatableObject>({
    createObj,
    deleteObj,
    objectName,
    objects,
    children,
    reset,
    defaultValues,
}: CreatorFormProps<T>) => {
    const [pic, setPic] = useState<string | undefined>(undefined);

    useEffect(() => {
        setPic(getImageUrl(defaultValues?.picUrl) || getImageUrl(defaultValues?.logoUrl) || undefined);
    }, [defaultValues]);

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
                        Here you can create a new {objectName}. If you want to delete, just unroll the element with <b>See all your {objectName}s</b>. To update an {objectName}, just select your actor and make changes.
                    </p>
                    {children && (
                        <div className="flex flex-row">
                            {children}
                            <button onClick={reset} type="button" className="btn btn-circle btn-error ml-2 mt-3"><FontAwesomeIcon icon={faXmark} width={15} /></button>
                        </div>
                    )}

                    <div className="flex flex-col sm:flex-row">
                        <CustomInput type="text" name="name" placeholder="Name..." defaultValue={defaultValues?.name} />
                        <div className="mx-2"></div>
                        {objectName !== "Studio" && (
                            <CustomInput type="text" name="lastname" placeholder="Last name..." defaultValue={defaultValues?.lastname} />
                        )}
                    </div>
                    {objectName === "Voice Actor" && (
                        <div className="md:w-fit pb-10">
                            <p className="text-gray-500 my-1 ml-1">Birthdate</p>
                            <CustomInput defaultValue={convertTime(defaultValues?.birthdate)} type="date" name="birthdate" placeholder="Birthdate..." />
                        </div>
                    )}
                    {objectName === "Character" && (
                        <>
                            <div className="pb-10">
                                <p className="text-gray-500 my-1 ml-1">Character information</p>
                                <textarea
                                    name="information"
                                    className="textarea textarea-bordered w-full"
                                    placeholder="Information about the character..."
                                    required
                                    defaultValue={defaultValues?.information}
                                />
                            </div>
                        </>
                    )}
                    {objectName === "Studio" && (
                        <div className="flex sm:flex-row flex-col">
                            <div className="sm:w-1/3 w-fit">
                                <CustomInput defaultValue={convertTime(defaultValues?.establishedDate)} type="date" name="establishedDate" placeholder="" />
                            </div>
                            <div className="mx-2"></div>
                            <div className="sm:w-2/3">
                                <CustomInput defaultValue={defaultValues?.website} type="text" name="website" placeholder="Website..." required={false} />
                            </div>
                        </div>
                    )}
                    <div className="w-fit">
                        <h1 className="text-lg font-semibold my-2 ml-1">{objectName === "Studio" ? "Add logo" : "Add image"}</h1>
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
                <div className="collapse bg-base-200 mt-10 shadow-xl">
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
