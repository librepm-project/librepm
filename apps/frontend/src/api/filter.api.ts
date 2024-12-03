import { Filter } from "@/lib/interfaces/filter.interface";
import api from "@/api/api";

const index = async (): Promise<Filter[]> => {
    const response = await api.apiCall().get("/filter");
    return response.data as Filter[];
}

const show = async (filterId: number): Promise<Filter> => {
    const response = await api.apiCall().get(`/filter/${filterId}`);
    return response.data as Filter;
}

export default {
    index,
    show
}