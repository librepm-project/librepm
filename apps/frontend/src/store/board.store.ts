import { defineStore } from 'pinia';
import { Board } from '../lib/interfaces/board.interface';

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
      this.current = {
        name: 'Board 1',
        description: 'Lorem ipsum dolor amet',
      };
    },

    getBoards() {
      this.index = [
        {
          name: 'Board 1',
          description: 'Lorem ipsum dolor amet',
        },
        {
          name: 'Board 2',
          description: 'Lorem ipsum dolor amet',
        },
      ];
    },
  },
});
