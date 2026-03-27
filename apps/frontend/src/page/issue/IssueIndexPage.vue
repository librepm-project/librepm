<template>
  <div>
    <div class="mb-4 d-flex align-center gap-2">
      <v-btn
        icon
        variant="tonal"
        color="primary"
        class="ml-auto"
        @click="showSettings = !showSettings"
        :title="showSettings ? 'Settings are visible' : 'Toggle display settings'"
      >
        <v-icon>mdi-cog</v-icon>
      </v-btn>
    </div>

    <v-expand-transition>
      <v-card v-if="showSettings" class="mb-4 pa-4 rounded-lg" elevation="0">
        <v-row class="align-center gap-4">
          <v-col cols="12" md="6">
            <label class="text-body2 font-weight-medium mb-2 d-block">Display Columns (drag to reorder)</label>
            <div class="columns-container d-flex flex-column gap-2">
              <div
                v-for="(column, index) in selectedColumns"
                :key="`col-${column}`"
                draggable="true"
                @dragstart="dragStart($event, index)"
                @dragover.prevent="dragOver($event, index)"
                @drop="dragDrop($event, index)"
                @dragend="dragEnd"
                class="column-item pa-3 rounded d-flex align-center justify-space-between"
                :style="{ background: draggedIndex === index ? 'rgba(63, 81, 181, 0.2)' : 'rgba(0,0,0,0.04)' }"
              >
                <span class="text-body2 flex-grow-1">{{ getColumnLabel(column) }}</span>
                <v-icon size="small" class="text-medium-emphasis me-2">mdi-drag</v-icon>
                <v-btn
                  icon
                  size="x-small"
                  variant="text"
                  color="error"
                  @click="removeColumn(index)"
                >
                  <v-icon size="small">mdi-close</v-icon>
                </v-btn>
              </div>
            </div>
            <v-divider class="my-3" />
            <v-select
              v-model="availableToAdd"
              :items="columnsNotSelected"
              label="Add Column"
              item-title="label"
              item-value="key"
              dense
              @update:model-value="addColumn"
            />
          </v-col>
          <v-col cols="12" md="6">
            <v-select
              v-model="selectedGroupBy"
              :items="groupByOptions"
              label="Group By"
              dense
            />
          </v-col>
        </v-row>
      </v-card>
    </v-expand-transition>

    <issue-table 
      :items="issueStore.index" 
      :columns="selectedColumns"
      :groupBy="selectedGroupBy"
    />
  </div>
</template>

<script lang="ts" setup>
import { useIssueStore } from '@/store/issue.store';
import { useLayoutStore } from '@/store/layout.store';
import { useSettingStore } from '@/store/setting.store';
import IssueTable from '@/component/IssueTable.vue';
import { onMounted, onUnmounted, ref, watch, computed } from 'vue';

const issueStore = useIssueStore();
const layoutStore = useLayoutStore();
const settingStore = useSettingStore();

const showSettings = ref(false);
const selectedColumns = ref<string[]>(['key', 'tracker', 'priority', 'summary', 'assignee', 'status']);
const selectedGroupBy = ref<string>('');
const availableToAdd = ref<string | null>(null);
const draggedIndex = ref<number | null>(null);

const allColumns = [
  { key: 'key', label: 'Key' },
  { key: 'tracker', label: 'Tracker' },
  { key: 'priority', label: 'Priority' },
  { key: 'summary', label: 'Summary' },
  { key: 'assignee', label: 'Assignee' },
  { key: 'status', label: 'Status' },
  { key: 'project', label: 'Project' },
];

const columnsNotSelected = computed(() => {
  return allColumns.filter(col => !selectedColumns.value.includes(col.key));
});

const groupByOptions = [
  { title: 'None', value: '' },
  { title: 'Status', value: 'status' },
  { title: 'Priority', value: 'priority' },
  { title: 'Assignee', value: 'assignee' },
  { title: 'Project', value: 'project' },
];

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

onMounted(async () => {
  await issueStore.getIssues();
  await settingStore.fetchSettings();
  
  const defaultColumns = settingStore.getSettingValue('default_issue_display_columns');
  const defaultGroupBy = settingStore.getSettingValue('default_issue_display_group_by');
  
  if (defaultColumns) {
    try {
      selectedColumns.value = JSON.parse(defaultColumns);
    } catch {
      selectedColumns.value = ['key', 'tracker', 'priority', 'summary', 'assignee', 'status'];
    }
  }
  
  if (defaultGroupBy) {
    selectedGroupBy.value = defaultGroupBy;
  }
  
  layoutStore.setActions([
    {
      text: 'global.create',
      to: '/issue/create',
      color: 'primary',
      icon: 'mdi-plus',
    },
  ]);
});

onUnmounted(() => {
  layoutStore.resetActions();
});

watch(selectedColumns, async (newColumns) => {
  await settingStore.updateSetting('default_issue_display_columns', JSON.stringify(newColumns));
}, { deep: true });

watch(selectedGroupBy, async (newGroupBy) => {
  await settingStore.updateSetting('default_issue_display_group_by', newGroupBy);
});
</script>
