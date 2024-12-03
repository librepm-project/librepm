<template>
  <div>
   
  </div>
</template>

<script lang="ts" setup>
  import { useTrackerStore } from '@/store/tracker.store';

  useTrackerStore().getTracker(1);
</script>
