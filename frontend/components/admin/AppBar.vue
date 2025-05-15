<script setup lang="ts">
import { useAppStore } from "~/stores/app";
const appBarHeight = 50;
const appStore = useAppStore();
const { toggleDrawer } = appStore;
</script>

<template>
  <v-app-bar app dense :height="appBarHeight" elevation="1" style="z-index: 100">
    <v-app-bar-nav-icon class="" @click="toggleDrawer" />
    <v-spacer />
    <TheRightAppBar />
  </v-app-bar>
</template>

<style scoped></style>
