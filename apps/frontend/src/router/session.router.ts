import { RouteRecordRaw } from 'vue-router';
import LoginPage from '@/page/session/LoginPage.vue';
import ProfilePage from '@/page/session/ProfilePage.vue';

export const sessionRouter: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'login',
    component: LoginPage,
    meta: { hideLayout: true },
  },
  {
    path: '/profile',
    name: 'profile',
    component: ProfilePage,
  },
];
