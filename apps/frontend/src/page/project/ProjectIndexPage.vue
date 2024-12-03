<template>
  <div>
    <project-table :items="projectStore.index" />
  </div>
</template>

<script lang="ts" setup>
import { useProjectStore } from '@/store/project.store';
import ProjectTable from '@/component/ProjectTable.vue';
import { onMounted } from 'vue';

const projectStore = useProjectStore();

onMounted(async () => {
  await projectStore.getProjects();
});
</script>
