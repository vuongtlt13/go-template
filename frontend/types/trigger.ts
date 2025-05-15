import type { Ref } from "vue";

export interface CommonTrigger<T> {
  triggers: Ref<T>;
  setTriggers: (_triggers: T) => void;
}

export interface DatatableTrigger {
  clearSelection?: () => void;
  reloadTable?: (delay?: boolean) => void;
  clearSelectionAndReload?: (delay?: boolean) => void;
  selectAllItems?: () => void;
}

export interface UpdateModalTrigger {
  showEditItem?: (item: any) => void;
  showEditItemByDoubleClick?: (_: any, { item }: any) => void;
}
