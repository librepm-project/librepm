<template>
  <div>
    <h1>TrackerDetails</h1>
   
  </div>
</template>

<script lang="ts" setup>
  import { useTrackerStore } from '@/store/tracker.store';

  useTrackerStore().getTracker(1);

  const project = useTrackerStore().current;

</script>
