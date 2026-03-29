<template>
  <project-table 
    :items="projectStore.index" 
    :on-edit="handleEdit" 
  />
</template>

<script lang="ts" setup>
import { useProjectStore } from '@/store/project.store';
import ProjectTable from '@/component/ProjectTable.vue';
import { onMounted, onUnmounted } from 'vue';
import { useLayoutStore } from '@/store/layout.store';
import { useRouter } from 'vue-router';
import { Project } from '@/lib/interfaces/project.interface';

const router = useRouter();
const layoutStore = useLayoutStore();
const projectStore = useProjectStore();

const handleEdit = (project: Project) => {
  router.push(`/admin/project/${project.id}`);
};

onMounted(async () => {
  await projectStore.getAll();
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
