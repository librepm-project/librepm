import { RouteRecordRaw } from 'vue-router';
import FilterIndexPage from '@/page/filter/FilterIndexPage.vue';
import FilterViewPage from '@/page/filter/FilterViewPage.vue';
import FilterShowPage from '@/page/filter/FilterShowPage.vue';
import FilterCreatePage from '@/page/filter/FilterCreatePage.vue';

export const filterRouter: RouteRecordRaw[] = [
  {
    path: '/filter',
    name: 'filterIndex',
    component: FilterIndexPage,
  },
  {
    path: '/filter/create',
    name: 'filterCreate',
    component: FilterCreatePage,
  },
  {
    path: '/filter/:filterId',
    name: 'filterView',
    component: FilterViewPage,
  },
  {
    path: '/filter/:filterId/edit',
    name: 'filterEdit',
    component: FilterShowPage,
  },
];
