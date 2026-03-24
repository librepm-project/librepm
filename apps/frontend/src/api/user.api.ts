import { User } from "@/lib/interfaces/user.interface";
import api from "@/api/api";

const index = async (): Promise<User[]> => {
    const response = await api.apiCall().get("/user");
    return response.data as User[];
}

const show = async (userId: string): Promise<User> => {
    const response = await api.apiCall().get(`/user/${userId}`);
    return response.data as User;
}

const create = async (user: Omit<User, 'id'>): Promise<User> => {
    const response = await api.apiCall().post("/user", user);
    return response.data as User;
}

const update = async (userId: string, user: Omit<User, 'id'>): Promise<User> => {
    await api.apiCall().put(`/user/${userId}`, user);
    return show(userId);
}

const destroy = async (userId: string): Promise<void> => {
    await api.apiCall().delete(`/user/${userId}`);
}

export default {
    index,
    show,
    create,
    update,
    destroy
}
