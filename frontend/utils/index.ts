import { isString } from "lodash-es";
import type { DataTableUpdateRowCallbackFunc, DataTableUpdateRowFunc, SelectedCellDataTable } from "~/types/datatable";

/**
 * Get cookie from request.
 *
 * @param  {Object} req
 * @param  {String} key
 * @return {String|undefined}
 */
export function cookieFromRequest(req: any, key: any) {
  if (!req.headers.cookie) {
    return;
  }

  const cookie = req.headers.cookie.split(";").find((c: string) => c.trim().startsWith(`${key}=`));

  if (cookie) {
    return cookie.split("=")[1];
  }
}

/**
 * https://router.vuejs.org/en/advanced/scroll-behavior.html
 */
export function scrollBehavior(to: any, _: any, savedPosition: any) {
  if (savedPosition) {
    return savedPosition;
  }

  let position = {};

  if (to.matched.length < 2) {
    position = { x: 0, y: 0 };
  } else if (to.matched.some((r: any) => r.components.default.options.scrollToTop)) {
    position = { x: 0, y: 0 };
  }
  if (to.hash) {
    position = { selector: to.hash };
  }

  return position;
}

/**
 * Check a href is a external link
 */
export const isExternalLink = (link: string | null | undefined): boolean => {
  link = link || "";
  link = link.trim();
  if (link === "") {
    return false;
  }
  return link.startsWith("http://") || link.startsWith("https://");
};

/**
 * Check a href is a internal link
 */
export const isInternalLink = (link: string | null | undefined): boolean => {
  link = link || "";
  link = link.trim();
  if (link === "") {
    return false;
  }
  return !isExternalLink(link);
};

export const sleep = (time: number) => {
  return new Promise((resolve: any) => {
    setTimeout(() => {
      resolve();
    }, time);
  });
};

export const generateTitle = (title: string): string => {
  const appStore = useAppStore();
  return `${title} - ${appStore.appName}`;
};

export function hasRequiredRule(rules: any) {
  if (isString(rules)) {
    return rules.split("|").includes("required");
  }

  return false;
}

export const findColumnKeyInHeader = (value: string, headers: WrapperDataTableHeader[]): string | undefined => {
  for (const i in headers) {
    if (headers[i].value == value && headers[i].columnKey) return headers[i].columnKey;
  }
  return undefined;
};

export const DefaultUpdateRowCallback: DataTableUpdateRowCallbackFunc = (
  invalid: boolean,
  item: any,
  _?: FetchOption,
): Promise<any> => {
  return new Promise((resolve, _) => {
    resolve({
      data: {
        data: item,
      },
    });
  });
};

export const makeUpdateRowDatatableFunc = (
  callback?: DataTableUpdateRowCallbackFunc,
  fetchOption?: FetchOption,
): DataTableUpdateRowFunc => {
  const finalFetchOption = {
    notifyWhenSuccess: false,
    ...fetchOption,
  };

  return (invalid: boolean, item: any, selectedCell: SelectedCellDataTable, fetchOption?: FetchOption) => {
    const _finalFetchOption = {
      ...finalFetchOption,
      ...fetchOption,
    };
    if (!invalid) {
      const oldValue = item[selectedCell.column];
      const newValue = selectedCell.value;
      if (oldValue == newValue) return;
      item[selectedCell.column] = selectedCell.value;
      const finalCallback: DataTableUpdateRowCallbackFunc = callback || DefaultUpdateRowCallback;
      return finalCallback!(selectedCell.row, item, _finalFetchOption)
        .then((resp) => {
          item[selectedCell.column] = resp.data.data[selectedCell.column];
        })
        .catch(() => {
          item[selectedCell.column] = oldValue;
        });
    }
  };
};

export const scrollIntoView = (target: HTMLElement, container: HTMLElement | null | undefined) => {
  if (!container) return;

  const containerTop = container.scrollTop;
  const containerBottom = containerTop + container.clientHeight;

  const elemTop = target.offsetTop;
  const elemBottom = elemTop + target.offsetHeight;

  if (elemTop < containerTop) {
    container.scrollTop = elemTop;
  } else if (elemBottom > containerBottom) {
    container.scrollTop = elemBottom - container.clientHeight;
  }
};
