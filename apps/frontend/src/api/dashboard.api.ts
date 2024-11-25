import { Dashboard } from "../lib/interfaces/dashboard.interface";

const index = async (): Promise<Dashboard[]> => {
    const response = await fetch("/api/dashboard");
    return await response.json() as Dashboard[];
}

export default {
    index
}