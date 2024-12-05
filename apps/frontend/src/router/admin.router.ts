import { RouteRecordRaw } from 'vue-router';
import AdminUserIndexPage from '@/page/admin/user/AdminUserIndexPage.vue';
import AdminUserShowPage from '@/page/admin/user/AdminUserShowPage.vue';
import AdminBoardIndexPage from '@/page/admin/board/AdminBoardIndexPage.vue';
import AdminBoardEditPage from '@/page/admin/board/AdminBoardEditPage.vue';

export const adminRouter: RouteRecordRaw[] = [
  {
    path: '/admin/user',
    name: 'adminUserIndex',
    component: AdminUserIndexPage,
  },
  {
    path: '/admin/user/{userId}',
    name: 'adminUserShow',
    component: AdminUserShowPage,
  },
  {
    path: '/admin/board',
    name: 'adminBoardIndex',
    component: AdminBoardIndexPage,
  },
  {
    path: '/admin/board/{boardId}/edit',
    name: 'adminBoardEdit',
    component: AdminBoardEditPage,
  },
];

