import { RouteRecordRaw } from 'vue-router';
import BoardShowView from '../pages/board/BoardShowPage.vue';

export const boardRouter: RouteRecordRaw[] = [
  {
    path: '/board',
    name: 'boardShow',
    component: BoardShowView
  },
];
