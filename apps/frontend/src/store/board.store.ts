import { defineStore } from 'pinia';
import { Board } from '@/lib/interfaces/board.interface';
import { boardMock } from '@/lib/mock-data';
import boardApi from '@/api/board.api';

interface BoardStore {
  current: Board | null;
  index: Board[];
}

export const useBoardStore = defineStore('board', {
  state: (): BoardStore => {
    return {
      current: null,
      index: [],
    };
  },
  actions: {
    async getBoard(boardId: number) {
      this.current = await boardApi.show(boardId);
    },

    async getBoards() {
      this.index = await boardApi.index();
    },
  },
});
