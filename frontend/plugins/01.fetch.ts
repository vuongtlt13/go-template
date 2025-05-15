import {
  createFetchFromConfig, initializeApi,
} from "~/utils/api";

export default defineNuxtPlugin((_) => {
  const api = createFetchFromConfig({
    notifyWhenError: true,
    notifyWhenSuccess: false,
    overlay: true,
  })

  initializeApi(api);
});
