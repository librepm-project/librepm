<template>
</template>

<script setup lang="ts">
import { onMounted, onUnmounted } from 'vue';
import { useDashboardStore } from '@/store/dashboard.store';
import { useLayoutStore } from '@/store/layout.store';


import { useI18n } from 'vue-i18n'

const { t } = useI18n();

const dashboardStore = useDashboardStore();
const layoutStore = useLayoutStore();

onMounted(async () => {
  await dashboardStore.getDashboards();
  layoutStore.setSidebar(dashboardStore.index.map(dashboard => ({
    key: dashboard.id,
    title: dashboard.name,
    link: `/dashboard/${dashboard.id}`,
  })));
  layoutStore.setActions([
    {
      text: 'global.create',
      to: '/admin/dashboard/create',
      color: 'primary',
      icon: 'mdi-plus'
    }
  ]);
});

onUnmounted(() => {
  layoutStore.resetSidebar();
  layoutStore.resetActions();

});
</script>