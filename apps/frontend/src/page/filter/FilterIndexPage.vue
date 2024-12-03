<template>
  <div>
    <filter-table :items="filterStore.index" />
  </div>
</template>

<script lang="ts" setup>
  import { useFilterStore } from '@/store/filter.store';
  import FilterTable from '@/component/FilterTable.vue';
  import { onMounted } from 'vue';

  const filterStore = useFilterStore();

  onMounted(async () => {
    await filterStore.getFilters()
  })

</script>