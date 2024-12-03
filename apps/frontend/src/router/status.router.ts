import { RouteRecordRaw } from 'vue-router';
import StatusIndexView from '@/pages/status/StatusIndexPage.vue';
import StatusShowView from '@/pages/status/StatusShowPage.vue';

export const statusRouter: RouteRecordRaw[] = [
  {
    path: '/status',
    name: 'statusIndex',
    component: StatusIndexView,
  },
  {
    path: '/status/{statusid}',
    name: 'statusShow',
    component: StatusShowView,
  },
];
