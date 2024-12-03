import { RouteRecordRaw } from 'vue-router';
import DashboardIndexPage from '@/page/dashboard/DashboardIndexPage.vue';
import { reactive } from 'vue';

export const dashboardRouter: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'dashboardIndex',
    component: DashboardIndexPage,
    meta: reactive({
      title: 'Dashboards',
    }),
  },
];

