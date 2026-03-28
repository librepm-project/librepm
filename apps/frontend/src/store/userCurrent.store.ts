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
    async get() {
      this.current = await userCurrentApi.getCurrent();
    },
    set(user: User | null) {
      this.current = user;
    },
    async updateCurrent(user: Partial<User>) {
      await userCurrentApi.updateCurrent(user);
      await this.get();
    },
  },
});
