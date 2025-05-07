<template>
  <status-table :items="statusStore.index" />
</template>

<script lang="ts" setup>
import { useStatusStore } from '@/store/status.store';
import StatusTable from '@/component/StatusTable.vue';
import { onMounted, onUnmounted } from 'vue';
import { useI18n } from 'vue-i18n'
import { useLayoutStore } from '@/store/layout.store';

const { t } = useI18n();

const layoutStore = useLayoutStore();

const statusStore = useStatusStore();

onMounted(async () => {
  await statusStore.getStatuses();
  layoutStore.setActions([
  {
    text: 'global.create',
    to: '/admin/status/create',
    color: 'primary',
    icon: 'mdi-plus'
  }
]);
});

onUnmounted(async () => {
  layoutStore.resetActions();
});
</script>
