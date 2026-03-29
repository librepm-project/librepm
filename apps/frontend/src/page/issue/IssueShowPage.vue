<template>
  <v-container
    v-if="issueStore.current"
    class="pa-0"
  >
    <!-- Description label -->
    <p class="text-subtitle2 font-weight-medium mb-3">
      <v-icon
        x-small
        class="mr-1"
      >
        mdi-text-box-outline
      </v-icon>
      {{ t('issue.description') }}
    </p>

    <!-- Description inline edit -->
    <div class="mb-6">
      <!-- eslint-disable vue/no-v-html -->
      <div
        v-if="!editingDescription"
        class="text-body1 cursor-pointer inline-hover pa-2 rounded"
        @click="startEditDescription"
        v-html="issueStore.current.description || t('global.no_data')"
      />
      <!-- eslint-enable vue/no-v-html -->
      <div
        v-else
        v-click-outside="saveDescription"
      >
        <RichTextField
          v-model="descriptionDraft"
          :label="t('issue.description')"
        />
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

    <v-divider class="my-4" />

    <!-- Tabs: Comments | History | Work Log -->
    <v-tabs
      v-model="activeTab"
      density="compact"
      class="mb-2"
    >
      <v-tab value="comments">
        {{ t('comment.comments') }}
      </v-tab>
      <v-tab value="audit">
        {{ t('audit_log.history') }}
      </v-tab>
      <v-tab value="worklogs">
        {{ t('worklog.work_log') }}
      </v-tab>
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
  </v-container>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { wsService } from '@/lib/websocket';
import { Issue } from '@/lib/interfaces/issue.interface';
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

// Summary editing in Layout Header
const startEditSummary = () => {
  if (!issueStore.current) return;
  layoutStore.setTitleEditing(true, issueStore.current.summary, saveSummary, cancelSummary);
};

const saveSummary = async (newSummary: string) => {
  if (!issueStore.current) return;
  layoutStore.setTitleEditing(false);
  if (newSummary && newSummary !== issueStore.current.summary) {
    await issueStore.update(issueStore.current.id, { summary: newSummary });
    layoutStore.setTitle(`${issueStore.current.key} - ${issueStore.current.summary}`, startEditSummary);
  }
};

const cancelSummary = () => {
  layoutStore.setTitleEditing(false);
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

let currentWsChannel: string | null = null;
const onIssueUpdate = (data: unknown) => {
  issueStore.$patch({ current: data as Issue });
};

const loadData = async () => {
  if (currentWsChannel) {
    wsService.unsubscribe(currentWsChannel, onIssueUpdate);
    currentWsChannel = null;
  }

  const key = route.params.key;
  const issueId = route.params.issueId;

  editingDescription.value = false;

  if (key) {
    // Load by key (e.g., /issue/key/PROJ-123)
    await issueStore.getByKey(key.toString());
  } else if (issueId) {
    // Load by ID (backward compatibility)
    await issueStore.get(issueId.toString());
  }

  if (issueStore.current) {
    const currentId = issueStore.current.id;
    await Promise.all([
      relatedIssueStore.get(currentId),
      worklogStore.get(currentId),
      attachmentStore.get(currentId),
      auditLogStore.get(currentId),
      commentStore.get(currentId)
    ]);

    layoutStore.setTitle(`${issueStore.current.key} - ${issueStore.current.summary}`, startEditSummary);
    layoutStore.setSidebarComponent(IssueSidebar, {});

    currentWsChannel = `issue:${currentId}`;
    wsService.subscribe(currentWsChannel, onIssueUpdate);
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
  if (currentWsChannel) {
    wsService.unsubscribe(currentWsChannel, onIssueUpdate);
  }
  layoutStore.resetActions();
  layoutStore.resetSidebarComponent();
  layoutStore.resetTitle();
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
