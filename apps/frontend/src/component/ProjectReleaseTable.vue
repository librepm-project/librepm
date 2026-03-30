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
            <th class="text-white font-weight-bold">{{ t('global.actions') }}</th>
          </tr>
        </thead>
      </template>

      <template #[`item.release.name`]="{ item }">
        {{ item.release?.name }}
      </template>

      <template #[`item.project.name`]="{ item }">
        {{ item.project?.name }}
      </template>

      <template #[`item.actions`]="{ item }">
        <v-icon
          small
          class="mr-2 cursor-pointer"
          @click.stop="onDelete(item.id)"
        >
          mdi-delete
        </v-icon>
      </template>
    </v-data-table>
  </v-card>
</template>

<script setup lang="ts">
import { ProjectRelease } from '@/lib/interfaces/project-release.interface';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();

defineProps<{
  items: ProjectRelease[];
  onDelete: (id?: string) => void;
}>();

const headers = [
  { title: t('admin.release.name'), key: 'release.name', align: 'start' },
  { title: t('admin.project.name'), key: 'project.name', align: 'start' },
];
</script>
