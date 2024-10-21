import { RouteRecordRaw } from 'vue-router';
import LoginView from '../views/session/LoginView.vue';
import ProfileView from '../views/session/ProfileView.vue';

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
