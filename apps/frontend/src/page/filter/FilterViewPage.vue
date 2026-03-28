<template>
  <v-container v-if="filter" fluid>
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

    <issue-list
      :filter-id="filterId"
      persist-mode="filter"
      show-settings-toggle
    />
  </v-container>
</template>

<script lang="ts" setup>
import { ref, onMounted, onUnmounted, computed } from 'vue';
import { useRoute } from 'vue-router';
import { useFilterStore } from '@/store/filter.store';
import { useLayoutStore } from '@/store/layout.store';
import { Filter } from '@/lib/interfaces/filter.interface';
import IssueList from '@/component/IssueList.vue';

const route = useRoute();
const filterStore = useFilterStore();
const layoutStore = useLayoutStore();

const filter = ref<Filter | null>(null);
const filterId = computed(() => route.params.filterId as string);

onMounted(async () => {
  filterStore.current = null;
  await filterStore.get(filterId.value);
  filter.value = filterStore.current;

  route.meta.title = filter.value?.name ?? 'Filter';

  layoutStore.setActions([
    {
      text: 'global.edit',
      to: `/filter/${filterId.value}/edit`,
      color: 'primary',
      icon: 'mdi-pencil',
    },
  ]);
});

onUnmounted(() => {
  layoutStore.resetActions();
});
</script>
