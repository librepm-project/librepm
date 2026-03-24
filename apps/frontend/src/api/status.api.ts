import { Status } from "@/lib/interfaces/status.interface";
import api from "@/api/api";

const index = async (): Promise<Status[]> => {
    const response = await api.apiCall().get("/status");
    return response.data as Status[];
}

const show = async (statusId: string): Promise<Status> => {
    const response = await api.apiCall().get(`/status/${statusId}`);
    return response.data as Status;
}

const create = async (status: Omit<Status, 'id'>): Promise<Status> => {
    const response = await api.apiCall().post("/status", status);
    return response.data as Status;
}

const update = async (statusId: string, status: Omit<Status, 'id'>): Promise<Status> => {
    await api.apiCall().put(`/status/${statusId}`, status);
    return show(statusId);
}

const destroy = async (statusId: string): Promise<void> => {
    await api.apiCall().delete(`/status/${statusId}`);
}

export default {
    index,
    show,
    create,
    update,
    destroy
}