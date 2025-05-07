<template>
  <tracker-table :items="trackerStore.index" />
</template>

<script lang="ts" setup>
import { useTrackerStore } from '@/store/tracker.store';
import TrackerTable from '@/component/TrackerTable.vue';
import { onMounted, onUnmounted } from 'vue';
import { useLayoutStore } from '@/store/layout.store';

const layoutStore = useLayoutStore();
const trackerStore = useTrackerStore();

onMounted(async () => {
  await trackerStore.getTrackers();
  layoutStore.setActions([
  {
    text: 'global.create',
    to: '/admin/tracker/create',
    color: 'primary',
    icon: 'mdi-plus'
  }
]);
});

onUnmounted(async () => {
  layoutStore.resetActions();
});

</script>
