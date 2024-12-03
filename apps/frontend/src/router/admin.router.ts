import { RouteRecordRaw } from 'vue-router';
import AdminUserIndexPage from '@/page/admin/user/AdminUserIndexPage.vue';
import AdminUserShowPage from '@/page/admin/user/AdminUserShowPage.vue';
import AdminBoardIndexPage from '@/page/admin/board/AdminBoardIndexPage.vue';
import AdminBoardEditPage from '@/page/admin/board/AdminBoardEditPage.vue';
import { reactive } from 'vue';

export const adminRouter: RouteRecordRaw[] = [
  {
    path: '/admin/user',
    name: 'adminUserIndex',
    component: AdminUserIndexPage,
    meta: reactive({
      title: 'Users',
    }),
  },
  {
    path: '/admin/user/{userId}',
    name: 'adminUserShow',
    component: AdminUserShowPage,
    meta: reactive({
      title: 'User details',
    }),
  },
  {
    path: '/admin/board',
    name: 'adminBoardIndex',
    component: AdminBoardIndexPage,
    meta: reactive({
      title: 'Boards',
    }),
  },
  {
    path: '/admin/board/{boardId}/edit',
    name: 'adminBoardEdit',
    component: AdminBoardEditPage,
    meta: reactive({
      title: 'Edit board',
    }),
  },
];

