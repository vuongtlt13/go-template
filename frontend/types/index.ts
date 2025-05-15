export type StringMap<T> = {
  [key: string]: T;
};

export type FormSubmitFunc = (values: object) => Promise<any>;
