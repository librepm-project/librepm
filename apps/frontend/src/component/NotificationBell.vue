<template>
  <v-menu
    v-model="open"
    :close-on-content-click="false"
    location="bottom end"
    min-width="320"
    max-width="400"
  >
    <template #activator="{ props }">
      <v-btn icon variant="text" class="text-white" v-bind="props" @click="onOpen">
        <v-badge
          v-if="store.items.length"
          :content="store.items.length"
          color="error"
          floating
        >
          <v-icon>mdi-bell-outline</v-icon>
        </v-badge>
        <v-icon v-else>mdi-bell-outline</v-icon>
      </v-btn>
    </template>

    <v-card>
      <v-card-title class="text-body-1 font-weight-bold pa-3 pb-1">
        {{ t('notification.title') }}
      </v-card-title>
      <v-divider />

      <v-list v-if="store.items.length" density="compact" class="pa-1">
        <v-list-item
          v-for="notification in store.items"
          :key="notification.id"
          class="rounded-lg mb-1"
        >
          <template #prepend>
            <v-icon size="small" :color="typeColor(notification.type)" class="mr-2">
              {{ typeIcon(notification.type) }}
            </v-icon>
          </template>
          <v-list-item-title class="text-body-2">
            {{ t(`notification.type.${notification.type}`, notification.type) }}
          </v-list-item-title>
          <v-list-item-subtitle class="text-caption">
            {{ formatDate(notification.createdAt) }}
          </v-list-item-subtitle>
          <template #append>
            <v-btn
              icon
              size="x-small"
              variant="text"
              :title="t('notification.dismiss')"
              @click.stop="store.markAsRead(notification.id)"
            >
              <v-icon size="small">mdi-close</v-icon>
            </v-btn>
          </template>
        </v-list-item>
      </v-list>

      <v-card-text v-else class="text-center text-body-2 text-medium-emphasis py-6">
        {{ t('notification.empty') }}
      </v-card-text>
    </v-card>
  </v-menu>
</template>

<script lang="ts" setup>
import { ref, onMounted } from 'vue';
import { useI18n } from 'vue-i18n';
import { useNotificationStore } from '@/store/notification.store';

const { t } = useI18n();
const store = useNotificationStore();
const open = ref(false);

const onOpen = () => {
  if (!open.value) {
    store.fetchUnread();
  }
};

const typeIcon = (type: string): string => {
  const icons: Record<string, string> = {
    issue_assigned: 'mdi-account-arrow-right',
    issue_created: 'mdi-ticket-outline',
    issue_updated: 'mdi-pencil-outline',
    comment_added: 'mdi-comment-outline',
    mention: 'mdi-at',
    issue_status_changed: 'mdi-swap-horizontal',
  };
  return icons[type] ?? 'mdi-bell-outline';
};

const typeColor = (type: string): string => {
  const colors: Record<string, string> = {
    issue_assigned: 'primary',
    issue_created: 'primary',
    issue_updated: 'info',
    comment_added: 'success',
    mention: 'warning',
    issue_status_changed: 'info',
  };
  return colors[type] ?? 'default';
};

const formatDate = (dateStr: string): string => {
  return new Date(dateStr).toLocaleString();
};

onMounted(() => {
  store.fetchUnread();
});
</script>
