import { defineStore } from 'pinia';
import { Filter } from '@/lib/interfaces/filter.interface';
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
    },

    async getFilters() {
      this.index = await filterApi.index();
    },
  },
});
