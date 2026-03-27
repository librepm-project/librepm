import { User } from "@/lib/interfaces/user.interface";
import api from "@/api/api";

const getCurrent = async (): Promise<User> => {
    const response = await api.apiCall().get("/user/current");
    return response.data as User;
}

const updateCurrent = async (user: Partial<User>): Promise<void> => {
    await api.apiCall().put("/user/current", user);
}

export default {
    getCurrent,
    updateCurrent,
}