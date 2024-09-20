import Link from "next/link";
import { FC, ReactNode, useState } from "react";
import Image from "next/image";
import React from "react";

export interface MenuProps {
    cards: ButtonWithBackgroundPicProps[];
    children?: ReactNode;
}

export const Menu: FC<MenuProps> = ({ cards, children }) => {
    return (
        <div>
            <>
                {children}
            </>
            <div className="py-5"></div>
            <div className={`flex flex-col sm:flex-row items-center justify-between mx-auto ${cards.length < 4 ? "max-w-3xl" : " max-w-5xl"}`}>
                {cards.map((e: ButtonWithBackgroundPicProps, i: number) => (
                    <React.Fragment key={e.title}>
                        <ButtonWithBackgroundPic link={e.link} imageUrl={e.imageUrl} title={e.title} content={e.content} btnText={e.btnText && e.btnText} addMargin={i % 2 == 1} />
                    </React.Fragment>
                ))}
            </div>
        </div>
    )
}

export interface ButtonWithBackgroundPicProps {
    imageUrl: string;
    title: string;
    content?: string;
    link: string;
    addMargin?: boolean;
    btnText?: string;
}

export const ButtonWithBackgroundPic: FC<ButtonWithBackgroundPicProps> = ({ link, imageUrl, title, content, addMargin, btnText = "Get started" }) => {
    const [isHovered, setIsHovered] = useState<boolean>(false);

    return (
        <div className={`flex-grow max-sm:my-5 hero h-96 transition-transform duration-300 ease-in-out hover:scale-105 ${addMargin && "mx-5"}`}
            onMouseOver={() => setIsHovered(true)}
            onMouseLeave={() => setIsHovered(false)}>
            <div className="relative w-full h-full">
                <Image
                    src={imageUrl}
                    alt={title}
                    id={title}
                    layout="fill"
                    objectFit="cover"
                    className={`hero-overlay transition-opacity duration-300 ease-in-out rounded-xl 
                        ${isHovered ? "dark:opacity-90 opacity-70" : "dark:opacity-60 opacity-20"}`}
                />
            </div>
            <div className="hero-content text-neutral-content text-center rounded-xl">
                <div className="max-w-md">
                    <h1 className="mb-5 text-4xl font-bold text-black dark:text-white">{title}</h1>
                    {content && <p className="mb-5">{content}</p>}
                    <Link className="btn glass" href={link}>{btnText}</Link>
                </div>
            </div>
        </div>
    );
};
