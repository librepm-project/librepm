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
  },
});
