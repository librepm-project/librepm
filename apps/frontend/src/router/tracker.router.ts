import { RouteRecordRaw } from 'vue-router';
import TrackerIndexPage from '@/page/tracker/TrackerIndexPage.vue';
import TrackerShowPage from '@/page/tracker/TrackerShowPage.vue';
import TrackerCreatePage from '@/page/tracker/TrackerCreatePage.vue';


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
  {
    path: '/tracker/create',
    name: 'trackerCreate',
    component: TrackerCreatePage,
  },
];
