import { RouteRecordRaw } from 'vue-router';
import AdminStatusIndexPage from '@/page/admin/status/AdminStatusIndexPage.vue';
import AdminStatusShowPage from '@/page/admin/status/AdminStatusShowPage.vue';
import AdminStatusCreatePage from '@/page/admin/status/AdminStatusCreatePage.vue';

export const adminStatusRouter: RouteRecordRaw[] = [
  {
    path: '/admin/status',
    name: 'adminStatusIndex',
    component: AdminStatusIndexPage,
  },
  {
    path: '/admin/status/create',
    name: 'adminStatusCreate',
    component: AdminStatusCreatePage,
  },
  {
    path: '/admin/status/:statusId',
    name: 'adminStatusShow',
    component: AdminStatusShowPage,
  },
];
