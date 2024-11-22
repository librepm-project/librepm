<template>
  <div>
    <h1>Issues</h1>
    <issue-table :items="issues" />
  </div>
</template>

<script lang="ts" setup>
  import Issue from '../../lib/interfaces/issue.interface.ts';
  import { useIssueStore } from '../../store/issue.store.ts';
  import IssueTable from '../../components/IssueTable.vue';

  useIssueStore().getIssues();

  const issues: Issue[] = useIssueStore().index;

</script>
