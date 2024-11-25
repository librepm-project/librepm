import { Filter } from "../lib/interfaces/filter.interface";

const index = async (): Promise<Filter[]> => {
    const response = await fetch("/api/filter");
    return await response.json() as Filter[];
}

export default {
    index
}