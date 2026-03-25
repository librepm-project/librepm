import { RouteRecordRaw } from 'vue-router';
import DashboardShowPage from '@/page/dashboard/DashboardShowPage.vue';
import DashboardCreatePage from '@/page/dashboard/DashboardCreatePage.vue';

export const dashboardRouter: RouteRecordRaw[] = [
    {
        path: '/',
        redirect: '/dashboard',
    },
    {
        // Fontos: /dashboard/create a wildcard elé kerül
        path: '/dashboard/create',
        name: 'dashboardCreate',
        component: DashboardCreatePage,
    },
    {
        path: '/dashboard/:dashboardId?',
        name: 'dashboardShow',
        component: DashboardShowPage,
    },
];
