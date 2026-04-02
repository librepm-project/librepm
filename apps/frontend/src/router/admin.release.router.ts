import { RouteRecordRaw } from 'vue-router';
import AdminReleaseIndexPage from '@/page/admin/release/AdminReleaseIndexPage.vue';
import AdminReleaseShowPage from '@/page/admin/release/AdminReleaseShowPage.vue';
import AdminReleaseCreatePage from '@/page/admin/release/AdminReleaseCreatePage.vue';
import { Permissions } from '@/lib/permissions';

export const adminReleaseRouter: RouteRecordRaw[] = [
  {
    path: '/admin/release',
    name: 'adminReleaseIndex',
    component: AdminReleaseIndexPage,
    meta: { permission: Permissions.ReleaseRead },
  },
  {
    path: '/admin/release/create',
    name: 'adminReleaseCreate',
    component: AdminReleaseCreatePage,
    meta: { permission: Permissions.ReleaseCreate },
  },
  {
    path: '/admin/release/:releaseId',
    name: 'adminReleaseShow',
    component: AdminReleaseShowPage,
    meta: { permission: Permissions.ReleaseRead },
  },
];
