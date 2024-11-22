<template>
  <div>
    <h1>IssueDetails</h1>
   
  </div>
</template>

<script lang="ts" setup>
  import Issue from '../../lib/interfaces/issue.interface.ts';
  import { useIssueStore } from '../../store/issue.store.ts';

  useIssueStore().getIssue(1);

  const issue: Issue = useIssueStore().current;

</script>
