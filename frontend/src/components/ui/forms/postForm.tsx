import { PostFormProps } from "@/types";
import { faPaperPlane, faPlus, faSave } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { ChangeEvent, FC, useState } from "react";
import Image from "next/image";
import { loadImage } from "@/lib/loadImage";


interface ImageProps {
    imagePreview: string | ArrayBuffer | null;
    defaultValues: { image?: string; title?: string };
    isHidden?: boolean;
}

const DisplayImage: FC<ImageProps> = ({ imagePreview, defaultValues, isHidden = false }) => {
    const imageUrl = imagePreview ? imagePreview as string : defaultValues?.image;
    const altText = defaultValues?.title || "Image preview";

    return (
        imageUrl ? (
            <div className="flex items-center justify-center">
                <div className={`mb-4 w-64 ${isHidden ? "md:hidden" : "max-md:hidden"}`}>
                    <Image
                        src={imageUrl}
                        alt={imagePreview ? "your image" : altText}
                        objectFit="cover"
                        layout="responsive"
                        width={400}
                        height={400}
                        className="rounded-xl shadow-lg text-center"
                    />
                </div>
            </div>
        ) : null
    );
}

export const PostForm: FC<PostFormProps> = ({ submitFunc, defaultValues }) => {
    const [defTitle, setDefTitle] = useState<string>(defaultValues?.title || "");
    const [defContent, setDefContent] = useState<string>(defaultValues?.content || "");
    const [defIsPublic, setDefIsPublic] = useState<boolean>(defaultValues ? defaultValues.isPublic : true);

    const [imagePreview, setImagePreview] = useState<string | ArrayBuffer | null>(null);

    return (
        <form className="form-control py-10 flex flex-col" encType="multipart/form-data" onSubmit={submitFunc}>
            <div className="flex">
                <h1 className="flex-1 text-4xl text-center md:text-left font-extrabold py-5">
                    {defaultValues ? "Edit" : "Create"} a <span className="text-sky-500">post <FontAwesomeIcon icon={faPlus} />
                    </span>
                </h1>
                <h1 className="flex-1 text-4xl text-center md:text-left font-extrabold py-5 max-md:hidden">
                    Image preview
                </h1>
            </div>
            <div className="flex flex-col md:flex-row flex-grow">
                <div className="flex-1 md:mr-4 mb-4 md:mb-0">
                    <input
                        type="text"
                        placeholder="Type topic/title"
                        className="input input-bordered w-full my-4"
                        name="title"
                        required={!defaultValues}
                        value={defTitle}
                        onChange={e => setDefTitle(e.target.value)}
                    />

                    <textarea
                        className="textarea textarea-bordered w-full h-1/3 my-4"
                        placeholder="Write your post here..."
                        value={defContent}
                        required={!defaultValues}
                        name="content"
                        onChange={e => setDefContent(e.target.value)}
                    ></textarea>
                    <input
                        type="file"
                        name="image"
                        onChange={(e: ChangeEvent<HTMLInputElement>,) => loadImage(e, setImagePreview)}
                        accept=".jpg,.jpeg,.png,.webp"
                        className="file-input file-input-bordered w-full my-4"
                    />
                    {imagePreview && <DisplayImage imagePreview={imagePreview} defaultValues={{ title: defaultValues?.title, image: defaultValues?.image }} isHidden />}
                    <label className="label cursor-pointer flex items-end justify-end">
                        <span className="label-text mr-2 lg:mb-0.5">Make public</span>
                        <input
                            name="isPublic"
                            type="checkbox"
                            checked={defIsPublic}
                            className="checkbox"
                            onChange={e => setDefIsPublic(e.target.checked)}
                        />
                    </label>
                </div>
                <div className="flex-1 md:ml-4 mt-4 md:mt-0">
                    {imagePreview && <DisplayImage imagePreview={imagePreview} defaultValues={{ title: defaultValues?.title, image: defaultValues?.image }} />}
                </div>
            </div>
            <div className="flex max-md:justify-center md:py-10">
                <div className="lg:w-1/2 w-full flex lg:flex-col items-center justify-center">
                    <button className="btn btn-outline btn-success w-1/4 max-sm:w-1/2" type="submit">
                        <span>{defaultValues ? "Save" : "Post"}</span>
                        <FontAwesomeIcon icon={defaultValues ? faSave : faPaperPlane} />
                    </button>
                </div>
            </div>
        </form>
    );
}
