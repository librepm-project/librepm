import { Issue } from "@/lib/interfaces/issue.interface";
import api from "@/api/api";

const index = async (): Promise<Issue[]> => {
    const response = await api.apiCall().get("/issue");
    return response.data as Issue[];
}

const show = async (issueId: string): Promise<Issue> => {
    const response = await api.apiCall().get(`/issue/${issueId}`);
    return response.data as Issue;
}

export default {
    index,
    show
}