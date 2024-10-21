import { RouteRecordRaw } from 'vue-router';
import BoardIndexView from '../views/board/BoardIndexView.vue';
import BoardShowView from '../views/board/BoardShowView.vue';
import BoardEditView from '../views/board/BoardEditView.vue';

export const boardRouter: RouteRecordRaw[] = [
  {
    path: '/board',
    name: 'boardIndex',
    component: BoardIndexView
  },
  {
    path: '/board/{boardId}',
    name: 'boardShow',
    component: BoardShowView
  },
  {
    path: '/board/{boardId}/edit',
    name: 'boardEdit',
    component: BoardEditView
  },
];
