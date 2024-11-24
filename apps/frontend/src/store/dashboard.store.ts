import { defineStore } from 'pinia';
import { Dashboard } from '../lib/interfaces/dashboard.interface';
import dashboardApi from '../api/dashboard.api';

interface DashboardStore {
  current: Dashboard | null;
  index: Dashboard[];
}

export const useDashboardStore = defineStore('dashboard', {
  state: (): DashboardStore => {
    return {
      current: null,
      index: [],
    };
  },
  actions: {
    getDashboard(dashboardId: number) {
    },

    async getDashboards() {
      this.index = await dashboardApi.indexDashboard();
    },
  },
});
