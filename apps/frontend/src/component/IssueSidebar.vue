<template>
  <div v-if="issueStore.current" class="pa-2">

    <div class="mb-6">
      <p class="text-subtitle2 font-weight-medium mb-2">
        <v-icon x-small class="mr-1">mdi-tag-outline</v-icon>
        {{ t('issue.key') }}
      </p>
      <v-chip color="primary" variant="outlined" class="font-weight-bold">
        {{ issueStore.current.key }}
      </v-chip>
    </div>

    <div class="mb-6">
      <p class="text-subtitle2 font-weight-medium mb-2">
        <v-icon x-small class="mr-1">mdi-folder-outline</v-icon>
        {{ t('issue.project') }}
      </p>
      <p class="text-body2 mb-0">
        {{ issueStore.current.project?.name }}
        <span class="text-caption">({{ issueStore.current.project?.codeName }})</span>
      </p>
    </div>

    <!-- Priority inline edit -->
    <inline-sidebar-edit
      :label="t('issue.priority')"
      icon="mdi-flag-outline"
      :model-value="issueStore.current.priorityId"
      :items="priorities"
      item-title="name"
      item-value="id"
      clearable
      @update:model-value="savePriority"
    >
      <priority-chip v-if="issueStore.current.priority" :priority="issueStore.current.priority" />
      <span v-else class="text-medium-emphasis">—</span>
    </inline-sidebar-edit>

    <!-- Tracker inline edit -->
    <inline-sidebar-edit
      :label="t('issue.tracker')"
      icon="mdi-lightning-bolt-outline"
      :model-value="issueStore.current.tracker?.id"
      :items="trackers"
      item-title="name"
      item-value="id"
      @update:model-value="saveTracker"
    >
      <tracker-chip :tracker="issueStore.current.tracker" />
    </inline-sidebar-edit>

    <!-- Status inline edit -->
    <inline-sidebar-edit
      :label="t('issue.status')"
      icon="mdi-circle-outline"
      :model-value="issueStore.current.status?.id"
      :items="statuses"
      item-title="name"
      item-value="id"
      @update:model-value="saveStatus"
    >
      <status-chip :status="issueStore.current.status" />
    </inline-sidebar-edit>

    <!-- Assignee inline edit -->
    <inline-sidebar-edit
      :label="t('issue.assignee')"
      icon="mdi-account-outline"
      :model-value="issueStore.current.assignedUserId"
      :items="users"
      :item-title="(u: any) => `${u.firstName} ${u.lastName}`"
      item-value="id"
      clearable
      @update:model-value="saveAssignee"
    >
      <span v-if="issueStore.current.assignedUser">
        {{ issueStore.current.assignedUser.firstName }} {{ issueStore.current.assignedUser.lastName }}
      </span>
      <span v-else class="text-medium-emphasis">—</span>
    </inline-sidebar-edit>

    <!-- Reporter inline edit -->
    <inline-sidebar-edit
      :label="t('issue.reporter')"
      icon="mdi-account-edit-outline"
      :model-value="issueStore.current.reporterUserId"
      :items="users"
      :item-title="(u: any) => `${u.firstName} ${u.lastName}`"
      item-value="id"
      clearable
      @update:model-value="saveReporter"
    >
      <span v-if="issueStore.current.reporterUser">
        {{ issueStore.current.reporterUser.firstName }} {{ issueStore.current.reporterUser.lastName }}
      </span>
      <span v-else class="text-medium-emphasis">—</span>
    </inline-sidebar-edit>

    <v-divider class="my-4"></v-divider>
  </div>
</template>

<script lang="ts" setup>
import { ref, watch, onMounted } from 'vue';
import { useI18n } from 'vue-i18n';
import { useIssueStore } from '@/store/issue.store';
import { useProjectStore } from '@/store/project.store';
import { useUserStore } from '@/store/user.store';
import { usePriorityStore } from '@/store/priority.store';
import StatusChip from '@/component/StatusChip.vue';
import TrackerChip from '@/component/TrackerChip.vue';
import PriorityChip from '@/component/PriorityChip.vue';
import InlineSidebarEdit from '@/component/InlineSidebarEdit.vue';

const { t } = useI18n();
const issueStore = useIssueStore();
const projectStore = useProjectStore();
const userStore = useUserStore();
const priorityStore = usePriorityStore();

const trackers = ref<any[]>([]);
const statuses = ref<any[]>([]);
const users = ref<any[]>([]);
const priorities = ref<any[]>([]);

onMounted(async () => {
  const projectId = issueStore.current?.project?.id;
  await Promise.all([
    projectId ? projectStore.getIssueProperty(projectId) : Promise.resolve(),
    userStore.getAll(),
    priorityStore.getAll(),
  ]);
  trackers.value = projectStore.currentIssueProperty?.trackers || [];
  const currentTrackerId = issueStore.current?.tracker?.id;
  const tracker = trackers.value.find(t => t.id === currentTrackerId);
  statuses.value = tracker?.statuses || [];
  users.value = userStore.index;
  priorities.value = priorityStore.index;
});

watch(() => issueStore.current?.tracker?.id, (newTrackerId) => {
  if (!newTrackerId || !projectStore.currentIssueProperty) return;
  const tracker = projectStore.currentIssueProperty.trackers.find(t => t.id === newTrackerId);
  statuses.value = tracker?.statuses || [];
});

const saveTracker = async (trackerId: string) => {
  if (trackerId !== issueStore.current?.tracker?.id) {
    await issueStore.update(issueStore.current!.id!, { trackerId });
  }
};

const saveStatus = async (statusId: string) => {
  if (statusId !== issueStore.current?.status?.id) {
    await issueStore.update(issueStore.current!.id!, { statusId });
  }
};

const saveAssignee = async (assignedUserId: string | null) => {
  if (assignedUserId !== issueStore.current?.assignedUserId) {
    await issueStore.update(issueStore.current!.id!, { assignedUserId });
  }
};

const saveReporter = async (reporterUserId: string | null) => {
  if (reporterUserId !== issueStore.current?.reporterUserId) {
    await issueStore.update(issueStore.current!.id!, { reporterUserId });
  }
};

const savePriority = async (priorityId: string | null) => {
  if (priorityId !== issueStore.current?.priorityId) {
    await issueStore.update(issueStore.current!.id!, { priorityId });
  }
};
</script>

