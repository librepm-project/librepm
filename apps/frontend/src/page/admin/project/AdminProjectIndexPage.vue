<template>
  <project-table :items="projectStore.index" />
</template>

<script lang="ts" setup>
import { useProjectStore } from '@/store/project.store';
import ProjectTable from '@/component/ProjectTable.vue';
import { onMounted, onUnmounted } from 'vue';
import { useLayoutStore } from '@/store/layout.store';

const layoutStore = useLayoutStore();
const projectStore = useProjectStore();

onMounted(async () => {
  await projectStore.getProjects();
  layoutStore.setActions([
    {
      text: 'global.create',
      to: '/admin/project/create',
      color: 'primary',
      icon: 'mdi-plus'
    }
  ]);
});

onUnmounted(async () => {
  layoutStore.resetActions();
});
</script>
