<template>
  <v-card elevation="0" class="rounded-xl">
    <v-data-table
      :items="items"
      :headers="headers"
      density="comfortable"
      class="transparent-table clickable-rows"
      item-key="id"
      @click:row="(e, { item }) => onEdit(item)"
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

      <template #item.name="{ item }">
        {{ item.name }}
      </template>

      <template #item.color="{ item }">
        <div class="d-flex justify-end">
          <div
            class="color-box"
            :style="{ backgroundColor: item.color }"
          ></div>
        </div>
      </template>
    </v-data-table>
  </v-card>
</template>

<script setup lang="ts">
import { Status } from '@/lib/interfaces/status.interface';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();

defineProps<{
  items: Status[];
  onEdit: (item: Status) => void;
}>()

const headers = [
  { key: 'name', title: t('status.name') },
  { key: 'color', title: t('status.color'), align: 'end' },
]

</script>

<style scoped>
.color-box {
  width: 20px;
  height: 20px;
  border-radius: 4px;
  border: 1px solid rgba(0, 0, 0, 0.1);
}

.clickable-rows :deep(tbody tr) {
  cursor: pointer;
}

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