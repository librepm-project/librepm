import { defineStore } from 'pinia';
import { Project } from '@/lib/interfaces/project.interface';
import projectApi from '@/api/project.api';


interface ProjectStore {
  current: Project | null;
  index: Project[];
}

export const useProjectStore = defineStore('project', {
  state: (): ProjectStore => {
    return {
      current: null,
      index: [],
    };
  },
  actions: {
    getProject(projectId: number) {
    },

    async getProjects() {
      this.index = await projectApi.index();
    },
  },
});
