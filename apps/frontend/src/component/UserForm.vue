<template>
    <v-form @submit.prevent="submit" ref="form">
        <v-text-field v-model="email" :rules="[requiredRule]" :label="$t('user.email')"></v-text-field>
        <v-text-field v-model="firstName" :rules="[requiredRule]" :label="$t('user.firstName')"></v-text-field>
        <v-text-field v-model="lastName" :rules="[requiredRule]" :label="$t('user.lastName')"></v-text-field>
        <v-text-field v-model="phone" :label="$t('user.phone')"></v-text-field>
        <v-text-field v-model="language" :label="$t('user.language')"></v-text-field>
        <v-text-field v-model="country" :label="$t('user.country')"></v-text-field>
        <v-btn class="mt-2" type="submit" color="primary" prepend-icon="mdi-floppy" block>{{ $t(props.submitButtonText) }}</v-btn>
    </v-form>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import { requiredRule } from '@/lib/formRule';

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

onMounted(() => {
    if (props.user) {
        email.value = props.user.email || "";
        firstName.value = props.user.firstName || "";
        lastName.value = props.user.lastName || "";
        phone.value = props.user.phone || "";
        language.value = props.user.language || "";
        country.value = props.user.country || "";
    }
})

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
