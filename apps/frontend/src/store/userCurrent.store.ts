import { defineStore } from 'pinia';
import { User } from '@/lib/interfaces/user.interface';
import userCurrentApi from '@/api/userCurrent.api';

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
    async getUser() {
      this.current = await userCurrentApi.create();
    },
    setUser(user: User) {
      this.current = user;
    }
  },
});
