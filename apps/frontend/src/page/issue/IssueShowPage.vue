<template>
  <v-container v-if="issueStore.current" class="pa-0">
    <v-card class="rounded-xl pa-6 mb-6">

      <!-- Summary inline edit -->
      <div class="mb-4">
        <div
          v-if="!editingSummary"
          class="text-h6 font-weight-bold cursor-pointer inline-hover pa-1 rounded"
          @click="startEditSummary"
        >
          {{ issueStore.current.summary }}
        </div>
        <v-text-field
          v-else
          v-model="summaryDraft"
          variant="underlined"
          density="compact"
          autofocus
          hide-details
          @blur="saveSummary"
          @keydown.enter.prevent="saveSummary"
          @keydown.escape="cancelSummary"
        />
      </div>

      <v-divider class="mb-4"></v-divider>

      <!-- Description label -->
      <p class="text-subtitle2 font-weight-medium mb-3">
        <v-icon x-small class="mr-1">mdi-text-box-outline</v-icon>
        {{ t('issue.description') }}
      </p>

      <!-- Description inline edit -->
      <div class="mb-6">
        <div
          v-if="!editingDescription"
          class="text-body1 cursor-pointer inline-hover pa-2 rounded"
          @click="startEditDescription"
          v-html="issueStore.current.description || t('global.no_data')"
        />
        <div
          v-else
          v-click-outside="saveDescription"
        >
          <RichTextField v-model="descriptionDraft" :label="t('issue.description')" />
        </div>
      </div>

      <!-- Attachments Section -->
      <attachment-panel
        :issue-id="issueStore.current.id || ''"
        :attachments="attachmentStore.attachments"
      />

      <!-- Related Issues Section -->
      <related-issues-panel
        :relations="relatedIssueStore.relations"
        :current-issue-id="issueStore.current.id || ''"
      />

      <!-- Worklog Section -->
      <worklog-panel
        :issue-id="issueStore.current.id || ''"
        :worklogs="worklogStore.worklogs"
      />
    </v-card>
  </v-container>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { useIssueStore } from '@/store/issue.store';
import { useRelatedIssueStore } from '@/store/related-issue.store';
import { useWorklogStore } from '@/store/worklog.store';
import { useAttachmentStore } from '@/store/attachment.store';
import { useLayoutStore } from '@/store/layout.store';
import { useRoute, useRouter } from 'vue-router';
import { onBeforeMount, onMounted, onUnmounted } from 'vue';
import { useI18n } from 'vue-i18n';
import IssueSidebar from '@/component/IssueSidebar.vue';
import RelatedIssuesPanel from '@/component/RelatedIssuesPanel.vue';
import WorklogPanel from '@/component/WorklogPanel.vue';
import AttachmentPanel from '@/component/AttachmentPanel.vue';
import RichTextField from '@/component/RichTextField.vue';

const { t } = useI18n();
const route = useRoute();
const router = useRouter();
const issueStore = useIssueStore();
const relatedIssueStore = useRelatedIssueStore();
const worklogStore = useWorklogStore();
const attachmentStore = useAttachmentStore();
const layoutStore = useLayoutStore();

// Summary inline edit
const editingSummary = ref(false);
const summaryDraft = ref('');

const startEditSummary = () => {
  summaryDraft.value = issueStore.current!.summary;
  editingSummary.value = true;
};

const saveSummary = async () => {
  if (!editingSummary.value) return;
  editingSummary.value = false;
  if (summaryDraft.value && summaryDraft.value !== issueStore.current!.summary) {
    await issueStore.update(issueStore.current!.id!, { summary: summaryDraft.value });
    route.meta.title = `${issueStore.current!.key} - ${issueStore.current!.summary}`;
  }
};

const cancelSummary = () => {
  editingSummary.value = false;
};

// Description inline edit
const editingDescription = ref(false);
const descriptionDraft = ref('');

const startEditDescription = () => {
  descriptionDraft.value = issueStore.current!.description || '';
  editingDescription.value = true;
};

const saveDescription = async () => {
  if (!editingDescription.value) return;
  editingDescription.value = false;
  if (descriptionDraft.value !== issueStore.current!.description) {
    await issueStore.update(issueStore.current!.id!, { description: descriptionDraft.value });
  }
};

const remove = async () => {
  if (confirm(t('global.delete_confirm'))) {
    await issueStore.destroy(route.params.issueId.toString());
    router.push('/issue');
  }
};

onBeforeMount(async () => {
  await issueStore.getIssue(route.params.issueId.toString());
  await relatedIssueStore.getRelated(route.params.issueId.toString());
  await worklogStore.getWorklogs(route.params.issueId.toString());
  await attachmentStore.getAttachments(route.params.issueId.toString());
  route.meta.title = `${issueStore.current!.key} - ${issueStore.current!.summary}`;
  layoutStore.setSidebarComponent(IssueSidebar, {});
});

onMounted(() => {
  layoutStore.setActions([
    {
      text: 'global.delete',
      onClick: remove,
      color: 'error',
      icon: 'mdi-delete',
    },
  ]);
});

onUnmounted(() => {
  layoutStore.resetActions();
  layoutStore.resetSidebarComponent();
});
</script>

<style scoped>
.inline-hover:hover {
  background-color: rgba(0, 0, 0, 0.04);
}

.cursor-pointer {
  cursor: pointer;
}
</style>
