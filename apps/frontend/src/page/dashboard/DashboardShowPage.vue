<template>
  <v-container fluid>
    <div class="d-flex align-center justify-space-between mb-5">
      <v-menu transition="slide-y-transition">
        <template #activator="{ props, isActive }">
          <v-btn
            v-bind="props"
            variant="tonal"
            color="primary"
            rounded="xl"
            size="large"
            prepend-icon="mdi-view-dashboard"
            :append-icon="isActive ? 'mdi-chevron-up' : 'mdi-chevron-down'"
            class="dashboard-switcher-btn"
          >
            {{ dashboardStore.current?.name ?? $t('title.dashboardShow') }}
          </v-btn>
        </template>
        <v-card
          rounded="xl"
          elevation="4"
          min-width="240"
          class="mt-1"
        >
          <v-list>
            <v-list-subheader class="text-caption font-weight-bold text-uppercase">
              {{ $t('dashboard.dashboards') }}
            </v-list-subheader>
            <v-list-item
              v-for="dashboard in dashboardStore.index"
              :key="dashboard.id"
              :to="`/dashboard/${dashboard.id}`"
              rounded="lg"
              :active="dashboard.id === dashboardStore.current?.id"
              active-color="primary"
              class="mx-2 mb-1"
            >
              <template #prepend>
                <v-icon :color="dashboard.id === dashboardStore.current?.id ? 'primary' : undefined">
                  mdi-view-dashboard-outline
                </v-icon>
              </template>
              <v-list-item-title class="font-weight-medium">
                {{ dashboard.name }}
              </v-list-item-title>
              <template
                v-if="dashboard.id === dashboardStore.current?.id"
                #append
              >
                <v-icon
                  size="small"
                  color="primary"
                >
                  mdi-check
                </v-icon>
              </template>
            </v-list-item>
          </v-list>
        </v-card>
      </v-menu>
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

    <div
      v-if="dashboardStore.widgets.length > 0"
      class="widget-grid"
    >
      <div
        v-for="(widget, index) in dashboardStore.widgets"
        :key="widget.id"
        class="widget-col"
        :class="[
          widthClass(widget.width),
          { 'widget--dragging': draggingIndex === index },
          { 'drop-before': dropTarget === index && dropPosition === 'before' },
          { 'drop-after': dropTarget === index && dropPosition === 'after' },
        ]"
        draggable="true"
        @dragstart="onDragStart($event, index)"
        @dragend="onDragEnd"
        @dragover.prevent="onDragOver($event, index)"
        @dragleave="onDragLeave($event)"
        @drop.prevent="onDrop(index)"
      >
        <component
          :is="widgetComponents[widget.type]"
          v-if="widgetComponents[widget.type]"
          :widget="widget"
          @remove="handleRemoveWidget"
          @resize="(width: string) => handleResizeWidget(widget.id!, width)"
        />
      </div>
    </div>

    <div
      v-else
      class="text-center py-12 text-grey"
    >
      <v-icon size="x-large">
        mdi-view-dashboard-outline
      </v-icon>
      <p class="text-h6 mt-3">
        {{ $t('dashboard.no_widgets') }}
      </p>
    </div>

    <!-- Add Widget Dialog -->
    <v-dialog
      v-model="addWidgetDialog"
      max-width="480"
    >
      <v-card
        rounded="xl"
        class="pa-2"
      >
        <v-card-title class="pa-4 pb-2">
          {{ $t('dashboard.add_widget') }}
        </v-card-title>
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
          <v-btn
            variant="text"
            @click="addWidgetDialog = false"
          >
            Cancel
          </v-btn>
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
import { useSettingStore } from '@/store/setting.store';
import { AnyDashboardWidget } from '@/lib/interfaces/dashboard.interface';
import DashboardFilterWidgetComponent from '@/component/DashboardFilterWidget.vue';

const { t } = useI18n();
const route = useRoute();
const router = useRouter();
const dashboardStore = useDashboardStore();
const filterStore = useFilterStore();
const layoutStore = useLayoutStore();
const settingStore = useSettingStore();

const widgetComponents: Record<string, Component> = {
    filter: DashboardFilterWidgetComponent,
};

const widgetTypeOptions = [
    { value: 'filter', label: 'Filter' },
];

// --- Width ---
function widthClass(width: string): string {
    switch (width) {
        case '1/2': return 'widget-col--half';
        case '2/3': return 'widget-col--two-thirds';
        case 'full': return 'widget-col--full';
        default:    return 'widget-col--third';
    }
}

// --- Add Widget dialog ---
const addWidgetDialog = ref(false);
const newWidget = ref({ type: 'filter', title: '', filterId: '' });

const canAddWidget = computed(() => {
    if (!newWidget.value.title) return false;
    if (newWidget.value.type === 'filter' && !newWidget.value.filterId) return false;
    return true;
});

// --- Drag & Drop ---
const draggingIndex = ref<number | null>(null);
const dropTarget = ref<number | null>(null);
const dropPosition = ref<'before' | 'after' | null>(null);

