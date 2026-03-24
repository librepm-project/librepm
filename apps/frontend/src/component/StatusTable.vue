<template>
  <v-card elevation="0" class="rounded-xl">
    <v-data-table
      :items="items"
      :headers="headers"
      density="comfortable"
      class="transparent-table"
      item-key="id"
    >
      <template #header>
        <thead class="bg-primary text-white">
          <tr>
            <th
              v-for="header in headers"
              :key="header.key"
              class="text-white font-weight-bold"
            >
              {{ header.title }}
            </th>
          </tr>
        </thead>
      </template>

      <template #item.color="{ item }">
        <v-chip
          :style="{ backgroundColor: item.color }"
          text-color="white"
          size="small"
          label
        >
          {{ item.name }}
        </v-chip>
      </template>

      <template #item.actions="{ item }">
        <general-record-actions :item="item" :onEdit="onEdit" :onDelete="onDelete" />
      </template>
    </v-data-table>
  </v-card>
</template>

<script setup lang="ts">
import GeneralRecordActions from '@/component/GeneralRecordActions.vue';
import { Status } from '@/lib/interfaces/status.interface';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();

defineProps<{
  items: Status[];
  onEdit: (item: Status) => void;
  onDelete: (item: Status) => void;
}>()

const headers = [
  { key: 'name', title: t('status.name') },
  { key: 'color', title: t('status.color') },
  { key: 'actions', title: t('global.actions') || 'Actions', align: 'end' },
]

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
</style>