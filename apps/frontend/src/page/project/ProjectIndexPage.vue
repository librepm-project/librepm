<template>
  <div>
    <div class="text-right">
      <v-btn :text="t('global.create')" to="/project/create" color="primary" prepend-icon="mdi-plus"></v-btn>
    </div>
    <project-table :items="projectStore.index" />
  </div>
</template>

<script lang="ts" setup>
import { useProjectStore } from '@/store/project.store';
import ProjectTable from '@/component/ProjectTable.vue';
import { onMounted } from 'vue';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();

const projectStore = useProjectStore();

onMounted(async () => {
  await projectStore.getProjects();
});
</script>
