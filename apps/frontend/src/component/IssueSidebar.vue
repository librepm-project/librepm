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
    <div class="mb-6">
      <p class="text-subtitle2 font-weight-medium mb-2">
        <v-icon x-small class="mr-1">mdi-flag-outline</v-icon>
        {{ t('issue.priority') }}
      </p>
      <priority-chip
        v-if="!editingPriority && issueStore.current.priority"
        :priority="issueStore.current.priority"
        class="cursor-pointer"
        @click="editingPriority = true"
      />
      <span
        v-else-if="!editingPriority"
        class="text-medium-emphasis cursor-pointer"
        @click="editingPriority = true"
      >—</span>
      <div
        v-else
        v-click-outside="{ handler: () => editingPriority = false, include: getOverlayContents }"
      >
        <v-select
          :model-value="issueStore.current.priorityId"
          :items="priorities"
          item-title="name"
          item-value="id"
          density="compact"
          variant="underlined"
          hide-details
          clearable
          autofocus
          @update:model-value="savePriority"
        />
      </div>
    </div>

    <!-- Tracker inline edit -->
    <div class="mb-6">
      <p class="text-subtitle2 font-weight-medium mb-2">
        <v-icon x-small class="mr-1">mdi-lightning-bolt-outline</v-icon>
        {{ t('issue.tracker') }}
      </p>
      <tracker-chip
        v-if="!editingTracker"
        :tracker="issueStore.current.tracker"
        class="cursor-pointer"
        @click="editingTracker = true"
      />
      <div
        v-else
        v-click-outside="{ handler: () => editingTracker = false, include: getOverlayContents }"
      >
        <v-select
          :model-value="issueStore.current.tracker?.id"
          :items="trackers"
          item-title="name"
          item-value="id"
          density="compact"
          variant="underlined"
          hide-details
          autofocus
          @update:model-value="saveTracker"
        />
      </div>
    </div>

    <!-- Status inline edit -->
    <div class="mb-6">
      <p class="text-subtitle2 font-weight-medium mb-2">
        <v-icon x-small class="mr-1">mdi-circle-outline</v-icon>
        {{ t('issue.status') }}
      </p>
      <status-chip
        v-if="!editingStatus"
        :status="issueStore.current.status"
        class="cursor-pointer"
        @click="editingStatus = true"
      />
      <div
        v-else
        v-click-outside="{ handler: () => editingStatus = false, include: getOverlayContents }"
      >
        <v-select
          :model-value="issueStore.current.status?.id"
          :items="statuses"
          item-title="name"
          item-value="id"
          density="compact"
          variant="underlined"
          hide-details
          autofocus
          @update:model-value="saveStatus"
        />
      </div>
    </div>

    <!-- Assignee inline edit -->
    <div class="mb-6">
      <p class="text-subtitle2 font-weight-medium mb-2">
        <v-icon x-small class="mr-1">mdi-account-outline</v-icon>
        {{ t('issue.assignee') }}
      </p>
      <div v-if="!editingAssignee" class="cursor-pointer" @click="editingAssignee = true">
        <span v-if="issueStore.current.assignedUser">
          {{ issueStore.current.assignedUser.firstName }} {{ issueStore.current.assignedUser.lastName }}
        </span>
        <span v-else class="text-medium-emphasis">—</span>
      </div>
      <div
        v-else
        v-click-outside="{ handler: () => editingAssignee = false, include: getOverlayContents }"
      >
        <v-select
          :model-value="issueStore.current.assignedUserId"
          :items="users"
          :item-title="(u: any) => `${u.firstName} ${u.lastName}`"
          item-value="id"
          density="compact"
          variant="underlined"
          hide-details
          clearable
          autofocus
          @update:model-value="saveAssignee"
        />
      </div>
    </div>

    <!-- Reporter inline edit -->
    <div class="mb-6">
      <p class="text-subtitle2 font-weight-medium mb-2">
        <v-icon x-small class="mr-1">mdi-account-edit-outline</v-icon>
        {{ t('issue.reporter') }}
      </p>
      <div v-if="!editingReporter" class="cursor-pointer" @click="editingReporter = true">
        <span v-if="issueStore.current.reporterUser">
          {{ issueStore.current.reporterUser.firstName }} {{ issueStore.current.reporterUser.lastName }}
        </span>
        <span v-else class="text-medium-emphasis">—</span>
      </div>
      <div
        v-else
        v-click-outside="{ handler: () => editingReporter = false, include: getOverlayContents }"
      >
        <v-select
          :model-value="issueStore.current.reporterUserId"
          :items="users"
          :item-title="(u: any) => `${u.firstName} ${u.lastName}`"
          item-value="id"
          density="compact"
          variant="underlined"
          hide-details
          clearable
          autofocus
          @update:model-value="saveReporter"
        />
      </div>
    </div>

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

const { t } = useI18n();
const issueStore = useIssueStore();
const projectStore = useProjectStore();
const userStore = useUserStore();
const priorityStore = usePriorityStore();

const editingTracker = ref(false);
const editingStatus = ref(false);
const editingAssignee = ref(false);
const editingReporter = ref(false);
const editingPriority = ref(false);
const trackers = ref<any[]>([]);
const statuses = ref<any[]>([]);
const users = ref<any[]>([]);
const priorities = ref<any[]>([]);

const getOverlayContents = () => Array.from(document.querySelectorAll('.v-overlay__content'));

onMounted(async () => {
  const projectId = issueStore.current?.project?.id;
  await Promise.all([
    projectId ? projectStore.getIssueProperty(projectId) : Promise.resolve(),
    userStore.getAllItems(),
    priorityStore.getPriorities(),
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
  editingTracker.value = false;
  if (trackerId !== issueStore.current?.tracker?.id) {
    await issueStore.update(issueStore.current!.id!, { trackerId });
  }
};

const saveStatus = async (statusId: string) => {
  editingStatus.value = false;
  if (statusId !== issueStore.current?.status?.id) {
    await issueStore.update(issueStore.current!.id!, { statusId });
  }
};

const saveAssignee = async (assignedUserId: string | null) => {
  editingAssignee.value = false;
  if (assignedUserId !== issueStore.current?.assignedUserId) {
    await issueStore.update(issueStore.current!.id!, { assignedUserId });
  }
};

const saveReporter = async (reporterUserId: string | null) => {
  editingReporter.value = false;
  if (reporterUserId !== issueStore.current?.reporterUserId) {
    await issueStore.update(issueStore.current!.id!, { reporterUserId });
  }
};

const savePriority = async (priorityId: string | null) => {
  editingPriority.value = false;
  if (priorityId !== issueStore.current?.priorityId) {
    await issueStore.update(issueStore.current!.id!, { priorityId });
  }
};
</script>

<style scoped>
.cursor-pointer {
  cursor: pointer;
}
</style>
