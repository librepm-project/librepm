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
            v-model="form.issue_id"
            :items="issues"
            :item-title="issueTitle"
            item-value="id"
            :label="t('admin.release.issue')"
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
import { ProjectReleaseIssue } from '@/lib/interfaces/project-release-issue.interface';
import { Issue } from '@/lib/interfaces/issue.interface';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();

interface Props {
  onSubmit: (releaseIssue: ProjectReleaseIssue) => Promise<void>;
  submitButtonText?: string;
  projectReleaseId: string;
  issues: Issue[];
}

const props = withDefaults(defineProps<Props>(), {
  submitButtonText: 'global.create',
});

const loading = ref(false);

const form = ref<ProjectReleaseIssue>({
  project_release_id: props.projectReleaseId,
  issue_id: '',
});

const issueTitle = (item: Issue) => `${item.key} - ${item.title}`;

const handleSubmit = async () => {
  loading.value = true;
  try {
    await props.onSubmit(form.value);
  } finally {
    loading.value = false;
  }
};
</script>
