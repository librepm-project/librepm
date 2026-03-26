<template>
  <v-container v-if="issueStore.current" class="pa-0">
    <!-- Description Section -->
    <v-card class="rounded-xl pa-6 mb-6">
      <v-card-title class="text-h6 font-weight-bold pa-0 mb-4">
        <v-icon start>mdi-text-box-outline</v-icon>
        {{ t('issue.description') }}
      </v-card-title>
      <v-divider class="mb-4"></v-divider>
      <div class="text-body1 mb-6" v-html="issueStore.current.description || t('global.no_data')"></div>

      <!-- Related Issues Section -->
      <related-issues-panel 
        :relations="relatedIssueStore.relations" 
        :current-issue-id="issueStore.current.id || ''" 
      />
    </v-card>
  </v-container>
</template>

<script lang="ts" setup>
import { useIssueStore } from '@/store/issue.store';
import { useRelatedIssueStore } from '@/store/related-issue.store';
import { useLayoutStore } from '@/store/layout.store';
import { useRoute, useRouter } from 'vue-router';
import { onBeforeMount, onMounted, onUnmounted } from 'vue';
import { useI18n } from 'vue-i18n';
import IssueSidebar from '@/component/IssueSidebar.vue';
import RelatedIssuesPanel from '@/component/RelatedIssuesPanel.vue';


const { t } = useI18n();
const route = useRoute();
const router = useRouter();
const issueStore = useIssueStore();
const relatedIssueStore = useRelatedIssueStore();
const layoutStore = useLayoutStore();

onBeforeMount(async () => {
  await issueStore.getIssue(route.params.issueId.toString());
  await relatedIssueStore.getRelated(route.params.issueId.toString());
  route.meta.title = `${ issueStore.current.key } - ${ issueStore.current.summary }`;
  
  // Set sidebar component
  layoutStore.setSidebarComponent(IssueSidebar, {
    issue: issueStore.current,
  });
});

const remove = async () => {
  if (confirm(t('global.delete_confirm'))) {
    await issueStore.destroy(route.params.issueId.toString());
    router.push('/issue');
  }
}

onMounted(() => {
  layoutStore.setActions([
    {
      text: 'global.edit',
      to: `/issue/${route.params.issueId}/edit`,
      color: 'primary',
      icon: 'mdi-pencil'
    },
    {
      text: 'global.delete',
      onClick: remove,
      color: 'error',
      icon: 'mdi-delete'
    },
  ]);
});

onUnmounted(() => {
  layoutStore.resetActions();
  layoutStore.resetSidebarComponent();
});
</script>