import { RouteRecordRaw } from 'vue-router';
import DashboardShowPage from '@/page/dashboard/DashboardShowPage.vue';
import DashboardCreatePage from '@/page/dashboard/DashboardCreatePage.vue';
import { Permissions } from '@/lib/permissions';

export const dashboardRouter: RouteRecordRaw[] = [
    {
        path: '/',
        redirect: '/dashboard',
    },
    {
        path: '/dashboard/create',
        name: 'dashboardCreate',
        component: DashboardCreatePage,
        meta: { permission: Permissions.DashboardCreate },
    },
    {
        path: '/dashboard/:dashboardId?',
        name: 'dashboardShow',
        component: DashboardShowPage,
        meta: { permission: Permissions.DashboardRead },
    },
];
