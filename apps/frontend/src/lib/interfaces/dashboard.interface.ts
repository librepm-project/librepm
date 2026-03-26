import { Filter } from './filter.interface';

export interface Dashboard {
    id?: string;
    name: string;
    description?: string;
}

// Base interface – minden widget type implementálja
export interface DashboardWidget {
    id?: string;
    dashboardId: string;
    type: string;
    title: string;
    weight: number;
    width: string;
}

// Filter widget – implementálja a DashboardWidget-et
export interface DashboardFilterWidget extends DashboardWidget {
    type: 'filter';
    filterId: string;
    filter?: Filter;
}

// Union type – ez a tényleges használati hely a store-ban és komponensekben
// Jövőbeli widget type-ok ide kerülnek: | DashboardChartWidget | ...
export type AnyDashboardWidget = DashboardFilterWidget;

export function isFilterWidget(w: AnyDashboardWidget): w is DashboardFilterWidget {
    return w.type === 'filter';
}
