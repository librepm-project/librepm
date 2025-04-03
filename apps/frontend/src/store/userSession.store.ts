import { defineStore } from 'pinia';
import userSessionApi from '@/api/userSession.api';
import { useUserCurrentStore } from './userCurrent.store';
import { LoginRequest } from '@/lib/interfaces/user.interface';

export const useUserSessionStore = defineStore('userSession', {
  state: () => ({}),
  actions: {
    async postLogin(data: LoginRequest) {
      const session = await userSessionApi.create(data);
      console.log("set localStorage", session.token)
      window.localStorage.setItem('accessToken', session.token);
      const userCurrentStore = useUserCurrentStore();
      userCurrentStore.setUser(session.user);
    },
  },
});
