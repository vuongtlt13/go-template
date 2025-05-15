import { BaseService } from "~/services/base-service";

const UserService = {
  async fetchUser(fetchDataOption: FetchDatatableOption) {
    return BaseService.fetchData("/crud/user/", fetchDataOption);
  },

  async loadCreateUserOption(params: any) {
    return await $api("/crud/user/create", {
      params,
    });
  },

  async loadUpdateUserOption(id: any, params: any) {
    return await $api(`/crud/user/${id}/edit`, {
      params,
    });
  },

  async addNewUser(data: any, fetchOption?: FetchOption) {
    const fetch = createFetchFromConfig(fetchOption);
    return await fetch("/crud/user/", {
      method: "POST",
      body: data,
    });
  },

  async updateUser(id: any, data: any, fetchOption?: FetchOption) {
    const fetch = createFetchFromConfig(fetchOption);
    return await fetch(`/crud/user/${id}`, {
      method: "PUT",
      body: data,
    });
  },

  async deleteUser(id: any, fetchOption?: FetchOption) {
    const fetch = createFetchFromConfig(fetchOption);
    return await fetch(`/crud/user/${id}`, {
      method: "DELETE",
    });
  },

  async deleteUsers(ids: any[], fetchOption?: FetchOption) {
    const fetch = createFetchFromConfig(fetchOption);
    return await fetch(`/crud/user/batch`, {
      method: "DELETE",
      body: {
        items: ids,
      },
    });
  },
};
export default UserService;
