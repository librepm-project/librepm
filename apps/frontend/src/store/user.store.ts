import { defineStore } from 'pinia';
import { User } from '../lib/interfaces/user.interface';

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
      this.current = {
        email: "example1@company.org",
        firstName: "John",
        lastName: "Doe",
      };
    },

    getUsers() {
      this.index = [
        {
            email: "example1@company.org",
            firstName: "John",
            lastName: "Doe",
        },
        {
            email: "example2@company.org",
            firstName: "Joe",
            lastName: "Smish",
        },
      ];
    },
  },
});
