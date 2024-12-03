<template>
  <div>
   
  </div>
</template>

<script lang="ts" setup>
  import { useIssueStore } from '@/store/issue.store';

  useIssueStore().getIssue(1);

</script>
