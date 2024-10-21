import { RouteRecordRaw } from 'vue-router';
import DashboardIndexView from '../views/dashboard/DashboardIndexView.vue';

export const dashboardRouter: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'dashboardIndex',
    component: DashboardIndexView,
  },
];

