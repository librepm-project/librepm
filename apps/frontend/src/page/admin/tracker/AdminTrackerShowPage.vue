<template>
  <tracker-form
    v-if="trackerStore.current"
    :tracker="trackerStore.current"
    :onSubmit="handleSubmit"
    :onDelete="handleDelete"
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
    await trackerStore.put(trackerId, tracker);
    router.push('/admin/tracker');
  } catch (error) {
    console.error('Error updating tracker:', error);
  }
};

const handleDelete = async () => {
  if (confirm('Are you sure you want to delete this tracker?')) {
    try {
      const trackerId = route.params.trackerId as string;
      await trackerStore.delete(trackerId);
      router.push('/admin/tracker');
    } catch (error) {
      console.error('Error deleting tracker:', error);
    }
  }
};

onMounted(async () => {
  const trackerId = route.params.trackerId as string;
  const fromIndex = trackerStore.index.find(t => t.id === trackerId);
  if (fromIndex) {
    trackerStore.current = fromIndex;
  }
  try {
    await trackerStore.get(trackerId);
  } catch (error) {
    console.error('Error loading tracker:', error);
  }
});
</script>
