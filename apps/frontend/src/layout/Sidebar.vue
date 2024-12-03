<template>
    <v-col cols="2">
        <v-sheet rounded="lg">
            <v-list rounded="lg">
                <v-list-item v-for="item in sidebarStore.items" :key="item.key" :title="item.title" :to="item.link"></v-list-item>
            </v-list>
        </v-sheet>
    </v-col>
</template>


<script setup lang="ts">
import { useSidebarStore } from '@/store/sidebar.store';

const sidebarStore = useSidebarStore();

</script>