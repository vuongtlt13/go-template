<template>
  <div>
    <v-data-table-server
      v-model:items-per-page="itemsPerPage"
      v-model="selectedRows"
      v-model:page="page"
      :class="dataTableClasses"
      :headers="headers"
      :height="height"
      :item-value="itemValue"
      :items="items"
      :items-length="totalItem"
      :loading="loading"
      :multi-sort="multiSort"
      :show-select="showSelect"
      :single-select="singleSelect"
      :style="dStyle"
      :row-props="generateRowProps"
      fixed-header
      hide-default-footer
      v-bind="$attrs"
      @update:options="fetchDatatableFunc"
      @click:row="selectOrUnselectRow"
    >
      <template v-for="(_, slot) of $slots" #[slot]="scope">
        <slot
          :name="slot"
          v-bind="{
            ...scope,
            ...slotProps,
          }"
        />
      </template>
      <template v-for="column in editableColumns" #[`item.${column.value}`]="slotProps">
        <slot
          :name="`item.${column.value}`"
          v-bind="{
            ...slotProps,
          }"
        >
          <VText
            v-if="editable && shouldShowEditor(slotProps, selectedCell, column, focused)"
            :key="`${column.value as string}-edited-input`"
            v-model="selectedCell.value"
            :autofocus="true"
            :hide-details="true"
            :name="column.value as string"
            :rules="column.rules"
            :width="column.inputWidth"
            :type="column.type || 'text'"
            dense
            single-line
            @change:meta="(meta: any) => (selectedCell.meta = meta)"
            @focus="$event.target.select()"
            @keydown.enter="updateRowFunc && updateRowFunc(!selectedCell.meta.valid, slotProps.item, selectedCell)"
            @keydown.esc="closeEditor"
          />
          <div
            v-else
            :key="`${column.value as string}-view-input`"
            style="width: 100%; height: 100%"
            :data-editable-row-index="slotProps.index"
            class="editable-row-data d-flex"
          >
            <span style="margin: auto">{{ (slotProps.item || {})[column.value as string] }}</span>
          </div>
        </slot>
      </template>
      <template #no-data>
        <p>{{ $t("crud.no_data") }}</p>
      </template>
    </v-data-table-server>
    <v-row class="pt-2">
      <v-col>
        <slot name="left-paginate" />
      </v-col>
      <v-col>
        <v-pagination
          v-if="paginate"
          v-model="page"
          :length="pageCount"
          :style="pStyle"
          :total-visible="totalVisible"
        />
      </v-col>
    </v-row>
  </div>
</template>

<script setup lang="ts">
import type {
  FetchDatatableFunc,
  WrapperDataTableHeader,
  DataTableUpdateRowCallbackFunc,
  SelectedCellDataTable,
  DataOptions,
  DataTableUpdateRowFunc,
} from "~/types/datatable";
import type { DatatableTrigger } from "~/types/trigger";
import { cloneDeep } from "lodash-es";
import { $i18n } from "~/utils/i18n";
import { makeOptionFromResponse } from "~/utils/api";
import { makeUpdateRowDatatableFunc, scrollIntoView } from "~/utils";
import type { FetchOption } from "~/types/api";

defineOptions({
  inheritAttrs: false,
});

const props = defineProps({
  fetchDataFunc: {
    type: Function as PropType<FetchDatatableFunc>,
    required: true,
  },
  headers: {
    type: Array as PropType<WrapperDataTableHeader[]>,
    required: true,
  },
  initFetchExtraParams: {
    type: Object,
    default: () => ({}),
  },
  filterRules: {
    type: Object,
    default: () => ({}),
  },
  initItemPerPage: {
    type: Number,
    default: 25,
  },
  editable: {
    type: Boolean,
    default: false,
  },
  editUpdateRowCallback: {
    type: Function as PropType<DataTableUpdateRowCallbackFunc>,
    default: null,
  },
  editValidateData: {
    type: Function as PropType<(item: any) => boolean>,
    default: null,
  },
  editFetchOption: {
    type: Object as PropType<FetchOption>,
    default: () => ({}),
  },
  locale: {
    type: String,
    default: () => "en-US",
  },
  height: {
    type: String,
    default: () => undefined,
  },
  dStyle: {
    type: String,
    default: () => "",
  },
  pStyle: {
    type: String,
    default: () => "",
  },
  totalVisible: {
    type: Number,
    default: () => 5,
  },
  showSelect: {
    type: Boolean,
    default: () => true,
  },
  singleSelect: {
    type: Boolean,
    default: () => false,
  },
  multiSort: {
    type: Boolean,
    default: () => true,
  },
  itemValue: { type: String, default: "id" },
  searchDelay: {
    type: Number,
    default: 300,
  },
  tableClasses: {
    type: String,
    default: "",
  },
  paginate: {
    type: Boolean,
    default: true,
  },
  setTrigger: {
    type: Function as PropType<(_triggers: DatatableTrigger) => void>,
    default: null,
  },
  // selectedItems: {
  //   type: Array,
  //   default: () => []
  // }
});