function onDragStart(event: DragEvent, index: number) {
    draggingIndex.value = index;
    if (event.dataTransfer) event.dataTransfer.effectAllowed = 'move';
}

function onDragEnd() {
    draggingIndex.value = null;
    dropTarget.value = null;
    dropPosition.value = null;
}

function onDragOver(event: DragEvent, index: number) {
    if (draggingIndex.value === null || draggingIndex.value === index) return;
    const rect = (event.currentTarget as Element).getBoundingClientRect();
    dropTarget.value = index;
    dropPosition.value = event.clientY < rect.top + rect.height / 2 ? 'before' : 'after';
}

function onDragLeave(event: DragEvent) {
    const related = event.relatedTarget as Node | null;
    if (related && (event.currentTarget as Element).contains(related)) return;
    dropTarget.value = null;
    dropPosition.value = null;
}

function onDrop(targetIndex: number) {
    const fromIndex = draggingIndex.value;
    const pos = dropPosition.value;
    draggingIndex.value = null;
    dropTarget.value = null;
    dropPosition.value = null;

    if (fromIndex === null || fromIndex === targetIndex || pos === null) return;

    const widgets = [...dashboardStore.widgets];
    const [moved] = widgets.splice(fromIndex, 1);

    let insertAt = targetIndex;
    if (fromIndex < targetIndex) insertAt = targetIndex - 1;
    if (pos === 'after') insertAt += 1;
    insertAt = Math.max(0, Math.min(insertAt, widgets.length));

    widgets.splice(insertAt, 0, moved);

    const dashboardId = route.params.dashboardId as string;
    dashboardStore.reorderWidgets(dashboardId, widgets);
}

// --- Load ---
const loadDashboard = async () => {
    const dashboardId = route.params.dashboardId as string | undefined;

    await Promise.all([
        dashboardStore.getAll(),
        settingStore.fetch(),
    ]);

    layoutStore.setActions([
        {
            text: 'global.create',
            to: '/dashboard/create',
            color: 'primary',
            icon: 'mdi-plus',
        },
    ]);

    if (dashboardId) {
        await dashboardStore.get(dashboardId);
        await dashboardStore.loadWidgets(dashboardId);
        await filterStore.getAll();
    } else {
        // Alapértelmezett oldal választása
        const defaultBoardId = settingStore.getValue('default_board_id');
        const defaultDashboardId = settingStore.getValue('default_dashboard_id');

        if (defaultBoardId) {
            router.replace(`/board/${defaultBoardId}`);
            return;
        }

        if (defaultDashboardId) {
            router.replace(`/dashboard/${defaultDashboardId}`);
            return;
        }

        if (dashboardStore.index.length > 0) {
            router.replace(`/dashboard/${dashboardStore.index[0].id}`);
        }
    }
};

onMounted(loadDashboard);
watch(() => route.params.dashboardId, loadDashboard);

onUnmounted(() => {
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

async function handleResizeWidget(widgetId: string, width: string) {
    const dashboardId = route.params.dashboardId as string;
    await dashboardStore.updateWidgetWidth(dashboardId, widgetId, width);
}

async function handleRemoveWidget(widgetId: string) {
    const dashboardId = route.params.dashboardId as string;
    await dashboardStore.removeWidget(dashboardId, widgetId);
}
</script>

<style scoped>
.dashboard-switcher-btn {
    font-size: 1.1rem;
    font-weight: 700;
    letter-spacing: 0;
    text-transform: none;
    padding-inline: 20px;
}

/* Grid layout */
.widget-grid {
    display: flex;
    flex-wrap: wrap;
    gap: 16px;
    align-items: flex-start;
}

.widget-col {
    min-width: 0;
    transition: opacity 0.15s;
}

.widget-col--third {
    flex: 0 0 calc(33.333% - 16px * 2 / 3);
    width: calc(33.333% - 16px * 2 / 3);
}

.widget-col--half {
    flex: 0 0 calc(50% - 8px);
    width: calc(50% - 8px);
}

.widget-col--two-thirds {
    flex: 0 0 calc(66.666% - 16px / 3);
    width: calc(66.666% - 16px / 3);
}

.widget-col--full {
    flex: 0 0 100%;
    width: 100%;
}

/* Responsive: small screens */
@media (max-width: 960px) {
    .widget-col--third,
    .widget-col--half,
    .widget-col--two-thirds {
        flex: 0 0 calc(50% - 8px);
        width: calc(50% - 8px);
    }
}

@media (max-width: 600px) {
    .widget-col--third,
    .widget-col--half,
    .widget-col--two-thirds,
    .widget-col--full {
        flex: 0 0 100%;
        width: 100%;
    }
}

.widget--dragging {
    opacity: 0.35;
}

/* Drop indicators via pseudo-elements */
.widget-col {
    position: relative;
}

.drop-before::before,
.drop-after::after {
    content: '';
    display: block;
    height: 3px;
    border-radius: 2px;
    background-color: rgb(var(--v-theme-primary));
    box-shadow: 0 0 6px rgba(var(--v-theme-primary), 0.5);
}

.drop-before::before {
    margin-bottom: 10px;
}

.drop-after::after {
    margin-top: 10px;
}
</style>
