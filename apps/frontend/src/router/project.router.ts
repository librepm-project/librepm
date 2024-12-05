import { RouteRecordRaw } from 'vue-router';
import ProjectIndexView from '@/page/project/ProjectIndexPage.vue';
import ProjectShowView from '@/page/project/ProjectShowPage.vue';

export const projectRouter: RouteRecordRaw[] = [
  {
    path: '/project',
    name: 'projectIndex',
    component: ProjectIndexView,
  },
  {
    path: '/project/{projectid}',
    name: 'projectShow',
    component: ProjectShowView,
  },
];
