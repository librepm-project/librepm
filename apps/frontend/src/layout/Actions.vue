<template>
  <div class="d-flex gap-3 flex-wrap">
    <v-btn
      v-for="button in layoutStore.actionButtons"
      :key="button.text"
      :text="t(button.text)"
      :to="button.to"
      :color="button.color || 'primary'"
      :prepend-icon="button.icon"
      elevation="1"
      rounded="lg"
      @click="button.onClick ? button.onClick() : null"
    />
  </div>
</template>

<script lang="ts" setup>
import { useLayoutStore } from '@/store/layout.store';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();
const layoutStore = useLayoutStore();

</script>

<style scoped>
.gap-3 {
    gap: 1rem;
}

:deep(.v-btn) {
    text-transform: none;
    font-weight: 500;
}
</style>