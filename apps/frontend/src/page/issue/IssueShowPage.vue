<template>
  <v-container v-if="issueStore.current" class="pa-6">
    <!-- Main Content -->
    <v-row class="mb-6">
      <v-col cols="12" lg="8">
        <!-- Description Section -->
        <v-card class="rounded-xl pa-6 mb-6">
          <v-card-title class="text-h6 font-weight-bold pa-0 mb-4">
            <v-icon start>mdi-text-box-outline</v-icon>
            {{ t('issue.description') }}
          </v-card-title>
          <v-divider class="mb-4"></v-divider>
          <div class="text-body1" v-html="issueStore.current.description || t('global.no_data')"></div>
        </v-card>
      </v-col>

      <!-- Sidebar: Metadata -->
      <v-col cols="12" lg="4">
        <v-card class="rounded-xl pa-6">
          <v-card-title class="text-h6 font-weight-bold pa-0 mb-4">
            <v-icon start>mdi-information-outline</v-icon>
            {{ t('global.details') }}
          </v-card-title>
          <v-divider class="mb-4"></v-divider>

          <!-- Key -->
          <div class="mb-6">
            <p class="text-subtitle2 font-weight-medium mb-2">
              <v-icon x-small class="mr-1">mdi-tag-outline</v-icon>
              {{ t('issue.key') }}
            </p>
            <v-chip
              :label="issueStore.current.key"
              color="primary"
              variant="outlined"
              class="font-weight-bold"
            ></v-chip>
          </div>

          <!-- Project -->
          <div class="mb-6">
            <p class="text-subtitle2 font-weight-medium mb-2">
              <v-icon x-small class="mr-1">mdi-folder-outline</v-icon>
              {{ t('issue.project') }}
            </p>
            {{issueStore.current.project.name}}
            
          </div>
          <!-- Status -->
          <div class="mb-6">
            <p class="text-subtitle2 font-weight-medium mb-2">
              <v-icon x-small class="mr-1">mdi-circle-outline</v-icon>
              {{ t('issue.status') }}
            </p>
            <status-chip :status="issueStore.current.status"></status-chip>
          </div>

          <!-- Tracker -->
          <div class="mb-6">
            <p class="text-subtitle2 font-weight-medium mb-2">
              <v-icon x-small class="mr-1">mdi-lightning-bolt-outline</v-icon>
              {{ t('issue.tracker') }}
            </p>
            <tracker-chip :tracker="issueStore.current.tracker"></tracker-chip>
          </div>

          <v-divider class="my-4"></v-divider>

          <!-- ID -->
          <p class="text-caption mb-0">
            <strong>ID:</strong> {{ issueStore.current.id }}
          </p>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts" setup>
import { useIssueStore } from '@/store/issue.store';
import { useLayoutStore } from '@/store/layout.store';
import { useRoute, useRouter } from 'vue-router';
import { onBeforeMount, onMounted, onUnmounted, computed } from 'vue';
import { useI18n } from 'vue-i18n';
import StatusChip from '@/component/StatusChip.vue';
import TrackerChip from '@/component/TrackerChip.vue';


const { t } = useI18n();
const route = useRoute();
const router = useRouter();
const issueStore = useIssueStore();
const layoutStore = useLayoutStore();

const projectName = computed(() => {
  return issueStore.current?.project?.name || 'N/A';
});

onBeforeMount(async () => {
  await issueStore.getIssue(route.params.issueId.toString());
  route.meta.title = `${ issueStore.current.key } - ${ issueStore.current.summary }`
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
});
</script>