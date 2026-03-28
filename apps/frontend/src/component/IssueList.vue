<template>
  <div class="issue-list-container">
    <div v-if="showSettingsToggle" class="mb-4 d-flex align-center gap-2">
      <slot name="actions-prepend"></slot>
      <v-spacer />
      <v-btn
        icon
        variant="tonal"
        color="primary"
        size="small"
        @click="showSettings = !showSettings"
        :title="showSettings ? 'Settings are visible' : 'Toggle display settings'"
      >
        <v-icon size="small">mdi-cog</v-icon>
      </v-btn>
    </div>

    <v-expand-transition>
      <v-card v-if="showSettings" class="mb-4 pa-4 rounded-xl" elevation="0" border>
        <v-row class="align-center">
          <v-col cols="12" md="6">
            <div class="text-subtitle-2 font-weight-bold mb-3 d-flex align-center">
              <v-icon size="small" class="me-2">mdi-view-column-outline</v-icon>
              {{ $t('issue.display_columns') || 'Display Columns' }}
            </div>
            <div class="columns-container d-flex flex-wrap gap-2 mb-4">
              <div
                v-for="(column, index) in selectedColumns"
                :key="`col-${column}`"
                draggable="true"
                @dragstart="dragStart($event, index)"
                @dragover.prevent="dragOver($event, index)"
                @drop="dragDrop($event, index)"
                @dragend="dragEnd"
                class="column-item pa-2 rounded-lg d-flex align-center gap-2"
                :style="{ background: draggedIndex === index ? 'rgba(var(--v-theme-primary), 0.1)' : 'rgba(0,0,0,0.03)', border: '1px solid rgba(0,0,0,0.05)' }"
              >
                <v-icon size="x-small" class="text-medium-emphasis cursor-grab">mdi-drag-vertical</v-icon>
                <span class="text-caption font-weight-medium">{{ getColumnLabel(column) }}</span>
                <v-btn
                  icon
                  size="x-small"
                  variant="text"
                  color="error"
                  @click="removeColumn(index)"
                >
                  <v-icon size="x-small">mdi-close</v-icon>
                </v-btn>
              </div>
            </div>
            <v-select
              v-model="availableToAdd"
              :items="columnsNotSelected"
              label="Add Column"
              item-title="label"
              item-value="key"
              density="compact"
              variant="outlined"
              hide-details
              rounded="lg"
              @update:model-value="addColumn"
            />
          </v-col>
          <v-col cols="12" md="6">
            <div class="text-subtitle-2 font-weight-bold mb-3 d-flex align-center">
              <v-icon size="small" class="me-2">mdi-format-list-group</v-icon>
              {{ $t('issue.group_by') || 'Group By' }}
            </div>
            <v-select
              v-model="selectedGroupBy"
              :items="groupByOptions"
              label="Group By"
              density="compact"
              variant="outlined"
              hide-details
              rounded="lg"
            />
            <p class="text-caption text-medium-emphasis mt-3">
              Organize issues into groups based on the selected field.
            </p>
          </v-col>
        </v-row>
      </v-card>
    </v-expand-transition>

    <div v-if="loading" class="d-flex justify-center py-8">
      <v-progress-circular indeterminate color="primary" />
    </div>
    <div v-else>
        <issue-table
          v-if="internalItems.length > 0"
          :items="internalItems"
          :columns="selectedColumns"
          :group-by="selectedGroupBy"
        />
        <div v-else class="text-center py-12 text-medium-emphasis">
            <v-icon size="x-large" class="mb-2">mdi-tray-blank</v-icon>
            <p>{{ $t('issue.no_issues_found') || 'No issues found' }}</p>
        </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, computed, onMounted } from 'vue';
import { useI18n } from 'vue-i18n';
import { Issue } from '@/lib/interfaces/issue.interface';
import { Filter } from '@/lib/interfaces/filter.interface';
import { useFilterStore } from '@/store/filter.store';
import { useSettingStore } from '@/store/setting.store';
import issueApi from '@/api/issue.api';
import filterApi from '@/api/filter.api';
import IssueTable from './IssueTable.vue';

interface Props {
  filterId?: string;
  items?: Issue[];
  showSettingsToggle?: boolean;
  persistMode?: 'global' | 'filter' | 'none';
}

const props = withDefaults(defineProps<Props>(), {
  showSettingsToggle: false,
  persistMode: 'none',
});

const { t } = useI18n();
const filterStore = useFilterStore();
const settingStore = useSettingStore();

const internalItems = ref<Issue[]>([]);
const filterData = ref<Filter | null>(null);
const loading = ref(false);
const showSettings = ref(false);

