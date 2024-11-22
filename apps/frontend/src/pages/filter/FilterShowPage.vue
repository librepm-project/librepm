<template>
  <div>
    <h1>Filter details</h1>
   
  </div>
</template>

<script lang="ts" setup>
  import Filter from '../../lib/interfaces/filter.interface.ts';
  import { useFilterStore } from '../../store/filter.store.ts';

  useFilterStore().getFilter(1);

  const filter: Filter = useFilterStore().current;

</script>
