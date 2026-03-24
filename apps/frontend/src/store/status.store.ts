import { defineStore } from 'pinia';
import { Status } from '@/lib/interfaces/status.interface';
import statusApi from '@/api/status.api';


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
    async getStatus(statusId: string) {
      this.current = await statusApi.show(statusId);
    },

    async getStatuses() {
      this.index = await statusApi.index()
    },

    async postStatus(status: Omit<Status, 'id'>) {
      const newStatus = await statusApi.create(status);
      this.index.push(newStatus);
      return newStatus;
    },

    async putStatus(statusId: string, status: Omit<Status, 'id'>) {
      const updatedStatus = await statusApi.update(statusId, status);
      const index = this.index.findIndex(s => s.id === statusId);
      if (index !== -1) {
        this.index[index] = updatedStatus;
      }
      return updatedStatus;
    },

    async deleteStatus(statusId: string) {
      await statusApi.destroy(statusId);
      this.index = this.index.filter(s => s.id !== statusId);
    },
  },
});
