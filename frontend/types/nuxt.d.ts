import type { VueI18nTranslation, VueI18nTranslationChoice } from "vue-i18n";
import type { Push } from "notivue";

export {};

declare module "#app" {
  interface NuxtApp {
    $notify: Push;
    $t: VueI18nTranslation;
    $tc: VueI18nTranslationChoice;
  }
}

declare module "@vue/runtime-core" {
  interface ComponentCustomProperties {
    $notify: Push;
    $t: VueI18nTranslation;
    $tc: VueI18nTranslationChoice;
  }
}

declare global {
  // eslint-disable-next-line @typescript-eslint/no-empty-object-type
  interface Window {}
}
