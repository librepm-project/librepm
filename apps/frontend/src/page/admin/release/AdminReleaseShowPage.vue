<template>
  <div class="space-y-6">
    <release-form
      :initial-data="releaseStore.current"
      :on-submit="handleSubmit"
      submit-button-text="global.update"
    />

    <div class="mt-6">
      <h3 class="text-h6 mb-4">
        {{ t('admin.release.projects') }}
      </h3>
      <project-release-table
        :items="projectReleaseStore.byRelease"
        :on-delete="handleDeleteProjectRelease"
      />
    </div>

    <div class="mt-6">
      <h3 class="text-h6 mb-4">
        {{ t('admin.release.add_project') }}
      </h3>
      <project-release-form
        :on-submit="handleAddProjectRelease"
        :release-id="releaseId"
        :projects="projectStore.index"
      />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { useReleaseStore } from '@/store/release.store';
import { useProjectReleaseStore } from '@/store/project-release.store';
import { useProjectStore } from '@/store/project.store';
import { useRouter, useRoute } from 'vue-router';
import ReleaseForm from '@/component/ReleaseForm.vue';
import ProjectReleaseTable from '@/component/ProjectReleaseTable.vue';
import ProjectReleaseForm from '@/component/ProjectReleaseForm.vue';
import { Release } from '@/lib/interfaces/release.interface';
import { ProjectRelease } from '@/lib/interfaces/project-release.interface';
import { useI18n } from 'vue-i18n';
import { onMounted, onUnmounted } from 'vue';
import { useLayoutStore } from '@/store/layout.store';

const { t } = useI18n();
const router = useRouter();
const route = useRoute();
const releaseStore = useReleaseStore();
const projectReleaseStore = useProjectReleaseStore();
const projectStore = useProjectStore();
const layoutStore = useLayoutStore();

const releaseId = route.params.releaseId as string;

const handleSubmit = async (release: Release) => {
  try {
    await releaseStore.put(releaseId, release);
  } catch (error) {
    console.error('Error updating release:', error);
  }
};

const handleAddProjectRelease = async (projectRelease: ProjectRelease) => {
  try {
    await projectReleaseStore.post(projectRelease);
    await projectReleaseStore.getByReleaseId(releaseId);
  } catch (error) {
    console.error('Error adding project to release:', error);
  }
};

const handleDeleteProjectRelease = async (id?: string) => {
  if (!id) return;
  try {
    await projectReleaseStore.delete(id);
    await projectReleaseStore.getByReleaseId(releaseId);
  } catch (error) {
    console.error('Error removing project from release:', error);
  }
};

onMounted(async () => {
  await releaseStore.get(releaseId);
  layoutStore.setTitle(releaseStore.current?.name ?? '');
  await projectReleaseStore.getByReleaseId(releaseId);
  await projectStore.getAll();
});

onUnmounted(() => {
  layoutStore.resetTitle();
});
</script>
