import { ReactNode } from "react";

export default function SettingsLayout({ children }: { children: ReactNode }) {
    return (
        <section className="py-28 md:px-20 px-10">
            {children}
        </section>
    )
}