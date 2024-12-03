import { RouteRecordRaw } from 'vue-router';
import DashboardIndexView from '@/pages/dashboard/DashboardIndexPage.vue';

export const dashboardRouter: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'dashboardIndex',
    component: DashboardIndexView,
  },
];

