import { ReactNode } from "react";

export default function UserLayout({ children }: { children: ReactNode }) {
    return (
        <section className="py-28 min-[1024px]:px-32 md:px-10 px-5">
            {children}
        </section>
    )
}