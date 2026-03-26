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

      <v-divider class="my-4"></v-divider>

      <!-- Tabs: Comments | History | Work Log -->
      <v-tabs v-model="activeTab" density="compact" class="mb-2">
        <v-tab value="comments">{{ t('comment.comments') }}</v-tab>
        <v-tab value="audit">{{ t('audit_log.history') }}</v-tab>
        <v-tab value="worklogs">{{ t('worklog.work_log') }}</v-tab>
      </v-tabs>
      <v-window v-model="activeTab">
        <v-window-item value="comments">
          <comment-panel
            :issue-id="issueStore.current.id || ''"
            :comments="commentStore.comments"
          />
        </v-window-item>
        <v-window-item value="audit">
          <issue-audit-log-panel
            :issue-id="issueStore.current.id || ''"
            :audit-logs="auditLogStore.auditLogs"
          />
        </v-window-item>
        <v-window-item value="worklogs">
          <worklog-panel
            :issue-id="issueStore.current.id || ''"
            :worklogs="worklogStore.worklogs"
          />
        </v-window-item>
      </v-window>
    </v-card>
  </v-container>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { useIssueStore } from '@/store/issue.store';
import { useRelatedIssueStore } from '@/store/related-issue.store';
import { useWorklogStore } from '@/store/worklog.store';
import { useAttachmentStore } from '@/store/attachment.store';
import { useIssueAuditLogStore } from '@/store/issue-audit-log.store';
import { useCommentStore } from '@/store/comment.store';
import { useLayoutStore } from '@/store/layout.store';
import { useRoute, useRouter } from 'vue-router';
import { onBeforeMount, onMounted, onUnmounted } from 'vue';
import { useI18n } from 'vue-i18n';
import IssueSidebar from '@/component/IssueSidebar.vue';
import RelatedIssuesPanel from '@/component/RelatedIssuesPanel.vue';
import WorklogPanel from '@/component/WorklogPanel.vue';
import AttachmentPanel from '@/component/AttachmentPanel.vue';
import IssueAuditLogPanel from '@/component/IssueAuditLogPanel.vue';
import CommentPanel from '@/component/CommentPanel.vue';
import RichTextField from '@/component/RichTextField.vue';

const { t } = useI18n();
const route = useRoute();
const router = useRouter();
const issueStore = useIssueStore();
const relatedIssueStore = useRelatedIssueStore();
const worklogStore = useWorklogStore();
const attachmentStore = useAttachmentStore();
const auditLogStore = useIssueAuditLogStore();
const commentStore = useCommentStore();
const layoutStore = useLayoutStore();

const activeTab = ref('comments');

// Summary inline edit
const editingSummary = ref(false);
const summaryDraft = ref('');

const startEditSummary = () => {
  if (!issueStore.current) return;
  summaryDraft.value = issueStore.current.summary;
  editingSummary.value = true;
};

const saveSummary = async () => {
  if (!editingSummary.value || !issueStore.current) return;
  editingSummary.value = false;
  if (summaryDraft.value && summaryDraft.value !== issueStore.current.summary) {
    await issueStore.update(issueStore.current.id, { summary: summaryDraft.value });
    route.meta.title = `${issueStore.current.key} - ${issueStore.current.summary}`;
  }
};

const cancelSummary = () => {
  editingSummary.value = false;
};

// Description inline edit
const editingDescription = ref(false);
const descriptionDraft = ref('');

const startEditDescription = () => {
  if (!issueStore.current) return;
  descriptionDraft.value = issueStore.current.description || '';
  editingDescription.value = true;
};

const saveDescription = async () => {
  if (!editingDescription.value || !issueStore.current) return;
  editingDescription.value = false;
  if (descriptionDraft.value !== issueStore.current.description) {
    await issueStore.update(issueStore.current.id, { description: descriptionDraft.value });
  }
};

const remove = async () => {
  if (!issueStore.current) return;
  if (confirm(t('global.delete_confirm'))) {
    await issueStore.destroy(issueStore.current.id);
    router.push('/issue');
  }
};

const loadData = async () => {
  const key = route.params.key;
  const issueId = route.params.issueId;
  
  editingSummary.value = false;
  editingDescription.value = false;

  if (key) {
    // Load by key (e.g., /issue/key/PROJ-123)
    await issueStore.getIssueByKey(key.toString());
  } else if (issueId) {
    // Load by ID (backward compatibility)
    await issueStore.getIssue(issueId.toString());
  }
  
  if (issueStore.current) {
    const currentId = issueStore.current.id;
    await Promise.all([
      relatedIssueStore.getRelated(currentId),
      worklogStore.getWorklogs(currentId),
      attachmentStore.getAttachments(currentId),
      auditLogStore.getAuditLogs(currentId),
      commentStore.getComments(currentId)
    ]);
    
    route.meta.title = `${issueStore.current.key} - ${issueStore.current.summary}`;
    layoutStore.setSidebarComponent(IssueSidebar, {});
  }
};

onBeforeMount(async () => {
  await loadData();
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

import { watch } from 'vue';
watch(() => route.params, async () => {
  await loadData();
}, { deep: true });

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
