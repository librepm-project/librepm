import { RouteRecordRaw } from 'vue-router';
import AdminProjectIndexPage from '@/page/admin/project/AdminProjectIndexPage.vue';
import AdminProjectShowPage from '@/page/admin/project/AdminProjectShowPage.vue';
import AdminProjectCreatePage from '@/page/admin/project/AdminProjectCreatePage.vue';
import { Permissions } from '@/lib/permissions';

export const adminProjectRouter: RouteRecordRaw[] = [
  {
    path: '/admin/project',
    name: 'adminProjectIndex',
    component: AdminProjectIndexPage,
    meta: { permission: Permissions.ProjectRead },
  },
  {
    path: '/admin/project/create',
    name: 'adminProjectCreate',
    component: AdminProjectCreatePage,
    meta: { permission: Permissions.ProjectCreate },
  },
  {
    path: '/admin/project/:projectId',
    name: 'adminProjectShow',
    component: AdminProjectShowPage,
    meta: { permission: Permissions.ProjectRead },
  },
];
