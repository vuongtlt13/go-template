import { configure, defineRule } from "vee-validate";
import { all } from "@vee-validate/rules";
import * as _ from "lodash-es";

interface FieldValidationMetaInfo {
  field: string;
  name: string;
  label?: string;
  value: unknown;
  form: Record<string, unknown>;
  rule?: {
    name: string;
    params?: Record<string, unknown> | unknown[];
  };
}

const renderParams = (params: any) => {
  if (_.isArray(params)) {
    const finalParams = {} as any;
    for (const index in params) {
      finalParams[`params${index}`] = params[index];
    }

    return finalParams;
  }

  return {
    ...params,
  };
};

export default defineNuxtPlugin((_) => {
  Object.entries(all).forEach(([name, rule]) => {
    defineRule(name, rule);
  });

  configure({
    generateMessage: (ctx: FieldValidationMetaInfo) => {
      if (ctx.rule) {
        return $i18n.t(`validations.${ctx.rule!.name}`, {
          attribute: $i18n.t(ctx.label || ctx.field),
          ...renderParams(ctx.rule!.params),
        });
      } else {
        return $i18n.t("validations.field_is_invalid", { field: ctx.field });
      }
    },
  });
});
