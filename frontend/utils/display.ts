export const isMobileScreen = function (width: number = 960) {
  if (import.meta.client) {
    return screen.width <= width
  }

  return false;
}
