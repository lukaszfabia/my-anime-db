import { faCircleExclamation, faEnvelope, faHeart, faLock, faLockOpen, faUser } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import Link from "next/link";
import { redirect } from "next/navigation";
import { FC, FormEvent, useEffect, useRef, useState } from "react";
import { useAuth } from "./providers/auth";
import { toast } from "react-toastify";
import { CustomInputProps, FormProps, StrengthLevel } from "@/types";
import validatePassword from "../../validators/validatePassword";
import { SubmitButton } from "./buttons/submitButton";


const CustomInput: FC<CustomInputProps> = ({ type, placeholder, inputRef, children, isPassword }) => {
    return (
        <label className="input input-bordered flex items-center gap-2 my-5" id={isPassword ? "passwordLabel" : ""}>
            {children}
            <input type={type} className="grow" placeholder={placeholder} ref={inputRef} autoComplete="off" required />
        </label>
    )
}


const Prompt: FC<{ text: string, link: string }> = ({ text, link }) => {
    return (
        <div className="flex items-center justify-center">
            <Link href={`/${link}`} className="btn btn-outline btn-info rounded-full w-3/4 lg:w-1/3">
                <FontAwesomeIcon icon={faHeart} width={15} />
                <span>{text}</span>
            </Link>
        </div>
    )
}

const AlertErr: FC<{ error: string }> = ({ error }) => {
    return (
        <div role="alert" className="alert alert-error">
            <FontAwesomeIcon icon={faCircleExclamation} width={15} height={15} className="animate-pulse" />
            <span>{error}</span>
        </div>
    )
}

const AbstractForm: FC<FormProps> = ({ type }) => {
    const [isPasswordVisible, setIsPasswordVisible] = useState<boolean>(false);
    const [error, setError] = useState<string | null>(null);

    const [passwordStrength, setPasswordStrength] = useState<StrengthLevel>("weak");

    const rememberRef = useRef<HTMLInputElement | null>(null);
    const passRef = useRef<HTMLInputElement | null>(null);
    const usernameRef = useRef<HTMLInputElement | null>(null);
    const emailRef = useRef<HTMLInputElement | null>(null);

    const fileRef = useRef<HTMLInputElement | null>(null);

    const { login, user, signup } = useAuth();

    if (user) {
        redirect("/profile");
    }

    const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();

        const username: string = usernameRef.current?.value!;
        const password: string = passRef.current?.value!;
        const email: string = emailRef.current?.value!;
        const pic: File = fileRef.current?.files![0]!;

        const remember: boolean = rememberRef.current?.checked!;

        if (type === "login") {
            await login({
                username: username,
                password: password,
                remember: remember,
            }).then((err: string | null) => {
                if (!err) {
                    toast.success('Successfully logged in 👌');
                } else {
                    setError(err);
                    toast.error('Something went wrong 🤯');
                }
            });

        } else if (type === "signup") {
            await signup({
                username: username,
                password: password,
                email: email,
                picUrl: pic,
                remember: false,
            }).then((err: string | null) => {
                if (!err) {
                    toast.success("Account created 🔥");
                } else {
                    setError(err);
                    toast.error("Couldn't create account 😿");
                }
            });


        }

    }

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
    }, []);

    return (
        <form className="max-sm:items-center max-sm:justify-center bg-base-300 p-10 rounded-3xl shadow-lg" onSubmit={handleSubmit} encType="multipart/form-data">
            {error && <AlertErr error={error} />}

            <CustomInput type="text" placeholder="Username..." inputRef={usernameRef}>
                <FontAwesomeIcon icon={faUser} width={12} height={12} />
            </CustomInput>

            {type === 'signup' && (
                <CustomInput type="email" placeholder="Email..." inputRef={emailRef}>
                    <FontAwesomeIcon icon={faEnvelope} width={12} height={12} />
                </CustomInput>
            )}

            <CustomInput type={isPasswordVisible ? 'text' : 'password'} placeholder="Password..." inputRef={passRef} isPassword={type == "signup"}>
                <div className="tooltip tooltip-left" data-tip="Show password">
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
                        id="pic"
                        ref={fileRef}
                        accept=".jpg,.jpeg,.png,.webp"
                        className="file-input file-input-bordered file-input-secondary w-full max-w-xs"
                    />
                </div>
            )}

            {type === 'login' && (
                <div className="flex max-md:flex-col items-center justify-center p-2 text-sm w-full">
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
        </form >
    );
};

export default AbstractForm;