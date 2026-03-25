<template>
  <status-form
    v-if="statusStore.current"
    :status="statusStore.current"
    :onSubmit="handleSubmit"
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

onMounted(async () => {
  const statusId = route.params.statusId as string;
  await statusStore.getStatus(statusId);
});
</script>
