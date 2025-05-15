import {i18n} from '~/plugins/i18n'
import {WrapperDataTableHeader} from "~/types";

const headerDataTable = (): WrapperDataTableHeader[] => [
    {
        text: i18n.t('models.user_role.fields.id').toString(),
        align: 'center',
        sortable: false ,
        columnKey: 'id',
        value: 'id',
        class: 'v-header',
        cellClass: '',
        show: true,
    },
    {
        text: i18n.t('models.user_role.fields.user_id').toString(),
        align: 'center',
        sortable: false,
        columnKey: 'userId',
        value: 'user.email',
        class: 'v-header',
        cellClass: '',
        show: true,
    },
    {
        text: i18n.t('models.user_role.fields.role_id').toString(),
        align: 'center',
        sortable: false,
        columnKey: 'roleId',
        value: 'role.name',
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
