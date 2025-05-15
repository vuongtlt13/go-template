import type { CommonTrigger } from "~/types/trigger";

export const useTrigger = <T>(initValues?: T): CommonTrigger<T> => {
  // @ts-expect-error ignore this error
  const triggers: Ref<T> = ref(initValues || ({} as T));

  const setTriggers = (_triggers: T) => {
    triggers.value = {
      ...triggers.value,
      ..._triggers,
    };
  };

  return {
    triggers,
    setTriggers,
  };
};
