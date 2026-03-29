<template>
  <div class="text-subtitle-1 mb-2">
    {{ label }}
  </div>
  <v-sheet elevation="0">
    <quill-editor
      v-model:content="localContent"
      content-type="html"
      theme="snow"
      style="min-height: 200px"
    />
  </v-sheet>
</template>

<script lang="ts" setup>
import { watch, ref } from 'vue';
import { QuillEditor } from '@vueup/vue-quill';
import '@vueup/vue-quill/dist/vue-quill.snow.css';

const props = defineProps<{
  label: string
  modelValue: string
}>()

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void
}>()

const localContent = ref(props.modelValue)

watch(localContent, (val) => {
  emit('update:modelValue', val)
})

watch(
  () => props.modelValue,
  (val) => {
    if (val !== localContent.value) {
      localContent.value = val
    }
  }
)
</script>