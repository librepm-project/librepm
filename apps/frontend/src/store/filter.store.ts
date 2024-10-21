import { defineStore } from 'pinia';
import { Filter } from '../lib/interfaces/filter.interface';

interface FilterStore {
  current: Filter | null;
  index: Filter[];
}

export const useFilterStore = defineStore('filter', {
  state: (): FilterStore => {
    return {
      current: null,
      index: [],
    };
  },
  actions: {
    getFilter(filterId: number) {
      this.current = {
        name: 'Filter 1',
        conditions: [],
        description: 'Lorem ipsum dolor amet',
      };
    },

    getFilters() {
      this.index = [
        {
          name: 'Filter 1',
          conditions: [],
          description: 'Lorem ipsum dolor amet',
        },
        {
          name: 'Filter 2',
          conditions: [],
          description: 'Lorem ipsum dolor amet',
        },
      ];
    },
  },
});
