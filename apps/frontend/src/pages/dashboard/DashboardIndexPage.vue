<template>
  <div>
    <h1>Dashboards</h1>
    <v-select
      label="Select"
      :items="dashboardStore.index"
      item-title="name"
      item-value="id"
    ></v-select>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { useDashboardStore } from '@/store/dashboard.store';


const dashboardStore = useDashboardStore();

onMounted(async () => {
  await dashboardStore.getDashboards();
});

</script>