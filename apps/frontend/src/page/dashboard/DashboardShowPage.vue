<template>
      <div class="text-right">
      <v-btn text="Create" to="/dashboard/create" variant="text"></v-btn>
    </div>
</template>

<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue';
import { useDashboardStore } from '@/store/dashboard.store';
import { useSidebarStore } from '@/store/sidebar.store';

import { useI18n } from 'vue-i18n'

const { t } = useI18n();

const dashboardStore = useDashboardStore();
const sidebarStore = useSidebarStore();

onMounted(async () => {
  await dashboardStore.getDashboards();
  sidebarStore.set(dashboardStore.index.map(dashboard => ({
    key: dashboard.id,
    title: dashboard.name,
    link: `/dashboard/${dashboard.id}`,
  })))
});

onUnmounted(() => {
  sidebarStore.reset()
});
</script>