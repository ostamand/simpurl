const configs = {
  //apiEndpoint: "http://localhost:3000/api", // for dev
  //apiEndpoint: "https://url.ostamand.com/api",
  apiEndpoint: null, // will get from host
};

export default function getConfigs() {
  if (!configs.apiEndpoint) {
    configs.apiEndpoint = window.location.origin + "/api";
  }
  return configs;
}