const selectedColumns = ref<string[]>(['key', 'tracker', 'priority', 'summary', 'assignee', 'status']);
const selectedGroupBy = ref<string>('');
const availableToAdd = ref<string | null>(null);
const draggedIndex = ref<number | null>(null);

const allColumns = [
  { key: 'key', label: t('issue.key') },
  { key: 'tracker', label: t('issue.tracker') },
  { key: 'priority', label: t('issue.priority') },
  { key: 'summary', label: t('issue.summary') },
  { key: 'assignee', label: t('issue.assignee') },
  { key: 'status', label: t('issue.status') },
  { key: 'project', label: t('issue.project') },
];

const groupByOptions = [
  { title: 'None', value: '' },
  { title: 'Status', value: 'status' },
  { title: 'Priority', value: 'priority' },
  { title: 'Assignee', value: 'assignee' },
  { title: 'Project', value: 'project' },
];

const columnsNotSelected = computed(() => {
  return allColumns.filter(col => !selectedColumns.value.includes(col.key));
});

const getColumnLabel = (key: string): string => {
  return allColumns.find(col => col.key === key)?.label || key;
};

const addColumn = (columnKey: string | null) => {
  if (columnKey) {
    selectedColumns.value.push(columnKey);
    availableToAdd.value = null;
  }
};

const removeColumn = (index: number) => {
  selectedColumns.value.splice(index, 1);
};

const dragStart = (event: DragEvent, index: number) => {
  draggedIndex.value = index;
  if (event.dataTransfer) {
    event.dataTransfer.effectAllowed = 'move';
  }
};

const dragOver = (event: DragEvent, index: number) => {
  event.preventDefault();
  if (event.dataTransfer) {
    event.dataTransfer.dropEffect = 'move';
  }
};

const dragDrop = (event: DragEvent, index: number) => {
  event.preventDefault();
  if (draggedIndex.value !== null && draggedIndex.value !== index) {
    const dragged = selectedColumns.value[draggedIndex.value];
    selectedColumns.value.splice(draggedIndex.value, 1);
    selectedColumns.value.splice(index, 0, dragged);
  }
};

const dragEnd = () => {
  draggedIndex.value = null;
};

const loadData = async () => {
  if (props.items) {
    internalItems.value = props.items;
  } else {
    loading.value = true;
    try {
      internalItems.value = await issueApi.index({ filterId: props.filterId });
    } catch (error) {
        console.error("Failed to load issues", error);
    } finally {
      loading.value = false;
    }
  }

  if (props.filterId) {
    try {
        const response = await filterApi.show(props.filterId);
        filterData.value = response;
        if (filterData.value) {
          if (filterData.value.columnList) {
            try {
              selectedColumns.value = JSON.parse(filterData.value.columnList);
            } catch {
              // keep default
            }
          }
          selectedGroupBy.value = filterData.value.groupBy || '';
        }
    } catch (error) {
        console.error("Failed to load filter settings", error);
    }
  } else if (props.persistMode === 'global') {
    await settingStore.fetchSettings();
    const defaultColumns = settingStore.getSettingValue('default_issue_display_columns');
    const defaultGroupBy = settingStore.getSettingValue('default_issue_display_group_by');
    
    if (defaultColumns) {
      try {
        selectedColumns.value = JSON.parse(defaultColumns);
      } catch {
        // keep default
      }
    }
    if (defaultGroupBy) {
      selectedGroupBy.value = defaultGroupBy;
    }
  }
};

onMounted(loadData);

watch(() => props.items, (newItems) => {
  if (newItems) {
    internalItems.value = newItems;
  }
}, { deep: true });

watch(() => props.filterId, loadData);

watch(selectedColumns, async (newColumns) => {
  if (props.persistMode === 'global') {
    await settingStore.updateSetting('default_issue_display_columns', JSON.stringify(newColumns));
  } else if (props.persistMode === 'filter' && props.filterId && filterData.value) {
    await filterStore.putFilter(props.filterId, {
      ...filterData.value,
      columnList: JSON.stringify(newColumns),
    });
  }
}, { deep: true });

watch(selectedGroupBy, async (newGroupBy) => {
  if (props.persistMode === 'global') {
    await settingStore.updateSetting('default_issue_display_group_by', newGroupBy);
  } else if (props.persistMode === 'filter' && props.filterId && filterData.value) {
    await filterStore.putFilter(props.filterId, {
      ...filterData.value,
      groupBy: newGroupBy,
    });
  }
});
</script>

<style scoped>
.column-item {
  transition: all 0.2s;
  user-select: none;
}
.cursor-grab {
  cursor: grab;
}
.cursor-grab:active {
  cursor: grabbing;
}
.gap-2 {
    gap: 0.5rem;
}
.gap-4 {
    gap: 1rem;
}
</style>
