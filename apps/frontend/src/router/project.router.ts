import { RouteRecordRaw } from 'vue-router';
import ProjectIndexView from '@/page/project/ProjectIndexPage.vue';
import ProjectShowView from '@/page/project/ProjectShowPage.vue';
import { reactive } from 'vue';

export const projectRouter: RouteRecordRaw[] = [
  {
    path: '/project',
    name: 'projectIndex',
    component: ProjectIndexView,
    meta: reactive({
      title: 'Projects',
    }),
  },
  {
    path: '/project/{projectid}',
    name: 'projectShow',
    component: ProjectShowView,
    meta: reactive({
      title: 'Project details',
    }),
  },
];
