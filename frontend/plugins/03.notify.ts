import { $notify } from "~/utils/notify";

export default defineNuxtPlugin((_) => {
  return {
    provide: {
      notify: $notify,
    },
  };
});
