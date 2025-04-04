import { RouteRecordRaw } from 'vue-router';
import DashboardShowPage from '@/page/dashboard/DashboardShowPage.vue';
import DashboardCreatePage from '@/page/dashboard/DashboardCreatePage.vue';

export const dashboardRouter: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'dashboardShow',
    component: DashboardShowPage,
  },
  {
    path: '/dashboard/create',
    name: 'dashboardCreate',
    component: DashboardCreatePage,
  },
];

