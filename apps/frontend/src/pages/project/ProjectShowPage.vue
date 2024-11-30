<template>
  <div>
    <h1>ProjectDetails</h1>
   
  </div>
</template>

<script lang="ts" setup>
  import Project from '../../lib/interfaces/project.interface.ts';
  import { useProjectStore } from '../../store/project.store.ts';

  useProjectStore().getProject(1);

  const project: Project = useProjectStore().current;

</script>
