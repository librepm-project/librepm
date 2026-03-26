<template>
  <v-data-table :items="items" :headers="headers">
    <template #item.key="{ item }">
      <router-link :to="`/issue/${item.id}`" class="key-link">{{ item.key }}</router-link>
    </template>
    <template #item.tracker="{ item }">
      <tracker-chip :tracker="item.tracker" />
    </template>
    <template #item.summary="{ item }">
      <router-link :to="`/issue/${item.id}`">{{ item.summary }}</router-link>
    </template>
    <template #item.status="{ item }">
      <status-chip :status="item.status" />
    </template>
  </v-data-table>
</template>

<script setup lang="ts">
import { Issue } from '@/lib/interfaces/issue.interface';
import { useI18n } from 'vue-i18n';
import TrackerChip from './TrackerChip.vue';
import StatusChip from './StatusChip.vue';

const { t } = useI18n();

defineProps<{ items: Issue[] }>();

const headers = [
  { key: 'key',     title: t('issue.key'),     width: '1px', cellProps: { class: 'text-no-wrap' } },
  { key: 'tracker', title: t('issue.tracker'),  width: '1px', cellProps: { class: 'text-no-wrap' } },
  { key: 'summary', title: t('issue.summary') },
  { key: 'status',  title: t('issue.status'),   width: '1px', align: 'end', cellProps: { class: 'text-no-wrap' } },
];
</script>
