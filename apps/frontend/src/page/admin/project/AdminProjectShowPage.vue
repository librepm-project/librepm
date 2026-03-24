<template>
  <div>
    <project-form 
      :project="projectStore.current" 
      :onSubmit="handleSubmit" 
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
    const projectId = route.params.projectid as string;
    await projectStore.putProject(projectId, project);
    router.push('/admin/project');
  } catch (error) {
    console.error('Error updating project:', error);
  }
};

onMounted(async () => {
  const projectId = route.params.projectid as string;
  await projectStore.getProject(projectId);
});
</script>
