<template>
    <v-form @submit.prevent="submit" ref="form">
        <v-text-field v-model="name" :rules="[requiredRule]" :label="$t('global.name')"></v-text-field>
        <v-color-picker v-model="color" class="mb-2"></v-color-picker>
        <v-btn class="mt-2" type="submit" color="primary" prepend-icon="mdi-floppy" block>{{ $t(props.submitButtonText) }}</v-btn>
    </v-form>
</template>

<script lang="ts" setup>
import { onMounted, ref } from 'vue';
import { requiredRule } from '@/lib/formRule';

const props = defineProps({
    onSubmit: Function,
    submitButtonText: String,
    status: Object as any,
})

const name = ref("");
const color = ref("#00FF00");
const form = ref(null);

onMounted(() => {
    if (props.status) {
        name.value = props.status.name || "";
        color.value = props.status.color || "#00FF00";
    }
})

const submit = async () => {
    const { valid } = await form.value.validate();
    if (valid) {
        const statusData = {
            name: name.value,
            color: color.value,
        };
        props.onSubmit(statusData);
    }
};
</script>
