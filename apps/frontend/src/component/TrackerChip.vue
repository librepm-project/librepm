<template>
  <v-chip
    :color="tracker.color"
    size="small"
    class="text-capitalize"
  >
    {{ tracker.name }}
  </v-chip>
</template>

<script lang="ts" setup>
import { Tracker } from '@/lib/interfaces/tracker.interface';

defineProps<{
    tracker: Tracker
}>()
</script>