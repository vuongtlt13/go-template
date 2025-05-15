import {i18n} from '~/plugins/i18n'
import {DataTableHeader} from 'vuetify';

const headerDataTable = (): DataTableHeader[] => [
  {
    text: i18n.t('models.role_permission.fields.id').toString(),
    align: 'center',
    sortable: false,
    value: 'id',
    class: 'v-header',
    cellClass: ''
  },
  // {
  //   text: i18n.t('models.role_permission.fields.role_id').toString(),
  //   align: 'center',
  //   sortable: false,
  //   value: 'roleId',
  //   class: 'v-header',
  //   cellClass: ''
  // },
  {
    text: i18n.t('models.permission.fields.name').toString(),
    align: 'center',
    sortable: false,
    value: 'permission.name',
    class: 'v-header',
    cellClass: ''
  },
  {
    text: i18n.t('models.permission.fields.method').toString(),
    align: 'center',
    sortable: false,
    value: 'permission.method',
    class: 'v-header',
    cellClass: ''
  },
  {
    text: i18n.t('models.permission.fields.url').toString(),
    align: 'center',
    sortable: false,
    value: 'permission.url',
    class: 'v-header',
    cellClass: ''
  },
  {
    text: i18n.t('crud.action').toString(),
    value: 'actions',
    sortable: false,
    class: 'v-header',
    align: 'center'
  }
]

export default headerDataTable
