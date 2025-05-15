<script setup lang="ts">
import AuthService from "~/services/auth";
import { sleep } from "~/utils";
import logoImg from "assets/images/logo.png";
import { useAuthStore } from "~/stores/auth";

const authStore = useAuthStore();
const router = useRouter();
const route = useRoute();
const { t } = useI18n();

onMounted(async () => {
  if (authStore.token) {
    await router.push({ path: "/" });
  }
});

const username = ref("");
const password = ref("");
const buttonTitle = ref(t("auth.login_button_text"));
const usernameField = "username";

const submit = async () => {
  try {
    buttonTitle.value = t("auth.logging_button_text");
    await sleep(300);
    const credential = {
      [usernameField]: username.value,
      password: password.value,
    };

    const resp = await AuthService.login(credential);
    const at = resp.access_token;
    authStore.saveToken(at, true);
    const redirect = (route.query.redirect || "/") as any;
    await router.push({ path: redirect! });
  } catch (err: any) {
    console.error(err);
  }

  buttonTitle.value = t("auth.login_button_text");
};
</script>

<template>
  <v-app class="login-background">
    <v-main>
      <v-container fluid class="auth">
        <div class="auth-wrapper">
          <v-card class="rounded-md login-card elevation-10" flat max-width="500px">
            <v-card-text class="v-card-item pa-sm-8">
              <div class="login-form-wrapper">
                <v-row style="padding-top: 16px !important; padding-bottom: 16px !important">
                  <v-col class="d-flex">
                    <img class="mx-auto" alt="logo" :src="logoImg" style="height: 64px" />
                  </v-col>
                </v-row>

                <v-row style="margin: auto">
                  <v-col style="margin: auto; text-align: center">
                    <h2>{{ t("Welcome! Sign in") }}</h2>
                  </v-col>
                </v-row>

                <v-row class="d-flex mb-3" style="margin: auto">
                  <VeeForm :callback="submit">
                    <v-col cols="24">
                      <label class="v-label">{{ t("auth.username") }}</label>
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
                    <v-col cols="24">
                      <label class="v-label">{{ t("auth.password") }}</label>
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

                    <v-col cols="24">
                      <v-btn :disabled="false" class="mr-4 primary" type="submit" width="100%">
                        {{ buttonTitle }}
                      </v-btn>
                    </v-col>
                  </VeeForm>
                </v-row>
              </div>
            </v-card-text>
          </v-card>
        </div>
      </v-container>
    </v-main>
  </v-app>
</template>

<style lang="scss" scoped>
.login-form-wrapper {
  min-width: 320px;
  max-width: 100%;
}

.login-background {
  background-color: #fafafa00;
}

.auth {
  height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
}

.login-card {
  background-color: white;
  overflow: visible;
}

.action-wrapper {
  width: 40vw;
  max-width: 500px;
}
</style>
