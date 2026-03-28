import { ProjectIssueProperty } from "@/lib/interfaces/project-issue-property.interface";
import api from "@/api/api";

const index = async (projectId: string): Promise<ProjectIssueProperty> => {
    const response = await api.apiCall().get(`/project/${projectId}/issue-property`);
    return response.data as ProjectIssueProperty;
}

export default {
    index,
}