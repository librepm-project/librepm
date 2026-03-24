import { Project } from "@/lib/interfaces/project.interface";
import api from "@/api/api";

const index = async (): Promise<Project[]> => {
    const response = await api.apiCall().get("/project");
    return response.data as Project[];
}

const show = async (projectId: string): Promise<Project> => {
    const response = await api.apiCall().get(`/project/${projectId}`);
    return response.data as Project;
}

const create = async (project: Omit<Project, 'id'>): Promise<Project> => {
    const response = await api.apiCall().post("/project", project);
    return response.data as Project;
}

const update = async (projectId: string, project: Omit<Project, 'id'>): Promise<Project> => {
    await api.apiCall().put(`/project/${projectId}`, project);
    return show(projectId);
}

const destroy = async (projectId: string): Promise<void> => {
    await api.apiCall().delete(`/project/${projectId}`);
}

export default {
    index,
    show,
    create,
    update,
    destroy
}
