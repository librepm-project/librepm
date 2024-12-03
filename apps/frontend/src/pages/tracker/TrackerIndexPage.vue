<template>
  <div>
    <h1>Trackers</h1>
    <tracker-table :items="trackerStore.index" />
  </div>
</template>

<script lang="ts" setup>
  import { useTrackerStore } from '@/store/tracker.store';
  import TrackerTable from '@/components/TrackerTable.vue';
  import { onMounted } from 'vue';


  const trackerStore = useTrackerStore();

  onMounted(async () => {
    await trackerStore.getTrackers();
  });

</script>
