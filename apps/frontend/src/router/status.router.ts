import { RouteRecordRaw } from 'vue-router';
import StatusIndexPage from '@/page/status/StatusIndexPage.vue';
import StatusShowPage from '@/page/status/StatusShowPage.vue';
import StatusCreatePage from '@/page/status/StatusCreatePage.vue';

export const statusRouter: RouteRecordRaw[] = [
  {
    path: '/status',
    name: 'statusIndex',
    component: StatusIndexPage,
  },
  {
    path: '/status/{statusid}',
    name: 'statusShow',
    component: StatusShowPage,
  },
  {
    path: '/status/create',
    name: 'statusCreate',
    component: StatusCreatePage,
  },
];
