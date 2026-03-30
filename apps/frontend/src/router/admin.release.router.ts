import { RouteRecordRaw } from 'vue-router';
import AdminReleaseIndexPage from '@/page/admin/release/AdminReleaseIndexPage.vue';
import AdminReleaseShowPage from '@/page/admin/release/AdminReleaseShowPage.vue';
import AdminReleaseCreatePage from '@/page/admin/release/AdminReleaseCreatePage.vue';

export const adminReleaseRouter: RouteRecordRaw[] = [
  {
    path: '/admin/release',
    name: 'adminReleaseIndex',
    component: AdminReleaseIndexPage,
  },
  {
    path: '/admin/release/create',
    name: 'adminReleaseCreate',
    component: AdminReleaseCreatePage,
  },
  {
    path: '/admin/release/:releaseId',
    name: 'adminReleaseShow',
    component: AdminReleaseShowPage,
  },
];
