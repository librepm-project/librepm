<template>
  <div>
    <v-card
      elevation="0"
      class="rounded-xl"
      :loading="loading"
    >
      <v-card-text class="pa-6">
        <v-form @submit.prevent="handleSubmit">
          <v-select
            v-if="!props.releaseId"
            v-model="form.release_id"
            :items="releases"
            item-title="name"
            item-value="id"
            :label="t('admin.release.name')"
            :rules="[v => !!v || t('validation.required')]"
            class="mb-4"
            required
          />

          <v-select
            v-model="form.project_id"
            :items="projects"
            item-title="name"
            item-value="id"
            :label="t('admin.project.name')"
            :rules="[v => !!v || t('validation.required')]"
            class="mb-4"
            required
          />

          <v-btn
            type="submit"
            color="primary"
            class="mt-4"
          >
            {{ t(submitButtonText) }}
          </v-btn>
        </v-form>
      </v-card-text>
    </v-card>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { ProjectRelease } from '@/lib/interfaces/project-release.interface';
import { Release } from '@/lib/interfaces/release.interface';
import { Project } from '@/lib/interfaces/project.interface';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();

interface Props {
  onSubmit: (projectRelease: ProjectRelease) => Promise<void>;
  submitButtonText?: string;
  releaseId?: string;
  releases?: Release[];
  projects: Project[];
}

const props = withDefaults(defineProps<Props>(), {
  submitButtonText: 'global.create',
  releaseId: undefined,
  releases: () => [],
});

const loading = ref(false);

const form = ref<ProjectRelease>({
  release_id: props.releaseId ?? '',
  project_id: '',
});

const handleSubmit = async () => {
  loading.value = true;
  try {
    await props.onSubmit(form.value);
  } finally {
    loading.value = false;
  }
};
</script>
