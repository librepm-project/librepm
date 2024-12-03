import { defineStore } from 'pinia';
import { Filter } from '@/lib/interfaces/filter.interface';
import { filterMock } from '@/lib/mock-data';
import filterApi from '@/api/filter.api';

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
      this.current = filterMock[0];
    },

    async getFilters() {
      this.index = await filterApi.index();
    },
  },
});
