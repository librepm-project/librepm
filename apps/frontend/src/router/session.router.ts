import { RouteRecordRaw } from 'vue-router';
import LoginPage from '@/page/session/LoginPage.vue';
import ProfilePage from '@/page/session/ProfilePage.vue';
import RegisterPage from '@/page/session/RegisterPage.vue';
import { useAppConfigStore } from '@/store/app-config.store';

export const sessionRouter: RouteRecordRaw[] = [
  {
    path: '/login',
    name: 'login',
    component: LoginPage,
    meta: { hideLayout: true },
  },
  {
    path: '/register',
    name: 'register',
    component: RegisterPage,
    meta: { hideLayout: true },
    beforeEnter: () => {
      const configStore = useAppConfigStore();
      if (!configStore.registerAllowed) {
        return { name: 'login' };
      }
    },
  },
  {
    path: '/profile',
    name: 'profile',
    component: ProfilePage,
  },
];
