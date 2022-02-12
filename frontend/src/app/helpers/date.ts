import { DateTime } from 'luxon';

/**
 * Converts a message timestamp into a cool format (like "Just Now", or "15:36" or "12-July-2017 14:52")
 * @param time
 * @returns
 */
export const TimeSinceMessage = (time: string) => {
  // Create a DateTime from the input type
  const msgDate = DateTime.fromISO(time);

  // Create the current DateTime
  const curDate = DateTime.now();

  // Determine the length of time since the message and send an appropriate text response
  if (msgDate.toSeconds() > curDate.toSeconds() - 60) {
    return 'Just Now';
  } else if (
    msgDate.day === curDate.day &&
    msgDate.month === curDate.month &&
    msgDate.year === curDate.year
  ) {
    return msgDate.toFormat('HH:mm');
  } else {
    return msgDate.toFormat('dd-LLL-yyyy HH:mm');
  }
};
