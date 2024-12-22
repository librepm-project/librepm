import { RouteRecordRaw } from 'vue-router';
import ProjectIndexPage from '@/page/project/ProjectIndexPage.vue';
import ProjectShowPage from '@/page/project/ProjectShowPage.vue';

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
];
