import { ChangeEvent, Dispatch, SetStateAction } from "react";

export const loadImage = (event: ChangeEvent<HTMLInputElement>, setImagePreview: Dispatch<SetStateAction<string | ArrayBuffer | null>>) => {
    const file = event.target.files?.[0];
    if (file) {
        const reader = new FileReader();
        reader.onloadend = () => {
            setImagePreview(reader.result);
        };
        reader.readAsDataURL(file);
    }
};