<template>
  <div>
    <status-table :items="statusStore.index" />
  </div>
</template>

<script lang="ts" setup>
  import { useStatusStore } from '@/store/status.store';
  import StatusTable from '@/component/StatusTable.vue';
  import { onMounted } from 'vue';


  const statusStore = useStatusStore();

  onMounted(async () => {
    await statusStore.getStatuses();
  });

</script>
