<template>
  <div v-if="issueStore.current">
    <issue-form :onSubmit="update" :initialData="issueStore.current" submitButtonText="global.update" />
  </div>
</template>

<script lang="ts" setup>
import { useIssueStore } from '@/store/issue.store';
import IssueForm from '@/component/IssueForm.vue';
import { Issue } from '@/lib/interfaces/issue.interface';
import { useRouter, useRoute } from 'vue-router';
import { onMounted } from 'vue';

const router = useRouter();
const route = useRoute();
const issueStore = useIssueStore();

onMounted(async () => {
    await issueStore.getIssue(route.params.issueId.toString());
});

const update = async (issue: Partial<Issue>) => {
  await issueStore.update(route.params.issueId.toString(), issue);
  router.push(`/issue/${route.params.issueId}`);
}

</script>
