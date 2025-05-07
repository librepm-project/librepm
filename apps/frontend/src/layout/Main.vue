<template>
    <v-col>
        <v-sheet rounded="lg" class="pa-5">
            <v-main>
                <div class="d-flex justify-space-between align-center mb-6">
                    <h2 class="text-h4 m-0">{{ $route.meta.title }}</h2>
                    <actions></actions>
                </div>
                <RouterView />
            </v-main>
        </v-sheet>
    </v-col>
</template>

<script lang="ts" setup>
import Actions from '@/layout/Actions.vue';
import { RouterView } from 'vue-router';

</script>