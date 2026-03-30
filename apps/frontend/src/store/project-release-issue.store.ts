import { defineStore } from 'pinia';
import { ProjectReleaseIssue } from '@/lib/interfaces/project-release-issue.interface';
import projectReleaseIssueApi from '@/api/project-release-issue.api';

interface ProjectReleaseIssueStore {
  current: ProjectReleaseIssue | null;
  index: ProjectReleaseIssue[];
  byProjectRelease: ProjectReleaseIssue[];
  byIssue: ProjectReleaseIssue | null;
}

export const useProjectReleaseIssueStore = defineStore('projectReleaseIssue', {
  state: (): ProjectReleaseIssueStore => {
    return {
      current: null,
      index: [],
      byProjectRelease: [],
      byIssue: null,
    };
  },
  actions: {
    async getAll() {
      this.index = await projectReleaseIssueApi.getAll();
      return this.index;
    },

    async getById(id: string) {
      this.current = await projectReleaseIssueApi.getById(id);
      return this.current;
    },

    async post(entity: ProjectReleaseIssue) {
      const data = await projectReleaseIssueApi.post(entity);
      this.index.push(data);
      return data;
    },

    async delete(id: string) {
      await projectReleaseIssueApi.delete(id);
      this.index = this.index.filter(item => item.id !== id);
    },

    async getByProjectReleaseId(projectReleaseId: string) {
      this.byProjectRelease = await projectReleaseIssueApi.getByProjectReleaseId(projectReleaseId);
      return this.byProjectRelease;
    },

    async getByIssueId(issueId: string) {
      this.byIssue = await projectReleaseIssueApi.getByIssueId(issueId);
      return this.byIssue;
    },
  },
});
