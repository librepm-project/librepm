import { defineStore } from 'pinia';
import { User } from '@/lib/interfaces/user.interface';
import userCurrentApi from '@/api/user-current.api';

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
      this.current = await userCurrentApi.get();
    },
    set(user: User | null) {
      this.current = user;
    },
    async update(user: Partial<User>) {
      await userCurrentApi.update(user);
      await this.get();
    },
  },
});
