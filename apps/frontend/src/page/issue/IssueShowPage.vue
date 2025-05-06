<template>
  <v-container v-if="issueStore.current">
    <v-row>
      <v-col cols="9">
        <v-card>
          <v-card-title>{{ issueStore.current.summary }}</v-card-title>
          <v-card-subtitle class="mb-2">{{ issueStore.current.id }}</v-card-subtitle>
          <v-card-text class="border rounded pa-2">{{ issueStore.current.description }}</v-card-text>
        </v-card>
      </v-col>
      <v-col cols="3">
        <v-card>
          <v-list dense nav>
            <v-list-item>
              <v-list-item-title class="text-bold">{{ t('issue.project') }}</v-list-item-title>
              <v-list-item-subtitle>{{ issueStore.current.project.name }}</v-list-item-subtitle>
            </v-list-item>
            <v-list-item>
              <v-list-item-title class="text-bold">{{ t('issue.status') }}</v-list-item-title>
              <v-list-item-subtitle><status-chip :status="issueStore.current.status"></status-chip></v-list-item-subtitle>
            </v-list-item>
            <v-list-item>
              <v-list-item-title class="text-bold">{{ t('issue.tracker') }}</v-list-item-title>
              <v-list-item-subtitle><tracker-chip :tracker="issueStore.current.tracker"></tracker-chip></v-list-item-subtitle>
            </v-list-item>
          </v-list>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts" setup>
import { useIssueStore } from '@/store/issue.store';
import { useRoute } from 'vue-router';
import { onBeforeMount } from 'vue';
import { useI18n } from 'vue-i18n';
import StatusChip from '@/component/StatusChip.vue';
import TrackerChip from '@/component/TrackerChip.vue';


const { t } = useI18n();
const route = useRoute();
const issueStore = useIssueStore();

onBeforeMount(async () => {
  await issueStore.getIssue(route.params.issueId.toString());
});
</script>