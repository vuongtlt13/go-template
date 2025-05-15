<script setup lang="ts">
definePageMeta({
  layout: "admin",
});
const { t } = useI18n();
const username = ref("sometext");
const password = ref("password");
const number = ref(123);
const money = ref("");
</script>

<template>
  <v-row>
    <v-col cols="12">
      <h1>Inputs</h1>
    </v-col>

    <v-col cols="12">
      <VText
        v-model="username"
        single-line
        :placeholder="t('auth.username')"
        :show-success="false"
        name="username"
        label="username"
        rules="required|min:4"
      />
    </v-col>

    <v-col cols="12">
      <VPassword
        v-model="password"
        single-line
        :show-success="false"
        :placeholder="t('auth.password')"
        name="password"
        label="Mật khẩu"
        rules="required|min:6"
      />
    </v-col>
    <v-col cols="12">
      <VNumber
        v-model="number"
        single-line
        :placeholder="t('Number')"
        :show-success="false"
        name="number"
        label="number"
        rules="required|min:4"
      />
    </v-col>
    <v-col cols="12">
      <VMoney
        v-model="money"
        single-line
        :placeholder="t('Money')"
        :show-success="false"
        name="money"
        label="money"
        rules="required|min:4"
      />
    </v-col>
  </v-row>
</template>

<style scoped></style>
