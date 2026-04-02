import { RouteRecordRaw } from 'vue-router';
import BoardShowPage from '@/page/board/BoardShowPage.vue';
import BoardCreatePage from '@/page/board/BoardCreatePage.vue';
import { useBoardStore } from '@/store/board.store';
import { useSettingStore } from '@/store/setting.store';
import { Permissions } from '@/lib/permissions';

export const boardRouter: RouteRecordRaw[] = [
  {
    path: '/board',
    name: 'boardRedirect',
    beforeEnter: async (to, from, next) => {
      const boardStore = useBoardStore();
      const settingStore = useSettingStore();

      await Promise.all([
        boardStore.getAll(),
        settingStore.fetch()
      ]);

      const defaultBoardId = settingStore.getValue('default_board_id');
      if (defaultBoardId && boardStore.index.some(b => b.id === defaultBoardId)) {
        next({ name: 'boardShow', params: { boardId: defaultBoardId } });
      } else if (boardStore.index && boardStore.index.length > 0) {
        next({ name: 'boardShow', params: { boardId: boardStore.index[0].id } });
      } else {
        next({ name: 'boardCreate' });
      }
    }
  },
  {
    path: '/board/:boardId',
    name: 'boardShow',
    component: BoardShowPage,
    meta: { permission: Permissions.BoardRead },
  },
  {
    path: '/board/create',
    name: 'boardCreate',
    component: BoardCreatePage,
    meta: { permission: Permissions.BoardCreate },
  },
];
