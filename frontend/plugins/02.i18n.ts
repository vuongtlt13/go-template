import CommonService from "~/services/common";
import type { VueI18n } from "vue-i18n";
import { useAppStore } from "~/stores/app";
import { switchLanguage, initializeI18n } from "~/utils/i18n";

export default defineNuxtPlugin((nuxtApp) => {
  const i18nInstance: VueI18n = nuxtApp.$i18n as VueI18n;
  // init accessor
  initializeI18n(i18nInstance);

  const appStore = useAppStore();

  if (Object.keys(i18nInstance.getLocaleMessage("en")).length === 0) {
    CommonService.fetchLanguages("en")
      .then((resp) => {
        const messages = resp.data;
        i18nInstance.setLocaleMessage("en", messages);
      })
      .then(() => {
        switchLanguage(appStore.appLocale);
      });
  }

  return {
    provide: {
      t: i18nInstance.t,
      tc: i18nInstance.tc,
    },
  };
});
