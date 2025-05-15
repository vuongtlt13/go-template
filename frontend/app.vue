<script setup lang="ts">
const customNotifyTheme = {
  ...lightTheme,
};

const appStore = useAppStore();
</script>

<template>
  <NuxtLayout>
    <NuxtPage />
  </NuxtLayout>

  <Notivue v-slot="item">
    <Notification :theme="customNotifyTheme" :item="item">
      <NotificationProgress :item="item" />
    </Notification>
  </Notivue>

  <v-overlay :value="appStore.overlay" z-index="100000">
    <v-progress-circular indeterminate size="64"></v-progress-circular>
  </v-overlay>
</template>

<style>
:root {
  /* Your variables */
  --header-height: 18px;
  --container-padding: 30px;

  /* Add the top padding and place it below the header */
  --nv-root-top: calc(var(--header-height) + var(--container-padding));

  /* Add the same left-right paddings of your app container */
  --nv-root-left: var(--container-padding);
  --nv-root-right: var(--container-padding);
}

/* Rules for mobile devices */
@media (max-width: 768px) {
  :root {
    --header-height: 60px;
    --container-padding: 20px;
  }
}
</style>
<style scoped lang="scss"></style>
