<template>
  <project-table 
    :items="projectStore.index" 
    :onEdit="handleEdit" 
    :onDelete="handleDelete"
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

const handleDelete = async (project: Project) => {
  if (confirm('Are you sure?')) {
    try {
      await projectStore.deleteProject(project.id!);
    } catch (error) {
      console.error('Error deleting project:', error);
    }
  }
};

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
