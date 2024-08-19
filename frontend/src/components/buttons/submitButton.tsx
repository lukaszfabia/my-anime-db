import { StrengthLevel } from "@/types";
import { faArrowRight } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { FC } from "react";

export const SubmitButton: FC<{ type: "signup" | "login", passwordStrength?: StrengthLevel }> = ({ type, passwordStrength }) => {

    const isSignup: boolean = passwordStrength !== undefined && type === "signup";

    const tooltipSyle: string = isSignup ? "tooltip tooltip-top" : "";

    const tooltipColor: string = isSignup ? passwordStrength === "weak" ? "tooltip-error" : passwordStrength === "good" ? "tooltip-warning" : "tooltip-success" : "";

    const dataTip = isSignup ? passwordStrength === "weak" ? "Password is too weak 🥺" : `Password is ${passwordStrength} 😎` : "";

    return (
        <div className={`flex items-center justify-center pt-5 ${tooltipSyle} ${tooltipColor}`} data-tip={dataTip}>
            <button className="btn btn-outline rounded-3xl w-auto" type="submit" id="submitButton">
                <FontAwesomeIcon icon={faArrowRight} width={15} height={15} />
                <span>{type === 'signup' ? 'Sign up' : 'Log in'}</span>
            </button>
        </div>
    )
}