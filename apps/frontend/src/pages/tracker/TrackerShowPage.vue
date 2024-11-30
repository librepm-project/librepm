<template>
  <div>
    <h1>TrackerDetails</h1>
   
  </div>
</template>

<script lang="ts" setup>
  import Tracker from '../../lib/interfaces/tracker.interface.ts';
  import { useTrackerStore } from '../../store/tracker.store.ts';

  useTrackerStore().getTracker(1);

  const project: Tracker = useTrackerStore().current;

</script>
