import { Project } from "../lib/interfaces/project.interface";
import api from "./api";

const index = async (): Promise<Project[]> => {
    const response = await api.apiCall().get("/project");
    return response.data as Project[];
}

const show = async (projectId: number): Promise<Project> => {
    const response = await api.apiCall().get(`/project/${projectId}`);
    return response.data as Project;
}

export default {
    index,
    show
}