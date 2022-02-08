import getConfigs from "../defaults.js";

const apiEndpoint = getConfigs().apiEndpoint;

export default class FetchWrapper {
  constructor(url = apiEndpoint) {
    this.url = url;
  }

  async _sendRequest(endpoint, method, body = null) {
    const options = {
      method: method,
      credentials: "include",
      headers: {
        "Content-Type": "application/json",
      },
    };
    if (body) {
      options["body"] = JSON.stringify(body);
    }
    const response = await fetch(this.url + endpoint, options);
    let data = {};
    try {
      data = await response.json();
    } catch (error) {
      console.error(error);
    }
    return [response.status, data];
  }

  post(endpoint, body) {
    return this._sendRequest(endpoint, "POST", body);
  }

  patch(endpoint, body) {
    return this._sendRequest(endpoint, "PATCH", body);
  }

  get(endpoint) {
    return this._sendRequest(endpoint, "GET");
  }
}
