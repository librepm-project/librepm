import { User } from "../lib/interfaces/user.interface";

const create = async (): Promise<User> => {
    const response = await fetch("/api/user/current");
    return await response.json() as User;
}

export default {
    create
}