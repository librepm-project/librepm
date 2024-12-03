import { RouteRecordRaw } from 'vue-router';
import BoardShowPage from '@/page/board/BoardShowPage.vue';
import { reactive } from 'vue';

export const boardRouter: RouteRecordRaw[] = [
  {
    path: '/board',
    name: 'boardShow',
    component: BoardShowPage,
    meta: reactive({
      title: 'Boards',
    }),
  },
];
