import { RouteRecordRaw } from 'vue-router';
import AdminStatusIndexPage from '@/page/admin/status/AdminStatusIndexPage.vue';
import AdminStatusShowPage from '@/page/admin/status/AdminStatusShowPage.vue';
import AdminStatusCreatePage from '@/page/admin/status/AdminStatusCreatePage.vue';
import { Permissions } from '@/lib/permissions';

export const adminStatusRouter: RouteRecordRaw[] = [
  {
    path: '/admin/status',
    name: 'adminStatusIndex',
    component: AdminStatusIndexPage,
    meta: { permission: Permissions.StatusRead },
  },
  {
    path: '/admin/status/create',
    name: 'adminStatusCreate',
    component: AdminStatusCreatePage,
    meta: { permission: Permissions.StatusCreate },
  },
  {
    path: '/admin/status/:statusId',
    name: 'adminStatusShow',
    component: AdminStatusShowPage,
    meta: { permission: Permissions.StatusRead },
  },
];
