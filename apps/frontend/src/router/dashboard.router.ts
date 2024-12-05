import { RouteRecordRaw } from 'vue-router';
import DashboardIndexPage from '@/page/dashboard/DashboardIndexPage.vue';

export const dashboardRouter: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'dashboardIndex',
    component: DashboardIndexPage,
  },
];

