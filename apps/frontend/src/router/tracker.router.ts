import { RouteRecordRaw } from 'vue-router';
import TrackerIndexPage from '@/page/tracker/TrackerIndexPage.vue';
import TrackerShowPage from '@/page/tracker/TrackerShowPage.vue';
import { reactive } from 'vue';

export const trackerRouter: RouteRecordRaw[] = [
  {
    path: '/tracker',
    name: 'trackerIndex',
    component: TrackerIndexPage,
    meta: reactive({
      title: 'Trackers',
    }),
  },
  {
    path: '/tracker/{trackerid}',
    name: 'trackerShow',
    component: TrackerShowPage,
    meta: reactive({
      title: '',
    }),
  },
];
