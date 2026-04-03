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
            <th class="text-white font-weight-bold">
              {{ t('global.actions') }}
            </th>
          </tr>
        </thead>
      </template>

      <template #[`item.issue.key`]="{ item }">
        <router-link
          :to="`/issue/${item.issue?.id}`"
          class="text-primary"
        >
          {{ item.issue?.key }}
        </router-link>
      </template>

      <template #[`item.issue.title`]="{ item }">
        {{ item.issue?.title }}
      </template>

      <template #[`item.actions`]="{ item }">
        <v-icon
          small
          class="cursor-pointer"
          @click.stop="onDelete(item.id)"
        >
          mdi-delete
        </v-icon>
      </template>
    </v-data-table>
  </v-card>
</template>

<script setup lang="ts">
import { ProjectReleaseIssue } from '@/lib/interfaces/project-release-issue.interface';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();

defineProps<{
  items: ProjectReleaseIssue[];
  onDelete: (id?: string) => void;
}>();

const headers = [
  { title: t('admin.release.issue_key'), key: 'issue.key', align: 'start' },
  { title: t('admin.release.issue_title'), key: 'issue.title', align: 'start' },
];
</script>
