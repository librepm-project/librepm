import { defineStore } from 'pinia';
import { User } from '@/lib/interfaces/user.interface';
import { userMock } from '@/lib/mock-data';

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
      this.current = userMock[0];
    },

    getUsers() {
      this.index = userMock;
    },
  },
});
