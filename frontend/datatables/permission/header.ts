import { i18n } from '~/plugins/i18n'
import {WrapperDataTableHeader} from "~/types";

const headerDataTable = (): WrapperDataTableHeader[] => [
  {
  text: i18n.t('models.permission.fields.id').toString(),
  align: 'center',
  sortable: false ,
  columnKey: 'id',
  value: 'id',
  class: 'v-header',
  cellClass: '',
  show: true,
},
  {
  text: i18n.t('models.permission.fields.name').toString(),
  align: 'center',
  sortable: false ,
  columnKey: 'name',
  value: 'name',
  class: 'v-header',
  cellClass: '',
  show: true,
},
  {
  text: i18n.t('models.permission.fields.method').toString(),
  align: 'center',
  sortable: false ,
  columnKey: 'method',
  value: 'method',
  class: 'v-header',
  cellClass: '',
  show: true,
},
  {
  text: i18n.t('models.permission.fields.url').toString(),
  align: 'center',
  sortable: false ,
  columnKey: 'url',
  value: 'url',
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