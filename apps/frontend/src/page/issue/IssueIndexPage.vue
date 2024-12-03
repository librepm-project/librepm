<template>
  <div>
    <issue-table :items="projectStore.index" />
  </div>
</template>

<script lang="ts" setup>
  import { useIssueStore } from '@/store/issue.store';
  import IssueTable from '@/component/IssueTable.vue';
  import { onMounted } from 'vue';

  const projectStore = useIssueStore()

  onMounted(async () => {
    await projectStore.getIssues();
  });

</script>
