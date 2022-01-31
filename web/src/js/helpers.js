const sessionToken = "session_token";

export const formatURL = (url) => {
  return url.replace("https://", "").replace("http://", "").replace("www.", "");
};

export const getSessionToken = () => {
  let session = "";
  document.cookie.split(";").forEach((cookie) => {
    const [key, value] = cookie.trim().split("=");
    if (key === sessionToken) {
      session = value;
    }
  });
  return session;
};

export const clearSessionToken = () => {
  document.cookie = `${sessionToken}=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/`;
};
