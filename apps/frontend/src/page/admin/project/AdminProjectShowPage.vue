<template>
  <div>
    <project-form
      v-if="projectStore.current"
      :project="projectStore.current"
      :onSubmit="handleSubmit"
      :onDelete="handleDelete"
      submitButtonText="global.update"
    />
  </div>
</template>

<script lang="ts" setup>
import { useProjectStore } from '@/store/project.store';
import { useRoute, useRouter } from 'vue-router';
import ProjectForm from '@/component/ProjectForm.vue';
import { Project } from '@/lib/interfaces/project.interface';
import { onMounted } from 'vue';

const route = useRoute();
const router = useRouter();
const projectStore = useProjectStore();

const handleSubmit = async (project: Omit<Project, 'id'>) => {
  try {
    const projectId = route.params.projectId as string;
    await projectStore.put(projectId, project);
    router.push('/admin/project');
  } catch (error) {
    console.error('Error updating project:', error);
  }
};

const handleDelete = async () => {
  if (confirm('Are you sure you want to delete this project?')) {
    try {
      const projectId = route.params.projectId as string;
      await projectStore.delete(projectId);
      router.push('/admin/project');
    } catch (error) {
      console.error('Error deleting project:', error);
    }
  }
};

onMounted(async () => {
  const projectId = route.params.projectId as string;
  await projectStore.get(projectId);
});
</script>
