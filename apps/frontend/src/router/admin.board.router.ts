import { RouteRecordRaw } from 'vue-router';
import AdminBoardIndexPage from '@/page/admin/board/AdminBoardIndexPage.vue';
import AdminBoardEditPage from '@/page/admin/board/AdminBoardEditPage.vue';
import BoardCreatePage from '@/page/board/BoardCreatePage.vue';

export const adminBoardRouter: RouteRecordRaw[] = [
  {
    path: '/admin/board',
    name: 'adminBoardIndex',
    component: AdminBoardIndexPage,
  },
  {
    path: '/admin/board/create',
    name: 'adminBoardCreate',
    component: BoardCreatePage,
  },
  {
    path: '/admin/board/:boardId/edit',
    name: 'adminBoardEdit',
    component: AdminBoardEditPage,
  },
];
