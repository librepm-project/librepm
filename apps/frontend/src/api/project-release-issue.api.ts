import api from '@/api/api';
import { ProjectReleaseIssue } from '@/lib/interfaces/project-release-issue.interface';

const basePath = 'project-release-issue';

export default {
  getAll: async () => {
    const response = await api.apiCall().get(`/${basePath}`);
    return response.data as ProjectReleaseIssue[];
  },
  getById: async (id: string) => {
    const response = await api.apiCall().get(`/${basePath}/${id}`);
    return response.data as ProjectReleaseIssue;
  },
  post: async (entity: ProjectReleaseIssue) => {
    const response = await api.apiCall().post(`/${basePath}`, entity);
    return response.data as ProjectReleaseIssue;
  },
  delete: async (id: string) => {
    await api.apiCall().delete(`/${basePath}/${id}`);
  },
  getByProjectReleaseId: async (projectReleaseId: string) => {
    const response = await api.apiCall().get(`/project-release/${projectReleaseId}/issue`);
    return response.data as ProjectReleaseIssue[];
  },
  getByIssueId: async (issueId: string) => {
    const response = await api.apiCall().get(`/issue/${issueId}/project-release-issue`);
    return response.data as ProjectReleaseIssue | null;
  },
};
