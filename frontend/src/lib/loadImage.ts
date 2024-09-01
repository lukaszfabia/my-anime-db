import { ChangeEvent, Dispatch, SetStateAction } from "react";

export const loadImage = (
    event: ChangeEvent<HTMLInputElement>,
    setImagePreview: Dispatch<SetStateAction<string | undefined>>
) => {
    const file = event.target.files?.[0];
    if (file) {
        const reader = new FileReader();
        reader.onloadend = () => {
            if (typeof reader.result === 'string') {
                setImagePreview(reader.result);
            }
        };
        reader.readAsDataURL(file);
    }
};