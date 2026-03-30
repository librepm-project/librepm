<template>
  <div>
    <release-form
      :on-submit="handleSubmit"
      submit-button-text="global.create"
    />
  </div>
</template>

<script lang="ts" setup>
import { useReleaseStore } from '@/store/release.store';
import { useRouter } from 'vue-router';
import ReleaseForm from '@/component/ReleaseForm.vue';
import { Release } from '@/lib/interfaces/release.interface';

const router = useRouter();
const releaseStore = useReleaseStore();

const handleSubmit = async (release: Release) => {
  try {
    await releaseStore.post(release);
    router.push('/admin/release');
  } catch (error) {
    console.error('Error creating release:', error);
  }
};
</script>
