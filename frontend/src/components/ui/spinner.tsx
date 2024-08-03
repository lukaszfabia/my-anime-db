import { FC } from "react";

export const Spinner: FC = () => {
    return (
        <div className="flex items-center justify-center">
            <span className="loading loading-dots loading-md"></span>
        </div>
    )
}