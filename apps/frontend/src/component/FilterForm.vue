<template>
  <v-form
    ref="form"
    class="form-container"
    @submit.prevent="submit"
  >
    <v-container>
      <v-card
        elevation="0"
        class="rounded-xl pa-6 mb-6"
      >
        <v-card-title class="pa-0 mb-6 text-h6 font-weight-bold">
          {{ props.submitButtonText === 'global.create' ? 'New Filter' : 'Edit Filter' }}
        </v-card-title>

        <v-row>
          <v-col
            cols="12"
            md="8"
          >
            <v-text-field
              v-model="name"
              :rules="[requiredRule]"
              :label="$t('filter.name')"
              hint="e.g., Open Issues, High Priority"
              outlined
              dense
            />
          </v-col>
        </v-row>

        <v-row>
          <v-col cols="12">
            <v-textarea
              v-model="description"
              :label="$t('global.description')"
              hint="Describe what this filter is used for"
              outlined
              dense
              rows="3"
            />
          </v-col>
        </v-row>

        <v-divider class="my-6" />

        <v-card-title class="d-flex align-center justify-space-between px-0 mb-4">
          <span class="text-subtitle-2 font-weight-bold">Display Options</span>
          <v-btn
            icon
            size="small"
            variant="text"
            color="primary"
            :title="showDisplayOptions ? 'Hide display options' : 'Show display options'"
            @click="showDisplayOptions = !showDisplayOptions"
          >
            <v-icon>{{ showDisplayOptions ? 'mdi-chevron-up' : 'mdi-chevron-down' }}</v-icon>
          </v-btn>
        </v-card-title>

        <v-expand-transition>
          <v-row
            v-if="showDisplayOptions"
            class="gap-4"
          >
            <v-col
              cols="12"
              md="6"
            >
              <label class="text-body2 font-weight-medium mb-2 d-block">Columns (drag to reorder)</label>
              <div class="columns-container d-flex flex-column gap-2 mb-3">
                <div
                  v-for="(column, index) in columnList"
                  :key="`col-${column}`"
                  draggable="true"
                  class="column-item pa-2 rounded d-flex align-center justify-space-between"
                  :style="{ background: draggedIndex === index ? 'rgba(63, 81, 181, 0.2)' : 'rgba(0,0,0,0.04)' }"
                  @dragstart="dragStart($event, index)"
                  @dragover.prevent="dragOver($event, index)"
                  @drop="dragDrop($event, index)"
                  @dragend="dragEnd"
                >
                  <span class="text-body2 text-truncate flex-grow-1">{{ getColumnLabel(column) }}</span>
                  <v-icon
                    size="x-small"
                    class="text-medium-emphasis ms-2"
                  >
                    mdi-drag
                  </v-icon>
                  <v-btn
                    icon
                    size="x-small"
                    variant="text"
                    color="error"
                    class="ms-1"
                    @click="removeColumn(index)"
                  >
                    <v-icon size="x-small">
                      mdi-close
                    </v-icon>
                  </v-btn>
                </div>
              </div>
            </v-col>
            <v-col
              cols="12"
              md="6"
            >
              <v-select
                v-model="columnToAdd"
                :items="columnsNotSelected"
                label="Add Column"
                item-title="label"
                item-value="key"
                dense
                @update:model-value="addColumn"
              />
              <v-select
                v-model="groupBy"
                :items="groupByOptions"
                label="Group By"
                dense
              />
            </v-col>
          </v-row>
        </v-expand-transition>

        <v-divider class="my-6" />

        <filter-condition-builder v-model="conditions" />

        <v-divider class="my-6" />

        <v-row class="mt-4">
          <v-col
            cols="12"
            class="d-flex gap-3 align-center"
          >
            <v-btn
              type="submit"
              color="primary"
              size="large"
              prepend-icon="mdi-check"
              rounded="lg"
              class="font-weight-bold"
            >
              {{ $t(props.submitButtonText) }}
            </v-btn>
            <v-btn
              v-if="onDelete"
              type="button"
              color="error"
              variant="text"
              size="large"
              prepend-icon="mdi-delete"
              rounded="lg"
              class="font-weight-bold"
              @click="onDelete"
            >
              {{ $t('global.delete') }}
            </v-btn>
            <v-spacer />
            <v-btn
              type="button"
              variant="outlined"
              color="default"
              size="large"
              rounded="lg"
              @click="$router.back()"
            >
              Cancel
            </v-btn>
          </v-col>
        </v-row>
      </v-card>
    </v-container>
  </v-form>
</template>

<script lang="ts" setup>
import { ref, watch, computed } from 'vue';
import { requiredRule } from '@/lib/formRule';
import { Filter, FilterCondition } from '@/lib/interfaces/filter.interface';
import FilterConditionBuilder from '@/component/FilterConditionBuilder.vue';

const props = defineProps({
    onSubmit: { type: Function, default: undefined },
    onDelete: { type: Function, default: undefined },
    submitButtonText: { type: String, default: undefined },
    filter: { type: Object, default: undefined },
});

const name = ref('');
const description = ref('');
const conditions = ref<FilterCondition[]>([]);
const columnList = ref<string[]>(['key', 'tracker', 'priority', 'summary', 'assignee', 'status']);
const groupBy = ref('');
const form = ref(null);
const columnToAdd = ref<string | null>(null);
const draggedIndex = ref<number | null>(null);
const showDisplayOptions = ref(false);

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
    return allColumns.filter(col => !columnList.value.includes(col.key));
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
    if (columnKey && !columnList.value.includes(columnKey)) {
        columnList.value.push(columnKey);
        columnToAdd.value = null;
    }
};

const removeColumn = (index: number) => {
    columnList.value.splice(index, 1);
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
        const dragged = columnList.value[draggedIndex.value];
        columnList.value.splice(draggedIndex.value, 1);
        columnList.value.splice(index, 0, dragged);
    }
};

const dragEnd = () => {
    draggedIndex.value = null;
};

watch(
    () => props.filter,
    (filter) => {
        if (filter) {
            name.value = filter.name || '';
            description.value = filter.description || '';
            conditions.value = filter.conditions ? [...filter.conditions] : [];
            if (filter.columnList) {
                try {
                    columnList.value = JSON.parse(filter.columnList);
                } catch {
                    columnList.value = ['key', 'tracker', 'priority', 'summary', 'assignee', 'status'];
                }
            }
            groupBy.value = filter.groupBy || '';
            showDisplayOptions.value = true; // <-- Set to true when editing
        }
    },
    { immediate: true },
);

const submit = async () => {
    const { valid } = await (form.value as { validate: () => Promise<{ valid: boolean }> }).validate();
    if (valid) {
        const filterData: Omit<Filter, 'id'> = {
            name: name.value,
            description: description.value,
            conditions: conditions.value,
            columnList: JSON.stringify(columnList.value),
            groupBy: groupBy.value,
        };
        props.onSubmit?.(filterData);
    }
};
</script>

<style scoped>
.form-container {
    max-width: 100%;
}

.gap-3 {
    gap: 1rem;
}
</style>
