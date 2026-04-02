import { RouteRecordRaw } from 'vue-router';
import AdminUserIndexPage from '@/page/admin/user/AdminUserIndexPage.vue';
import AdminUserShowPage from '@/page/admin/user/AdminUserShowPage.vue';
import AdminUserCreatePage from '@/page/admin/user/AdminUserCreatePage.vue';
import { Permissions } from '@/lib/permissions';

export const adminUserRouter: RouteRecordRaw[] = [
  {
    path: '/admin/user',
    name: 'adminUserIndex',
    component: AdminUserIndexPage,
    meta: { permission: Permissions.UserRead },
  },
  {
    path: '/admin/user/create',
    name: 'adminUserCreate',
    component: AdminUserCreatePage,
    meta: { permission: Permissions.UserCreate },
  },
  {
    path: '/admin/user/:userId',
    name: 'adminUserShow',
    component: AdminUserShowPage,
    meta: { permission: Permissions.UserRead },
  },
];
