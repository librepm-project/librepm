import { Login, LoginRequest } from "../lib/interfaces/user.interface";

const create = async (data: LoginRequest): Promise<Login> => {
    const response = await fetch("/api/user/session", {
        method: "post",
        body: JSON.stringify(data),
    });
    return await response.json() as Login;
}

export default {
    create
}