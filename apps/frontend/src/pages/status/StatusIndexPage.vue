<template>
  <div>
    <h1>Statuses</h1>
    <status-table :items="statusStore.index" />
  </div>
</template>

<script lang="ts" setup>
  import { useStatusStore } from '@/store/status.store';
  import StatusTable from '@/components/StatusTable.vue';
  import { onMounted } from 'vue';


  const statusStore = useStatusStore();

  onMounted(async () => {
    await statusStore.getStatuses();
  });

</script>
