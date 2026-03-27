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
      this.current = await userCurrentApi.getCurrent();
    },
    setUser(user: User | null) {
      this.current = user;
    },
    async updateCurrent(user: Partial<User>) {
      await userCurrentApi.updateCurrent(user);
      await this.getUser();
    },
  },
});
