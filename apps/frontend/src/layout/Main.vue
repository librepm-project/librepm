<template>
    <v-col cols="12" md="10" class="px-3">
        <v-card elevation="2" class="rounded-xl overflow-hidden">
            <v-card-title class="bg-primary text-white d-flex justify-space-between align-center">
                <div class="d-flex align-center gap-3">
                    <v-icon size="large">mdi-file-document</v-icon>
                    <h2 class="text-h5 m-0">{{ $route.meta.title }}</h2>
                </div>
                <actions />
            </v-card-title>
            
            <v-divider />
            
            <v-card-text class="pa-6">
                <RouterView />
            </v-card-text>
        </v-card>
    </v-col>
</template>

<script lang="ts" setup>
import Actions from '@/layout/Actions.vue';
import { RouterView } from 'vue-router';

</script>

<style scoped>
.gap-3 {
    gap: 1rem;
}

:deep(.v-card) {
    background: linear-gradient(135deg, #ffffff 0%, #f8f9fa 100%);
}
</style>