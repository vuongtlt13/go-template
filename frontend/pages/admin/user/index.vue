<template>
  <v-card elevation="2" rounded="lg" style="height: 100%">
    <DataTable
      height="77vh"
      :show-select="false"
      :fetch-data-func="UserService.fetchUser"
      :headers="adminHeader()"
      :editable="false"
      :set-trigger="setUserDatatableTrigger"
      d-style="border-radius: inherit;"
      p-style=""
      @dblclick:row="userUpdateModalTrigger.showEditItemByDoubleClick"
      @change:selected-rows="updateSelectedRows"
    >
    </DataTable>

    <!--    region create modal-->
    <!--    <FormModal-->
    <!--        width="90vw"-->
    <!--        max-width="100vw"-->
    <!--        v-model="showUserCreateDialog"-->
    <!--        :default-form-data="defaultFormData"-->
    <!--        :form-title="$t('crud.add_new_modal_title', { model: $t('models.user.singular') })"-->
    <!--        :handle-submit-fn="UserService.addNewUser"-->
    <!--        :load-option-fn="UserService.loadCreateUserOption"-->
    <!--        :success-callback-option="{reloadTableFn: userDatatableTrigger.clearSelectionAndReload}"-->
    <!--    >-->
    <!--      <template #fields="{initData, options}">-->
    <!--        <UserFields :init-data="initData" :options="options"/>-->
    <!--      </template>-->
    <!--    </FormModal>-->
    <!--    endregion create modal-->

    <!--    region update modal-->
    <!--    <BaseFormModalWithKey-->
    <!--        width="90vw"-->
    <!--        max-width="100vw"-->
    <!--        v-model="showUserUpdateDialog"-->
    <!--        :default-form-data="defaultFormData"-->
    <!--        :form-title="$t('crud.add_new_modal_title', { model: $t('models.user.singular') })"-->
    <!--        :load-option-fn="UserService.loadUpdateUserOption"-->
    <!--        :handle-submit-fn="UserService.updateUser"-->
    <!--        :set-trigger="setUserUpdateModalTrigger"-->
    <!--        :success-callback-option="{reloadTableFn: userDatatableTrigger.clearSelectionAndReload}"-->
    <!--    >-->
    <!--      <template #fields="{initData, options}">-->
    <!--        <UserFields :init-data="initData" :options="options"/>-->
    <!--      </template>-->
    <!--    </BaseFormModalWithKey>-->
    <!--    endregion update modal-->
  </v-card>
</template>

<script setup lang="ts">
import { adminHeader } from "@/datatables/user/header";
import UserService from "@/services/admin/user";
import { generateTitle } from "@/utils";
import { useTrigger } from "~/composables/useTrigger";
import type { DatatableTrigger, UpdateModalTrigger } from "~/types/trigger";

useHead({
  title: generateTitle($i18n.t("models.user.singular").toString()),
});

definePageMeta({
  layout: "admin",
  middleware: ["auth"],
  meta: {
    rp: ["get_list_api_crud_user__get", "get_list_api_crud_permission__get"],
  },
});

// region User

const defaultFormData = {
  id: undefined,
  email: undefined,
  password: undefined,
  fullName: undefined,
  isActive: true,
  isAdmin: false,
  createdAt: undefined,
  updatedAt: undefined,
};

// region datatable init
const { triggers: userDatatableTrigger, setTriggers: setUserDatatableTrigger } = useTrigger<DatatableTrigger>();
// endregion

// region create modal
const showUserCreateDialog = ref(false);
// endregion

// region update modal
const showUserUpdateDialog = ref(false);
const { triggers: userUpdateModalTrigger, setTriggers: setUserUpdateModalTrigger } = useTrigger<UpdateModalTrigger>({
  showEditItemByDoubleClick: () => {},
});
// endregion

// region delete confirm
const selectedRows = ref([] as any[]);
const updateSelectedRows = (rows: any[]) => {
  selectedRows.value = rows;
};
// const userDeleteConfirm = useConfirmDelete({
//   deleteRecordFn: UserService.deleteUser,
//   deleteRecordsFn: UserService.deleteUsers,
//   selectedRows: selectedRows,
//   successCallbackOption: {
//     reloadTableFn: () =>
//       userDatatableTrigger.value.clearSelectionAndReload && userDatatableTrigger.value.clearSelectionAndReload(),
//   },
// });
// endregion

// endregion
</script>

<style scoped></style>
