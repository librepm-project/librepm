<template>
  <status-form
    v-if="statusStore.current"
    :status="statusStore.current"
    :onSubmit="handleSubmit"
    :onDelete="handleDelete"
    submitButtonText="global.update"
  />
</template>

<script lang="ts" setup>
import { useStatusStore } from '@/store/status.store';
import { useRoute, useRouter } from 'vue-router';
import StatusForm from '@/component/StatusForm.vue';
import { Status } from '@/lib/interfaces/status.interface';
import { onMounted } from 'vue';

const route = useRoute();
const router = useRouter();
const statusStore = useStatusStore();

const handleSubmit = async (status: Status) => {
  try {
    const statusId = route.params.statusId as string;
    await statusStore.putStatus(statusId, status);
    router.push('/admin/status');
  } catch (error) {
    console.error('Error updating status:', error);
  }
};

const handleDelete = async () => {
  if (confirm('Are you sure you want to delete this status?')) {
    try {
      const statusId = route.params.statusId as string;
      await statusStore.deleteStatus(statusId);
      router.push('/admin/status');
    } catch (error) {
      console.error('Error deleting status:', error);
    }
  }
};

onMounted(async () => {
  const statusId = route.params.statusId as string;
  await statusStore.getStatus(statusId);
});
</script>
