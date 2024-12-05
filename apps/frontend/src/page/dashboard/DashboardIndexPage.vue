<template>
  <div>
    <v-select
      :label="t('global.select')"
      :items="dashboardStore.index"
      item-title="name"
      item-value="id"
    ></v-select>
  </div>
</template>

<script setup lang="ts">
import { onMounted } from 'vue';
import { useDashboardStore } from '@/store/dashboard.store';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();

const dashboardStore = useDashboardStore();

onMounted(async () => {
  await dashboardStore.getDashboards();
});

</script>