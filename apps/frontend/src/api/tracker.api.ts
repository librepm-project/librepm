import { Tracker } from "@/lib/interfaces/tracker.interface";
import api from "@/api/api";

const index = async (): Promise<Tracker[]> => {
    const response = await api.apiCall().get("/tracker");
    return response.data as Tracker[];
}

const show = async (trackerId: string): Promise<Tracker> => {
    const response = await api.apiCall().get(`/tracker/${trackerId}`);
    return response.data as Tracker;
}

const create = async (tracker: Omit<Tracker, 'id'>): Promise<Tracker> => {
    const response = await api.apiCall().post("/tracker", tracker);
    return response.data as Tracker;
}

const update = async (trackerId: string, tracker: Omit<Tracker, 'id'>): Promise<Tracker> => {
    await api.apiCall().put(`/tracker/${trackerId}`, tracker);
    return show(trackerId);
}

const destroy = async (trackerId: string): Promise<void> => {
    await api.apiCall().delete(`/tracker/${trackerId}`);
}

export default {
    index,
    show,
    create,
    update,
    destroy
}