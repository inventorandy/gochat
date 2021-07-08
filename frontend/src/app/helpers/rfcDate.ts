const shortMonths = [
  'Jan',
  'Feb',
  'Mar',
  'Apr',
  'May',
  'Jun',
  'Jul',
  'Aug',
  'Sep',
  'Oct',
  'Nov',
  'Dec'
  ];
  
  /**
   * Converts a datetime string in RFC3339 format to dd-mmm-yy hh:mm
   * @param datetime Datetime string in RFC3339 Format (2006-05-04T01:02:03Z)
   * @returns dd-mmm-yy hh:mm
   */
  export const rfc3339ToSystemDate = (datetime: string | undefined): string => {
    if (datetime === undefined) return "";
    
    // Convert to a Date Object
    const dateObj = new Date(datetime);
    
    // Format the Day
    let day: string = dateObj.getDate().toString();
    if (dateObj.getDate() < 10) {
      day = "0" + dateObj.getDate().toString();
    }
    
    // Format the Month
    let month: string = shortMonths[dateObj.getMonth()];
    
    // Format the Hours
    let hours: string = dateObj.getHours().toString();
    if (dateObj.getHours() < 10) {
      hours = "0" + dateObj.getHours().toString();
    }
    
    // Format the Minutes
    let mins: string = dateObj.getMinutes().toString();
    if (dateObj.getMinutes() < 10) {
      mins = "0" + dateObj.getMinutes().toString();
    }
    
    // Return the String
    return day + "-" + month + " " + dateObj.getFullYear().toString() + " " + hours + ":" + mins;
  }