const emits = defineEmits(["change:selected-rows"]);

// region init

// region for general
const itemKey = props.itemValue || "id";
const headers = toRef(props, "headers");
const items = ref([]);
const page = ref(1);
const loading = ref(false);
const extraData = ref({
  options: {},
} as any);

const totalItem = ref(0);

const itemsPerPage = ref(props.initItemPerPage || 25);

const pageCount = computed(() => (totalItem.value - 1) / itemsPerPage.value + 1);
const searchKeyword = ref("");

const options = ref<DataOptions>({
  page: 1,
  itemsPerPage: itemsPerPage.value,
  sortBy: [],
  sortDesc: [],
} as DataOptions);

// const selectedRows = useVModel(props, "selectedItems");
const selectedRows = ref([] as any[]);

const generateRowProps = (item: any) => {
  return {
    class: selectedRows.value.includes(item.internalItem.value) ? "v-data-table__selected" : "",
  };
};

const defaultSelectedCell: SelectedCellDataTable = {
  row: "",
  column: "",
  value: "",
  origin: "",
  item: {},
  meta: {},
  render: false,
} as SelectedCellDataTable;

const selectedCell = ref(cloneDeep(defaultSelectedCell));

const fetchExtraParams = ref<any>(props.initFetchExtraParams || ({} as any));

let searchJob: any = null;

// delayTime for refresh
const delayTime = ref(0);

// endregion

// region for editable
const focused = ref(false);
// endregion

// endregion

// region function declaration
// region for general
const selectOrUnselectRow = (_: any, row: any) => {
  row.toggleSelect(row.internalItem);
};

const resetSelectedCell = () => {
  selectedCell.value = cloneDeep(defaultSelectedCell);
};

const reloadTable = (delay = false) => {
  if (+new Date() / 1000 - delayTime.value < 0.5) {
    $notify.warning({
      title: $i18n.t("notification.warning_title").toString(),
      message: $i18n.t("crud.reload_to_fast").toString(),
    });
    return;
  }

  if (delay) {
    delayTime.value = +new Date() / 1000;
  }

  fetchDatatableFunc(options.value);
};

const clearSelection = () => {
  selectedRows.value = [];
  resetSelectedCell();
};

const clearSelectionAndReload = (delay = false) => {
  clearSelection();
  reloadTable(delay);
};

const selectAllItems = () => {
  nextTick(() => {
    if (selectedRows.value.length > 0) {
      selectedRows.value = [];
    } else {
      selectedRows.value.push(...items.value);
    }
  });
};

const updateSelectedCell = (row: any, column: any) => {
  selectedCell.value.row = row;
  selectedCell.value.column = column;
};

console.log(updateSelectedCell);

const shouldUseParams = (type: string, value: any) => {
  switch (type) {
    case FilterRuleType.BETWEEN:
      return value[0] && value[1];
    default:
      return value;
  }
};

const buildFetchParams = (): any => {
  const builtParams = {} as any;
  Object.keys(fetchExtraParams.value).map((k) => {
    if (props.filterRules[k]) {
      if (shouldUseParams(props.filterRules[k], fetchExtraParams.value[k])) {
        builtParams[k] = `::${props.filterRules[k]}(${JSON.stringify(fetchExtraParams.value[k])})`;
      }
    } else {
      builtParams[k] = fetchExtraParams.value[k];
    }
  });
  return builtParams;
};

const fetchDatatableFunc = (options: DataOptions) => {
  loading.value = true;
  props
    .fetchDataFunc({
      options,
      keyword: searchKeyword.value,
      headers: headers.value,
      params: buildFetchParams(),
    })
    .then((data: any) => {
      items.value = data.items;
      totalItem.value = data.total;
      const extra = cloneDeep(data.extra || {});
      extra.options = makeOptionFromResponse(data.options || {});
      extraData.value = extra;
    })
    .catch((err) => {
      console.error(err);
    })
    .finally(() => {
      loading.value = false;
    });
};

