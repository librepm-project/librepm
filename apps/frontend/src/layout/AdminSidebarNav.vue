<template>
  <v-list class="pa-1">
    <v-list-subheader class="font-weight-bold text-uppercase text-caption px-2">
      Administration
    </v-list-subheader>
    <v-list-item
      v-for="link in adminLinks"
      :key="link.key"
      :to="link.to"
      :prepend-icon="link.icon"
      rounded="lg"
      class="mb-0"
      active-color="primary"
    >
      <v-list-item-title class="font-weight-medium">
        {{ link.text }}
      </v-list-item-title>
    </v-list-item>
  </v-list>
</template>

<script setup lang="ts">
import { navigationLinks } from '@/lib/navigation';

const adminLinks = navigationLinks.find(l => l.key === 'admin')?.sublinks ?? [];
</script>
