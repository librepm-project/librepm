<template>
  <div>
    <h1>Filters</h1>
    <filter-table :items="filterStore.index" />
  </div>
</template>

<script lang="ts" setup>
  import { useFilterStore } from '../../store/filter.store';
  import FilterTable from '../../components/FilterTable.vue';
  import { onMounted } from 'vue';

  const filterStore = useFilterStore();

  onMounted(async () => {
    await filterStore.getFilters()
  })

</script>