import { RouteRecordRaw } from 'vue-router';
import LoginPage from '@/page/session/LoginPage.vue';
import ProfilePage from '@/page/session/ProfilePage.vue';
import { reactive } from 'vue';

export const sessionRouter: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'login',
    component: LoginPage,
    meta: reactive({
      title: '',
    }),
  },
  {
    path: '/profile',
    name: 'profile',
    component: ProfilePage,
    meta: reactive({
      title: '',
    }),
  },
];
