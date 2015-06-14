export function formatDateTime(date) {
  const year = date.getFullYear();
  const month = date.getMonth() + 1;
  const day = date.getDate();
  const hours = date.getHours();
  const minutes = padZero(date.getMinutes());
  return `${year}/${month}/${day} ${hours}:${minutes}`;
}

export function formatDate(date) {
  const month = date.getMonth() + 1;
  const day = date.getDate();
  return `${month}/${day}`;
}

export function formatTime(date) {
  const hours = date.getHours();
  const minutes = padZero(date.getMinutes());
  return `${hours}:${minutes}`;
}

function padZero(num) {
  const str = num.toString();
  if (str.length === 1) {
    return `0${str}`;
  } else {
    return str;
  }
}
