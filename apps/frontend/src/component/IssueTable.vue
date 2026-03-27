<template>
  <div>
    <div v-if="groupedIssues.length > 0" class="grouped-issues">
      <div v-for="group in groupedIssues" :key="group.key" class="mb-6">
        <h3 class="text-h6 mb-4 pa-3 rounded" style="background: rgba(63, 81, 181, 0.08)">
          {{ group.label }}
        </h3>
        <v-data-table :items="group.issues" :headers="computedHeaders">
          <template #item.key="{ item }">
            <router-link :to="`/issue/key/${item.key}`" class="key-link">{{ item.key }}</router-link>
          </template>
          <template #item.tracker="{ item }">
            <tracker-chip :tracker="item.tracker" />
          </template>
          <template #item.priority="{ item }">
            <priority-chip v-if="item.priority" :priority="item.priority" />
            <span v-else class="text-medium-emphasis">—</span>
          </template>
          <template #item.summary="{ item }">
            <router-link :to="`/issue/key/${item.key}`">{{ item.summary }}</router-link>
          </template>
          <template #item.status="{ item }">
            <status-chip :status="item.status" />
          </template>
          <template #item.assignee="{ item }">
            <span v-if="item.assignedUser">{{ item.assignedUser.firstName }} {{ item.assignedUser.lastName }}</span>
            <span v-else class="text-medium-emphasis">—</span>
          </template>
          <template #item.project="{ item }">
            <span v-if="item.project">{{ item.project.name }}</span>
            <span v-else class="text-medium-emphasis">—</span>
          </template>
        </v-data-table>
      </div>
    </div>
    <v-data-table v-else :items="items" :headers="computedHeaders">
      <template #item.key="{ item }">
        <router-link :to="`/issue/key/${item.key}`" class="key-link">{{ item.key }}</router-link>
      </template>
      <template #item.tracker="{ item }">
        <tracker-chip :tracker="item.tracker" />
      </template>
      <template #item.priority="{ item }">
        <priority-chip v-if="item.priority" :priority="item.priority" />
        <span v-else class="text-medium-emphasis">—</span>
      </template>
      <template #item.summary="{ item }">
        <router-link :to="`/issue/key/${item.key}`">{{ item.summary }}</router-link>
      </template>
      <template #item.status="{ item }">
        <status-chip :status="item.status" />
      </template>
      <template #item.assignee="{ item }">
        <span v-if="item.assignedUser">{{ item.assignedUser.firstName }} {{ item.assignedUser.lastName }}</span>
        <span v-else class="text-medium-emphasis">—</span>
      </template>
      <template #item.project="{ item }">
        <span v-if="item.project">{{ item.project.name }}</span>
        <span v-else class="text-medium-emphasis">—</span>
      </template>
    </v-data-table>
  </div>
</template>

<script setup lang="ts">
import { Issue } from '@/lib/interfaces/issue.interface';
import { useI18n } from 'vue-i18n';
import { computed } from 'vue';
import TrackerChip from './TrackerChip.vue';
import StatusChip from './StatusChip.vue';
import PriorityChip from './PriorityChip.vue';

const { t } = useI18n();

interface Props {
  items: Issue[];
  columns?: string[];
  groupBy?: string;
}

const props = withDefaults(defineProps<Props>(), {
  columns: () => ['key', 'tracker', 'priority', 'summary', 'assignee', 'status'],
  groupBy: '',
});

const allHeaders = {
  key: { key: 'key', title: t('issue.key'), width: '1px', cellProps: { class: 'text-no-wrap' } },
  tracker: { key: 'tracker', title: t('issue.tracker'), width: '1px', cellProps: { class: 'text-no-wrap' } },
  priority: { key: 'priority', title: t('issue.priority'), width: '1px', cellProps: { class: 'text-no-wrap' } },
  summary: { key: 'summary', title: t('issue.summary') },
  assignee: { key: 'assignee', title: t('issue.assignee'), width: '1px', cellProps: { class: 'text-no-wrap' } },
  status: { key: 'status', title: t('issue.status'), width: '1px', align: 'end', cellProps: { class: 'text-no-wrap' } },
  project: { key: 'project', title: t('issue.project'), width: '1px', cellProps: { class: 'text-no-wrap' } },
};

const computedHeaders = computed(() => {
  return props.columns
    .map((col: string) => (allHeaders as any)[col])
    .filter(Boolean);
});

const getGroupKey = (issue: Issue): string => {
  if (!props.groupBy) return '';
  
  switch (props.groupBy) {
    case 'status':
      return issue.status?.name || 'Unassigned';
    case 'priority':
      return issue.priority?.name || 'No Priority';
    case 'assignee':
      return issue.assignedUser
        ? `${issue.assignedUser.firstName} ${issue.assignedUser.lastName}`
        : 'Unassigned';
    case 'project':
      return issue.project?.name || 'No Project';
    default:
      return '';
  }
};

const groupedIssues = computed(() => {
  if (!props.groupBy) return [];

  const groups = new Map<string, Issue[]>();
  
  for (const issue of props.items) {
    const key = getGroupKey(issue);
    if (!groups.has(key)) {
      groups.set(key, []);
    }
    groups.get(key)!.push(issue);
  }

  return Array.from(groups.entries()).map(([key, issues]) => ({
    key,
    label: key,
    issues,
  }));
});
</script>
