import api from "@/api/api";

export interface AppConfig {
    registerAllowed: boolean;
}

const show = async (): Promise<AppConfig> => {
    const response = await api.apiCall().get("/config");
    return response.data as AppConfig;
}

export default { show }
