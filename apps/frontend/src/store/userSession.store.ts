import { defineStore } from 'pinia';
import userSessionApi from '@/api/userSession.api';
import { useUserCurrentStore } from './userCurrent.store';
import { LoginRequest } from '@/lib/interfaces/user.interface';

export const useCurrentUserStore = defineStore('userSession', {
  state: () => ({}),
  actions: {
    async postLogin(data: LoginRequest) {
      const session = await userSessionApi.create(data);
      window.localStorage.setItem('accessToken', session.token);
      const userCurrentStore = useUserCurrentStore();
      userCurrentStore.setUser(session.user);
    },
  },
});
