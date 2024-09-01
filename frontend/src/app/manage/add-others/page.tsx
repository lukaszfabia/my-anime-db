"use client";
import { ButtonWithBackgroundPicProps, Menu } from "@/components/manage/menu";
import { useAuth } from "@/components/providers/auth";
import { Spinner } from "@/components/ui/spinner";

export default function AddOthers() {
    const { user, loading } = useAuth();

    if (!user || loading) return <Spinner />;

    if (!user.isMod) return null;

    const props: ButtonWithBackgroundPicProps[] = [
        { imageUrl: "/images/blue_lock.jpg", link: "/manage/add-others/voice-actor", title: "Voice Actor" },
        { imageUrl: "/images/hyouka.png", link: "/manage/add-others/character", title: "Character", },
        { imageUrl: "/images/kimi.jpg", link: "/manage/add-others/studio", title: "Studio", },
        { imageUrl: "/images/tengoku.jpeg", link: "/manage/add-others/genre", title: "Genre", },
    ]

    return (
        <Menu cards={props}>
            <h1 className="text-center text-5xl font-extrabold">Add others</h1>
            <p className="text-sm text-center text-gray-500 py-4">
                To enrich anime and our db. To update date go to page with given object.
            </p>
        </Menu>
    );
}
