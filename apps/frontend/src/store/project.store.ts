import { defineStore } from 'pinia';
import { Project } from '../lib/interfaces/project.interface';
import { projectMock } from '../lib/mock-data';
import projectAPI from '../api/project.api';


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
      this.current = projectMock[0];
    },

    async getProjects() {
      this.index = await projectAPI.indexProject()
    },
  },
});
