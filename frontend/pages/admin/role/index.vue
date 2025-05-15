<template>
  <v-container>
    <h1>Quản lý Roles</h1>
    <BaseDataTable
      :headers="headers"
      :items="roles"
      :loading="loading"
      :fetch-data-func="fetchRoles"
      :editable="true"
      @edit="editItem"
      @delete="deleteItem"
    >
      <template #top>
        <v-toolbar flat>
          <v-toolbar-title>Danh sách Roles</v-toolbar-title>
          <v-divider class="mx-4" inset vertical></v-divider>
          <v-spacer></v-spacer>
          <v-btn color="primary" dark @click="openDialog()">Thêm Role</v-btn>
        </v-toolbar>
      </template>
    </BaseDataTable>

    <v-dialog v-model="dialog" max-width="500px">
      <v-card>
        <v-card-title>
          <span class="text-h5">{{ formTitle }}</span>
        </v-card-title>
        <v-card-text>
          <v-container>
            <v-row>
              <v-col cols="12" sm="6" md="4">
                <v-text-field
                  :model-value="editedItem.name"
                  label="Tên Role"
                  @update:model-value="(val) => (editedItem.name = val)"
                ></v-text-field>
              </v-col>
              <v-col cols="12" sm="6" md="4">
                <v-text-field
                  :model-value="editedItem.code"
                  label="Mã Role"
                  @update:model-value="(val) => (editedItem.code = val)"
                ></v-text-field>
              </v-col>
            </v-row>
          </v-container>
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="blue darken-1" variant="text" @click="closeDialog">Hủy</v-btn>
          <v-btn color="blue darken-1" variant="text" @click="save">Lưu</v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
  </v-container>
</template>

<script setup lang="ts">
import { ref, onMounted, computed } from "vue";
import BaseDataTable from "~/components/base/datatable/DataTable.vue";
import type { WrapperDataTableHeader, FetchDatatableFunc } from "~/types/datatable";

interface Role {
  id: number;
  name: string;
  code: string;
}

const roles = ref<Role[]>([]);
const loading = ref<boolean>(false);
const dialog = ref<boolean>(false);
const editedIndex = ref<number>(-1);
const editedItem = ref<Partial<Role>>({
  name: "",
  code: "",
});
const defaultItem: Partial<Role> = {
  name: "",
  code: "",
};

const headers: WrapperDataTableHeader[] = [
  { title: "ID", key: "id" },
  { title: "Tên Role", key: "name" },
  { title: "Mã Role", key: "code" },
];

const formTitle = computed<string>(() => {
  return editedIndex.value === -1 ? "Thêm Role" : "Sửa Role";
});

const fetchRoles: FetchDatatableFunc = async (options) => {
  loading.value = true;
  try {
    const response = await fetch("/api/roles");
    if (!response.ok) throw new Error("Failed to fetch roles");
    const data = await response.json();
    roles.value = data;
    return {
      items: data,
      total: data.length,
    };
  } catch (error) {
    console.error("Error fetching roles:", error);
    return {
      items: [],
      total: 0,
    };
  } finally {
    loading.value = false;
  }
};

const editItem = (item: Role) => {
  editedIndex.value = roles.value.indexOf(item);
  editedItem.value = Object.assign({}, item);
  dialog.value = true;
};

const deleteItem = async (item: Role) => {
  const index = roles.value.indexOf(item);
  if (confirm("Bạn có chắc chắn muốn xóa role này?")) {
    try {
      const response = await fetch(`/api/roles/${item.id}`, {
        method: "DELETE",
      });
      if (!response.ok) throw new Error("Failed to delete role");
      roles.value.splice(index, 1);
    } catch (error) {
      console.error("Error deleting role:", error);
    }
  }
};

const closeDialog = () => {
  dialog.value = false;
  editedItem.value = Object.assign({}, defaultItem);
  editedIndex.value = -1;
};

const save = async () => {
  try {
    if (editedIndex.value > -1) {
      const response = await fetch(`/api/roles/${editedItem.value.id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(editedItem.value),
      });
      if (!response.ok) throw new Error("Failed to update role");
      const data = await response.json();
      Object.assign(roles.value[editedIndex.value], data);
    } else {
      const response = await fetch("/api/roles", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(editedItem.value),
      });
      if (!response.ok) throw new Error("Failed to create role");
      const data = await response.json();
      roles.value.push(data);
    }
    closeDialog();
  } catch (error) {
    console.error("Error saving role:", error);
  }
};

const openDialog = () => {
  editedItem.value = Object.assign({}, defaultItem);
  editedIndex.value = -1;
  dialog.value = true;
};

onMounted(() => {
  fetchRoles({
    options: { page: 1, itemsPerPage: 25, sortBy: [], sortDesc: [] },
    keyword: null,
    headers: headers,
  });
});
</script>
