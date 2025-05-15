export type DataOptions = any;

interface InternalItem<T = any> {
  value: any;
  raw: T;
}
type FilterMatch = boolean | number | [number, number] | [number, number][];
type FilterFunction = (value: string, query: string, item?: InternalItem) => FilterMatch;
type DataTableCompareFunction<T = any> = (a: T, b: T) => number | null;
type SelectItemKey<T = Record<string, any>> =
  | boolean
  | null
  | undefined
  | string
  | readonly (string | number)[]
  | ((item: T, fallback?: any) => any);
type DataTableHeader<T = Record<string, any>> = {
  key?: "data-table-group" | "data-table-select" | "data-table-expand" | (string & {});
  value?: SelectItemKey<T>;
  title?: string;
  fixed?: boolean;
  align?: "start" | "end" | "center";
  width?: number | string;
  minWidth?: string;
  maxWidth?: string;
  nowrap?: boolean;
  headerProps?: Record<string, any>;
  cellProps?: any;
  sortable?: boolean;
  sort?: DataTableCompareFunction;
  sortRaw?: DataTableCompareFunction;
  filter?: FilterFunction;
  children?: DataTableHeader<T>[];
};
export type DataTableFetchDataFunc = (options: DataOptions, keyword: string, ...args: any[]) => Promise<any>;

export interface EditableColumnDataTable {
  value: string;
  width?: string;
  type?: string;
  rules?: string;
}

export interface SelectedCellDataTable {
  row: any;
  column: any;
  value: any;
  origin: any;
  item: any;
  meta: any;
  render: boolean;
}

export type DataTableUpdateRowFunc = (
  invalid: boolean,
  item: any,
  selectedCell: SelectedCellDataTable,
  fetchOption?: FetchOption,
) => Promise<any> | undefined;

export type DataTableUpdateRowCallbackFunc = (invalid: boolean, item: any, fetchOption?: FetchOption) => Promise<any>;

export interface FetchDatatableOption {
  options: DataOptions;
  keyword: string | null;
  headers: DataTableHeader[];
  params?: any;
  action?: string;
}

export type FetchDatatableFunc = (fetchDataOption: FetchDatatableOption) => Promise<any>;

export interface DataTableHandler {
  items: Ref<any[]>;
  totalItem: Ref<number>;
  extraData: Ref;
  headers: DataTableHeader[];
  itemsPerPage: Ref<number>;
  selectedRows: Ref<any>;
  selectedCell: Ref<SelectedCellDataTable>;
  draw: Ref<number>;
  focused: Ref<boolean>;
  options: Ref<DataOptions>;
  exportData: (action: string) => Promise<any>;
  searchKeyword: Ref<any>;
  fetchExtraParams: Ref<any>;
  selectAllItems: () => void;
  reloadTableFn: (delay?: boolean) => void;
  clearSelectionAndReload: (delay?: boolean) => void;
  updateSelectedCell: (row: any, column: any) => void;
  resetSelectedCell: () => void;
  updateRowFunc: DataTableUpdateRowFunc | undefined;
  closeEditor: () => void;
  fetchDatatableFunc: (options: DataOptions) => void;
  loading: Ref<boolean>;
  page: Ref<number>;
  pageCount: Ref<number>;
}

export interface WrapperDataTableHeader extends DataTableHeader {
  show?: boolean;
  columnKey?: string;

  // region for editable
  editable?: boolean;
  inputWidth?: string;
  type?: string;
  rules?: string;
  // endregion for editable
}
