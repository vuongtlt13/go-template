<script setup lang="ts">
import AdminAppBar from "~/components/admin/AppBar.vue";
import menuComponents from "~/utils/menu";
</script>

<template>
  <v-responsive class="border rounded">
    <v-app>
      <TheNavigation home-path="/admin" :menu-components="menuComponents" />
      <AdminAppBar />
      <v-main>
        <v-container>
          <slot />
        </v-container>
      </v-main>
      <v-footer app :height="1" />
    </v-app>
  </v-responsive>
</template>
