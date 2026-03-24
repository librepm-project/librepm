<template>
    <v-btn-group v-if="layoutStore.actionButtons.length > 0">
        <v-btn v-for="button in layoutStore.actionButtons" :key="button.text" :text="t(button.text)" :to="button.to"
            :color="button.color" :prepend-icon="button.icon" class="ml-2" @click="button.onClick ? button.onClick() : null" />
    </v-btn-group>
</template>

<script lang="ts" setup>
import { useLayoutStore } from '@/store/layout.store';
import { useI18n } from 'vue-i18n'

const { t } = useI18n();

const layoutStore = useLayoutStore();

</script>