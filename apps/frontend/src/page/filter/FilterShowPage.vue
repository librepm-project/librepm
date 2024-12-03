<template>
  <div>
   
  </div>
</template>

<script lang="ts" setup>
  import { useFilterStore } from '@/store/filter.store';

  useFilterStore().getFilter(1);

</script>
