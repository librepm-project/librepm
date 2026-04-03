<template>
  <v-container
    class="fill-height"
    fluid
  >
    <v-row
      align="center"
      justify="center"
    >
      <v-col
        cols="12"
        sm="6"
        md="4"
        class="text-center"
      >
        <v-icon
          size="64"
          color="error"
          class="mb-4"
        >
          mdi-lock-outline
        </v-icon>
        <h1 class="text-h5 font-weight-bold mb-2">
          {{ t('forbidden.title') }}
        </h1>
        <p class="text-body-2 text-medium-emphasis mb-6">
          {{ t('forbidden.description') }}
        </p>
        <v-btn
          :to="{ name: 'dashboardShow' }"
          color="primary"
          variant="tonal"
        >
          {{ t('forbidden.back') }}
        </v-btn>
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts" setup>
import { useI18n } from 'vue-i18n';
const { t } = useI18n();
</script>
