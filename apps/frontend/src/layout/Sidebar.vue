<template>
  <v-col
    cols="12"
    md="2"
    class="pe-3"
  >
    <v-sheet
      rounded="xl"
      elevation="2"
      color="surface"
      class="sidebar-sheet"
    >
      <!-- Custom component if set -->
      <component 
        :is="layoutStore.sidebarComponent" 
        v-if="layoutStore.sidebarComponent"
        v-bind="layoutStore.sidebarProps"
      />
            
      <!-- Default navigation items -->
      <v-list
        v-else
        class="pa-1"
      >
        <v-list-item
          v-for="item in layoutStore.sidebarItems"
          :key="item.key"
          :to="item.link"
          rounded="lg"
          class="mb-0"
          active-color="primary"
        >
          <template #prepend>
            <v-icon>{{ item.icon || 'mdi-chevron-right' }}</v-icon>
          </template>
          <v-list-item-title class="font-weight-medium">
            {{ item.title }}
          </v-list-item-title>
        </v-list-item>
      </v-list>
    </v-sheet>
  </v-col>
</template>

<script setup lang="ts">
import { useLayoutStore } from '@/store/layout.store';

const layoutStore = useLayoutStore();
</script>

<style scoped>
.sidebar-sheet {
    background: linear-gradient(135deg, #ffffff 0%, #f8f9fa 100%);
    border: 1px solid rgba(0, 0, 0, 0.05);
}
</style>
