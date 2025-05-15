const CommonService = {
  async fetchLanguages(locale: string) {
    return await $api(`/lang/${locale}.json`)
      .catch(() => {
        return { data: [] };
      });
  },
};

export default CommonService;
