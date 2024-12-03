import { RouteRecordRaw } from 'vue-router';
import FilterIndexPage from '@/page/filter/FilterIndexPage.vue';
import FilterShowPage from '@/page/filter/FilterShowPage.vue';
import { reactive } from 'vue';

export const filterRouter: RouteRecordRaw[] = [
  {
    path: '/filter',
    name: 'filterIndex',
    component: FilterIndexPage,
    meta: reactive({
      title: 'Filters',
    }),

  },
  {
    path: '/filter/{issueid}',
    name: 'filterShow',
    component: FilterShowPage,
    meta: reactive({
      title: 'Filter details',
    }),
  },
];
