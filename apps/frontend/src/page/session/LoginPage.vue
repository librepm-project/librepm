<template>
  <v-container class="fill-height d-flex align-center justify-center" fluid>
    <v-col cols="12" sm="9" md="6" lg="5">
      <v-card class="pa-10 rounded-xl mx-auto" elevation="4" max-width="600">
        <div class="text-center mb-6">
          <v-icon size="48" color="primary" class="mb-3">mdi-bug-outline</v-icon>
          <h1 class="text-h5 font-weight-bold">LibrePM</h1>
        </div>

        <v-alert
          v-if="error"
          type="error"
          variant="tonal"
          density="compact"
          class="mb-4"
        >
          {{ error }}
        </v-alert>

        <v-form @submit.prevent="submit" ref="form">
          <v-text-field
            v-model="email"
            :rules="[requiredRule, emailRule]"
            :label="t('user.email')"
            prepend-inner-icon="mdi-email-outline"
            variant="outlined"
            density="comfortable"
            class="mb-3"
            min-width="200px"
          />
          <v-text-field
            v-model="password"
            type="password"
            :rules="[requiredRule]"
            :label="t('user.password')"
            prepend-inner-icon="mdi-lock-outline"
            variant="outlined"
            density="comfortable"
            class="mb-5"
            min-width="200px"
          />
          <v-btn
            type="submit"
            block
            color="primary"
            size="large"
            :loading="loading"
          >
            {{ t('login.submit') }}
          </v-btn>
        </v-form>
      </v-card>
    </v-col>
  </v-container>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useI18n } from 'vue-i18n';
import { emailRule, requiredRule } from '@/lib/formRule';
import { useUserSessionStore } from '@/store/userSession.store';

const { t } = useI18n();
const router = useRouter();
const userSessionStore = useUserSessionStore();

const email = ref('');
const password = ref('');
const loading = ref(false);
const error = ref('');
const form = ref<{ validate: () => Promise<{ valid: boolean }> } | null>(null);

const submit = async () => {
  const { valid } = await form.value!.validate();
  if (!valid) return;

  loading.value = true;
  error.value = '';
  try {
    await userSessionStore.postLogin({ email: email.value, password: password.value });
    router.push('/dashboard');
  } catch {
    error.value = t('login.invalid_credentials');
  } finally {
    loading.value = false;
  }
};
</script>