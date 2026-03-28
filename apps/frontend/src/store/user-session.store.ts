import { defineStore } from 'pinia';
import userSessionApi from '@/api/user-session.api';
import { useUserCurrentStore } from './user-current.store';
import { LoginRequest } from '@/lib/interfaces/user.interface';
import { setToken, removeToken } from '@/lib/cookie';

export const useUserSessionStore = defineStore('userSession', {
  state: () => ({}),
  actions: {
    async postLogin(data: LoginRequest) {
      const session = await userSessionApi.create(data);
      setToken(session.token);
      const userCurrentStore = useUserCurrentStore();
      userCurrentStore.set(session.user);
    },
    logout() {
      removeToken();
      const userCurrentStore = useUserCurrentStore();
      userCurrentStore.set(null);
    },
  },
});
