import { defineStore } from 'pinia';
import { User } from '@/lib/interfaces/user.interface';
import userApi from '@/api/user.api';
import { createCrudActions } from './utils/crudMixin';

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
    ...createCrudActions<User>(userApi),

    // Alias methods for backward compatibility
    async getUser(userId: string) {
      return this.getCurrentItem(userId);
    },

    async getUsers() {
      return this.getAllItems();
    },

    async postUser(user: Omit<User, 'id'>) {
      return this.createItem(user);
    },

    async putUser(userId: string, user: Omit<User, 'id'>) {
      return this.updateItem(userId, user);
    },

    async deleteUser(userId: string) {
      return this.deleteItem(userId);
    },
  },
});
