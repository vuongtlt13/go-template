export type RuleCheckFunction = (curRules: string[], targetRules: string[]) => boolean;

export const RequireAll: RuleCheckFunction = (curRules: string[], targetRules: string[]) => {
  return targetRules.every((rule) => curRules.includes(rule));
};

export const RequireOne: RuleCheckFunction = (curRules: string[], targetRules: string[]) => {
  return targetRules.some((rule) => curRules.includes(rule));
};
