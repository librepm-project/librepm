import { defineStore } from 'pinia';
import { Dashboard, AnyDashboardWidget } from '@/lib/interfaces/dashboard.interface';
import dashboardApi from '@/api/dashboard.api';
import dashboardWidgetApi from '@/api/dashboard-widget.api';
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

        async loadWidgets(dashboardId: string) {
            this.widgets = await dashboardWidgetApi.index(dashboardId);
        },

        async addWidget(dashboardId: string, widget: { type: string; title: string; filterId?: string }) {
            const created = await dashboardWidgetApi.create(dashboardId, widget as any);
            this.widgets.push(created);
            return created;
        },

        async updateWidgetWidth(dashboardId: string, widgetId: string, width: string) {
            const w = this.widgets.find(w => w.id === widgetId);
            if (w) w.width = width;
            await dashboardWidgetApi.update(dashboardId, widgetId, { width });
        },

        async reorderWidgets(dashboardId: string, ordered: AnyDashboardWidget[]) {
            const items = ordered.map((w, i) => ({ id: w.id!, weight: i }));
            this.widgets = ordered.map((w, i) => ({ ...w, weight: i }));
            await dashboardWidgetApi.reorder(dashboardId, items);
        },

        async removeWidget(dashboardId: string, widgetId: string) {
            await dashboardWidgetApi.destroy(dashboardId, widgetId);
            this.widgets = this.widgets.filter(w => w.id !== widgetId);
        },
    },
});
