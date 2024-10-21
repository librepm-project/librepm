import { RouteRecordRaw } from 'vue-router';
import FilterIndexView from '../views/filter/FilterIndexView.vue';
import FilterShowView from '../views/filter/FilterShowView.vue';

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
