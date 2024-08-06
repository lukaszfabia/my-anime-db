import { StrengthLevel } from "@/types";
import { MutableRefObject } from "react";

const regexForGood: RegExp = new RegExp("^(?=.*[A-Za-z])(?=.*\\d)[A-Za-z\\d]{8,}$");
const regexForStrong: RegExp = new RegExp("^(?=.*[A-Za-z])(?=.*\\d)(?=.*[@$!%*#?&])[A-Za-z\\d@$!%*#?&]{8,}$");

export default function validatePassword(
    setPasswordStrength: (lvl: StrengthLevel) => void,
    passRef: MutableRefObject<HTMLInputElement | null>
) {
    const password = (passRef.current as HTMLInputElement).value as string;

    const submitButton = document.getElementById("submitButton") as HTMLButtonElement;

    if (regexForStrong.test(password)) {
        setPasswordStrength("strong");
        submitButton.disabled = false;
    } else if (regexForGood.test(password)) {
        setPasswordStrength("good");
        submitButton.disabled = false;
    } else {
        setPasswordStrength("weak");
        submitButton.disabled = true;
    }

}
