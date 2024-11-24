import { Filter } from "../lib/interfaces/filter.interface";

const indexFilter = async (): Promise<Filter[]> => {
    const response = await fetch("/api/filter");
    return await response.json() as Filter[];
}

export default {
    indexFilter
}