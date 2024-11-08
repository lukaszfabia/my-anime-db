"use client";

import { Avatar } from "@/components/person";
import { useAuth } from "@/components/providers/auth";
import { DialogWindow } from "@/components/ui/dialog";
import { CustomInput } from "@/components/ui/forms/accountForm";
import { Spinner } from "@/components/ui/spinner";
import api from "@/lib/api";
import { ACCEPTED_IMAGE_TYPES } from "@/lib/config";
import { getImageUrl } from "@/lib/getImageUrl";
import { loadImage } from "@/lib/loadImage";
import { Anchor } from "@/types";
import { User } from "@/types/models";
import { GoResponse } from "@/types/responses";
import { faCheck, faEnvelope, faLock, faLockOpen, faUser } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Link from "next/link";
import { redirect, useRouter } from "next/navigation";
import React, { ChangeEvent, FC, useRef, useState } from "react";
import OTPInput from "react-otp-input";
import { toast } from "react-toastify";

const chapters: Anchor[] = [
    { title: "Basic info", id: "basic-info" },
    { title: "About you", id: "about-you" },
    { title: "Security", id: "security" },
    { title: "Save", id: "save" },
];

const Contents: FC = () => {
    return (
        <aside className="max-md:hidden bg-base-100 w-1/6 h-screen">
            <ul className="flex flex-col fixed">
                {chapters.map((chapter: Anchor, i: number) => (
                    <React.Fragment key={chapter.id}>
                        {chapter.id === "save" && (<div className="divider" key={`divider-${i}`}></div>)}
                        <Link href={`#${chapter.id}`}>
                            <li className={`btn btn-ghost btn-md lg:btn-lg my-2 text-lg`}>
                                {chapter.title}<span className="text-violet-500"> #</span>
                            </li>
                        </Link>
                    </React.Fragment>
                ))}
            </ul>
        </aside>

    );
};

const AnchorElement: FC<{ anchor: Anchor }> = ({ anchor }) => {
    return (
        <span id={anchor.id} className="relative pt-20">
            {anchor.title}
            <Link href={`#${anchor.id}`} className="text-violet-500 transition-all duration-300 ease-in-out opacity-0 hover:opacity-100"> #</Link>
        </span>
    );
}

const BasicInfo: FC = () => {

    const [isPasswordVisible, setIsPasswordVisible] = useState<boolean>(false);
    const passRef = useRef<HTMLInputElement | null>(null);

    return (
        <div>
            <h1 className="text-3xl font-extrabold text-black dark:text-white"><AnchorElement anchor={chapters[0]} /></h1>
            <p className="text-sm text-gray-600 py-1">You can change your basic information, including your password and email, here. Make sure to save your changes once you're done.</p>

            <div className="form-control md:w-1/2 w-full">
                <CustomInput name="username" type="text" placeholder="Username..." required={false}>
                    <FontAwesomeIcon icon={faUser} width={12} height={12} />
                </CustomInput>


                <CustomInput name="email" type="email" placeholder="Email..." required={false}>
                    <FontAwesomeIcon icon={faEnvelope} width={12} height={12} />
                </CustomInput>

                <CustomInput name="password" required={false} type={isPasswordVisible ? 'text' : 'password'} placeholder="Password..." inputRef={passRef}>
                    <div className="tooltip md:tooltip-left tooltip-bottom" data-tip="Show password">
                        <button
                            type="button"
                            onClick={() => setIsPasswordVisible(!isPasswordVisible)}
                            className="animate-pulse"
                            tabIndex={-1}
                        >
                            <FontAwesomeIcon
                                icon={isPasswordVisible ? faLockOpen : faLock}
                                width={12}
                                height={12}
                            />
                        </button>
                    </div>
                </CustomInput>
            </div>
        </div>
    );
}

const Security: FC<{ user: User }> = ({ user }) => {
    const modalRef = useRef<HTMLDialogElement | null>(null);
    const { refreshUser } = useAuth();
    const [pin, setPin] = useState<string>("");

    const handleSend = () => {
        api.post<GoResponse>("/auth/account/send-code/")
            .then(() => {
                modalRef.current?.showModal();
                toast.success("Check your email for a verification link!");
                setPin("");
            })
            .catch((_: any) => {
                toast.error("An error occurred");
            });
    }

    const handleVerify = () => {
        const formData = new FormData();
        formData.append("code", pin);
        api.post<GoResponse>("/auth/account/verify/", formData)
            .then(() => {
                refreshUser();
            })
            .catch((_: any) => {
                toast.error("An error occurred");
            });
    }

    return (
        <div>
            <h1 className="text-3xl font-extrabold text-black dark:text-white"><AnchorElement anchor={chapters[2]} /></h1>
            <>
                <p className="text-sm text-gray-600 pb-3 pt-1">Here you can verify your account. After click, you will get a message on your email with a code to confirm.</p>
                {user.isVerified ? (
                    <div className="flex-col flex items-center justify-center">
                        <div role="alert" className="alert alert-success md:w-1/2 w-full flex-col mt-4 px-15">
                            <FontAwesomeIcon icon={faCheck} />
                            <span>Your account has been confirmed!</span>
                        </div>
                    </div>
                ) : (
                    <div className="flex-col flex items-center justify-center">
                        <button className="btn btn-info btn-wide mt-5" onClick={handleSend}>
                            Verify me
                        </button>
                        <DialogWindow modalRef={modalRef} short>
                            <form encType="multipart/form-data" onSubmit={handleVerify}>
                                <p className="text-sm dark:text-gray-600 text-gray-400">We sent confirmation code on <span className="font-semibold text-black dark:text-white">{user.email}</span>. Enter your received code down below.</p>

                                <div className="flex items-center justify-center pt-10 w-full h-full">
                                    <OTPInput
                                        value={pin}
                                        onChange={setPin}
                                        numInputs={6}
                                        inputType="number"
                                        inputStyle={{ height: "3rem", width: "3rem", fontSize: "1.5rem", borderRadius: "10px" }}
                                        renderSeparator={<span className="px-3">-</span>}
                                        renderInput={(props) => <input {...props} />}
                                    />
                                </div>

                                <div className="text-center pt-10">
                                    <button className="btn btn-outline btn-success btn-wide" type="submit">Confirm</button>
                                </div>
                            </form>
                        </DialogWindow>
                    </div>
                )}
            </>
        </div>
    )
}