const exportData = (action: string) => {
  return props.fetchDataFunc({
    options: options.value,
    keyword: searchKeyword.value,
    headers: headers.value,
    params: buildFetchParams(),
    action: action,
  });
};
console.log(exportData);

// endregion

// region for editable
const closeEditor = () => {
  nextTick(() => {
    selectedCell.value.value = selectedCell.value.origin;
    nextTick(() => {
      selectedCell.value.render = false;
    });
  });
};
let updateRowFunc: DataTableUpdateRowFunc | undefined;
// eslint-disable-next-line @typescript-eslint/no-invalid-void-type
let changeActiveCell: (evt: any) => void | undefined;
if (props.editable) {
  updateRowFunc = makeUpdateRowDatatableFunc(props.editUpdateRowCallback, props.editFetchOption);

  changeActiveCell = (evt: Event) => {
    const tableWrapper = document.querySelector("div.editable-datatable");
    if (!tableWrapper) return;

    // Remove all "selectedCell" classes
    const allCells = tableWrapper.querySelectorAll("table > tbody > tr > td");
    allCells.forEach((cell) => cell.classList.remove("selectedCell"));

    // Find the closest TD from event target
    const activeCell = (evt.target as HTMLElement).closest("td");
    if (!activeCell) return;

    activeCell.classList.add("selectedCell");

    // Find data index
    const activeRow = activeCell.parentElement; // <tr>
    if (!activeRow) {
      resetSelectedCell();
      return;
    }

    const child = activeRow.querySelector("div.editable-row-data");
    if (!child) {
      resetSelectedCell();
      return;
    }

    const dataIndex = child.getAttribute("data-editable-row-index");
    if (!dataIndex || isNaN(+dataIndex) || items.value[+dataIndex] == undefined || items.value[+dataIndex] == null) {
      resetSelectedCell();
      return;
    }

    const data: any = items.value[+dataIndex];
    const row = data[itemKey];

    const cellIndex = Array.from(activeCell.parentElement.children).indexOf(activeCell);
    const column: string = headers.value[cellIndex].value as string;

    if (selectedCell.value.row !== row || selectedCell.value.column !== column) {
      selectedCell.value.row = row;
      selectedCell.value.column = column;
      selectedCell.value.value = data[column];
      selectedCell.value.origin = data[column];
      selectedCell.value.item = data;
      selectedCell.value.render = isMobileScreen();
    }
  };
}
// endregion

// endregion

// region watch changes

// region for general
watch(
  options,
  (c) => {
    fetchDatatableFunc(c);
  },
  { deep: true },
);

watch(itemsPerPage, (currentValue) => {
  options.value.itemsPerPage = currentValue;
  options.value.page = 1;
});

watch(
  fetchExtraParams,
  () => {
    if (searchJob !== null) {
      clearTimeout(searchJob);
    }

    searchJob = setTimeout(() => {
      options.value.page = 1;
      clearSelectionAndReload();
    }, 100);
  },
  {
    deep: true,
  },
);

watch(searchKeyword, (currentVal, oldVal) => {
  console.log("change", currentVal);
  if (currentVal !== oldVal) {
    if (searchJob !== null) {
      clearTimeout(searchJob);
    }

    searchJob = setTimeout(() => {
      options.value.page = 1;
      fetchDatatableFunc(options.value);
    }, props.searchDelay);
  }
});

watch(selectedRows, (c) => {
  emits("change:selected-rows", c);
});
// endregion

