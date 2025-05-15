import { findColumnKeyInHeader } from "~/utils";
import { cloneDeep } from "lodash-es";
import { $api } from "~/utils/api";

export const BaseService = {
  async fetchData(url: string, fetchDataOption: FetchDatatableOption) {
    const clone = cloneDeep(fetchDataOption.options);
    const { page, itemsPerPage, sortBy, sortDesc } = clone;
    const keyword = (fetchDataOption.keyword || "").trim();

    // TODO: Sort column in future
    const finalSortBy: string[] = [];
    const finalSortDesc: string[] = [];
    sortBy.map((value: string, index: number) => {
      const columnKey = findColumnKeyInHeader(value, fetchDataOption.headers);
      const sortType = sortDesc[index] ? "desc" : "asc";
      if (columnKey) {
        finalSortBy.push(columnKey);
        finalSortDesc.push(sortType);
      }
    });

    const fetchParams = {
      q: keyword,
      s: Math.max((page - 1) * itemsPerPage, 0),
      ipp: itemsPerPage,
      sb: finalSortBy,
      sd: finalSortDesc,
      ...fetchDataOption.params,
    };

    if (fetchDataOption.action) fetchParams.action = fetchDataOption.action;

    return await $api(url, {
      params: fetchParams,
    }).then((resp) => {
      if (!fetchDataOption.action) {
        const total = resp.data.totalRecords;
        const items = resp.data.items;
        const options = resp.data.options;
        const extra = resp.data.extra;
        return { total, items, options, extra };
      } else {
        const url = $api.defaults.baseURL + `download/export/${resp.data.link}`;
        window.open(url, "_blank");
      }
    });
  },
};
