<template>
  <tracker-form
    v-if="trackerStore.current"
    :tracker="trackerStore.current"
    :onSubmit="handleSubmit"
    submitButtonText="global.update"
  />
</template>

<script lang="ts" setup>
import { useTrackerStore } from '@/store/tracker.store';
import { useRoute, useRouter } from 'vue-router';
import TrackerForm from '@/component/TrackerForm.vue';
import { Tracker } from '@/lib/interfaces/tracker.interface';
import { onMounted } from 'vue';

const route = useRoute();
const router = useRouter();
const trackerStore = useTrackerStore();

const handleSubmit = async (tracker: Tracker) => {
  try {
    const trackerId = route.params.trackerId as string;
    await trackerStore.putTracker(trackerId, tracker);
    router.push('/admin/tracker');
  } catch (error) {
    console.error('Error updating tracker:', error);
  }
};

onMounted(async () => {
  const trackerId = route.params.trackerId as string;
  await trackerStore.getTracker(trackerId);
});
</script>
