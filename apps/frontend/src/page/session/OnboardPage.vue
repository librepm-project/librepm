<template>
  <v-container
    class="fill-height d-flex align-center justify-center"
    fluid
  >
    <v-col
      cols="12"
      sm="9"
      md="6"
      lg="5"
    >
      <v-card
        class="pa-10 rounded-xl mx-auto"
        elevation="4"
        max-width="600"
      >
        <div class="text-center mb-6">
          <v-icon
            size="48"
            color="primary"
            class="mb-3"
          >
            mdi-bug-outline
          </v-icon>
          <h1 class="text-h5 font-weight-bold">
            {{ t('onboard.title') }}
          </h1>
          <p class="text-body-2 text-medium-emphasis mt-1">
            {{ t('onboard.subtitle') }}
          </p>
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

        <v-form
          ref="form"
          @submit.prevent="submit"
        >
          <v-text-field
            v-model="siteTitle"
            :rules="[requiredRule]"
            :label="t('onboard.site_name')"
            prepend-inner-icon="mdi-web"
            variant="outlined"
            density="comfortable"
            class="mb-6"
            min-width="200px"
          />

          <div class="text-subtitle-2 text-medium-emphasis mb-3">
            {{ t('onboard.admin_account') }}
          </div>

          <v-text-field
            v-model="firstName"
            :rules="[requiredRule]"
            :label="t('user.firstName')"
            prepend-inner-icon="mdi-account-outline"
            variant="outlined"
            density="comfortable"
            class="mb-3"
            min-width="200px"
          />
          <v-text-field
            v-model="lastName"
            :rules="[requiredRule]"
            :label="t('user.lastName')"
            prepend-inner-icon="mdi-account-outline"
            variant="outlined"
            density="comfortable"
            class="mb-3"
            min-width="200px"
          />
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
            class="mb-3"
            min-width="200px"
          />
          <v-text-field
            v-model="passwordConfirm"
            type="password"
            :rules="[requiredRule, passwordMatchRule]"
            :label="t('onboard.password_confirm')"
            prepend-inner-icon="mdi-lock-check-outline"
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
            {{ t('onboard.submit') }}
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
import onboardApi from '@/api/onboard.api';
import { useSettingStore } from '@/store/setting.store';

const { t } = useI18n();
const router = useRouter();
const settingStore = useSettingStore();

const siteTitle = ref('');
const firstName = ref('');
const lastName = ref('');
const email = ref('');
const password = ref('');
const passwordConfirm = ref('');
const loading = ref(false);
const error = ref('');
const form = ref<{ validate: () => Promise<{ valid: boolean }> } | null>(null);

const passwordMatchRule = (value: string) =>
  value === password.value || t('onboard.password_mismatch');

const submit = async () => {
  const { valid } = await form.value!.validate();
  if (!valid) return;

  loading.value = true;
  error.value = '';
  try {
    await onboardApi.create({
      siteTitle: siteTitle.value,
      email: email.value,
      password: password.value,
      firstName: firstName.value,
      lastName: lastName.value,
    });
    await settingStore.fetch();
    router.push({ name: 'login' });
  } catch {
    error.value = t('onboard.error');
  } finally {
    loading.value = false;
  }
};
</script>
