import { i18n } from '~/plugins/i18n'
import {WrapperDataTableHeader} from "~/types";

const headerDataTable = (): WrapperDataTableHeader[] => [
  {
  text: i18n.t('models.role.fields.id').toString(),
  align: 'center',
  sortable: false ,
  columnKey: 'id',
  value: 'id',
  class: 'v-header',
  cellClass: '',
  show: true,
},
  {
  text: i18n.t('models.role.fields.name').toString(),
  align: 'center',
  sortable: false ,
  columnKey: 'name',
  value: 'name',
  class: 'v-header',
  cellClass: '',
  show: true,
},
  {
  text: i18n.t('models.role.fields.code').toString(),
  align: 'center',
  sortable: false ,
  columnKey: 'code',
  value: 'code',
  class: 'v-header',
  cellClass: '',
  show: true,
},
  {
    text: i18n.t('crud.action').toString(),
    value: 'actions',
    sortable: false,
    class: 'v-header',
    align: 'center',
    show: true,
  }
]

export default headerDataTable