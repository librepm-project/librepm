import { Issue } from "@/lib/interfaces/issue.interface";
import api from "@/api/api";

const index = async (params?: { filterId?: string }): Promise<Issue[]> => {
    const response = await api.apiCall().get("/issue", { params });
    return response.data as Issue[];
}

const show = async (issueId: string): Promise<Issue> => {
    const response = await api.apiCall().get(`/issue/${issueId}`);
    return response.data as Issue;
}

const create = async (issue: Partial<Issue>): Promise<Issue> => {
    const response = await api.apiCall().post(`/issue`, issue);
    return response.data as Issue;
}

const update = async (issueId: string, issue: Partial<Issue>): Promise<Issue> => {
    const response = await api.apiCall().put(`/issue/${issueId}`, issue);
    return response.data as Issue;
}

const destroy = async (issueId: string): Promise<void> => {
    await api.apiCall().delete(`/issue/${issueId}`);
}

export default {
    index,
    show,
    create,
    update,
    destroy
}