const AboutYou: FC<{ user: User }> = ({ user }) => {
    const [pic, setPic] = useState<string | undefined>(user.picUrl ? getImageUrl(user.picUrl) : undefined);

    return (
        <div className="flex">
            <div className="flex-col py-5">
                <h1 className="text-3xl font-extrabold text-black dark:text-white"><AnchorElement anchor={chapters[1]} /></h1>
                <p className="text-sm text-gray-600 mt-2 mb-4">Change your additional informations like bio, avatar or website.</p>
                <label htmlFor="pic" className="form-control  w-full md:max-w-md">
                    <div className="label">
                        <span className="label-text">Pick a new avatar</span>
                    </div>
                    <input
                        type="file"
                        name="pic"
                        id="pic"
                        className="file-input file-input-bordered  w-full md:max-w-md"
                        onChange={(e: ChangeEvent<HTMLInputElement>) => loadImage(e, setPic)}
                        accept={ACCEPTED_IMAGE_TYPES}
                    />
                </label>

                <div className="md:hidden flex-col flex items-center justify-center py-5">
                    <h1 className="text-3xl font-extrabold py-4">Current avatar</h1>
                    <div className="w-1/2">
                        <Avatar picUrl={pic} name={user.username} />
                    </div>
                </div>

                <label htmlFor="bio" className="form-control  w-full md:max-w-md py-5">
                    <div className="label">
                        <span className="label-text">Your bio</span>
                        <span className="label-text-alt">Max length is 100</span>
                    </div>
                    <textarea id="bio" name="bio" className="textarea textarea-bordered h-24 max-w-md" placeholder="Tell us something about yourself..." maxLength={100}></textarea>
                </label>

                <label htmlFor="website" className="form-control w-full md:max-w-md py-5">
                    <div className="label">
                        <span className="label-text">Any website?</span>
                    </div>
                    <input id="website" name="website" type="text" placeholder="Type here..." className="input input-bordered w-full max-w-md" maxLength={50} />
                </label>
            </div>
            <div className="md:px-10"></div>
            <div className="max-md:hidden w-1/3 flex items-center justify-center">
                <div className="flex-col py-5">
                    <h1 className="text-3xl font-extrabold py-4 text-center">Current avatar</h1>
                    <Avatar picUrl={pic} name={user.username} />
                </div>
            </div>
        </div>
    );
}

const Summary: FC = () => {
    return (
        <div className="py-5">
            <h1 className="text-3xl font-extrabold text-black dark:text-white"><AnchorElement anchor={chapters[3]} /></h1>
            <p className="text-sm text-gray-600 py-1">If you're happy with them, save them and you're done!</p>
            <div className="flex md:items-end md:justify-end justify-center items-center py-5">
                <Link href="/profile" className="btn">Cancel</Link>
                <div className="px-4"></div>
                <button className="btn btn-warning" type="submit" form="edit">Save changes</button>
            </div>
        </div>
    );
}


export default function Settings() {
    const { user, loading, refreshUser } = useAuth();
    const router = useRouter();

    if (!user) {
        redirect("/login");
    }

    if (loading) {
        return <Spinner />
    }

    const handleEdit = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const formData = new FormData(e.currentTarget);

        api.put<GoResponse>("/auth/account/me/", formData)
            .then(() => {
                toast.success("Your account has been updated!");
                refreshUser();
                router.push("/profile");
            })
            .catch((_: any) => {
                toast.error("Too weak password or username or email already exists");
            });
    }

    return (
        <>
            <div className="flex">
                <Contents />

                <div className="md:divider md:divider-horizontal"></div>
                <div className="md:w-5/6">
                    <form className="pb-5" id="edit" encType="multipart/form-data" onSubmit={(e) => handleEdit(e)}>
                        <BasicInfo />
                        <div className="py-5"></div>
                        <AboutYou user={user} />
                    </form>
                    <Security user={user} />
                    <Summary />
                </div>
            </div>
        </>
    )
}
