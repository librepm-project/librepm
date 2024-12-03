import { Dashboard } from "@/lib/interfaces/dashboard.interface";
import api from "@/api/api";

const index = async (): Promise<Dashboard[]> => {
    const response = await api.apiCall().get("/dashboard");
    return response.data as Dashboard[];
}

const show = async (dashboardId: number): Promise<Dashboard> => {
    const response = await api.apiCall().get(`/dashboard/${dashboardId}`);
    return response.data as Dashboard;
}

export default {
    index,
    show
}