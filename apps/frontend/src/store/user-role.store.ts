import { defineStore } from 'pinia';
import userRoleApi from '@/api/user-role.api';

interface UserRoleState {
  roles: string[];
  userId: string | null;
}

export const useUserRoleStore = defineStore('userRole', {
  state: (): UserRoleState => ({
    roles: [],
    userId: null,
  }),
  actions: {
    async fetch(userId: string) {
      this.userId = userId;
      this.roles = await userRoleApi.index(userId);
    },
    async add(role: string) {
      if (!this.userId) return;
      await userRoleApi.create(this.userId, role);
      await this.fetch(this.userId);
    },
    async remove(role: string) {
      if (!this.userId) return;
      await userRoleApi.destroy(this.userId, role);
      await this.fetch(this.userId);
    },
  },
});
