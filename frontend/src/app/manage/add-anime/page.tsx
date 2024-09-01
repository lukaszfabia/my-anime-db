"use client";

import { CustomInput } from "@/components/ui/forms/accountForm";
import api from "@/lib/api";
import { ACCEPTED_IMAGE_TYPES } from "@/lib/config";
import { GoResponse } from "@/types/responses";
import { faAdd, faXmark } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Link from "next/link";
import { FC, useEffect, useRef, useState } from "react";
import { toast } from "react-toastify";

interface SelectorProps {
    text: string;
    collection: string[];
}

export default function AddAnime() {
    const [alternativeTitles, setAlternativeTitles] = useState<string[]>([]);
    const altRef = useRef<HTMLInputElement | null>(null);

    const [animeType, setAnimeType] = useState<string>("");
    const [pegi, setPegi] = useState<string>("");
    const [status, setStatus] = useState<string>("");

    const [animeTypes, setAnimeTypes] = useState<string[]>([]);
    const [pegis, setPegis] = useState<string[]>([]);
    const [statuses, setStatuses] = useState<string[]>([]);

    useEffect(() => {
        api
            .get<GoResponse>("/categories/?category=anime_type&category=pegi&category=anime_status")
            .then((res) => {
                if (res.status === 200) {
                    setAnimeTypes(res.data.data.anime_type);
                    setPegis(res.data.data.pegi);
                    setStatuses(res.data.data.anime_status);
                }
            })
            .catch((_: any) => toast.error("failed to fetch anime types, pegi and status!"));
    }, []);

    const Selector: FC<SelectorProps> = ({ collection, text }) => {
        const handleChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
            const value: string = e.target.value;
            if (text === "Choose anime type") setAnimeType(value);
            if (text === "What is target group?") setPegi(value);
            if (text === "Set anime status") setStatus(value);
        };

        return (
            <div className="w-full max-w-xs py-3">
                <label className="label">
                    <span>{text}</span>
                </label>
                <select className="select select-bordered w-full max-w-xs" onChange={(e) => handleChange(e)} defaultValue={collection[0]}>
                    {collection.map((item: string) => (
                        <option key={item} value={item}>{item}</option>
                    ))}
                </select>
            </div>
        );
    };

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
                    <CustomInput placeholder="Other titles..." type="text" name="altTitles" inputRef={altRef} />
                    <button type="button" className="btn btn-primary btn-circle ml-3 mt-4" onClick={addAltTitle}><FontAwesomeIcon icon={faAdd} width={15} /></button>
                </div>
                <div className="flex flex-row flex-wrap justify-center md:justify-start">
                    {alternativeTitles.map((title: string) => (
                        <div key={title} className="badge badge-lg badge-secondary mr-2 mt-2 gap-1 rounded-lg">
                            <button type="button" onClick={() => removeAltTitle(title)}>
                                <FontAwesomeIcon icon={faXmark} width={15} />
                            </button>{title}
                        </div>
                    ))}
                </div>
            </div>
        );
    };

    const addAnime = (e: React.FormEvent<HTMLFormElement>) => {
        const form = new FormData(e.currentTarget);

        // iterate 
        form.forEach((value, key) => console.log(key, value));
    };

    return (
        <form encType="multipart/form-data" onSubmit={(e) => addAnime(e)} className="p-1 max-w-full overflow-x-hidden">
            <h1 className="font-extrabold text-4xl text-center md:text-left">Add <span className="text-blue-400">Anime</span></h1>
            <p className="text-gray-500 py-2 text-center md:text-left">Please remember that you can add <Link href="#" className="link link-info">roles</Link> - characters and voice actors to anime.</p>
            <div className="flex flex-col md:flex-row pt-10 justify-center md:justify-between">
                <div className="w-full md:w-1/3">
                    <p className="text-warning text-center md:text-left">Notice: Title must be <b>unique</b>.</p>
                    <CustomInput placeholder="" type="text" name="title" required>
                        <b>Title</b>
                    </CustomInput>
                </div>
                <div className="w-full md:w-1/2">
                    <AlternativeTitles />
                </div>
            </div>

            <div className="flex flex-col md:flex-row pt-10 justify-center md:justify-between">
                <div className="md:w-1/2 w-full">
                    <h1 className="text-2xl font-semibold text-center md:text-left">Select parameters</h1>
                    <p className="text-gray-500 text-center md:text-left">Contains important technical information.</p>
                    <Selector collection={animeTypes} text="Choose anime type" />
                    <Selector collection={pegis} text="What is target group?" />
                    <Selector collection={statuses} text="Set anime status" />
                </div>
                <div className="md:w-1/2 w-full">
                    <h1 className="text-2xl font-semibold text-center md:text-left">Some numbers</h1>
                    <p className="text-gray-500 text-center md:text-left">
                        Set number of episodes and duration of one episode.
                    </p>
                    <div className="w-72 md:w-fit mx-auto md:mx-0">
                        <CustomInput placeholder="e.g 12" type="number" name="episodes" required={status === "currently-airing" || status === "finished"}>
                            <b className="text-sm md:text-base">Episodes</b>
                        </CustomInput>
                        <CustomInput placeholder="e.g 24" type="number" name="episodesLength" required={status === "currently-airing" || status === "finished"}>
                            <b className="text-sm md:text-base">Duration</b>
                        </CustomInput>
                        <div className="divider divider-vertical"></div>
                        <h1 className="text-2xl font-semibold text-center md:text-left">Dates</h1>
                        <p className="text-gray-500 text-center md:text-left">You can set <b>optionally</b> start and finish date.</p>
                        <CustomInput placeholder="" type="date" name="startDate" disabled={status === "unknown"}>
                            <b>Start</b>
                        </CustomInput>

                        <CustomInput placeholder="" type="date" name="finishDate" disabled={status === "unknown" || status === "planned"}>
                            <b>Finish</b>
                        </CustomInput>
                    </div>
                </div>
            </div>

            <div className="flex flex-col md:flex-row pt-10 justify-center md:justify-between">
                <div className="md:w-2/5 w-full max-sm:py-5">
                    <h1 className="text-2xl font-semibold text-center md:text-left">Description</h1>
                    <p className="text-gray-500 text-center md:text-left">Write something about anime.</p>
                    <textarea className="textarea textarea-bordered w-full mt-2" placeholder="Description..." name="description"></textarea>
                </div>

                <div className="md:w-1/2 w-full">
                    <h1 className="text-2xl font-semibold text-center md:text-left">Upload cover</h1>
                    <p className="text-gray-500 text-center md:text-left">You can upload image.</p>
                    <input type="file" name="pic" id="pic" className="file file-input mt-5 w-full" accept={ACCEPTED_IMAGE_TYPES} />
                </div>
            </div>

            <div className="flex justify-center md:justify-end w-full md:w-5/6 mt-5">
                <button type="submit" className="btn btn-secondary">
                    <FontAwesomeIcon icon={faAdd} width={15} />
                    <span>Add anime</span>
                </button>
            </div>
        </form>
    );
}
