<template>
  <priority-form
    v-if="priorityStore.current"
    :priority="priorityStore.current"
    :onSubmit="handleSubmit"
    :onDelete="handleDelete"
    submitButtonText="global.update"
  />
</template>

<script lang="ts" setup>
import { usePriorityStore } from '@/store/priority.store';
import { useRoute, useRouter } from 'vue-router';
import PriorityForm from '@/component/PriorityForm.vue';
import { Priority } from '@/lib/interfaces/priority.interface';
import { onMounted } from 'vue';

const route = useRoute();
const router = useRouter();
const priorityStore = usePriorityStore();

const handleSubmit = async (priority: Priority) => {
  try {
    const priorityId = route.params.priorityId as string;
    await priorityStore.put(priorityId, priority);
    router.push('/admin/priority');
  } catch (error) {
    console.error('Error updating priority:', error);
  }
};

const handleDelete = async () => {
  if (confirm('Are you sure you want to delete this priority?')) {
    try {
      const priorityId = route.params.priorityId as string;
      await priorityStore.delete(priorityId);
      router.push('/admin/priority');
    } catch (error) {
      console.error('Error deleting priority:', error);
    }
  }
};

onMounted(async () => {
  const priorityId = route.params.priorityId as string;
  await priorityStore.get(priorityId);
});
</script>
