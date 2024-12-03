import { RouteRecordRaw } from 'vue-router';
import StatusIndexPage from '@/page/status/StatusIndexPage.vue';
import StatusShowPage from '@/page/status/StatusShowPage.vue';
import { reactive } from 'vue';

export const statusRouter: RouteRecordRaw[] = [
  {
    path: '/status',
    name: 'statusIndex',
    component: StatusIndexPage,
    meta: reactive({
      title: 'Statuses',
    }),
  },
  {
    path: '/status/{statusid}',
    name: 'statusShow',
    component: StatusShowPage,
    meta: reactive({
      title: '',
    }),
  },
];
