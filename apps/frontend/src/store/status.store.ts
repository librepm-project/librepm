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
    getStatus(statusId: number) {
    },

    async getStatuses() {
      this.index = await statusApi.index()
    },
  },
});
