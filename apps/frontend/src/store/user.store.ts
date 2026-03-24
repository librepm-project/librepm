import { defineStore } from 'pinia';
import { User } from '@/lib/interfaces/user.interface';
import userApi from '@/api/user.api';

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
    async getUser(userId: string) {
      this.current = await userApi.show(userId);
    },

    async getUsers() {
      this.index = await userApi.index();
    },

    async postUser(user: Omit<User, 'id'>) {
      const newUser = await userApi.create(user);
      this.index.push(newUser);
      return newUser;
    },

    async putUser(userId: string, user: Omit<User, 'id'>) {
      const updatedUser = await userApi.update(userId, user);
      const index = this.index.findIndex(u => u.id === userId);
      if (index !== -1) {
        this.index[index] = updatedUser;
      }
      return updatedUser;
    },

    async deleteUser(userId: string) {
      await userApi.destroy(userId);
      this.index = this.index.filter(u => u.id !== userId);
    },
  },
});
