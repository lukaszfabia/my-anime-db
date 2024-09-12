import { FC } from "react";

interface SelectorProps<> {
    text: string;
    collection: string[];
    name: string;
    handler?: (e: React.ChangeEvent<HTMLSelectElement>, ...props: any) => void;
    lastElem?: string;
    disablePrompt?: boolean;
    disableValue?: string;
}

export const Selector: FC<SelectorProps> = ({ collection, text, name, handler, lastElem, disableValue, disablePrompt = true }) => {
    return (
        <div className="w-full max-w-xs py-3">
            <select name={name} className="select select-bordered w-full max-w-xs" value={lastElem ? lastElem : ""} onChange={handler ? (e, ...props) => handler(e, ...props) : () => { }}>
                <option value="" disabled={!disablePrompt}>{text}</option>
                {collection.map((item: string, i: number) => (
                    <option key={`${i}_${item}`} value={item} disabled={item === disableValue}>{item}</option>
                ))}
            </select>
        </div>
    );
};