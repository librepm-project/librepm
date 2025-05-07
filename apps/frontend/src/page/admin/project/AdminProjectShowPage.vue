<template>
  <div>
   
  </div>
</template>

<script lang="ts" setup>
  import { useProjectStore } from '@/store/project.store';

  useProjectStore().getProject(1);

</script>
