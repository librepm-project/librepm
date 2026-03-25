import { defineStore } from 'pinia';
import { Dashboard, AnyDashboardWidget } from '@/lib/interfaces/dashboard.interface';
import dashboardApi from '@/api/dashboard.api';
import dashboardWidgetApi from '@/api/dashboardWidget.api';
import { createCrudActions } from './utils/crudMixin';

interface DashboardStore {
    current: Dashboard | null;
    index: Dashboard[];
    widgets: AnyDashboardWidget[];
}

export const useDashboardStore = defineStore('dashboard', {
    state: (): DashboardStore => ({
        current: null,
        index: [],
        widgets: [],
    }),
    actions: {
        ...createCrudActions<Dashboard>(dashboardApi),

        async getDashboard(dashboardId: string) {
            return this.getCurrentItem(dashboardId);
        },

        async getDashboards() {
            return this.getAllItems();
        },

        async loadWidgets(dashboardId: string) {
            this.widgets = await dashboardWidgetApi.index(dashboardId);
        },

        async addWidget(dashboardId: string, widget: { type: string; title: string; filterId?: string }) {
            const created = await dashboardWidgetApi.create(dashboardId, widget as any);
            this.widgets.push(created);
            return created;
        },

        async removeWidget(dashboardId: string, widgetId: string) {
            await dashboardWidgetApi.destroy(dashboardId, widgetId);
            this.widgets = this.widgets.filter(w => w.id !== widgetId);
        },
    },
});
