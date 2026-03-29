<template>
  <v-form
    ref="form"
    class="form-container"
    @submit.prevent="submit"
  >
    <v-card
      elevation="0"
      class="rounded-xl pa-6 mb-6"
    >
      <v-card-title class="pa-0 mb-6 text-h6 font-weight-bold">
        {{ props.submitButtonText === 'global.create' ? 'New Priority' : 'Edit Priority' }}
      </v-card-title>

      <v-row>
        <v-col
          cols="12"
          md="8"
        >
          <v-text-field
            v-model="name"
            :rules="[requiredRule]"
            :label="$t('global.name')"
            outlined
            dense
          />
        </v-col>
      </v-row>

      <v-row>
        <v-col cols="12">
          <label class="text-subtitle-2 font-weight-bold mb-3">
            {{ $t('priority.color') }}
          </label>
          <v-color-picker
            v-model="color"
            mode="hex"
            dot-size="30"
            width="300"
            hide-canvas
          />
        </v-col>
      </v-row>

      <v-divider class="my-6" />

      <v-row class="mt-4">
        <v-col
          cols="12"
          class="d-flex gap-3 align-center"
        >
          <v-btn
            type="submit"
            color="primary"
            size="large"
            prepend-icon="mdi-check"
            rounded="lg"
            class="font-weight-bold"
          >
            {{ $t(props.submitButtonText) }}
          </v-btn>
          <v-btn
            v-if="onDelete"
            type="button"
            color="error"
            variant="text"
            size="large"
            prepend-icon="mdi-delete"
            rounded="lg"
            class="font-weight-bold"
            @click="onDelete"
          >
            {{ $t('global.delete') }}
          </v-btn>
          <v-spacer />
          <v-btn
            type="button"
            variant="outlined"
            color="default"
            size="large"
            rounded="lg"
            @click="$router.back()"
          >
            Cancel
          </v-btn>
        </v-col>
      </v-row>
    </v-card>
  </v-form>
</template>

<script lang="ts" setup>
import { ref, watch } from 'vue';
import { requiredRule } from '@/lib/formRule';

const props = defineProps({
    onSubmit: { type: Function, default: undefined },
    onDelete: { type: Function, default: undefined },
    submitButtonText: { type: String, default: undefined },
    priority: { type: Object, default: undefined },
})

const name = ref('');
const color = ref('#FF5722');
const form = ref(null);

watch(() => props.priority, (priority) => {
    if (priority) {
        name.value = priority.name || '';
        color.value = priority.color || '#FF5722';
    }
}, { immediate: true })

const submit = async () => {
    const { valid } = await (form.value as { validate: () => Promise<{ valid: boolean }> }).validate();
    if (valid) {
        props.onSubmit({ name: name.value, color: color.value });
    }
};
</script>

<style scoped>
.form-container {
    max-width: 100%;
}

.gap-3 {
    gap: 1rem;
}

:deep(.v-field) {
    background-color: rgba(255, 255, 255, 0.8) !important;
}
</style>
