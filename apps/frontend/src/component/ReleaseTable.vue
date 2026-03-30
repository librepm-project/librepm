<template>
  <v-card
    elevation="0"
    class="rounded-xl"
  >
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

      <template #[`item.status`]="{ item }">
        <v-chip
          :color="item.status === 'Released' ? 'success' : 'warning'"
          text-color="white"
          size="small"
        >
          {{ item.status }}
        </v-chip>
      </template>

      <template #[`item.released_at`]="{ item }">
        {{ item.released_at ? new Date(item.released_at).toLocaleDateString() : '-' }}
      </template>
    </v-data-table>
  </v-card>
</template>

<script setup lang="ts">
import { Release } from '@/lib/interfaces/release.interface';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();

defineProps<{
  items: Release[];
  onEdit: (item: Release) => void;
}>();

const headers = [
  { title: t('admin.release.name'), key: 'name', align: 'start' },
  { title: t('admin.release.status'), key: 'status', align: 'center' },
  { title: t('admin.release.released_at'), key: 'released_at', align: 'center' },
];
</script>

<style scoped>
.color-box {
  width: 20px;
  height: 20px;
  border-radius: 4px;
}
</style>
