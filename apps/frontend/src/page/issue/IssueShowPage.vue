<template>
  <v-container v-if="issueStore.current">
    <v-row>
      <v-col cols="9">
        <v-card>
          <v-card-title>{{ t('issue.description') }}</v-card-title>
          <v-card-text class="border rounded pa-2" v-html="issueStore.current.description"></v-card-text>
        </v-card>
      </v-col>
      <v-col cols="3">
        <v-card>
          <v-list dense nav>
            <v-list-item>
              <v-list-item-title class="text-bold">{{ t('issue.key') }}</v-list-item-title>
              <v-list-item-subtitle>{{ issueStore.current.key }}</v-list-item-subtitle>
            </v-list-item>
            <v-list-item>
              <v-list-item-title class="text-bold">{{ t('issue.project') }}</v-list-item-title>
              <v-list-item-subtitle>{{ issueStore.current.project.name }}</v-list-item-subtitle>
            </v-list-item>
            <v-list-item>
              <v-list-item-title class="text-bold">{{ t('issue.status') }}</v-list-item-title>
              <v-list-item-subtitle><status-chip
                  :status="issueStore.current.status"></status-chip></v-list-item-subtitle>
            </v-list-item>
            <v-list-item>
              <v-list-item-title class="text-bold">{{ t('issue.tracker') }}</v-list-item-title>
              <v-list-item-subtitle><tracker-chip
                  :tracker="issueStore.current.tracker"></tracker-chip></v-list-item-subtitle>
            </v-list-item>
          </v-list>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts" setup>
import { useIssueStore } from '@/store/issue.store';
import { useLayoutStore } from '@/store/layout.store';
import { useRoute, useRouter } from 'vue-router';
import { onBeforeMount, onMounted, onUnmounted } from 'vue';
import { useI18n } from 'vue-i18n';
import StatusChip from '@/component/StatusChip.vue';
import TrackerChip from '@/component/TrackerChip.vue';


const { t } = useI18n();
const route = useRoute();
const router = useRouter();
const issueStore = useIssueStore();
const layoutStore = useLayoutStore();

onBeforeMount(async () => {
  await issueStore.getIssue(route.params.issueId.toString());
  route.meta.title = `${ issueStore.current.key } - ${ issueStore.current.summary }`
});

const remove = async () => {
  if (confirm(t('global.delete_confirm'))) {
    await issueStore.remove(route.params.issueId.toString());
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