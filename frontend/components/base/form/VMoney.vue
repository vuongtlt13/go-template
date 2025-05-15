<script setup lang="ts">
import { hasRequiredRule } from "~/utils";
import { moneyToNumber } from "~/utils/money";

const props = defineProps({
  modelValue: { type: [String, Number], default: "" },
  rules: { type: [Object, String], default: "" },
  showSuccess: { type: Boolean, default: true },
  name: { type: String, default: "" },
  width: { type: String, default: "unset" },
  disabled: { type: Boolean, default: () => false },
  label: { type: String, default: () => "" },
  singleLine: { type: Boolean, default: () => false },
  valueWhenIsEmpty: {
    type: String,
    default: "", // "0" or "" or null
  },
});

const { value, errors, meta, setValue } = useField(props.name, props.rules, {
  label: props.label || props.name,
  syncVModel: true,
});

const isRequired = computed(() => {
  return hasRequiredRule(props.rules);
});

const moneyRenderValue = computed({
  get() {
    return value.value !== null && value.value !== ""
      ? moneyFormat(+value.value!.toString() || 0)
      : props.valueWhenIsEmpty;
  },
  set(newValue) {
    if (newValue === "") {
      setValue("");
      return;
    }
    const newNumber = moneyToNumber(newValue.toString());
    setValue(newNumber);
  },
});
</script>

<template>
  <div class="input-wrapper">
    <v-text-field v-model="value" style="display: none" v-bind="$attrs"></v-text-field>
    <label v-if="!singleLine" class="v-label">
      <template v-if="isRequired"> {{ label }} <span style="color: red">*</span> </template>
      <template v-else>
        {{ label }}
      </template>
    </label>
    <v-text-field
      v-model="moneyRenderValue"
      single-line
      :label="label"
      :disabled="disabled"
      :style="`width: ${width};margin: auto`"
      :error-messages="errors"
      :success="meta.validated && meta.dirty && meta.valid && showSuccess"
      hide-details="auto"
    />
  </div>
</template>
