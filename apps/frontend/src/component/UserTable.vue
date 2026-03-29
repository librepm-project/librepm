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

      <template #[`item.email`]="{ item }">
        {{ item.email }}
      </template>

      <template #[`item.fullName`]="{ item }">
        <v-chip
          variant="elevated"
          size="small"
        >
          {{ item.firstName }} {{ item.lastName }}
        </v-chip>
      </template>
    </v-data-table>
  </v-card>
</template>

<script setup lang="ts">
import { User } from '@/lib/interfaces/user.interface';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();

defineProps<{
  items: User[];
  onEdit: (item: User) => void;
}>()

const headers = [
  { key: 'email', title: t('user.email') },
  { key: 'firstName', title: t('user.firstName') },
  { key: 'lastName', title: t('user.lastName') },
]

</script>

<style scoped>
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
