import { ChangeEvent, Dispatch, SetStateAction } from "react";

export const loadImage = (
    event: ChangeEvent<HTMLInputElement>,
    setImagePreview: Dispatch<SetStateAction<string | null>>
) => {
    const file = event.target.files?.[0];
    if (file) {
        const reader = new FileReader();
        reader.onloadend = () => {
            if (typeof reader.result === 'string') {
                setImagePreview(reader.result);
            } else {
                // Możesz obsłużyć przypadek, gdy reader.result nie jest stringiem
                console.error("Unexpected result type from FileReader");
            }
        };
        reader.readAsDataURL(file);
    }
};