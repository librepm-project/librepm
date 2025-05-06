<template>
  <div>
    <tracker-form :onSubmit="trackerStore.postTracker" submitButtonText="global.create" />
  </div>
</template>

<script lang="ts" setup>
import { useTrackerStore } from '@/store/tracker.store';

import TrackerForm from '@/component/TrackerForm.vue';

const trackerStore = useTrackerStore();

</script>