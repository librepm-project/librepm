import { defineStore } from 'pinia';
import { Board } from '../lib/interfaces/board.interface';
import { boardMock } from '../lib/mock-data';
import boardApi from '../api/board.api';

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
    getBoard(boardId: number) {
      this.current = boardMock[0];
    },

    async getBoards() {
      this.index = await boardApi.index();
    },
  },
});
