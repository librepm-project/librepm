<template>
  <div class="mb-6">
    <p class="text-subtitle2 font-weight-medium mb-2">
      <v-icon x-small class="mr-1">{{ icon }}</v-icon>
      {{ label }}
    </p>
    <div
      v-if="!editing"
      class="cursor-pointer"
      @click="editing = true"
    >
      <slot />
    </div>
    <div
      v-else
      v-click-outside="{ handler: () => editing = false, include: getOverlayContents }"
    >
      <v-select
        :model-value="modelValue"
        :items="items"
        :item-title="itemTitle"
        :item-value="itemValue"
        :clearable="clearable"
        density="compact"
        variant="underlined"
        hide-details
        autofocus
        @update:model-value="onSave"
      />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';

const props = defineProps<{
  label: string;
  icon: string;
  modelValue: string | null | undefined;
  items: any[];
  itemTitle?: string | ((item: any) => string);
  itemValue?: string;
  clearable?: boolean;
}>();

const emit = defineEmits<{
  'update:modelValue': [value: string | null];
}>();

const editing = ref(false);

const getOverlayContents = () => Array.from(document.querySelectorAll('.v-overlay__content'));

const onSave = (value: string | null) => {
  editing.value = false;
  emit('update:modelValue', value);
};
</script>

<style scoped>
.cursor-pointer {
  cursor: pointer;
}
</style>
