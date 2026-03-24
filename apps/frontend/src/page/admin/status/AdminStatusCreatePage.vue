<template>
  <div>
    <status-form :onSubmit="handleSubmit" submitButtonText="global.create" />
  </div>
</template>

<script lang="ts" setup>
import { useStatusStore } from '@/store/status.store';
import { useRouter } from 'vue-router';
import StatusForm from '@/component/StatusForm.vue';
import { Status } from '@/lib/interfaces/status.interface';

const router = useRouter();
const statusStore = useStatusStore();

const handleSubmit = async (status: Status) => {
  try {
    await statusStore.postStatus(status);
    router.push('/admin/status');
  } catch (error) {
    console.error('Error creating status:', error);
  }
};

</script>