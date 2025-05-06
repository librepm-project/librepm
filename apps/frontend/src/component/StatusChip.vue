<template>
    <v-chip :color="status.color" size="small" class="text-capitalize">
        {{ status.name }}
    </v-chip>
</template>

<script lang="ts" setup>
import { Status } from '@/lib/interfaces/status.interface';

defineProps<{
    status: Status
}>()
</script>