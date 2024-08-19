import { ReactNode } from "react";

export default function LoginLayout({ children }: { children: ReactNode }) {
    return (
        <section className="md:py-40 py-20 md:px-32 px-5">
            {children}
        </section>
    )
}