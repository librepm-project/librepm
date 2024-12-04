<template>
  <v-container>
    <v-row>
      <v-col>
        <v-card>
          <v-card-title>{{ issueStore.current.summary }}</v-card-title>
          <v-card-subtitle>{{ issueStore.current.id }}</v-card-subtitle>
          <v-card-text>{{ issueStore.current.description }}</v-card-text>
        </v-card>
      </v-col>
      <v-col>
        <v-card>
          <v-card-title>Details</v-card-title>
          <v-card-text>{{ issueStore.current.project.name }}</v-card-text>
          <v-card-text>{{ issueStore.current.status.name }}</v-card-text>
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

const route = useRoute();
const issueStore = useIssueStore();

onBeforeMount(async () => {
  await issueStore.getIssue(route.params.issueId);
});
</script>
