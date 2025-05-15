import type { FetchOption } from "~/utils/api";

const AuthService = {
  async login(credential: any) {
    const formData = new FormData();

    for (const key in credential) {
      formData.append(key, credential[key]);
    }

    return await $api("/auth/login", {
      method: "POST",
      body: formData,
    });
  },

  async fetchUserInfo(fetchOption?: FetchOption) {
    const fetch = createFetchFromConfig({
      notifyWhenSuccess: false,
      notifyWhenError: false,
      ...(fetchOption || {}),
    });
    return await fetch("/profile/me");
  },

  async updateProfile(data: any, fetchOption?: FetchOption) {
    const fetch = createFetchFromConfig(fetchOption);
    return await fetch(`/profile/me`, {
      method: "PUT",
      body: data,
    });
  },
};

export default AuthService;
