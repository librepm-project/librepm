import { RouteRecordRaw } from 'vue-router';
import AdminPriorityIndexPage from '@/page/admin/priority/AdminPriorityIndexPage.vue';
import AdminPriorityShowPage from '@/page/admin/priority/AdminPriorityShowPage.vue';
import AdminPriorityCreatePage from '@/page/admin/priority/AdminPriorityCreatePage.vue';
import { Permissions } from '@/lib/permissions';

export const adminPriorityRouter: RouteRecordRaw[] = [
  {
    path: '/admin/priority',
    name: 'adminPriorityIndex',
    component: AdminPriorityIndexPage,
    meta: { title: 'title.adminPriorityIndex', permission: Permissions.PriorityRead },
  },
  {
    path: '/admin/priority/create',
    name: 'adminPriorityCreate',
    component: AdminPriorityCreatePage,
    meta: { title: 'title.adminPriorityCreate', permission: Permissions.PriorityCreate },
  },
  {
    path: '/admin/priority/:priorityId',
    name: 'adminPriorityShow',
    component: AdminPriorityShowPage,
    meta: { title: 'title.adminPriorityShow', permission: Permissions.PriorityRead },
  },
];
