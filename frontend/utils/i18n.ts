import CommonService from "~/services/common";
import type { VueI18n } from "vue-i18n";

let $i18n: VueI18n; // eslint-disable-line import/no-mutable-exports

export function initializeI18n(i18n: VueI18n) {
  $i18n = i18n;
}

export default defineI18nConfig(() => ({
  legacy: false,
  fallbackLocale: "en",
  silentTranslationWarn: true,
  defaultLocale: process.env.APP_LOCALE || "en",
  messages: {
    en: {},
    vi: {},
  },
}));

export async function switchLanguage(locale: string) {
  if (Object.keys($i18n.getLocaleMessage(locale)).length === 0) {
    const resp = await CommonService.fetchLanguages(locale);
    const messages = resp.data;
    $i18n.setLocaleMessage(locale, messages);
  }

  if ($i18n.locale !== locale) {
    $i18n.locale = locale;
  }
}

export { $i18n };
