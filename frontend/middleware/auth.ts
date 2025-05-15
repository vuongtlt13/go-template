import { useAuthStore } from "~/stores/auth";
import { HttpCode } from "~/utils/constants";

export default defineNuxtRouteMiddleware(async (to) => {
  const authStore = useAuthStore();
  if (authStore.token) {
    if (!authStore.isLogged) {
      await authStore.fetchUser();
      // check permission
      if (authStore.isSuperAdmin) return;

      // const route = useRoute();
      const requiredPermissions: string[] = [];
      // console.log(route);
      // route.meta.map((meta: any) => {
      //   if (meta.rp) requiredPermissions.push(...meta.rp);
      // });

      if (requiredPermissions.length == 0) return;

      const permissions: string[] = authStore.userPermissions;
      let hasPermission = true;
      requiredPermissions.map((permission) => {
        if (!permissions.includes(permission)) hasPermission = false;
      });

      if (!hasPermission) throw { statusCode: HttpCode.PERMISSION_DENIED, message: $i18n.t("auth.permission_denied") };
    }
  } else {
    const router = useRouter();
    await router.push({
      path: "/auth/login",
      query: {
        redirect: to.fullPath,
      },
    });
  }
});
