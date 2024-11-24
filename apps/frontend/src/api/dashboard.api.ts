import { Dashboard } from "../lib/interfaces/dashboard.interface";

const indexDashboard = async (): Promise<Dashdashboard[]> => {
    const response = await fetch("/api/dashboard");
    return await response.json() as Dashboard[];
}

export default {
    indexDashboard
}