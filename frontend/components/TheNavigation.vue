<script setup lang="ts">
import TheMenu from "~/components/TheMenu.vue";
import { useAppStore } from "~/stores/app";
import logoImg from "assets/images/logo.png";

defineProps({
  menuComponents: {
    type: Array as PropType<Menu[]>,
    required: true,
  },
  homePath: {
    type: String,
    default: "/",
  },
});

const appStore = useAppStore();

const drawer = computed({
  get(): boolean {
    return appStore.isDrawer;
  },
  set(isDrawer: boolean) {
    appStore.setDrawer(isDrawer);
  },
});

const useMini = computed({
  get(): boolean {
    return appStore.useMini;
  },
  set(isUseMini: boolean) {
    appStore.setUseMini(isUseMini);
  },
});
</script>

<template>
  <v-navigation-drawer
    v-model="drawer"
    :rail="useMini"
    rail-width="56"
    permanent
    :width="200"
    style="max-width: 215px; top: auto"
    app
  >
    <NuxtLink :to="homePath">
      <img alt="logo" :src="logoImg" class="logo-img" />
    </NuxtLink>
    <TheMenu :menu-components="menuComponents" />
  </v-navigation-drawer>
</template>

<style scoped>
.logo-img {
  cursor: pointer;
  height: 48px;
  padding-left: 5px;
  margin-top: auto;
  margin-bottom: auto;
}
</style>
