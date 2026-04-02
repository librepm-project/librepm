import { RouteRecordRaw } from 'vue-router';
import AdminTrackerIndexPage from '@/page/admin/tracker/AdminTrackerIndexPage.vue';
import AdminTrackerShowPage from '@/page/admin/tracker/AdminTrackerShowPage.vue';
import AdminTrackerCreatePage from '@/page/admin/tracker/AdminTrackerCreatePage.vue';
import { Permissions } from '@/lib/permissions';

export const adminTrackerRouter: RouteRecordRaw[] = [
  {
    path: '/admin/tracker',
    name: 'adminTrackerIndex',
    component: AdminTrackerIndexPage,
    meta: { permission: Permissions.TrackerRead },
  },
  {
    path: '/admin/tracker/create',
    name: 'adminTrackerCreate',
    component: AdminTrackerCreatePage,
    meta: { permission: Permissions.TrackerCreate },
  },
  {
    path: '/admin/tracker/:trackerId',
    name: 'adminTrackerShow',
    component: AdminTrackerShowPage,
    meta: { permission: Permissions.TrackerRead },
  },
];
