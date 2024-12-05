import { RouteRecordRaw } from 'vue-router';
import FilterIndexPage from '@/page/filter/FilterIndexPage.vue';
import FilterShowPage from '@/page/filter/FilterShowPage.vue';

export const filterRouter: RouteRecordRaw[] = [
  {
    path: '/filter',
    name: 'filterIndex',
    component: FilterIndexPage,

  },
  {
    path: '/filter/{issueid}',
    name: 'filterShow',
    component: FilterShowPage,
  },
];
