<template>
    <v-col cols="12" :md="layoutStore.hasSidebar ? 10 : 12">
        <v-card elevation="2" class="rounded-xl overflow-hidden">
            <v-card-title class="bg-primary text-white d-flex justify-space-between align-center">
                <div class="d-flex align-center gap-3 flex-grow-1 mr-4">
                    <v-icon size="large">mdi-file-document</v-icon>
                    <div v-if="!layoutStore.isEditingTitle" 
                         class="text-h5 m-0 cursor-pointer flex-grow-1" 
                         @click="layoutStore.onTitleClick?.()"
                    >
                        {{ layoutStore.title || $route.meta.title }}
                    </div>
                    <v-text-field
                        v-else
                        v-model="layoutStore.titleDraft"
                        variant="underlined"
                        density="compact"
                        autofocus
                        hide-details
                        dark
                        class="title-edit-field"
                        @blur="layoutStore.onTitleSave?.(layoutStore.titleDraft)"
                        @keydown.enter.prevent="layoutStore.onTitleSave?.(layoutStore.titleDraft)"
                        @keydown.escape="layoutStore.onTitleCancel?.()"
                    />
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
import { useLayoutStore } from '@/store/layout.store';

const layoutStore = useLayoutStore();
</script>

<style scoped>
.gap-3 {
    gap: 1rem;
}

.title-edit-field :deep(input) {
    color: white !important;
    font-size: 1.5rem;
    font-weight: 400;
}

:deep(.v-card) {
    background: linear-gradient(135deg, #ffffff 0%, #f8f9fa 100%);
}
</style>
