<template>
    <v-col :cols="props.cols || 12" :md="props.md || (layoutStore.hasSidebar ? 10 : 12)" class="pa-1">
        <v-card elevation="2" class="rounded-xl overflow-hidden">
            <v-card-title class="bg-primary text-white d-flex justify-space-between align-center py-1 px-3">
                <div class="d-flex align-center gap-2 flex-grow-1 mr-4">
                    <v-icon size="small">mdi-file-document</v-icon>
                    <div v-if="!layoutStore.isEditingTitle" 
                         class="text-h6 m-0 cursor-pointer flex-grow-1 font-weight-bold" 
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

            <v-card-text class="pa-3">
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

const props = defineProps<{
    cols?: number | string;
    md?: number | string;
}>();
</script>

<style scoped>
.gap-2 {
    gap: 0.5rem;
}

.title-edit-field :deep(input) {
    color: white !important;
    font-size: 1.25rem;
    font-weight: 500;
}

:deep(.v-card) {
    background: linear-gradient(135deg, #ffffff 0%, #f8f9fa 100%);
}
</style>
