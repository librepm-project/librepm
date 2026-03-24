import { defineStore } from 'pinia';
import { Filter } from '@/lib/interfaces/filter.interface';
import filterApi from '@/api/filter.api';
import { createCrudActions } from './utils/crudMixin';

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
    ...createCrudActions<Filter>(filterApi),

    // Alias methods for backward compatibility
    async getFilter(filterId: string) {
      return this.getCurrentItem(filterId);
    },

    async getFilters() {
      return this.getAllItems();
    },

    async postFilter(filter: Omit<Filter, 'id'>) {
      return this.createItem(filter);
    },

    async putFilter(filterId: string, filter: Omit<Filter, 'id'>) {
      return this.updateItem(filterId, filter);
    },

    async deleteFilter(filterId: string) {
      return this.deleteItem(filterId);
    },
  },
});
