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
  },
});
