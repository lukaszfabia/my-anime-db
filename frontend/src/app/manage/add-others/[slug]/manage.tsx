import { GoResponse } from "@/types/responses";
import { CreatorProps } from "./page";
import api from "@/lib/api";
import { toast } from "react-toastify";
import { format } from "date-fns";

export const deleteObj = (id: number, model: CreatorProps, innerFunc: () => void) => {
    api.delete<GoResponse>(`/auth/manage/${model.entity}/${id}`)
        .then((res) => {
            if (res.data.code === 200) {
                innerFunc();
            }
        })
        .catch((_: any) => toast.error("failed to delete studio"));
}

export const createObj = async (e: React.FormEvent<HTMLFormElement>, model: CreatorProps) => {
    e.preventDefault();

    const form = new FormData(e.currentTarget);

    const updateDateField = (fieldName: string) => {
        const date = form.get(fieldName) as string;
        if (date) {
            form.set(fieldName, formatDate(date));
        }
    };

    updateDateField("birthdate");
    updateDateField("establishedDate");

    return await api.post<GoResponse>(`/auth/manage/${model.entity}/`, form)
        .then((res) => {
            if (res.data.code === 201) {
                return res.data;
            }
        })
        .catch(() => {
            toast.error("Failed to create character");
            return null;
        });
}


function formatDate(date: string) {
    return format(new Date(date), "yyyy-MM-dd");
}