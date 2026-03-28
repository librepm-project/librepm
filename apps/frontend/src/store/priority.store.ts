import { defineStore } from 'pinia';
import { Priority } from '@/lib/interfaces/priority.interface';
import priorityApi from '@/api/priority.api';
import { createCrudActions } from './utils/crudMixin';

interface PriorityStore {
  current: Priority | null;
  index: Priority[];
}

export const usePriorityStore = defineStore('priority', {
  state: (): PriorityStore => ({
    current: null,
    index: [],
  }),
  actions: {
    ...createCrudActions<Priority>(priorityApi),
  },
});
