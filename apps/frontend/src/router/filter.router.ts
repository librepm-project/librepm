import { RouteRecordRaw } from 'vue-router';
import FilterIndexPage from '@/page/filter/FilterIndexPage.vue';
import FilterViewPage from '@/page/filter/FilterViewPage.vue';
import FilterShowPage from '@/page/filter/FilterShowPage.vue';
import FilterCreatePage from '@/page/filter/FilterCreatePage.vue';
import { Permissions } from '@/lib/permissions';

export const filterRouter: RouteRecordRaw[] = [
  {
    path: '/filter',
    name: 'filterIndex',
    component: FilterIndexPage,
    meta: { permission: Permissions.FilterRead },
  },
  {
    path: '/filter/create',
    name: 'filterCreate',
    component: FilterCreatePage,
    meta: { permission: Permissions.FilterCreate },
  },
  {
    path: '/filter/:filterId',
    name: 'filterView',
    component: FilterViewPage,
    meta: { permission: Permissions.FilterRead },
  },
  {
    path: '/filter/:filterId/edit',
    name: 'filterEdit',
    component: FilterShowPage,
    meta: { permission: Permissions.FilterUpdate },
  },
];
