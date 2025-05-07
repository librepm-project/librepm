import { defineStore } from 'pinia';
import { User } from '@/lib/interfaces/user.interface';

interface UserStore {
  current: User | null;
  index: User[];
}

export const useUserStore = defineStore('user', {
  state: (): UserStore => {
    return {
      current: null,
      index: [],
    };
  },
  actions: {
    getUser(userId: number) {
    },

    getUsers() {
      this.index = [];
    },
  },
});
