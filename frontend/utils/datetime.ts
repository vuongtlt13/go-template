import dayjs from "dayjs";

// import locales
import "dayjs/locale/vi";
import "dayjs/locale/en";

// import plugins
import customParseFormat from "dayjs/plugin/customParseFormat";

export const VNDateFormat = "DD/MM/YYYY";
export const VNDashDateFormat = "DD-MM-YYYY";
export const NoDashDateFormat = "YYYYMMDD";
export const VNDateTimeFormat = "DD/MM/YYYY HH:mm:ss";
export const VNDashDateTimeFormat = "DD-MM-YYYY HH:mm:ss";

dayjs.locale("en"); // use locale
dayjs.extend(customParseFormat);

export const now = () => {
  return dayjs();
};

export const setDateTimeLocale = (locale: string) => {
  dayjs.locale(locale);
};

export const standardDatetime = (datetimeStr: string, format: string | null = null) => {
  if (format === null) {
    return dayjs(datetimeStr);
  }
  return dayjs(datetimeStr, format);
};

export const diffTimeFromNow = (dtObj: dayjs.Dayjs): string => {
  const currentTime = now();
  const diffInSeconds = currentTime.diff(dtObj, "seconds");
  if (diffInSeconds <= 10) {
    return $i18n.tc("time.just_now");
  }

  if (diffInSeconds < 60) {
    return $i18n.tc("time.x_seconds_ago", diffInSeconds);
  }

  const diffInMinutes = currentTime.diff(dtObj, "minutes");
  if (diffInMinutes < 60) {
    return $i18n.tc("time.x_minutes_ago", diffInMinutes);
  }

  const diffInHours = currentTime.diff(dtObj, "hours");
  if (diffInHours < 24) {
    return $i18n.tc("time.x_hours_ago", diffInHours);
  }

  const diffInDays = currentTime.diff(dtObj, "days");
  if (diffInDays < 31) {
    return $i18n.tc("time.x_days_ago", diffInDays);
  }

  const diffInMonths = currentTime.diff(dtObj, "months");
  if (diffInMonths < 12) {
    return $i18n.tc("time.x_months_ago", diffInMonths);
  }

  const diffInYears = currentTime.diff(dtObj, "years");
  return $i18n.tc("time.x_years_ago", diffInYears);
};

export const formatDate = (date: Date | dayjs.Dayjs, format: string) => {
  const dateObj = dayjs(date);

  return dateObj.format(format);
};

export const parseDate = (dateStr: string, format: string) => {
  return dayjs(dateStr, format);
};

export const formatDateTime = (date: Date | dayjs.Dayjs, format: string) => {
  const dateObj = dayjs(date);

  return dateObj.format(format);
};

export const parseDateTime = (dateStr: string, format: string) => {
  return dayjs(dateStr, format);
};
