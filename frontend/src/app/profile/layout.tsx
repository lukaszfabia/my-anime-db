import { ReactNode } from "react";

export default function ProfileLayout({ children }: { children: ReactNode }) {
    return (
        <section className="py-28 md:px-32 px-5">
            {children}
        </section>
    )
}