import { RouteRecordRaw } from 'vue-router';
import AdminTrackerIndexPage from '@/page/admin/tracker/AdminTrackerIndexPage.vue';
import AdminTrackerShowPage from '@/page/admin/tracker/AdminTrackerShowPage.vue';
import AdminTrackerCreatePage from '@/page/admin/tracker/AdminTrackerCreatePage.vue';

export const adminTrackerRouter: RouteRecordRaw[] = [
  {
    path: '/admin/tracker',
    name: 'adminTrackerIndex',
    component: AdminTrackerIndexPage,
  },
  {
    path: '/admin/tracker/create',
    name: 'adminTrackerCreate',
    component: AdminTrackerCreatePage,
  },
  {
    path: '/admin/tracker/:trackerId',
    name: 'adminTrackerShow',
    component: AdminTrackerShowPage,
  },
];
