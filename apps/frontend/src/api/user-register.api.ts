import { RegisterRequest, User } from "@/lib/interfaces/user.interface";
import api from "@/api/api";

const create = async (data: RegisterRequest): Promise<User> => {
    const response = await api.apiCall().post("/user/register", data);
    return response.data as User;
}

export default { create }
