<template>
  <v-container v-if="filter">
    <v-row class="mb-4">
      <v-col>
        <p v-if="filter.description" class="text-body-1 text-medium-emphasis">
          {{ filter.description }}
        </p>

        <div v-if="filter.conditions?.length" class="d-flex flex-wrap gap-2 mt-3">
          <v-chip
            v-for="(condition, i) in filter.conditions"
            :key="i"
            size="small"
            variant="tonal"
            color="primary"
          >
            {{ condition.field }} {{ condition.op }} {{ condition.value }}
          </v-chip>
        </div>
      </v-col>
    </v-row>

    <v-divider class="mb-4" />

    <div v-if="loading" class="d-flex justify-center py-8">
      <v-progress-circular indeterminate color="primary" />
    </div>

    <issue-table
      v-else
      :items="issues"
      :columns="filterColumns"
      :groupBy="filterGroupBy"
    />
  </v-container>
</template>

<script lang="ts" setup>
import { ref, onMounted, onUnmounted, computed } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { useFilterStore } from '@/store/filter.store';
import { useLayoutStore } from '@/store/layout.store';
import { Issue } from '@/lib/interfaces/issue.interface';
import { Filter } from '@/lib/interfaces/filter.interface';
import issueApi from '@/api/issue.api';
import IssueTable from '@/component/IssueTable.vue';

const route = useRoute();
const router = useRouter();
const filterStore = useFilterStore();
const layoutStore = useLayoutStore();

const filter = ref<Filter | null>(null);
const issues = ref<Issue[]>([]);
const loading = ref(true);

const filterColumns = computed(() => {
  if (!filter.value?.columnList) {
    return ['key', 'tracker', 'priority', 'summary', 'assignee', 'status'];
  }
  try {
    return JSON.parse(filter.value.columnList);
  } catch {
    return ['key', 'tracker', 'priority', 'summary', 'assignee', 'status'];
  }
});

const filterGroupBy = computed(() => {
  return filter.value?.groupBy || '';
});

onMounted(async () => {
  const filterId = route.params.filterId as string;

  filterStore.current = null;
  await filterStore.getFilter(filterId);
  filter.value = filterStore.current;

  route.meta.title = filter.value?.name ?? 'Filter';

  layoutStore.setActions([
    {
      text: 'global.edit',
      to: `/filter/${filterId}/edit`,
      color: 'primary',
      icon: 'mdi-pencil',
    },
  ]);

  try {
    issues.value = await issueApi.index({ filterId });
  } finally {
    loading.value = false;
  }
});

onUnmounted(() => {
  layoutStore.resetActions();
});
</script>
