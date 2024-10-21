import { RouteRecordRaw } from 'vue-router';
import AdminUserIndexView from '../views/admin/user/AdminUserIndexView.vue';
import AdminUserShowView from '../views/admin/user/AdminUserShowView.vue';

export const adminRouter: RouteRecordRaw[] = [
  {
    path: '/admin/user',
    name: 'adminUserIndex',
    component: AdminUserIndexView,
  },
  {
    path: '/admin/user/{userId}',
    name: 'adminUserShow',
    component: AdminUserShowView,
  },
];

