<template>
  <div>
    <v-card
      elevation="0"
      class="rounded-xl"
      :loading="loading"
    >
      <v-card-text class="pa-6">
        <v-form
          ref="formRef"
          @submit.prevent="handleSubmit"
        >
          <v-text-field
            v-model="form.name"
            :label="t('admin.release.name')"
            :rules="[v => !!v || t('validation.required')]"
            class="mb-4"
            required
          />

          <v-select
            v-model="form.status"
            :items="statusOptions"
            :label="t('admin.release.status')"
            :rules="[v => !!v || t('validation.required')]"
            class="mb-4"
            required
          />

          <v-text-field
            v-if="form.status === 'Released'"
            v-model="form.released_at"
            type="datetime-local"
            :label="t('admin.release.released_at')"
            class="mb-4"
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
import { ref, watch } from 'vue';
import { Release } from '@/lib/interfaces/release.interface';
import { useI18n } from 'vue-i18n';

const { t } = useI18n();

const formRef = ref();

interface Props {
  onSubmit: (release: Release) => Promise<void>;
  submitButtonText?: string;
  initialData?: Release;
}

const props = withDefaults(defineProps<Props>(), {
  submitButtonText: 'global.create',
});

const loading = ref(false);

const form = ref<Release>({
  name: '',
  status: 'Planned',
  released_at: null,
});

const statusOptions = ['Planned', 'Released'];

watch(() => props.initialData, (newVal) => {
  if (newVal) {
    form.value = { ...newVal };
  }
}, { immediate: true, deep: true });

const handleSubmit = async () => {
  const { valid } = await formRef.value.validate();
  if (!valid) return;

  loading.value = true;
  try {
    await props.onSubmit(form.value);
  } finally {
    loading.value = false;
  }
};
</script>
