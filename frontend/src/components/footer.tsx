import Link from "next/link"
import { FC } from "react"

export const Footer: FC = () => {
    return (
        <footer className="footer footer-center text-base-content p-4 text-wrap">
            <aside>
                <p>Copyright &copy; {new Date().getFullYear()} - All right reserved by <Link href="https://lukaszfabia.vercel.app" target="_blank" className="text-blue-600 no-underline font-semibold">Lukasz Fabia</Link></p>
            </aside>
        </footer>
    )
} 