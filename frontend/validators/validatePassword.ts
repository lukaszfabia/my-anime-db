import { StrengthLevel } from "@/types";
import { MutableRefObject } from "react";

const regexForGood: RegExp = new RegExp("^(?=.*[A-Za-z])(?=.*\\d)[A-Za-z\\d]{8,}$");
const regexForStrong: RegExp = new RegExp("^(?=.*[A-Za-z])(?=.*\\d)(?=.*[@$!%*#?&])[A-Za-z\\d@$!%*#?&]{8,}$");

export default function validatePassword(
    setPasswordStrength: (lvl: StrengthLevel) => void,
    passRef: MutableRefObject<HTMLInputElement | null>
) {
    const passwordField = passRef.current as HTMLInputElement;

    if (!passwordField) {
        console.error("Password field is null");
        return;
    }

    const submitButton = document.getElementById("submitButton") as HTMLButtonElement;
    if (!submitButton) {
        console.error("Submit button not found");
        return;
    }

    const password = passwordField.value;

    const passwordLabel = document.getElementById("passwordLabel") as HTMLLabelElement;

    passwordLabel.classList.remove("input-warning", "input-error", "input-success");

    if (regexForStrong.test(password)) {
        setPasswordStrength("strong");
        passwordLabel.classList.add("input-success");
        submitButton.disabled = false;
    } else if (regexForGood.test(password)) {
        setPasswordStrength("good");
        passwordLabel.classList.add("input-warning");
        submitButton.disabled = false;
    } else {
        setPasswordStrength("weak");
        passwordLabel.classList.add("input-error");
        submitButton.disabled = true;
    }

}
