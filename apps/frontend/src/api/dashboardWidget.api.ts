import api from '@/api/api';
import { AnyDashboardWidget, DashboardWidget } from '@/lib/interfaces/dashboard.interface';

export default {
    async index(dashboardId: string): Promise<AnyDashboardWidget[]> {
        const response = await api.apiCall().get(`/dashboard/${dashboardId}/widget`);
        return response.data;
    },

    async create(dashboardId: string, widget: Omit<DashboardWidget, 'id' | 'dashboardId'>): Promise<AnyDashboardWidget> {
        const response = await api.apiCall().post(`/dashboard/${dashboardId}/widget`, widget);
        return response.data;
    },

    async destroy(dashboardId: string, widgetId: string): Promise<void> {
        await api.apiCall().delete(`/dashboard/${dashboardId}/widget/${widgetId}`);
    },
};
