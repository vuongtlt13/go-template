import Cookies from "js-cookie";

// In Setup Stores:
//   ref()s become state properties
//   computed()s become getters
//   function()s become actions
export const useAppStore = defineStore("app", () => {
  const overlay = ref(false);
  const drawer = ref(!isMobileScreen());
  const useMini = ref(isMobileScreen());
  const isDrawer = computed(() => drawer.value);
  const appName = ref("");
  const companyName = ref("");
  const apiURL = ref(null);
  const _appLocale = ref(process.env.APP_LOCALE || "en");

  const setConfig = (config: { appName: string; appLocale: string }) => {
    appName.value = config.appName || "";
    _appLocale.value = config.appLocale || "en";
  };

  const resetConfig = () => {
    appName.value = "";
    _appLocale.value = process.env.APP_LOCALE || "en";
  };

  const setOverlay = (value: boolean) => {
    overlay.value = value;
  };

  const setDrawer = (value: boolean) => {
    drawer.value = value;
  };

  const setUseMini = (value: any = null) => {
    if (value !== undefined && value !== null) {
      useMini.value = value;
    } else useMini.value = !useMini.value;
  };

  const toggleDrawer = () => {
    // setDrawer(!drawer.value);
    setUseMini(!useMini.value);
  };

  const appLocale = computed(() => {
    let langInCookie: string | undefined = "";
    if (import.meta.client) {
      langInCookie = Cookies.get("locale");
    }
    return _appLocale.value || langInCookie || "en";
  });

  const setLocale = (lang: string) => {
    _appLocale.value = lang;
    if (import.meta.client) {
      Cookies.set("locale", lang, { expires: 365 });
    }
  };

  return {
    overlay,
    appName,
    companyName,
    apiURL,
    appLocale,
    isDrawer,
    useMini,
    setConfig,
    resetConfig,
    setOverlay,
    setDrawer,
    toggleDrawer,
    setLocale,
    setUseMini,
  };
});
