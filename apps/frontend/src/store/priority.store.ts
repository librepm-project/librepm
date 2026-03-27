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

    async getPriority(priorityId: string) {
      return this.getCurrentItem(priorityId);
    },

    async getPriorities() {
      return this.getAllItems();
    },

    async postPriority(priority: Omit<Priority, 'id'>) {
      return this.createItem(priority);
    },

    async putPriority(priorityId: string, priority: Omit<Priority, 'id'>) {
      return this.updateItem(priorityId, priority);
    },

    async deletePriority(priorityId: string) {
      return this.deleteItem(priorityId);
    },
  },
});
