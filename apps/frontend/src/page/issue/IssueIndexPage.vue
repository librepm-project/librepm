<template>
  <div>
    <issue-table :items="issueStore.index" />
  </div>
</template>

<script lang="ts" setup>
import { useIssueStore } from '@/store/issue.store';
import { useLayoutStore } from '@/store/layout.store';
import IssueTable from '@/component/IssueTable.vue';
import { onMounted, onUnmounted } from 'vue';

const issueStore = useIssueStore();
const layoutStore = useLayoutStore();

onMounted(async () => {
  await issueStore.getIssues();
  layoutStore.setActions([
    {
      text: 'global.create',
      to: '/issue/create',
      color: 'primary',
      icon: 'mdi-plus',
    },
  ]);
});

onUnmounted(() => {
  layoutStore.resetActions();
});
</script>