// region for editable
if (props.editable) {
  const computeSelectedCell = () => {
    return [
      selectedCell.value.render,
      selectedCell.value.item,
      selectedCell.value.row,
      selectedCell.value.column,
      selectedCell.value.value,
      selectedCell.value.origin,
      selectedCell.value.meta,
    ];
  };

  watch(computeSelectedCell, (n, o) => {
    if (o[0] && !n[0] && updateRowFunc) {
      const oldSelectedCell: SelectedCellDataTable = {
        row: o[2],
        column: o[3],
        value: o[4],
        origin: o[5],
        item: o[1],
        meta: o[6],
        render: false,
      };
      const isValid = !props.editValidateData || props.editValidateData(oldSelectedCell.item);
      updateRowFunc!(!isValid, oldSelectedCell.item, oldSelectedCell);
    }
  });

  document.addEventListener("DOMContentLoaded", () => {
    const tableWrapper = document.querySelector("div.editable-datatable");
    const body = document.body;
    const table = document.querySelector("div.editable-datatable table");

    if (!tableWrapper || !table) return;

    // Remove old click listeners then add new
    tableWrapper.addEventListener("click", (e) => {
      const target = e.target as HTMLElement;
      if (target && target.matches("table tbody tr td")) {
        changeActiveCell(e);
      }
    });

    // Custom event "changeCell"
    tableWrapper.addEventListener("changeCell", (e) => {
      const target = e.target as HTMLElement;
      if (target && target.matches("table tbody tr td")) {
        changeActiveCell(e);
      }
    });

    // KeyDown events
    body.addEventListener("keydown", (evt: KeyboardEvent) => {
      if (
        [
          KeyCode.DownArrow,
          KeyCode.UpArrow,
          KeyCode.LeftArrow,
          KeyCode.RightArrow,
          KeyCode.Enter,
          KeyCode.Escape,
        ].includes(evt.keyCode)
      ) {
        evt.preventDefault();

        const activeCell = document.querySelector("div.editable-datatable table td.selectedCell") as HTMLElement;

        if (activeCell) {
          let newCell: HTMLElement | null = null;
          const currentRow = activeCell.parentElement as HTMLTableRowElement;
          const activeCellIndex = Array.from(currentRow.children).indexOf(activeCell);
          const getSiblingCell = (row: HTMLTableRowElement | null, index: number) => {
            if (!row) return null;
            return row.children[index] as HTMLElement;
          };

          switch (evt.keyCode) {
            case KeyCode.DownArrow:
              newCell = getSiblingCell(
                activeCell.parentElement?.nextElementSibling as HTMLTableRowElement,
                activeCellIndex,
              );
              break;
            case KeyCode.UpArrow:
              newCell = getSiblingCell(
                activeCell.parentElement?.previousElementSibling as HTMLTableRowElement,
                activeCellIndex,
              );
              break;
            case KeyCode.RightArrow:
              newCell = activeCell.nextElementSibling as HTMLElement;
              break;
            case KeyCode.LeftArrow:
              newCell = activeCell.previousElementSibling as HTMLElement;
              break;
            case KeyCode.Enter:
              if (!selectedCell.value.render) {
                selectedCell.value.render = true;
              } else {
                newCell = getSiblingCell(
                  activeCell.parentElement?.nextElementSibling as HTMLTableRowElement,
                  activeCellIndex,
                );
              }
              break;
            case KeyCode.Escape:
              break;
          }

          if (newCell) {
            const event = new Event("changeCell", { bubbles: true });
            newCell.dispatchEvent(event);
            scrollIntoView(
              newCell,
              document.querySelector("div.editable-datatable > div.v-data-table__wrapper") as HTMLElement,
            );
          }
        }
      } else if (!selectedCell.value.render) {
        selectedCell.value.render = true;
      }
    });

    // Focus check on click
    window.addEventListener("click", (e) => {
      focused.value = e.target instanceof Node && table.contains(e.target as Node);
    });
  });
}
// endregion

// endregion

const triggers: DatatableTrigger = {
  clearSelection,
  reloadTable,
  clearSelectionAndReload,
  selectAllItems,
};

onMounted(() => {
  if (props.setTrigger) {
    props.setTrigger!(triggers);
  }
});

const dataTableClasses = computed((): string => {
  let classes = props.tableClasses;
  if (props.editable) {
    classes = classes + " editable-datatable";
  }
  return classes;
});

const editableColumns = computed((): WrapperDataTableHeader[] => {
  return props.headers.filter((header: WrapperDataTableHeader) => {
    return header.editable;
  });
});

const shouldShowEditor = (props: any, selectedCell: SelectedCellDataTable, column: any, focused: boolean): boolean => {
  return (
    (props.item || {})[props.itemKey] == selectedCell.row &&
    column.value === selectedCell.column &&
    selectedCell.render &&
    focused
  );
};

const slotProps = {
  items,
  totalItem,
  extraData,
  options,
  headers,
  itemsPerPage,
  selectedRows,
  selectedCell,
  loading,
  page,
  pageCount,
  searchKeyword,
  focused,
  selectOrUnselectRow,
  resetSelectedCell,
  fetchDatatableFunc,
  closeEditor,
  selectAllItems,
  clearSelectionAndReload,
  updateRowFunc,
};
</script>

<style scoped></style>
