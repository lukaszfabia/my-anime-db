import { FC } from "react"

export const Footer: FC = () => {
    return (
        <footer className="footer footer-center text-base-content p-7">
            <aside>
                <p>myanime.db &copy; {new Date().getFullYear()} - All right reserved by <a className="link" href="https://lukaszfabia.vercel.app">Lukasz Fabia</a></p>
            </aside>
        </footer>
    )
} 