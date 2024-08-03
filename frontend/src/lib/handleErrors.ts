import axios from "axios";

export default function readError(error: any): string {
    return axios.isAxiosError(error) ? error.response?.data.error : "Something went wrong";
}