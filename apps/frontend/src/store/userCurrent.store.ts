import { defineStore } from 'pinia';
import { User } from '../lib/interfaces/user.interface';
import { userMock } from '../lib/mock-data';

interface UserCurrentStore {
  current: User | null;
}

export const useUserCurrentStore = defineStore('userCurrent', {
  state: (): UserCurrentStore => {
    return {
      current: null,
    };
  },
  actions: {
    getUser() {
      this.current = userMock[0];
    },
  },
});
