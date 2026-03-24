import { defineStore } from 'pinia';
import { Project } from '@/lib/interfaces/project.interface';
import projectApi from '@/api/project.api';
import projectIssuePropertyApi from '@/api/projectIssueProperty.api';
import { ProjectIssueProperty } from "@/lib/interfaces/projectIssueProperty.interface";
import { createCrudActions } from './utils/crudMixin';

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
    ...createCrudActions<Project>(projectApi),

    // Alias methods for backward compatibility
    async getProject(projectId: string) {
      return this.getCurrentItem(projectId);
    },

    async getProjects() {
      return this.getAllItems();
    },

    async postProject(project: Omit<Project, 'id'>) {
      return this.createItem(project);
    },

    async putProject(projectId: string, project: Omit<Project, 'id'>) {
      return this.updateItem(projectId, project);
    },

    async deleteProject(projectId: string) {
      return this.deleteItem(projectId);
    },

    async getIssueProperty(projectId: string): Promise<void> {
      this.currentIssueProperty = await projectIssuePropertyApi.index(projectId)
    }
  },
});
