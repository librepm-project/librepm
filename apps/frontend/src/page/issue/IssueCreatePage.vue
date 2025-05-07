<template>
  <div>
    <issue-form :onSubmit="issueStore.create" submitButtonText="global.create" />
  </div>
</template>

<script lang="ts" setup>
import { useIssueStore } from '@/store/issue.store';

import IssueForm from '@/component/IssueForm.vue';

const issueStore = useIssueStore();

</script>