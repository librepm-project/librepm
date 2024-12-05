<template>
  <v-container v-if="issueStore.current">
    <v-row>
      <v-col>
        <v-card>
          <v-card-title>{{ issueStore.current.summary }}</v-card-title>
          <v-card-subtitle>{{ issueStore.current.id }}</v-card-subtitle>
          <v-card-text class="text-bold">{{ t('issue.description') }}</v-card-text>
          <v-card-text>{{ issueStore.current.description }}</v-card-text>
        </v-card>
      </v-col>
      <v-col>
        <v-card>
          <v-card-text class="text-bold">{{ t('issue.project') }}</v-card-text>
          <v-card-text>{{ issueStore.current.project.name }}</v-card-text>
          <v-card-text class="text-bold">{{ t('issue.status') }}</v-card-text>
          <v-card-text>{{ issueStore.current.status.name }}</v-card-text>
          <v-card-text class="text-bold">{{ t('issue.tracker') }}</v-card-text>
          <v-card-text>{{ issueStore.current.tracker.name }}</v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>

</template>

<script lang="ts" setup>
import { useIssueStore } from '@/store/issue.store';
import { useRoute } from 'vue-router';
import { onBeforeMount } from 'vue';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();
const route = useRoute();
const issueStore = useIssueStore();

onBeforeMount(async () => {
  await issueStore.getIssue(route.params.issueId);
});
</script>
