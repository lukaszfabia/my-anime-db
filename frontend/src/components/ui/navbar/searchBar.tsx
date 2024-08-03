import { FC } from "react"

import { faMagnifyingGlass } from "@fortawesome/free-solid-svg-icons"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"

export const SearchBar: FC = () => {
    return (
        <label className="hidden md:visible input input-bordered md:input-md md:flex items-center gap-2">
            <FontAwesomeIcon icon={faMagnifyingGlass} width={15} height={15} />
            <input type="text" className="grow" placeholder="Search..." autoComplete="off" />
        </label>
    )
}