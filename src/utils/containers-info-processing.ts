/**
 * Transform date value to string value
 * If the value only has one digit, append a 0
 * @param value
 */
function dateValueToString(value: number): string {
  let valueToString = value.toString();
  if (valueToString.length < 2) valueToString = "0" + valueToString;
  return valueToString;
}

/**
 * Transform UNIX Timestamp into readable date
 * @param timestamp - Unix timestamp
 */
function timestampToDate(timestamp: number): string {
  const date = new Date(timestamp * 1000);
  const formattedDate = [
    date.getFullYear(),
    date.getMonth() + 1,
    date.getDate(),
  ];
  const formattedDateToString = formattedDate.map((item) =>
    dateValueToString(item)
  );
  const formattedTime = [date.getHours(), date.getMinutes(), date.getSeconds()];
  const formattedTimeToString = formattedTime.map((item) =>
    dateValueToString(item)
  );

  return (
    formattedDateToString.join("-") + " " + formattedTimeToString.join(":")
  );
}

export { timestampToDate };
