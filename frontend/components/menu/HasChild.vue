<script setup lang="ts">
import { MenuType } from "~/utils/constants";
import type { PropType } from 'vue'

defineProps({
  activeClass: {
    type: String,
    default: "",
  },
  menuClass: {
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
  children: {
    type: Array as PropType<Array<Menu>>,
    required: true,
  },
  subGroup: {
    type: Boolean,
    default: false,
  },
  value: {
    type: Boolean,
    default: true,
  },
});
</script>

<template>
  <v-list-group
    no-action
    :value="value"
    :sub-group="subGroup"
    :class="menuClass"
    :prepend-icon="!subGroup ? icon : undefined"
    append-icon="mdi-chevron-down"
  >
    <template #activator="{ props }">
      <v-list-item
        v-bind="props"
        :title="title"
        class="menu-parent-title text-left"
      />
    </template>

    <template v-for="menu in children">
      <MenuNoChild
        v-if="menu.type === MenuType.NO_CHILD"
        :key="`mnc-${menu.title}`"
        :to="menu.to!"
        :title="$t(menu.title)"
        :menu-class="menu.menuClass"
        :active-class="menu.activeClass"
        :icon="menu.icon"
        style="padding-left: 32px"
      />
      <MenuHasChild
        v-else-if="menu.type === MenuType.HAS_CHILD"
        :key="`mhc-${menu.title}`"
        :to="menu.to"
        :title="$t(menu.title)"
        :menu-class="menu.menuClass"
        :active-class="menu.activeClass"
        :icon="menu.icon"
        :children="menu.children || []"
        style="padding-left: 8px"
        :sub-group="true"
      />
    </template>
  </v-list-group>
</template>

<style scoped></style>
