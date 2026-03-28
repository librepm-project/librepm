<template>
  <issue-list
    :items="issueStore.index"
    persist-mode="global"
    show-settings-toggle
  />
</template>

<script lang="ts" setup>
import { useIssueStore } from '@/store/issue.store';
import { useLayoutStore } from '@/store/layout.store';
import IssueList from '@/component/IssueList.vue';
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
