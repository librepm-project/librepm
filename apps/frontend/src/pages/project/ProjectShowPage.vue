<template>
  <div>
    <h1>ProjectDetails</h1>
   
  </div>
</template>

<script lang="ts" setup>
  import { useProjectStore } from '@/store/project.store';

  useProjectStore().getProject(1);

  const project  = useProjectStore().current;

</script>
