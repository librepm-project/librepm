import { RouteRecordRaw } from 'vue-router';
import LoginView from '../pages/session/LoginPage.vue';
import ProfileView from '../pages/session/ProfilePage.vue';

export const sessionRouter: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'login',
    component: LoginView,
  },
  {
    path: '/profile',
    name: 'profile',
    component: ProfileView,
  },
];
