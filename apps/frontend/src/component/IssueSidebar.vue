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

    <v-divider class="my-4"></v-divider>
  </div>
</template>

<script lang="ts" setup>
import { ref, watch, onMounted } from 'vue';
import { useI18n } from 'vue-i18n';
import { useIssueStore } from '@/store/issue.store';
import { useProjectStore } from '@/store/project.store';
import StatusChip from '@/component/StatusChip.vue';
import TrackerChip from '@/component/TrackerChip.vue';

const { t } = useI18n();
const issueStore = useIssueStore();
const projectStore = useProjectStore();

const editingTracker = ref(false);
const editingStatus = ref(false);
const trackers = ref<any[]>([]);
const statuses = ref<any[]>([]);

const getOverlayContents = () => Array.from(document.querySelectorAll('.v-overlay__content'));

onMounted(async () => {
  const projectId = issueStore.current?.project?.id;
  if (!projectId) return;
  await projectStore.getIssueProperty(projectId);
  trackers.value = projectStore.currentIssueProperty?.trackers || [];
  const currentTrackerId = issueStore.current?.tracker?.id;
  const tracker = trackers.value.find(t => t.id === currentTrackerId);
  statuses.value = tracker?.statuses || [];
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
</script>

<style scoped>
.cursor-pointer {
  cursor: pointer;
}
</style>
