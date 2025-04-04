<template>
    <v-col>
        <v-sheet min-height="80vh" rounded="lg" class="pa-5">
            <v-main>
                <h2 class="text-h2 mb-10">{{ $route.meta.title }}</h2>
                <RouterView />
            </v-main>
        </v-sheet>
    </v-col>
</template>

<script lang="ts" setup>
import { RouterView } from 'vue-router';
</script>