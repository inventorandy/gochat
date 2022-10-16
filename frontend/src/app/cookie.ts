/**
 * Utilities for managing cookies
 */

export const setCookie = (name: string, value: string, expires: Date) => {
  // Set the Cookie
  document.cookie =
    name + '=' + value + '; expires=' + expires.toUTCString() + '; path=/';
};

export const getCookie = (name: string): string | undefined => {
  const value = '; ' + document.cookie;
  const parts = value.split('; ' + name + '=');
  if (parts.length === 2) {
    return parts.pop()?.split(';').shift();
  }
};

export const deleteCookie = (name: string) => {
  const date = new Date();

  // Set it expire in -1 days
  date.setTime(date.getTime() + -1 * 24 * 60 * 60 * 1000);

  // Set it
  document.cookie = name + '=; expires=' + date.toUTCString() + '; path=/';
};
