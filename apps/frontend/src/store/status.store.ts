import { defineStore } from 'pinia';
import { Status } from '@/lib/interfaces/status.interface';
import statusApi from '@/api/status.api';
import { createCrudActions } from './utils/crudMixin';

interface StatusStore {
  current: Status | null;
  index: Status[];
}

export const useStatusStore = defineStore('status', {
  state: (): StatusStore => {
    return {
      current: null,
      index: [],
    };
  },
  actions: {
    ...createCrudActions<Status>(statusApi),
  },
});
