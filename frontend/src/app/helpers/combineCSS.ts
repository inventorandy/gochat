/**
 * Combines a base CSS class with additional specified classes.
 * @param base Base CSS class - always returned
 * @param additions Optional additional CSS classes - only returned when specified
 * @returns combined string of base + additions
 */
const combineCSS = (base: string, additions?: string) => {
  if (additions && additions !== "") {
    return base + ' ' + additions;
  }
  return base;
}

export default combineCSS;