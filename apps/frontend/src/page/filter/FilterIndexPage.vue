<template>
  <filter-table
    :items="filterStore.index"
    :on-edit="handleEdit"
    :on-view="handleView"
  />
</template>

<script lang="ts" setup>
import { useFilterStore } from '@/store/filter.store';
import FilterTable from '@/component/FilterTable.vue';
import { onMounted, onUnmounted } from 'vue';
import { useLayoutStore } from '@/store/layout.store';
import { useRouter } from 'vue-router';
import { Filter } from '@/lib/interfaces/filter.interface';

const router = useRouter();
const layoutStore = useLayoutStore();
const filterStore = useFilterStore();

const handleEdit = (filter: Filter) => {
  router.push(`/filter/${filter.id}/edit`);
};

const handleView = (filter: Filter) => {
  router.push(`/filter/${filter.id}`); // Navigate to view page
};

onMounted(async () => {
  await filterStore.getAll();
  layoutStore.setActions([
    {
      text: 'global.create',
      to: '/filter/create',
      color: 'primary',
      icon: 'mdi-plus'
    },
  ]);
})

onUnmounted(async () => {
  layoutStore.resetActions();
});

</script>