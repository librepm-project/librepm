import { RouteRecordRaw } from 'vue-router';
import AdminPriorityIndexPage from '@/page/admin/priority/AdminPriorityIndexPage.vue';
import AdminPriorityShowPage from '@/page/admin/priority/AdminPriorityShowPage.vue';
import AdminPriorityCreatePage from '@/page/admin/priority/AdminPriorityCreatePage.vue';

export const adminPriorityRouter: RouteRecordRaw[] = [
  {
    path: '/admin/priority',
    name: 'adminPriorityIndex',
    component: AdminPriorityIndexPage,
    meta: { title: 'title.adminPriorityIndex' },
  },
  {
    path: '/admin/priority/create',
    name: 'adminPriorityCreate',
    component: AdminPriorityCreatePage,
    meta: { title: 'title.adminPriorityCreate' },
  },
  {
    path: '/admin/priority/:priorityId',
    name: 'adminPriorityShow',
    component: AdminPriorityShowPage,
    meta: { title: 'title.adminPriorityShow' },
  },
];
