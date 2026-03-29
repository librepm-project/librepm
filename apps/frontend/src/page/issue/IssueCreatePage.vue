<template>
  <div>
    <issue-form
      :on-submit="create"
      submit-button-text="global.create"
    />
  </div>
</template>

<script lang="ts" setup>
import { useIssueStore } from '@/store/issue.store';

import IssueForm from '@/component/IssueForm.vue';
import { Issue } from '@/lib/interfaces/issue.interface';
import { useRouter } from 'vue-router';

const router = useRouter();

const issueStore = useIssueStore();

const create = async (issue: Partial<Issue>) => {
  const record = await issueStore.create(issue);
  console.log("REDIREKT");
  router.push(`/issue/key/${record.key}`);
}

</script>