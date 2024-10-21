import { defineStore } from 'pinia';
import { User } from '../lib/interfaces/user.interface';

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
      this.current = {
        email: "example1@company.org",
        firstName: "John",
        lastName: "Doe",
      };
    },
  },
});
