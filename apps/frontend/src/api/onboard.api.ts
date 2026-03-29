import api from "@/api/api";

export interface OnboardRequest {
    siteTitle: string;
    email: string;
    password: string;
    firstName: string;
    lastName: string;
}

const create = async (data: OnboardRequest): Promise<void> => {
    await api.apiCall().post("/onboard", data);
}

export default { create }
