import { MenuType } from "~/utils/constants";
import type { RuleCheckFunction } from "~/utils/rule-check";

export interface Menu {
  type: string;
  title: string;
  icon?: string;
  to?: string | object;
  activeClass?: string;
  class?: string;
  tooltip?: string;
  menuClass?: string;
  children?: Menu[];
  permissions?: string[];
  permissionRuleCheckFn?: RuleCheckFunction;
  roles?: string[];
  roleRuleCheckFn?: RuleCheckFunction;
}

const menuComponents: Menu[] = [
  {
    type: MenuType.DIVIDER,
    title: "",
    roles: ["super_admin"],
  },
  {
    type: MenuType.SUB_HEADER,
    title: "Super Admin",
    roles: ["super_admin"],
  },
  // {
  //   type: MenuType.NO_CHILD,
  //   title: $i18n.t('models.user.menu_title').toString(),
  //   icon: 'mdi-account', // default: 'mdi-checkbox-blank-circle'
  //   to: {name: 'user'},
  //   activeClass: 'light-blue lighten-4 text--accent-4',
  //   children: [],
  //   roles: [
  //     'super_admin'
  //   ]
  // },
  // {
  //   type: MenuType.NO_CHILD,
  //   title: $i18n.t('models.user_role.menu_title').toString(),
  //   icon: 'mdi-security', // default: 'mdi-checkbox-blank-circle'
  //   to: {name: 'user-role'},
  //   activeClass: 'light-blue lighten-4 text--accent-4',
  //   children: [],
  //   roles: [
  //     'super_admin'
  //   ]
  // },
  // {
  //   type: MenuType.NO_CHILD,
  //   title: $i18n.t('models.role.menu_title').toString(),
  //   icon: 'mdi-security', // default: 'mdi-checkbox-blank-circle'
  //   to: {name: 'role'},
  //   activeClass: 'light-blue lighten-4 text--accent-4',
  //   children: [],
  //   roles: [
  //     'super_admin'
  //   ]
  // },
  // {
  //   type: MenuType.NO_CHILD,
  //   title: $i18n.t('models.permission.menu_title').toString(),
  //   icon: 'mdi-security', // default: 'mdi-checkbox-blank-circle'
  //   to: {name: 'permission'},
  //   activeClass: 'light-blue lighten-4 text--accent-4',
  //   children: [],
  //   roles: [
  //     'super_admin'
  //   ]
  // },
  // {
  //   type: MenuType.NO_CHILD,
  //   title: $i18n.t('models.role_permission.menu_title').toString(),
  //   icon: 'mdi-security', // default: 'mdi-checkbox-blank-circle'
  //   to: {name: 'role-permission'},
  //   activeClass: 'light-blue lighten-4 text--accent-4',
  //   children: [],
  //   roles: [
  //     'super_admin'
  //   ]
  // }
];

export default menuComponents;
