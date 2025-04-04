<template>
  <div>
    <v-form @submit.prevent="submit" ref="form">
      <v-text-field v-model="email" :rules="[requiredRule, emailRule]" :label="$t('user.email')"></v-text-field>
      <v-text-field v-model="password" type="password" :rules="[requiredRule]"
        :label="$t('user.password')"></v-text-field>
      <v-btn class="mt-2" type="submit" block>{{ $t("login.submit") }}</v-btn>
    </v-form>
  </div>
</template>

<script lang="ts" setup>
import { ref } from "vue";
import { emailRule, requiredRule } from '@/lib/formRule';
import { useUserSessionStore } from '@/store/userSession.store';

const email = ref("");
const password = ref("");

const form = ref(null);
const userSessionStore = useUserSessionStore();

const submit = async () => {
  const { valid } = await form.value.validate();
  if (valid) {
    await userSessionStore.postLogin({ email: email.value, password: password.value });
  }
};
</script>