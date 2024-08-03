import Link from "next/link"
import { FC } from "react"

export const Footer: FC = () => {
    return (
        <footer className="footer footer-center text-base-content p-4">
            <aside>
                <p>Copyright &copy; {new Date().getFullYear()} - All right reserved by <Link href="lukaszfabia.vercel.app" className="text-blue-600 no-underline font-semibold">Lukasz Fabia</Link></p>
            </aside>
        </footer>
    )
} 