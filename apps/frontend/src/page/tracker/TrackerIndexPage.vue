<template>
  <div>
    <div class="text-right">
      <v-btn text="Create" to="/tracker/create" variant="text"></v-btn>
    </div>
    <tracker-table :items="trackerStore.index" />
  </div>
</template>

<script lang="ts" setup>
  import { useTrackerStore } from '@/store/tracker.store';
  import TrackerTable from '@/component/TrackerTable.vue';
  import { onMounted } from 'vue';


  const trackerStore = useTrackerStore();

  onMounted(async () => {
    await trackerStore.getTrackers();
  });

</script>
