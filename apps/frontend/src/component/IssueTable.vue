<template>
  <v-data-table class="issue-table" density="compact">
    <thead>
      <tr>
        <th class="text-left key-col">
          {{ t('issue.key') }}
        </th>
        <th class="text-left">
          {{ t('issue.tracker') }}
        </th>
        <th class="text-left summary-col">
          {{ t('issue.summary') }}
        </th>
        <th class="text-left">
          {{ t('issue.status') }}
        </th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="item in items" :key="item.id">
        <td class="key-col">{{ item.key }}</td>
        <td><tracker-chip :tracker="item.tracker" /></td>
        <td class="summary-col">
          <router-link :to="`/issue/${item.id}`">{{ item.summary }}</router-link>
        </td>
        <td><status-chip :status="item.status" /></td>
      </tr>
    </tbody>
  </v-data-table>
</template>

<script setup lang="ts">
import { Issue } from '@/lib/interfaces/issue.interface';
import { useI18n } from 'vue-i18n'
import TrackerChip from './TrackerChip.vue';
import StatusChip from './StatusChip.vue';

const { t } = useI18n();

defineProps<{
  items: Issue[]
}>()

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