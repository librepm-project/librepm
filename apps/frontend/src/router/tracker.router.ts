import { RouteRecordRaw } from 'vue-router';
import TrackerIndexPage from '@/page/tracker/TrackerIndexPage.vue';
import TrackerShowPage from '@/page/tracker/TrackerShowPage.vue';

export const trackerRouter: RouteRecordRaw[] = [
  {
    path: '/tracker',
    name: 'trackerIndex',
    component: TrackerIndexPage,
  },
  {
    path: '/tracker/{trackerid}',
    name: 'trackerShow',
    component: TrackerShowPage,
  },
];
