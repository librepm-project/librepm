<template>
    <v-form @submit.prevent="submit" ref="form" class="form-container">
        <v-card elevation="0" class="rounded-xl pa-6 mb-6">
            <v-card-title class="pa-0 mb-6 text-h6 font-weight-bold">
                {{ props.submitButtonText === 'global.create' ? 'New User' : 'Edit User' }}
            </v-card-title>

            <v-row>
                <v-col cols="12" md="6">
                    <v-text-field
                        v-model="email"
                        :rules="[requiredRule]"
                        :label="$t('user.email')"
                        type="email"
                        outlined
                        dense
                        class="rounded-lg"
                    />
                </v-col>
                <v-col cols="12" md="6">
                    <v-text-field
                        v-model="firstName"
                        :rules="[requiredRule]"
                        :label="$t('user.firstName')"
                        outlined
                        dense
                    />
                </v-col>
            </v-row>

            <v-row>
                <v-col cols="12" md="6">
                    <v-text-field
                        v-model="lastName"
                        :rules="[requiredRule]"
                        :label="$t('user.lastName')"
                        outlined
                        dense
                    />
                </v-col>
                <v-col cols="12" md="6">
                    <v-text-field
                        v-model="phone"
                        :label="$t('user.phone')"
                        outlined
                        dense
                    />
                </v-col>
            </v-row>

            <v-row>
                <v-col cols="12" md="6">
                    <v-text-field
                        v-model="language"
                        :label="$t('user.language')"
                        hint="ISO 639-1 code (e.g., en, hu, de)"
                        outlined
                        dense
                    />
                </v-col>
                <v-col cols="12" md="6">
                    <v-text-field
                        v-model="country"
                        :label="$t('user.country')"
                        hint="ISO 3166-1 code (e.g., US, HU, DE)"
                        outlined
                        dense
                    />
                </v-col>
            </v-row>

            <v-divider class="my-6" />

            <v-row class="mt-4">
                <v-col cols="12" class="d-flex gap-3">
                    <v-btn
                        type="submit"
                        color="primary"
                        size="large"
                        prepend-icon="mdi-check"
                        rounded="lg"
                        class="font-weight-bold"
                    >
                        {{ $t(props.submitButtonText) }}
                    </v-btn>
                    <v-btn
                        type="button"
                        variant="outlined"
                        color="default"
                        size="large"
                        rounded="lg"
                        @click="$router.back()"
                    >
                        Cancel
                    </v-btn>
                </v-col>
            </v-row>
        </v-card>
    </v-form>
</template>

<script lang="ts" setup>
import { ref, watch } from 'vue';
import { requiredRule } from '@/lib/formRule';
import { useRouter } from 'vue-router';

const router = useRouter();

const props = defineProps({
    onSubmit: Function,
    submitButtonText: String,
    user: Object as any,
})

const email = ref("");
const firstName = ref("");
const lastName = ref("");
const phone = ref("");
const language = ref("");
const country = ref("");
const form = ref(null);

watch(() => props.user, (user) => {
    if (user) {
        email.value = user.email || "";
        firstName.value = user.firstName || "";
        lastName.value = user.lastName || "";
        phone.value = user.phone || "";
        language.value = user.language || "";
        country.value = user.country || "";
    }
}, { immediate: true })

const submit = async () => {
    const { valid } = await form.value.validate();
    if (valid) {
        const userData = {
            email: email.value,
            firstName: firstName.value,
            lastName: lastName.value,
            phone: phone.value,
            language: language.value,
            country: country.value,
        };
        props.onSubmit(userData);
    }
};
</script>

<style scoped>
.form-container {
    max-width: 100%;
}

.gap-3 {
    gap: 1rem;
}

:deep(.v-field) {
    background-color: rgba(255, 255, 255, 0.8) !important;
}
</style>
