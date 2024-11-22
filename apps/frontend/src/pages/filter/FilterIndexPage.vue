<template>
  <div>
    <h1>Filters</h1>
    <filter-table :items="filters" />
  </div>
</template>

<script lang="ts" setup>
  import Filter from '../../lib/interfaces/filter.interface.ts';
  import { useFilterStore } from '../../store/filter.store.ts';
  import FilterTable from '../../components/FilterTable.vue';

  useFilterStore().getFilters();

  const filters: Filter[] = useFilterStore().index;

</script>