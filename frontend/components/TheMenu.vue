<script setup lang="ts">
import MenuNoChild from "~/components/menu/NoChild.vue";
import MenuHasChild from "~/components/menu/HasChild.vue";
import { MenuType } from "~/utils/constants";
import type { Menu } from "~/utils/menu";
import { useAuthStore } from "~/stores/auth";

defineProps({
  menuComponents: {
    type: Array as PropType<Menu[]>,
    required: true,
  },
});

const authStore = useAuthStore();

const isValidPermission = (userPermissions: string[], menu: Menu) => {
  const menuPermissions = menu.permissions || [];

  if (userPermissions.length == 0 && menuPermissions.length > 0) return false;

  let result = true;
  menuPermissions.map(function (permission: string) {
    if (!userPermissions.includes(permission)) result = false;
  });

  return result;
};

const isValidRole = (userRoles: string[], menu: Menu) => {
  const menuRoles = menu.roles || [];
  if (userRoles.length == 0 && menuRoles.length > 0) return false;

  let result = true;
  menuRoles.map(function (role: string) {
    if (!userRoles.includes(role)) result = false;
  });

  return result;
};

const isShowable = (menu: Menu) => {
  const userRoles = (authStore.userRoles || []) as string[];
  const userPermissions = (authStore.userPermissions || []) as string[];

  if ((userRoles || []).includes("super_admin")) return true;

  return isValidPermission(userPermissions, menu) && isValidRole(userRoles, menu);
};
</script>

<template>
  <v-list nav dense shaped subheader>
    <template v-for="menu in menuComponents">
      <template v-if="isShowable(menu)">
        <MenuNoChild
          v-if="menu.type === MenuType.NO_CHILD"
          :key="`mnc-${menu.title}`"
          :to="menu.to || '#'"
          :title="$t(menu.title)"
          :active-class="menu.activeClass"
          :icon="menu.icon"
        />
        <MenuHasChild
          v-else-if="menu.type === MenuType.HAS_CHILD"
          :key="`mhc-${menu.title}`"
          :to="menu.to"
          :title="$t(menu.title)"
          :active-class="menu.activeClass"
          :icon="menu.icon"
          :children="menu.children || []"
        />
        <v-divider v-else-if="menu.type === MenuType.DIVIDER" :key="`m-${menu.title}`" />
        <v-list-subheader v-else-if="menu.type === MenuType.SUB_HEADER" :key="`msh-${menu.title}`" :class="menu.class">
          <span class="">{{ menu.title }}</span>
        </v-list-subheader>
      </template>
    </template>
  </v-list>
</template>

<style scoped>
.no-active:hover {
  cursor: pointer;
}
</style>
