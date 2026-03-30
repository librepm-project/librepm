import api from '@/api/api';
import { ProjectRelease } from '@/lib/interfaces/project-release.interface';

const basePath = 'project-release';

export default {
  getAll: async () => {
    const response = await api.apiCall().get(`/${basePath}`);
    return response.data as ProjectRelease[];
  },
  getById: async (id: string) => {
    const response = await api.apiCall().get(`/${basePath}/${id}`);
    return response.data as ProjectRelease;
  },
  post: async (entity: ProjectRelease) => {
    const response = await api.apiCall().post(`/${basePath}`, entity);
    return response.data as ProjectRelease;
  },
  delete: async (id: string) => {
    await api.apiCall().delete(`/${basePath}/${id}`);
  },
  getByReleaseId: async (releaseId: string) => {
    const response = await api.apiCall().get(`/release/${releaseId}/project`);
    return response.data as ProjectRelease[];
  },
  getByProjectId: async (projectId: string) => {
    const response = await api.apiCall().get(`/project/${projectId}/release`);
    return response.data as ProjectRelease[];
  },
};
