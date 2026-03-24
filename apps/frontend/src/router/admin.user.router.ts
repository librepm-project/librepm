import { RouteRecordRaw } from 'vue-router';
import AdminUserIndexPage from '@/page/admin/user/AdminUserIndexPage.vue';
import AdminUserShowPage from '@/page/admin/user/AdminUserShowPage.vue';
import AdminUserCreatePage from '@/page/admin/user/AdminUserCreatePage.vue';

export const adminUserRouter: RouteRecordRaw[] = [
  {
    path: '/admin/user',
    name: 'adminUserIndex',
    component: AdminUserIndexPage,
  },
  {
    path: '/admin/user/create',
    name: 'adminUserCreate',
    component: AdminUserCreatePage,
  },
  {
    path: '/admin/user/:userId',
    name: 'adminUserShow',
    component: AdminUserShowPage,
  },
];

