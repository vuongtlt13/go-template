<script setup lang="ts">
const authStore = useAuthStore();
const router = useRouter();

const items = [
  {
    title: $i18n.t("auth.app.profile"),
    icon: "mdi-account",
    onclick: () => {
      router.push("/profile");
    },
  },
  {
    type: "button",
    title: $i18n.t("auth.logout"),
    divider: false,
    onclick: () => {
      authStore.logout().then(() => {
        setTimeout(() => {
          window.location.reload();
        }, 100);
      });
    },
  },
];
</script>

<template>
  <div class="d-flex">
    <v-menu
      class="white-background"
      offset-y
      bottom
      left
      transition="scale-transition"
      origin="right top"
      nudge-top="-9"
    >
      <template #activator="{ props }">
        <v-avatar size="32" style="cursor: pointer; margin: auto; display: table-cell" v-bind="props" class="mx-2">
          <img src="../assets/images/default_avatar.jpg" alt="Avatar" />
        </v-avatar>
      </template>
      <v-list class="pa-0 rounded-md mt-2" style="width: 200px">
        <v-list-item key="account_name" style="text-align: left" class="text-sm-body-2">
          <template #title>
            Xin chào <strong>{{ authStore.userName }}</strong>
          </template>
        </v-list-item>
        <v-divider />
        <template v-for="(item, index) in items" :key="`g-${index}`">
          <v-divider v-if="item.divider" />
          <v-list-item
            v-if="item.type != 'button'"
            :prepend-icon="item.icon ? item.icon : undefined"
            link
            style="text-align: left; cursor: pointer"
            class="text-sm-body-2"
            @click="item.onclick"
          >
            <template #title>{{ item.title }}</template>
          </v-list-item>

          <template v-else>
            <v-list-item key="btn" style="text-align: left" class="text-sm-body-2">
              <v-btn
                variant="outlined"
                class="mx-auto"
                :icon="item.icon"
                color="primary"
                style="width: 160px"
                @click="item.onclick"
              >
                {{ item.title }}
              </v-btn>
            </v-list-item>
          </template>
        </template>
      </v-list>
    </v-menu>
  </div>
</template>

<style scoped></style>
