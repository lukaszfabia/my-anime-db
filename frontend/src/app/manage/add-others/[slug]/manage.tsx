import { GoResponse } from "@/types/responses";
import { CreatorProps } from "./page";
import api from "@/lib/api";
import { toast } from "react-toastify";
import { convertTime } from "@/lib/computeTime";

export const deleteObj = (id: number, model: CreatorProps, innerFunc: () => void) => {
    api.delete<GoResponse>(`/auth/manage/${model.entity}/${id}`)
        .then((res) => {
            if (res.data.code === 200) {
                innerFunc();
            }
        })
        .catch((_: any) => toast.error("failed to delete studio"));
}

export const createObj = async (e: React.FormEvent<HTMLFormElement>, model: CreatorProps, id?: string | number | null) => {
    e.preventDefault();

    const form = new FormData(e.currentTarget);

    const updateDateField = (fieldName: string) => {
        const date = form.get(fieldName) as string;
        if (date) {
            form.set(fieldName, convertTime(date));
        }
    };

    updateDateField("establishedDate");

    const method = id ? api.put<GoResponse>(`/auth/manage/${model.entity}/${id}`, form) : api.post<GoResponse>(`/auth/manage/${model.entity}/`, form);

    return await method
        .then((res) => {
            if (res.data.code === 201) {
                return res.data;
            }
        })
        .catch(() => null);
}
