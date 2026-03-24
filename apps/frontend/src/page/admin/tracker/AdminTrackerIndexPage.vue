<template>
  <tracker-table 
    :items="trackerStore.index" 
    :onEdit="handleEdit" 
    :onDelete="handleDelete"
  />
</template>

<script lang="ts" setup>
import { useTrackerStore } from '@/store/tracker.store';
import TrackerTable from '@/component/TrackerTable.vue';
import { onMounted, onUnmounted } from 'vue';
import { useLayoutStore } from '@/store/layout.store';
import { useRouter } from 'vue-router';
import { Tracker } from '@/lib/interfaces/tracker.interface';

const router = useRouter();
const layoutStore = useLayoutStore();
const trackerStore = useTrackerStore();

const handleEdit = (tracker: Tracker) => {
  router.push(`/admin/tracker/${tracker.id}`);
};

const handleDelete = async (tracker: Tracker) => {
  if (confirm('Are you sure?')) {
    try {
      await trackerStore.deleteTracker(tracker.id);
    } catch (error) {
      console.error('Error deleting tracker:', error);
    }
  }
};

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
