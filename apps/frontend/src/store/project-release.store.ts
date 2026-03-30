import { defineStore } from 'pinia';
import { ProjectRelease } from '@/lib/interfaces/project-release.interface';
import projectReleaseApi from '@/api/project-release.api';

interface ProjectReleaseStore {
  current: ProjectRelease | null;
  index: ProjectRelease[];
  byRelease: ProjectRelease[];
  byProject: ProjectRelease[];
}

export const useProjectReleaseStore = defineStore('projectRelease', {
  state: (): ProjectReleaseStore => {
    return {
      current: null,
      index: [],
      byRelease: [],
      byProject: [],
    };
  },
  actions: {
    async getAll() {
      this.index = await projectReleaseApi.getAll();
      return this.index;
    },

    async getById(id: string) {
      this.current = await projectReleaseApi.getById(id);
      return this.current;
    },

    async post(entity: ProjectRelease) {
      const data = await projectReleaseApi.post(entity);
      this.index.push(data);
      return data;
    },

    async delete(id: string) {
      await projectReleaseApi.delete(id);
      this.index = this.index.filter(item => item.id !== id);
    },

    async getByReleaseId(releaseId: string) {
      this.byRelease = await projectReleaseApi.getByReleaseId(releaseId);
      return this.byRelease;
    },

    async getByProjectId(projectId: string) {
      this.byProject = await projectReleaseApi.getByProjectId(projectId);
      return this.byProject;
    },
  },
});
