<template>
  <v-chip
    :color="priority.color"
    size="small"
    class="text-capitalize"
  >
    {{ priority.name }}
  </v-chip>
</template>

<script lang="ts" setup>
import { Priority } from '@/lib/interfaces/priority.interface';

defineProps<{
    priority: Priority
}>()
</script>
