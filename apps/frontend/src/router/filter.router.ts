import { RouteRecordRaw } from 'vue-router';
import FilterIndexView from '../pages/filter/FilterIndexPage.vue';
import FilterShowView from '../pages/filter/FilterShowPage.vue';

export const filterRouter: RouteRecordRaw[] = [
  {
    path: '/filter',
    name: 'filterIndex',
    component: FilterIndexView,
  },
  {
    path: '/filter/{issueid}',
    name: 'filterShow',
    component: FilterShowView,
  },
];
