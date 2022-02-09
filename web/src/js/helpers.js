const sessionToken = "session";

export function formatURL(url) {
  return url.replace("https://", "").replace("http://", "").replace("www.", "");
}

/**
 * Assume that the user is logged in if the session cookie is found
 * @returns
 */
export function isLoggedIn() {
  const cookies = document.cookie.split(";").map((text) => {
    return text.split("=")[0].trim();
  });
  return (
    cookies.includes(sessionToken) && cookies.includes(`${sessionToken}.sig`)
  );
}
