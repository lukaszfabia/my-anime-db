import { Model } from "@/types/models";
import { faTrash, faXmark } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { FC, ReactNode, useRef } from "react";

export interface DialogProps {
    title?: string;
    handler?: () => void;
    children?: ReactNode;
    icon?: ReactNode;
    model?: Model;
    modalRef?: React.RefObject<HTMLDialogElement>;
    errorColor?: boolean;
    wantButton?: boolean;
    short?: boolean;
    actionOrClose?: boolean;
}

export const DialogWindow: FC<DialogProps> = ({ title, handler, children, icon, modalRef = useRef<HTMLDialogElement>(null), errorColor = false, wantButton = false, short = false, actionOrClose = false }) => {

    return (
        <div className="preview">
            {wantButton && (
                <button onClick={() => modalRef.current?.showModal()} className={errorColor ? "text-red-500" : ""}>
                    {icon ? icon : <FontAwesomeIcon icon={faTrash} width={10} className="mr-2" />}
                    <span>{title}</span>
                </button>
            )}

            <dialog className="modal" ref={modalRef} id="modal">
                <div className={`modal-box w-5/6 ${short ? "max-w-xl" : "max-w-full"}`}>
                    <h3 className="font-bold text-lg">Attention!</h3>
                    <div className="py-4">
                        {children}
                    </div>
                    <div className="modal-action flex items-center justify-center">
                        <form method="dialog">
                            {actionOrClose ? (
                                <>
                                    <button className="btn btn-error mr-2" onClick={handler}><FontAwesomeIcon icon={faTrash} width={10} />Remove</button>
                                    <button className="btn btn-ghost ml-2" onClick={() => modalRef.current?.close()}><FontAwesomeIcon icon={faXmark} width={10} />Close</button>
                                </>
                            ) :
                                <button className="btn btn-sm btn-circle btn-ghost absolute right-2 top-2">✕</button>
                            }
                        </form>
                    </div>
                </div>
            </dialog>
        </div>
    );
}