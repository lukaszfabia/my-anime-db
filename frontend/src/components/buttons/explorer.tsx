import { faLightbulb } from "@fortawesome/free-solid-svg-icons"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import Link from "next/link"
import { FC } from "react"

export const ExplorerButton: FC = () => {
    return (
        <div className="flex justify-center items-center px-5">
            <button className="btn btn-ghost lg:btn-sm btn-md">
                <Link href="/explorer"><FontAwesomeIcon icon={faLightbulb} width={15} height={15} className="mr-1" /><span>Explorer</span></Link>
            </button>
        </div>
    )
}