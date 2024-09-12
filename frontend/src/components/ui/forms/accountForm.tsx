import React, { FC, FormEvent, useEffect, useRef, useState } from "react";
import { CustomInputProps, FormProps, StrengthLevel } from "@/types";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faCircleExclamation, faEnvelope, faHeart, faLock, faLockOpen, faUser } from "@fortawesome/free-solid-svg-icons";
import Link from "next/link";
import { redirect } from "next/navigation";
import { SubmitButton } from "@/components/buttons/submitButton";
import { useAuth } from "@/components/providers/auth";
import validatePassword from "../../../../validators/validatePassword";
import { ACCEPTED_IMAGE_TYPES } from "@/lib/config";

export const CustomInput: FC<CustomInputProps> = ({ type, placeholder, inputRef, children, name, defaultValue, required = true, disabled = false }) => (
    <label className={`input input-bordered flex items-center gap-2 my-4`}>
        {children}
        <input type={type} name={name} className="grow" placeholder={placeholder} ref={inputRef} defaultValue={defaultValue} autoComplete="off" required={required} disabled={disabled} />
    </label>
);

const Prompt: FC<{ text: string, link: string }> = ({ text, link }) => (
    <div className="flex items-center justify-center">
        <Link href={`/${link}`} className="btn btn-outline btn-info rounded-full w-auto">
            <FontAwesomeIcon icon={faHeart} width={15} />
            <span>{text}</span>
        </Link>
    </div>
);

const AlertErr: FC<{ error: string }> = ({ error }) => (
    <div role="alert" className="alert alert-error">
        <FontAwesomeIcon icon={faCircleExclamation} width={15} height={15} className="animate-pulse" />
        <span>{error}</span>
    </div>
);

const AbstractForm: FC<FormProps> = ({ type }) => {
    const [isPasswordVisible, setIsPasswordVisible] = useState<boolean>(false);
    const [error, setError] = useState<string | null>(null);
    const [passwordStrength, setPasswordStrength] = useState<StrengthLevel>("weak");

    const rememberRef = useRef<HTMLInputElement | null>(null);
    const passRef = useRef<HTMLInputElement | null>(null); // for password validating 

    const { login, user, createAccount } = useAuth();

    if (user) {
        redirect("/profile");
    }

    const handleSubmit = (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        if (type === "login") {
            login(setError, rememberRef.current?.checked ?? false, e)
        } else if (type === "signup") {
            createAccount(e, setError)
        }
    };


    useEffect(() => {
        if (type === "signup") {
            const handlePasswordStrength = () => {
                validatePassword(setPasswordStrength, passRef);
            };

            const inputElement = passRef.current;
            if (inputElement) {
                inputElement.addEventListener("input", handlePasswordStrength);
            }

            return () => {
                if (inputElement) {
                    inputElement.removeEventListener("input", handlePasswordStrength);
                }
            };
        }
    }, [type]);

    return (
        <form className="max-sm:items-center max-sm:justify-center bg-base-300 md:p-10 p-5 rounded-3xl shadow-lg" onSubmit={handleSubmit} encType="multipart/form-data">
            {error && <AlertErr error={error} />}

            <CustomInput name="username" type="text" placeholder="Username...">
                <FontAwesomeIcon icon={faUser} width={12} height={12} />
            </CustomInput>

            {type === 'signup' && (
                <CustomInput name="email" type="email" placeholder="Email...">
                    <FontAwesomeIcon icon={faEnvelope} width={12} height={12} />
                </CustomInput>
            )}

            <CustomInput name="password" type={isPasswordVisible ? 'text' : 'password'} placeholder="Password..." inputRef={passRef}>
                <div className="tooltip md:tooltip-left tooltip-bottom" data-tip="Show password">
                    <button
                        type="button"
                        onClick={() => setIsPasswordVisible(!isPasswordVisible)}
                        className="animate-pulse"
                        tabIndex={-1}
                    >
                        <FontAwesomeIcon
                            icon={isPasswordVisible ? faLockOpen : faLock}
                            width={12}
                            height={12}
                        />
                    </button>
                </div>
            </CustomInput>

            {type === 'signup' && (
                <div className="py-5">
                    <input
                        type="file"
                        name="pic"
                        accept={ACCEPTED_IMAGE_TYPES}
                        className="file-input file-input-bordered file-input-secondary w-full max-w-xs"
                    />
                </div>
            )}

            {type === 'login' && (
                <div className="flex max-sm:flex-col items-center justify-center p-2 text-sm w-full">
                    <div className="flex-1">
                        <Link href="/" className="link">Forgot password?</Link>
                    </div>
                    <div className="form-control">
                        <label className="label cursor-pointer">
                            <span className="label-text">Remember me</span>
                            <input type="checkbox" className="toggle rounded-full ms-3" ref={rememberRef} />
                        </label>
                    </div>
                </div>
            )}

            <SubmitButton passwordStrength={passwordStrength} type={type} />

            <div className="divider py-4 text-gray-500">
                {type === 'login' ? "Don't have an account?" : 'Already have account?'}
            </div>

            <Prompt
                text={type === 'login' ? 'Create account' : 'Log in'}
                link={type === 'login' ? 'signup' : 'login'}
            />
        </form>
    );
};

export default AbstractForm;
