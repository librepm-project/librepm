<template>
  <div>
    <issue-table :items="issueStore.index" :onEdit="onEdit" :onDelete="onDelete" />
  </div>
</template>

<script lang="ts" setup>
import { useIssueStore } from '@/store/issue.store';
import { useLayoutStore } from '@/store/layout.store';
import IssueTable from '@/component/IssueTable.vue';
import { onMounted, onUnmounted } from 'vue';
import { useRouter } from 'vue-router';
import { Issue } from '@/lib/interfaces/issue.interface';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();
const router = useRouter();
const issueStore = useIssueStore();
const layoutStore = useLayoutStore();

const onEdit = (item: Issue) => {
  router.push(`/issue/${item.id}`);
}

const onDelete = async (item: Issue) => {
  if (confirm(t('global.delete_confirm'))) {
    await issueStore.destroy(item.id);
    await issueStore.getIssues();
  }
}

onMounted(async () => {
  await issueStore.getIssues();
  layoutStore.setActions([
    {
      text: 'global.create',
      to: '/issue/create',
      color: 'primary',
      icon: 'mdi-plus'
    },
  ]);
});


onUnmounted(async () => {
  layoutStore.resetActions();
});
</script>
