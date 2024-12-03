import { User } from "@/lib/interfaces/user.interface";
import api from "@/api/api";

const create = async (): Promise<User> => {
    const response = await api.apiCall().get("/user/current");
    return await response.data as User;
}

export default {
    create
}