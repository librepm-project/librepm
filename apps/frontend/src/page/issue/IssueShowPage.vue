<template>
  <v-container v-if="issueStore.current" class="pa-0">
    <!-- Description Section -->
    <div class="text-body1" v-html="issueStore.current.description || t('global.no_data')"></div>
  </v-container>
</template>

<script lang="ts" setup>
import { useIssueStore } from '@/store/issue.store';
import { useLayoutStore } from '@/store/layout.store';
import { useRoute, useRouter } from 'vue-router';
import { onBeforeMount, onMounted, onUnmounted, watch } from 'vue';
import { useI18n } from 'vue-i18n';
import IssueSidebar from '@/component/IssueSidebar.vue';


const { t } = useI18n();
const route = useRoute();
const router = useRouter();
const issueStore = useIssueStore();
const layoutStore = useLayoutStore();

onBeforeMount(async () => {
  await issueStore.getIssue(route.params.issueId.toString());
  route.meta.title = `${ issueStore.current.key } - ${ issueStore.current.summary }`;
  
  // Set sidebar component
  layoutStore.setSidebarComponent(IssueSidebar, {
    issue: issueStore.current,
  });
});

watch(() => issueStore.current, (newIssue) => {
  if (newIssue) {
    layoutStore.setSidebarComponent(IssueSidebar, {
      issue: newIssue,
    });
  }
}, { deep: true });

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
  layoutStore.resetSidebarComponent();
});
</script>