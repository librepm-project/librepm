<template>
  <div>
    <user-form
      v-if="userStore.current"
      :user="userStore.current"
      :on-submit="handleSubmit"
      :on-delete="handleDelete"
      submit-button-text="global.update"
    />

    <v-card elevation="0" class="rounded-xl pa-6">
      <v-card-title class="pa-0 mb-4 text-h6 font-weight-bold">
        {{ t('role.title') }}
      </v-card-title>

      <div class="d-flex flex-wrap gap-2 mb-4">
        <template v-if="userRoleStore.roles.length">
          <v-chip
            v-for="role in userRoleStore.roles"
            :key="role"
            closable
            color="primary"
            @click:close="userRoleStore.remove(role)"
          >
            {{ t(`role.options.${role}`, role) }}
          </v-chip>
        </template>
        <span v-else class="text-body-2 text-medium-emphasis">{{ t('role.empty') }}</span>
      </div>

      <v-row align="center">
        <v-col cols="12" sm="6" md="4">
          <v-select
            v-model="selectedRole"
            :items="availableRoles"
            item-title="label"
            item-value="value"
            :label="t('role.add')"
            hide-details
          />
        </v-col>
        <v-col cols="auto">
          <v-btn
            color="primary"
            variant="tonal"
            :disabled="!selectedRole"
            prepend-icon="mdi-plus"
            @click="addRole"
          >
            {{ t('role.add') }}
          </v-btn>
        </v-col>
      </v-row>
    </v-card>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted } from 'vue';
import { useI18n } from 'vue-i18n';
import { useUserStore } from '@/store/user.store';
import { useUserRoleStore } from '@/store/user-role.store';
import { useRoute, useRouter } from 'vue-router';
import UserForm from '@/component/UserForm.vue';
import { User } from '@/lib/interfaces/user.interface';

const { t } = useI18n();
const route = useRoute();
const router = useRouter();
const userStore = useUserStore();
const userRoleStore = useUserRoleStore();

const selectedRole = ref<string | null>(null);

const allRoles = ['admin', 'project_manager', 'developer', 'viewer'];

const availableRoles = computed(() =>
  allRoles
    .filter(r => !userRoleStore.roles.includes(r))
    .map(r => ({ value: r, label: t(`role.options.${r}`, r) }))
);

const addRole = async () => {
  if (!selectedRole.value) return;
  await userRoleStore.add(selectedRole.value);
  selectedRole.value = null;
};

const handleSubmit = async (user: User) => {
  try {
    const userId = route.params.userId as string;
    await userStore.put(userId, user);
    router.push('/admin/user');
  } catch (error) {
    console.error('Error updating user:', error);
  }
};

const handleDelete = async () => {
  if (confirm('Are you sure you want to delete this user?')) {
    try {
      const userId = route.params.userId as string;
      await userStore.delete(userId);
      router.push('/admin/user');
    } catch (error) {
      console.error('Error deleting user:', error);
    }
  }
};

onMounted(async () => {
  const userId = route.params.userId as string;
  await Promise.all([
    userStore.get(userId),
    userRoleStore.fetch(userId),
  ]);
});
</script>
