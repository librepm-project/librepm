import { RouteRecordRaw } from 'vue-router';
import AdminUserIndexView from '@/pages/admin/user/AdminUserIndexPage.vue';
import AdminUserShowView from '@/pages/admin/user/AdminUserShowPage.vue';
import AdminBoardIndexView from '@/pages/admin/board/AdminBoardIndexPage.vue';
import AdminBoardEditView from '@/pages/admin/board/AdminBoardEditPage.vue';

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
  {
    path: '/admin/board',
    name: 'adminBoardIndex',
    component: AdminBoardIndexView
  },
  {
    path: '/admin/board/{boardId}/edit',
    name: 'adminBoardEdit',
    component: AdminBoardEditView
  },
];

