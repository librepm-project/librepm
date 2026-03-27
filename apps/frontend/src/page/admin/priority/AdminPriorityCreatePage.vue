<template>
  <div>
    <priority-form :onSubmit="handleSubmit" submitButtonText="global.create" />
  </div>
</template>

<script lang="ts" setup>
import { usePriorityStore } from '@/store/priority.store';
import { useRouter } from 'vue-router';
import PriorityForm from '@/component/PriorityForm.vue';
import { Priority } from '@/lib/interfaces/priority.interface';

const router = useRouter();
const priorityStore = usePriorityStore();

const handleSubmit = async (priority: Priority) => {
  try {
    await priorityStore.postPriority(priority);
    router.push('/admin/priority');
  } catch (error) {
    console.error('Error creating priority:', error);
  }
};
</script>
