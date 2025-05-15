import type { FetchContext, ResponseMap, FetchResponse, $Fetch } from "ofetch";
import { ofetch } from "ofetch";

let $api: $Fetch; // eslint-disable-line import/no-mutable-exports

export const showNotificationFromResponse = (response: FetchResponse<any>) => {
  console.log("showNotificationFromResponse", response);
  // if (response.status !== HttpCode.SUCCESS || !response.success) {
  //   Vue.notify({
  //     type: "error",
  //     title: i18n.t("notification.error_title").toString(),
  //     text: i18n.t(response.data.message || response.data.error).toString(),
  //     duration: NOTIFICATION_DURATION,
  //   });
  // } else {
  //   Vue.notify({
  //     type: "success",
  //     title: i18n.t("notification.success_title").toString(),
  //     text: i18n.t(response.data.message).toString(),
  //     duration: NOTIFICATION_DURATION,
  //   });
  // }
};

export const showNotificationFromErrorResponse = async (err: any, fetchOption?: FetchOption) => {
  const resp: FetchResponse<any> = err.response;
  if (!resp) {
    console.log("error", err);
    return;
  }
  console.log("Show error", resp);
  if (!resp.bodyUsed) {
    await resp.json();
  }
  const respData = resp._data;
  if (resp.status >= HttpCode.SERVER_ERROR) {
    if (fetchOption?.notifyWhenError)
      $notify.error({
        title: $i18n.t("notification.server_error_title"),
        message: $i18n.t(respData.message || respData.error || err.toString()).toString(),
      });
  } else if (resp.status === HttpCode.UNAUTHORIZED) {
    if (fetchOption?.notifyWhenError)
      $notify.error({
        title: $i18n.t("notification.token_expired_alert_title"),
        message: $i18n.t("notification.token_expired_alert_text"),
      });
    if (!fetchOption?.disableRedirect) {
      const authStore = useAuthStore();
      await authStore.logout();
    }
  } else {
    if (fetchOption?.notifyWhenError) {
      $notify.error({
        title: $i18n.t("notification.error_title"),
        message: $i18n.t(respData.message || respData.error).toString(),
      });
    }
  }
};

type ResponseType = keyof ResponseMap | "json";

type onRequestOption = FetchContext<any, ResponseType>;
type onRequestErrorOption = onRequestOption & { error: Error };
type onResponseOption<T> = onRequestOption & { response: FetchResponse<T> };
type onResponseErrorOption<T> = onResponseOption<T>;

export const onRequestDefault = ({ options }: onRequestOption) => {
  const config = useRuntimeConfig();
  const authStore = useAuthStore();

  options.baseURL = config.public.baseURL;
  if (authStore.token) {
    options.headers.set("Authorization", `Bearer ${authStore.token}`);
  }
};
export const onRequestErrorDefault = ({ error }: onRequestErrorOption) => {
  console.error("request error", error);
};
export const onResponseDefault = ({ response }: onResponseOption<any>, fetchOption?: FetchOption) => {
  // Log response
  // console.log("[fetch response]", request, response.status, response.body);
  if (fetchOption?.overlay) {
    const appStore = useAppStore();
    appStore.setOverlay(false);
  }

  if (fetchOption?.notifyWhenSuccess) {
    showNotificationFromResponse(response);
  }
};
export const onResponseErrorDefault = (err: onResponseErrorOption<any>, fetchOption?: FetchOption) => {
  if (fetchOption?.overlay) {
    const appStore = useAppStore();
    appStore.setOverlay(false);
  }

  showNotificationFromErrorResponse(err, fetchOption);
};

export const createFetchFromConfig = (option?: FetchOption): $Fetch => {
  if (option == undefined && $api != undefined) return $api;

  return ofetch.create({
    onRequest: onRequestDefault,
    onRequestError: onRequestErrorDefault,
    onResponse: (resOption) => {
      onResponseDefault(resOption, {
        notifyWhenError: true,
        notifyWhenSuccess: false,
        overlay: true,
        ...option,
      });
    },
    onResponseError: (resOption) => {
      onResponseErrorDefault(resOption, {
        notifyWhenError: true,
        notifyWhenSuccess: false,
        overlay: true,
        ...option,
      });
    },
  });
};

export function initializeApi(initApi: $Fetch) {
  $api = initApi;
}

export const convertToVSelectOption = (options: any): object => {
  const selectOptions = [] as any[];
  for (const key in options) {
    let value: any = key;
    if (!isNaN(+key)) value = +key;
    selectOptions.push({
      text: options[key],
      value: value,
    });
  }

  return selectOptions;
};

export const makeOptionFromResponse = (optionResp: any) => {
  const finalOptionResp = {} as any;
  Object.keys(optionResp).forEach((key) => {
    finalOptionResp[key] = convertToVSelectOption(optionResp[key]);
  });
  return finalOptionResp;
};
export { $api };
