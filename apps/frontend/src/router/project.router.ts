import { RouteRecordRaw } from 'vue-router';
import ProjectIndexPage from '@/page/project/ProjectIndexPage.vue';
import ProjectShowPage from '@/page/project/ProjectShowPage.vue';
import ProjectCreatePage from '@/page/project/ProjectCreatePage.vue';

export const projectRouter: RouteRecordRaw[] = [
  {
    path: '/project',
    name: 'projectIndex',
    component: ProjectIndexPage,
  },
  {
    path: '/project/{projectid}',
    name: 'projectShow',
    component: ProjectShowPage,
  },

  {
    path: '/project/create',
    name: 'projectCreate',
    component: ProjectCreatePage,
  },
];
