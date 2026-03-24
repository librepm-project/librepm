<template>
  <div>
    <tracker-form :onSubmit="handleSubmit" submitButtonText="global.create" />
  </div>
</template>

<script lang="ts" setup>
import { useTrackerStore } from '@/store/tracker.store';
import { useRouter } from 'vue-router';
import TrackerForm from '@/component/TrackerForm.vue';
import { Tracker } from '@/lib/interfaces/tracker.interface';

const router = useRouter();
const trackerStore = useTrackerStore();

const handleSubmit = async (tracker: Tracker) => {
  try {
    await trackerStore.postTracker(tracker);
    router.push('/admin/tracker');
  } catch (error) {
    console.error('Error creating tracker:', error);
  }
};

</script>