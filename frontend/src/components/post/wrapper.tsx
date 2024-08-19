import { Post, User } from "@/types/models"
import { faEarth, faLock, faPenToSquare } from "@fortawesome/free-solid-svg-icons"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import Link from "next/link"
import { FC, ReactNode } from "react"
import Image from "next/image"
import transformTime from "@/lib/computeTime"

export const PostWrapper: FC<{ post: Post, user: User, children?: ReactNode, isReadOnly?: boolean }> = ({ post, user, children, isReadOnly = false }) => {
    return (
        <div className="pt-5">
            <div className="bg-base-300 md:p-10 p-5 rounded-2xl w-full relative shadow-lg">
                {/* manage menu */}
                {children}

                <div className="flex">
                    <div className="avatar flex-col flex justify-center items-center">
                        <div className="ring-primary ring-offset-base-100 w-14 rounded-full ring ring-offset-1 transition-all duration-200 ease-in-out hover:ring-offset-0">
                            <Link href={!isReadOnly ? "/profile" : `/user/${post.userId}`}>
                                <Image src={user?.picUrl} alt={`${user.username}'s avatar`} width={150} height={150} />
                            </Link>
                        </div>
                    </div>
                    <div className="sm:divider sm:divider-horizontal max-sm:px-2"></div>
                    <div className="flex flex-col">
                        <Link href={!isReadOnly ? "/profile" : `/user/${post.userId}`}>
                            <h1 className="text-blue-500 text-lg hover:underline">{user.username}</h1>
                        </Link>
                        <h1 className="text-sm">
                            <span>created {transformTime(post.createdAt)}</span>
                            <div className="tooltip tooltip-right" data-tip={`post is ${post.isPublic ? "public" : "private"}`}>
                                <FontAwesomeIcon icon={post.isPublic ? faEarth : faLock} className="ml-1 w-4" />
                            </div>
                        </h1>
                        {post.updatedAt !== post.createdAt && (
                            <h1 className="text-sm">
                                <span>edited {transformTime(post.updatedAt)}</span>
                                <FontAwesomeIcon icon={faPenToSquare} className="ml-1 w-4" />
                            </h1>
                        )}
                    </div>
                </div>
                <div className="py-5">
                    <h1 className="text-2xl font-semibold dark:text-black text-white">{post.title}</h1>
                    <p className="font-lato">{post.content}</p>
                </div>
                {post.image && (
                    <div className="flex items-center justify-center w-full">
                        <div className="py-5 md:w-1/2">
                            <div className="relative w-full max-w-xl">
                                <Image
                                    src={post.image}
                                    alt={post.title}
                                    objectFit="cover"
                                    layout="responsive"
                                    width={600}
                                    height={400}
                                    className="rounded-xl"
                                />
                            </div>
                        </div>
                    </div>
                )}

            </div>
        </div >
    )
}
