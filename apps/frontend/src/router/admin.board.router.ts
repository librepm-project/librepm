import { RouteRecordRaw } from 'vue-router';
import AdminBoardIndexPage from '@/page/admin/board/AdminBoardIndexPage.vue';
import AdminBoardEditPage from '@/page/admin/board/AdminBoardEditPage.vue';
import BoardCreatePage from '@/page/board/BoardCreatePage.vue';
import { Permissions } from '@/lib/permissions';

export const adminBoardRouter: RouteRecordRaw[] = [
  {
    path: '/admin/board',
    name: 'adminBoardIndex',
    component: AdminBoardIndexPage,
    meta: { permission: Permissions.BoardRead },
  },
  {
    path: '/admin/board/create',
    name: 'adminBoardCreate',
    component: BoardCreatePage,
    meta: { permission: Permissions.BoardCreate },
  },
  {
    path: '/admin/board/:boardId/edit',
    name: 'adminBoardEdit',
    component: AdminBoardEditPage,
    meta: { permission: Permissions.BoardUpdate },
  },
];
