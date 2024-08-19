import { FC } from "react"

import { faMagnifyingGlass } from "@fortawesome/free-solid-svg-icons"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"

export const SearchBar: FC = () => {
    return (
        <div className="flex items-center justify-center lg:py-0 py-5">
            <label className="input input-sm lg:input-bordered md:input-md md:flex items-center gap-2">
                <FontAwesomeIcon icon={faMagnifyingGlass} width={15} height={15} className="max-md:mr-1" />
                <input type="text" className="grow" placeholder="Search..." autoComplete="off" name="searchbar" />
            </label>
        </div>
    )
}