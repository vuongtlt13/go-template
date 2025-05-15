import type { WrapperDataTableHeader } from "~/types/datatable";

export const adminHeader = (): WrapperDataTableHeader[] => [
  {
    title: $i18n.t("models.user.fields.id").toString(),
    align: "center",
    sortable: false,
    columnKey: "id",
    value: "id",
    key: "id",
    editable: true,
    headerProps: {
      class: "v-header",
    },
    cellProps: {
      class: "",
    },
    show: true,
  },
  {
    title: $i18n.t("models.user.fields.email").toString(),
    align: "center",
    sortable: false,
    columnKey: "email",
    value: "email",
    key: "email",
    editable: true,
    headerProps: {
      class: "v-header",
    },
    cellProps: {
      class: "",
    },
    show: true,
  },
  {
    title: $i18n.t("models.user.fields.full_name").toString(),
    align: "center",
    sortable: false,
    columnKey: "fullName",
    value: "fullName",
    key: "fullName",
    headerProps: {
      class: "v-header",
    },
    cellProps: {
      class: "",
    },
    show: true,
  },
  {
    title: $i18n.t("models.user.fields.is_active").toString(),
    align: "center",
    sortable: false,
    columnKey: "isActive",
    value: "isActive",
    key: "v-header",
    headerProps: {
      class: "v-header",
    },
    cellProps: {
      class: "",
    },
    show: true,
  },
  {
    title: $i18n.t("models.user.fields.is_admin").toString(),
    align: "center",
    sortable: false,
    columnKey: "isAdmin",
    value: "isAdmin",
    key: "isAdmin",
    headerProps: {
      class: "v-header",
    },
    cellProps: {
      class: "",
    },
    show: true,
  },
  {
    title: $i18n.t("crud.action").toString(),
    value: "actions",
    key: "actions",
    sortable: false,
    headerProps: {
      class: "v-header",
    },
    cellProps: {
      class: "",
    },
    align: "center",
    show: true,
  },
];
