<template>
    <v-container fluid>
        <div class="d-flex align-center justify-space-between mb-5">
            <span class="text-h6 font-weight-bold">
                {{ dashboardStore.current?.name ?? $t('title.dashboardShow') }}
            </span>
            <v-btn
                color="primary"
                variant="tonal"
                prepend-icon="mdi-plus"
                rounded="lg"
                @click="openAddWidget"
            >
                {{ $t('dashboard.add_widget') }}
            </v-btn>
        </div>

        <v-row v-if="dashboardStore.widgets.length > 0">
            <v-col
                v-for="widget in dashboardStore.widgets"
                :key="widget.id"
                cols="12"
                sm="6"
                lg="4"
            >
                <component
                    :is="widgetComponents[widget.type]"
                    v-if="widgetComponents[widget.type]"
                    :widget="widget"
                    @remove="handleRemoveWidget"
                />
            </v-col>
        </v-row>

        <div v-else class="text-center py-12 text-grey">
            <v-icon size="x-large">mdi-view-dashboard-outline</v-icon>
            <p class="text-h6 mt-3">{{ $t('dashboard.no_widgets') }}</p>
        </div>

        <!-- Add Widget Dialog -->
        <v-dialog v-model="addWidgetDialog" max-width="480">
            <v-card rounded="xl" class="pa-2">
                <v-card-title class="pa-4 pb-2">{{ $t('dashboard.add_widget') }}</v-card-title>
                <v-card-text>
                    <v-select
                        v-model="newWidget.type"
                        :items="widgetTypeOptions"
                        item-title="label"
                        item-value="value"
                        :label="$t('dashboard.widget_type')"
                        class="mb-3"
                    />
                    <v-text-field
                        v-model="newWidget.title"
                        :label="$t('dashboard.widget_title')"
                        class="mb-3"
                    />
                    <v-select
                        v-if="newWidget.type === 'filter'"
                        v-model="newWidget.filterId"
                        :items="filterStore.index"
                        item-title="name"
                        item-value="id"
                        :label="$t('dashboard.select_filter')"
                    />
                </v-card-text>
                <v-card-actions class="pa-4 pt-0">
                    <v-spacer />
                    <v-btn variant="text" @click="addWidgetDialog = false">Cancel</v-btn>
                    <v-btn
                        color="primary"
                        variant="flat"
                        rounded="lg"
                        :disabled="!canAddWidget"
                        @click="handleAddWidget"
                    >
                        {{ $t('global.create') }}
                    </v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </v-container>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted, watch, type Component } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useI18n } from 'vue-i18n';
import { useDashboardStore } from '@/store/dashboard.store';
import { useFilterStore } from '@/store/filter.store';
import { useLayoutStore } from '@/store/layout.store';
import DashboardFilterWidgetComponent from '@/component/DashboardFilterWidget.vue';

const { t } = useI18n();
const route = useRoute();
const router = useRouter();
const dashboardStore = useDashboardStore();
const filterStore = useFilterStore();
const layoutStore = useLayoutStore();

// Widget type → komponens mapping
const widgetComponents: Record<string, Component> = {
    filter: DashboardFilterWidgetComponent,
};

const widgetTypeOptions = [
    { value: 'filter', label: 'Filter' },
];

// Add Widget dialog state
const addWidgetDialog = ref(false);
const newWidget = ref({ type: 'filter', title: '', filterId: '' });

const canAddWidget = computed(() => {
    if (!newWidget.value.title) return false;
    if (newWidget.value.type === 'filter' && !newWidget.value.filterId) return false;
    return true;
});

const loadDashboard = async () => {
    const dashboardId = route.params.dashboardId as string | undefined;

    await dashboardStore.getDashboards();
    layoutStore.setSidebar(dashboardStore.index.map(d => ({
        key: d.id,
        title: d.name,
        link: `/dashboard/${d.id}`,
    })));
    layoutStore.setActions([
        {
            text: 'global.create',
            to: '/dashboard/create',
            color: 'primary',
            icon: 'mdi-plus',
        },
    ]);

    if (dashboardId) {
        await dashboardStore.getDashboard(dashboardId);
        await dashboardStore.loadWidgets(dashboardId);
        await filterStore.getFilters();
    } else if (dashboardStore.index.length > 0) {
        router.replace(`/dashboard/${dashboardStore.index[0].id}`);
    }
};

onMounted(loadDashboard);
watch(() => route.params.dashboardId, loadDashboard);

onUnmounted(() => {
    layoutStore.resetSidebar();
    layoutStore.resetActions();
});

function openAddWidget() {
    newWidget.value = { type: 'filter', title: '', filterId: '' };
    addWidgetDialog.value = true;
}

async function handleAddWidget() {
    const dashboardId = route.params.dashboardId as string;
    await dashboardStore.addWidget(dashboardId, {
        type: newWidget.value.type,
        title: newWidget.value.title,
        filterId: newWidget.value.filterId || undefined,
    } as any);
    addWidgetDialog.value = false;
}

async function handleRemoveWidget(widgetId: string) {
    const dashboardId = route.params.dashboardId as string;
    await dashboardStore.removeWidget(dashboardId, widgetId);
}
</script>
