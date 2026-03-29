<template>
  <v-col>
    <p class="text-center">
      LibrePM 2024 - {{ currentYear }}
    </p>
  </v-col>
</template>

<script setup lang="ts">
const currentYear = new Date().getFullYear();
</script>
