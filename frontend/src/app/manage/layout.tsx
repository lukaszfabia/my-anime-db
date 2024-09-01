import { ReactNode } from "react";

export default function ManageLayout({ children }: { children: ReactNode }) {
    return (
        <section className="py-28 lg:px-32 px-5">
            {children}
        </section>
    )
}