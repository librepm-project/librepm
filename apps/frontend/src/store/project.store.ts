import { defineStore } from 'pinia';
import { Project } from '@/lib/interfaces/project.interface';
import projectApi from '@/api/project.api';
import projectIssuePropertyApi from '@/api/project-issue-property.api';
import { ProjectIssueProperty } from "@/lib/interfaces/project-issue-property.interface";
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

    async getIssueProperty(projectId: string): Promise<void> {
      this.currentIssueProperty = await projectIssuePropertyApi.index(projectId)
    }
  },
});
