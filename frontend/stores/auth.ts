import Cookies from "js-cookie";
import AuthService from "~/services/auth";
import { TOKEN_KEY } from "~/utils/constants";
import * as _ from "lodash-es";

// In Setup Stores:
//   ref()s become state properties
//   computed()s become getters
//   function()s become actions
export const useAuthStore = defineStore("auth", () => {
  const initUserValue = {
    name: "",
    fullName: "",
    permissions: [],
    userRoles: [],
  };

  const user = ref(_.cloneDeep(initUserValue));
  const _token = ref("");

  const userRoles = computed(() => {
    return ((user.value || {}).userRoles || []) as string[];
  });

  const userPermissions = computed(() => {
    return ((user.value || {}).permissions || []) as string[];
  });

  const userName = computed(() => {
    return (user.value || {}).name || (user.value || {}).fullName;
  });

  const token = computed(() => {
    let tokenInCookie: string | undefined = "";
    if (import.meta.client) {
      tokenInCookie = Cookies.get(TOKEN_KEY);
    }
    return _token.value || tokenInCookie;
  });

  const isLogged = computed(() => {
    return token.value && user.value && userName.value;
  });

  const saveToken = (token: string, remember: boolean) => {
    _token.value = token;
    if (import.meta.client) {
      Cookies.set(TOKEN_KEY, token, { expires: remember ? 365 : undefined });
    }
  };

  const logout = async () => {
    user.value = _.cloneDeep(initUserValue);
    _token.value = "";
    await removeToken();
  };

  const removeToken = () => {
    return new Promise((resolve) => {
      if (import.meta.client) {
        Cookies.remove(TOKEN_KEY);
      }
      resolve(true);
    });
  };

  const setUser = (userInfo: any) => {
    user.value = userInfo;
  };

  const fetchUser = async () => {
    try {
      const { data } = await AuthService.fetchUserInfo();
      setUser(data);
    } catch (err) {
      console.error(err);
      await logout();
    }
  };

  const isSuperAdmin = computed(() => {
    return isLogged && userRoles.value.includes("super_admin");
  });

  return {
    user,
    token,
    userRoles,
    userPermissions,
    userName,
    isLogged,
    isSuperAdmin,
    saveToken,
    removeToken,
    setUser,
    logout,
    fetchUser,
  };
});
