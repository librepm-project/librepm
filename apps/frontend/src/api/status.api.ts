import { Status } from "@/lib/interfaces/status.interface";
import api from "@/api/api";

const index = async (): Promise<Status[]> => {
    const response = await api.apiCall().get("/status");
    return response.data as Status[];
}

const show = async (statusId: number): Promise<Status> => {
    const response = await api.apiCall().get(`/status/${statusId}`);
    return response.data as Status;
}

export default {
    index,
    show
}