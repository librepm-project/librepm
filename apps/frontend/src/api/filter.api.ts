import { Filter } from "@/lib/interfaces/filter.interface";
import api from "@/api/api";

const index = async (): Promise<Filter[]> => {
    const response = await api.apiCall().get("/filter");
    return response.data as Filter[];
}

const show = async (filterId: string): Promise<Filter> => {
    const response = await api.apiCall().get(`/filter/${filterId}`);
    return response.data as Filter;
}

const create = async (filter: Omit<Filter, 'id'>): Promise<Filter> => {
    const response = await api.apiCall().post("/filter", filter);
    return response.data as Filter;
}

const update = async (filterId: string, filter: Omit<Filter, 'id'>): Promise<Filter> => {
    await api.apiCall().put(`/filter/${filterId}`, filter);
    return show(filterId);
}

const destroy = async (filterId: string): Promise<void> => {
    await api.apiCall().delete(`/filter/${filterId}`);
}

export default {
    index,
    show,
    create,
    update,
    destroy
}
