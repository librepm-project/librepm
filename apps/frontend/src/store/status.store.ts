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

    // Alias methods for backward compatibility
    async getStatus(statusId: string) {
      return this.getCurrentItem(statusId);
    },

    async getStatuses() {
      return this.getAllItems();
    },

    async postStatus(status: Omit<Status, 'id'>) {
      return this.createItem(status);
    },

    async putStatus(statusId: string, status: Omit<Status, 'id'>) {
      return this.updateItem(statusId, status);
    },

    async deleteStatus(statusId: string) {
      return this.deleteItem(statusId);
    },
  },
});
