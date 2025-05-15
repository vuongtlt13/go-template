<script setup lang="ts">
const props = defineProps({
  to: {
    type: [String, Object],
    required: true,
  },
  menuClass: {
    type: String,
    default: "",
  },
  activeClass: {
    type: String,
    default: "",
  },
  icon: {
    type: String,
    default: "",
  },
  title: {
    type: String,
    required: true,
  },
  tooltip: {
    type: String,
    default: "",
  },
});

const computedTooltip = computed(() => {
  return props.tooltip || props.title;
});
</script>

<template>
  <v-tooltip right disabled>
    <template #activator="{ props: _props }">
      <v-list-item :to="to" :class="menuClass" :active-class="activeClass" v-bind="_props" :prepend-icon="icon">
        <v-list-item-title class="text-left">{{ title }}</v-list-item-title>
      </v-list-item>
    </template>
    <span>{{ computedTooltip }}</span>
  </v-tooltip>
</template>

<style scoped></style>
