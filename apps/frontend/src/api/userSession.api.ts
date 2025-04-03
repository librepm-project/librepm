import { Login, LoginRequest } from "@/lib/interfaces/user.interface";
import api from "@/api/api";

const create = async (data: LoginRequest): Promise<Login> => {
    const response = await api.apiCall().post("/user/session", data);
    return await response.data as Login;
}

export default {
    create
}