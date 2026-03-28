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

    async update(dashboardId: string, widgetId: string, fields: { width?: string }): Promise<void> {
        await api.apiCall().patch(`/dashboard/${dashboardId}/widget/${widgetId}`, fields);
    },

    async reorder(dashboardId: string, items: { id: string; weight: number }[]): Promise<void> {
        await api.apiCall().put(`/dashboard/${dashboardId}/widget/reorder`, items);
    },

    async destroy(dashboardId: string, widgetId: string): Promise<void> {
        await api.apiCall().delete(`/dashboard/${dashboardId}/widget/${widgetId}`);
    },
};
