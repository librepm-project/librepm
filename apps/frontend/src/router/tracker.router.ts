import { RouteRecordRaw } from 'vue-router';
import TrackerIndexView from '@/pages/tracker/TrackerIndexPage.vue';
import TrackerShowView from '@/pages/tracker/TrackerShowPage.vue';

export const trackerRouter: RouteRecordRaw[] = [
  {
    path: '/tracker',
    name: 'trackerIndex',
    component: TrackerIndexView,
  },
  {
    path: '/tracker/{trackerid}',
    name: 'trackerShow',
    component: TrackerShowView,
  },
];
