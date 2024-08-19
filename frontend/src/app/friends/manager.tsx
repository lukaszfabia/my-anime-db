import { AxiosError, AxiosResponse } from "axios";
import api from "../../lib/api";
import { toast } from "react-toastify";
import { FriendRequest, RequestStatus } from "@/types/models";

export const respondToFriendRequest = async (friendRequest: FriendRequest | null, action: RequestStatus, callback?: (response?: AxiosResponse) => void) => {
    if (friendRequest) {
        try {
            const response = await api.post(`/auth/friend/${friendRequest.id}/respond/?status=${action}`);
            callback && callback(response);
        } catch (err) {
            const axiosError = err as AxiosError<GoResponse>;
            toast.error(axiosError.response?.data.error || "Something went wrong");
        }
    }
}

export const removeFriend = async (friendId: number, callback?: (response?: AxiosResponse) => void) => {
    try {
        const response = await api.delete(`/auth/friend/${friendId}`);
        callback && callback(response);
    } catch (err) {
        const axiosError = err as AxiosError<GoResponse>;
        toast.error(axiosError.response?.data.error || "Something went wrong");
    }
}
