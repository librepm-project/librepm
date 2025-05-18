import { defineStore } from 'pinia';
import { Project } from '@/lib/interfaces/project.interface';
import projectApi from '@/api/project.api';
import projectIssuePropertyApi from '@/api/projectIssueProperty.api';
import { ProjectIssueProperty } from "@/lib/interfaces/projectIssueProperty.interface";

interface ProjectStore {
  current: Project | null;
  index: Project[];
  currentIssueProperty: ProjectIssueProperty | null;
}

export const useProjectStore = defineStore('project', {
  state: (): ProjectStore => {
    return {
      current: null,
      index: [],
      currentIssueProperty: null,
    };
  },
  actions: {
    getProject(projectId: string) {
    },

    async getProjects() {
      this.index = await projectApi.index();
    },

    async getIssueProperty(projectId: string): Promise<void> {
      this.currentIssueProperty = await projectIssuePropertyApi.index(projectId)
    }
  },
});
