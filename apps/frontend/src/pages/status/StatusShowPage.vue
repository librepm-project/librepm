<template>
  <div>
    <h1>StatusDetails</h1>
   
  </div>
</template>

<script lang="ts" setup>
  import { useStatusStore } from '@/store/status.store';

  useStatusStore().getStatus(1);

  const status = useStatusStore().current;

</script>
