import { RouteRecordRaw } from 'vue-router';
import BoardShowPage from '@/page/board/BoardShowPage.vue';
import BoardCreatePage from '@/page/board/BoardCreatePage.vue';


export const boardRouter: RouteRecordRaw[] = [
  {
    path: '/board',
    name: 'boardShow',
    component: BoardShowPage,
  },
  {
    path: '/board/create',
    name: 'boardCreate',
    component: BoardCreatePage,
  },
];
