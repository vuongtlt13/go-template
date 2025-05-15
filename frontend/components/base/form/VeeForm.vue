<script setup lang="ts">
import type { FormSubmitFunc } from "~/types";

const props = defineProps({
  callback: { type: Function as PropType<FormSubmitFunc>, default: undefined },
});
const { handleSubmit } = useForm({});
const onSubmit = handleSubmit((values) => {
  if (props.callback) props.callback(values);
});
</script>

<template>
  <form style="width: 100%" @submit.prevent="onSubmit">
    <slot />
  </form>
</template>

<style scoped></style>
