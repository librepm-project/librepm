<template>
  <div class="pa-2">
    <div class="mb-6">
      <p class="text-subtitle2 font-weight-medium mb-2">
        <v-icon x-small class="mr-1">mdi-tag-outline</v-icon>
        {{ t('issue.key') }}
      </p>
      <v-chip
        :label="issue.key"
        color="primary"
        variant="outlined"
        class="font-weight-bold"
      ></v-chip>
    </div>

    <div class="mb-6">
      <p class="text-subtitle2 font-weight-medium mb-2">
        <v-icon x-small class="mr-1">mdi-folder-outline</v-icon>
        {{ t('issue.project') }}
      </p>
      <p class="text-body2 mb-0">
        {{ issue.project?.name }} <span class="text-caption">({{ issue.project?.codeName }})</span>
      </p>
    </div>

    <div class="mb-6">
      <p class="text-subtitle2 font-weight-medium mb-2">
        <v-icon x-small class="mr-1">mdi-circle-outline</v-icon>
        {{ t('issue.status') }}
      </p>
      <status-chip :status="issue.status"></status-chip>
    </div>

    <div class="mb-6">
      <p class="text-subtitle2 font-weight-medium mb-2">
        <v-icon x-small class="mr-1">mdi-lightning-bolt-outline</v-icon>
        {{ t('issue.tracker') }}
      </p>
      <tracker-chip :tracker="issue.tracker"></tracker-chip>
    </div>

    <v-divider class="my-4"></v-divider>

    <p class="text-caption mb-0">
      <strong>ID:</strong> {{ issue.id }}
    </p>
  </div>
</template>

<script lang="ts" setup>
import { useI18n } from 'vue-i18n';
import StatusChip from '@/component/StatusChip.vue';
import TrackerChip from '@/component/TrackerChip.vue';
import { Issue } from '@/lib/interfaces/issue.interface';

const { t } = useI18n();

defineProps({
  issue: {
    type: Object as () => Issue,
    required: true,
  },
});
</script>

<style scoped>
</style>
