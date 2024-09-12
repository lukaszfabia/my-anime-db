import { ReactNode } from "react";

export default function AnimeLayout({ children }: { children: ReactNode }) {
    return (
        <section className="py-28 lg:px-20 px-5">
            {children}
        </section>
    )
}