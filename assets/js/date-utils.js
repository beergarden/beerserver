/* @flow */

export function formatDateTime(date: Date): string {
  const year = date.getFullYear();
  const month = date.getMonth() + 1;
  const day = date.getDate();
  const hours = date.getHours();
  const minutes = padZero(date.getMinutes());
  return `${year}/${month}/${day} ${hours}:${minutes}`;
}

export function formatDate(date: Date): string {
  const month = date.getMonth() + 1;
  const day = date.getDate();
  return `${month}/${day}`;
}

export function formatTime(date: Date): string {
  const hours = date.getHours();
  const minutes = padZero(date.getMinutes());
  return `${hours}:${minutes}`;
}

function padZero(num: number): string {
  const str = num.toString();
  if (str.length === 1) {
    return `0${str}`;
  } else {
    return str;
  }
}
