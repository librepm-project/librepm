<template>
  <div>
    <filter-form :onSubmit="filterStore.postDashboard" submitButtonText="global.create" />
  </div>
</template>

<script lang="ts" setup>
import { useFilterStore } from '@/store/filter.store';

import FilterForm from '@/component/FilterForm.vue';

const filterStore = useFilterStore();

</script>