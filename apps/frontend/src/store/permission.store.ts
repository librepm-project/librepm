import { defineStore } from 'pinia';
import userPermissionApi from '@/api/user-permission.api';
import type { Permission } from '@/lib/permissions';

interface PermissionState {
  permissions: string[];
  loaded: boolean;
}

export const usePermissionStore = defineStore('permission', {
  state: (): PermissionState => ({
    permissions: [],
    loaded: false,
  }),
  actions: {
    async fetch() {
      this.permissions = await userPermissionApi.index();
      this.loaded = true;
    },
    clear() {
      this.permissions = [];
      this.loaded = false;
    },
  },
  getters: {
    can: (state) => (permission: Permission): boolean => {
      return state.permissions.includes(permission);
    },
  },
});
