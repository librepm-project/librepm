<template>
  <div>
    <project-form :onSubmit="projectStore.postProject" submitButtonText="global.create" />
  </div>
</template>

<script lang="ts" setup>
import { useProjectStore } from '@/store/project.store';

import ProjectForm from '@/component/ProjectForm.vue';

const projectStore = useProjectStore();

</script>