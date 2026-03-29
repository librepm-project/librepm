<template>
  <div>
    <project-form
      :on-submit="handleSubmit"
      submit-button-text="global.create"
    />
  </div>
</template>

<script lang="ts" setup>
import { useProjectStore } from '@/store/project.store';
import { useRouter } from 'vue-router';
import ProjectForm from '@/component/ProjectForm.vue';
import { Project } from '@/lib/interfaces/project.interface';

const router = useRouter();
const projectStore = useProjectStore();

const handleSubmit = async (project: Omit<Project, 'id'>) => {
  try {
    await projectStore.post(project);
    router.push('/admin/project');
  } catch (error) {
    console.error('Error creating project:', error);
  }
};

</script>