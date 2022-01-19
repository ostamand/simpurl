export const formatURL = (url) => {
  return url.replace("https://", "").replace("http://", "").replace("www.", "");
};
