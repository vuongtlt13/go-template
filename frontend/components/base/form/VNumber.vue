<script setup lang="ts">
import { hasRequiredRule } from "~/utils";

const props = defineProps({
  modelValue: { type: [Number, String], default: "" },
  rules: { type: [Object, String], default: "" },
  showSuccess: { type: Boolean, default: true },
  name: { type: String, default: "" },
  width: { type: String, default: "unset" },
  disabled: { type: Boolean, default: () => false },
  label: { type: String, default: () => "" },
  singleLine: { type: Boolean, default: () => false },
});

const { value, errors, meta } = useField(props.name, props.rules, {
  label: props.label || props.name,
  syncVModel: true,
});

const isRequired = computed(() => {
  return hasRequiredRule(props.rules);
});
</script>

<template>
  <div class="input-wrapper">
    <label v-if="!singleLine" class="v-label">
      <template v-if="isRequired"> {{ label }} <span style="color: red">*</span> </template>
      <template v-else>
        {{ label }}
      </template>
    </label>
    <v-text-field
      v-model.number="value"
      type="number"
      single-line
      :label="label"
      :disabled="disabled"
      :style="`width: ${width};margin: auto`"
      :error-messages="errors"
      :success="meta.validated && meta.dirty && meta.valid && showSuccess"
      hide-details="auto"
      v-bind="$attrs"
    />
  </div>
</template>
