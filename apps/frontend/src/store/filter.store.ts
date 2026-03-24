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
    async getFilter(filterId: string) {
      this.current = await filterApi.show(filterId);
    },

    async getFilters() {
      this.index = await filterApi.index();
    },

    async postFilter(filter: Omit<Filter, 'id'>) {
      const newFilter = await filterApi.create(filter);
      this.index.push(newFilter);
      return newFilter;
    },

    async putFilter(filterId: string, filter: Omit<Filter, 'id'>) {
      const updatedFilter = await filterApi.update(filterId, filter);
      const index = this.index.findIndex(f => f.id === filterId);
      if (index !== -1) {
        this.index[index] = updatedFilter;
      }
      return updatedFilter;
    },

    async deleteFilter(filterId: string) {
      await filterApi.destroy(filterId);
      this.index = this.index.filter(f => f.id !== filterId);
    },
  },
});
