import { ReactNode } from "react";

export default function FriendsLayout({ children }: { children: ReactNode }) {
    return (
        <section className="py-28 md:px-32 px-10">
            {children}
        </section>
    )
}