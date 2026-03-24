import { defineStore } from 'pinia';
import { Board } from '@/lib/interfaces/board.interface';
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
    async getBoard(boardId: string) {
      this.current = await boardApi.show(boardId);
    },

    async getBoards() {
      this.index = await boardApi.index();
    },

    async create(board: Partial<Board>) {
      return boardApi.create(board);
    },

    async update(boardId: string, board: Partial<Board>) {
      await boardApi.update(boardId, board);
    },

    async remove(boardId: string) {
      await boardApi.remove(boardId);
    }
  },
});
