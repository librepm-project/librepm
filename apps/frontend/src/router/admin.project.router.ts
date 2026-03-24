import { RouteRecordRaw } from 'vue-router';
import AdminProjectIndexPage from '@/page/admin/project/AdminProjectIndexPage.vue';
import AdminProjectShowPage from '@/page/admin/project/AdminProjectShowPage.vue';
import AdminProjectCreatePage from '@/page/admin/project/AdminProjectCreatePage.vue';

export const adminProjectRouter: RouteRecordRaw[] = [
  {
    path: '/admin/project',
    name: 'adminProjectIndex',
    component: AdminProjectIndexPage,
  },
  {
    path: '/admin/project/create',
    name: 'adminProjectCreate',
    component: AdminProjectCreatePage,
  },
  {
    path: '/admin/project/:projectid',
    name: 'adminProjectShow',
    component: AdminProjectShowPage,
  },
];
