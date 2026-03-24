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
    async getProject(projectId: string) {
      this.current = await projectApi.show(projectId);
    },

    async getProjects() {
      this.index = await projectApi.index();
    },

    async postProject(project: Omit<Project, 'id'>) {
      const newProject = await projectApi.create(project);
      this.index.push(newProject);
      return newProject;
    },

    async putProject(projectId: string, project: Omit<Project, 'id'>) {
      const updatedProject = await projectApi.update(projectId, project);
      const index = this.index.findIndex(p => p.id === projectId);
      if (index !== -1) {
        this.index[index] = updatedProject;
      }
      return updatedProject;
    },

    async deleteProject(projectId: string) {
      await projectApi.destroy(projectId);
      this.index = this.index.filter(p => p.id !== projectId);
    },

    async getIssueProperty(projectId: string): Promise<void> {
      this.currentIssueProperty = await projectIssuePropertyApi.index(projectId)
    }
  },
});
