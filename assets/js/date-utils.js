/* @flow */

export function formatDateTime(date: Date): string {
  var year = date.getFullYear();
  var month = date.getMonth() + 1;
  var day = date.getDate();
  var hours = date.getHours();
  var minutes = padZero(date.getMinutes());
  return `${year}/${month}/${day} ${hours}:${minutes}`;
}

export function formatDate(date: Date): string {
  var month = date.getMonth() + 1;
  var day = date.getDate();
  return `${month}/${day}`;
}

export function formatTime(date: Date): string {
  var hours = date.getHours();
  var minutes = padZero(date.getMinutes());
  return `${hours}:${minutes}`;
}

function padZero(num: number): string {
  var str = num.toString();
  if (str.length === 1) {
    return `0${str}`;
  } else {
    return str;
  }
}
