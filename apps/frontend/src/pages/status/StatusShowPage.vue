<template>
  <div>
    <h1>StatusDetails</h1>
   
  </div>
</template>

<script lang="ts" setup>
  import Status from '../../lib/interfaces/status.interface.ts';
  import { useStatusStore } from '../../store/status.store.ts';

  useStatusStore().getStatus(1);

  const project: Status = useStatusStore().current;

</script>
