<template>
  <v-data-table class="issue-table" :items="items" :headers="headers">
    <template #item.summary="{ item }">
      <router-link :to="`/issue/${item.id}`">{{ item.summary }}</router-link>
    </template>
    <template #item.tracker="{ item }">
      <TrackerChip :tracker="item.tracker" />
    </template>
    <template #item.status="{ item }">
      <StatusChip :status="item.status" />
    </template>
  </v-data-table>
</template>

<script setup lang="ts">
import { Issue } from '@/lib/interfaces/issue.interface';
import { useI18n } from 'vue-i18n'
import TrackerChip from './TrackerChip.vue';
import StatusChip from './StatusChip.vue';

const { t } = useI18n();

const props = defineProps<{
  items: Issue[]
}>()

const headers = [
  { key: 'key', title: t('issue.key') },
  { key: 'tracker', title: t('issue.tracker') },
  { key: 'summary', title: t('issue.summary') },
  { key: 'status', title: t('issue.status') },
]

</script>

<style scoped>
.issue-table {
  table-layout: auto;
  width: 100%;
}

.summary-col {
  width: 100%;
}

.key-col {
  white-space: nowrap;
}
</style>