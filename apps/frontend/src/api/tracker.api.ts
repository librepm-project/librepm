import { Tracker } from "../lib/interfaces/tracker.interface";
import api from "./api";

const index = async (): Promise<Tracker[]> => {
    const response = await api.apiCall().get("/tracker");
    return response.data as Tracker[];
}

const show = async (trackerId: number): Promise<Tracker> => {
    const response = await api.apiCall().get(`/tracker/${trackerId}`);
    return response.data as Tracker;
}

export default {
    index,
    show
}