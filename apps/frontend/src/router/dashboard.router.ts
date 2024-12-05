import { RouteRecordRaw } from 'vue-router';
import DashboardShowPage from '@/page/dashboard/DashboardShowPage.vue';

export const dashboardRouter: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'dashboardShow',
    component: DashboardShowPage,
  },
];

