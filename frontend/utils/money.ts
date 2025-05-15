export const moneyToNumber = (money: string, t = ","): number => {
  return parseInt(String(money).split(t).join(""));
};

export const numberToMoney = (number: number, c = 0): string => {
  return (number || 0).toFixed(c).replace(/(\d)(?=(\d{3})+(?:\.\d+)?$)/g, "$1,");
};

export const moneyFormat = numberToMoney;

// export const moneyToWords = (number: string | number): string => {
//   let money = number;
//   if (isString(number)) {
//     money = moneyToNumber(number)
//   }
//   return String(number);
// }
