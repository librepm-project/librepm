import { RouteRecordRaw } from 'vue-router';
import AdminBoardIndexPage from '@/page/admin/board/AdminBoardIndexPage.vue';
import AdminBoardEditPage from '@/page/admin/board/AdminBoardEditPage.vue';

export const adminBoardRouter: RouteRecordRaw[] = [
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

