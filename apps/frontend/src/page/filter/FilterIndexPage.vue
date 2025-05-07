<template>
  <filter-table :items="filterStore.index" />
</template>

<script lang="ts" setup>
import { useFilterStore } from '@/store/filter.store';
import FilterTable from '@/component/FilterTable.vue';
import { onMounted, onUnmounted } from 'vue';
import { useLayoutStore } from '@/store/layout.store';

const layoutStore = useLayoutStore();
const filterStore = useFilterStore();

onMounted(async () => {
  await filterStore.getFilters();
  layoutStore.setActions([
    {
      text: 'global.create',
      to: '/admin/filter/create',
      color: 'primary',
      icon: 'mdi-plus'
    },
  ]);
})

onUnmounted(async () => {
  layoutStore.resetActions();
});

</script>