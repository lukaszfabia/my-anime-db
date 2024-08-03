import { ReactNode } from "react";

export default function ProfileLayout({ children }: { children: ReactNode }) {
    return (
        <section className="py-32 px-20">
            {children}
        </section>
    )
}