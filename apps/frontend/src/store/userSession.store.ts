import { defineStore } from 'pinia';
import userSessionApi from '@/api/userSession.api';
import { useUserCurrentStore } from './userCurrent.store';
import { LoginRequest } from '@/lib/interfaces/user.interface';
import { setToken, removeToken } from '@/lib/cookie';

export const useUserSessionStore = defineStore('userSession', {
  state: () => ({}),
  actions: {
    async postLogin(data: LoginRequest) {
      const session = await userSessionApi.create(data);
      setToken(session.token);
      const userCurrentStore = useUserCurrentStore();
      userCurrentStore.setUser(session.user);
    },
    logout() {
      removeToken();
      const userCurrentStore = useUserCurrentStore();
      userCurrentStore.setUser(null);
    },
  },
});
