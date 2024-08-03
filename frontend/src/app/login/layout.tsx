import { ReactNode } from "react";

export default function LoginLayout({ children }: { children: ReactNode }) {
    return (
        <section className="py-40 md:px-32 px-10">
            {children}
        </section>
    )
}