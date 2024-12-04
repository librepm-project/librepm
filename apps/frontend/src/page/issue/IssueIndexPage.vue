<template>
  <div>
    <issue-table :items="issueStore.index" />
  </div>
</template>

<script lang="ts" setup>
import { useIssueStore } from '@/store/issue.store';
import IssueTable from '@/component/IssueTable.vue';
import { onMounted } from 'vue';

const issueStore = useIssueStore();

onMounted(async () => {
  await issueStore.getIssues();
});
</script>
