<template>
  <div>
    <h1>Issues</h1>
    <issue-table :items="projectStore.index" />
  </div>
</template>

<script lang="ts" setup>
  import { useIssueStore } from '../../store/issue.store';
  import IssueTable from '../../components/IssueTable.vue';
  import { onMounted } from 'vue';

  const projectStore = useIssueStore()

  onMounted(async () => {
    await projectStore.getIssues();
  });

</script>
