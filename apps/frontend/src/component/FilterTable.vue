<template>
  <v-card
    elevation="0"
    class="rounded-xl"
  >
    <v-data-table
      :items="items"
      :headers="headers"
      density="comfortable"
      class="transparent-table"
      item-key="id"
      @click:row="(e, { item }) => onView(item)"
    >
      <template #header>
        <thead class="bg-primary text-white">
          <tr>
            <th
              v-for="header in headers"
              :key="header.key"
              :class="header.align ? `text-${header.align}` : ''"
              class="text-white font-weight-bold"
            >
              {{ header.title }}
            </th>
          </tr>
        </thead>
      </template>

      <template #[`item.name`]="{ item }">
        {{ item.name }}
      </template>

      <template #[`item.conditions`]="{ item }">
        <v-chip
          v-if="item.conditions && item.conditions.length > 0"
          size="small"
          label
        >
          {{ item.conditions.length }} condition(s)
        </v-chip>
        <v-chip
          v-else
          size="small"
          variant="outlined"
          label
        >
          No conditions
        </v-chip>
      </template>

      <template #[`item.actions`]="{ item }">
        <div class="d-flex justify-end">
          <v-icon
            icon="mdi-pencil"
            class="cursor-pointer"
            @click.stop="onEdit(item)"
          />
        </div>
      </template>
    </v-data-table>
  </v-card>
</template>

<script setup lang="ts">
import { Filter } from '@/lib/interfaces/filter.interface';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();

defineProps<{
  items: Filter[],
  onEdit: (item: Filter) => void,
  onView: (item: Filter) => void, // New prop for viewing
}>();

const headers = [
  { key: 'name', title: t('filter.name') },
  { key: 'description', title: t('global.description') },
  { key: 'conditions', title: 'Conditions' },
  { key: 'actions', title: '', align: 'right' }, // Added actions header
];
</script>

<style scoped>
.transparent-table {
  background-color: transparent !important;
}

:deep(.v-data-table) {
  background-color: transparent !important;
}

:deep(.v-data-table thead tr) {
  background-color: transparent !important;
}

:deep(.v-data-table tbody tr:hover) {
  background-color: rgba(63, 81, 181, 0.05) !important;
}

:deep(.v-data-table td) {
  border-bottom: 1px solid rgba(0, 0, 0, 0.08) !important;
}

.cursor-pointer {
    cursor: pointer;
}
</style>