<template>
  <release-table 
    :items="releaseStore.index" 
    :on-edit="handleEdit" 
  />
</template>

<script lang="ts" setup>
import { useReleaseStore } from '@/store/release.store';
import ReleaseTable from '@/component/ReleaseTable.vue';
import { onMounted, onUnmounted } from 'vue';
import { useLayoutStore } from '@/store/layout.store';
import { useRouter } from 'vue-router';
import { Release } from '@/lib/interfaces/release.interface';

const router = useRouter();
const layoutStore = useLayoutStore();
const releaseStore = useReleaseStore();

const handleEdit = (release: Release) => {
  router.push(`/admin/release/${release.id}`);
};

onMounted(async () => {
  await releaseStore.getAll();
  layoutStore.setActions([
    {
      text: 'global.create',
      to: '/admin/release/create',
      color: 'primary',
      icon: 'mdi-plus'
    }
  ]);
});

onUnmounted(async () => {
  layoutStore.resetActions();
});
</script>
